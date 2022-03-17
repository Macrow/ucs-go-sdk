package ucs

import (
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewValidator(publicKey []byte) *Validator {
	if len(publicKey) == 0 {
		log.Fatalf("please provide rsa public key")
	}
	validator := &Validator{publicKey: publicKey}
	return validator
}

func NewRpcClient(addr string, port int) Client {
	client := &RpcClient{
		addr:    addr,
		port:    port,
		options: make([]grpc.DialOption, 0),
		timeout: DefaultTimeout,
	}
	client.options = append(client.options, grpc.WithTransportCredentials(insecure.NewCredentials()))

	return client
}

func NewTLSRpcClient(cert []byte, addr string, port int) Client {
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(cert) {
		log.Fatalf("credentials: failed to append certificates")
	}
	client := &RpcClient{
		addr:    addr,
		port:    port,
		options: make([]grpc.DialOption, 0),
		timeout: DefaultTimeout,
	}
	client.options = append(client.options, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(cp, "")))

	return client
}

func NewHttpClient(addr string, port int, ssl bool, accessCode string) Client {
	client := &HttpClient{
		addr:             addr,
		port:             port,
		ssl:              ssl,
		accessCode:       accessCode,
		accessCodeHeader: DefaultHeaderAccessCode,
		randomKeyHeader:  DefaultHeaderRandomKey,
		timeout:          DefaultTimeout,
	}
	return client
}
