package floating_client

import (
	"gitlab.com/l3178/sdk-go/floating_client"
)

func GetLicense() {
	client := InitializeClient()

	// create registration request
	product := floating_client.ProductRequest{
		Product: "<your_product>",
	}

	// unregister user
	// this is also done automatically as configured in floating server
	client.GetLicense(product)
}

func Ping() {
	client := InitializeClient()
	client.Ping()
}
