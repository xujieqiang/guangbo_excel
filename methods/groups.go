package methods

import (
	"fmt"
	"gexcel/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
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

func (g *GG) Exportgroup(c *gin.Context) {
	var group []models.Groups
	data.Order("ty").Find(&group)

	ef := excelize.NewFile()
	//ef.NewSheet("sheet1")
	ef.SetCellValue("sheet1", "A1", "groupName")
	ef.SetCellValue("sheet1", "B1", "val")

	line := 2
	for _, val := range group {
		ef.SetCellValue("sheet1", fmt.Sprintf("A%d", line), val.Groupname)
		ef.SetCellValue("sheet1", fmt.Sprintf("B%d", line), val.Val)

		line += 1
	}
	err := ef.SaveAs("groups.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	c.Redirect(302, "/groups/")
}
