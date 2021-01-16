// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostStockCompanyCreatedCode is the HTTP code returned for type PostStockCompanyCreated
const PostStockCompanyCreatedCode int = 201

/*PostStockCompanyCreated created

swagger:response postStockCompanyCreated
*/
type PostStockCompanyCreated struct {
}

// NewPostStockCompanyCreated creates PostStockCompanyCreated with default headers values
func NewPostStockCompanyCreated() *PostStockCompanyCreated {

	return &PostStockCompanyCreated{}
}

// WriteResponse to the client
func (o *PostStockCompanyCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostStockCompanyBadRequestCode is the HTTP code returned for type PostStockCompanyBadRequest
const PostStockCompanyBadRequestCode int = 400

/*PostStockCompanyBadRequest bad request

swagger:response postStockCompanyBadRequest
*/
type PostStockCompanyBadRequest struct {
}

// NewPostStockCompanyBadRequest creates PostStockCompanyBadRequest with default headers values
func NewPostStockCompanyBadRequest() *PostStockCompanyBadRequest {

	return &PostStockCompanyBadRequest{}
}

// WriteResponse to the client
func (o *PostStockCompanyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PostStockCompanyInternalServerErrorCode is the HTTP code returned for type PostStockCompanyInternalServerError
const PostStockCompanyInternalServerErrorCode int = 500

/*PostStockCompanyInternalServerError internal error

swagger:response postStockCompanyInternalServerError
*/
type PostStockCompanyInternalServerError struct {
}

// NewPostStockCompanyInternalServerError creates PostStockCompanyInternalServerError with default headers values
func NewPostStockCompanyInternalServerError() *PostStockCompanyInternalServerError {

	return &PostStockCompanyInternalServerError{}
}

// WriteResponse to the client
func (o *PostStockCompanyInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
