// package main

// import (
// 	"encoding/json"
// 	"errors"
// 	"flag"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// )

// type User struct {
// 	Id    string `json:"id"`
// 	Email string `json:"email"`
// 	Age   int    `json:"age"`
// }

// type Users []User

// type Arguments map[string]string

// const fileJSON = "users.json"
// const fPermission = 0644

// func main() {
// 	err := Perform(parseArgs(), os.Stdout)
// 	if err != nil {
// 		panic(err)
// 		// fmt.Println(err)
// 	}
// }

// func Perform(args Arguments, writer io.Writer) error {

// 	if args["fileName"] == "" {
// 		return errors.New("-fileName flag has to be specified")
// 	}
// 	if args["operation"] == "" {
// 		return errors.New("-operation flag has to be specified")
// 	}

// 	users, err := NewUsers(args["fileName"])
// 	if err != nil {
// 		return err
// 	}

// 	if args["operation"] == "add" {
// 		if args["item"] == "" {
// 			return errors.New("-item flag has to be specified")
// 		}
// 		err := users.Add(args["item"])
// 		if err != nil {
// 			writer.Write([]byte(err.Error()))
// 			return nil
// 		}
// 	} else if args["operation"] == "list" {
// 		bytes, err := users.List()
// 		if err != nil {
// 			return err
// 		}
// 		writer.Write(bytes)

// 	} else if args["operation"] == "findById" {
// 		if args["id"] == "" {
// 			return errors.New("-id flag has to be specified")
// 		}
// 		bytes, err := users.FindById(args["id"])
// 		if err != nil {
// 			return err
// 		}
// 		writer.Write(bytes)
// 		return nil

// 	} else if args["operation"] == "remove" {
// 		if args["id"] == "" {
// 			return errors.New("-id flag has to be specified")
// 		}

// 		err := users.Remove(args["id"])
// 		if err != nil {
// 			writer.Write([]byte(err.Error()))
// 			return nil
// 		}

// 	} else {
// 		// return errors.New(fmt.Sprintf("Operation %s not allowed!", args["operation"]))
// 		// return fmt.Errorf("operation %s not allowed", args["operation"])
// 		return fmt.Errorf(fmt.Sprintf("Operation %s not allowed!", args["operation"]))
// 		// return nil
// 	}

// 	return nil
// }

// func NewUsers(fileName string) (Users, error) {

// 	users := Users{}

// 	if fileName != "" {
// 		err := users.Load(fileName)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return users, nil

// }

// func (u *Users) Load(fileName string) error {

// 	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0644)
// 	if err != nil {
// 		return err
// 	}

// 	defer func() {
// 		err := file.Close()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}()

// 	content, err := io.ReadAll(file)
// 	if err != nil {
// 		return err
// 	}

// 	if len(content) > 0 {
// 		err = json.Unmarshal(content, u)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	return nil

// }

// func (u *Users) readeFile(fileName string) error {
// 	// users := Users{}
// 	file, err := os.OpenFile(fileJSON, os.O_RDONLY|os.O_CREATE, fPermission)
// 	if err != nil {
// 		return err
// 	}

// 	defer file.Close()

// 	jsonBlob, err := io.ReadAll(file)
// 	if err != nil {
// 		return err
// 	}

// 	err = json.Unmarshal(jsonBlob, u)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func writeToFile(Users) error {

// 	var user Users
// 	bytes, err := json.Marshal(user)
// 	if err != nil {
// 		return err
// 	}
// 	os.Stdout.Write(bytes)
// 	return nil
// }

// func (u *Users) Add(userJSON string) error {

// 	user := User{}

// 	err := json.Unmarshal([]byte(userJSON), &user)
// 	if err != nil {
// 		return err
// 	}

// 	*u = append(*u, user)

// 	err = writeToFile(*u)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (u *Users) List() ([]byte, error) {

// 	if len(*u) == 0 {
// 		return nil, nil
// 	}

// 	bytes, err := json.Marshal(*u)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return bytes, nil

// }

// func (u *Users) FindById(id string) ([]byte, error) {

// 	for _, user := range *u {
// 		if user.Id == id {
// 			bytes, err := json.Marshal(&user)
// 			if err != nil {
// 				return nil, err
// 			}
// 			return bytes, nil
// 		}
// 	}
// 	return nil, nil
// }

// func (u *Users) Remove(id string) error {

// 	bytes, err := u.FindById(id)
// 	if err != nil {
// 		return err
// 	}
// 	if bytes == nil {
// 		// return errors.New(fmt.Sprintf("Item with id %s not found", id))
// 		return fmt.Errorf(fmt.Sprintf("Item with id %s not found", id))
// 	}
// 	for i, user := range *u {
// 		if user.Id == id {
// 			*u = append((*u)[:i], (*u)[i+1:]...)
// 			return nil
// 		}
// 	}
// 	return nil
// }

// func parseArgs() Arguments {

// 	// var flagFilename, flagOperation, flagItem, flagId string
// 	// flag.StringVar(&flagId, "id", "", "identification")
// 	// flag.StringVar(&flagFilename, "filename", fileJSON, "file name")
// 	// flag.StringVar(&flagOperation, "operation", "list", "operacions type")
// 	// flag.StringVar(&flagItem, "item", "", "item")
// 	// flag.Parse()

// 	// return Arguments{
// 	// 	"id":        flagId,
// 	// 	"operation": flagOperation,
// 	// 	"item":      flagItem,
// 	// 	"filename":  flagFilename,
// 	// }
// 	flagId := flag.String("id", "", "User Id")
// 	flagOperation := flag.String("operation", "", "Operation")
// 	flagItem := flag.String("item", "", "Item")
// 	flagFileName := flag.String("fileName", "", "Path to file")

// 	flag.Parse()

// 	return Arguments{
// 		"id":        *flagId,
// 		"operation": *flagOperation,
// 		"item":      *flagItem,
// 		"fileName":  *flagFileName,
// 	}
// }

package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	argumentId        = "id"
	argumentOperation = "operation"
	argumentItem      = "item"
	argumentFileName  = "fileName"
)

const (
	opearionAdd      = "add"
	operationRemove  = "remove"
	operationList    = "list"
	opearionFindById = "findById"
)

type Arguments map[string]string

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type Users []User

func main() {

	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}

}

func Perform(args Arguments, writer io.Writer) error {

	if args[argumentOperation] == "" {
		return errors.New("-operation flag has to be specified")
	}

	if args[argumentFileName] == "" {
		return errors.New("-fileName flag has to be specified")
	}

	users, err := NewUsers(args[argumentFileName])
	if err != nil {
		return err
	}

	switch args[argumentOperation] {
	case operationList:

		bytes, err := users.List()
		if err != nil {
			return err
		}

		if bytes != nil {
			writer.Write(bytes)
		}

	case opearionFindById:

		if args[argumentId] == "" {
			return errors.New("-id flag has to be specified")
		}

		bytes, err := users.FindById(args[argumentId])
		if err != nil {
			return err
		}

		if bytes != nil {
			writer.Write(bytes)
		}

	case opearionAdd:

		if args[argumentItem] == "" {
			return errors.New("-item flag has to be specified")
		}

		err := users.Add(args[argumentItem])
		if err != nil {
			writer.Write([]byte(err.Error()))
			return nil
		}

		err = users.Save(args[argumentFileName])
		if err != nil {
			return err
		}

	case operationRemove:

		if args[argumentId] == "" {
			return errors.New("-id flag has to be specified")
		}

		err := users.Remove(args[argumentId])
		if err != nil {
			writer.Write([]byte(err.Error()))
			return nil
		}

		err = users.Save(args[argumentFileName])
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf(fmt.Sprintf("Operation %s not allowed!", args[argumentOperation]))

	}

	return nil

}

func NewUser(userJson string) (User, error) {

	user := User{}

	if userJson != "" {

		err := json.Unmarshal([]byte(userJson), &user)
		if err != nil {
			return user, err
		}

	}

	return user, nil

}

func NewUsers(fileName string) (Users, error) {

	users := Users{}

	if fileName != "" {
		err := users.Load(fileName)
		if err != nil {
			return nil, err
		}
	}

	return users, nil

}

func (u *Users) Load(fileName string) error {

	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	content, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if len(content) > 0 {
		err = json.Unmarshal(content, u)
		if err != nil {
			log.Fatal(err)
		}
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

			bytes, err := json.Marshal(&user)
			if err != nil {
				return nil, err
			}

			return bytes, nil

		}

	}

	return nil, nil

}

func (u *Users) Add(userJson string) error {

	user, err := NewUser(userJson)
	if err != nil {
		return err
	}

	bytes, err := u.FindById(user.Id)
	if err != nil {
		return err
	}

	if bytes != nil {
		return fmt.Errorf(fmt.Sprintf("Item with id %s already exists", user.Id))
	}

	*u = append(*u, user)

	return nil

}

func (u *Users) Remove(id string) error {

	bytes, err := u.FindById(id)
	if err != nil {
		return err
	}

	if bytes == nil {
		return fmt.Errorf(fmt.Sprintf("Item with id %s not found", id))
	}

	for i, user := range *u {

		if user.Id == id {

			*u = append((*u)[:i], (*u)[i+1:]...)

			return nil

		}

	}

	return nil

}

func (u *Users) Save(fileName string) error {

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	bytes, err := json.Marshal(*u)
	if err != nil {
		return err
	}

	file.Write(bytes)

	return nil

}

func parseArgs() Arguments {

	flagId := flag.String(argumentId, "", "User Id")
	flagOperation := flag.String(argumentOperation, "", "Operation")
	flagItem := flag.String(argumentItem, "", "Item")
	flagFileName := flag.String(argumentFileName, "", "Path to file")

	flag.Parse()

	return Arguments{
		argumentId:        *flagId,
		argumentOperation: *flagOperation,
		argumentItem:      *flagItem,
		argumentFileName:  *flagFileName,
	}

}
