// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteStockCompanyHandlerFunc turns a function with the right signature into a delete stock company handler
type DeleteStockCompanyHandlerFunc func(DeleteStockCompanyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteStockCompanyHandlerFunc) Handle(params DeleteStockCompanyParams) middleware.Responder {
	return fn(params)
}

// DeleteStockCompanyHandler interface for that can handle valid delete stock company params
type DeleteStockCompanyHandler interface {
	Handle(DeleteStockCompanyParams) middleware.Responder
}

// NewDeleteStockCompany creates a new http.Handler for the delete stock company operation
func NewDeleteStockCompany(ctx *middleware.Context, handler DeleteStockCompanyHandler) *DeleteStockCompany {
	return &DeleteStockCompany{Context: ctx, Handler: handler}
}

/*DeleteStockCompany swagger:route DELETE /stock/{company} deleteStockCompany

stop crawling a company

*/
type DeleteStockCompany struct {
	Context *middleware.Context
	Handler DeleteStockCompanyHandler
}

func (o *DeleteStockCompany) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteStockCompanyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
