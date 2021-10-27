# apigw-cognito-example

An example using the client credential workflow with Amazon Cognito and API
Gateway.

## Usage

```text
Usage:
  apigw-cognito-example [flags]

Flags:
      --appURL string          Application URL
      --clientID string        Client ID
      --clientSecret string    Client Secret
  -h, --help                   help for apigw-cognito-example
      --scope string           OAuth2 scope
      --tokenEndpoint string   OAuth2 token endpoint
```

## Setup

Run the following command to create an empty `.envrc.local` file:

```shell
% make envrc_local
```

Update `.envrc.local` with actual values.

Run the following to build and run the client:

```shell
% make build
% make run
```
