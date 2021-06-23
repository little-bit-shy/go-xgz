# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api.proto](#api.proto)
    - [HelloReq](#app.framework.v1.HelloReq)
    - [HelloResp](#app.framework.v1.HelloResp)
    - [SayResp](#app.framework.v1.SayResp)
    - [SayResp.CallSelfMsg](#app.framework.v1.SayResp.CallSelfMsg)
    - [SayResp.DbMsg](#app.framework.v1.SayResp.DbMsg)
    - [SayResp.EsMsg](#app.framework.v1.SayResp.EsMsg)
    - [SayResp.JrpcMsg](#app.framework.v1.SayResp.JrpcMsg)
    - [SayResp.JrpcMsg.ResultEntry](#app.framework.v1.SayResp.JrpcMsg.ResultEntry)
    - [SayResp.JrpcMsg.language](#app.framework.v1.SayResp.JrpcMsg.language)
  
    - [App](#app.framework.v1.App)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api.proto
注意：参数统一采用&#34;_&#34;拼接，比如ClusterName=cluster_name，且如果RPC服务开放了网关，方法为读是GET为写则为POST


<a name="app.framework.v1.HelloReq"></a>

### HelloReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Name | [string](#string) |  |  |






<a name="app.framework.v1.HelloResp"></a>

### HelloResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Content | [string](#string) |  |  |






<a name="app.framework.v1.SayResp"></a>

### SayResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Db | [SayResp.DbMsg](#app.framework.v1.SayResp.DbMsg) |  |  |
| Es | [SayResp.EsMsg](#app.framework.v1.SayResp.EsMsg) |  |  |
| Redis | [string](#string) |  |  |
| Jrpc | [SayResp.JrpcMsg](#app.framework.v1.SayResp.JrpcMsg) |  |  |
| Client | [SayResp.CallSelfMsg](#app.framework.v1.SayResp.CallSelfMsg) |  |  |
| Hbase | [bytes](#bytes) | repeated |  |






<a name="app.framework.v1.SayResp.CallSelfMsg"></a>

### SayResp.CallSelfMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Content | [string](#string) |  |  |






<a name="app.framework.v1.SayResp.DbMsg"></a>

### SayResp.DbMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Id | [int64](#int64) |  |  |
| Name | [string](#string) |  |  |






<a name="app.framework.v1.SayResp.EsMsg"></a>

### SayResp.EsMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ClusterName | [string](#string) |  |  |
| Status | [string](#string) |  |  |
| TimedOut | [bool](#bool) |  |  |
| NumberOfNodes | [double](#double) |  |  |
| NumberOfDataNodes | [double](#double) |  |  |
| ActivePrimaryShards | [double](#double) |  |  |
| ActiveShards | [double](#double) |  |  |
| RelocatingShards | [double](#double) |  |  |
| InitializingShards | [double](#double) |  |  |
| UnassignedShards | [double](#double) |  |  |
| DelayedUnassignedShards | [double](#double) |  |  |
| NumberOfPendingTasks | [double](#double) |  |  |
| NumberOfInFlightFetch | [double](#double) |  |  |
| TaskMaxWaitingInQueue | [string](#string) |  |  |
| TaskMaxWaitingInQueueMillis | [double](#double) |  |  |
| ActiveShardsPercent | [string](#string) |  |  |
| ActiveShardsPercentAsNumber | [double](#double) |  |  |






<a name="app.framework.v1.SayResp.JrpcMsg"></a>

### SayResp.JrpcMsg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Jsonrpc | [string](#string) |  |  |
| Result | [SayResp.JrpcMsg.ResultEntry](#app.framework.v1.SayResp.JrpcMsg.ResultEntry) | repeated |  |






<a name="app.framework.v1.SayResp.JrpcMsg.ResultEntry"></a>

### SayResp.JrpcMsg.ResultEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [SayResp.JrpcMsg.language](#app.framework.v1.SayResp.JrpcMsg.language) |  |  |






<a name="app.framework.v1.SayResp.JrpcMsg.language"></a>

### SayResp.JrpcMsg.language



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| En | [string](#string) |  |  |





 

 

 


<a name="app.framework.v1.App"></a>

### App
服务方法

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Ping | [.google.protobuf.Empty](#google.protobuf.Empty) | [.google.protobuf.Empty](#google.protobuf.Empty) | ping |
| Say | [HelloReq](#app.framework.v1.HelloReq) | [SayResp](#app.framework.v1.SayResp) | say |
| CallSelf | [HelloReq](#app.framework.v1.HelloReq) | [HelloResp](#app.framework.v1.HelloResp) | call_self |

 



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

