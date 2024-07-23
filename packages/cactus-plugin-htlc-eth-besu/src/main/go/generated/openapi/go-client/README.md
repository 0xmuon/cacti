# Go API client for cactus-plugin-htlc-eth-besu

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import cactus-plugin-htlc-eth-besu "github.com/hyperledger/cactus-plugin-htlc-eth-besu/src/main/go/generated/openapi/go-client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), cactus-plugin-htlc-eth-besu.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), cactus-plugin-htlc-eth-besu.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), cactus-plugin-htlc-eth-besu.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), cactus-plugin-htlc-eth-besu.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**GetSingleStatusV1**](docs/DefaultApi.md#getsinglestatusv1) | **Post** /api/v1/plugins/@hyperledger/cactus-plugin-htlc-eth-besu/get-single-status | 
*DefaultApi* | [**GetStatusV1**](docs/DefaultApi.md#getstatusv1) | **Post** /api/v1/plugins/@hyperledger/cactus-plugin-htlc-eth-besu/get-status | 
*DefaultApi* | [**InitializeV1**](docs/DefaultApi.md#initializev1) | **Post** /api/v1/plugins/@hyperledger/cactus-plugin-htlc-eth-besu/initialize | 
*DefaultApi* | [**NewContractV1**](docs/DefaultApi.md#newcontractv1) | **Post** /api/v1/plugins/@hyperledger/cactus-plugin-htlc-eth-besu/new-contract | 
*DefaultApi* | [**RefundV1**](docs/DefaultApi.md#refundv1) | **Post** /api/v1/plugins/@hyperledger/cactus-plugin-htlc-eth-besu/refund | 
*DefaultApi* | [**WithdrawV1**](docs/DefaultApi.md#withdrawv1) | **Post** /api/v1/plugins/@hyperledger/cactus-plugin-htlc-eth-besu/withdraw | 


## Documentation For Models

 - [GetSingleStatusRequest](docs/GetSingleStatusRequest.md)
 - [GetStatusRequest](docs/GetStatusRequest.md)
 - [InitializeRequest](docs/InitializeRequest.md)
 - [InvokeContractV1Response](docs/InvokeContractV1Response.md)
 - [NewContractObj](docs/NewContractObj.md)
 - [NewContractObjGas](docs/NewContractObjGas.md)
 - [RefundReq](docs/RefundReq.md)
 - [RunTransactionResponse](docs/RunTransactionResponse.md)
 - [Web3SigningCredential](docs/Web3SigningCredential.md)
 - [Web3SigningCredentialCactusKeychainRef](docs/Web3SigningCredentialCactusKeychainRef.md)
 - [Web3SigningCredentialNone](docs/Web3SigningCredentialNone.md)
 - [Web3SigningCredentialPrivateKeyHex](docs/Web3SigningCredentialPrivateKeyHex.md)
 - [Web3SigningCredentialType](docs/Web3SigningCredentialType.md)
 - [Web3TransactionReceipt](docs/Web3TransactionReceipt.md)
 - [WithdrawReq](docs/WithdrawReq.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


