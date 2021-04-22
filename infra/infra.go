package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awslogs"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

type InfraStackProps struct {
	awscdk.StackProps
}

func NewInfraStack(scope constructs.Construct, id string, props *InfraStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// The code that defines your stack goes here
	api := awsapigateway.NewRestApi(stack, jsii.String("api"), &awsapigateway.RestApiProps{
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowMethods: awsapigateway.Cors_ALL_METHODS(),
			AllowOrigins: awsapigateway.Cors_ALL_ORIGINS(),
		},
	})
	var lambda awslambda.IFunction
	lambda = awslambda.NewFunction(stack, jsii.String("SampleFunc"), &awslambda.FunctionProps{
		Timeout:      awscdk.Duration_Millis(jsii.Number(60.0 * 1000)),
		FunctionName: jsii.String("SampleFile"),
		Code:         awslambda.AssetCode_Asset(jsii.String("../bin")),
		Handler:      jsii.String("sampleFunc"),
		Runtime:      awslambda.Runtime_GO_1_X(),
		LogRetention: awslogs.RetentionDays_FIVE_DAYS,
		MemorySize:   jsii.Number(256),
	})
	mapping := awsapigateway.NewLambdaIntegration(lambda, &awsapigateway.LambdaIntegrationOptions{})
	api.Root().AddMethod(jsii.String("GET"), mapping, api.Root().DefaultMethodOptions())
	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewInfraStack(app, "InfraStack", &InfraStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
