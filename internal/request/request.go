package request

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type secType int

const (
	SecTypeNone secType = iota
	SecTypeAPIKey
	SecTypeSigned // if the 'timestamp' parameter is required
)

type params map[string]interface{}

// Request define an API Request
type Request struct {
	Method     string
	Endpoint   string
	Query      url.Values
	Form       url.Values
	RecvWindow int64
	SecType    secType
	Header     http.Header
	Body       io.Reader
	FullURL    string
}

// Init initialize Request's query
func (r *Request) Init() *Request {
	if r.Query == nil {
		r.Query = url.Values{}
	}
	if r.Form == nil {
		r.Form = url.Values{}
	}
	return r
}

// AddParam add param with key/value to query string
func (r *Request) AddParam(key string, value interface{}) *Request {
	r.Query.Add(key, fmt.Sprintf("%v", value))
	return r
}

// SetParam set param with key/value to query string, if param is nil it will be ignored
func (r *Request) SetParam(key string, value interface{}) *Request {
	// better to use reflection to handle all types
	var param string
	switch v := value.(type) {
	case nil:
		return r
	case string:
		param = v
	case int:
		param = strconv.Itoa(v)
	case int64:
		param = strconv.FormatInt(v, 10)
	case float32:
		param = strconv.FormatFloat(float64(v), 'f', -1, 64)
	case float64:
		param = strconv.FormatFloat(v, 'f', -1, 64)
	case *string:
		if v != nil {
			param = *v
		}
	case *int:
		if v != nil {
			param = strconv.Itoa(*v)
		}
	case *int64:
		if v != nil {
			param = strconv.FormatInt(*v, 10)
		}
	case *float32:
		if v != nil {
			param = strconv.FormatFloat(float64(*v), 'f', -1, 64)
		}
	case *float64:
		if v != nil {
			param = strconv.FormatFloat(*v, 'f', -1, 64)
		}
	default:
		param = fmt.Sprintf("%v", value)
	}
	if param != "" {
		r.Query.Set(key, param)
	}

	return r
}

// setParams set params with key/values to query string
func (r *Request) SetParams(m params) *Request {
	for k, v := range m {
		r.SetParam(k, v)
	}
	return r
}

func (r *Request) Validate() (err error) {
	if r.Query == nil {
		r.Query = url.Values{}
	}
	if r.Form == nil {
		r.Form = url.Values{}
	}
	return nil
}

// Append `WithRecvWindow(insert_recvwindow)` to Request to modify the default recvWindow value
func WithRecvWindow(recvWindow int64) RequestOption {
	return func(r *Request) {
		r.RecvWindow = recvWindow
	}
}

// RequestOption define option type for Request
type RequestOption func(*Request)
