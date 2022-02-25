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

// V3ClusterNodeStatus v3 cluster node status
//
// swagger:model v3ClusterNodeStatus
type V3ClusterNodeStatus struct {

	// allocatable
	Allocatable *V3Resources `json:"allocatable,omitempty"`

	// allocated
	Allocated *V3Resources `json:"allocated,omitempty"`

	// capacity
	Capacity *V3Resources `json:"capacity,omitempty"`

	// conditions
	Conditions []*V1NodeCondition `json:"conditions"`

	// ips
	Ips []*V3ClusterNodeIP `json:"ips"`

	// node info
	NodeInfo *V1NodeSystemInfo `json:"nodeInfo,omitempty"`

	// state
	State *V3ClusterNodeState `json:"state,omitempty"`
}

// Validate validates this v3 cluster node status
func (m *V3ClusterNodeStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocatable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllocated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapacity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3ClusterNodeStatus) validateAllocatable(formats strfmt.Registry) error {
	if swag.IsZero(m.Allocatable) { // not required
		return nil
	}

	if m.Allocatable != nil {
		if err := m.Allocatable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocatable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("allocatable")
			}
			return err
		}
	}

	return nil
}

func (m *V3ClusterNodeStatus) validateAllocated(formats strfmt.Registry) error {
	if swag.IsZero(m.Allocated) { // not required
		return nil
	}

	if m.Allocated != nil {
		if err := m.Allocated.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocated")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("allocated")
			}
			return err
		}
	}

	return nil
}

func (m *V3ClusterNodeStatus) validateCapacity(formats strfmt.Registry) error {
	if swag.IsZero(m.Capacity) { // not required
		return nil
	}

	if m.Capacity != nil {
		if err := m.Capacity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capacity")
			}
			return err
		}
	}

	return nil
}

func (m *V3ClusterNodeStatus) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3ClusterNodeStatus) validateIps(formats strfmt.Registry) error {
	if swag.IsZero(m.Ips) { // not required
		return nil
	}

	for i := 0; i < len(m.Ips); i++ {
		if swag.IsZero(m.Ips[i]) { // not required
			continue
		}

		if m.Ips[i] != nil {
			if err := m.Ips[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ips" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ips" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3ClusterNodeStatus) validateNodeInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeInfo) { // not required
		return nil
	}

	if m.NodeInfo != nil {
		if err := m.NodeInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeInfo")
			}
			return err
		}
	}

	return nil
}

func (m *V3ClusterNodeStatus) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v3 cluster node status based on the context it is used
func (m *V3ClusterNodeStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllocatable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAllocated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCapacity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3ClusterNodeStatus) contextValidateAllocatable(ctx context.Context, formats strfmt.Registry) error {

	if m.Allocatable != nil {
		if err := m.Allocatable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocatable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("allocatable")
			}
			return err
		}
	}

	return nil
}

func (m *V3ClusterNodeStatus) contextValidateAllocated(ctx context.Context, formats strfmt.Registry) error {

	if m.Allocated != nil {
		if err := m.Allocated.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allocated")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("allocated")
			}
			return err
		}
	}

	return nil
}

func (m *V3ClusterNodeStatus) contextValidateCapacity(ctx context.Context, formats strfmt.Registry) error {

	if m.Capacity != nil {
		if err := m.Capacity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("capacity")
			}
			return err
		}
	}

	return nil
}

func (m *V3ClusterNodeStatus) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Conditions); i++ {

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3ClusterNodeStatus) contextValidateIps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Ips); i++ {

		if m.Ips[i] != nil {
			if err := m.Ips[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ips" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ips" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3ClusterNodeStatus) contextValidateNodeInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeInfo != nil {
		if err := m.NodeInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeInfo")
			}
			return err
		}
	}

	return nil
}

func (m *V3ClusterNodeStatus) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {
		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3ClusterNodeStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3ClusterNodeStatus) UnmarshalBinary(b []byte) error {
	var res V3ClusterNodeStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}