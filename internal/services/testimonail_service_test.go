package services

import (
	"testing"

	"github.com/jobenevi/vibe-trip-experience/internal/database"
	"github.com/jobenevi/vibe-trip-experience/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() {
	database.DB, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	database.DB.AutoMigrate(&models.Testimonial{})
}

func TestCreateTestimonial(t *testing.T) {
	setupTestDB()

	testimonial := &models.Testimonial{Content: "New Testimonial"}
	err := CreateTestimonial(testimonial)
	assert.NoError(t, err)

	var result models.Testimonial
	database.DB.First(&result, testimonial.ID)
	assert.Equal(t, "New Testimonial", result.Content)
}

func TestGetAllTestimonials(t *testing.T) {
	setupTestDB()

	database.DB.Create(&models.Testimonial{Content: "Testimonial 1"})
	database.DB.Create(&models.Testimonial{Content: "Testimonial 2"})

	testimonials, err := GetAllTestimonials()
	assert.NoError(t, err)
	assert.Len(t, testimonials, 2)
}

func TestGetTestimonialByID(t *testing.T) {
	setupTestDB()

	database.DB.Create(&models.Testimonial{Content: "Testimonial 1"})
	database.DB.Create(&models.Testimonial{Content: "Testimonial 2"})

	testimonial, err := GetTestimonialByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "Testimonial 1", testimonial.Content)
}

func TestUpdateTestimonial(t *testing.T) {
	setupTestDB()
	testimonial := &models.Testimonial{Content: "Testimonial 1"}
	database.DB.Create(testimonial)

	updatedTestimonial := &models.Testimonial{Content: "Updated Testimonial"}
	err := UpdateTestimonialByID(testimonial.ID, updatedTestimonial)
	assert.NoError(t, err)

	var result models.Testimonial
	database.DB.First(&result, testimonial.ID)
	assert.Equal(t, "Updated Testimonial", result.Content)
}

func TestDeleteTestimonial(t *testing.T) {
	setupTestDB()
	testimonial := &models.Testimonial{Content: "Testimonial 1"}
	database.DB.Create(testimonial)

	err := DeleteTestimonialByID(testimonial.ID)
	assert.NoError(t, err)

	var result models.Testimonial
	database.DB.First(&result, testimonial.ID)
	assert.Equal(t, models.Testimonial{}, result)
}
