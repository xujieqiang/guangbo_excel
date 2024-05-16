package methods

import (
	"fmt"
	"gexcel/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type GG struct {
}

func NewGg() GG {
	return GG{}
}

func (g *GG) Index(c *gin.Context) {

	dd := time.Now().String()
	arr := strings.Split(dd, " ")
	var gl []models.Groups
	data.Find(&gl)

	c.HTML(200, "groups/index.html", gin.H{"tt": arr[0], "groups": gl})
}

// func (g *GG) ListGroups(c *gin.Context) {
// 	c.HTML(200, "", "")
// }

func (g *GG) Modgroups(c *gin.Context) {
	id, e := c.Params.Get("id")
	if !e {
		c.Redirect(302, "/groups/")
		return
	}
	dd := time.Now().String()
	arr := strings.Split(dd, " ")
	idd, _ := strconv.Atoi(id)
	var ss models.Groups
	data.Where("id=?", idd).Find(&ss)
	c.HTML(200, "groups/modgroup.html", gin.H{"tt": arr[0], "onegroup": ss})

}

func (g *GG) Delgroups(c *gin.Context) {
	id, e := c.Params.Get("id")
	if !e {
		c.Redirect(302, "/groups/")
		return
	}
	idd, _ := strconv.Atoi(id)
	fmt.Println(idd)
	// var ss models.Groups
	// data.Delete(&ss, idd)
	c.Redirect(302, "/groups/")
}

func (g *GG) Submod(c *gin.Context) {
	var gg models.Groups
	id := c.PostForm("id")
	groupname := c.PostForm("groupname")
	val := c.PostForm("val")
	ty := c.PostForm("ty")
	nty, _ := strconv.Atoi(ty)
	idd, _ := strconv.Atoi(id)
	gg.Id = idd
	gg.Groupname = groupname
	gg.Val = val
	gg.Ty = nty

	data.Save(&gg)
	c.Redirect(302, "/groups/")
}

func (g *GG) Addpage(c *gin.Context) {
	dd := time.Now().String()
	arr := strings.Split(dd, " ")
	c.HTML(200, "groups/modgroup.html", gin.H{"tt": arr[0]})
}

func (g *GG) Subgroup(c *gin.Context) {
	groupname := c.PostForm("groupname")
	val := c.PostForm("val")
	ty := c.PostForm("ty")
	nty, _ := strconv.Atoi(ty)
	fmt.Println(ty)
	fmt.Println("===>", nty)
	var gr models.Groups
	gr.Groupname = groupname
	gr.Val = val
	gr.Ty = nty
	data.Create(&gr)
	c.Redirect(302, "/groups")
}
