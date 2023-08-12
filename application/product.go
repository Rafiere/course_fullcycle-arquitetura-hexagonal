package application

import (
	"errors"
	uuid "github.com/satori/go.uuid"
)
import "github.com/asaskevich/govalidator"

func init() {

	/* Esse é um pacote para tornar os pacotes "required" por padrão. */
	govalidator.SetFieldsRequiredByDefault(true)
}

/* Nesse arquivo, teremos o produto. Isso faz parte do coração da aplicação. É como
se fosse a "entidade". */

/* As regras de negócio ficarão dentro da entidade. */

/* Estamos definindo os métodos que o "Product" deverá implementar. */
type ProductInterface interface {

	/* Se o erro for vazio, é que deu tudo certo. */

	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

/* Nesse exemplo, o "service" servirá para salvarmos no banco de dados. */
type ProductServiceInterface interface {
	Get(id string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

/*
	Caso queiramos implementar a interface do "Reader", basta termos um método

"get" implementando uma "ProductInterface".

	Se quisermos implementar a interface do "Writer", basta termos um método

"create" implementando essa interface.
*/
type ProductReader interface {
	Get(id string) (ProductInterface, error)
}

type ProductWriter interface {
	Save(product ProductInterface) (ProductInterface, error)
}

// Podemos trabalhar com o service e definir que ele implementará tanto o
//método "get" ou "save".

// Sempre trabalharemos com interfaces.
type ProductPersistenceInterface interface {
	ProductReader
	ProductWriter
}

/* Todas as vezes que setarmos um status do produto, utilizaremos essa constante. */
const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

/* No Go, não temos classes, e sim "structs". */

/*
	A "struct" não precisa implementar uma interface. Somente dela ter os métodos, ela

implementa a interface.
*/
type Product struct {

	/* Tags são anotações. Estamos adicionando as tags do "GoValidator" para realizar
	as validações. */

	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required"`
}

/*
	É recomendado criarmos uma função que retorne um ponteiro para o "Product". Ela

inicializará os valores de um "Product".
*/
func NewProduct() *Product {
	product := Product{
		ID:     uuid.NewV4().String(),
		Status: DISABLED,
	}

	return &product
}

/* Esse é um método da struct. */
func (p *Product) IsValid() (bool, error) {

	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("The status must be enabled or disabled")
	}

	if p.Price < 0 {
		return false, errors.New("The price must be greater or equal zero")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("The price must be greater than zero to enable the product.")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}
	return errors.New("The price must be zero to disable the product.")
}

func (p *Product) GetId() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
