# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [v1/annotation.proto](#v1_annotation-proto)
    - [MethodExtend](#api-v1-MethodExtend)
  
    - [AuthMethod](#api-v1-AuthMethod)
  
    - [File-level Extensions](#v1_annotation-proto-extensions)
  
- [v1/hello_service.proto](#v1_hello_service-proto)
    - [Req](#api-v1-Req)
    - [User](#api-v1-User)
  
    - [HelloService](#api-v1-HelloService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="v1_annotation-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## v1/annotation.proto



<a name="api-v1-MethodExtend"></a>

### MethodExtend



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| method_signature | [string](#string) |  | Method signature 方法签名 |
| allow_without_credential | [bool](#bool) |  | Whether to allow access without credentials 是否允许无凭证访问 |
| permission | [string](#string) |  | Permission required to access the method 访问该方法所需的权限 |
| auth_method | [AuthMethod](#api-v1-AuthMethod) |  | Authorization method 授权方法 |
| audit | [bool](#bool) |  | Whether to audit the method 是否审计该方法 |
| rpm | [int32](#int32) |  | Rate limit per minute 每分钟的速率限制 |
| timeout | [int32](#int32) |  | Timeout in milliseconds 超时控制（毫秒） |





 


<a name="api-v1-AuthMethod"></a>

### AuthMethod


| Name | Number | Description |
| ---- | ------ | ----------- |
| AUTH_METHOD_UNSPECIFIED | 0 |  |
| IAM | 1 | IAM uses the standard IAM authorization check on the organizational resources. |
| CUSTOM | 2 | Custom authorization method. |


 


<a name="v1_annotation-proto-extensions"></a>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| method_extend | MethodExtend | .google.protobuf.MethodOptions | 63501 |  |

 

 



<a name="v1_hello_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## v1/hello_service.proto



<a name="api-v1-Req"></a>

### Req
请求参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| age | [int32](#int32) |  | 年龄 3. 年龄限制 |
| purchase_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 4. 日期限制 |
| delivery_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| name | [string](#string) |  | 姓名 The name of the branch. Format: projects/{project}/branches/{branch} {branch} should be the id of a sheet. |






<a name="api-v1-User"></a>

### User
返回参数。用户信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the user. Format: users/{user}. {user} is a system-generated unique ID. |





 

 

 


<a name="api-v1-HelloService"></a>

### HelloService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetUser | [Req](#api-v1-Req) | [User](#api-v1-User) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

