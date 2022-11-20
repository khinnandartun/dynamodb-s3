package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/jsii-runtime-go"
)

type DynamodbS3StackProps struct {
	awscdk.StackProps
}

func NewDynamodbS3Stack(scope constructs.Construct, id string, props *DynamodbS3StackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)
	//create s3 bucket
	awss3.NewBucket(stack, jsii.String("MyBucket"), &awss3.BucketProps{
		Versioned: jsii.Bool(false),
		RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
		AutoDeleteObjects: jsii.Bool(true),
	})

	//create dynamoDB
	awsdynamodb.NewTable(stack, jsii.String("dynamodb-table"),
		&awsdynamodb.TableProps{
			PartitionKey: &awsdynamodb.Attribute{
				Name: jsii.String("id"),
				Type: awsdynamodb.AttributeType_STRING},
			RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
			// BillingMode: awsdynamodb.billingMode_PAY_PER_REQUEST
		})

	// The code that defines your stack goes here

	// example resource
	// queue := awssqs.NewQueue(stack, jsii.String("DynamodbS3Queue"), &awssqs.QueueProps{
	// 	VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewDynamodbS3Stack(app, "DynamodbS3Stack", &DynamodbS3StackProps{
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
	//return &awscdk.Environment{
	// Account: jsii.String("926233248796"),
	// Region:  jsii.String("ap-southeast-1"),
	//}

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
	 Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	 Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
