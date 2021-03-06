// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetCompaniesOKCode is the HTTP code returned for type GetCompaniesOK
const GetCompaniesOKCode int = 200

/*GetCompaniesOK ok

swagger:response getCompaniesOK
*/
type GetCompaniesOK struct {

	/*
	  In: Body
	*/
	Payload *GetCompaniesOKBody `json:"body,omitempty"`
}

// NewGetCompaniesOK creates GetCompaniesOK with default headers values
func NewGetCompaniesOK() *GetCompaniesOK {

	return &GetCompaniesOK{}
}

// WithPayload adds the payload to the get companies o k response
func (o *GetCompaniesOK) WithPayload(payload *GetCompaniesOKBody) *GetCompaniesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get companies o k response
func (o *GetCompaniesOK) SetPayload(payload *GetCompaniesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCompaniesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCompaniesPaymentRequiredCode is the HTTP code returned for type GetCompaniesPaymentRequired
const GetCompaniesPaymentRequiredCode int = 402

/*GetCompaniesPaymentRequired method not allowed

swagger:response getCompaniesPaymentRequired
*/
type GetCompaniesPaymentRequired struct {
}

// NewGetCompaniesPaymentRequired creates GetCompaniesPaymentRequired with default headers values
func NewGetCompaniesPaymentRequired() *GetCompaniesPaymentRequired {

	return &GetCompaniesPaymentRequired{}
}

// WriteResponse to the client
func (o *GetCompaniesPaymentRequired) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(402)
}

// GetCompaniesInternalServerErrorCode is the HTTP code returned for type GetCompaniesInternalServerError
const GetCompaniesInternalServerErrorCode int = 500

/*GetCompaniesInternalServerError Internal error.

swagger:response getCompaniesInternalServerError
*/
type GetCompaniesInternalServerError struct {
}

// NewGetCompaniesInternalServerError creates GetCompaniesInternalServerError with default headers values
func NewGetCompaniesInternalServerError() *GetCompaniesInternalServerError {

	return &GetCompaniesInternalServerError{}
}

// WriteResponse to the client
func (o *GetCompaniesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
