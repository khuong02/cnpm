package api

import (
	"cnpm/pkg/model"
	"github.com/gin-gonic/gin"
)

func (r Route) getByID(c *gin.Context) {
	id := c.Param("id")
	centers, err := r.CenterSvc.GetByID(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, centers)
}

func (r Route) list(c *gin.Context) {
	var payload model.GetListCenter
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Sai định dạng dữ liệu đầu vào!",
		})
		return
	}

	centers, err := r.CenterSvc.List(c, payload)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Lấy danh sách trung tâm thất bại!",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Lấy danh sách trung tâm thành công!",
		"data":    centers,
	})
}

func (r Route) insert(c *gin.Context) {
	var payload model.UpsertCenter
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Sai định dạng dữ liệu đầu vào!",
		})
		return
	}

	center, err := r.CenterSvc.Insert(c, payload)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Thêm trung tâm thất bại!",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Thêm trung tâm thành công!",
		"data":    center,
	})
}

func (r Route) update(c *gin.Context) {
	id := c.Param("id")
	var payload model.UpsertCenter
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Sai định dạng dữ liệu đầu vào!",
		})
		return
	}

	center, err := r.CenterSvc.Update(c, id, payload)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Cập nhật trung tâm thất bại!",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Cập nhật trung tâm thành công!",
		"data":    center,
	})
}
