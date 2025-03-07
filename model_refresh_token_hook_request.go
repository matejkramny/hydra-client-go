/*
 * Ory Oathkeeper API
 *
 * Documentation for all of Ory Oathkeeper's APIs. 
 *
 * API version: v1.11.7
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// RefreshTokenHookRequest struct for RefreshTokenHookRequest
type RefreshTokenHookRequest struct {
	// ClientID is the identifier of the OAuth 2.0 client.
	ClientId *string `json:"client_id,omitempty"`
	// GrantedAudience is the list of audiences granted to the OAuth 2.0 client.
	GrantedAudience []string `json:"granted_audience,omitempty"`
	// GrantedScopes is the list of scopes granted to the OAuth 2.0 client.
	GrantedScopes []string `json:"granted_scopes,omitempty"`
	// Subject is the identifier of the authenticated end-user.
	Subject *string `json:"subject,omitempty"`
}

// NewRefreshTokenHookRequest instantiates a new RefreshTokenHookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshTokenHookRequest() *RefreshTokenHookRequest {
	this := RefreshTokenHookRequest{}
	return &this
}

// NewRefreshTokenHookRequestWithDefaults instantiates a new RefreshTokenHookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshTokenHookRequestWithDefaults() *RefreshTokenHookRequest {
	this := RefreshTokenHookRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *RefreshTokenHookRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshTokenHookRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *RefreshTokenHookRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *RefreshTokenHookRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetGrantedAudience returns the GrantedAudience field value if set, zero value otherwise.
func (o *RefreshTokenHookRequest) GetGrantedAudience() []string {
	if o == nil || o.GrantedAudience == nil {
		var ret []string
		return ret
	}
	return o.GrantedAudience
}

// GetGrantedAudienceOk returns a tuple with the GrantedAudience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshTokenHookRequest) GetGrantedAudienceOk() ([]string, bool) {
	if o == nil || o.GrantedAudience == nil {
		return nil, false
	}
	return o.GrantedAudience, true
}

// HasGrantedAudience returns a boolean if a field has been set.
func (o *RefreshTokenHookRequest) HasGrantedAudience() bool {
	if o != nil && o.GrantedAudience != nil {
		return true
	}

	return false
}

// SetGrantedAudience gets a reference to the given []string and assigns it to the GrantedAudience field.
func (o *RefreshTokenHookRequest) SetGrantedAudience(v []string) {
	o.GrantedAudience = v
}

// GetGrantedScopes returns the GrantedScopes field value if set, zero value otherwise.
func (o *RefreshTokenHookRequest) GetGrantedScopes() []string {
	if o == nil || o.GrantedScopes == nil {
		var ret []string
		return ret
	}
	return o.GrantedScopes
}

// GetGrantedScopesOk returns a tuple with the GrantedScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshTokenHookRequest) GetGrantedScopesOk() ([]string, bool) {
	if o == nil || o.GrantedScopes == nil {
		return nil, false
	}
	return o.GrantedScopes, true
}

// HasGrantedScopes returns a boolean if a field has been set.
func (o *RefreshTokenHookRequest) HasGrantedScopes() bool {
	if o != nil && o.GrantedScopes != nil {
		return true
	}

	return false
}

// SetGrantedScopes gets a reference to the given []string and assigns it to the GrantedScopes field.
func (o *RefreshTokenHookRequest) SetGrantedScopes(v []string) {
	o.GrantedScopes = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *RefreshTokenHookRequest) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshTokenHookRequest) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *RefreshTokenHookRequest) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *RefreshTokenHookRequest) SetSubject(v string) {
	o.Subject = &v
}

func (o RefreshTokenHookRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.GrantedAudience != nil {
		toSerialize["granted_audience"] = o.GrantedAudience
	}
	if o.GrantedScopes != nil {
		toSerialize["granted_scopes"] = o.GrantedScopes
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	return json.Marshal(toSerialize)
}

type NullableRefreshTokenHookRequest struct {
	value *RefreshTokenHookRequest
	isSet bool
}

func (v NullableRefreshTokenHookRequest) Get() *RefreshTokenHookRequest {
	return v.value
}

func (v *NullableRefreshTokenHookRequest) Set(val *RefreshTokenHookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshTokenHookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshTokenHookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshTokenHookRequest(val *RefreshTokenHookRequest) *NullableRefreshTokenHookRequest {
	return &NullableRefreshTokenHookRequest{value: val, isSet: true}
}

func (v NullableRefreshTokenHookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshTokenHookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


