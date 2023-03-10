package gps

import (
	"fmt"
	"os"
	"strings"
)

func (c *CreateStructure) CreateControllerExampleFile(projectName string, entitieName string) {
	path, _ := os.Getwd()
	localFile := fmt.Sprintf("%s/%s", path, "src/controllers")
	firstCapitalCase := strings.Title(entitieName)
	entitieNameString := fmt.Sprintf("%s/%s.go", localFile, entitieName)
	file, err := os.Create(entitieNameString)
	CheckIfIndexModelWasCreated(projectName, entitieName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileText := fmt.Sprintf(`package controllers

	import (
		"%s/src/entities"
		"%s/src/models"
		"encoding/json"
		"fmt"
		"io/ioutil"
		"net/http"
		"strconv"
	
		"github.com/swaggo/swag/example/celler/httputil"
	
		"github.com/gin-gonic/gin"
	)
	
	func GetAll%s(ctx *gin.Context) {
			%s, err := models.GetAll%s()
			if err != nil {
				httputil.NewError(ctx, http.StatusBadRequest, err)
				return
			}
			ctx.JSON(http.StatusOK, %s)
	}
	
	func Get%sByID(ctx *gin.Context) {
		Id, _ := strconv.ParseUint(ctx.Param("id"), 10, 16)
		%s, err := models.Get%sById(Id)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}
		ctx.JSON(http.StatusOK, %s)
	}
	
	func Create%s(ctx *gin.Context) {
		var new entities.%s
	
		data, _ := ioutil.ReadAll(ctx.Request.Body)
		json.Unmarshal(data, &new)
	
		%sEmail, _ := models.Get%sByEmail(new.Email)
	
		if %sEmail.Email == new.Email {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "this email already exists",
			})
			return
		}
	
		err := models.Create%s(new)
		if err != nil {
			fmt.Println("Error to create a %s controller")
		}
	
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "successfully create %s",
		})
	}
	
	func Delete%s(ctx *gin.Context) {
		Id, _ := strconv.ParseUint(ctx.Param("id"), 10, 16)
		err := models.Delete%s(Id)
		if err != nil {
			httputil.NewError(ctx, http.StatusBadRequest, err)
			return
		}
	
		ctx.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "successfully deleted %s",
		})
	}
	`, projectName, projectName, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase)

	_, err = file.WriteString(fileText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = file.Sync()
	if err != nil {
		fmt.Println(err)
		return
	}
}
