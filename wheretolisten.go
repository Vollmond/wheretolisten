package wheretolisten

import "gopkg.in/gin-gonic/gin.v1"
import "net/http"

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "fick you!")
	})
	router.Run()
}
