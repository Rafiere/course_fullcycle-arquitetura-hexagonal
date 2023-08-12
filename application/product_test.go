package application_test

import (
	"github.com/rafiere/course_fullcycle-arquitetura-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

/* Caso queiramos testar as funções internas do pacote, temos que utilizar o pacote
com o mesmo nome. Caso queiramos testar de forma externa, podemos colocar o package
"application_test. */

/* Vamos testar a struct "Product" com o método "Enable". */
func TestProduct_Enable(t *testing.T) {

	product := application.Product{}

	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	/* Estamos verificando que o "err" deve ter recebido um erro de
	"nil". */
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()

	require.Equal(t, "The price must be greater than zero to enable the product.", err)
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}

	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()

	/* Estamos verificando que o "err" deve ter recebido um erro de
	"nil". */
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "The price must be zero to disable the product.", err)
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}

	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "The status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "The price must be greater or equal zero", err.Error())
}
