package cfg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitCfg(t *testing.T) {
	cfg := new(Cfg)
	cfg.InitCfg("config.conf")
	fmt.Printf("cfg: %v", cfg)
	assert.Equal(t, nil, nil)
}

func TestRead(t *testing.T) {
	cfg := new(Cfg)
	cfg.InitCfg("config.conf")
	daemon := cfg.Read("ServerSet", "daemon")
	port := cfg.Read("ServerSet", "port")
	fmt.Printf("daemon: %v, port: %v \n", daemon, port)
	assert.Equal(t, nil, nil)
}
