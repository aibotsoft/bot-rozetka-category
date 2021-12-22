# \CategoriesApi

All URIs are relative to *https://xl-catalog-api.rozetka.com.ua*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChildren**](CategoriesApi.md#GetChildren) | **Get** /v4/categories/getChildren | getChildren



## GetChildren

> InlineResponse200 GetChildren(ctx).CategoryId(categoryId).Lang(lang).Execute()

getChildren



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    categoryId := "83850" // string | 
    lang := "ru" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.GetChildren(context.Background()).CategoryId(categoryId).Lang(lang).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.GetChildren``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChildren`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.GetChildren`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryId** | **string** |  | 
 **lang** | **string** |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

