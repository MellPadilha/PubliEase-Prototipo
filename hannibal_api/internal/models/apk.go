package models

import "github.com/ccesarfp/hannibal/internal/config"

type Installer struct{}

func (i *Installer) InstallApk(apk string) error {
	cmd := "adb install" + apk + ".apk"
	c := config.New()

	err := c.Docker.ExecuteCommand(cmd)
	if err != nil {
		return err
	}

	return nil
}
