package main

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"logframeworks/logrus/log"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var printCount = 0

func main() {
	start := time.Now()
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: log.RFC3339NanoFixed,
		FullTimestamp:   true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%s:%d",filepath.Base(f.File), f.Line)
		},
	})
	logrus.SetLevel(logrus.DebugLevel)
	mainCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mainCtx = context.WithValue(mainCtx, "ReqID", uuid.New().String())
	mainReqID, _ := mainCtx.Value("ReqID").(string)
	v := logrus.Fields{
		"ReqID": mainReqID,
	}
	mainCtx = log.WithLogger(mainCtx, logrus.WithFields(v))
	log.GetLogger(mainCtx).Infof("main function")
	printCount++
	for i := 0; i < 10000; i++ {
		func1(mainCtx)
		func2(mainCtx)
		func1(mainCtx)
		func2(mainCtx)
	}
	log.GetLogger(mainCtx).Infof("main function end")
	printCount++
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(fmt.Sprintf("elapsed: %v", elapsed.Seconds()))
	printCount++
	fmt.Println(fmt.Sprintf("printCount: %v", printCount))
}

func func1(ctx context.Context) {
	log.GetLogger(ctx).Infof("function1 printing with parent ctx")
	printCount++
	childReqID :=  uuid.New().String()
	logrusFields := logrus.Fields{
		"ChildReqID": childReqID,
	}
	childCtx := log.WithFields(ctx, logrusFields)
	log.GetLogger(childCtx).Infof("function1 printing with child ctx")
	printCount++
}

func func2(ctx context.Context) {
	log.GetLogger(ctx).Infof("function2 printing with parent ctx")
	printCount++
	childReqID :=  uuid.New().String()
	logrusFields := logrus.Fields{
		"ChildReqID": childReqID,
	}
	childCtx := log.WithFields(ctx, logrusFields)
	log.GetLogger(childCtx).Infof("function2 printing with child ctx")
	printCount++
}
