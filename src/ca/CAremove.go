// certificateManager : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/ca/removeRootCA.go
// 4/22/23 08:55:20

package ca

import (
	"cm/helpers"
	"os"
	"path/filepath"
)

// This is a stub, really, before we get to the actual removal in branch 0.600
func RemoveCACertificate() error {
	cfg, err := helpers.Json2Config()

	if err != nil {
		return err
	}

	err = os.Remove(filepath.Join(cfg.CertificateDirectory, cfg.CertificateName, ".key"))
	if err != nil {
		return err
	}
	err = os.Remove(filepath.Join(cfg.CertificateDirectory, cfg.CertificateName, ".crt"))

	return err
}
