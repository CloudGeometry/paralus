// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfigRepoFile config repo file
//
// swagger:model configRepoFile
type ConfigRepoFile struct {

	// file type
	FileType *ConfigFileType `json:"fileType,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// rel path
	RelPath string `json:"relPath,omitempty"`
}

// Validate validates this config repo file
func (m *ConfigRepoFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigRepoFile) validateFileType(formats strfmt.Registry) error {
	if swag.IsZero(m.FileType) { // not required
		return nil
	}

	if m.FileType != nil {
		if err := m.FileType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fileType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fileType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config repo file based on the context it is used
func (m *ConfigRepoFile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFileType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigRepoFile) contextValidateFileType(ctx context.Context, formats strfmt.Registry) error {

	if m.FileType != nil {
		if err := m.FileType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fileType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fileType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigRepoFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigRepoFile) UnmarshalBinary(b []byte) error {
	var res ConfigRepoFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}