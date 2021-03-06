// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/tupyy/stock/models"
)

// GetStocksCompanyReader is a Reader for the GetStocksCompany structure.
type GetStocksCompanyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStocksCompanyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStocksCompanyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetStocksCompanyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStocksCompanyOK creates a GetStocksCompanyOK with default headers values
func NewGetStocksCompanyOK() *GetStocksCompanyOK {
	return &GetStocksCompanyOK{}
}

/*GetStocksCompanyOK handles this case with default header values.

ok
*/
type GetStocksCompanyOK struct {
	Payload *models.StockValues
}

func (o *GetStocksCompanyOK) Error() string {
	return fmt.Sprintf("[GET /stocks/{company}][%d] getStocksCompanyOK  %+v", 200, o.Payload)
}

func (o *GetStocksCompanyOK) GetPayload() *models.StockValues {
	return o.Payload
}

func (o *GetStocksCompanyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StockValues)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStocksCompanyNotFound creates a GetStocksCompanyNotFound with default headers values
func NewGetStocksCompanyNotFound() *GetStocksCompanyNotFound {
	return &GetStocksCompanyNotFound{}
}

/*GetStocksCompanyNotFound handles this case with default header values.

company not found
*/
type GetStocksCompanyNotFound struct {
}

func (o *GetStocksCompanyNotFound) Error() string {
	return fmt.Sprintf("[GET /stocks/{company}][%d] getStocksCompanyNotFound ", 404)
}

func (o *GetStocksCompanyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
