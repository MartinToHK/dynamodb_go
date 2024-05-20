package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/MartinToHK/dynamodb_go/internal/repository/adapter"
	"github.com/MartinToHK/dynamodb_go/internal/repository/instance"
	"github.com/MartinToHK/dynamodb_go/internal/routes"
	"github.com/MartinToHK/dynamodb_go/internal/rules"
	"github.com/MartinToHK/dynamodb_go/internal/rules/product"
	"github.com/MartinToHK/dynamodb_go/utils/logger"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main(){
	config := Config.GetConfig()
	connection := instance.GetConnection()
	repository := adapter.NewAdapter(connection)

	logger.INFO("waiting for the service to start ......", nil)
	errors := Migrate(connection)
	if len(errors) > 0 {
		for _, err := range errors{
			logger.PANIC("Error on migration.....", err)
		}
	}
	logger.PANIC("", checkTables(connection))

	port := fmt.Sprintf(":%v", config.Port)
	router := routes.NewRouter().SetRouters(repository)
	logger.INFO("service is running on port", port)

	server := http.ListenAndServe(port, router)
	log.Fatal(server)

}

func Migrate(connection *dynamodb.DynamoDB) []error{
	var errors []error
	callMigrateAndAppendError(&errors, connection, &RulesProduct.Rules{})
	return errors


}

func callMigrateAndAppendError(errors *[]error, connection *dynamodb.DynamoDB, rule rules.Interface){
	err := rule.Migrate(connection)
	if err!= nil{
		*errors = apped=nd(*errors, err)
	}
}

func checkTables(connection *dynamodb.DynamoDB) error{
	response, err := connection.ListTables(&dynamodb.ListTables(&dynamodb.ListTablesInput{}))

	if (response) != nil{
		if len(response.TableNames) == 0 {
			logger.INFO("Table not found:", nil)

		}
		for _, tableName := range response.TableNames{
			logger.INFO("Table found:", *tableName)
		}
	}
	return err


}