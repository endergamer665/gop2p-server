package main

import ( //imports

	"net/http"

	"github.com/gin-gonic/gin"
)

// represents the devices data
type devVar struct {
	ID       int32  `json:"id"`
	Hostname string `json:"hostname"`
	IPAdress string `json:"ipadress"`
}

var devs = []devVar{}

func main() {
	router := gin.Default()

	router.GET("/devs", qurDevs)
	router.POST("/devs", regDev)
	router.Run(":8080")
}

func regDev(c *gin.Context) {
	var newDev devVar

	if err := c.BindJSON(&newDev); err != nil {
		return
	}

	devs = append(devs, newDev)
	c.IndentedJSON(http.StatusCreated, newDev)
}

// qurDevs responds with contence of json file
func qurDevs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, devs)
}
