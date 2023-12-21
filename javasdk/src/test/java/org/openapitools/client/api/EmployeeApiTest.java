/*
 * Employee API
 * Endpoints of CRUD of employee.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import org.openapitools.client.model.Employee;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for EmployeeApi
 */
@Disabled
public class EmployeeApiTest {

    private final EmployeeApi api = new EmployeeApi();

    /**
     * Returns all employee.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void apiEmployeeGetTest() throws ApiException {
        Employee response = api.apiEmployeeGet();
        // TODO: test validations
    }

    /**
     * @throws ApiException if the Api call fails
     */
    @Test
    public void apiEmployeeIdDeleteTest() throws ApiException {
        Integer id = null;
        List<Employee> response = api.apiEmployeeIdDelete(id);
        // TODO: test validations
    }

    /**
     * Returns employee of given id
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void apiEmployeeIdGetTest() throws ApiException {
        Integer id = null;
        Employee response = api.apiEmployeeIdGet(id);
        // TODO: test validations
    }

    /**
     * Updates employee details
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void apiEmployeeIdPutTest() throws ApiException {
        Integer id = null;
        Employee response = api.apiEmployeeIdPut(id);
        // TODO: test validations
    }

    /**
     * Adds new employee
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void apiEmployeePostTest() throws ApiException {
        Employee employee = null;
        List<Employee> response = api.apiEmployeePost(employee);
        // TODO: test validations
    }

}