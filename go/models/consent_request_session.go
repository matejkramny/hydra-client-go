// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ConsentRequestSession ConsentRequestSession ConsentRequestSession Used to pass session data to a consent request.
// swagger:model consentRequestSession
type ConsentRequestSession struct {

	// AccessToken sets session data for the access and refresh token, as well as any future tokens issued by the
	// refresh grant. Keep in mind that this data will be available to anyone performing OAuth 2.0 Challenge Introspection.
	// If only your services can perform OAuth 2.0 Challenge Introspection, this is usually fine. But if third parties
	// can access that endpoint as well, sensitive data from the session might be exposed to them. Use with care!
	AccessToken map[string]interface{} `json:"access_token,omitempty"`

	// IDToken sets session data for the OpenID Connect ID token. Keep in mind that the session'id payloads are readable
	// by anyone that has access to the ID Challenge. Use with care!
	IDToken map[string]interface{} `json:"id_token,omitempty"`
}

// Validate validates this consent request session
func (m *ConsentRequestSession) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConsentRequestSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsentRequestSession) UnmarshalBinary(b []byte) error {
	var res ConsentRequestSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
