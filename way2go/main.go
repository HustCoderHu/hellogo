package main

import (
	"time"
	"way2go/ch14"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		// DisableColors:   false,
		ForceQuote:      false,
		TimestampFormat: "2006-01-02 15:04:05.3",
		FullTimestamp:   true,
		ForceColors:     true,
	})
	logrus.SetLevel(logrus.TraceLevel)

	// logrus.Trace("trace msg")
	// logrus.Debug("debug msg")
	// logrus.Info("info msg")
	// logrus.Warn("warn msg")
	// logrus.Error("error msg")
	// logrus.Fatal("fatal msg")
	// logrus.Panic("panic msg")
}

func test() {
	logrus.Infof("%v", 1)
	logrus.Info(2)
	out := make(chan int, 1)
	go func(in chan int) {
		logrus.Info(<-in)
		logrus.Info(<-in)
		// fmt.Println(<-in)
	}(out)
	out <- 1
	out <- 2
	time.Sleep(1e9)
}

func main() {
	// ch11.F1_interface()
	// ch11.F2_poly()
	// ch11.F3_valuable()
	// ch11.F4_type_interface()
	// ch11.F7_sort()
	// ch11.F9_case_type()
	// ch11.F9_interface_array()
	// ch11.F12_overload()
	// ch12.F1_reader()
	// ch12.F2_file_rw()
	// ch12.F5_read_write_file1()
	// ch12.F11_os_args()
	// ch14.F1_main()
	// test()
	// ch14.F4_select()
	// test()
	// ch14.F4_select()
	// ch14.F5_ticker()
	ch14.F6_recover()
}
