package api

import (
	"github.com/gin-gonic/gin"
	"novaro-server/model"
)

type TagsApi struct {
}

// GetTagsList godoc
// @Summary Get list of tags
// @Description Retrieve a list of all available tags
// @Tags tags
// @Produce json
// @Success 200 {array} model.Tags
// @Failure 400
// @Router /api/tags/getTagsList [get]
func (TagsApi) GetTagsList(c *gin.Context) {
	tags, err := model.GetTagsList()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tags)
}
