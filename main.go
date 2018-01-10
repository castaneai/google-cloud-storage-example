// Sample storage-quickstart creates a Google Cloud Storage bucket.
package main

import (
	"fmt"
	"log"

	// Imports the Google Cloud Storage client package.
	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
	"io/ioutil"
)

const (
	ProjectID  = "morning-tide"
	BucketName = ProjectID + "-test-bucket"
	Region     = "asia-northeast1"
)

func createBucket() {
	ctx := context.Background()

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name for the new bucket.

	// Creates a Bucket instance.
	bucket := client.Bucket(BucketName)
	attrs := &storage.BucketAttrs{Location: Region}

	// Creates the new bucket.
	if err := bucket.Create(ctx, ProjectID, attrs); err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
	}

	fmt.Printf("Bucket %v created.\n", BucketName)
}

func uploadFile(bucketName, objectName string, bytes []byte) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	bucket := client.Bucket(bucketName)
	obj := bucket.Object(objectName)
	wc := obj.NewWriter(ctx)
	wc.ContentType = "image/jpeg"
	if _, err := wc.Write(bytes); err != nil {
		log.Fatalf("Failed to write object: %s (bucket: %s)", objectName, bucketName)
	}
	if err := wc.Close(); err != nil {
		log.Fatalf("Failed to close object: %s (bucket: %s)", objectName, bucketName)
	}
	// make public
	if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		log.Fatalf("Failed to set ACL, %v", err)
	}

	publicURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, objectName)
	log.Printf("object created: %s\n", publicURL)
}

func main() {
	path := "C:\\Users\\phone\\Desktop\\ed07b2d2d33317dd64e8d60ea2318fd5.jpg"
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read %s", path)
	}
	uploadFile(BucketName, "test-public.jpg", b)
}
