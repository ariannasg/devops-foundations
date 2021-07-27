// File: application.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	sparta "github.com/mweagle/Sparta"
)

func paramVal(keyName string, defaultValue string) string {
	value := os.Getenv(keyName)
	if "" == value {
		value = defaultValue
	}
	return value
}

var s3Bucket = paramVal("S3_TEST_BUCKET", "arn:aws:s3:::howdy-world")

////////////////////////////////////////////////////////////////////////////////
// Hello world event handler
//
func helloWorld(event *json.RawMessage,
	context *sparta.LambdaContext,
	w http.ResponseWriter,
	logger *logrus.Logger) {
	logger.Info("raw msg: ", string(*event))
	fmt.Fprint(w, "Howdy, Infrastructure as Code Class\n\n\n"+string(*event))
}

func appendS3Lambda(api *sparta.API, lambdaFunctions []*sparta.LambdaAWSInfo) []*sparta.LambdaAWSInfo {
	options := new(sparta.LambdaFunctionOptions)
	options.Timeout = 30
	lambdaFn := sparta.NewLambda(sparta.IAMRoleDefinition{}, helloWorld, options)
	apiGatewayResource, _ := api.NewResource("/greeting", lambdaFn)
	apiGatewayResource.NewMethod("GET", http.StatusOK)

	lambdaFn.Permissions = append(lambdaFn.Permissions, sparta.S3Permission{
		BasePermission: sparta.BasePermission{
			SourceArn: s3Bucket,
		},
		Events: []string{"s3:ObjectCreated:*", "s3:ObjectRemoved:*"},
	})
	return append(lambdaFunctions, lambdaFn)
}

func spartaLambdaData(api *sparta.API) []*sparta.LambdaAWSInfo {

	var lambdaFunctions []*sparta.LambdaAWSInfo
	lambdaFunctions = appendS3Lambda(api, lambdaFunctions)
	return lambdaFunctions
}

////////////////////////////////////////////////////////////////////////////////
// Main
func main() {
	stage := sparta.NewStage("prod")
	apiGateway := sparta.NewAPIGateway("howdy", stage)
	apiGateway.CORSEnabled = true

	//lambda info
	os.Setenv("AWS_PROFILE", "sparta")
	os.Setenv("AWS_REGION", "us-east-1")

	// Deploy it
	sparta.Main("HowdyWorld",
		"Saying howdy through the API Gateway using lambdas",
		spartaLambdaData(apiGateway),
		apiGateway,
		nil)
}
