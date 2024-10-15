package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jobenevi/vibe-trip-experience/internal/models"
	"github.com/jobenevi/vibe-trip-experience/internal/services"
	"github.com/jobenevi/vibe-trip-experience/pkg/utils"
)

func CreateTestimonial(c *gin.Context) {
	var testimonial models.Testimonial
	if err := c.ShouldBindJSON(&testimonial); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateTestimonial(&testimonial); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, testimonial)
}

func GetAllTestimonials(c *gin.Context) {
	testimonials, err := services.GetAllTestimonials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, testimonials)
}

func GetTestimonialByID(c *gin.Context) {
	id := c.Param("id")
	uintID := utils.ConvertStringToUint(id)
	testimonial, err := services.GetTestimonialByID(uint(uintID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, testimonial)
}

func UpdateTestimonialByID(c *gin.Context) {
	id := c.Param("id")
	uintID := utils.ConvertStringToUint(id)

	var testimonial models.Testimonial
	if err := c.ShouldBindJSON(&testimonial); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateTestimonialByID(uint(uintID), &testimonial); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, testimonial)
}

func DeleteTestimonialByID(c *gin.Context) {
	id := c.Param("id")
	uintID := utils.ConvertStringToUint(id)

	if err := services.DeleteTestimonialByID(uint(uintID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
