package cmd

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/kyokomi/emoji/v2"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var Doctor = &cli.Command{
	Name: "doctor",
	Action: doctor,
}

func doctor(ctx *cli.Context) error {
	checkMac()
	checkAndroid()

	return nil
}

func checkMac() {
	if runtime.GOOS != "darwin" {
		return
	}
 	log.Info("Checking iOS commands are on path...")
	checkCmd("xcrun")
}

func checkAndroid() {
	log.Info("Checking Android environment variables...")
	checkEnv("JAVA_HOME")
	checkEnv("ANDROID_HOME")
	checkEnv("ANDROID_SDK_ROOT")
	log.Info("Checking Android commands are on path...")
	checkCmd("emulator")
	checkCmd("adb")
}

func checkEnv(env string) {
	_, ok := os.LookupEnv(env)
	log.Infof("%s: %s", env, getState(ok))
}

func checkCmd(cmd string) {
	_, err := exec.LookPath(cmd)
	log.Infof("%s: %s", cmd, getState(err == nil))
}

func getState(ok bool) string {
	if ok {
		return emoji.Sprint(":check_mark_button:")
	}

	return emoji.Sprint(":cross_mark:")
}
