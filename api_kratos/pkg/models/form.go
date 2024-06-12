package models

import (
	"mime/multipart"
)

type Form struct {
	Name                  string `json:"name"`
	Package               string `json:"package"`
	HasLocationAccess     bool   `json:"hasLocationAccess"`
	LocationJustification string `json:"locationJustification"`
	HasCameraAccess       bool   `json:"hasCameraAccess"`
	CameraJustification   string `json:"cameraJustification"`
	Apk                   multipart.File
}
