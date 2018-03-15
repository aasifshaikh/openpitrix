// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixAddClusterNodesRequest openpitrix add cluster nodes request
// swagger:model openpitrixAddClusterNodesRequest
type OpenpitrixAddClusterNodesRequest struct {

	// advanced param
	AdvancedParam []string `json:"advanced_param"`

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// node count
	NodeCount *ProtobufUint32Value `json:"node_count,omitempty"`

	// role
	Role string `json:"role,omitempty"`
}

// Validate validates this openpitrix add cluster nodes request
func (m *OpenpitrixAddClusterNodesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdvancedParam(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNodeCount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixAddClusterNodesRequest) validateAdvancedParam(formats strfmt.Registry) error {

	if swag.IsZero(m.AdvancedParam) { // not required
		return nil
	}

	return nil
}

func (m *OpenpitrixAddClusterNodesRequest) validateNodeCount(formats strfmt.Registry) error {

	if swag.IsZero(m.NodeCount) { // not required
		return nil
	}

	if m.NodeCount != nil {

		if err := m.NodeCount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node_count")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixAddClusterNodesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixAddClusterNodesRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixAddClusterNodesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
