package services

import (
	"github.com/jobenevi/vibe-trip-experience/internal/database"
	"github.com/jobenevi/vibe-trip-experience/internal/models"
	"github.com/jobenevi/vibe-trip-experience/internal/models/dto"
)

var TestimonialService TestimonialServiceInterface

type TestimonialServiceInterface interface {
	CreateTestimonial(testimonial *models.Testimonial) error
	GetAllTestimonials() ([]dto.TestimonialDTO, error)
	GetTestimonialById(id uint) (*models.Testimonial, error)
}

func CreateTestimonial(testimominal *models.Testimonial) error {
	if err := database.DB.Create(testimominal).Error; err != nil {
		return err
	}
	return nil
}

func GetAllTestimonials() ([]dto.TestimonialDTO, error) {
	var testimonials []models.Testimonial
	if err := database.DB.Find(&testimonials).Error; err != nil {
		return nil, err
	}

	var testimonialsDTO []dto.TestimonialDTO
	for _, testimonial := range testimonials {
		testimonialsDTO = append(testimonialsDTO, dto.ConvertTestimonialToDTO(&testimonial))
	}
	return testimonialsDTO, nil
}

func GetTestimonialByID(id uint) (dto.TestimonialDTO, error) {
	var testimonial models.Testimonial
	if err := database.DB.First(&testimonial, id).Error; err != nil {
		return dto.TestimonialDTO{}, err
	}

	testimonialDTO := dto.ConvertTestimonialToDTO(&testimonial)
	return testimonialDTO, nil
}

func UpdateTestimonialByID(id uint, testimonial *models.Testimonial) error {
	if err := database.DB.Model(&models.Testimonial{}).Where("id = ?", id).Updates(testimonial).Error; err != nil {
		return err
	}
	return nil

}

func DeleteTestimonialByID(id uint) error {
	if err := database.DB.Delete(&models.Testimonial{}, id).Error; err != nil {
		return err
	}
	return nil
}
