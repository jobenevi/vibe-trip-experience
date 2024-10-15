package dto

import (
	"time"

	"github.com/jobenevi/vibe-trip-experience/internal/models"
)

type TestimonialDTO struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Content      string    `json:"content"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func ConvertTestimonialToDTO(testimonial *models.Testimonial) TestimonialDTO {
	return TestimonialDTO{
		ID:           testimonial.ID,
		Name:         testimonial.Name,
		Content:      testimonial.Content,
		ProfileImage: testimonial.ProfileImage,
		CreatedAt:    testimonial.CreatedAt,
		UpdatedAt:    testimonial.UpdatedAt,
	}
}
