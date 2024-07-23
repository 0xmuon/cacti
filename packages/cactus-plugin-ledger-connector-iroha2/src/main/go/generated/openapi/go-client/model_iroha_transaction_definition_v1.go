/*
Hyperledger Cactus Plugin - Connector Iroha V2

Can perform basic tasks on a Iroha V2 ledger

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-iroha2

import (
	"encoding/json"
)

// checks if the IrohaTransactionDefinitionV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IrohaTransactionDefinitionV1{}

// IrohaTransactionDefinitionV1 Iroha V2 transaction definition
type IrohaTransactionDefinitionV1 struct {
	Instruction IrohaTransactionDefinitionV1Instruction `json:"instruction"`
	Params *IrohaTransactionParametersV1 `json:"params,omitempty"`
}

// NewIrohaTransactionDefinitionV1 instantiates a new IrohaTransactionDefinitionV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIrohaTransactionDefinitionV1(instruction IrohaTransactionDefinitionV1Instruction) *IrohaTransactionDefinitionV1 {
	this := IrohaTransactionDefinitionV1{}
	this.Instruction = instruction
	return &this
}

// NewIrohaTransactionDefinitionV1WithDefaults instantiates a new IrohaTransactionDefinitionV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIrohaTransactionDefinitionV1WithDefaults() *IrohaTransactionDefinitionV1 {
	this := IrohaTransactionDefinitionV1{}
	return &this
}

// GetInstruction returns the Instruction field value
func (o *IrohaTransactionDefinitionV1) GetInstruction() IrohaTransactionDefinitionV1Instruction {
	if o == nil {
		var ret IrohaTransactionDefinitionV1Instruction
		return ret
	}

	return o.Instruction
}

// GetInstructionOk returns a tuple with the Instruction field value
// and a boolean to check if the value has been set.
func (o *IrohaTransactionDefinitionV1) GetInstructionOk() (*IrohaTransactionDefinitionV1Instruction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instruction, true
}

// SetInstruction sets field value
func (o *IrohaTransactionDefinitionV1) SetInstruction(v IrohaTransactionDefinitionV1Instruction) {
	o.Instruction = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *IrohaTransactionDefinitionV1) GetParams() IrohaTransactionParametersV1 {
	if o == nil || IsNil(o.Params) {
		var ret IrohaTransactionParametersV1
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IrohaTransactionDefinitionV1) GetParamsOk() (*IrohaTransactionParametersV1, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *IrohaTransactionDefinitionV1) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given IrohaTransactionParametersV1 and assigns it to the Params field.
func (o *IrohaTransactionDefinitionV1) SetParams(v IrohaTransactionParametersV1) {
	o.Params = &v
}

func (o IrohaTransactionDefinitionV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IrohaTransactionDefinitionV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["instruction"] = o.Instruction
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	return toSerialize, nil
}

type NullableIrohaTransactionDefinitionV1 struct {
	value *IrohaTransactionDefinitionV1
	isSet bool
}

func (v NullableIrohaTransactionDefinitionV1) Get() *IrohaTransactionDefinitionV1 {
	return v.value
}

func (v *NullableIrohaTransactionDefinitionV1) Set(val *IrohaTransactionDefinitionV1) {
	v.value = val
	v.isSet = true
}

func (v NullableIrohaTransactionDefinitionV1) IsSet() bool {
	return v.isSet
}

func (v *NullableIrohaTransactionDefinitionV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIrohaTransactionDefinitionV1(val *IrohaTransactionDefinitionV1) *NullableIrohaTransactionDefinitionV1 {
	return &NullableIrohaTransactionDefinitionV1{value: val, isSet: true}
}

func (v NullableIrohaTransactionDefinitionV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIrohaTransactionDefinitionV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

