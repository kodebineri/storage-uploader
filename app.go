package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/storage"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

// const STORAGE_BUCKET = "kodebineri-web.appspot.com"
// const CREDENTIAL_PATH = "kodebineri-web-firebase-adminsdk-xcq5v-7a27152afe.json"

func main() {
	fmt.Println("Initiating...")
	var config *firebase.Config
	if os.Getenv("STORAGE_BUCKET") == "" {
		log.Fatalln("Error: No STORAGE_BUCKET found!")
	}
	config = &firebase.Config{
		StorageBucket: os.Getenv("STORAGE_BUCKET"),
	}

	var opt option.ClientOption
	if os.Getenv("CREDENTIAL_JSON") != "" {
		opt = option.WithCredentialsJSON([]byte(os.Getenv("CREDENTIAL_JSON")))
	} else if os.Getenv("CREDENTIAL_PATH") != "" {
		opt = option.WithCredentialsFile(os.Getenv("CREDENTIAL_PATH"))
	} else {
		log.Fatalln("Error: No CREDENTIAL_JSON or CREDENTIAL_PATH found!")
	}

	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	_, err = client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}

	if len(os.Args) < 2 {
		log.Fatalln("Error: file not found!")
	}

	uploadFile(os.Getenv("STORAGE_BUCKET"), os.Args[1], opt)
}

func uploadFile(bucket, object string, opt option.ClientOption) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, opt)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	// Open local file.
	fmt.Println("Opening local file...")

	f, err := os.Open(object)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("os.Open: %v", err)
	}
	defer f.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	filenames := strings.Split(object, "/")
	filename := filenames[len(filenames)-1]

	wc := client.Bucket(bucket).Object(filename).NewWriter(ctx)
	if _, err = io.Copy(wc, f); err != nil {
		fmt.Println(err)
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		fmt.Println(err)
		return fmt.Errorf("Writer.Close: %v", err)
	}
	fmt.Sprintf("Blob %v uploaded.\n", object)
	return nil
}
