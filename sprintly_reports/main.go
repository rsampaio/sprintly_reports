package main

import (
	"github.com/rsampaio/sprintly_reports"
	"fmt"
	"os"
)

func main() {
	client := sprintly_reports.NewClient(os.Args[1], os.Args[2])
	
	resp, _ := client.Products()
	
	items, _ := client.Items(resp[0].Id)

	fmt.Println(items[0].What)

	for _, v := range items {
		fmt.Println(v.Tags)
	}
}










