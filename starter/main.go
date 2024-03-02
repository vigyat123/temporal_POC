package main

import (
	"context"
	"log"
	s3Actions "temporal_POC"

	"go.temporal.io/sdk/client"
)

func main() {
	//cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	//if err != nil {
	//	log.Fatal(err)
	//}

	//cfg.Region = "us-east-1"
	//s3Client := s3.NewFromConfig(cfg)
	//listBucketInput := &s3.ListBucketsInput{}
	//output, err := s3Client.ListBuckets(context.Background(), listBucketInput)
	//if err != nil {
	//	log.Fatalln("problem with s3", err)
	//}
	//fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!")
	//fmt.Println(*output.Buckets[4].Name)
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "hello_world_workflowID",
		TaskQueue: "hello-world",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, s3Actions.S3Workflow, "temporal")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	// Synchronously wait for the workflow completion.
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
