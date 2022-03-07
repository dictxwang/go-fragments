package file

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func SampleMain()  {

	// 获取当前路径
	currentDir, _ := os.Getwd()
	fmt.Printf("currentDir:%s\n", currentDir)

	filePath := "file_sample_test.log"
	fi, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND, 0)
	if err != nil {
		fmt.Printf("Error:%s\n", err)
		return
	}

	defer fi.Close()

	bw := bufio.NewWriter(fi)
	bw.WriteString("123\n")
	bw.Flush()

	b, err := ioutil.ReadFile(filePath)
	content := string(b)
	fmt.Println(content)
}