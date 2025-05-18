package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fileInfo, err := f.Stat()
	// if(err != nil){
	// 	panic(err)
	// }

	// fmt.Println(fileInfo)
	// fmt.Println("File name", fileInfo.Name())
	// fmt.Println("File or folder", fileInfo.IsDir())
	// fmt.Println("File size", fileInfo.Size())
	// fmt.Println("file permission", fileInfo.Mode())
	// fmt.Println("file modified at", fileInfo.ModTime())

	//read File

	// f, err := os.Open("example.txt")
	// if err != nil{
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte, 20)

	// d, err := f.Read(buf)
	// if(err != nil){
	// 	panic(err)
	// }
	// for i:= 0; i < len(buf); i++{
	// 	println("data", d, string(buf[i]))
	// }

	// f, err := os.ReadFile("example.txt") //not to be used as loads the file completely at once
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(f))

	//reading folders
	// dir, err := os.Open("../")
	// if err != nil{
	// 	panic(err)
	// }
	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1) //when <0 then returns all

	// for _, fi := range fileInfo{
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	//create a file

	f, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// f.WriteString("hi golang")
	// f.WriteString("nice golang")
	// bytes := []byte("hello go lang")
	// f.Write(bytes)

	//read and write to another file streaming fashion

	sourceFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create("destiFile.txt")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		er := writer.WriteByte(b)
		if er != nil {
			panic(er)
		}
	}
	writer.Flush()

	fmt.Println("written to new file successfully")

	//deleting file
	eror := os.Remove("example2.txt")
	if eror != nil{
		panic(eror)
	}
	fmt.Println("File deleted successfully")

}
