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

// GetChildrenResponseDataChildren struct for GetChildrenResponseDataChildren
type GetChildrenResponseDataChildren struct {
	AllowIndexThreeParameters *bool `json:"allow_index_three_parameters,omitempty"`
	Attach *string `json:"attach,omitempty"`
	ChildId *float64 `json:"child_id,omitempty"`
	Count int64 `json:"count"`
	CountChildren float64 `json:"count_children"`
	GoodsCount int64 `json:"goods_count"`
	GoodsId *int64 `json:"goods_id,omitempty"`
	Href *string `json:"href,omitempty"`
	Id int64 `json:"id"`
	IsAppendGoods *bool `json:"is_append_goods,omitempty"`
	IsRozetkaTop *bool `json:"is_rozetka_top,omitempty"`
	LeftKey *int64 `json:"left_key,omitempty"`
	Level *float64 `json:"level,omitempty"`
	Mpath *string `json:"mpath,omitempty"`
	Name string `json:"name"`
	ParentId int64 `json:"parent_id"`
	RightKey *int64 `json:"right_key,omitempty"`
	RzMpath *string `json:"rz_mpath,omitempty"`
	Status *string `json:"status,omitempty"`
	Title string `json:"title"`
	TitlesMode *string `json:"titles_mode,omitempty"`
}

// NewGetChildrenResponseDataChildren instantiates a new GetChildrenResponseDataChildren object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetChildrenResponseDataChildren(count int64, countChildren float64, goodsCount int64, id int64, name string, parentId int64, title string) *GetChildrenResponseDataChildren {
	this := GetChildrenResponseDataChildren{}
	this.Count = count
	this.CountChildren = countChildren
	this.GoodsCount = goodsCount
	this.Id = id
	this.Name = name
	this.ParentId = parentId
	this.Title = title
	return &this
}

// NewGetChildrenResponseDataChildrenWithDefaults instantiates a new GetChildrenResponseDataChildren object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetChildrenResponseDataChildrenWithDefaults() *GetChildrenResponseDataChildren {
	this := GetChildrenResponseDataChildren{}
	return &this
}

// GetAllowIndexThreeParameters returns the AllowIndexThreeParameters field value if set, zero value otherwise.
func (o *GetChildrenResponseDataChildren) GetAllowIndexThreeParameters() bool {
	if o == nil || o.AllowIndexThreeParameters == nil {
		var ret bool
		return ret
	}
	return *o.AllowIndexThreeParameters
}

// GetAllowIndexThreeParametersOk returns a tuple with the AllowIndexThreeParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetAllowIndexThreeParametersOk() (*bool, bool) {
	if o == nil || o.AllowIndexThreeParameters == nil {
		return nil, false
	}
	return o.AllowIndexThreeParameters, true
}

// HasAllowIndexThreeParameters returns a boolean if a field has been set.
func (o *GetChildrenResponseDataChildren) HasAllowIndexThreeParameters() bool {
	if o != nil && o.AllowIndexThreeParameters != nil {
		return true
	}

	return false
}

// SetAllowIndexThreeParameters gets a reference to the given bool and assigns it to the AllowIndexThreeParameters field.
func (o *GetChildrenResponseDataChildren) SetAllowIndexThreeParameters(v bool) {
	o.AllowIndexThreeParameters = &v
}

// GetAttach returns the Attach field value if set, zero value otherwise.
func (o *GetChildrenResponseDataChildren) GetAttach() string {
	if o == nil || o.Attach == nil {
		var ret string
		return ret
	}
	return *o.Attach
}

// GetAttachOk returns a tuple with the Attach field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetAttachOk() (*string, bool) {
	if o == nil || o.Attach == nil {
		return nil, false
	}
	return o.Attach, true
}

// HasAttach returns a boolean if a field has been set.
func (o *GetChildrenResponseDataChildren) HasAttach() bool {
	if o != nil && o.Attach != nil {
		return true
	}

	return false
}

// SetAttach gets a reference to the given string and assigns it to the Attach field.
func (o *GetChildrenResponseDataChildren) SetAttach(v string) {
	o.Attach = &v
}

// GetChildId returns the ChildId field value if set, zero value otherwise.
func (o *GetChildrenResponseDataChildren) GetChildId() float64 {
	if o == nil || o.ChildId == nil {
		var ret float64
		return ret
	}
	return *o.ChildId
}

// GetChildIdOk returns a tuple with the ChildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetChildIdOk() (*float64, bool) {
	if o == nil || o.ChildId == nil {
		return nil, false
	}
	return o.ChildId, true
}

// HasChildId returns a boolean if a field has been set.
func (o *GetChildrenResponseDataChildren) HasChildId() bool {
	if o != nil && o.ChildId != nil {
		return true
	}

	return false
}

// SetChildId gets a reference to the given float64 and assigns it to the ChildId field.
func (o *GetChildrenResponseDataChildren) SetChildId(v float64) {
	o.ChildId = &v
}

// GetCount returns the Count field value
func (o *GetChildrenResponseDataChildren) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *GetChildrenResponseDataChildren) SetCount(v int64) {
	o.Count = v
}

// GetCountChildren returns the CountChildren field value
func (o *GetChildrenResponseDataChildren) GetCountChildren() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.CountChildren
}

// GetCountChildrenOk returns a tuple with the CountChildren field value
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetCountChildrenOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CountChildren, true
}

// SetCountChildren sets field value
func (o *GetChildrenResponseDataChildren) SetCountChildren(v float64) {
	o.CountChildren = v
}

// GetGoodsCount returns the GoodsCount field value
func (o *GetChildrenResponseDataChildren) GetGoodsCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.GoodsCount
}

// GetGoodsCountOk returns a tuple with the GoodsCount field value
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetGoodsCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GoodsCount, true
}

// SetGoodsCount sets field value
func (o *GetChildrenResponseDataChildren) SetGoodsCount(v int64) {
	o.GoodsCount = v
}

// GetGoodsId returns the GoodsId field value if set, zero value otherwise.
func (o *GetChildrenResponseDataChildren) GetGoodsId() int64 {
	if o == nil || o.GoodsId == nil {
		var ret int64
		return ret
	}
	return *o.GoodsId
}

// GetGoodsIdOk returns a tuple with the GoodsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetGoodsIdOk() (*int64, bool) {
	if o == nil || o.GoodsId == nil {
		return nil, false
	}
	return o.GoodsId, true
}

// HasGoodsId returns a boolean if a field has been set.
func (o *GetChildrenResponseDataChildren) HasGoodsId() bool {
	if o != nil && o.GoodsId != nil {
		return true
	}

	return false
}

// SetGoodsId gets a reference to the given int64 and assigns it to the GoodsId field.
func (o *GetChildrenResponseDataChildren) SetGoodsId(v int64) {
	o.GoodsId = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *GetChildrenResponseDataChildren) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *GetChildrenResponseDataChildren) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *GetChildrenResponseDataChildren) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value
func (o *GetChildrenResponseDataChildren) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetChildrenResponseDataChildren) SetId(v int64) {
	o.Id = v
}

// GetIsAppendGoods returns the IsAppendGoods field value if set, zero value otherwise.
func (o *GetChildrenResponseDataChildren) GetIsAppendGoods() bool {
	if o == nil || o.IsAppendGoods == nil {
		var ret bool
		return ret
	}
	return *o.IsAppendGoods
}

// GetIsAppendGoodsOk returns a tuple with the IsAppendGoods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetIsAppendGoodsOk() (*bool, bool) {
	if o == nil || o.IsAppendGoods == nil {
		return nil, false
	}
	return o.IsAppendGoods, true
}

// HasIsAppendGoods returns a boolean if a field has been set.
func (o *GetChildrenResponseDataChildren) HasIsAppendGoods() bool {
	if o != nil && o.IsAppendGoods != nil {
		return true
	}

	return false
}

// SetIsAppendGoods gets a reference to the given bool and assigns it to the IsAppendGoods field.
func (o *GetChildrenResponseDataChildren) SetIsAppendGoods(v bool) {
	o.IsAppendGoods = &v
}

// GetIsRozetkaTop returns the IsRozetkaTop field value if set, zero value otherwise.
func (o *GetChildrenResponseDataChildren) GetIsRozetkaTop() bool {
	if o == nil || o.IsRozetkaTop == nil {
		var ret bool
		return ret
	}
	return *o.IsRozetkaTop
}

// GetIsRozetkaTopOk returns a tuple with the IsRozetkaTop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetIsRozetkaTopOk() (*bool, bool) {
	if o == nil || o.IsRozetkaTop == nil {
		return nil, false
	}
	return o.IsRozetkaTop, true
}

// HasIsRozetkaTop returns a boolean if a field has been set.
func (o *GetChildrenResponseDataChildren) HasIsRozetkaTop() bool {
	if o != nil && o.IsRozetkaTop != nil {
		return true
	}

	return false
}

// SetIsRozetkaTop gets a reference to the given bool and assigns it to the IsRozetkaTop field.
func (o *GetChildrenResponseDataChildren) SetIsRozetkaTop(v bool) {
	o.IsRozetkaTop = &v
}

// GetLeftKey returns the LeftKey field value if set, zero value otherwise.
func (o *GetChildrenResponseDataChildren) GetLeftKey() int64 {
	if o == nil || o.LeftKey == nil {
		var ret int64
		return ret
	}
	return *o.LeftKey
}

// GetLeftKeyOk returns a tuple with the LeftKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetLeftKeyOk() (*int64, bool) {
	if o == nil || o.LeftKey == nil {
		return nil, false
	}
	return o.LeftKey, true
}

// HasLeftKey returns a boolean if a field has been set.
func (o *GetChildrenResponseDataChildren) HasLeftKey() bool {
	if o != nil && o.LeftKey != nil {
		return true
	}

	return false
}

// SetLeftKey gets a reference to the given int64 and assigns it to the LeftKey field.
func (o *GetChildrenResponseDataChildren) SetLeftKey(v int64) {
	o.LeftKey = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *GetChildrenResponseDataChildren) GetLevel() float64 {
	if o == nil || o.Level == nil {
		var ret float64
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetLevelOk() (*float64, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *GetChildrenResponseDataChildren) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given float64 and assigns it to the Level field.
func (o *GetChildrenResponseDataChildren) SetLevel(v float64) {
	o.Level = &v
}

// GetMpath returns the Mpath field value if set, zero value otherwise.
func (o *GetChildrenResponseDataChildren) GetMpath() string {
	if o == nil || o.Mpath == nil {
		var ret string
		return ret
	}
	return *o.Mpath
}

// GetMpathOk returns a tuple with the Mpath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetMpathOk() (*string, bool) {
	if o == nil || o.Mpath == nil {
		return nil, false
	}
	return o.Mpath, true
}

// HasMpath returns a boolean if a field has been set.
func (o *GetChildrenResponseDataChildren) HasMpath() bool {
	if o != nil && o.Mpath != nil {
		return true
	}

	return false
}

// SetMpath gets a reference to the given string and assigns it to the Mpath field.
func (o *GetChildrenResponseDataChildren) SetMpath(v string) {
	o.Mpath = &v
}

// GetName returns the Name field value
func (o *GetChildrenResponseDataChildren) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetChildrenResponseDataChildren) SetName(v string) {
	o.Name = v
}

// GetParentId returns the ParentId field value
func (o *GetChildrenResponseDataChildren) GetParentId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetParentIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *GetChildrenResponseDataChildren) SetParentId(v int64) {
	o.ParentId = v
}

// GetRightKey returns the RightKey field value if set, zero value otherwise.
func (o *GetChildrenResponseDataChildren) GetRightKey() int64 {
	if o == nil || o.RightKey == nil {
		var ret int64
		return ret
	}
	return *o.RightKey
}

// GetRightKeyOk returns a tuple with the RightKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetRightKeyOk() (*int64, bool) {
	if o == nil || o.RightKey == nil {
		return nil, false
	}
	return o.RightKey, true
}

// HasRightKey returns a boolean if a field has been set.
func (o *GetChildrenResponseDataChildren) HasRightKey() bool {
	if o != nil && o.RightKey != nil {
		return true
	}

	return false
}

// SetRightKey gets a reference to the given int64 and assigns it to the RightKey field.
func (o *GetChildrenResponseDataChildren) SetRightKey(v int64) {
	o.RightKey = &v
}

// GetRzMpath returns the RzMpath field value if set, zero value otherwise.
func (o *GetChildrenResponseDataChildren) GetRzMpath() string {
	if o == nil || o.RzMpath == nil {
		var ret string
		return ret
	}
	return *o.RzMpath
}

// GetRzMpathOk returns a tuple with the RzMpath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetRzMpathOk() (*string, bool) {
	if o == nil || o.RzMpath == nil {
		return nil, false
	}
	return o.RzMpath, true
}

// HasRzMpath returns a boolean if a field has been set.
func (o *GetChildrenResponseDataChildren) HasRzMpath() bool {
	if o != nil && o.RzMpath != nil {
		return true
	}

	return false
}

// SetRzMpath gets a reference to the given string and assigns it to the RzMpath field.
func (o *GetChildrenResponseDataChildren) SetRzMpath(v string) {
	o.RzMpath = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetChildrenResponseDataChildren) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetChildrenResponseDataChildren) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetChildrenResponseDataChildren) SetStatus(v string) {
	o.Status = &v
}

// GetTitle returns the Title field value
func (o *GetChildrenResponseDataChildren) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *GetChildrenResponseDataChildren) SetTitle(v string) {
	o.Title = v
}

// GetTitlesMode returns the TitlesMode field value if set, zero value otherwise.
func (o *GetChildrenResponseDataChildren) GetTitlesMode() string {
	if o == nil || o.TitlesMode == nil {
		var ret string
		return ret
	}
	return *o.TitlesMode
}

// GetTitlesModeOk returns a tuple with the TitlesMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChildrenResponseDataChildren) GetTitlesModeOk() (*string, bool) {
	if o == nil || o.TitlesMode == nil {
		return nil, false
	}
	return o.TitlesMode, true
}

// HasTitlesMode returns a boolean if a field has been set.
func (o *GetChildrenResponseDataChildren) HasTitlesMode() bool {
	if o != nil && o.TitlesMode != nil {
		return true
	}

	return false
}

// SetTitlesMode gets a reference to the given string and assigns it to the TitlesMode field.
func (o *GetChildrenResponseDataChildren) SetTitlesMode(v string) {
	o.TitlesMode = &v
}

func (o GetChildrenResponseDataChildren) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowIndexThreeParameters != nil {
		toSerialize["allow_index_three_parameters"] = o.AllowIndexThreeParameters
	}
	if o.Attach != nil {
		toSerialize["attach"] = o.Attach
	}
	if o.ChildId != nil {
		toSerialize["child_id"] = o.ChildId
	}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["count_children"] = o.CountChildren
	}
	if true {
		toSerialize["goods_count"] = o.GoodsCount
	}
	if o.GoodsId != nil {
		toSerialize["goods_id"] = o.GoodsId
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.IsAppendGoods != nil {
		toSerialize["is_append_goods"] = o.IsAppendGoods
	}
	if o.IsRozetkaTop != nil {
		toSerialize["is_rozetka_top"] = o.IsRozetkaTop
	}
	if o.LeftKey != nil {
		toSerialize["left_key"] = o.LeftKey
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.Mpath != nil {
		toSerialize["mpath"] = o.Mpath
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["parent_id"] = o.ParentId
	}
	if o.RightKey != nil {
		toSerialize["right_key"] = o.RightKey
	}
	if o.RzMpath != nil {
		toSerialize["rz_mpath"] = o.RzMpath
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if o.TitlesMode != nil {
		toSerialize["titles_mode"] = o.TitlesMode
	}
	return json.Marshal(toSerialize)
}

type NullableGetChildrenResponseDataChildren struct {
	value *GetChildrenResponseDataChildren
	isSet bool
}

func (v NullableGetChildrenResponseDataChildren) Get() *GetChildrenResponseDataChildren {
	return v.value
}

func (v *NullableGetChildrenResponseDataChildren) Set(val *GetChildrenResponseDataChildren) {
	v.value = val
	v.isSet = true
}

func (v NullableGetChildrenResponseDataChildren) IsSet() bool {
	return v.isSet
}

func (v *NullableGetChildrenResponseDataChildren) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetChildrenResponseDataChildren(val *GetChildrenResponseDataChildren) *NullableGetChildrenResponseDataChildren {
	return &NullableGetChildrenResponseDataChildren{value: val, isSet: true}
}

func (v NullableGetChildrenResponseDataChildren) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetChildrenResponseDataChildren) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

