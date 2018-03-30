package client

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"reflect"

	"github.com/google/go-querystring/query"
)

const (
	libbraryVersion   = "1.0"
	defaultUserAgent  = "imgx-go/" + libbraryVersion
	defaultBaseURL    = ""
	defaultApiVersion = "v1"
)

type Client struct {
	Image      *ImageService
	HttpClient *http.Client // HTTP client used to communicate with the API.
	BaseURL    *url.URL
	UserAgent  string
	ApiVersion string
	Debug      bool
}

func NewClient(apiUrl string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	if len(apiUrl) == 0 {
		apiUrl = defaultBaseURL
	}
	baseURL, _ := url.Parse(apiUrl)

	c := &Client{
		HttpClient: httpClient,
		BaseURL:    baseURL,
		UserAgent:  defaultUserAgent,
		ApiVersion: defaultApiVersion,
	}
	c.Image = &ImageService{client: c}
	return c
}

func (c *Client) versioned(url string) string {
	return path.Join(c.ApiVersion, url)
}

// addOptions adds the parameters in opt as URL query parameters to s. opt
// must be a struct whose fields may contain "url" tags.
func AddOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(c.versioned(urlStr))
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

func (c *Client) NewUploadRequest(urlStr string, reader io.Reader, mediaType string) (*http.Request, error) {
	rel, err := url.Parse(c.versioned(urlStr))
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest("POST", u.String(), reader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", mediaType)
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	if c.Debug {
		log.Printf("Executing request (%v): %#v", req.URL, req)
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if c.Debug {
		log.Printf("Response received: %#v", resp)
	}

	err = checkResponse(resp)
	if err != nil {
		return resp, err
	}

	// If v implements the io.Writer,
	// the response body is decoded into v.
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return resp, err
}

func checkResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	switch r.StatusCode {
	case 400:
		errsResp := &ErrorsResponse{Response: r}
		if err := json.NewDecoder(r.Body).Decode(errsResp); err != nil {
			return err
		}
		return errsResp
	default:
		errResp := &ErrorResponse{Response: r}
		if err := json.NewDecoder(r.Body).Decode(errResp); err != nil {
			return err
		}
		return errResp

	}
}
