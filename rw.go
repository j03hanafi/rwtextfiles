package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
) 
  
func CreateFile(fileName string) { 
  
	fmt.Printf("Writing to a file in Go lang\n") 
	
	if !strings.Contains(fileName, ".txt") {
		fileName += ".txt"
	}
	
	file, err := os.Create(fileName) 
		
	if err != nil { 
			log.Fatalf("failed creating file: %s", err) 
	} 
	
	defer file.Close() 
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	content := scanner.Text()

	len, err := file.WriteString(content) 

	if err != nil { 
			log.Fatalf("failed writing to file: %s", err) 
	} 

	fmt.Printf("\nFile Name: %s", file.Name()) 
	fmt.Printf("\nLength: %d bytes", len) 

	fmt.Printf("\nPress any key to back\n")

	var anyKey string
	fmt.Scanln(&anyKey)
	
	if anyKey != "" {
		fmt.Printf("\n\n")
		Init()
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
  
func ReadFile(fileName string) { 
  
	fmt.Printf("\n\nReading a file in Go lang\n") 
	
	data, err := ioutil.ReadFile(fileName) 
	if err != nil { 
			log.Panicf("failed reading data from file: %s", err) 
	} 
	fmt.Printf("\nFile Name: %s", fileName) 
	fmt.Printf("\nSize: %d bytes", len(data)) 
	fmt.Printf("\nData: %s", data) 

	fmt.Printf("\nPress any key to back\n")
	
	var anyKey string
	fmt.Scanln(&anyKey)

	if anyKey != "" {
		fmt.Printf("\n\n")
		Init()
	} else {
		fmt.Printf("\n\n")
		Init()
	}
  
} 
	
func Init() {
	fmt.Println("Read or Write a text file program")
	fmt.Println("1. Read a text file")
	fmt.Println("2. Write a text file")
	fmt.Println("3. Exit")
	fmt.Println("Choose 1-3 : ")

	var choose int

	fmt.Scanln(&choose)

	fmt.Println()
	var namaFile string

	if choose == 1 {
		fmt.Println("Type your file's name")
		fmt.Scanln(&namaFile)
		if FuncCheckExist(namaFile) {
			ReadFile(namaFile)
		} else {
			fmt.Printf("\nFile not found\n")
			Init()
		}

	} else if choose == 2 {
		fmt.Println("Type your file's name")
		fmt.Scanln(&namaFile)
		CreateFile(namaFile) 
	} else if choose == 3 {
		// Exit()
	}	else {
		fmt.Println("Please choose 1-3")
		fmt.Printf("\n")
		Init()
	}
}

func main() { 
	Init()
}