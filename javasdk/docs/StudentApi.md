# StudentApi

All URIs are relative to *http://localhost:5283*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**apiStudentGet**](StudentApi.md#apiStudentGet) | **GET** /api/Student | Returns all student details |
| [**apiStudentIdGet**](StudentApi.md#apiStudentIdGet) | **GET** /api/Student/{id} | Returns student of given id |


<a id="apiStudentGet"></a>
# **apiStudentGet**
> List&lt;Student&gt; apiStudentGet()

Returns all student details

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.StudentApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:5283");

    StudentApi apiInstance = new StudentApi(defaultClient);
    try {
      List<Student> result = apiInstance.apiStudentGet();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling StudentApi#apiStudentGet");
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

[**List&lt;Student&gt;**](Student.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |

<a id="apiStudentIdGet"></a>
# **apiStudentIdGet**
> List&lt;Student&gt; apiStudentIdGet(id)

Returns student of given id

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.StudentApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:5283");

    StudentApi apiInstance = new StudentApi(defaultClient);
    Integer id = 56; // Integer | 
    try {
      List<Student> result = apiInstance.apiStudentIdGet(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling StudentApi#apiStudentIdGet");
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

[**List&lt;Student&gt;**](Student.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | Success |  -  |

