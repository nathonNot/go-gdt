package log

import "github.com/sirupsen/logrus"

type Wrapper struct {
	*logrus.Entry
}

var Logger *Wrapper

func ServerLog() *Wrapper {
	return Logger
}
