package main

	import (
		"io/fs"
		"os"
		"path"
		"path/filepath"
		"strings"
		"github.com/aws/aws-cdk/awscdk/v2"
		"github.com/aws/aws-cdk/awscdk/v2/awssqs"
		"github.com/aws/constructs-go/constructs/v10"	
		"github.com/aws/jsii-runtime-go"
	)

	func main (event, context) {
		dynamodb := boto3.resource('dynamodb')
		table := dynamodb.Table('StockData')
		response := table.scan()
		return {
			"statusCode": 200,
			"body": json.dumps(response['Items'])
		}
	}
	