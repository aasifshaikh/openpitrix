// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixCreateTaskRequest openpitrix create task request
// swagger:model openpitrixCreateTaskRequest
type OpenpitrixCreateTaskRequest struct {

	//
	string `json:"_,omitempty"`

	// directive
	Directive string `json:"directive,omitempty"`

	// job id
	JobID string `json:"job_id,omitempty"`

	// node id
	NodeID string `json:"node_id,omitempty"`

	// target
	Target string `json:"target,omitempty"`

	// task action
	TaskAction string `json:"task_action,omitempty"`
}

// Validate validates this openpitrix create task request
func (m *OpenpitrixCreateTaskRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixCreateTaskRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixCreateTaskRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixCreateTaskRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
