// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/tupyy/stock/models"
)

// GetStocksCompanyOKCode is the HTTP code returned for type GetStocksCompanyOK
const GetStocksCompanyOKCode int = 200

/*GetStocksCompanyOK ok

swagger:response getStocksCompanyOK
*/
type GetStocksCompanyOK struct {

	/*
	  In: Body
	*/
	Payload *models.StockValues `json:"body,omitempty"`
}

// NewGetStocksCompanyOK creates GetStocksCompanyOK with default headers values
func NewGetStocksCompanyOK() *GetStocksCompanyOK {

	return &GetStocksCompanyOK{}
}

// WithPayload adds the payload to the get stocks company o k response
func (o *GetStocksCompanyOK) WithPayload(payload *models.StockValues) *GetStocksCompanyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get stocks company o k response
func (o *GetStocksCompanyOK) SetPayload(payload *models.StockValues) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStocksCompanyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetStocksCompanyNotFoundCode is the HTTP code returned for type GetStocksCompanyNotFound
const GetStocksCompanyNotFoundCode int = 404

/*GetStocksCompanyNotFound company not found

swagger:response getStocksCompanyNotFound
*/
type GetStocksCompanyNotFound struct {
}

// NewGetStocksCompanyNotFound creates GetStocksCompanyNotFound with default headers values
func NewGetStocksCompanyNotFound() *GetStocksCompanyNotFound {

	return &GetStocksCompanyNotFound{}
}

// WriteResponse to the client
func (o *GetStocksCompanyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
