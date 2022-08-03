package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"
)

type handler struct{}

//각 http 요청 url 경로에 대응할 수 있는 핸들러를 등록
//URL 수신되었을 때 그것을 처리하는 함수 또는 객체
//Handle() 함수로는 http.Hanlder 인터페이스를 구현한 객체를 등록
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Wooribank")
}

func main() {

	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)

	subject := pkix.Name{
		Organization:       []string{"test Organization"},
		OrganizationalUnit: []string{"test"},
		CommonName:         "Go Web Programming",
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}

	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()

	//=======================서버 연결===============================

	handler := handler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}

	//웹 서버 시작
	//func ListenAndServer(addr string, handler handler) error
	//func ListenAndServeTLS("cert.pem", "key.pem")
	server.ListenAndServeTLS("cert.pem", "key.pem")

}
