// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"ascii-pet/internal/gen/models"
)

// NewUploadPetParams creates a new UploadPetParams object
//
// There are no default values defined in the spec.
func NewUploadPetParams() UploadPetParams {

	return UploadPetParams{}
}

// UploadPetParams contains all the bound params for the upload pet operation
// typically these are obtained from a http.Request
//
// swagger:parameters UploadPet
type UploadPetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Pet properties
	  Required: true
	  In: body
	*/
	PetProperties *models.PetProperties
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUploadPetParams() beforehand.
func (o *UploadPetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PetProperties
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("petProperties", "body", ""))
			} else {
				res = append(res, errors.NewParseError("petProperties", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.PetProperties = &body
			}
		}
	} else {
		res = append(res, errors.Required("petProperties", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
