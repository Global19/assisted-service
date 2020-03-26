// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"github.com/filanov/bm-inventory/models"
)

// NewUpdateClusterParams creates a new UpdateClusterParams object
// no default values defined in spec.
func NewUpdateClusterParams() UpdateClusterParams {

	return UpdateClusterParams{}
}

// UpdateClusterParams contains all the bound params for the update cluster operation
// typically these are obtained from a http.Request
//
// swagger:parameters UpdateCluster
type UpdateClusterParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*New cluster parameters
	  Required: true
	  In: body
	*/
	ClusterUpdateParams *models.ClusterUpdateParams
	/*The ID of the cluster to retrieve
	  Required: true
	  In: path
	*/
	ClusterID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateClusterParams() beforehand.
func (o *UpdateClusterParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ClusterUpdateParams
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("clusterUpdateParams", "body"))
			} else {
				res = append(res, errors.NewParseError("clusterUpdateParams", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ClusterUpdateParams = &body
			}
		}
	} else {
		res = append(res, errors.Required("clusterUpdateParams", "body"))
	}
	rClusterID, rhkClusterID, _ := route.Params.GetOK("cluster_id")
	if err := o.bindClusterID(rClusterID, rhkClusterID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindClusterID binds and validates parameter ClusterID from path.
func (o *UpdateClusterParams) bindClusterID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ClusterID = raw

	return nil
}
