package ostest

import (
	"fmt"
	"os"
	"reflect"
)

type User struct {
	Id   int
	Name string
	//addr string
}

type FileUtil struct {
	IMAGE_DIC      string
	LIMITED_LENGTH int
}

// var fileUtilInstance = &fileUtil{}

func FileUtilInstance(path string) *FileUtil {
	return &FileUtil{path, 0}
}

func reflcts(path string) {
	u := User{Id: 1001, Name: "xxx" /*, addr:"xxx"*/}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	for k := 0; k < t.NumField(); k++ {
		fmt.Printf("%s -- %v \n", t.Field(k).Name, v.Field(k).Interface())
	}

	fileUtil := FileUtilInstance(path)
	fileType := reflect.TypeOf(*fileUtil)
	fileValue := reflect.ValueOf(*fileUtil)
	for k := 0; k < fileType.NumField(); k++ {
		if fileType.Field(k).Name != "LIMITED_LENGTH" {
			fmt.Printf("%s -- %v \n", fileType.Field(k).Name, fileValue.Field(k).Interface())
		}
	}
}

func createDir(name string) error {
	err := os.Mkdir(name, 0777)
	if err != nil {
		return err
	}
	return nil
}

func createAllDir(path interface{}) error {
	err := os.MkdirAll(path.(string), 0777)
	if err != nil {
		return err
	}
	return nil
}
