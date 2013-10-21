package gestalt

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
}

func (cfg *Config) get(name string) string {
	key := strings.ToUpper(name)
	key = strings.Replace(key, ".", "_", -1)
	return os.Getenv(key)
}

// Bool defines a bool config with specified name and default value.
// Config value can been "true", "false", "1", "0"
func (cfg *Config) Bool(name string, value bool) bool {
  n := cfg.get(name)

  switch n {
    case "true", "1":
      return true
    case "false", "0":
      return false
  }

  return value
}

// Float64 defines a float64 config with specified name and default value.
func (cfg *Config) Float64(name string, value float64) float64 {
  n := cfg.get(name)

  var floatVal float64
  var err error
  if len(n) != 0 {
    if floatVal, err = strconv.ParseFloat(n, 64); err != nil {
      floatVal = value
    }
  } else {
    floatVal = value
  }

  return floatVal
}

// Int defines an integer config with specified name and default value.
func (cfg *Config) Int(name string, value int) int {
	n := cfg.get(name)

	var intVal int
	var err error
	if len(n) != 0 {
		if intVal, err = strconv.Atoi(n); err != nil {
			intVal = value
		}
	} else {
		intVal = value
	}

	return intVal
}

// String defines a string config with specified name and default value.
func (cfg *Config) String(name, value string) string {
	strVal := cfg.get(name)

	if len(strVal) != 0 {
		strVal = value
	}

	return strVal
}

func New() *Config {
	return &Config{}
}
