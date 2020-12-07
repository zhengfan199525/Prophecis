// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-controlcenter-go/pkg/models"
)

// ReceiveTaskAlertOKCode is the HTTP code returned for type ReceiveTaskAlertOK
const ReceiveTaskAlertOKCode int = 200

/*ReceiveTaskAlertOK Detailed role and role information.

swagger:response receiveTaskAlertOK
*/
type ReceiveTaskAlertOK struct {

	/*
	  In: Body
	*/
	Payload *models.Result `json:"body,omitempty"`
}

// NewReceiveTaskAlertOK creates ReceiveTaskAlertOK with default headers values
func NewReceiveTaskAlertOK() *ReceiveTaskAlertOK {

	return &ReceiveTaskAlertOK{}
}

// WithPayload adds the payload to the receive task alert o k response
func (o *ReceiveTaskAlertOK) WithPayload(payload *models.Result) *ReceiveTaskAlertOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the receive task alert o k response
func (o *ReceiveTaskAlertOK) SetPayload(payload *models.Result) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReceiveTaskAlertOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReceiveTaskAlertUnauthorizedCode is the HTTP code returned for type ReceiveTaskAlertUnauthorized
const ReceiveTaskAlertUnauthorizedCode int = 401

/*ReceiveTaskAlertUnauthorized Unauthorized

swagger:response receiveTaskAlertUnauthorized
*/
type ReceiveTaskAlertUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReceiveTaskAlertUnauthorized creates ReceiveTaskAlertUnauthorized with default headers values
func NewReceiveTaskAlertUnauthorized() *ReceiveTaskAlertUnauthorized {

	return &ReceiveTaskAlertUnauthorized{}
}

// WithPayload adds the payload to the receive task alert unauthorized response
func (o *ReceiveTaskAlertUnauthorized) WithPayload(payload *models.Error) *ReceiveTaskAlertUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the receive task alert unauthorized response
func (o *ReceiveTaskAlertUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReceiveTaskAlertUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReceiveTaskAlertNotFoundCode is the HTTP code returned for type ReceiveTaskAlertNotFound
const ReceiveTaskAlertNotFoundCode int = 404

/*ReceiveTaskAlertNotFound url to add role not found.

swagger:response receiveTaskAlertNotFound
*/
type ReceiveTaskAlertNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReceiveTaskAlertNotFound creates ReceiveTaskAlertNotFound with default headers values
func NewReceiveTaskAlertNotFound() *ReceiveTaskAlertNotFound {

	return &ReceiveTaskAlertNotFound{}
}

// WithPayload adds the payload to the receive task alert not found response
func (o *ReceiveTaskAlertNotFound) WithPayload(payload *models.Error) *ReceiveTaskAlertNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the receive task alert not found response
func (o *ReceiveTaskAlertNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReceiveTaskAlertNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}