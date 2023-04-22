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

	config.CertConfig, err = config.Json2Config()
	if err != nil {
		return err
	}
	// We cannot allow a certificate to last 0yr, 0mt, 0d, so we set a default value of 1 year
	if config.CertConfig.Duration == 0 {
		config.CertConfig.Duration = 1
	}
	// Create a new self-signed root certificate template
	template := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: config.CertConfig.CommonName, Locality: []string{config.CertConfig.Locality}, Country: []string{config.CertConfig.Country}, Organization: []string{config.CertConfig.Organization}, OrganizationalUnit: []string{config.CertConfig.OrganizationalUnit}, Province: []string{config.CertConfig.Province}},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(config.CertConfig.Duration, 0, 0),
		KeyUsage:              config.GetKeyUsageFromStrings(config.CertConfig.KeyUsage),
		IsCA:                  config.CertConfig.IsCA,
		BasicConstraintsValid: true,
		DNSNames:              config.CertConfig.DNSNames,
		IPAddresses:           config.CertConfig.IPAddresses,
		EmailAddresses:        config.CertConfig.EmailAddresses,
	}
	if config.CertConfig.IsCA {
		template.KeyUsage = config.ReindexKeyUsage(config.CertConfig)
	}
	// Create the root certificate using the template and the private key
	rootCertBytes, err := x509.CreateCertificate(rand.Reader, template, template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return err
	}

	// Write the private key and root certificate to files
	privateKeyFile, err := os.Create(filepath.Join(config.CertConfig.CertificateDirectory, config.CertConfig.CertificateName) + ".key")
	if err != nil {
		return err
	}
	defer privateKeyFile.Close()

	err = pem.Encode(privateKeyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
	if err != nil {
		return err
	}

	rootCertFile, err := os.Create(filepath.Join(config.CertConfig.CertificateDirectory, config.CertConfig.CertificateName) + ".crt")
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
