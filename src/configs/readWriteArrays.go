// certificateManager
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// readWriteArrays.go, jfgratton : 2023-03-20

package configs

//import (
//"encoding/json"
//"fmt"
//"os"
//)
//
//type MyStruct struct {
//	MyInt   int
//	MySlice []string
//	MyBool  bool
//}
//
//func main() {
//	myData := MyStruct{
//		MyInt:   42,
//		MySlice: []string{"apple", "banana", "cherry"},
//		MyBool:  true,
//	}
//
//	file, err := os.Create("data.json")
//	if err != nil {
//		fmt.Println("Error creating file:", err)
//		return
//	}
//	defer file.Close()
//
//	jsonData, err := json.Marshal(myData)
//	if err != nil {
//		fmt.Println("Error marshaling data:", err)
//		return
//	}
//
//	_, err = file.Write(jsonData)
//	if err != nil {
//		fmt.Println("Error writing data:", err)
//		return
//	}
//
//	fmt.Println("Data written to data.json")
//}

/*
This program creates a struct called MyStruct with three fields: an integer (MyInt), a slice of strings (MySlice),
and a boolean (MyBool). It then creates an instance of MyStruct called myData, and populates it with some sample data.


Next, the program creates a file called data.json and uses json.Marshal to encode the myData struct into JSON.
It then writes the encoded JSON data to the file.

When you run this program, it will create a new file called data.json in the current directory and write the encoded
JSON data to it.
*/
