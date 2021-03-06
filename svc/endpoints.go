// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 72999ebd2f
// Version Date: Wed Mar 17 08:36:51 UTC 2021

package svc

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats.

import (
	"context"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	pb "authentication"
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
//
// In a server, it's useful for functions that need to operate on a per-endpoint
// basis. For example, you might pass an Endpoints to a function that produces
// an http.Handler, with each method (endpoint) wired up to a specific path. (It
// is probably a mistake in design to invoke the Service methods on the
// Endpoints struct in a server.)
//
// In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them into an Endpoints, and return it to the caller as a Service.
type Endpoints struct {
	httpServerOptions    map[string][]httptransport.ServerOption
	httpRequestDecoders  map[string]httptransport.DecodeRequestFunc
	httpResponseEncoders map[string]httptransport.EncodeResponseFunc
	httpHandlerFuncs     map[string]func(http.ResponseWriter, *http.Request)

	RegisterEndpoint       endpoint.Endpoint
	SignInEndpoint         endpoint.Endpoint
	SignOutEndpoint        endpoint.Endpoint
	RefreshEndpoint        endpoint.Endpoint
	GetPermissionsEndpoint endpoint.Endpoint
}

func NewEndpoints() Endpoints {
	return Endpoints{
		httpServerOptions:    make(map[string][]httptransport.ServerOption),
		httpRequestDecoders:  make(map[string]httptransport.DecodeRequestFunc),
		httpResponseEncoders: make(map[string]httptransport.EncodeResponseFunc),
		httpHandlerFuncs:     make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

// Endpoints

func (e Endpoints) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	response, err := e.RegisterEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.RegisterResponse), nil
}

func (e Endpoints) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	response, err := e.SignInEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.SignInResponse), nil
}

func (e Endpoints) SignOut(ctx context.Context, in *pb.SignOutRequest) (*pb.SignOutResponse, error) {
	response, err := e.SignOutEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.SignOutResponse), nil
}

func (e Endpoints) Refresh(ctx context.Context, in *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	response, err := e.RefreshEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.RefreshResponse), nil
}

func (e Endpoints) GetPermissions(ctx context.Context, in *pb.GetPermissionsRequest) (*pb.GetPermissionsResponse, error) {
	response, err := e.GetPermissionsEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.GetPermissionsResponse), nil
}

// Make Endpoints

func MakeRegisterEndpoint(s pb.AuthenticationServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.RegisterRequest)
		v, err := s.Register(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeSignInEndpoint(s pb.AuthenticationServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.SignInRequest)
		v, err := s.SignIn(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeSignOutEndpoint(s pb.AuthenticationServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.SignOutRequest)
		v, err := s.SignOut(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeRefreshEndpoint(s pb.AuthenticationServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.RefreshRequest)
		v, err := s.Refresh(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeGetPermissionsEndpoint(s pb.AuthenticationServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.GetPermissionsRequest)
		v, err := s.GetPermissions(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

// WrapAllExcept wraps each Endpoint field of struct Endpoints with a
// go-kit/kit/endpoint.Middleware.
// Use this for applying a set of middlewares to every endpoint in the service.
// Optionally, endpoints can be passed in by name to be excluded from being wrapped.
// WrapAllExcept(middleware, "Status", "Ping")
func (e *Endpoints) WrapAllExcept(middleware endpoint.Middleware, excluded ...string) {
	included := map[string]struct{}{
		"Register":       {},
		"SignIn":         {},
		"SignOut":        {},
		"Refresh":        {},
		"GetPermissions": {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "Register" {
			e.RegisterEndpoint = middleware(e.RegisterEndpoint)
		}
		if inc == "SignIn" {
			e.SignInEndpoint = middleware(e.SignInEndpoint)
		}
		if inc == "SignOut" {
			e.SignOutEndpoint = middleware(e.SignOutEndpoint)
		}
		if inc == "Refresh" {
			e.RefreshEndpoint = middleware(e.RefreshEndpoint)
		}
		if inc == "GetPermissions" {
			e.GetPermissionsEndpoint = middleware(e.GetPermissionsEndpoint)
		}
	}
}

// LabeledMiddleware will get passed the endpoint name when passed to
// WrapAllLabeledExcept, this can be used to write a generic metrics
// middleware which can send the endpoint name to the metrics collector.
type LabeledMiddleware func(string, endpoint.Endpoint) endpoint.Endpoint

// WrapAllLabeledExcept wraps each Endpoint field of struct Endpoints with a
// LabeledMiddleware, which will receive the name of the endpoint. See
// LabeldMiddleware. See method WrapAllExept for details on excluded
// functionality.
func (e *Endpoints) WrapAllLabeledExcept(middleware func(string, endpoint.Endpoint) endpoint.Endpoint, excluded ...string) {
	included := map[string]struct{}{
		"Register":       {},
		"SignIn":         {},
		"SignOut":        {},
		"Refresh":        {},
		"GetPermissions": {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "Register" {
			e.RegisterEndpoint = middleware("Register", e.RegisterEndpoint)
		}
		if inc == "SignIn" {
			e.SignInEndpoint = middleware("SignIn", e.SignInEndpoint)
		}
		if inc == "SignOut" {
			e.SignOutEndpoint = middleware("SignOut", e.SignOutEndpoint)
		}
		if inc == "Refresh" {
			e.RefreshEndpoint = middleware("Refresh", e.RefreshEndpoint)
		}
		if inc == "GetPermissions" {
			e.GetPermissionsEndpoint = middleware("GetPermissions", e.GetPermissionsEndpoint)
		}
	}
}

// WrapAllWithHttpOptionExcept wraps each Endpoint entry of filed HttpServerOptions of struct Endpoints with a
// httptransport.ServerOption.
// Use this for applying a set of server options to every endpoint in the service.
// Optionally, endpoints can be passed in by name to be excluded from being wrapped.
// WrapAllWithHttpOptionExcept(serverOption, "Status", "Ping")
func (e *Endpoints) WrapAllWithHttpOptionExcept(serverOption httptransport.ServerOption, excluded ...string) {
	included := map[string]struct{}{
		"Register":       {},
		"SignIn":         {},
		"SignOut":        {},
		"Refresh":        {},
		"GetPermissions": {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		var options []httptransport.ServerOption
		if o, ok := e.httpServerOptions[inc]; ok {
			options = append(o, serverOption)
		} else {
			options = make([]httptransport.ServerOption, 1)
			options[0] = serverOption
		}
		e.httpServerOptions[inc] = options
	}
}

// WrapWithHttpOption wraps one Endpoint entry of filed HttpServerOptions of struct Endpoints with a
// httptransport.ServerOption.
// WrapWithHttpOption(serverOption, "Status")
func (e *Endpoints) WrapWithHttpOption(endpoint string, serverOption httptransport.ServerOption) {
	var options []httptransport.ServerOption
	if o, ok := e.httpServerOptions[endpoint]; ok {
		options = append(o, serverOption)
	} else {
		options = []httptransport.ServerOption{
			serverOption,
		}
	}
	e.httpServerOptions[endpoint] = options
}

// GetHttpServerOptions returns all httptransport.ServerOption associated with the given endpoint.
func (e Endpoints) GetHttpServerOptions(endpoint string) []httptransport.ServerOption {
	if options, ok := e.httpServerOptions[endpoint]; ok {
		return options
	}
	return make([]httptransport.ServerOption, 0)
}

// SetHttpRequestDecoder assigns a httptransport.DecodeRequestFunc to an endpoint.
func (e Endpoints) SetHttpRequestDecoder(endpoint string, decoder httptransport.DecodeRequestFunc) {
	e.httpRequestDecoders[endpoint] = decoder
}

// GetHttpRequestDecoder returns the httptransport.DecodeRequestFunc associated with the given endpoint.
func (e Endpoints) GetHttpRequestDecoder(endpoint string, fallback httptransport.DecodeRequestFunc) httptransport.DecodeRequestFunc {
	if decoder, ok := e.httpRequestDecoders[endpoint]; ok {
		return decoder
	}
	return fallback
}

// SetHttpResponseEncoder assigns a httptransport.EncodeResponseFunc to an endpoint.
func (e Endpoints) SetHttpResponseEncoder(endpoint string, encoder httptransport.EncodeResponseFunc) {
	e.httpResponseEncoders[endpoint] = encoder
}

// GetHttpResponseEncoder returns the httptransport.EncodeResponseFunc associated with the given endpoint.
func (e Endpoints) GetHttpResponseEncoder(endpoint string, fallback httptransport.EncodeResponseFunc) httptransport.EncodeResponseFunc {
	if encoder, ok := e.httpResponseEncoders[endpoint]; ok {
		return encoder
	}
	return fallback
}

// SetHttpHandlerFunc assigns a custom http HandlerFunc to an endpoint instead of using the default one.
func (e Endpoints) SetHttpHandlerFunc(endpoint string, handler func(http.ResponseWriter, *http.Request)) {
	e.httpHandlerFuncs[endpoint] = handler
}

// GetHttpHandlerFunc returns the http HandlerFunc for the given endpoint.
func (e Endpoints) GetHttpHandlerFunc(endpoint string) func(http.ResponseWriter, *http.Request) {
	if handler, ok := e.httpHandlerFuncs[endpoint]; ok {
		return handler
	}
	return nil
}

// HasHttpHandlerFunc checks if a custom http HandlerFunc is associated with the given endpoint.
func (e Endpoints) HasHttpHandlerFunc(endpoint string) bool {
	_, ok := e.httpHandlerFuncs[endpoint]
	return ok
}
