package controllers

import (
	"fmt"

	"github.com/ionnotion/echo-api/models"
	"github.com/labstack/echo/v4"
)

func GetAll(c echo.Context) error {
	var properties []models.PropertyPromotion
	models.DB.Find(&properties)
	fmt.Println(models.DB)
	return c.JSON(200, properties)
}

func GetById(c echo.Context) error {
	var property models.PropertyPromotion
	id := c.Param("id")
	if err := models.DB.Where("id= ?", id).First(&property).Error; err != nil{
		return echo.NewHTTPError(404, "Data Not Found")
	}

	return c.JSON(200, property)
}

func PostPromotion(c echo.Context) error {
	var input models.PropertyPromotionForm 
	
	if err := c.Bind(&input); err!= nil{
		return c.JSON(400,"Bad Request")
	}

	newPromotion := models.PropertyPromotion{Title: input.Title, Description: input.Description, Image_Banner: input.Image_Banner, Start_Time: input.Start_Time, End_Time: input.End_Time}
	models.DB.Create(&newPromotion)

	return c.JSON(201, newPromotion)
}

func UpdatePromotion(c echo.Context) error {
	id := c.Param("id")
	var input models.PropertyPromotion

	if err := models.DB.Where("id= ?", id).First(&input).Error; err != nil{
		return echo.NewHTTPError(404, "Data Not Found")
	}

	err := c.Bind(&input); if err!= nil{
		return echo.NewHTTPError(400, "Bad Request")
	}

	models.DB.Model(&input).Updates(models.PropertyPromotion{Title: input.Title, Description: input.Description, Image_Banner: input.Image_Banner, Start_Time: input.Start_Time, End_Time: input.End_Time})
	return c.JSON(200, input)
}

func DeletePromotion(c echo.Context) error {
	var property models.PropertyPromotion

	id := c.Param("id")
	if err := models.DB.Where("id= ?", id).First(&property).Error; err != nil{
		return echo.NewHTTPError(404, "Data Not Found")
	}

	models.DB.Delete(&property)

	message := fmt.Sprintf("Promotion with ID %s Deleted.",id)
	return c.JSON(200,message)
}