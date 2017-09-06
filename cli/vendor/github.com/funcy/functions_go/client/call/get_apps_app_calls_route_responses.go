package call

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/funcy/functions_go/models"
)

// GetAppsAppCallsRouteReader is a Reader for the GetAppsAppCallsRoute structure.
type GetAppsAppCallsRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppsAppCallsRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAppsAppCallsRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetAppsAppCallsRouteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAppsAppCallsRouteOK creates a GetAppsAppCallsRouteOK with default headers values
func NewGetAppsAppCallsRouteOK() *GetAppsAppCallsRouteOK {
	return &GetAppsAppCallsRouteOK{}
}

/*GetAppsAppCallsRouteOK handles this case with default header values.

Calls found
*/
type GetAppsAppCallsRouteOK struct {
	Payload *models.CallsWrapper
}

func (o *GetAppsAppCallsRouteOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app}/calls/{route}][%d] getAppsAppCallsRouteOK  %+v", 200, o.Payload)
}

func (o *GetAppsAppCallsRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CallsWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppsAppCallsRouteNotFound creates a GetAppsAppCallsRouteNotFound with default headers values
func NewGetAppsAppCallsRouteNotFound() *GetAppsAppCallsRouteNotFound {
	return &GetAppsAppCallsRouteNotFound{}
}

/*GetAppsAppCallsRouteNotFound handles this case with default header values.

Calls not found.
*/
type GetAppsAppCallsRouteNotFound struct {
	Payload *models.Error
}

func (o *GetAppsAppCallsRouteNotFound) Error() string {
	return fmt.Sprintf("[GET /apps/{app}/calls/{route}][%d] getAppsAppCallsRouteNotFound  %+v", 404, o.Payload)
}

func (o *GetAppsAppCallsRouteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}