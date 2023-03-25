// certificateManager
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// configStructs.go, jfgratton : 2023-03-20

package configs

import "net"

type BaseCertinfo struct {
	CountryCode            string
	StateOrProvinceName    string
	Location               string
	OrganizationalName     string
	OrganizationalUnitName string
	EmailAddress           string
}

type RootCAconfig struct {
	CommonName    string
	ValidForYears int
	DNSNames      []string
	IPAddresses   []net.IP
	KeyFilePath   string
	CertFilePath  string
	Basefilename  string
	BasicInfos    BaseCertinfo
}
