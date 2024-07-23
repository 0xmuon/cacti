/*
Hyperledger Cactus Plugin - Object Store - IPFS 

Contains/describes the Hyperledger Cactus Object Store IPFS plugin.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-object-store-ipfs

import (
	"encoding/json"
)

// checks if the SetObjectRequestV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetObjectRequestV1{}

// SetObjectRequestV1 struct for SetObjectRequestV1
type SetObjectRequestV1 struct {
	// The key for the entry to set in the object store.
	Key string `json:"key"`
	// The value that will be associated with the key in the object store.
	Value string `json:"value"`
}

// NewSetObjectRequestV1 instantiates a new SetObjectRequestV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetObjectRequestV1(key string, value string) *SetObjectRequestV1 {
	this := SetObjectRequestV1{}
	this.Key = key
	this.Value = value
	return &this
}

// NewSetObjectRequestV1WithDefaults instantiates a new SetObjectRequestV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetObjectRequestV1WithDefaults() *SetObjectRequestV1 {
	this := SetObjectRequestV1{}
	return &this
}

// GetKey returns the Key field value
func (o *SetObjectRequestV1) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SetObjectRequestV1) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SetObjectRequestV1) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value
func (o *SetObjectRequestV1) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SetObjectRequestV1) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SetObjectRequestV1) SetValue(v string) {
	o.Value = v
}

func (o SetObjectRequestV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetObjectRequestV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableSetObjectRequestV1 struct {
	value *SetObjectRequestV1
	isSet bool
}

func (v NullableSetObjectRequestV1) Get() *SetObjectRequestV1 {
	return v.value
}

func (v *NullableSetObjectRequestV1) Set(val *SetObjectRequestV1) {
	v.value = val
	v.isSet = true
}

func (v NullableSetObjectRequestV1) IsSet() bool {
	return v.isSet
}

func (v *NullableSetObjectRequestV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetObjectRequestV1(val *SetObjectRequestV1) *NullableSetObjectRequestV1 {
	return &NullableSetObjectRequestV1{value: val, isSet: true}
}

func (v NullableSetObjectRequestV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetObjectRequestV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

