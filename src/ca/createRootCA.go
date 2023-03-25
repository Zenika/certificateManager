package ca

import (
	"certificateManager/configs"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

/*
The createCACert function takes a RootCAconfig struct as input, which contains the following fields:

CommonName: The common name to use for the root certificate.
ValidForYears: The number of years that the root certificate should be valid for.
DNSNames: A list of DNS names to include in the root certificate.
IPAddresses: A list of IP addresses to include in the root certificate.
KeyFilePath: The path to write the private key to.
CertFilePath: The path to write the root certificate to.
The function generates a new RSA private key
*/

func createCACert(caconfig *configs.RootCAconfig) error {
	// Generate a new private key for the CA
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	// Create a new self-signed root certificate template
	template := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: caconfig.CommonName},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(caconfig.ValidForYears, 0, 0),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		IsCA:                  true,
		BasicConstraintsValid: true,
		DNSNames:              caconfig.DNSNames,
		IPAddresses:           caconfig.IPAddresses,
	}

	// Create the root certificate using the template and the private key
	rootCertBytes, err := x509.CreateCertificate(rand.Reader, template, template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return err
	}

	// Write the private key and root certificate to files
	privateKeyFile, err := os.Create(caconfig.KeyFilePath)
	if err != nil {
		return err
	}
	defer privateKeyFile.Close()

	err = pem.Encode(privateKeyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
	if err != nil {
		return err
	}

	rootCertFile, err := os.Create(caconfig.CertFilePath)
	if err != nil {
		return err
	}
	defer rootCertFile.Close()

	err = pem.Encode(rootCertFile, &pem.Block{Type: "CERTIFICATE", Bytes: rootCertBytes})
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Example usage
	config := &configs.RootCAconfig{
		CommonName:    "My Root CA",
		ValidForYears: 10,
		DNSNames:      []string{"example.com", "www.example.com"},
		IPAddresses:   []net.IP{net.ParseIP("127.0.0.1")},
		KeyFilePath:   "/path/to/root.key",
		CertFilePath:  "/path/to/root.crt",
	}
	err := createCACert(config)
	if err != nil {
		panic(err)
	}
}
