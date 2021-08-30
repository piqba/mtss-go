package main

import (
	"context"
	"fmt"
	"log"

	mtssgo "github.com/piqba/mtss-go"
)

func main() {
	baseURL := mtssgo.URL_BASE
	skipVerify := true
	client := mtssgo.NewClient(
		baseURL,
		skipVerify,
		nil, //custom http client, defaults to http.DefaultClient
		nil, //io.Writer (os.Stdout) to output debug messages
	)
	jobs, err := client.GetMtssJobs(context.TODO())
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, v := range jobs[:2] {
		fmt.Println(v.Company)
	}
}
