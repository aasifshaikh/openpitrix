// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixModifyMarketResponse openpitrix modify market response
// swagger:model openpitrixModifyMarketResponse
type OpenpitrixModifyMarketResponse struct {

	// market id
	MarketID string `json:"market_id,omitempty"`
}

// Validate validates this openpitrix modify market response
func (m *OpenpitrixModifyMarketResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixModifyMarketResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixModifyMarketResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixModifyMarketResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
