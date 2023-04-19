// certificateManager : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/config/configHelpers.go
// 4/16/23 13:52:19

package config

import (
	"crypto/x509"
	"encoding/json"
	"os"
	"strings"
)

var CertConfigFile = "defaultConfig.json"

//var ServerCertEnvironment = "serverCert-default.json"

func Json2Config() (CertConfigStruct, error) {
	var payload CertConfigStruct
	var err error
	rcDir, _ := os.UserHomeDir()
	rcFile := rcDir + "/.config/certificateManager/" + CertConfigFile
	jFile, err := os.ReadFile(rcFile)
	if err != nil {
		return CertConfigStruct{}, err
	}
	err = json.Unmarshal(jFile, &payload)
	if err != nil {
		return CertConfigStruct{}, err
	} else {
		return payload, nil
	}
}

// GetKeyUsageFromStrings() : converts a slice of strings into
// A x509.KeyUsage value. We use slices of strings because x509.KeyUsage
// Can hold multiple operations at once
func GetKeyUsageFromStrings(usageStrings []string) x509.KeyUsage {
	keyUsage := x509.KeyUsage(0)
	for _, usage := range usageStrings {
		switch strings.ToLower(usage) {
		case "digital signature":
			keyUsage |= x509.KeyUsageDigitalSignature
		case "content commitment":
			keyUsage |= x509.KeyUsageContentCommitment
		case "key encipherment":
			keyUsage |= x509.KeyUsageKeyEncipherment
		case "data encipherment":
			keyUsage |= x509.KeyUsageDataEncipherment
		case "key agreement":
			keyUsage |= x509.KeyUsageKeyAgreement
		case "cert sign":
		case "certificate sign":
			keyUsage |= x509.KeyUsageCertSign
		case "crl sign":
		case "crl":
			keyUsage |= x509.KeyUsageCRLSign
		case "encipheronly":
		case "encipher":
			keyUsage |= x509.KeyUsageEncipherOnly
		case "decipheronly":
		case "decipher":
			keyUsage |= x509.KeyUsageDecipherOnly
		}
	}
	return keyUsage
}

// GetStringsFromKeyUsage(): takes the x509.KeyUsage numerical value
// And converts it in a slice of human-readable strings,
// As KeyUsage can hold multiple operations at once.
func GetStringsFromKeyUsage(keyUsage x509.KeyUsage) []string {
	var usages []string

	if keyUsage&x509.KeyUsageDigitalSignature != 0 {
		usages = append(usages, "digital signature")
	}
	if keyUsage&x509.KeyUsageContentCommitment != 0 {
		usages = append(usages, "content commitment")
	}
	if keyUsage&x509.KeyUsageKeyEncipherment != 0 {
		usages = append(usages, "key encipherment")
	}
	if keyUsage&x509.KeyUsageDataEncipherment != 0 {
		usages = append(usages, "data encipherment")
	}
	if keyUsage&x509.KeyUsageKeyAgreement != 0 {
		usages = append(usages, "key agreement")
	}
	if keyUsage&x509.KeyUsageCertSign != 0 {
		usages = append(usages, "cert sign")
	}
	if keyUsage&x509.KeyUsageCRLSign != 0 {
		usages = append(usages, "crl sign")
	}
	if keyUsage&x509.KeyUsageEncipherOnly != 0 {
		usages = append(usages, "encipher only")
	}
	if keyUsage&x509.KeyUsageDecipherOnly != 0 {
		usages = append(usages, "decipher only")
	}
	return usages
}
