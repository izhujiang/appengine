package endpoint

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/izhujiang/appengine/endpoint/common"
	"github.com/izhujiang/appengine/module/buz"
)

func addEntry(c *gin.Context) {
	var input buz.Entry
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := common.CurrentUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = user.ID

	savedEntry, err := buz.SaveEntry(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func getAllEntries(context *gin.Context) {
	user, err := common.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("user form token: ", user)
	entries, err := buz.FindEntries(user.ID)
	context.JSON(http.StatusOK, gin.H{"data": entries})
}
