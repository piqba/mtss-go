package main

import (
	"context"
	"fmt"
	"log"

	"github.com/piqba/mtss-go"
)

func main() {
	baseURL := mtss.URL_BASE
	skipVerify := true
	api := mtss.NewAPIClient(
		baseURL,
		skipVerify,
		nil, //custom http client, defaults to http.DefaultClient
		nil, //io.Writer (os.Stdout) to output debug messages
	)
	jobs, err := api.MtssJobs(context.TODO())
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, v := range jobs[:1] {
		fmt.Printf("%+v", v)
	}
}
