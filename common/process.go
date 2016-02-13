package common

import (
	log "github.com/Sirupsen/logrus"
	conf "github.com/andboson/configlog"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

func ChangeOldAppPort() {
	existsPid := ReadPid()
	if existsPid != 0 {
		Log.Printf("Exists Pid %d", existsPid)
		process, _ := os.FindProcess(existsPid)
		process.Signal(syscall.SIGTERM)
	}
}

func ReplacePid() {
	pid := os.Getpid()
	WritePid(pid)
	Log.Printf("Started with pid: %d", pid)
}

//pid writer
func WritePid(pid int) {
	var pidfile string
	pidfile, pidnameerr := conf.AppConfig.String("pidfile")
	if pidnameerr != nil {
		pidfile = conf.CurrDirectory + string(filepath.Separator) + "pidfile"
	}
	f, err := os.OpenFile(pidfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v file:", err, pidfile)
	}
	f.WriteString(strconv.Itoa(pid))
	f.Close()
}

//pid reader
func ReadPid() int {
	var pidfile string
	pidfile, pidnameerr := conf.AppConfig.String("pidfile")
	if pidnameerr != nil {
		pidfile = conf.CurrDirectory + string(filepath.Separator) + "pidfile"
	}
	text, _ := ioutil.ReadFile(pidfile)
	pid, _ := strconv.Atoi(string(text))
	return pid
}

//clear pid
func ClearPid() {
	var pidfile string
	pidfile, pidnameerr := conf.AppConfig.String("pidfile")
	if pidnameerr != nil {
		pidfile = conf.CurrDirectory + string(filepath.Separator) + "pidfile"
	}
	error := os.Remove(pidfile)
	if error != nil {
		Log.Printf("error clear pid")
	}
}
