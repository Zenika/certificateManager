// certificateManager : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/configs/configHelpers.go
// 4/16/23 13:52:19

package configs

import (
	"encoding/json"
	"fmt"
	"os"
)

var CertConfigFile = "rootCA-default.json"
var ServerCertEnvironment = "serverCert-default.json"

func Json2Config() CertConfigStruct {
	var payload CertConfigStruct
	rcDir, _ := os.UserHomeDir()
	rcFile := rcDir + "/.config/certificatemanager/" + CertConfigFile
	jFile, _ := os.ReadFile(rcFile)
	err := json.Unmarshal(jFile, &payload)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return payload
}

func ReadConfigCA() error {
	return nil
}
