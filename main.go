package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Arguments map[string]string

const filePermission1 = 0644

func Perform(args Arguments, writer io.Writer) error {

	operation := args["operation"]
	fileName := args["fileName"]
	item := args["item"]
	id := args["id"]

	if operation == "" {
		return fmt.Errorf("-operation flag has to be specified")
	}

	if fileName == "" {
		return fmt.Errorf("-fileName flag has to be specified")
	}

	switch {
	case operation == "list":
		file, err := os.OpenFile(args["fileName"], os.O_RDWR, filePermission1)
		if err != nil {
			return fmt.Errorf("-fileName flag has to be specified")
		}

		bytes, err := ioutil.ReadAll(file)

		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("read %d bytes: %q\n", count, data[:count])
		// fmt.Printf("Operation to be performed %s \n", operation)
		writer.Write(bytes)

	case operation == "add":

		if item == "" {
			return fmt.Errorf("-item flag has to be specified")
		}

		file, err := os.OpenFile(args["fileName"], os.O_RDWR|os.O_CREATE, filePermission1)
		if err != nil {
			log.Fatal(err)
		}

		data := make([]byte, 4096)
		count, err := file.Read(data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("read %d bytes: %q\n", count, data[:count])
		fmt.Printf("Operation to be performed %s \n", operation)

	case operation == "findById":
		file, err := os.OpenFile(args["fileName"], os.O_RDWR, filePermission1)
		if err != nil {
			return fmt.Errorf("-fileName flag has to be specified")
		}
		data := make([]byte, 4096)
		count, err := file.Read(data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("read %d bytes: %q\n", count, data[:count])

	case operation == "remove":
		if id == "" {
			return fmt.Errorf("-id flag has to be specified")
		}
	default:
		return fmt.Errorf("Operation %s not allowed!", operation)
	}

	return nil
}

func parseArgs() map[string]string {
	opr := flag.String("operation", "", "Operation on the file")
	fileTobeListed := flag.String("fileName", "", "File to be listed")
	flag.Parse()
	// fmt.Printf("Operation is %s, fileName is %s \n", *opr, *fileTobeListed)

	var Arguments = make(map[string]string)
	Arguments["fileName"] = *fileTobeListed
	Arguments["operation"] = *opr
	Arguments["id"] = ""
	Arguments["item"] = ""

	return Arguments
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
