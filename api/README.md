# Proto

Using [prototool](https://github.com/uber/prototool) for actual implementation, all protos will be in single repo to make it easier to manage.

## Repo layout

```sh
➜  protos git:(master) tree .           
.
├── README.md
├── audit
│   └── v1
│       ├── audit_api.pb.go
│       └── audit_api.proto
└── prototool.yaml

2 directories, 4 files

```

## prototool.yaml

```yaml
protoc:
  version: 3.8.0
lint:
  group: uber2
generate:
  go_options:
    import_path: github.com/<ORG>/protos
  plugins:
    - name: go
      type: go
      flags: plugins=grpc
      output: .
create:
  packages:
    - directory: audit
      name: audit
```

## audit_api.proto

```go
syntax = "proto3";

package audit.v1;

option csharp_namespace = "Audit.V1";
option go_package = "auditv1";
option java_multiple_files = true;
option java_outer_classname = "AuditApiProto";
option java_package = "com.audit.v1";
option objc_class_prefix = "AXX";
option php_namespace = "Audit\\V1";

// AuditAPI should only create and read user and org audits.
service AuditAPI {
  // Create a new audit.
  rpc Create(CreateRequest) returns (CreateResponse);
  // ReadUser audit log.
  rpc ReadUser(ReadUserRequest) returns (ReadUserResponse);
  // ReadOrg audit log.
  rpc ReadOrg(ReadOrgRequest) returns (ReadOrgResponse);
}

message CreateRequest {
  string user_id = 1;
  string org_id = 2;
  string ip_address = 3;
  string target = 4;
  string action = 5;
  string action_type = 6;
  string event_name = 7;
}

message CreateResponse {
  string response = 1;
}

message ReadUserRequest {
  string user_id = 1;
}

message ReadUserResponse {
  string response = 1;
}

message ReadOrgRequest {
  string org_id = 1;
}

message ReadOrgResponse {
  string response = 1;
}
```