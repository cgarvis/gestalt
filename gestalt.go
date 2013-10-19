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

// Bool defines a bool config with specified name and default value.
// Config value can been "true", "false", "1", "0"
func (cfg *Config) Bool(name string, value bool) bool {
  n := cfg.Get(name)

  switch n {
    case "true", "1":
      return true
    case "false", "0":
      return false
  }

  return value
}

// Int defines an integer config with specificed name and default value.
func (cfg *Config) Int(name string, value int) int {
	n := cfg.Get(name)

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
