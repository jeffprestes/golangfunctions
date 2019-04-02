# Google Cloud Functions examples in Golang

## How to setup an account at Google Cloud and setup gcloud tools for Golang

Please visit JBD https://medium.com/google-cloud/google-cloud-functions-for-go-57e4af9b10da and Benjamin's tutorials https://benjamincongdon.me/blog/2019/01/21/Getting-Started-with-Golang-Google-Cloud-Functions/


## Comand to deploy functions

```bash
gcloud functions deploy golangfunctions --entry-point OlaMundo --runtime go111 --trigger-http
```

```bash
gcloud functions deploy testfunctions --entry-point SouUmTeste --runtime go111 --trigger-http
```

## URLs generated for these examples

https://us-central1-golang-functions.cloudfunctions.net/golangfunctions

https://us-central1-golang-functions.cloudfunctions.net/testfunctions
