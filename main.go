package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



// kiritishda : go run main.go file1.txt file2.txt file3.txt kiritish kerak 



func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("Iltimos, biror fayl nomi kiriting.")
		return
	}
	

	var birlashtirilgan string
	for _, fayl := range files {
		file, err := os.Open(fayl)
		if err != nil {
			fmt.Printf("Xatolik: %v\n", err)
			continue
		}


		defer file.Close()



		skaner := bufio.NewScanner(file)
		for skaner.Scan() {
			birlashtirilgan += skaner.Text() + "\n"
		}
		if err := skaner.Err(); err != nil {
			fmt.Printf("Xatolik: %v\n", err)
		}
	}

	fmt.Println("Birlashtirilgan malumot:")
	fmt.Println(strings.TrimSpace(birlashtirilgan))
}
