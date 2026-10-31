package handlers

import (
	// "food-ordering-system-project/repository"
	"net/http"
	"github.com/gin-gonic/gin"
	"food-ordering-system-project/services"
)

type MenuHandler struct {
	MenuService *services.MenuService
}

// func MenuPage(c *gin.Context) {
func (h *MenuHandler) MenuPage(c *gin.Context) {

	// menus, err := repository.GetMenus()
	menus, err := h.MenuService.GetMenus()

	if err != nil {
		c.String(500, err.Error())
		return
	}




	c.HTML(http.StatusOK, "menu.html", gin.H{
		"menus": menus,
		"title":          "Menu",
	})
}

// func AddMenu(c *gin.Context) {
func (h *MenuHandler) AddMenu(c *gin.Context) {

	name := c.PostForm("name")
	description := c.PostForm("description")
	price := c.PostForm("price")
	status := c.PostForm("statuss")


	file, err := c.FormFile("image")

	if err != nil {
		c.String(500, err.Error())
		return
	}

	imagePath := "static/uploads/" + file.Filename

	c.SaveUploadedFile(file, imagePath)

		// err = repository.CreateMenu(
		err = h.MenuService.CreateMenu(
		name,
		description,
		price,
		"/"+imagePath,
		status,
	)

	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.Redirect(302, "/admin/menu")
}

// func UpdateMenu(c *gin.Context) {
func (h *MenuHandler) UpdateMenu(c *gin.Context) {

	id := c.Param("id")

	name := c.PostForm("name")
	description := c.PostForm("description")
	price := c.PostForm("price")
	status := c.PostForm("statuss")


	// err := repository.UpdateMenu(
	err := h.MenuService.UpdateMenu(
	id,
	name,
	description,
	price,
	status,
)

	if err != nil {

		c.String(500, err.Error())
		return
	}

	c.Redirect(302, "/admin/menu")
}