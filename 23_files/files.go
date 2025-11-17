package main

import "os"

func main() {
	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }
	// fileInfo, err := f.Stat()

	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// // Meta data about files
	// fmt.Println("file name: ", fileInfo.Name())
	// fmt.Println("isDir: ", fileInfo.IsDir())
	// fmt.Println("file size: ", fileInfo.Size(), "bytes")
	// fmt.Println("file permission: ", fileInfo.Mode())
	// fmt.Println("file modified at: ", fileInfo.ModTime())

	// // content data of file
	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// buf := make([]byte, 10)

	// d, err := f.Read(buf)

	// if err != nil {
	// 	panic(err)
	// }

	// for i := range len(buf) {
	// 	println("data: ", string(buf[i]))
	// }
	// println("data: ", d, string(buf))

	// // something simple => not recommended
	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// println(string(data))

	// // read folders
	// // dir, err := os.Open(".")
	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)

	// for _, fi := range fileInfo {
	// 	println(fi.Name(), fi.IsDir())
	// }

	// create a file
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// f.WriteString("hi go ")
	// f.WriteString("nice language")

	// bytes := []byte("Hello Golang")
	// n, err := f.Write(bytes)
	// if err != nil {
	// 	panic(err)
	// }

	// println("length of data: ", n)

	// read and write to another file (streaming fashion)

	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }

	// writer.Flush()

	// println("written to new file successfully...")

	// delete a file

	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

	println("file deleted successfully...")
}
