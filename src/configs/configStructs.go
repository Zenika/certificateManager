// certificateManager
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// configStructs.go, jfgratton : 2023-03-20

package configs

import "net"

//type RootCAconfig struct {
//	CommonName    string
//	ValidForYears int
//	DNSNames      []string
//	IPAddresses   []net.IP
//	KeyFilePath   string
//	CertFilePath  string
//	Basefilename  string
//	BasicInfos    BaseCertinfo
//}

type RootCAconfig struct {
	Country              string   `json:"C"`
	Province             string   `json:"ST"`
	Locality             string   `json:"L"`
	Organization         string   `json:"O"`
	OrganizationalUnit   string   `json:"OU,omitempty"`
	CommonName           string   `json:"CN"`
	EmailAddresses       []string `json:"Email,omitempty"`
	Duration             int      `json:"Duration"`
	Usage                []string `json:"Usage"`
	DNS                  []string `json:"DNS,omitempty"`
	IPS                  []net.IP `json:"IPS,omitempty"`
	CertificateDirectory string   `json:"CertificateDirectory"`
	CertificateName      string   `json:"CertificateName"`
}
