package main

import (
	"github.com/gin-gonic/gin"
)

// Index func
func (conn *Conn) Index(c *gin.Context) {
	rows, err := conn.db.Query("select * from menu")
	if err != nil {
		c.Error(err)
	}

	menu := []Menu{}
	for rows.Next() {
		m := Menu{}
		err = rows.Scan(&m.MenuID, &m.MenuName, &m.MenuType, &m.MenuPrice)
		if err != nil {
			c.Error(err)
		}
		menu = append(menu, m)
	}

	c.JSON(200, gin.H{
		"menu": menu,
	})
}
