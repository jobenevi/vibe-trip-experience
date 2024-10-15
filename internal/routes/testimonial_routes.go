package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jobenevi/vibe-trip-experience/internal/controllers"
)

func HandlerRequests() {
	router := gin.Default()
	router.POST("/testimonial", controllers.CreateTestimonial)
	router.GET("/testimonials", controllers.GetAllTestimonials)
	router.GET("/testimonial/:id", controllers.GetTestimonialByID)
	router.PUT("/testimonial/:id", controllers.UpdateTestimonialByID)
	router.DELETE("/testimonial/:id", controllers.DeleteTestimonialByID)
	router.Run()
}
