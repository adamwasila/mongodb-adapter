package main

import (
	"fmt"
	"os"

	mongodbadapter "github.com/adamwasila/mongodb-adapter/v2"

	"github.com/casbin/casbin/v2"
)

func main() {
	ce, err := casbin.NewEnforcer()
	if err != nil {
		fmt.Printf("Fatal error: %v\n", err)
		os.Exit(1)
	}
	mo, err := casbin.NewModel("./model.conf", "")
	if err != nil {
		fmt.Printf("Fatal error: %v\n", err)
		os.Exit(1)
	}
	ma := mongodbadapter.NewAdapter("mongodb://localhost:27017/casbin", mongodbadapter.OptionCollectionName("casbin_rule"))
	ce.InitWithModelAndAdapter(mo, ma)

	fmt.Printf("Policies: %v\n", ce.GetPolicy())
}
