package testlog

import "github.com/sirupsen/logrus"

func F() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		// DisableColors:   false,
		ForceQuote:      false,
		TimestampFormat: "2006-01-02 15:04:05.3",
		FullTimestamp:   true,
		ForceColors:     true,
	})
	logrus.SetLevel(logrus.TraceLevel)

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Fatal("fatal msg")
	logrus.Panic("panic msg")
}
