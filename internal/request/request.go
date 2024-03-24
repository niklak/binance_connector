package request

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
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
	v := reflect.ValueOf(value)

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return r
		} else {
			value = v.Elem().Interface()
		}
	}
	var param string
	switch v := value.(type) {
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

func SecType(secType secType) RequestOption {
	return func(r *Request) {
		r.SecType = secType
	}
}

func Method(method string) RequestOption {
	return func(r *Request) {
		r.Method = method
	}
}

// New create a new Request, prefer to use this function to create a new Request, because it will initialize query and form
func New(endpoint string, options ...RequestOption) *Request {
	r := &Request{
		Endpoint: endpoint,
	}
	for _, option := range options {
		option(r)
	}
	return r.Init()
}
