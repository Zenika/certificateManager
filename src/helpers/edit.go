// certificateManager : Écrit par Jean-François Gratton (jean-francois@famillegratton.net)
// src/helpers/edit.go
// 4/29/23 17:36:16

package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetDuration(prompt string, value *int) {
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

func GetStringValFromPrompt(prompt string, value *string) {
	inputScanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%s [%s]: ", prompt, *value)
	inputScanner.Scan()
	nval := inputScanner.Text()

	if nval != "" {
		*value = nval
	}
}

func GetStringSliceFromPrompt(prompt string, valuesPointer *[]string) {
	slice := *valuesPointer
	scanner := bufio.NewScanner(os.Stdin)

	// If slice is empty, prompt for first element
	if len(slice) == 0 {
		fmt.Printf("%s [], ENTER to ignore: ", prompt)
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			return
		}
		slice = append(slice, input)
	}

	// Update existing elements
	for i := range slice[1:] {
		fmt.Println("A value of '' (empty string) means that we keep the current value")
		fmt.Println("A value of '.' means an empty string")
		fmt.Printf("A value of '.. ' means that we exit of this loop. %s [%s]: ", prompt, slice[i])
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			continue
		}
		if input == "." {
			slice[i] = ""
		}
		if input == ".." {
			*valuesPointer = slice
			return
		}
	}

	// Prompt for new elements
	for {
		fmt.Print("Enter value for new element: ")
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			return
		}
		slice = append(slice, input)
	}

	*valuesPointer = slice
}

func GetKeyUsageFromPrompt() []string {
	inputScanner := bufio.NewScanner(os.Stdin)
	ku := []string{"decipher only", "encipher only", "crl sign", "cert sign", "key agreement",
		"data encipherment", "key encipherment", "content commitment", "digital signature"}
	inputs := []string{}

	fmt.Println("The valid key usage values are:")
	for i, j := range ku {
		if i%5 == 0 && i != 0 {
			fmt.Println()
		}
		fmt.Printf("'%s' ", White(j))
	}
	fmt.Println()
	for {
		input := ""
		fmt.Print("Please enter a value from the above, just press ENTER to end : ")
		inputScanner.Scan()
		input = inputScanner.Text()
		if input == "" {
			break
		}
		if valueInList(input, ku) {
			inputs = append(inputs, input)
		}
	}
	// if the array is empty, we return a default value
	if len(inputs) == 0 {
		return []string{"digital signature"}
	}
	// now we need to ensure that we do not have any duplicates
	s := make([]string, 0, len(inputs))
	m := make(map[string]bool)

	for _, value := range inputs {
		if _, ok := m[value]; !ok {
			m[value] = true
			s = append(s, value)
		}
	}
	return s
}

func valueInList(in string, list []string) bool {
	for _, x := range list {
		if x == in {
			return true
		}
	}
	return false
}
