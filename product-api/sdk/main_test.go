package main

import (
	"fmt"
	"testing"

	"mado/sdk/client"
	"mado/sdk/client/products"
)

func TestOutClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(prod.GetPayload()[0])

}
