package controllers

import (
	"bytes"
	"encoding/gob"
	"github.com/ccesarfp/hannibal/internal/config/cache"
	"github.com/ccesarfp/hannibal/internal/services"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
)

// FormController handles HTTP requests related to forms.
type FormController struct{}

// NewFormController creates and returns a new instance of FormController.
func NewFormController() *FormController {
	return &FormController{}
}

// ValidateApplication handles GET requests to retrieve forms.
func (c *FormController) ValidateApplication(ctx *gin.Context) {

	name := ctx.PostForm("name")
	packageName := ctx.PostForm("package")

	// verifying if exists cache
	ch := cache.New()
	ch.Name = packageName
	cacheExists := ch.VerifyFile()
	if cacheExists == true {
		err := ch.GetCacheFile()
		if err != nil {
			ctx.String(http.StatusBadRequest, "Error: "+err.Error())
			return
		}

		var permissionsList []string
		buf := bytes.NewBuffer(ch.Content)
		err = gob.NewDecoder(buf).Decode(&permissionsList)
		if err != nil {
			ctx.String(http.StatusBadRequest, "Error: "+err.Error())
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"permissions": permissionsList,
		})
		return
	}

	// setting file
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	src, err := file.Open()
	if err != nil {
		ctx.String(http.StatusBadRequest, "Error: "+err.Error())
		return
	}
	defer src.Close()

	// saving file
	fileCreate, err := os.Create("./sdk/" + name + ".apk")
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Error: "+err.Error())
		return
	}
	defer fileCreate.Close()
	io.Copy(fileCreate, src)

	// installing apk
	_, err = services.InstallApk(name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to install apk. Error: " + err.Error(),
		})
		return
	}

	// verifying permissions
	permissions, err := services.GetPermissions(packageName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to capture permissions: " + err.Error(),
		})
		return
	}

	permissionsList := strings.Split(permissions, ",")

	// creating cache
	if permissions != "" && permissionsList != nil {
		buf := &bytes.Buffer{}
		err = gob.NewEncoder(buf).Encode(permissionsList)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to encode cache: " + err.Error(),
			})
		}
		ch.Content = buf.Bytes()

		err = ch.CreateCacheFile()
		if err != nil {
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "failed to create cache: " + err.Error(),
				})
			}
		}
	}

	// uninstalling apk
	_, err = services.UninstallApk(packageName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to remove apk: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"permissions": permissionsList,
	})
}
