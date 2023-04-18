// certificateManager : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/ca/verifyRootCAcert.go
// 4/18/23 05:37:11

package ca

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

var CaVerifyVerbose = false

func VerifyCACertificate(certFilePath string) error {
	// Read the certificate file
	certPEMBlock, err := os.ReadFile(certFilePath)
	if err != nil {
		return err
	}

	// Decode the PEM block into a certificate
	cert, _ := pem.Decode(certPEMBlock)
	if cert == nil {
		return fmt.Errorf("failed to decode certificate PEM block")
	}

	// Parse the certificate
	parsedCert, err := x509.ParseCertificate(cert.Bytes)
	if err != nil {
		return err
	}

	// Print certificate information
	fmt.Printf("Certificate:\n")
	if CaVerifyVerbose {
		fmt.Printf("    Data:\n%s\n", string(certPEMBlock))
	}
	fmt.Printf("    Subject: %v\n", parsedCert.Subject)
	fmt.Printf("    Issuer: %v\n", parsedCert.Issuer)
	fmt.Printf("    Serial Number: %v\n", parsedCert.SerialNumber)
	fmt.Printf("    Not Before: %v\n", parsedCert.NotBefore)
	fmt.Printf("    Not After : %v\n", parsedCert.NotAfter)
	//fmt.Printf("    DNS Names : %v\n", parsedCert.DNSNames)
	//fmt.Printf("    Email Addresses : %v\n", parsedCert.EmailAddresses)
	//fmt.Printf("    IP Addresses : %v\n", parsedCert.IPAddresses)
	if len(parsedCert.URIs) > 0 {
		fmt.Printf("    URIs : %v\n", parsedCert.URIs)
	}
	fmt.Printf("    Signature Algorithm: %v\n", parsedCert.SignatureAlgorithm)
	if CaVerifyVerbose {
		fmt.Printf("    Signature: %v\n", parsedCert.Signature)
	}

	// Print X509v3 Key Usage information
	if parsedCert.KeyUsage != 0 {
		fmt.Printf("    X509v3 Key Usage: %s\n", parsedCert.KeyUsage)
	}

	// Print X509v3 Basic Constraints information
	if parsedCert.BasicConstraintsValid {
		if parsedCert.IsCA {
			fmt.Printf("    X509v3 Basic Constraints:\n\tIs CA: true\n")
		} else {
			fmt.Printf("    X509v3 Basic Constraints:\n\tIs CA: false\n")
		}
	}

	// Print X509v3 Subject Alternative Name information
	if len(parsedCert.DNSNames) > 0 {
		fmt.Printf("    X509v3 Subject Alternative Name(s):\n")
		fmt.Println("        DNS:")
		for _, dns := range parsedCert.DNSNames {
			fmt.Printf("        \t- %s\n", dns)
		}
	}
	if len(parsedCert.IPAddresses) > 0 {
		fmt.Println("        IP Address(es):")
		for _, ipa := range parsedCert.IPAddresses {
			fmt.Printf("        \t- %s\n", ipa)
		}
	}
	if len(parsedCert.EmailAddresses) > 0 {
		fmt.Println("        Email Address(es):")
		for _, email := range parsedCert.EmailAddresses {
			fmt.Printf("        \t- %s\n", email)
		}
	}

	return nil
}
