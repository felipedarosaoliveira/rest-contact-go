package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/felipedarosaoliveira/rest-contact-go/domain"
	"github.com/gin-gonic/gin"
)

type result struct {
	ID int64 `json:"id,omitempty"`
}

func (h handler) getContacts(c *gin.Context) {
	service := h.contactService
	contacts := service.LoadAllContacts()
	c.JSON(http.StatusOK, contacts)
}

func (h handler) getContactByID(c *gin.Context) {
	ID := c.Param("id")
	fmt.Println(ID)
	service := h.contactService
	i, _ := strconv.ParseInt(ID, 10, 64)
	contact := service.FindContactByID(i)
	status := http.StatusOK
	if contact.ID == 0 && contact.Name == "" && contact.Email == "" && contact.Phone == "" {
		status = http.StatusNotFound
		c.AbortWithStatus(status)
	} else {
		c.JSON(status, contact)
	}

}

func (h handler) insertContact(c *gin.Context) {
	var contact *domain.Contact
	//Desserializa  o json através do UnMarshal se der ero retorna vazio
	if err := c.BindJSON(&contact); err != nil {
		return
	}
	id := h.contactService.Insert(contact)
	if id > -1 {
		c.JSON(http.StatusCreated, result{ID: id})
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}

}

func (h handler) updateContact(c *gin.Context) {
	ID := c.Param("id")
	fmt.Println(ID)
	i, _ := strconv.ParseInt(ID, 10, 64)
	var contact *domain.Contact
	//Desserializa  o json através do UnMarshal se der ero retorna vazio
	if err := c.BindJSON(&contact); err != nil {
		return
	}
	contact.ID = i
	if h.contactService.Update(contact) {
		c.AbortWithStatus(http.StatusAccepted)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

func (h handler) removeContact(c *gin.Context) {
	ID := c.Param("id")
	fmt.Println(ID)
	i, _ := strconv.ParseInt(ID, 10, 64)
	if h.contactService.Remove(i) {
		c.AbortWithStatus(http.StatusNoContent)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
