# EmployeeApi

All URIs are relative to *https://virtserver.swaggerhub.com/ISHAAABDUL23/EmployeeAPI/1.0.0*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiEmployeeGet**](EmployeeApi.md#apiEmployeeGet) | **GET** /api/Employee | Returns all employee. |
| [**apiEmployeeIdDelete**](EmployeeApi.md#apiEmployeeIdDelete) | **DELETE** /api/Employee/{id} |  |
| [**apiEmployeeIdGet**](EmployeeApi.md#apiEmployeeIdGet) | **GET** /api/Employee/{id} | Returns employee of given id |
| [**apiEmployeeIdPut**](EmployeeApi.md#apiEmployeeIdPut) | **PUT** /api/Employee/{id} | Updates employee details |
| [**apiEmployeePost**](EmployeeApi.md#apiEmployeePost) | **POST** /api/Employee | Adds new employee |


<a id="apiEmployeeGet"></a>
# **apiEmployeeGet**
> Employee apiEmployeeGet()

Returns all employee.

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
    defaultClient.setBasePath("https://virtserver.swaggerhub.com/ISHAAABDUL23/EmployeeAPI/1.0.0");

    EmployeeApi apiInstance = new EmployeeApi(defaultClient);
    try {
      Employee result = apiInstance.apiEmployeeGet();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling EmployeeApi#apiEmployeeGet");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**Employee**](Employee.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |

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
    defaultClient.setBasePath("https://virtserver.swaggerhub.com/ISHAAABDUL23/EmployeeAPI/1.0.0");

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
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |

<a id="apiEmployeeIdGet"></a>
# **apiEmployeeIdGet**
> Employee apiEmployeeIdGet(id)

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
    defaultClient.setBasePath("https://virtserver.swaggerhub.com/ISHAAABDUL23/EmployeeAPI/1.0.0");

    EmployeeApi apiInstance = new EmployeeApi(defaultClient);
    Integer id = 56; // Integer | 
    try {
      Employee result = apiInstance.apiEmployeeIdGet(id);
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

[**Employee**](Employee.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |

<a id="apiEmployeeIdPut"></a>
# **apiEmployeeIdPut**
> Employee apiEmployeeIdPut(id)

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
    defaultClient.setBasePath("https://virtserver.swaggerhub.com/ISHAAABDUL23/EmployeeAPI/1.0.0");

    EmployeeApi apiInstance = new EmployeeApi(defaultClient);
    Integer id = 56; // Integer | 
    try {
      Employee result = apiInstance.apiEmployeeIdPut(id);
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

### Return type

[**Employee**](Employee.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

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
    defaultClient.setBasePath("https://virtserver.swaggerhub.com/ISHAAABDUL23/EmployeeAPI/1.0.0");

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

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |

