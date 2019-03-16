package main

import (
	"github.com/gin-gonic/gin"
)

// Edit func
func (conn *Conn) Edit(c *gin.Context) {
	menuID := c.PostForm("MenuID")
	menuName := c.PostForm("MenuName")
	menuType := c.PostForm("MenuType")
	menuPrice := c.PostForm("MenuPrice")

	_, err := conn.db.Query(
		"update menu set menu_name=$1, menu_type=$2, menu_price=$3 where menu_id=$4",
		menuName, menuType, menuPrice, menuID,
	)
	if err != nil {
		c.Error(err)
		c.JSON(200, gin.H{
			"status": "error",
			"msg":    "แก้ไขรายการอาหารไม่สำเร็จ",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "success",
		"msg":    "แก้ไขรายการอาหารสำเร็จ",
	})
}
