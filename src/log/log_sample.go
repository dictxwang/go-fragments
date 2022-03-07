package log

import (
	"fmt"
	"log"
	"os"
)

/*
 *created by wangqiang at 2018/12/19
 */

func doLog1() {

	fmt.Println("begin doLog1 ...")
	file := "log_sample_test.log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Fatalln("fail to create test.log file!")
	}

	// 注意log.LUTC 只是对日期时间的修饰，如果未设置日期时间，LUTC相当于无效
	// log.LstdFlags = Ldate | Ltime
	logger := log.New(logFile, "[doLog1] ", log.LstdFlags | log.Lshortfile | log.LUTC)
	logger.Println("log sample doLog1")

	// 执行log.Fatal，程序将退出
	//logger.Fatal("log sample fatal")
}

func doLog2() {

	fmt.Println("begin doLog2 ...")
	file := "log_sample_test.log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Fatalln("fail to create test.log file!")
	}

	log.SetOutput(logFile)
	log.SetPrefix("[doLog2]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)

	log.Println("log sample doLog2")
}

func SampleMainLog() {
	fmt.Println("\n[log_sample]")

	doLog1()
	doLog2()
}
