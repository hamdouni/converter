package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gohouse/converter"
)

func main() {
	parser()
}

func parser() {
	dsn := flag.String("dsn", "", "database dsn configuration")
	file := flag.String("file", "", "save route")
	table := flag.String("table", "", "table to migrate")
	realNameMethod := flag.String("realNameMethod", "", "The table name corresponding to the structure")
	dateToTime := flag.Bool("dateToTime", true, "Whether to convert date to Time, default true")
	packageName := flag.String("packageName", "model", "Generated struct package name")
	tagKey := flag.String("tagKey", "orm", "the key of the field tag")
	prefix := flag.String("prefix", "", "table prefix")
	version := flag.Bool("version", false, "version number")
	v := flag.Bool("v", false, "version number")
	enableJsonTag := flag.Bool("enableJsonTag", false, "Whether to add json tag, default false")
	h := flag.Bool("h", false, "help")
	help := flag.Bool("help", false, "help")

	// start
	flag.Parse()

	if *h || *help {
		flag.Usage()
		return
	}

	// version number
	if *version || *v {
		fmt.Printf("\n version: %s\n %s\n using -h param for more help \n", converter.VERSION, converter.VERSION_TEXT)
		return
	}

	// initialization
	t2t := converter.NewTable2Struct()
	// personalized configuration
	t2t.Config(&converter.T2tConfig{
		// If the first letter of the field is originally capitalized, the tag will not be added. By default, false is added, and true is not added.
		RmTagIfUcFirsted: false,
		// Whether the field name of the tag is converted to lowercase, if it has uppercase letters, the default false is not converted
		TagToLower: false,
		// When the first letter of the field is capitalized, whether to convert other letters to lowercase, the default false is not converted
		UcFirstOnly: false,
		//// Put each struct into a separate file, the default is false, put into the same file (not provided yet)
		//SeperatFile: false,
	})
	// Start migration
	err := t2t.
		// Specify a table, if not specified, all tables will be migrated by default
		Table(*table).
		// table prefix
		Prefix(*prefix).
		// Whether to add json tag
		EnableJsonTag(*enableJsonTag).
		// The package name of the generated struct (if it is empty by default, it will be named: package model)
		PackageName(*packageName).
		// The key value of the tag field, the default is orm
		TagKey(*tagKey).
		// Whether to add a structure method to get the table name
		RealNameMethod(*realNameMethod).
		// Generated structure save path
		SavePath(*file).
		// database dsn
		Dsn(*dsn).
		// use time for sql date
		DateToTime(*dateToTime).
		// and run
		Run()

	if err != nil {
		log.Println(err.Error())
	}
}
