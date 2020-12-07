// Code generated by go-swagger; DO NOT EDIT.

package storages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	models "mlss-controlcenter-go/pkg/models"
)

// NewAddStorageParams creates a new AddStorageParams object
// no default values defined in spec.
func NewAddStorageParams() AddStorageParams {

	return AddStorageParams{}
}

// AddStorageParams contains all the bound params for the add storage operation
// typically these are obtained from a http.Request
//
// swagger:parameters AddStorage
type AddStorageParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The Storage Request
	  Required: true
	  In: body
	*/
	Storage *models.Storage
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewAddStorageParams() beforehand.
func (o *AddStorageParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Storage
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("storage", "body"))
			} else {
				res = append(res, errors.NewParseError("storage", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Storage = &body
			}
		}
	} else {
		res = append(res, errors.Required("storage", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}