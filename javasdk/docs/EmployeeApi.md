# EmployeeApi

All URIs are relative to *http://localhost:5283*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiEmployeeIdDelete**](EmployeeApi.md#apiEmployeeIdDelete) | **DELETE** /api/Employee/{id} |  |
| [**apiEmployeeIdGet**](EmployeeApi.md#apiEmployeeIdGet) | **GET** /api/Employee/{id} | Returns employee of given id |
| [**apiEmployeeIdPut**](EmployeeApi.md#apiEmployeeIdPut) | **PUT** /api/Employee/{id} | Updates employee details |
| [**apiEmployeePost**](EmployeeApi.md#apiEmployeePost) | **POST** /api/Employee | Adds new employee |


<a id="apiEmployeeIdDelete"></a>
# **apiEmployeeIdDelete**
> List&lt;Employee&gt; apiEmployeeIdDelete(id)



### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.EmployeeApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:5283");

    EmployeeApi apiInstance = new EmployeeApi(defaultClient);
    Integer id = 56; // Integer | 
    try {
      List<Employee> result = apiInstance.apiEmployeeIdDelete(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling EmployeeApi#apiEmployeeIdDelete");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | **Integer**|  | |

### Return type

[**List&lt;Employee&gt;**](Employee.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |

<a id="apiEmployeeIdGet"></a>
# **apiEmployeeIdGet**
> List&lt;Employee&gt; apiEmployeeIdGet(id)

Returns employee of given id

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.EmployeeApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:5283");

    EmployeeApi apiInstance = new EmployeeApi(defaultClient);
    Integer id = 56; // Integer | 
    try {
      List<Employee> result = apiInstance.apiEmployeeIdGet(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling EmployeeApi#apiEmployeeIdGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | **Integer**|  | |

### Return type

[**List&lt;Employee&gt;**](Employee.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |

<a id="apiEmployeeIdPut"></a>
# **apiEmployeeIdPut**
> List&lt;Employee&gt; apiEmployeeIdPut(id, employee)

Updates employee details

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.EmployeeApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:5283");

    EmployeeApi apiInstance = new EmployeeApi(defaultClient);
    Integer id = 56; // Integer | 
    List<Employee> employee = Arrays.asList(); // List<Employee> | 
    try {
      List<Employee> result = apiInstance.apiEmployeeIdPut(id, employee);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling EmployeeApi#apiEmployeeIdPut");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | **Integer**|  | |
| **employee** | [**List&lt;Employee&gt;**](Employee.md)|  | [optional] |

### Return type

[**List&lt;Employee&gt;**](Employee.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/json, text/json
 - **Accept**: text/plain, application/json, text/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |

<a id="apiEmployeePost"></a>
# **apiEmployeePost**
> List&lt;Employee&gt; apiEmployeePost(employee)

Adds new employee

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.EmployeeApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:5283");

    EmployeeApi apiInstance = new EmployeeApi(defaultClient);
    Employee employee = new Employee(); // Employee | 
    try {
      List<Employee> result = apiInstance.apiEmployeePost(employee);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling EmployeeApi#apiEmployeePost");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **employee** | [**Employee**](Employee.md)|  | [optional] |

### Return type

[**List&lt;Employee&gt;**](Employee.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/*+json
 - **Accept**: text/plain, application/json, text/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |

