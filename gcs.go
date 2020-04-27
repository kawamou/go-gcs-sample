package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type Gcs struct {
	bucketName string
	objectpath string
	object     *storage.ObjectHandle
}

func NewGcs(bucketName, objectPath, credentialFilePath string) *Gcs {
	gcs := new(Gcs)
	gcs.bucketName = bucketName
	gcs.objectpath = objectPath
	ctx := context.Background()
	client, err := storage.NewClient(ctx,
		option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		log.Fatal(err)
	}
	gcs.object = client.Bucket(bucketName).Object(objectPath)
	return gcs
}

func (gcs *Gcs) PutObject(file *os.File) error {
	ctx := context.Background()
	writer := gcs.object.NewWriter(ctx)
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
		log.Fatal(err)
	}
	return nil
}

func (gcs *Gcs) GetObject() ([]byte, error) {
	ctx := context.Background()
	reader, err := gcs.object.NewReader(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}
