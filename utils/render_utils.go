package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Render one of HTML, JSON or XML based on the 'Accept' header of the request
// If the header doesn't specify, HTML is rendered
// the template name is present
func RenderRequest(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}
