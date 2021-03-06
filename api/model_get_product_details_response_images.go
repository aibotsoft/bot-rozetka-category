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

// GetProductDetailsResponseImages struct for GetProductDetailsResponseImages
type GetProductDetailsResponseImages struct {
	AllImages *[]string `json:"all_images,omitempty"`
	Hover *string `json:"hover,omitempty"`
	Main *string `json:"main,omitempty"`
	Preview *string `json:"preview,omitempty"`
}

// NewGetProductDetailsResponseImages instantiates a new GetProductDetailsResponseImages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductDetailsResponseImages() *GetProductDetailsResponseImages {
	this := GetProductDetailsResponseImages{}
	return &this
}

// NewGetProductDetailsResponseImagesWithDefaults instantiates a new GetProductDetailsResponseImages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductDetailsResponseImagesWithDefaults() *GetProductDetailsResponseImages {
	this := GetProductDetailsResponseImages{}
	return &this
}

// GetAllImages returns the AllImages field value if set, zero value otherwise.
func (o *GetProductDetailsResponseImages) GetAllImages() []string {
	if o == nil || o.AllImages == nil {
		var ret []string
		return ret
	}
	return *o.AllImages
}

// GetAllImagesOk returns a tuple with the AllImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductDetailsResponseImages) GetAllImagesOk() (*[]string, bool) {
	if o == nil || o.AllImages == nil {
		return nil, false
	}
	return o.AllImages, true
}

// HasAllImages returns a boolean if a field has been set.
func (o *GetProductDetailsResponseImages) HasAllImages() bool {
	if o != nil && o.AllImages != nil {
		return true
	}

	return false
}

// SetAllImages gets a reference to the given []string and assigns it to the AllImages field.
func (o *GetProductDetailsResponseImages) SetAllImages(v []string) {
	o.AllImages = &v
}

// GetHover returns the Hover field value if set, zero value otherwise.
func (o *GetProductDetailsResponseImages) GetHover() string {
	if o == nil || o.Hover == nil {
		var ret string
		return ret
	}
	return *o.Hover
}

// GetHoverOk returns a tuple with the Hover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductDetailsResponseImages) GetHoverOk() (*string, bool) {
	if o == nil || o.Hover == nil {
		return nil, false
	}
	return o.Hover, true
}

// HasHover returns a boolean if a field has been set.
func (o *GetProductDetailsResponseImages) HasHover() bool {
	if o != nil && o.Hover != nil {
		return true
	}

	return false
}

// SetHover gets a reference to the given string and assigns it to the Hover field.
func (o *GetProductDetailsResponseImages) SetHover(v string) {
	o.Hover = &v
}

// GetMain returns the Main field value if set, zero value otherwise.
func (o *GetProductDetailsResponseImages) GetMain() string {
	if o == nil || o.Main == nil {
		var ret string
		return ret
	}
	return *o.Main
}

// GetMainOk returns a tuple with the Main field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductDetailsResponseImages) GetMainOk() (*string, bool) {
	if o == nil || o.Main == nil {
		return nil, false
	}
	return o.Main, true
}

// HasMain returns a boolean if a field has been set.
func (o *GetProductDetailsResponseImages) HasMain() bool {
	if o != nil && o.Main != nil {
		return true
	}

	return false
}

// SetMain gets a reference to the given string and assigns it to the Main field.
func (o *GetProductDetailsResponseImages) SetMain(v string) {
	o.Main = &v
}

// GetPreview returns the Preview field value if set, zero value otherwise.
func (o *GetProductDetailsResponseImages) GetPreview() string {
	if o == nil || o.Preview == nil {
		var ret string
		return ret
	}
	return *o.Preview
}

// GetPreviewOk returns a tuple with the Preview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductDetailsResponseImages) GetPreviewOk() (*string, bool) {
	if o == nil || o.Preview == nil {
		return nil, false
	}
	return o.Preview, true
}

// HasPreview returns a boolean if a field has been set.
func (o *GetProductDetailsResponseImages) HasPreview() bool {
	if o != nil && o.Preview != nil {
		return true
	}

	return false
}

// SetPreview gets a reference to the given string and assigns it to the Preview field.
func (o *GetProductDetailsResponseImages) SetPreview(v string) {
	o.Preview = &v
}

func (o GetProductDetailsResponseImages) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllImages != nil {
		toSerialize["all_images"] = o.AllImages
	}
	if o.Hover != nil {
		toSerialize["hover"] = o.Hover
	}
	if o.Main != nil {
		toSerialize["main"] = o.Main
	}
	if o.Preview != nil {
		toSerialize["preview"] = o.Preview
	}
	return json.Marshal(toSerialize)
}

type NullableGetProductDetailsResponseImages struct {
	value *GetProductDetailsResponseImages
	isSet bool
}

func (v NullableGetProductDetailsResponseImages) Get() *GetProductDetailsResponseImages {
	return v.value
}

func (v *NullableGetProductDetailsResponseImages) Set(val *GetProductDetailsResponseImages) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductDetailsResponseImages) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductDetailsResponseImages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductDetailsResponseImages(val *GetProductDetailsResponseImages) *NullableGetProductDetailsResponseImages {
	return &NullableGetProductDetailsResponseImages{value: val, isSet: true}
}

func (v NullableGetProductDetailsResponseImages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductDetailsResponseImages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


