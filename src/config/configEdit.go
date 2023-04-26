// certificateManager : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/ca/editRootCA.go
// 4/20/23 18:10:04

package config

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func EditConfig() error {
	var err error
	CertConfig, err = Json2Config()
	if err != nil {
		if !os.IsNotExist(err) {
			//if err != os.ErrNotExist {
			return err
		}
	}
	err = prompt4values(&CertConfig)
	if err != nil {
		return err
	}

	// we now need to reinject the config in a json
	err = Config2Json(CertConfig)
	if err != nil {
		return err
	}
	return nil
}

func prompt4values(cfg *CertConfigStruct) error {
	fmt.Println(`
You will now be prompted to provide values to all of the fields that should
be part of your config file. If a prompt shows a value between [brackets],
this means that this value is either already present, or a suggested default
value that can be accepted by just pressing ENTER.`)

	// this is beyond ugly....
	retVal := ""
	fmt.Print("\nIs this a certificate authority (CA) ?\nPlease enter one of the following for true: true, t, 1.")
	fmt.Print("Any other answer will be treated as false: ")
	fmt.Scanln(&retVal)
	if strings.ToLower(retVal) == "true" || strings.ToLower(retVal) == "t" || retVal == "1" {
		cfg.IsCA = true
	}
	getDuration("Please enter the certification duration, in years.\nAn invalid duration would default to 1 year", &cfg.Duration)
	getStringValFromPrompt("Please enter the certificate name", &cfg.CertificateName)
	getStringValFromPrompt("Please enter the certificate rootdir", &cfg.CertificateDirectory)
	getStringValFromPrompt("Please enter the country (C)", &cfg.Country)
	getStringValFromPrompt("Please enter the province/state (ST)", &cfg.Province)
	getStringValFromPrompt("Please enter the locality (L)", &cfg.Locality)
	getStringValFromPrompt("Please enter the organization (O)", &cfg.Organization)
	getStringValFromPrompt("Please enter the organizational unit (OU)", &cfg.OrganizationalUnit)
	getStringValFromPrompt("Please enter the common name (CN)", &cfg.CommonName)

	// A non-CA cert should not have KeyUsage

	//if cfg.IsCA {
	//	cfg.KeyUsage = getKeyUsageFromPrompt()
	//} else {
	//	cfg.KeyUsage = []string{}
	//}
	getStringSliceFromPrompt("Please enter the email address(es) to be included in this certicate", &cfg.EmailAddresses)
	getStringSliceFromPrompt("Please enter the DNS name(s) to be included in this certicate", &cfg.DNSNames)
	getStringSliceFromPrompt("Please enter the comments to be included in this certicate\n(Note: those are for documentation purposes only, not part of the cert)", &cfg.Comments)

	// Still need net.IP...
	netip := []string{}
	if len(cfg.IPAddresses) > 0 {
		for _, x := range cfg.IPAddresses {
			netip = append(netip, x.String())
		}
	}
	getStringSliceFromPrompt("Please enter the IP address(es) to be included in this certicate\n(Note: this is NOT recommended in a CA)", &netip)
	cfg.IPAddresses = []net.IP{}
	if len(netip) > 0 {
		for _, x := range netip {
			cfg.IPAddresses = append(cfg.IPAddresses, net.ParseIP(x))
		}
	}

	return nil
}

func getDuration(prompt string, value *int) {
	var err error
	inputScanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s [%d]: ", prompt, *value)
	inputScanner.Scan()
	nval := inputScanner.Text()

	if nval != "" {
		*value, err = strconv.Atoi(nval)
		if err != nil {
			*value = 1
		}
	}
}

func getStringValFromPrompt(prompt string, value *string) {
	inputScanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s [%s]: ", prompt, *value)
	inputScanner.Scan()
	nval := inputScanner.Text()

	if nval != "" {
		*value = nval
	}
}

func getStringSliceFromPrompt(prompt string, values *[]string) {
	inputScanner := bufio.NewScanner(os.Stdin)
	//proceed := true
	for {
		for _, input := range *values {
			fmt.Printf("%s [%s]\n", prompt, input)
			fmt.Println("Press ENTER to keep the current value, enter a single dot to leave line empty, and two dots to stop ")
			inputScanner.Scan()
			nval := inputScanner.Text()
			if nval == ".." {
				break
			} else {
				if nval == "." {
					nval = input
				}
			//*value = append(*value, nval)
		}
	}
}

//func getKeyUsageFromPrompt() []string {
//	inputScanner := bufio.NewScanner(os.Stdin)
//	ku := []string{"decipher only", "encipher only", "crl sign", "cert sign", "key agreement",
//		"data encipherment", "key encipherment", "content commitment", "digital signature"}
//	inputs := []string{}
//
//	fmt.Println("The valid key usage values are:")
//	for i, j := range ku {
//		if i%5 == 0 && i != 0 {
//			fmt.Println()
//		}
//		fmt.Printf("'%s' ", misc.White(j))
//	}
//	fmt.Println()
//	for {
//		input := ""
//		fmt.Print("Please enter a value from the above, just press ENTER to end : ")
//		inputScanner.Scan()
//		input = inputScanner.Text()
//		if input == "" {
//			break
//		}
//		if valueInList(input, ku) {
//			inputs = append(inputs, input)
//		}
//	}
//	// if the array is empty, we return a default value
//	if len(inputs) == 0 {
//		return []string{"digital signature"}
//	}
//	// now we need to ensure that we do not have any duplicates
//	s := make([]string, 0, len(inputs))
//	m := make(map[string]bool)
//
//	for _, value := range inputs {
//		if _, ok := m[value]; !ok {
//			m[value] = true
//			s = append(s, value)
//		}
//	}
//	return s
//}
//
//func valueInList(in string, list []string) bool {
//	for _, x := range list {
//		if x == in {
//			return true
//		}
//	}
//	return false
//}