package main

import (
	"io"
	"os"
	"fmt"
	"flag" 
	"io/ioutil"
)

type Arguments map[string]string

// 0000 нет разрешений
// 0700 чтение, запись и выполнение только для владельца
// 0770 чтение, запись и выполнение для владельца и группы
// 0777 чтение, запись и выполнение для владельца, группы и других
// 0111 выполнить
// 0222 написать
// 0333 написать и выполнить
// 0444 читать
// 0555 читать и выполнять
// 0666 читать и писать
// 0740 владелец может читать, писать и выполнять; группа может только читать; у других нет разрешений

const fileJSON = "users.json"
const fPermission = 0644

func Perform(args Arguments, writer io.Writer) error {

	var filename, operation,item string
	//название, значение, описание
	//varoperationFlag := flag.String("operation", "-operation", "типы операций")

	flag.StringVar(&filename, "filename", "", "файл")
	flag.StringVar(&operation, "operation", "list", "типы операций")
	flag.StringVar(&item, "item", "", "значение")

	flag.Parse()

	file, err := os.OpenFile(fileJSON, os.O_RDWR|os.O_CREATE, fPermission)
	//os.O_RDWR - чтение/запись
	//os.O_CREATE - чтение/запись
	defer os.Remove(fileJSON)//удаляем файл
	if err != nil {
		fmt.Errorf("Error")
	}
	fmt.Print(file)
	//проверяем операции

	if(filename == "list"){
		file, err := os.OpenFile(fileJSON, os.O_RDONLY, fPermission)
		//defer os.Remove(fileJSON)
		fmt.Print(file)
		fmt.Print(err)

	}
	args = Arguments{
		"id":        "",
		"operation": "",
		"item":      "",
		"fileName":  "fileName",
	}
	return nil
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
