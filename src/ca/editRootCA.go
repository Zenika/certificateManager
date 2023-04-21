// certificateManager : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/ca/editRootCA.go
// 4/20/23 18:10:04

package ca

import (
	"certificateManager/config"
	"fmt"
	"os"
)

func EditCACertificate() error {
	var err error
	config.CertConfig, err = config.Json2Config()
	if err != nil {
		if err != os.ErrNotExist {
			return err
		}
	}

	err = prompt4values(&config.CertConfig)
	if err != nil {
		return err
	}
	return nil
}

func prompt4values(cfg *config.CertConfigStruct) error {
	fmt.Print(`
You will now be prompted to provide values to all of the fields that should
be part of your config file. If a prompt shows a value between [brackets],
this means that this value is either already present, or a suggested default
value that can be accepted by just pressing ENTER.\n`)

	// this is beyond ugly....

	getStringValFromPrompt("Please enter the certificate name", &cfg.CertificateName)
	getStringValFromPrompt("Please enter the certificate rootdir", &cfg.CertificateDirectory)
	getStringValFromPrompt("Please enter the country (C)", &cfg.Country)
	getStringValFromPrompt("Please enter the province/state (ST)", &cfg.Province)
	getStringValFromPrompt("Please enter the organization (O)", &cfg.Organization)
	getStringValFromPrompt("Please enter the organizational unit (OU)", &cfg.OrganizationalUnit)
	getStringValFromPrompt("Please enter the organizational unit (OU)", &cfg.OrganizationalUnit)
	getStringValFromPrompt("Please enter the common name (CN)", &cfg.CommonName)
	return nil
}

func getStringValFromPrompt(prompt string, value *string) {
	nval := ""
	fmt.Printf("%s [%s]: ", prompt, *value)
	fmt.Scanln(&nval)

	if nval != "" {
		*value = nval
	}
}
