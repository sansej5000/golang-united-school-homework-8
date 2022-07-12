package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type Users []User

type Arguments map[string]string

const fPermission = 0644

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}

func Perform(args Arguments, writer io.Writer) error {

	if args["fileName"] == "" {
		return errors.New("-fileName flag has to be specified")
	}
	if args["operation"] == "" {
		return errors.New("-operation flag has to be specified")
	}

	if args["operation"] == "add" {
		if args["item"] == "" {
			return errors.New("-item flag has to be specified")
		}
		err := Add(args["item"], args["fileName"], writer)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return nil
		}
	} else if args["operation"] == "list" {
		bytes, err := List(args["fileName"])
		if err != nil {
			return err
		}
		writer.Write(bytes)

	} else if args["operation"] == "findById" {
		if args["id"] == "" {
			return errors.New("-id flag has to be specified")
		}
		bytes, err := FindById(args["id"], args["fileName"], writer)
		if err != nil {
			return err
		}
		writer.Write(bytes)
		return nil

	} else if args["operation"] == "remove" {
		if args["id"] == "" {
			return errors.New("-id flag has to be specified")
		}

		err := Remove(args["id"], args["fileName"], writer)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return nil
		}

	} else {
		return fmt.Errorf(fmt.Sprintf("Operation %s not allowed!", args["operation"]))
	}

	return nil
}

func readefromFile(fileName string) (Users, error) {
	users := Users{}
	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, fPermission)
	if err != nil {
		return nil, err
	}
	jsonBlob, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonBlob, &users)
	if err != nil {
		return nil, err
	}
	file.Close()
	return users, nil
}

func writeToFile(value Users, fileName string) error {

	file, err := os.OpenFile(fileName, os.O_RDWR, fPermission)
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fileName, bytes, 0777)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

func Add(item string, fileName string, writer io.Writer) error {

	user := User{}

	err := json.Unmarshal([]byte(item), &user)
	if err != nil {
		return err
	}

	data, _ := FindById(user.Id, fileName, writer)
	if data != nil {
		return fmt.Errorf(fmt.Sprintf("Item with id %s already exists", user.Id))
	}

	users, _ := readefromFile(fileName)

	users = append(users, user)

	err = writeToFile(users, fileName)
	if err != nil {
		return err
	}
	return nil
}

func List(fileName string) ([]byte, error) {

	users, _ := readefromFile(fileName)

	bytes, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}

	return bytes, nil

}

func FindById(id string, fileName string, writer io.Writer) ([]byte, error) {

	users, _ := readefromFile(fileName)

	for _, user := range users {
		if user.Id == id {
			data, err := json.Marshal(user)
			if err != nil {
				return nil, err
			}
			return data, nil
		}
	}
	return nil, nil
}

func Remove(id string, fileName string, writer io.Writer) error {
	dat, _ := FindById(id, fileName, writer)
	if dat == nil {
		return fmt.Errorf(fmt.Sprintf("Item with id %s not found", id))
	}

	data, _ := readefromFile(fileName)

	for index, user := range data {
		if user.Id == id {
			data = append(data[:index], data[index+1:]...)
		}
	}
	err := writeToFile(data, fileName)
	if err != nil {
		return err
	}
	return nil
}

func parseArgs() Arguments {

	flagId := flag.String("id", "", "User Id")
	flagOperation := flag.String("operation", "", "Operation")
	flagItem := flag.String("item", "", "Item")
	flagFileName := flag.String("fileName", "", "Path to file")

	flag.Parse()

	return Arguments{
		"id":        *flagId,
		"operation": *flagOperation,
		"item":      *flagItem,
		"fileName":  *flagFileName,
	}
}
