package main

import (
    //"fmt"
	"com.craig/testgomod/pkg/impt"
	"git.chotot.org/go-common/kit/v2/logger"
)

var log = logger.GetLogger("etcd")

func main() {
	log.Infof("testgomod")
	impt.CallCLog()
}
