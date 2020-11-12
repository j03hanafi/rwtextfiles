package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Init() {
	fmt.Println("Read or Write a text file program")
	fmt.Println("1. Read a text file")
	fmt.Println("2. Write a text file")
	fmt.Println("3. Exit")
	fmt.Println("Choose 1-3 : ")

	var choose int

	fmt.Scanln(&choose)

	fmt.Println()

	if choose == 1 {
		FuncRead()
	} else if choose == 2 {
		FuncWrite()
	} else if choose == 3 {
		// Exit()
	}	else {
		fmt.Println("Silahkan pilih 1-3")
		fmt.Printf("\n\n")
		Init()
	}
}

func FuncRead() {
	fmt.Println("Masukan nama file text")
	var namaFile string
	fmt.Scanln(&namaFile)

	if FuncCheckExist(namaFile) {
		ReadContent(namaFile)
		fmt.Println("Press any key to back")
		
		var anyKey string
		fmt.Scanln(&anyKey)

		if anyKey == "" {
			fmt.Printf("\n\n")
			Init()
		} else {
			fmt.Printf("\n\n")
			Init()
		}
	} else {
		fmt.Printf("\n\n")
		Init()
	}
}

func FuncCheckExist(namaFile string) bool {
	info, err := os.Stat(namaFile)
		if os.IsNotExist(err) {
			return false
		}
		return !info.IsDir()
}

func ReadContent(namaFile string) {
	content, err := ioutil.ReadFile(namaFile)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}

func FuncWrite()  {
	fmt.Println("Masukan nama file text")
	var namaFile string
	fmt.Scanln(&namaFile)

	if strings.Contains(namaFile, ".txt") {
		WriteContent(namaFile)
		fmt.Println("Press any key to back")
		
		var anyKey string
		fmt.Scanln(&anyKey)

		if anyKey == "" {
			fmt.Printf("\n\n")
			Init()
		} else {
			fmt.Printf("\n\n")
			Init()
		}
	} else {
		namaFile += ".txt"
		WriteContent(namaFile)
		fmt.Println("Press any key to back")
		
		var anyKey string
		fmt.Scanln(&anyKey)

		if anyKey == "" {
			fmt.Printf("\n\n")
			Init()
		} else {
			fmt.Printf("\n\n")
			Init()
		}
	}
}

func WriteContent(namaFile string) {
	
	fmt.Println("Silakan ketik isi konten")

	f, err := os.Create(namaFile)
	
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	content := scanner.Text()

	// if err2 != nil {
	// 	log.Fatal(err2)
	// }
	fmt.Println("done")
}

func main()  {

	Init()
}