package main

import (
	"handler"
)

func main() {
	//data := []byte("hello,world!")
	//filePath:="/Users/fengya/pro/go-learn/scryinfo/data1.json";
	//content,err:=handler.ProcessData(filePath,data)
	//if err != nil {
	//	fmt.Println("proccess data:",err)
	//}
	//fmt.Println("content:",string(content))

	sc := &handler.ThreadScheduler{}
	sc.Start()
	sc.ExecuteTask()

	//time.Sleep(time.Second * 5)

	sc1 := &handler.ThreadScheduler{}
	sc1.Start()
	sc1.ExecuteTask()
	sc1.ExecuteTask()

	//defer sc.Stop()w

}
