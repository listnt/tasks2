package mymodule

import (
	"os"
	"os/user"
	"testing"
)

// pwd
func TestCase1(t *testing.T) {
	pwd, _ := os.Getwd()
	us, _ := user.Current()
	name, _ := os.Hostname()
	app := MyTerminalNew(pwd, us.Name, name)
	st, _ := app.Pwd("", os.Stdout, os.Stdin)
	if st != pwd {
		t.Error("expected", pwd, "got", st)
	}
}

// Cd
func TestCase2(t *testing.T) {
	pwd, _ := os.Getwd()
	us, _ := user.Current()
	name, _ := os.Hostname()
	app := MyTerminalNew(pwd, us.Name, name)
	app.Cd("../", os.Stdout, os.Stdin)
	newDir, _ := os.Getwd()
	if app.WorkDir != newDir {
		t.Error("expected", newDir, "got", app.WorkDir)
	}
}

// echo
func TestCase3(t *testing.T) {
	pwd, _ := os.Getwd()
	us, _ := user.Current()
	name, _ := os.Hostname()
	app := MyTerminalNew(pwd, us.Name, name)
	_, err := app.Echo("sd", os.Stdout, os.Stdin)
	if err != nil {
		t.Error("something broke")
	}
}

// ps
func TestCase4(t *testing.T) {
	pwd, _ := os.Getwd()
	us, _ := user.Current()
	name, _ := os.Hostname()
	app := MyTerminalNew(pwd, us.Name, name)
	_, err := app.Ps("", os.Stdout, os.Stdin)
	if err != nil {
		t.Error("something broke")
	}
}
