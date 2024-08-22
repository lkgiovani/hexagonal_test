package app_test

import (
	"github.com/lkgiovani/hexagonal_test/app"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := app.Product{}
	product.Name = "hello"
	product.Status = app.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to be enable the product", err.Error())

}

func TestProduct_Disable(t *testing.T) {

	product := app.Product{}
	product.Name = "hello"
	product.Status = app.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disable", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := app.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "hello"
	product.Status = app.DISABLED
	product.Price = 10.0

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = app.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

func TestProduct_GetName(t *testing.T) {

	product := app.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "hello"
	product.Status = app.DISABLED
	product.Price = 10.0

	id := product.GetId()
	name := product.GetName()
	status := product.GetStatus()
	price := product.GetPrice()

	require.Equal(t, id, product.ID)
	require.Equal(t, name, product.Name)
	require.Equal(t, status, product.Status)
	require.Equal(t, price, product.Price)
}
