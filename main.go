package main

import (
	"context"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func main() {
	f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = PutObject(f)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("done")
}

func PutObject(file *os.File) error {
	credentialFilePath := "./key.json"
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		return err
	}
	bucketName := "BUCKET-NAME"
	objectPath := "OBJECT-PATH"
	object := client.Bucket(bucketName).Object(objectPath)
	writer := object.NewWriter(ctx)
	writer.ObjectAttrs.ContentType = "text/plain"
	writer.ObjectAttrs.CacheControl = "no-cache"
	writer.ObjectAttrs.ACL = []storage.ACLRule{
		storage.ACLRule{
			Entity: storage.AllUsers,
			Role:   storage.RoleReader,
		},
	}
	defer writer.Close()
	if _, err := io.Copy(writer, file); err != nil {
		return err
	}
	return err
}
