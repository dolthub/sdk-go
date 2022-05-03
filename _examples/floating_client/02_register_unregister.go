package floating_client

import (
	"gitlab.com/l3178/sdk/sdk-go/v1/floating_client"
)

func Register() {
	client := InitializeClient()

	// create registration request
	registration := floating_client.FloatingRegistrationRequest{
		User:    "user@company.com",
		Product: "<your_product>",
	}

	// register user
	client.Register(registration)
}

func Unregister() {
	client := InitializeClient()

	// create registration request
	registration := floating_client.FloatingRegistrationRequest{
		User:    "user@company.com",
		Product: "<your_product>",
	}

	// unregister user
	// this is also done automatically as configured in floating server
	client.Unregister(registration)
}
