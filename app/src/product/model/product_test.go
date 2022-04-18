package model_test

import (
	"loja/app/src/product/model"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := model.Product{}
	product.Name = "Hello"
	product.Status = model.DISABLE
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()

	require.Equal(t, "O preço do produto não pode ser menor que zero para ser ativado", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := model.Product{}
	product.Name = "Hello"
	product.Status = model.ENABLE
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()

	require.Equal(t, "O preço do produto de ser zero para ser desativado", err.Error())
}

func TestProduct_IsValid(t *testing.T){
	product := model.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = model.DISABLE
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t,err)

	product.Status = "INVALID"

	_, err = product.IsValid()

	
	require.Equal(t, "O status po ser apenas "+ model.ENABLE + " ou " + model.DISABLE, err.Error())

	product.Status = model.ENABLE
	_, err = product.IsValid()
	require.Nil(t,err)

	product.Price = -10
	_, err = product.IsValid()

	require.Equal(t, "O preço não pode ser menor que zero", err.Error())

}
