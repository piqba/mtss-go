/*
	Package mtssgo is the no-official mtss SDK for Go.
	Use it to interact with the mtss API.
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
		for _, v := range jobs[:1] {
			fmt.Printf("%+v", v)
		}
	}
	Examples can be found at
	https://github.com/piqba/mtss-go/tree/main/example
	If you find an issue with the SDK, please report through
	https://github.com/piqba/mtss-go/issues/new
*/
package mtss
