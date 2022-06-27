// Package private provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Aborts a catchpoint catchup.
	// (DELETE /v2/catchup/{catchpoint})
	AbortCatchup(ctx echo.Context, catchpoint string) error
	// Starts a catchpoint catchup.
	// (POST /v2/catchup/{catchpoint})
	StartCatchup(ctx echo.Context, catchpoint string) error
	// Return a list of participation keys
	// (GET /v2/participation)
	GetParticipationKeys(ctx echo.Context) error
	// Add a participation key to the node
	// (POST /v2/participation)
	AddParticipationKey(ctx echo.Context) error
	// Delete a given participation key by ID
	// (DELETE /v2/participation/{participation-id})
	DeleteParticipationKeyByID(ctx echo.Context, participationId string) error
	// Get participation key info given a participation ID
	// (GET /v2/participation/{participation-id})
	GetParticipationKeyByID(ctx echo.Context, participationId string) error
	// Append state proof keys to a participation key
	// (POST /v2/participation/{participation-id})
	AppendKeys(ctx echo.Context, participationId string) error

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AbortCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) AbortCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AbortCatchup(ctx, catchpoint)
	return err
}

// StartCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) StartCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StartCatchup(ctx, catchpoint)
	return err
}

// GetParticipationKeys converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeys(ctx)
	return err
}

// AddParticipationKey converts echo context to params.
func (w *ServerInterfaceWrapper) AddParticipationKey(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddParticipationKey(ctx)
	return err
}

// DeleteParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteParticipationKeyByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteParticipationKeyByID(ctx, participationId)
	return err
}

// GetParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeyByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeyByID(ctx, participationId)
	return err
}

// AppendKeys converts echo context to params.
func (w *ServerInterfaceWrapper) AppendKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AppendKeys(ctx, participationId)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty":  true,
		"timeout": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------
	if paramValue := ctx.QueryParam("timeout"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE("/v2/catchup/:catchpoint", wrapper.AbortCatchup, m...)
	router.POST("/v2/catchup/:catchpoint", wrapper.StartCatchup, m...)
	router.GET("/v2/participation", wrapper.GetParticipationKeys, m...)
	router.POST("/v2/participation", wrapper.AddParticipationKey, m...)
	router.DELETE("/v2/participation/:participation-id", wrapper.DeleteParticipationKeyByID, m...)
	router.GET("/v2/participation/:participation-id", wrapper.GetParticipationKeyByID, m...)
	router.POST("/v2/participation/:participation-id", wrapper.AppendKeys, m...)
	router.POST("/v2/shutdown", wrapper.ShutdownNode, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"xrQAAA==",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
