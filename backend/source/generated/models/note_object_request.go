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

// NoteObjectRequest note obj
//
// swagger:model noteObjectRequest
type NoteObjectRequest struct {

	// description of the note
	Description string `json:"description,omitempty"`

	// path of file
	// Required: true
	NoteReference *string `json:"note_reference"`

	// style of the note
	// Required: true
	Style *string `json:"style"`

	// tags of the note
	// Required: true
	Tag *string `json:"tag"`

	// title of the note
	Title string `json:"title,omitempty"`

	// type of the note file, public, shared, private
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this note object request
func (m *NoteObjectRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNoteReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStyle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NoteObjectRequest) validateNoteReference(formats strfmt.Registry) error {

	if err := validate.Required("note_reference", "body", m.NoteReference); err != nil {
		return err
	}

	return nil
}

func (m *NoteObjectRequest) validateStyle(formats strfmt.Registry) error {

	if err := validate.Required("style", "body", m.Style); err != nil {
		return err
	}

	return nil
}

func (m *NoteObjectRequest) validateTag(formats strfmt.Registry) error {

	if err := validate.Required("tag", "body", m.Tag); err != nil {
		return err
	}

	return nil
}

func (m *NoteObjectRequest) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this note object request based on context it is used
func (m *NoteObjectRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NoteObjectRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NoteObjectRequest) UnmarshalBinary(b []byte) error {
	var res NoteObjectRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
