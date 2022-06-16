package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type Users []User

type Arguments map[string]string

// const fileJSON = "users.json"
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
		err := Add(args["item"], args["fileName"], args["id"], writer)
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
		// return errors.New(fmt.Sprintf("Operation %s not allowed!", args["operation"]))
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

	defer file.Close()

	jsonBlob, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBlob, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func writeToFile(value Users, fileName string) error {

	file, err := os.OpenFile(fileName, os.O_RDWR, fPermission)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	file.Write(bytes)
	return nil
}

//Adding new item
func Add(item string, fileName string, id string, writer io.Writer) error {

	data, _ := FindById(id, fileName, writer)
	if data != nil {
		return fmt.Errorf(fmt.Sprintf("Item with id %s already exists", id))
	}

	user := User{}

	err := json.Unmarshal([]byte(item), &user)
	if err != nil {
		return err
	}

	users, err := readefromFile(fileName)

	users = append(users, user)

	err = writeToFile(users, fileName)
	if err != nil {
		return err
	}
	return nil
}

func List(fileName string) ([]byte, error) {

	users, err := readefromFile(fileName)

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

	_, err := FindById(id, fileName, writer)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Item with id %s not found", id))
	}

	data, _ := readefromFile(fileName)

	for i, user := range data {
		if user.Id == id {
			data = append((data)[:i], (data)[i+1:]...)
			return nil
		}
	}
	err = writeToFile(data, fileName)
	if err != nil {
		return err
	}
	return nil
}

func parseArgs() Arguments {

	// var flagFilename, flagOperation, flagItem, flagId string
	// flag.StringVar(&flagId, "id", "", "identification")
	// flag.StringVar(&flagFilename, "filename", fileJSON, "file name")
	// flag.StringVar(&flagOperation, "operation", "list", "operacions type")
	// flag.StringVar(&flagItem, "item", "", "item")
	// flag.Parse()

	// return Arguments{
	// 	"id":        flagId,
	// 	"operation": flagOperation,
	// 	"item":      flagItem,
	// 	"filename":  flagFilename,
	// }
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
