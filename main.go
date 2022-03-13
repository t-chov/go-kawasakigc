package main

import (
	"bytes"
	_ "embed"
	"fmt"

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
	db, err := db.InitDb(csvBytes)
	if err != nil {
		panic(err)
	}
	garbage, _ := db.Find("IH調理器")
	fmt.Println(pretty(*garbage))
}
