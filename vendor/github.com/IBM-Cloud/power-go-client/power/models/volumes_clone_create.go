// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VolumesCloneCreate volumes clone create
//
// swagger:model VolumesCloneCreate
type VolumesCloneCreate struct {

	// Unique name within a cloud instance used to identify a volumes-clone request
	// name can be used in replace of a volumesCloneID when used as a URL path parameter
	//
	// Required: true
	Name *string `json:"name"`

	// Enables performance path for creating volume clones.
	PerformancePath *bool `json:"performancePath,omitempty"`

	// List of volumes to be cloned
	// Required: true
	VolumeIDs []string `json:"volumeIDs"`
}

// Validate validates this volumes clone create
func (m *VolumesCloneCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeIDs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumesCloneCreate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VolumesCloneCreate) validateVolumeIDs(formats strfmt.Registry) error {

	if err := validate.Required("volumeIDs", "body", m.VolumeIDs); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this volumes clone create based on context it is used
func (m *VolumesCloneCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VolumesCloneCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumesCloneCreate) UnmarshalBinary(b []byte) error {
	var res VolumesCloneCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
