package gestalt

import "os"
import "testing"

func TestBool(t *testing.T) {
    os.Clearenv()
    config := New()

    boolCfg := config.Bool("bool", false)
    if boolCfg != false {
        t.Errorf("expected %v, got %v", false, boolCfg)
    }

    os.Setenv("BOOL", "true")

    boolCfg = config.Bool("bool", false)
    if boolCfg != true {
        t.Errorf("expected %v, got %v", true, boolCfg)
    }
}

func TestFloat64(t *testing.T) {
    os.Clearenv()
    config := New()

    float64Cfg := config.Float64("float64", 14.15)
    if float64Cfg != 14.15 {
        t.Errorf("expected %v, got %v", 14.15, float64Cfg)
    }

    os.Setenv("FLOAT64", "16.17")

    float64Cfg = config.Float64("float64", 14.15)
    if float64Cfg != 16.17 {
        t.Errorf("expected %v, got %v", 16.17, float64Cfg)
    }
}

func TestInt(t *testing.T) {
    os.Clearenv()
    config := New()

    intCfg := config.Int("int", 3)
    if intCfg != 3 {
        t.Errorf("expected %v, got %v", 3, intCfg)
    }

    os.Setenv("INT", "5")

    intCfg = config.Int("int", 3)
    if intCfg != 5 {
        t.Errorf("expected %v, got %v", 5, intCfg)
    }
}

func TestString(t *testing.T) {
    os.Clearenv()
    config := New()

    stringCfg := config.String("string", "foobar")
    if stringCfg != "foobar" {
        t.Errorf("expected %v, got %v", "foobar", stringCfg)
    }

    os.Setenv("STRING", "barfoo")

    stringCfg = config.String("string", "foobar")
    if stringCfg != "barfoo" {
        t.Errorf("expected %v, got %v", "barfoo", stringCfg)
    }
}
