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


/*
import requests
import boto3
import json


def main(event, context):
    try:
        body = json.loads(event['body'])
    except:
        return{
            'statusCode': 200,
            'body': "unable to read event body"
        }

    dynamodb = boto3.resource('dynamodb')
    dateString = list(body['Meta Data'].values())[2].split(" ")[0]
    timeString = list(body['Meta Data'].values())[2].split(" ")[1]
    print(dateString)
    print(timeString)

    table = dynamodb.Table('StockData')

    table.put_item(
        Item={
            'dateString': dateString,
            'timeString': timeString,
            'stock_data': body['Time Series (1min)']
        }
    )

    return {
        "statusCode": 200,
        'body': "Successfully created item in table"
    }

	*/


func main (event, context) {
	var body map[string]interface{}
	var err error
	if err = json.Unmarshal([]byte(event["body"]), &body); err != nil {
		return
	}
	dynamodb := boto3.resource('dynamodb')
	dateString := list(body["Meta Data"].values())[2].split(" ")[0]
	timeString := list(body["Meta Data"].values())[2].split(" ")[1]
	print(dateString)
	print(timeString)

	table := dynamodb.Table('StockData')

	table.put_item(
		Item={
			'dateString': dateString,
			'timeString': timeString,
			'stock_data': body["Time Series (1min)"]
		}
	)

	return {
		"statusCode": 200,
		'body': "Successfully created item in table"
	}
