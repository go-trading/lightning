package core

import (
	"errors"
	"io"
	"os"
	"reflect"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func logAndReturnError(msg string) error {
	log.Error(msg)
	return errors.New(msg)
}

func sliceContainsNotNil(slice interface{}) bool {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(slice)
		for i := 0; i < s.Len(); i++ {
			if !s.Index(i).IsNil() {
				return true
			}
		}
		return false
	default:
		log.Error("sliceContainsNotNil wait []interface input")
	}
	return false
}

//https://groups.google.com/forum/#!topic/golang-nuts/PnLnr4bc9Wo
func iPow(a, b uint64) uint64 {
	var result uint64 = 1
	for 0 != b {
		if 0 != (b & 1) {
			result *= a
		}
		b >>= 1
		a *= a
	}
	return result
}

func InitLog(cliContext *cli.Context) error {
	logrus.SetLevel(logrus.TraceLevel) //TODO read level from contex
	logFileName := time.Now().UTC().Format("20060102150405.log")

	f, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	mw := io.MultiWriter(os.Stdout, f)
	logrus.SetOutput(mw)

	logrus.WithFields(logrus.Fields{
		"logFileName": logFileName,
		"GITHUB_SHA":  cliContext.App.Version,
	}).Info("File logging initialized")
	return nil
}
