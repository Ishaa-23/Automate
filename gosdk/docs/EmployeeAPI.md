# \EmployeeAPI

All URIs are relative to *http://localhost:5283*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiEmployeeIdDelete**](EmployeeAPI.md#ApiEmployeeIdDelete) | **Delete** /api/Employee/{id} | 
[**ApiEmployeeIdGet**](EmployeeAPI.md#ApiEmployeeIdGet) | **Get** /api/Employee/{id} | Returns employee of given id
[**ApiEmployeeIdPut**](EmployeeAPI.md#ApiEmployeeIdPut) | **Put** /api/Employee/{id} | Updates employee details
[**ApiEmployeePost**](EmployeeAPI.md#ApiEmployeePost) | **Post** /api/Employee | Adds new employee



## ApiEmployeeIdDelete

> []Employee ApiEmployeeIdDelete(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmployeeAPI.ApiEmployeeIdDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmployeeAPI.ApiEmployeeIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEmployeeIdDelete`: []Employee
    fmt.Fprintf(os.Stdout, "Response from `EmployeeAPI.ApiEmployeeIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEmployeeIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Employee**](Employee.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEmployeeIdGet

> []Employee ApiEmployeeIdGet(ctx, id).Execute()

Returns employee of given id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmployeeAPI.ApiEmployeeIdGet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmployeeAPI.ApiEmployeeIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEmployeeIdGet`: []Employee
    fmt.Fprintf(os.Stdout, "Response from `EmployeeAPI.ApiEmployeeIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEmployeeIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Employee**](Employee.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEmployeeIdPut

> []Employee ApiEmployeeIdPut(ctx, id).Employee(employee).Execute()

Updates employee details

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := int32(56) // int32 | 
    employee := []openapiclient.Employee{*openapiclient.NewEmployee()} // []Employee |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmployeeAPI.ApiEmployeeIdPut(context.Background(), id).Employee(employee).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmployeeAPI.ApiEmployeeIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEmployeeIdPut`: []Employee
    fmt.Fprintf(os.Stdout, "Response from `EmployeeAPI.ApiEmployeeIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiEmployeeIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **employee** | [**[]Employee**](Employee.md) |  | 

### Return type

[**[]Employee**](Employee.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain, application/json, text/json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiEmployeePost

> []Employee ApiEmployeePost(ctx).Employee(employee).Execute()

Adds new employee

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    employee := *openapiclient.NewEmployee() // Employee |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmployeeAPI.ApiEmployeePost(context.Background()).Employee(employee).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmployeeAPI.ApiEmployeePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiEmployeePost`: []Employee
    fmt.Fprintf(os.Stdout, "Response from `EmployeeAPI.ApiEmployeePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiEmployeePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **employee** | [**Employee**](Employee.md) |  | 

### Return type

[**[]Employee**](Employee.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

