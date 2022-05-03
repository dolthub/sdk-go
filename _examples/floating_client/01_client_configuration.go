package floating_client

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/l3178/sdk-go/floating_client"
)

func Configuration() {
	// url to your floating server
	floatingServerUrl := "<url>"

	// initialize config
	config := floating_client.NewFloatingClientConfiguration(floatingServerUrl)

	// we can further setup our config, or we can use the default values

	// setup logging
	config.LoggingLevel(logrus.DebugLevel)
	// disable multiple retries on failed requests
	config.RetryCount = 0
	// set request timeout
	config.RequestTimeout = 10
}

func InitializeClient() *floating_client.FloatingClient {
	config := floating_client.NewFloatingClientConfiguration("<url>")

	// initialize client
	return floating_client.NewFloatingClient(config)
}
