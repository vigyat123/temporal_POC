package helloworld

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

// Workflow is a Hello World workflow definition.
func S3Workflow(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("HelloWorld workflow started")

	var result string
	err := workflow.ExecuteActivity(ctx, S3Activity, name).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	logger.Info("HelloWorld workflow completed.", "result", result)

	return result, nil
}

func S3Activity(ctx context.Context, name string) (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	}
	logger := activity.GetLogger(ctx)
	logger.Info("Activity")
	listBucketInput := &s3.ListBucketsInput{}

	s3Client := s3.NewFromConfig(cfg)
	output, err := s3Client.ListBuckets(ctx, listBucketInput)
	if err != nil {
		return "buckets not found", err
	}
	//createBucketInput := &s3.CreateBucketInput{
	//	Bucket: aws.String("testBucket123"),
	//}
	//_, err := client.CreateBucket(ctx, createBucketInput)
	//if err != nil {
	//	return "bucket creation failed", err
	//}
	fmt.Println(*output.Buckets[4].Name)
	return "buckets found", nil
}
