package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addMessage(c *gin.Context, message string, code int) {
	c.AbortWithError(code, errors.New(message))
	c.JSON(code, gin.H{
		"message": message,
	})

}

func getMaterial(c *gin.Context) {

	grp := c.Query("grp")

	if len(materials) == 0 {
		addMessage(c, "Not Found", http.StatusNotFound)
		return
	}

	var listMaterial []Material

	if grp != "" {
		for _, material := range materials {
			if material.Group == grp {
				listMaterial = append(listMaterial, material)
			}
		}
	} else {
		listMaterial = materials
	}

	c.JSON(http.StatusOK, listMaterial)
	return
}

func getMaterialDet(c *gin.Context) {

	id := c.Param("id")

	var respMaterial Material

	for _, material := range materials {
		if material.Code == id {
			respMaterial = material
			break
		}
	}
	if (respMaterial == Material{}) {
		addMessage(c, "Not Found", http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, respMaterial)
	return
}

func createMaterial(c *gin.Context) {

	var reqMaterial Material

	c.Bind(&reqMaterial)

	var materialExist bool = false
	for _, material := range materials {
		if material.Code == reqMaterial.Code {
			materialExist = true
			break
		}
	}
	if materialExist {
		addMessage(c, "Material exists already", http.StatusBadRequest)
		return
	}

	reqMaterial.Release = false

	materials = append(materials, reqMaterial)

	c.JSON(http.StatusCreated, reqMaterial)
	return

}

func updateMaterial(c *gin.Context) {

	id := c.Param("id")

	var reqMaterial Material

	c.Bind(&reqMaterial)

	var respMaterial Material
	for i, material := range materials {
		if material.Code == id {
			materials[i].Desc = reqMaterial.Desc
			materials[i].Uom = reqMaterial.Uom
			materials[i].Release = reqMaterial.Release
			materials[i].Group = reqMaterial.Group
			respMaterial = materials[i]
			break
		}
	}
	if (respMaterial == Material{}) {
		addMessage(c, "Material not found", http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, respMaterial)
	return

}

func deleteMaterial(c *gin.Context) {

	id := c.Param("id")

	var idx int
	var materialExist bool = false
	for i, material := range materials {
		if material.Code == id {
			materialExist = true
			idx = i
			break
		}
	}
	if !materialExist {
		addMessage(c, "Material not found", http.StatusNotFound)
		return
	}

	materials = append(materials[:idx], materials[idx+1:]...)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s deleted", id),
	})
	return

}

func releaseMaterial(c *gin.Context) {

	id := c.Param("id")

	var reqMaterial Material

	c.Bind(&reqMaterial)

	var respMaterial Material

	for i, material := range materials {
		if material.Code == id {
			materials[i].Release = reqMaterial.Release
			respMaterial = materials[i]
		}
	}

	if (respMaterial == Material{}) {
		addMessage(c, "Material not found", http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, respMaterial)
	return

}
