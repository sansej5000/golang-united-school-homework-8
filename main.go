package main

import (
	"io"
	"os"
	"fmt"
	"flag" 
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

var operationFlag = flag.String("-operation", "-operation", "типы операций")
var itemFlag = flag.String("-item", "-item", "значение")
var fileNameFlag = flag.String("-fileName", "-fileName", "файл")


func Perform(args Arguments, writer io.Writer) error {

	file, err := os.OpenFile(fileJSON, os.O_RDWR|os.O_CREATE, fPermission)
	//os.O_RDWR - чтение/запись
	//os.O_CREATE - чтение/запись
	defer os.Remove(fileJSON)//удаляем файл
	if err != nil {
		fmt.Errorf("Error")
	}
	//fmt.Print(file)
	//проверяем операции
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
