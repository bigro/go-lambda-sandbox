package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bigro/go-lambda-sandbox/greeting"
)
func executeFunction()  {
	greeting.SayHello()
}

func main() {
	lambda.Start(executeFunction)
}
