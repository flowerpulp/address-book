package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

type entry struct {
	Name         string `json:"Name"`
	Email        string `json:"E-Mail,omitempty"`
	Telefon      string `json:"Telefon,omitempty"`
	Birthdate    string `json:"Geburtsdatum,omitempty"`
}

func sortJsonFile(fp string) error {
	f, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	fmt.Println(f)
	fmt.Println(string(b))

	e := &[]entry{}
	fmt.Println(e)
	err = json.Unmarshal(b, e)
	if err != nil {
		return err
	}

	fmt.Println(e)

	bb, err := json.MarshalIndent(e, "", "        ")
	if err != nil {
		return err
	}

	fmt.Println(string(bb))

	return nil
}

func main() {
	err := sortJsonFile("book.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
