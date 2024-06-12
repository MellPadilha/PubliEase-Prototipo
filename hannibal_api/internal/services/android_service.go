package services

import (
	"github.com/ccesarfp/hannibal/internal/config/docker"
)

// InstallApk install apk in android emulator
func InstallApk(apk string) (bool, error) {
	// setting command
	cmd := "adb install " + apk + ".apk"

	_, err := executeCommand(cmd)
	if err != nil {
		return false, err
	}

	return true, nil
}

// GetPermissions get all package permissions in android emulator
func GetPermissions(apkPackage string) (string, error) {
	// setting command
	cmd := "adb shell dumpsys package " +
		apkPackage +
		" | awk '/requested permissions:/{flag=1;next}/install permissions:/{flag=0}flag' " +
		" | sed 's/^ *//' | paste -sd ',' - | sed 's/ *, */,/g'"

	output, err := executeCommand(cmd)
	if err != nil {
		return "", err
	}

	return output, nil
}

// CleanApkData clean all package data
func CleanApkData(apkPackage string) (bool, error) {
	// setting command
	cmd := "adb shell pm clear " + apkPackage

	_, err := executeCommand(cmd)
	if err != nil {
		return false, err
	}

	return true, nil
}

// UninstallApk uninstall package from android emulator
func UninstallApk(apkPackage string) (bool, error) {
	// setting command
	cmd := "adb uninstall " + apkPackage

	_, err := executeCommand(cmd)
	if err != nil {
		return false, err
	}

	return true, nil
}

// executeCommand execute command
func executeCommand(cmd string) (string, error) {
	// getting docker instance
	d, err := docker.New()
	if err != nil {
		return "", err
	}

	// executing command
	output, err := d.Execute(cmd)
	if err != nil {
		return "", err
	}

	return output, nil
}
