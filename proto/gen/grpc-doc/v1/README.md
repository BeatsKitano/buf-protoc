# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [v1/annotation.proto](#v1_annotation-proto)
    - [MethodExtend](#api-v1-MethodExtend)
  
    - [AuthMethod](#api-v1-AuthMethod)
  
    - [File-level Extensions](#v1_annotation-proto-extensions)
  
- [v1/common.proto](#v1_common-proto)
    - [State](#api-v1-State)
  
- [v1/hello_service.proto](#v1_hello_service-proto)
    - [Database](#api-v1-Database)
    - [Database.LabelsEntry](#api-v1-Database-LabelsEntry)
    - [GetDatabaseRequest](#api-v1-GetDatabaseRequest)
    - [GetUserRequest](#api-v1-GetUserRequest)
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
| concurrent | [int32](#int32) |  | Concurrent 并发控制 |





 


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

 

 



<a name="v1_common-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## v1/common.proto


 


<a name="api-v1-State"></a>

### State


| Name | Number | Description |
| ---- | ------ | ----------- |
| STATE_UNSPECIFIED | 0 |  |
| ACTIVE | 1 |  |
| DELETED | 2 |  |


 

 

 



<a name="v1_hello_service-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## v1/hello_service.proto



<a name="api-v1-Database"></a>

### Database



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the database. Format: instances/{instance}/databases/{database} {database} is the database name in the instance. |
| sync_state | [State](#api-v1-State) |  | The existence of a database on latest sync. |
| successful_sync_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | The latest synchronization time. |
| project | [string](#string) |  | The project for a database. Format: projects/{project} |
| schema_version | [string](#string) |  | The version of database schema. |
| environment | [string](#string) |  | The environment resource. Format: environments/prod where prod is the environment resource ID. |
| effective_environment | [string](#string) |  | The effective environment based on environment tag above and environment tag on the instance. Inheritance follows https://cloud.google.com/resource-manager/docs/tags/tags-overview. |
| labels | [Database.LabelsEntry](#api-v1-Database-LabelsEntry) | repeated | Labels will be used for deployment and policy control. |
| backup_available | [bool](#bool) |  | The database is available for DML prior backup. |






<a name="api-v1-Database-LabelsEntry"></a>

### Database.LabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="api-v1-GetDatabaseRequest"></a>

### GetDatabaseRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name of the database to retrieve. Format: instances/{instance}/databases/{database} |






<a name="api-v1-GetUserRequest"></a>

### GetUserRequest
请求参数


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| age | [int32](#int32) |  | 字段解释描述文本

syntax. CEL is a C-like expression language. The syntax and semantics of CEL

are documented at https://github.com/google/cel-spec. { &#34;age&#34;: 25, &#34;purchase_date&#34;: { &#34;seconds&#34;: 1680054400, // 2023-03-27T00:00:00Z &#34;nanos&#34;: 0 }, &#34;delivery_date&#34;: { &#34;seconds&#34;: 1680140800, // 2023-03-28T00:00:00Z &#34;nanos&#34;: 0 }, &#34;name&#34;: &#34;productA&#34;, &#34;custom_expr&#34;: { // 自定义表达式的JSON表示，根据具体需求填充 &#34;expr&#34;: &#34;value &gt; 10&#34;, &#34;type&#34;: &#34;BOOL&#34; } }

段落1:

 标题1 这就是描述1 document.summary.size() &lt; 100

段落2: 

 标题2 这就是描述2 document.owner == request.auth.claims.email |
| purchase_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 购买日期 |
| delivery_date | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 交付日期 |
| name | [string](#string) |  | 物品名 |
| custom_expr | [google.type.Expr](#google-type-Expr) |  | 自定义表达式 |






<a name="api-v1-User"></a>

### User
用户信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | 用户id |
| name | [string](#string) |  | 用户名 |
| son | [string](#string) |  |  |





 

 

 


<a name="api-v1-HelloService"></a>

### HelloService
提供问候语服务。
1. 服务级别的注解 option (google.api.default_host) = &#34;https://huige.api.com&#34;;
2. 服务级别的注解 option (google.api.oauth_scopes) = &#34;https://www.huige.com/auth/user&#34;;
3. 服务级别的注解
4. 服务级别的注解
5. 服务级别的注解

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetUser | [GetUserRequest](#api-v1-GetUserRequest) | [User](#api-v1-User) | 获取用户信息 |
| GetDatabase | [GetDatabaseRequest](#api-v1-GetDatabaseRequest) | [Database](#api-v1-Database) | 使用&lt;br&gt;换行 instances/*/databases/*的注解&lt;br&gt; Gets the database with the given name.&lt;br&gt; The name must be in the format: instances/{instance}/databases/{database}. |

 



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

