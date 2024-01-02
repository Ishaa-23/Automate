# Student

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Age** | Pointer to **int32** |  | [optional] 

## Methods

### NewStudent

`func NewStudent() *Student`

NewStudent instantiates a new Student object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudentWithDefaults

`func NewStudentWithDefaults() *Student`

NewStudentWithDefaults instantiates a new Student object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Student) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Student) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Student) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Student) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Student) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Student) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Student) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Student) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAge

`func (o *Student) GetAge() int32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Student) GetAgeOk() (*int32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Student) SetAge(v int32)`

SetAge sets Age field to given value.

### HasAge

`func (o *Student) HasAge() bool`

HasAge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


