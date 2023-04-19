package ca

import (
	"certificateManager/config"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"path/filepath"
	"strings"
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

// func CreateRootCA(caconfig *config.CertConfigStruct) error {
func CreateRootCA(privateKeySize int) error {
	// Generate a new private key for the CA
	privateKey, err := rsa.GenerateKey(rand.Reader, privateKeySize)
	if err != nil {
		return err
	}
	if !strings.HasSuffix(config.CertConfigFile, ".json") {
		config.CertConfigFile += ".json"
	}
	CertConfig, err := config.Json2Config()
	if err != nil {
		return err
	}
	// We cannot allow a certificate to last 0yr, 0mt, 0d, so we set a defa
	if CertConfig.Duration == 0 {
		CertConfig.Duration = 1
	}
	// Create a new self-signed root certificate template
	template := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: CertConfig.CommonName, Locality: []string{CertConfig.Locality}, Country: []string{CertConfig.Country}, Organization: []string{CertConfig.Organization}, OrganizationalUnit: []string{CertConfig.OrganizationalUnit}, Province: []string{CertConfig.Province}},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(CertConfig.Duration, 0, 0),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign | x509.KeyUsageDigitalSignature,
		IsCA:                  CertConfig.IsCA,
		BasicConstraintsValid: true,
		DNSNames:              CertConfig.DNSNames,
		IPAddresses:           CertConfig.IPAddresses,
		EmailAddresses:        CertConfig.EmailAddresses,
	}
	if CertConfig.IsCA {
		template.KeyUsage = x509.KeyUsageCertSign | x509.KeyUsageCRLSign | x509.KeyUsageDigitalSignature
	}
	// Create the root certificate using the template and the private key
	rootCertBytes, err := x509.CreateCertificate(rand.Reader, template, template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return err
	}

	// Write the private key and root certificate to files
	privateKeyFile, err := os.Create(filepath.Join(CertConfig.CertificateDirectory, CertConfig.CertificateName) + ".key")
	if err != nil {
		return err
	}
	defer privateKeyFile.Close()

	err = pem.Encode(privateKeyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
	if err != nil {
		return err
	}

	rootCertFile, err := os.Create(filepath.Join(CertConfig.CertificateDirectory, CertConfig.CertificateName) + ".crt")
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

//func CreateRootCA1(nBytes int) error {
//	// Example usage
//	config := &config.CertConfigStruct{
//		IsCA:                 true,
//		CommonName:           "Famille Gratton",
//		Country:              "CA",
//		Province:             "Quebec",
//		Locality:             "Blainville",
//		Organization:         "famillegratton.net",
//		OrganizationalUnit:   "famillegratton",
//		Duration:             10,
//		DNSNames:             []string{"famillegratton.net", "nas.famillegratton.net", "lan.famillegratton.net"},
//		IPAddresses:          []net.IP{net.ParseIP("10.1.1.11"), net.ParseIP("127.0.0.1")},
//		CertificateDirectory: "/tmp",
//		CertificateName:      "rootCA",
//	}
//
//	err := createCACert(config)
//	if err != nil {
//		return err
//	}
//
//	fmt.Printf("Certificate files %s/%s.crt and %s/%s.key were created.\n",
//		config.CertificateDirectory, config.CertificateName, config.CertificateDirectory, config.CertificateName)
//	return nil
//}
