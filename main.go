package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/xerardoo/hash-example/file"
	"github.com/xerardoo/hash-example/imt"
	"os"
	"strings"
)

func main() {
	fmt.Println("===== Welcome =====")
	var newFile file.File

	for {
		fmt.Println("\n\n")
		fmt.Println("1) A URL")
		fmt.Println("2) A destination path for a file containing a hash")
		fmt.Println("3) An optional value for throttling the download")
		fmt.Println("4) Exit")
		fmt.Println("\n Please select an option:")
		reader := bufio.NewReader(os.Stdin)
		option, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Err: %s\n", err.Error())
			return
		}

		switch strings.TrimSuffix(option, "\n") {
		case "1":
			fmt.Println("Enter URL:")
			url, _ := reader.ReadString('\n')
			url = strings.TrimSuffix(url, "\n")

			_, err = newFile.Fetch(url)
			if err != nil {
				fmt.Printf("Err: %s\n", err.Error())
				break
			}
			newFile.Data, err = imt.Generate(newFile.Data)
			if err != nil {
				fmt.Printf("Err: %s\n", err.Error())
				break
			}
			fmt.Println("File fetched")
			break
		case "2":
			fmt.Println("Enter Path:")
			path, _ := reader.ReadString('\n')
			path = strings.TrimSuffix(path, "\n")

			newFile.Data = []byte(hex.EncodeToString(newFile.Data[:]))
			err = newFile.Write(path)
			if err != nil {
				fmt.Printf("Err: %s\n", err.Error())
				break
			}
			fmt.Println("File saved")
			break
		case "3":
			var limit float64
			fmt.Println("Enter Max File Size [MB]:")
			fmt.Scan(&limit)

			err = file.SetLimitSize(limit)
			if err != nil {
				fmt.Printf("Err: %s\n", err.Error())
				break
			}
			fmt.Println("Modified file size limit")
			break
		case "4":
			fmt.Println("Are you sure? [Y/n]")
			response, _ := reader.ReadString('\n')
			if strings.ToLower(response) == "y\n" {
				os.Exit(0)
			}
			break
		default:
			fmt.Println("Option invalid, try again...")
		}
	}
}
