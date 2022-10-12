package main

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type ApiLambdaDynamoGoStackProps struct {
	awscdk.StackProps
}

func NewApiLambdaDynamoGoStack(scope constructs.Construct, id string, props *ApiLambdaDynamoGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	//define stock data dynamo table
	dynamoTable := dynamodb.NewTable(stack, "StockData", &dynamodb.TableProps{
		PartitionKey: dynamodb.NewAttributeDefinition("dateString", dynamodb.AttributeTypeString),
		SortKey:      dynamodb.NewAttributeDefinition("timeString", dynamodb.AttributeTypeString),
		TableName:    "StockData",
		RemovalPolicy: policies.DESTROY,
	})

	//define create one lambda function
	createOneLambda := lambda.NewFunction(stack, "createOneFunction", &lambda.FunctionProps{
		FunctionName: "createOneFunction",
		Code: lambda.NewInlineCode(
			fs.ReadDirFS("./lambdas/create-one.go", &fs.ReadDirOptions{Encoding: fs.EncodingUTF8}),
		),
		Handler: "index.main",
		Timeout: awscdk.Duration(300),
		Runtime: lambda.RuntimeGo1X,
		MemorySize: 128,
		Environment: &lambda.Environment{
			Variables: map[string]string{
				"PRIMARY_KEY": "dateString",
				"SORT_KEY":    "timeString",
				"TABLE_NAME":  "StockData",
			},
		},
	})

	//define get one lambda function
	getOneLambda := lambda.NewFunction(stack, "getOneFunction", &lambda.FunctionProps{
		FunctionName: "getOneFunction",
		Code: lambda.NewInlineCode(
			fs.ReadDirFS("./lambdas/get-one.go", &fs.ReadDirOptions{Encoding: fs.EncodingUTF8}),
		),
		Handler: "index.main",
		Timeout: awscdk.Duration(300),
		Runtime: lambda.RuntimeGo1X,
		MemorySize: 128,
		Environment: &lambda.Environment{
			Variables: map[string]string{
				"PRIMARY_KEY": "dateString",
				"SORT_KEY":    "timeString",
				"TABLE_NAME":  "StockData",
			},
		},
	})

	//define get all lambda function
	getAllLambda := lambda.NewFunction(stack, "getAllFunction", &lambda.FunctionProps{
		FunctionName: "getAllFunction",
		Code: lambda.NewInlineCode(
			fs.ReadDirFS("./lambdas/get-all.go", &fs.ReadDirOptions{Encoding: fs.EncodingUTF8}),
		),
		Handler: "index.main",
		Timeout: awscdk.Duration(300),
		Runtime: lambda.RuntimeGo1X,
		MemorySize: 128,
		Environment: &lambda.Environment{
			Variables: map[string]string{
				"PRIMARY_KEY": "dateString",
				"SORT_KEY":    "timeString",
				"TABLE_NAME":  "StockData",
			},
		},
	})

	//grant read/write data permission to lambda functions
	dynamoTable.GrantReadWriteData(getOneLambda)
	dynamoTable.GrantReadWriteData(createOneLambda)
	dynamoTable.GrantReadWriteData(getAllLambda)

	
	//define integration for lambda functions
	createOneIntegration := lambda.NewLambdaIntegration(createOneLambda)
	getOneIntegration := lambda.NewLambdaIntegration(getOneLambda)
	getAllIntegration := lambda.NewLambdaIntegration(getAllLambda)


	//define api gateway
	api := awscdk.NewRestAPI(stack, "stockDataApi", &awscdk.RestAPIConfig{
		Name: "stockDataApi",
	})

	//define resources for api gateway
	stockData := api.Root.AddResource("StockData")
	stockData.AddMethod("POST", createOneIntegration)
	addCorsOptions(stockData)

	singleDateString := stockData.AddResource("{dateString}")
	singleDateString.AddMethod("GET", getOneIntegration)
	addCorsOptions(singleDateString)

	allDateStrings := stockData.AddResource("all")
	allDateStrings.AddMethod("GET", getAllIntegration)
	addCorsOptions(allDateStrings)
	

	return stack
}

	//define addCorsOptions function
	func addCorsOptions(apiResource *awscdk.Resource) {
		apiResource.AddMethod("OPTIONS", &awscdk.MockIntegration{
			IntegrationResponses: []awscdk.IntegrationResponse{
				{
					StatusCode: "200",
					ResponseParameters: map[string]string{
						"method.response.header.Access-Control-Allow-Headers": "'Content-Type,X-Amz-date,Authorization,X-Api-Key,X-Amz-Security-Token,X-Amz-User-Agent'",
						"method.response.header.Access-Control-Allow-Origin": "'*'",
						"method.response.header.Access-Control-Allow-Methods": "'OPTIONS,GET,PUT,POST,DELETE'",
					},
				},
			},
			PassthroughBehavior: awscdk.PassthroughBehavior.NEVER,
			RequestTemplates: map[string]string{
				"application/json": `{"statusCode": 200}`,
			},
		}, &awscdk.MethodOptions{
			MethodResponses: []awscdk.MethodResponse{
				{
					StatusCode: "200",
					ResponseParameters: map[string]bool{
						"method.response.header.Access-Control-Allow-Headers": true,
						"method.response.header.Access-Control-Allow-Methods": true,
						"method.response.header.Access-Control-Allow-Origin": true,
					},
				},
			},
		}
		)
	}



func main() {
	app := awscdk.NewApp(nil)

	NewApiLambdaDynamoGoStack(app, "ApiLambdaDynamoGoStack", &ApiLambdaDynamoGoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
