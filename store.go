package main

import (
	"github.com/gin-gonic/gin"
)

// Store func
func (conn *Conn) Store(c *gin.Context) {
	menuID := c.PostForm("MenuID")
	menuName := c.PostForm("MenuName")
	menuType := c.PostForm("MenuType")
	menuPrice := c.PostForm("MenuPrice")

	rows, _ := conn.db.Query("select * from menu where menu_id = $1", menuID)
	count := 0
	for rows.Next() {
		count++
	}
	if count > 0 {
		c.JSON(200, gin.H{"status": "error", "msg": "รหัสเมนูนี้มีอยู่ในระบบแล้ว"})
		return
	}

	_, err := conn.db.Exec(
		"insert into menu (menu_id, menu_name, menu_type, menu_price) values ($1,$2,$3,$4)",
		menuID, menuName, menuType, NewNullString(menuPrice),
	)
	if err != nil {
		c.Error(err)
		c.JSON(200, gin.H{"status": "error", "msg": "เพิ่มรายการอาหารไม่สำเร็จ"})
		return
	}
	c.JSON(200, gin.H{"status": "success", "msg": "เพิ่มรายการอาหารสำเร็จ"})
}
