/*
Employee API

Endpoints of CRUD of employee.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Combinedapi

import (
	"encoding/json"
)

// checks if the Employee type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Employee{}

// Employee struct for Employee
type Employee struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Age *int32 `json:"age,omitempty"`
}

// NewEmployee instantiates a new Employee object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployee() *Employee {
	this := Employee{}
	return &this
}

// NewEmployeeWithDefaults instantiates a new Employee object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployeeWithDefaults() *Employee {
	this := Employee{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Employee) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Employee) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Employee) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Employee) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Employee) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Employee) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Employee) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Employee) SetName(v string) {
	o.Name = &v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *Employee) GetAge() int32 {
	if o == nil || IsNil(o.Age) {
		var ret int32
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Employee) GetAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.Age) {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *Employee) HasAge() bool {
	if o != nil && !IsNil(o.Age) {
		return true
	}

	return false
}

// SetAge gets a reference to the given int32 and assigns it to the Age field.
func (o *Employee) SetAge(v int32) {
	o.Age = &v
}

func (o Employee) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Employee) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Age) {
		toSerialize["age"] = o.Age
	}
	return toSerialize, nil
}

type NullableEmployee struct {
	value *Employee
	isSet bool
}

func (v NullableEmployee) Get() *Employee {
	return v.value
}

func (v *NullableEmployee) Set(val *Employee) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployee) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployee) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployee(val *Employee) *NullableEmployee {
	return &NullableEmployee{value: val, isSet: true}
}

func (v NullableEmployee) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployee) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


