# Knative Functions

## Create a function

```bash
rm -rf siemens-function
func create -l go siemens-function
export FUNC_PATH=/Users/rlehmann/code/retocode/demos/siemens-oss-days/siemens-function
```

```text
Created go function in ./siemens-function
```

```bash
tree siemens-function
```

```text
siemens-function
├── README.md
├── func.yaml
├── go.mod
├── handle.go
└── handle_test.go

1 directory, 5 files
```

## Building the function
```bash
func build
```
```text
🙌 Function image built: quay.io/rlehmann/siemens-function:latest
```

## Running and invoking the function
```bash
func run --build=false
```

```bash
cd siemens-function
func invoke
```

```text
Received response
POST / HTTP/1.1 localhost:8080
  User-Agent: Go-http-client/1.1
  Content-Length: 25
  Content-Type: application/json
  Accept-Encoding: gzip
Body:
```

## Deploying and invoking the function

```bash
watch -n 1 kubectl get pod -n default
```

```bash
func deploy
```

```text
⬆️  Deploying function to the cluster
✅ Function updated in namespace "default" and exposed at URL: 
http://siemens-function.default.10.89.0.200.sslip.io
```

```bash
func invoke --target=remote
```

```text
Received response
POST / HTTP/1.1 siemens-function.default.10.89.0.200.sslip.io
  X-Request-Id: 03f94c2c-eb05-4775-b190-e9aa8e7f3294
  Forwarded: for=10.244.3.4;proto=http
  K-Proxy-Request: activator
  X-Forwarded-For: 10.244.3.4, 10.244.3.3
  X-Forwarded-Proto: http
  User-Agent: Go-http-client/1.1
  Content-Length: 25
  Accept-Encoding: gzip
  Content-Type: application/json
Body:
```

## Update and redeploy

> 📝 Change something in ./siemens-function/handle.go

```bash
func deploy
```

We can also use `curl`
```bash
curl -X POST http://siemens-function.default.10.89.0.200.sslip.io
```

```text
POST / HTTP/1.1 siemens-function.default.10.89.0.200.sslip.io
  User-Agent: curl/7.87.0
  Accept: */*
  K-Proxy-Request: activator
  X-Forwarded-Proto: http
  X-Request-Id: 044636c7-9b44-411f-8e84-a4d42cfc8bd9
  Content-Length: 0
  Forwarded: for=10.244.3.4;proto=http
  X-Forwarded-For: 10.244.3.4, 10.244.3.3
Hello Siemens Folks 🙋‍♂
```
