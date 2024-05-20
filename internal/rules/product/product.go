package product

import(
	"encoding/json"
	"errors"
	"strings"
	"io"
	"time"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	Validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/MartinToHK/dynamodb_go/internal/entities"
	"github.com/MartinToHK/dynamodb_go/internal/entities/product"
)

type Rules struct{}

func NewRules() *Rules{
	return &Rules{}
}

func (r *Rules) ConvertIoReaderToStruct (data io.Reader, model interface{})(interface{}, error){
	if data == nil{
		return nil, errors.New("body is invalid")
	}
	return model, json.NewDecoder(data).Decode(model)

}

func (r *Rules) GetMock() interface{}{
	return product.Product{
		Base: entities.Base{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		Name: uuid.New().String()
	}

}

func (r *Rules) Migrate(connection *dynamodb.DynamoDB) error{
	return r.createTable(connection)

}

func (r *Rules) Validate(model interface{}) error{
	productModel, err := product.InterfaceToModel(model)
	if err != nil{
		return err
	}
	return Validate.ValidateStruct(productModel, Validation.Field(&productModel.ID, Validation.Required, is.UUIDv4), Validation.Field(&productModel.Name, Validation.Required,Validation.Length(3,50)),)

}

func (r *Rules) createTable() {

}