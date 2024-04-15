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

	value = indirect(value)
	if value == nil {
		return r
	}

	val := toString(value)

	if val != "" {
		r.Query.Set(key, val)
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

// SecType sets security type for Request
func SecType(secType secType) RequestOption {
	return func(r *Request) {
		r.SecType = secType
	}
}

// Method sets method for Request
func Method(method string) RequestOption {
	return func(r *Request) {
		r.Method = method
	}
}

// RequiredParams sets required params for Request
func RequiredParams(params ...string) RequestOption {
	return func(r *Request) {
		r.RequiredParams = params
	}
}

// RequiredOneOfParams sets required one of the set params for Request
func RequiredOneOfParams(params ...[]string) RequestOption {
	return func(r *Request) {
		r.RequiredOneOfParams = params
	}
}

// SetParam sets query parameter with key/value to Request
func SetParam(key string, value interface{}) RequestOption {
	return func(r *Request) {
		r.SetParam(key, value)
	}
}

// Signed sets security type to SecTypeSigned. This request require to be signed with a secret key
func Signed() RequestOption {
	return SecType(SecTypeSigned)
}

// New create a new Request, prefer to use this function to create a new Request, because it will initialize query and form
func New(endpoint string, options ...RequestOption) *Request {
	r := &Request{
		Endpoint: endpoint,
		Query:    url.Values{},
	}
	for _, option := range options {
		option(r)
	}

	if r.Method == "" {
		r.Method = http.MethodGet
	}

	return r
}

// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
// indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
func indirect(a interface{}) interface{} {
	if a == nil {
		return nil
	}
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}
	return v.Interface()
}

func toString(v interface{}) string {
	// At this point v can't be nil, because it is invoked after `indirect` function

	switch s := v.(type) {
	case string:
		return s
	case bool:
		return strconv.FormatBool(s)
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32)
	case int:
		return strconv.Itoa(s)
	case int64:
		return strconv.FormatInt(s, 10)
	case int32:
		return strconv.Itoa(int(s))
	case int16:
		return strconv.FormatInt(int64(s), 10)
	case int8:
		return strconv.FormatInt(int64(s), 10)
	case uint:
		return strconv.FormatUint(uint64(s), 10)
	case uint64:
		return strconv.FormatUint(uint64(s), 10)
	case uint32:
		return strconv.FormatUint(uint64(s), 10)
	case uint16:
		return strconv.FormatUint(uint64(s), 10)
	case uint8:
		return strconv.FormatUint(uint64(s), 10)
	case fmt.Stringer:
		return s.String()
	default:
		return fmt.Sprintf("%v", s)
	}

}
