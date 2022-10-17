package github

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/msolimans/gitcomp/pkg/xhttp"
)

//represents client that goes to github
type GitClient struct {
	client *http.Client // HTTP client used to communicate with the API.
	BaseURL *url.URL
	UserAgent string
}

func NewClient(httpClient *http.Client) *GitClient {
	//use this for testing purposes
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	//use passed in http client 
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &GitClient{client: httpClient, BaseURL: baseURL, UserAgent: defaultUserAgent}	
	return c
}

func (c *GitClient) InitRequest(method, urlStr string) (*http.Request, error) {
	
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	 
	headers := map[string]string {
		"Accept": acceptHeader,
		"User-Agent": c.UserAgent,
	}
	
	return xhttp.NewRequest(method, u.String(), headers, nil)
}

// If ctx is canceled or times out, ctx.Err() will be returned.
func (c *GitClient) Send(ctx context.Context, req *http.Request, v interface{}) error {

	if ctx == nil {
		return errors.New("context must be non-nil")
	}
	
	resp, err := c.client.Do(req)
	if err != nil {
		// If we got error, while the context has been canceled, return context's error 
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		//otherwise return the error
		return err
	}
	defer resp.Body.Close()

	if err = checkResponseForErrors(resp); err != nil {
		return err
	}

	decErr := json.NewDecoder(resp.Body).Decode(v)
	if decErr == io.EOF {
		decErr = nil // ignore EOF errors caused by empty response body
	} else {
		return decErr
	}
	return nil 
}

func checkResponseForErrors(r *http.Response) error {
	 
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	//otherwise return error
	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}
	
	return errorResponse
}
