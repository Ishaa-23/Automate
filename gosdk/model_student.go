/*
Combined API

Endpoints of CRUD of employee and displaying results of student.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package EmployeeAPI

import (
	"encoding/json"
)

// checks if the Student type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Student{}

// Student struct for Student
type Student struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Age *int32 `json:"age,omitempty"`
}

// NewStudent instantiates a new Student object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStudent() *Student {
	this := Student{}
	return &this
}

// NewStudentWithDefaults instantiates a new Student object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStudentWithDefaults() *Student {
	this := Student{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Student) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Student) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Student) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Student) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Student) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Student) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Student) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Student) SetName(v string) {
	o.Name = &v
}

// GetAge returns the Age field value if set, zero value otherwise.
func (o *Student) GetAge() int32 {
	if o == nil || IsNil(o.Age) {
		var ret int32
		return ret
	}
	return *o.Age
}

// GetAgeOk returns a tuple with the Age field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Student) GetAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.Age) {
		return nil, false
	}
	return o.Age, true
}

// HasAge returns a boolean if a field has been set.
func (o *Student) HasAge() bool {
	if o != nil && !IsNil(o.Age) {
		return true
	}

	return false
}

// SetAge gets a reference to the given int32 and assigns it to the Age field.
func (o *Student) SetAge(v int32) {
	o.Age = &v
}

func (o Student) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Student) ToMap() (map[string]interface{}, error) {
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

type NullableStudent struct {
	value *Student
	isSet bool
}

func (v NullableStudent) Get() *Student {
	return v.value
}

func (v *NullableStudent) Set(val *Student) {
	v.value = val
	v.isSet = true
}

func (v NullableStudent) IsSet() bool {
	return v.isSet
}

func (v *NullableStudent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStudent(val *Student) *NullableStudent {
	return &NullableStudent{value: val, isSet: true}
}

func (v NullableStudent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStudent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

