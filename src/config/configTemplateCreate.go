// certificateManager
// Écrit par J.F. Gratton <jean-francois@famillegratton.net>
// Orininal name: src/config/configTemplateCreate.go
// Original time: 2023/06/03 07:01

package config

import (
	"cm/helpers"
	"net"
	"os"
	"path/filepath"
)

func TemplateConfigCreate() error {
	if err := createExplanationfile(); err != nil {
		return err
	}
	if err := createSampleTemplate(); err != nil {
		return err
	}
	return nil
}

func createExplanationfile() error {
	expText := `{
	"Country" : "CA", -> This is the country of origin for the certificate
	"Provice" : "Quebec", -> State of province of origin
	"Locality" : "Blainville", -> City of origin
	"Organization" : "myorg.net", -> Organization of origin
	"OrganizationalUnit" : "myorg", -> Sub-organization of origin
	"CommonName" : "myorg.net root CA", -> The name of the certificate
	"EmailAddresses" : ["certs@myorg.net", "certificates@myorg.net"], -> Array of email addresses responsible for this cert
	"Duration" : 10, -> CA duration, in years
	"KeyUsage" : ["Digital Signature", "Certificate Sign", "CRL Sign"], -> Certificate usage. This here are common values for CAs
	"DNSNames" : ["myorg.net","myorg.com","lan.myorg.net"], -> DNS names assigned to this cert
	"IPAddresses" : ["10.1.1.11", "127.0.0.1"], -> IP addresses assigned to this cert (never a good idea to assign IPs to a CA)
	"CertificateDirectory" : "/tmp/", -> directory where to write the cert
	"CertificateName" : "sample_cert", -> cert filename, no extension to the filename
	"IsCA": true, -> Are we creating a CA or a "normal" server cert ?
	"Comments": ["To see which values to put in the Usage field, see https://pkg.go.dev/crypto/x509#KeyUsage", "Strip off 'KeyUsage' from the const name and there you go.", "", "Please note that this field offers no functionality and is strictly here for documentation purposes"] -> Those won't appear in the certificate file
}`

	expFile, err := os.Create(filepath.Join(os.Getenv("HOME"), ".config", "certificatemanager", "template-README.txt"))
	if err != nil {
		return err
	}
	defer expFile.Close()

	_, err = expFile.WriteString(expText)
	if err != nil {
		return err
	}

	return nil
}

// Note : it is a bit ugly: this function is the same as the one above, without the explanations, and with a different filename
func createSampleTemplate2() error {
	expText := `{
	"Country" : "CA",
	"Provice" : "Quebec",
	"Locality" : "Blainville",
	"Organization" : "myorg.net",
	"OrganizationalUnit" : "myorg",
	"CommonName" : "myorg.net root CA",
	"EmailAddresses" : ["certs@myorg.net"],
	"Duration" : 10,
	"KeyUsage" : ["Digital Signature", "Certificate Sign", "CRL Sign"],
	"DNSNames" : ["myorg.net","myorg.com","lan.myorg.net"],
	"IPAddresses" : ["10.1.1.11"],
	"CertificateDirectory" : "/tmp/",
	"CertificateName" : "dev_cert",
	"IsCA": true,
	"Comments": ["To see which values to put in the Usage field, see https://pkg.go.dev/crypto/x509#KeyUsage", "Strip off 'KeyUsage' from the const name and there you go.", "", "Please note that this field offers no functionality and is strictly here for documentation purposes"]
}`

	expFile, err := os.Create(filepath.Join(os.Getenv("HOME"), ".config", "certificatemanager", "template-README.txt"))
	if err != nil {
		return err
	}
	defer expFile.Close()

	_, err = expFile.WriteString(expText)
	if err != nil {
		return err
	}

	return nil
}

func createSampleTemplate() error {
	var sampleCertConfig = helpers.CertConfigStruct{
		Country:              "CA",
		Province:             "Quebec",
		Locality:             "Blainville",
		Organization:         "myorg.net",
		OrganizationalUnit:   "myorg",
		CommonName:           "myorg.net root CA",
		EmailAddresses:       []string{"certs@myorg.net", "certificates@myorg.net"},
		Duration:             10,
		KeyUsage:             []string{"cert sign", "crl sign", "digital signature"},
		DNSNames:             []string{"myorg.net", "myorg.com", "lan.myorg.net"},
		IPAddresses:          []net.IP{net.ParseIP("10.1.1.11"), net.ParseIP("127.0.0.1")},
		CertificateDirectory: "/tmp",
		CertificateName:      "sample_cert",
		IsCA:                 true,
		Comments: []string{"To see which values to put in the Usage field, see https://pkg.go.dev/crypto/x509#KeyUsage",
			"Strip off 'KeyUsage' from the const name and there you go.",
			"",
			"Please note that this field offers no functionality and is strictly here for documentation purposes"},
	}
	if err := sampleCertConfig.Config2Json("template.json"); err != nil {
		return err
	}
	return nil
}
