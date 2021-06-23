不确定  
```protobuf
message SearchRequest {
string query = 1;
int32 page_number = 2;
int32 result_per_page = 3;
enum Corpus {
UNIVERSAL = 0;
WEB = 1;
IMAGES = 2;
LOCAL = 3;
NEWS = 4;
PRODUCTS = 5;
VIDEO = 6;
}
Corpus corpus = 4;
}
```
引用  
```protobuf
message SearchResponse {
repeated Result results = 1;
}

message Result {
string url = 1;
string title = 2;
repeated string snippets = 3;
}
```
嵌套  
```protobuf
message SearchResponse {
message Result {
string url = 1;
string title = 2;
repeated string snippets = 3;
}
repeated Result results = 1;
}
message SomeOtherMessage {
SearchResponse.Result result = 1;
}
```
任意类型  
```protobuf
import "google/protobuf/any.proto";

message ErrorStatus {
string message = 1;
repeated google.protobuf.Any details = 2;
}
```
包  
```protobuf
package foo.bar;
message Open { ... }
message Foo {
...
required foo.bar.Open open = 1;
...
}
```