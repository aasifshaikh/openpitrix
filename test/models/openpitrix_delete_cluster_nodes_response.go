// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixDeleteClusterNodesResponse openpitrix delete cluster nodes response
// swagger:model openpitrixDeleteClusterNodesResponse
type OpenpitrixDeleteClusterNodesResponse struct {

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// job id
	JobID string `json:"job_id,omitempty"`
}

// Validate validates this openpitrix delete cluster nodes response
func (m *OpenpitrixDeleteClusterNodesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixDeleteClusterNodesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixDeleteClusterNodesResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixDeleteClusterNodesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
