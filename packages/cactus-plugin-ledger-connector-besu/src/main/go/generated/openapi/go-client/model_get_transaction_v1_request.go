/*
Hyperledger Cactus Plugin - Connector Besu

Can perform basic tasks on a Besu ledger

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cactus-plugin-ledger-connector-besu

import (
	"encoding/json"
)

// checks if the GetTransactionV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTransactionV1Request{}

// GetTransactionV1Request struct for GetTransactionV1Request
type GetTransactionV1Request struct {
	TransactionHash string `json:"transactionHash"`
}

// NewGetTransactionV1Request instantiates a new GetTransactionV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransactionV1Request(transactionHash string) *GetTransactionV1Request {
	this := GetTransactionV1Request{}
	this.TransactionHash = transactionHash
	return &this
}

// NewGetTransactionV1RequestWithDefaults instantiates a new GetTransactionV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransactionV1RequestWithDefaults() *GetTransactionV1Request {
	this := GetTransactionV1Request{}
	return &this
}

// GetTransactionHash returns the TransactionHash field value
func (o *GetTransactionV1Request) GetTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value
// and a boolean to check if the value has been set.
func (o *GetTransactionV1Request) GetTransactionHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionHash, true
}

// SetTransactionHash sets field value
func (o *GetTransactionV1Request) SetTransactionHash(v string) {
	o.TransactionHash = v
}

func (o GetTransactionV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTransactionV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transactionHash"] = o.TransactionHash
	return toSerialize, nil
}

type NullableGetTransactionV1Request struct {
	value *GetTransactionV1Request
	isSet bool
}

func (v NullableGetTransactionV1Request) Get() *GetTransactionV1Request {
	return v.value
}

func (v *NullableGetTransactionV1Request) Set(val *GetTransactionV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransactionV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransactionV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransactionV1Request(val *GetTransactionV1Request) *NullableGetTransactionV1Request {
	return &NullableGetTransactionV1Request{value: val, isSet: true}
}

func (v NullableGetTransactionV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransactionV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

