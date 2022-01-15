/*
api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetProductsInCategoryResponseDataAlternate struct for GetProductsInCategoryResponseDataAlternate
type GetProductsInCategoryResponseDataAlternate struct {
	Domain *string `json:"domain,omitempty"`
	Hreflang *string `json:"hreflang,omitempty"`
	Lang *string `json:"lang,omitempty"`
	Subdomain *string `json:"subdomain,omitempty"`
}

// NewGetProductsInCategoryResponseDataAlternate instantiates a new GetProductsInCategoryResponseDataAlternate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductsInCategoryResponseDataAlternate() *GetProductsInCategoryResponseDataAlternate {
	this := GetProductsInCategoryResponseDataAlternate{}
	return &this
}

// NewGetProductsInCategoryResponseDataAlternateWithDefaults instantiates a new GetProductsInCategoryResponseDataAlternate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductsInCategoryResponseDataAlternateWithDefaults() *GetProductsInCategoryResponseDataAlternate {
	this := GetProductsInCategoryResponseDataAlternate{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *GetProductsInCategoryResponseDataAlternate) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductsInCategoryResponseDataAlternate) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *GetProductsInCategoryResponseDataAlternate) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *GetProductsInCategoryResponseDataAlternate) SetDomain(v string) {
	o.Domain = &v
}

// GetHreflang returns the Hreflang field value if set, zero value otherwise.
func (o *GetProductsInCategoryResponseDataAlternate) GetHreflang() string {
	if o == nil || o.Hreflang == nil {
		var ret string
		return ret
	}
	return *o.Hreflang
}

// GetHreflangOk returns a tuple with the Hreflang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductsInCategoryResponseDataAlternate) GetHreflangOk() (*string, bool) {
	if o == nil || o.Hreflang == nil {
		return nil, false
	}
	return o.Hreflang, true
}

// HasHreflang returns a boolean if a field has been set.
func (o *GetProductsInCategoryResponseDataAlternate) HasHreflang() bool {
	if o != nil && o.Hreflang != nil {
		return true
	}

	return false
}

// SetHreflang gets a reference to the given string and assigns it to the Hreflang field.
func (o *GetProductsInCategoryResponseDataAlternate) SetHreflang(v string) {
	o.Hreflang = &v
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *GetProductsInCategoryResponseDataAlternate) GetLang() string {
	if o == nil || o.Lang == nil {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductsInCategoryResponseDataAlternate) GetLangOk() (*string, bool) {
	if o == nil || o.Lang == nil {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *GetProductsInCategoryResponseDataAlternate) HasLang() bool {
	if o != nil && o.Lang != nil {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *GetProductsInCategoryResponseDataAlternate) SetLang(v string) {
	o.Lang = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *GetProductsInCategoryResponseDataAlternate) GetSubdomain() string {
	if o == nil || o.Subdomain == nil {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductsInCategoryResponseDataAlternate) GetSubdomainOk() (*string, bool) {
	if o == nil || o.Subdomain == nil {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *GetProductsInCategoryResponseDataAlternate) HasSubdomain() bool {
	if o != nil && o.Subdomain != nil {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *GetProductsInCategoryResponseDataAlternate) SetSubdomain(v string) {
	o.Subdomain = &v
}

func (o GetProductsInCategoryResponseDataAlternate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Hreflang != nil {
		toSerialize["hreflang"] = o.Hreflang
	}
	if o.Lang != nil {
		toSerialize["lang"] = o.Lang
	}
	if o.Subdomain != nil {
		toSerialize["subdomain"] = o.Subdomain
	}
	return json.Marshal(toSerialize)
}

type NullableGetProductsInCategoryResponseDataAlternate struct {
	value *GetProductsInCategoryResponseDataAlternate
	isSet bool
}

func (v NullableGetProductsInCategoryResponseDataAlternate) Get() *GetProductsInCategoryResponseDataAlternate {
	return v.value
}

func (v *NullableGetProductsInCategoryResponseDataAlternate) Set(val *GetProductsInCategoryResponseDataAlternate) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductsInCategoryResponseDataAlternate) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductsInCategoryResponseDataAlternate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductsInCategoryResponseDataAlternate(val *GetProductsInCategoryResponseDataAlternate) *NullableGetProductsInCategoryResponseDataAlternate {
	return &NullableGetProductsInCategoryResponseDataAlternate{value: val, isSet: true}
}

func (v NullableGetProductsInCategoryResponseDataAlternate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductsInCategoryResponseDataAlternate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

