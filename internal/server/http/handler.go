package http

import (
	"net/http"

	"github.com/felipedarosaoliveira/rest-contact-go/domain"
	"github.com/gin-gonic/gin"
)

type handler struct {
	contactService domain.ContactService
}

//NewHandler create a new Handler with routes
func NewHandler(contactService domain.ContactService) http.Handler {
	handler := &handler{
		contactService: contactService,
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(handler.recovery())
	v1 := router.Group("/v1")
	v1.GET("/contacts", handler.getContacts)
	v1.GET("/contacts/:id", handler.getContactByID)
	v1.POST("/contacts", handler.insertContact)
	v1.PUT("/contacts/:id", handler.updateContact)
	v1.DELETE("/contacts/:id", handler.removeContact)

	return router
}

func (h *handler) recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
	}
}
