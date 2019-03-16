package main

import (
	"github.com/gin-gonic/gin"
)

// Search func
func (conn *Conn) Search(c *gin.Context) {
	key := c.PostForm("key")

	rows, err := conn.db.Query("select * from menu where menu_id like '%" + key + "%' or menu_name like '%" + key + "%'")
	if err != nil {
		c.Error(err)
		c.JSON(200, gin.H{"status": "error"})
		return
	}

	menu := []Menu{}
	for rows.Next() {
		m := Menu{}
		rows.Scan(&m.MenuID, &m.MenuName, &m.MenuType, &m.MenuPrice)
		menu = append(menu, m)
	}
	c.JSON(200, gin.H{"menu": menu})
}
