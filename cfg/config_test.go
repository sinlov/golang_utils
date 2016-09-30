package cfg

import (
	"testing"
	"github.com/bmizerany/assert"
	"fmt"
)

func TestInitCfg(t *testing.T) {
	cfg := new(Cfg)
	cfg.InitCfg("config.conf")
	fmt.Printf("cfg: %v", cfg)
}

func TestRead(t *testing.T) {
	cfg := new(Cfg)
	cfg.InitCfg("config.conf")
	daemon := cfg.Read("ServerSet", "daemon")
	port := cfg.Read("ServerSet", "port")
	fmt.Printf("daemon: %v, port: %v \n", daemon, port)
	assert.Equal(t, nil, nil)
}
