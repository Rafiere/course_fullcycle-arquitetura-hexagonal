package application

/* Trabalharemos com o conceito de injeção de dependências, ou seja, não faremos
que o "service" utilize uma persistência em específico. Ele apenas deve implementar
essa interface. */

type ProductService struct {
	Persistence ProductPersistenceInterface
}

/*
	Não sabemos o banco de dados, sabemos apenas que temos um banco

de dados que buscará um serviço.
*/
func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistence.Get(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {

	product := NewProduct()
	product.Name = name
	product.Price = price
	_, err := product.IsValid()
	if err != nil {
		return &Product{}, err
	}

	result, err := s.Persistence.Save(product)

	if err != nil {
		return &Product{}, err
	}
	return result, nil
}
