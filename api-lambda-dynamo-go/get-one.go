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
from boto3.dynamodb.conditions import Key

def main(event, context):
    try:
        dateString = event['pathParameters']['dateString']
    except:
        return{
            'statusCode': 200,
            'body': "unable to read event path parameters"
        }

    dynamodb = boto3.resource('dynamodb')
    table = dynamodb.Table('StockData')

    response = table.query(
        KeyConditionExpression=Key('dateString').eq(dateString)
    )

    openAvg = 0
    highAvg = 0
    lowAvg = 0
    closeAvg = 0
    volAvg = 0

    for x in response['Items']:

        print("\n")
        stockObj = x['stock_data']
        if not openAvg:
            openAvg = float(list(stockObj.items())[0][1]['1. open'])
        if not highAvg:
            highAvg = float(list(stockObj.items())[0][1]['1. open'])
        if not lowAvg:
            lowAvg = float(list(stockObj.items())[0][1]['1. open'])
        if not closeAvg:
            closeAvg = float(list(stockObj.items())[0][1]['1. open'])
        if not volAvg:
            volAvg = float(list(stockObj.items())[0][1]['1. open'])

        for key, value in stockObj.items():
            openAvg = (openAvg + float(value['1. open'])) / 2
            highAvg = (highAvg + float(value['2. high']))/2
            lowAvg = (lowAvg + float(value['3. low']))/2
            closeAvg = (closeAvg + float(value['4. close']))/2
            volAvg = (volAvg + float(value['5. volume']))/2
            print("\n")

    print("Open:   {:.2f}".format(openAvg))
    print("High:   {:.2f}".format(highAvg))
    print("Low:    {:.2f}".format(lowAvg))
    print("Close:  {:.2f}".format(closeAvg))
    print("Volume: {:.2f}".format(volAvg))

    dayStats = json.dumps({
            'Opening Average':   "{:.2f}".format(openAvg),
            'High  Average':   "{:.2f}".format(highAvg),
            'Low Average':    "{:.2f}".format(lowAvg),
            'Close Average':  "{:.2f}".format(closeAvg),
            'Volume Average': "{:.2f}".format(volAvg)
            })

    return {
        'statusCode': 200,
        'body': dayStats
    }
*/


func getOneFunction(){
	
}
