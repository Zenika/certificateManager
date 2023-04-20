// certificateManager
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// configStructs.go, jfgratton : 2023-03-20

package config

import (
	"net"
)

type CertConfigStruct struct {
	Country              string   `json:"Country"`
	Province             string   `json:"Province"`
	Locality             string   `json:"Locality"`
	Organization         string   `json:"Organization"`
	OrganizationalUnit   string   `json:"OrganizationalUnit,omitempty"`
	CommonName           string   `json:"CommonName"`
	IsCA                 bool     `json:"IsCA,omitempty"`
	EmailAddresses       []string `json:"EmailAddresses,omitempty"`
	Duration             int      `json:"Duration"`
	KeyUsage             []string `json:"KeyUsage"`
	DNSNames             []string `json:"DNSNames,omitempty"`
	IPAddresses          []net.IP `json:"IPAddresses,omitempty"`
	CertificateDirectory string   `json:"CertificateDirectory"`
	CertificateName      string   `json:"CertificateName"`
}

type EnvConfigStruct struct {
	CertificateRootDir string `json:"CertificateRootDir"`
	RootCAdir          string `json:"RootCAdir"`
	ServerCertsDir     string `json:"ServerCertsDir"`
	RemoveDuplicates   bool   `json:"RemoveDuplicates"`
}

var CertConfig = CertConfigStruct{Duration: 1, KeyUsage: []string{"cert sign", "crl sign", "digital signature"}}
var EnvConfig EnvConfigStruct
