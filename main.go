package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"

	"github.com/t-chov/kawasakigc/db"
)

//go:embed go_gc.csv
var csvBytes []byte

func pretty(g db.Garbage) string {
	var buf bytes.Buffer

	if len(g.DetailType) > 0 {
		buf.WriteString(fmt.Sprintf("%s(%s)\n", g.GarbageType, g.DetailType))
	} else {
		buf.WriteString(g.GarbageType + "\n")
	}
	if len(g.Description) > 0 {
		buf.WriteString(g.Description + "\n")
	}
	if len(g.Url) > 0 {
		buf.WriteString(fmt.Sprintf("URL: %s", g.Url))
	}
	return buf.String()
}

func main() {
	dbp, err := db.InitDb(csvBytes)
	if err != nil {
		panic(err)
	}

	args := os.Args
	if len(args) == 1 {
		fmt.Fprintf(os.Stderr, "need arguments to search.\n")
		os.Exit(1)
	}

	db := *dbp
	garbage, found := db.Find(args[1])
	if !found {
		fmt.Fprintf(os.Stderr, "not found %s.\n", args[1])
		os.Exit(1)
	}
	fmt.Println(pretty(*garbage))
}
