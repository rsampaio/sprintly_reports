package main

import (
	"github.com/rsampaio/sprintly_reports"
	"fmt"
	"os"
)

func main() {
	client := sprintly_reports.NewClient(os.Args[1], os.Args[2]])
	resp, _ := client.Products()
	fmt.Println(resp)
}
