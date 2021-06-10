package main

import (
	"github.com/KathurimaKimathi/hapikey"
	"github.com/sirupsen/logrus"
)

func main() {
	err := hapikey.CreateCRMContact()
	if err != nil {
		logrus.Print(err)
	}
}
