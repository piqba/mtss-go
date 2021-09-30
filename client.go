package mtss

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/piqba/mtss-go/pkg/errors"
)

const (
	// OFFERS are the path that show all jobs on the API
	OFFERS = "necesidades"
	// URL_BASE are the url for the jobs endpoints
	URL_BASE = "https://apiempleo.xutil.net/apiempleo"
)

// Client is an interface that implements Mtss's API
type Client interface {
	// MtssJobs returns the corresponding jobs on fetch call, or an error.
	MtssJobs(
		ctx context.Context,
	) ([]Mtss, error)
}

// Client
// client represents a Mtss client. If the Debug field is set to an io.Writer
// (for example os.Stdout), then the client will dump API requests and responses
// to it.  To use a non-default HTTP client (for example, for testing, or to set
// a timeout), assign to the HTTPClient field. To set a non-default URL (for
// example, for testing), assign to the URL field.
type client struct {
	url        string
	httpClient *http.Client
	debug      io.Writer
}

func NewAPIClient(
	// mtss API's base url
	baseURL string,
	// skipVerify
	skipVerify bool,
	//optional, defaults to http.DefaultClient
	httpClient *http.Client,
	debug io.Writer,
) Client {

	c := &client{
		url:        baseURL,
		httpClient: httpClient,
		debug:      debug,
	}
	if httpClient != nil {
		c.httpClient = httpClient
	} else {
		c.httpClient = http.DefaultClient
	}
	if skipVerify {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		c.httpClient.Transport = tr
	}
	return c
}

// dumpResponse writes the raw response data to the debug output, if set, or
// standard error otherwise.
func (c *client) dumpResponse(resp *http.Response) {
	// ignore errors dumping response - no recovery from this
	responseDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatalf("mtss/clientHttp: dumpResponse: " + err.Error())
	}
	fmt.Fprintln(c.debug, string(responseDump))
	fmt.Fprintln(c.debug)
}

// apiCall define how you can make a call to Mtss API
func (c *client) apiCall(
	ctx context.Context,
	method string,
	URL string,
	data []byte,
) (statusCode int, response string, err error) {
	requestURL := c.url + "/" + URL
	req, err := http.NewRequest(method, requestURL, bytes.NewBuffer(data))
	if err != nil {
		return 0, "", fmt.Errorf("mtss/clientHttp: failed to create HTTP request: %v", err)
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Set("User-Agent", "mtssgo-client/0.0")
	if c.debug != nil {
		requestDump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return 0, "", errors.Errorf("mtss/clientHttp: error dumping HTTP request: %v", err)
		}
		fmt.Fprintln(c.debug, string(requestDump))
		fmt.Fprintln(c.debug)
	}
	req = req.WithContext(ctx)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, "", errors.Errorf("mtss/clientHttp: HTTP request failed with: %v", err)
	}
	defer resp.Body.Close()
	if c.debug != nil {
		c.dumpResponse(resp)
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", errors.Errorf("mtss/clientHttp: HTTP request failed: %v", err)
	}
	return resp.StatusCode, string(res), nil
}
