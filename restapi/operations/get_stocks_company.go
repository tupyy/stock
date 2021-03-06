// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetStocksCompanyHandlerFunc turns a function with the right signature into a get stocks company handler
type GetStocksCompanyHandlerFunc func(GetStocksCompanyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetStocksCompanyHandlerFunc) Handle(params GetStocksCompanyParams) middleware.Responder {
	return fn(params)
}

// GetStocksCompanyHandler interface for that can handle valid get stocks company params
type GetStocksCompanyHandler interface {
	Handle(GetStocksCompanyParams) middleware.Responder
}

// NewGetStocksCompany creates a new http.Handler for the get stocks company operation
func NewGetStocksCompany(ctx *middleware.Context, handler GetStocksCompanyHandler) *GetStocksCompany {
	return &GetStocksCompany{Context: ctx, Handler: handler}
}

/*GetStocksCompany swagger:route GET /stocks/{company} getStocksCompany

get daily stock values for a company

*/
type GetStocksCompany struct {
	Context *middleware.Context
	Handler GetStocksCompanyHandler
}

func (o *GetStocksCompany) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetStocksCompanyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
