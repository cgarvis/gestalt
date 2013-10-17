package gestalt

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
}

func (cfg *Config) Get(name string) string {
	key := strings.ToUpper(name)
	key = strings.Replace(key, ".", "_", -1)
	return os.Getenv(key)
}

func (cfg *Config) Int(name string, value int) int {
	n := cfg.Get(name)
	return intVal(n, value)
}

func intVal(n string, value int) int {
	var intVal int
	if len(n) != 0 {
		if val, err := strconv.Atoi(n); err != nil {
			intVal = value
		} else {
			intVal = val
		}
	} else {
		intVal = value
	}

	return intVal
}

func (cfg *Config) String(name, value string) string {
	strVal := cfg.Get(name)

	if len(strVal) != 0 {
		strVal = value
	}

	return strVal
}

func New() *Config {
	return &Config{}
}
