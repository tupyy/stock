// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetStocksHandlerFunc turns a function with the right signature into a get stocks handler
type GetStocksHandlerFunc func(GetStocksParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetStocksHandlerFunc) Handle(params GetStocksParams) middleware.Responder {
	return fn(params)
}

// GetStocksHandler interface for that can handle valid get stocks params
type GetStocksHandler interface {
	Handle(GetStocksParams) middleware.Responder
}

// NewGetStocks creates a new http.Handler for the get stocks operation
func NewGetStocks(ctx *middleware.Context, handler GetStocksHandler) *GetStocks {
	return &GetStocks{Context: ctx, Handler: handler}
}

/*GetStocks swagger:route GET /stocks getStocks

List all watched .

*/
type GetStocks struct {
	Context *middleware.Context
	Handler GetStocksHandler
}

func (o *GetStocks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetStocksParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetStocksOKBodyTuple0 GetStocksOKBodyTuple0 a representation of an anonymous Tuple type
//
// swagger:model GetStocksOKBodyTuple0
type GetStocksOKBodyTuple0 struct {

	// p0
	// Required: true
	P0 *string `json:"-"` // custom serializer

}

// UnmarshalJSON unmarshals this tuple type from a JSON array
func (o *GetStocksOKBodyTuple0) UnmarshalJSON(raw []byte) error {
	// stage 1, get the array but just the array
	var stage1 []json.RawMessage
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&stage1); err != nil {
		return err
	}

	// stage 2: hydrates struct members with tuple elements
	if len(stage1) > 0 {
		var dataP0 string
		buf = bytes.NewBuffer(stage1[0])
		dec := json.NewDecoder(buf)
		dec.UseNumber()
		if err := dec.Decode(&dataP0); err != nil {
			return err
		}
		o.P0 = &dataP0

	}

	return nil
}

// MarshalJSON marshals this tuple type into a JSON array
func (o GetStocksOKBodyTuple0) MarshalJSON() ([]byte, error) {
	data := []interface{}{
		o.P0,
	}

	return json.Marshal(data)
}

// Validate validates this get stocks o k body tuple0
func (o *GetStocksOKBodyTuple0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateP0(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetStocksOKBodyTuple0) validateP0(formats strfmt.Registry) error {

	if err := validate.Required("P0", "body", o.P0); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetStocksOKBodyTuple0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStocksOKBodyTuple0) UnmarshalBinary(b []byte) error {
	var res GetStocksOKBodyTuple0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
