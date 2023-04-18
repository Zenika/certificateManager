// certificateManager : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/servercerts/verifyServerCert.go
// 4/18/23 05:34:48

package servercerts

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func verifyCertificate(certFilePath string) error {
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

	// Verify the certificate
	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		return err
	}
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}
	if ok := rootCAs.AppendCertsFromPEM(certPEMBlock); !ok {
		return fmt.Errorf("failed to append certificate to root CAs")
	}
	_, err = parsedCert.Verify(x509.VerifyOptions{
		Roots:         rootCAs,
		Intermediates: x509.NewCertPool(),
	})
	if err != nil {
		return err
	}

	return nil
}
