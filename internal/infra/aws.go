package infra

import "os"

func IsRunningOnLambda() bool {
	_, exists := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME")
	return exists
}

func EndpointURL() string {
	if value, ok := os.LookupEnv("AWS_ENDPOINT_URL"); ok {
		return value
	}
	return "http://localhost:4566"
}
