package adapter

import(
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/dynamodb/expression"

)

type Database struct{
	connection *dynamodb.DynamoDB
	logMode bool
}

type Interface interface{
	Health() bool
	FindAll(condition expression.Expression, tableName string)(response *dynamodb.ScanOutPut, err error)
	FindOne(condition map[string]interface{}, tableName string)(response *dynamodb.GetItemOutput, err error)
	CreateOrUpdate(entity interface{}, tableName string)(response *dynamodb.PutItemOutput, err error)
	Delete(condition map[string]interface{},tableName string)(response *dynamodb.DeleteItemOutput, err error)

}

func NewAdapter( con *dynamodb.DynamoDB) Interface{
	return &Database{
		connection : con,
		logMode : false,
	}

}

func (db *Database) Health() bool{
	_, err := db.connection.ListTables(&dynamodb.ListTablesInput{})
	return err == nil

}

func (db *Database) FindAll(condition expression.Expression, tableName string)(response *dynamodb.ScanOutPut, err error){
	input := &dynamodb.ScanInput{
		ExpressionAttributeNames: condition.Names(),
		ExpressionAttributeValue: condition.Values(),
		FilerExpression:condition.Filter(),
		ProjectionExpression:condition.Projection(),
		TableName : aws.string(tableName),
	}

}

func (db *Database) FindOne(condition map[String]interface{}, tableName string)(response *dynamodb.GetItemOutput, err error){

	conditionParsed, err:= dynamodbattribute.MarshalMap(condition)
	if err !=nil{
		return nil, err
	}
	input:= &dynamodb.GetItemInput{
		TableName : aws.String(tableName),
		key : conditionParsed,
	}
	return db.connection.GetItem(input)
}

func (db *Database) CreateOrUpdate( entity interface{}, tableName string)(response *dynamodb.PutItemOutput, err error){

	entityParsed, err := dynamodbattribute.MarshalMap(entity)
	if err !=nil{
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item: entityParsed,
		TableName: aws.String(tableName),
	}
	return db.connection.PutItem(input)
}

func (db *Database) Delete(condition map[String]interface{}, tableName string)(response *dynamodb.DeleteItemOutput, err error){

	conditionParsed, err:= dynamodbattribute.MarshalMap(condition)
	if err != nil {
		return nil, err
	}


	input := &dynamodb.DeleteItemInput{
		key : conditionParsed,
		TableName: aws.String(tableName),
	}
	return db.connection.DeleteItem(input)
}
