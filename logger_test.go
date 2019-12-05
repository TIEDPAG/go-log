package log

import "testing"

func TestLog(t *testing.T){
	Debug("test,%v",1)

	log:= NewPrefixLogger("service")
	log.Info("test prefix")
}
