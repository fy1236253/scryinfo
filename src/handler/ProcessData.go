package handler

import (
	"fmt"
	"io/ioutil"
)

func ProcessData(path string, data []byte) ([]byte,error) {
	content,err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("read file filed:",err)
		return nil, err
	}
	fmt.Println("read content:",string(content))
	err = ioutil.WriteFile(path,data,0666)
	if err != nil {
		fmt.Println("write file filed",err)
	}
	fmt.Println("wirte success:",string(data))
	return content, err
}
