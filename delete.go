package main

import (
	"github.com/gin-gonic/gin"
)

// Delete func
func (conn *Conn) Delete(c *gin.Context) {
	id := c.Param("id")

	_, err := conn.db.Query("delete from menu where menu_id = $1", id)
	if err != nil {
		c.Error(err)
		c.JSON(200, gin.H{"status": "error", "msg": "ลบรายการอาหารไม่สำเร็จ"})
		return
	}
	c.JSON(200, gin.H{"status": "success", "msg": "ลบรายการอาหารสำเร็จ"})

}
