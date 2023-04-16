package ca

import (
	"certificateManager/configs"
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

func createCACert(caconfig *configs.RootCAconfig) error {
	// Generate a new private key for the CA
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}

	// Create a new self-signed root certificate template
	template := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: caconfig.CommonName, Locality: []string{caconfig.Locality}, Country: []string{caconfig.Country}, Organization: []string{caconfig.Organization}, OrganizationalUnit: []string{caconfig.OrganizationalUnit}, Province: []string{caconfig.Province}},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(caconfig.Duration, 0, 0),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign | x509.KeyUsageDigitalSignature,
		IsCA:                  true,
		BasicConstraintsValid: true,
		DNSNames:              caconfig.DNS,
		IPAddresses:           caconfig.IPS,
		EmailAddresses:        []string{"certs@famillegratton.net", "jfgratton@famillegratton.net"},
	}

	// Create the root certificate using the template and the private key
	rootCertBytes, err := x509.CreateCertificate(rand.Reader, template, template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return err
	}

	// Write the private key and root certificate to files
	privateKeyFile, err := os.Create(filepath.Join(caconfig.CertificateDirectory, caconfig.CertificateName+".key"))
	if err != nil {
		return err
	}
	defer privateKeyFile.Close()

	err = pem.Encode(privateKeyFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
	if err != nil {
		return err
	}

	rootCertFile, err := os.Create(filepath.Join(caconfig.CertificateDirectory, caconfig.CertificateName+".crt"))
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

func CreateRootCA() {
	// Example usage
	config := &configs.RootCAconfig{
		CommonName:         "Famille Gratton",
		Country:            "CA",
		Province:           "Quebec",
		Locality:           "Blainville",
		Organization:       "famillegratton.net",
		OrganizationalUnit: "famillegratton",
		Duration:           10,
		//DNSNames:      []string{"famillegratton.net", "nas.famillegratton.net", "lan.famillegratton.net"},
		//IPAddresses:  []net.IP{},
		CertificateDirectory: "/tmp/",
		CertificateName:      "rootCA",
	}
	err := createCACert(config)
	if err != nil {
		panic(err)
	}
}
