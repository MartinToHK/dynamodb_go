package product

import()

type Controller struct{
	repository adapter.Interface

}

type Interface interface{
	ListOne(id uuid.UUID) (entity product.Product, err error)
	ListAll()(entites []product.Product, err error)
	Create(entity product.Product)(uuid.UUID, error)
	Update(ID uuid.UUID, entity *product.Product) error
	Remove(ID uuid.UUID) error

}

func NewControler(repository adapter.Interface) Interface{
	return &Controller(repository:repository)

}

func (c *Controller) ListOne(id uuid.UUID) (entity product.Product, err error){
	entity.ID = id
	response, err := c.repository.FindOne(entity.GetFilterId(), entity.TableName())

	if err != nil{
		return entity, err
	}
	return product.ParseDynamoAttributeToStruct(response.Item)

}

func (c *Controller) ListAll()(entites []product.Product, err error){

}

func (c *Controller) Create(entity product.Product)(uuid.UUID, error){

}

func (c *Controller) Update(id uuid.UUID, entity *product.Product) {

}

func (c *Controller) Remove(id uuid.UUID) error{

}
