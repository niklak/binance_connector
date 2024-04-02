package request

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type secType int

const (
	SecTypeNone secType = iota
	SecTypeAPIKey
	SecTypeSigned // if the 'timestamp' parameter is required
)

// Request define an API Request
type Request struct {
	Method              string
	Endpoint            string
	Query               url.Values
	RecvWindow          int64
	SecType             secType
	Header              http.Header
	Body                io.Reader
	FullURL             string
	RequiredParams      []string
	RequiredOneOfParams [][]string
}

// Init initialize Request's query
// if query is nil, it will be initialized as url.Values{}
// if method is empty, it will be initialized as http.MethodGet
func (r *Request) Init() *Request {
	if r.Query == nil {
		r.Query = url.Values{}
	}
	if r.Method == "" {
		r.Method = http.MethodGet
	}
	return r
}

// SetParam set param with key/value to query string, if param is nil it will be ignored
// if param is nil pointer, it will be ignored. if param is pointer it will be dereferenced
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
	case uint64:
		param = strconv.FormatUint(v, 10)
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

func (r *Request) checkRequiredParams() error {
	var errs []error
	for _, key := range r.RequiredParams {
		if len(r.Query[key]) == 0 {
			e := fmt.Errorf("%w: %q", ErrMissingParam, key)
			errs = append(errs, e)
		}
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func (r *Request) checkRequiredEitherParams() error {
	var errs []error
	for _, keys := range r.RequiredOneOfParams {
		var found bool
		for _, key := range keys {
			if len(r.Query[key]) > 0 {
				found = true
				break
			}
		}
		if !found {
			e := fmt.Errorf("%w: one of: %q", ErrMissingParam, strings.Join(keys, ", "))
			errs = append(errs, e)
		}
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

// Validate checks if the request's query has all required parameters.
// Returns all missing parameters as an error.
// Since it checks only required parameters and `required one of the parameters`
// this package doesn't need a heavy validation library such as go-playground/validator or its alternatives.
func (r *Request) Validate() error {
	var errs []error
	if err := r.checkRequiredParams(); err != nil {
		errs = append(errs, err)
	}
	if err := r.checkRequiredEitherParams(); err != nil {
		errs = append(errs, err)

	}
	if len(errs) > 0 {
		return errors.Join(errs...)
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
