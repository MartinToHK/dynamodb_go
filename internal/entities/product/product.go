package product

import(
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
	"errors"
	"time"
)

type Product struct{
	entities.Base
	Name string `json:"name"`
}

func InterfaceToModel(data interface{})(instance *Product, err error){
	bytes, err := json.Marshal(data)
	if err != nil{
		return instance, err
	}
	return instance, json.Unmarshal(bytes, &instance)
}

func (p *Product) GetFilterId() map[string]interface{}{

}

func (p *Product) TableName() string {
	return "product"
}

func (p *Product) Bytes() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Product) GetMap() map[string]interface{}{

}

func ParseDynamoAttributeToStruct()(){
	
}