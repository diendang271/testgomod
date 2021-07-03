package utils

import "git.chotot.org/go-common/kit/v2/logger"

var log = logger.GetLogger("etcd")

func CLog(s string) {
	log.Infof(s)
	log.Errorf(s)
	log.Debugf(s)
}
