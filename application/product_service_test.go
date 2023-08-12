package application

import (
	"github.com/golang/mock/gomock"
	mock_application "github.com/rafiere/course_fullcycle-arquitetura-hexagonal/application/mocks"
	"testing"
)

/* Um teste de unidade não depende de objetos externos. */

/* Um mock permite que falsifiquemos um objeto externo, que será utilizado por uma
struct, para controlarmos o comportamento, partindo do princípio que achamos que
isso não vai afetar os testes. */

/* Uma stub é uma camada de código feita para trabalharmos com os mocks. */

func TestProductService_Get(t *testing.T) {

	ctrl := gomock.NewController(t)

	/* O "defer" espera tudo acontecer dentro do método para executar. */

	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	/* Sempre que o método "Get" for chamado, um produto fake será retornado. */
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()
}
