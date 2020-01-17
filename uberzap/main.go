package main

import (
	"context"
	"fmt"
	"time"

	"logframeworks/uberzap/log"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var printCount = 0
func main() {
	start := time.Now()
	mainCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mainCtx = log.WithFields(mainCtx, zap.String("ReqID", uuid.New().String()))
	log.GetLogger(mainCtx).Sugar().Infof("main function")
	printCount++
	for i := 0; i < 10000; i++ {
		func1(mainCtx)
		func2(mainCtx)
		func1(mainCtx)
		func2(mainCtx)
	}
	log.GetLogger(mainCtx).Sugar().Infof("main function end")
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(fmt.Sprintf("elapsed: %v", elapsed.Seconds()))
	printCount++
	fmt.Println(fmt.Sprintf("printCount: %v", printCount))
}

func func1(ctx context.Context) {
	log.GetLogger(ctx).Sugar().Infof("function1 printing with parent ctx")
	printCount++
	childReqID :=  uuid.New().String()
	childCtx := log.WithFields(ctx, zap.String("childReqID", childReqID))
	log.GetLogger(childCtx).Sugar().Infof("function1 printing with child ctx")
	printCount++
}

func func2(ctx context.Context) {
	log.GetLogger(ctx).Sugar().Infof("function2 printing with parent ctx")
	printCount++
	childReqID :=  uuid.New().String()
	childCtx := log.WithFields(ctx, zap.String("childReqID", childReqID))
	log.GetLogger(childCtx).Sugar().Infof("function2 printing with child ctx")
	printCount++
}
