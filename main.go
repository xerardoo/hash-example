package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("===== Welcome =====")
	fmt.Println("\n Please select an option:")
	fmt.Println("\n1) A URL")
	fmt.Println("2) A destination path for a file containing a hash")
	fmt.Println("3) An optional value for throttling the download")
	fmt.Println("4) Exit")

	for {
		reader := bufio.NewReader(os.Stdin)
		option, _ := reader.ReadString('\n')

		switch strings.TrimSuffix(option, "\n") {
		case "1":
			fmt.Println("Enter URL:")
			url, _ := reader.ReadString('\n')
			url = strings.TrimSuffix(url, "\n")
			break
		case "2":
			fmt.Println("Enter Path:")
			path, _ := reader.ReadString('\n')
			path = strings.TrimSuffix(path, "\n")
			break
		case "3":
			var limit int
			fmt.Println("Enter Max File Size [MB]:")
			fmt.Scan(&limit)
			fmt.Println(limit)
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
