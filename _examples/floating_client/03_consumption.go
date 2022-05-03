package floating_client

import (
	"gitlab.com/l3178/sdk/sdk-go/v1/floating_client"
)

func Consumption() {
	client := InitializeClient()

	consumption := floating_client.ProductRequest{
		Product: "<your_product>",
	}

	// add consumption
	client.AddConsumption(consumption)
}
