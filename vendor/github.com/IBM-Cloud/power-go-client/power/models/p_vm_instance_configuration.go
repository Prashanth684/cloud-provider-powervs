// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PVMInstanceConfiguration p VM instance configuration
//
// swagger:model PVMInstanceConfiguration
type PVMInstanceConfiguration struct {

	// Console language and code
	ConsoleLanguage *ConsoleLanguage `json:"consoleLanguage,omitempty"`

	// The VTL license repository capacity TB value
	LicenseRepositoryCapacity int64 `json:"licenseRepositoryCapacity,omitempty"`

	// If this is an SAP pvm-instance the profile reference will link to the SAP profile
	SapProfile *SAPProfileReference `json:"sapProfile,omitempty"`

	// The pvm instance Software Licenses
	SoftwareLicenses *SoftwareLicenses `json:"softwareLicenses,omitempty"`

	// The pvm instance system reference code lists
	SystemReferenceCodes [][]*SRC `json:"systemReferenceCodes"`
}

// Validate validates this p VM instance configuration
func (m *PVMInstanceConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsoleLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSapProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareLicenses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemReferenceCodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PVMInstanceConfiguration) validateConsoleLanguage(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsoleLanguage) { // not required
		return nil
	}

	if m.ConsoleLanguage != nil {
		if err := m.ConsoleLanguage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consoleLanguage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consoleLanguage")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstanceConfiguration) validateSapProfile(formats strfmt.Registry) error {
	if swag.IsZero(m.SapProfile) { // not required
		return nil
	}

	if m.SapProfile != nil {
		if err := m.SapProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sapProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sapProfile")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstanceConfiguration) validateSoftwareLicenses(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareLicenses) { // not required
		return nil
	}

	if m.SoftwareLicenses != nil {
		if err := m.SoftwareLicenses.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("softwareLicenses")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("softwareLicenses")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstanceConfiguration) validateSystemReferenceCodes(formats strfmt.Registry) error {
	if swag.IsZero(m.SystemReferenceCodes) { // not required
		return nil
	}

	for i := 0; i < len(m.SystemReferenceCodes); i++ {

		for ii := 0; ii < len(m.SystemReferenceCodes[i]); ii++ {
			if swag.IsZero(m.SystemReferenceCodes[i][ii]) { // not required
				continue
			}

			if m.SystemReferenceCodes[i][ii] != nil {
				if err := m.SystemReferenceCodes[i][ii].Validate(formats); err != nil {
					if ve, ok := err.(*errors.Validation); ok {
						return ve.ValidateName("systemReferenceCodes" + "." + strconv.Itoa(i) + "." + strconv.Itoa(ii))
					} else if ce, ok := err.(*errors.CompositeError); ok {
						return ce.ValidateName("systemReferenceCodes" + "." + strconv.Itoa(i) + "." + strconv.Itoa(ii))
					}
					return err
				}
			}

		}

	}

	return nil
}

// ContextValidate validate this p VM instance configuration based on the context it is used
func (m *PVMInstanceConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConsoleLanguage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSapProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareLicenses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSystemReferenceCodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PVMInstanceConfiguration) contextValidateConsoleLanguage(ctx context.Context, formats strfmt.Registry) error {

	if m.ConsoleLanguage != nil {
		if err := m.ConsoleLanguage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consoleLanguage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consoleLanguage")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstanceConfiguration) contextValidateSapProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.SapProfile != nil {
		if err := m.SapProfile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sapProfile")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sapProfile")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstanceConfiguration) contextValidateSoftwareLicenses(ctx context.Context, formats strfmt.Registry) error {

	if m.SoftwareLicenses != nil {
		if err := m.SoftwareLicenses.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("softwareLicenses")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("softwareLicenses")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstanceConfiguration) contextValidateSystemReferenceCodes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SystemReferenceCodes); i++ {

		for ii := 0; ii < len(m.SystemReferenceCodes[i]); ii++ {

			if m.SystemReferenceCodes[i][ii] != nil {
				if err := m.SystemReferenceCodes[i][ii].ContextValidate(ctx, formats); err != nil {
					if ve, ok := err.(*errors.Validation); ok {
						return ve.ValidateName("systemReferenceCodes" + "." + strconv.Itoa(i) + "." + strconv.Itoa(ii))
					} else if ce, ok := err.(*errors.CompositeError); ok {
						return ce.ValidateName("systemReferenceCodes" + "." + strconv.Itoa(i) + "." + strconv.Itoa(ii))
					}
					return err
				}
			}

		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PVMInstanceConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PVMInstanceConfiguration) UnmarshalBinary(b []byte) error {
	var res PVMInstanceConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
