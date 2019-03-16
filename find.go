package main

import "github.com/gin-gonic/gin"

// Find func
func (conn *Conn) Find(c *gin.Context) {
	id := c.PostForm("id")

	menu := Menu{}
	err := conn.db.QueryRow("select * from menu where menu_id = $1", id).Scan(&menu.MenuID, &menu.MenuName, &menu.MenuType, &menu.MenuPrice)
	if err != nil {
		c.Error(err)
		c.JSON(200, gin.H{"status": "error", "msg": "ไม่พบข้อมูล"})
		return
	}
	c.JSON(200, gin.H{
		"MenuID":    menu.MenuID,
		"MenuName":  menu.MenuName,
		"MenuType":  menu.MenuType,
		"MenuPrice": menu.MenuPrice.Int64,
	})
}
