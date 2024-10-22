# Secure Three Party Computation via Replicated Secret Sharing

This project demonstrates Secure Three Party Computation via Replicated Secret Sharing using Go with TLS communication between patients (Alice, Bob, Charlie) and the Hospital (server).

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.19 or higher)
- OpenSSL (for generating certificates)

## Generate TLS Certificates
**Make key and self-signed certification**:
    
Run these three commands to create a key and a self-signed certificate
```bash
openssl genrsa -out key.pem 2048
openssl req -new -key key.pem -out cert.csr
openssl x509 -req -days 365 -in cert.csr -signkey key.pem -out cert.pem
```
Make sure that the files **cert.pem** and **key.pem** now is in the directory of the project.
## Build the Project
To be able to compile the go project, initialize the go.mod:
```bash
go mod init handin2
```
Now build the project:
```bash
go build -o mpc
```
Now you can run the code:
```bash
./mpc
```