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

const fileJSON = "users.json"
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

	users, err := NewUsers(args["fileName"])
	if err != nil {
		return err
	}

	if args["operation"] == "add" {
		if args["item"] == "" {
			return errors.New("-item flag has to be specified")
		}
		err := users.Add(args["item"])
		if err != nil {
			writer.Write([]byte(err.Error()))
			return nil
		}
	} else if args["operation"] == "list" {
		bytes, err := users.List()
		if err != nil {
			return err
		}
		writer.Write(bytes)

	} else if args["operation"] == "findById" {
		if args["id"] == "" {
			return errors.New("-id flag has to be specified")
		}
		bytes, err := users.FindById(args["id"])
		if err != nil {
			return err
		}
		writer.Write(bytes)
		return nil

	} else if args["operation"] == "remove" {
		if args["id"] == "" {
			return errors.New("-id flag has to be specified")
		}

		err := users.Remove(args["id"])
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

func NewUsers(fileName string) (Users, error) {

	users := Users{}

	if fileName != "" {
		_, err := users.readeFile(fileName)
		if err != nil {
			return nil, err
		}
	}

	return users, nil

}

func (u *Users) readeFile(fileName string) (Users, error) {
	users := Users{}
	file, err := os.OpenFile(fileJSON, os.O_RDONLY|os.O_CREATE, fPermission)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	jsonBlob, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBlob, u)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func writeToFile(value Users) error {

	file, err := os.OpenFile(fileJSON, os.O_RDWR, fPermission)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	file.Write(bytes)
	// os.Stdout.Write(bytes)
	return nil
}

func (u *Users) Add(userJSON string) error {

	user := User{}

	err := json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		return err
	}

	*u = append(*u, user)

	err = writeToFile(*u)
	if err != nil {
		return err
	}
	return nil
}

func (u *Users) List() ([]byte, error) {

	if len(*u) == 0 {
		return nil, nil
	}

	bytes, err := json.Marshal(*u)
	if err != nil {
		return nil, err
	}

	return bytes, nil

}

func (u *Users) FindById(id string) ([]byte, error) {

	for _, user := range *u {
		if user.Id == id {
			data, err := json.Marshal(&user)
			if err != nil {
				return nil, err
			}
			return data, nil
		}
	}
	return nil, nil
}

func (u *Users) Remove(id string) error {

	_, err := u.FindById(id)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Item with id %s not found", id))
	}

	for i, user := range *u {
		if user.Id == id {
			*u = append((*u)[:i], (*u)[i+1:]...)
			return nil
		}
	}
	err = writeToFile(*u)
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
