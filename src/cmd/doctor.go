package cmd

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/kyokomi/emoji/v2"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const (
	checkingIosCommandsMsg     = "Checking iOS commands are on path..."
	checkingAndroidEnvsMsg     = "Checking Android environment variables..."
	checkingAndroidCommandsMsg = "Checking Android commands are on path..."

	iosCommandXcrun = "xcrun"

	androidEnvJavaHome       = "JAVA_HOME"
	androidEnvAndroidHome    = "ANDROID_HOME"
	androidEnvAndroidSdkRoot = "ANDROID_SDK_ROOT"
	androidCommandEmulator   = "emulator"
	androidCommandAdb        = "adb"
)

var Doctor = &cli.Command{
	Name:   "doctor",
	Action: doctor,
}

func doctor(ctx *cli.Context) error {
	checkiOS(runtime.GOOS == "darwin")
	checkAndroid()

	return nil
}

func checkiOS(isMac bool) {
	if !isMac {
		return
	}
	log.Info(checkingIosCommandsMsg)
	checkCmd(iosCommandXcrun)
}

func checkAndroid() {
	log.Info(checkingAndroidEnvsMsg)
	checkEnv(androidEnvJavaHome)
	checkEnv(androidEnvAndroidHome)
	checkEnv(androidEnvAndroidSdkRoot)
	log.Info(checkingAndroidCommandsMsg)
	checkCmd(androidCommandEmulator)
	checkCmd(androidCommandAdb)
}

func checkEnv(env string) {
	_, ok := os.LookupEnv(env)
	log.Infof("%s: %s", env, getResultIcon(ok))
}

func checkCmd(cmd string) {
	_, err := exec.LookPath(cmd)
	log.Infof("%s: %s", cmd, getResultIcon(err == nil))
}

func getResultIcon(ok bool) string {
	if ok {
		return emoji.Sprint(":check_mark_button:")
	}

	return emoji.Sprint(":cross_mark:")
}
