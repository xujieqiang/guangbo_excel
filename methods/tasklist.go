package methods

import (
	"encoding/json"
	"fmt"
	"gexcel/db"
	"gexcel/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Tlist struct{}

func NewTlist() Tlist {
	return Tlist{}
}

var (
	data *gorm.DB
)

func init() {
	data = db.DB

}
func (t *Tlist) Index(c *gin.Context) {

	dd := time.Now().String()
	arr := strings.Split(dd, " ")
	id, _ := c.Params.Get("id")
	fmt.Println(id)

	msg, e := c.Params.Get("msg")
	if !e {
		msg = ""
	}
	iid, err := strconv.Atoi(id)
	if err != nil {
		c.Redirect(301, "/")
		return
	}
	// 获取tasklist
	var tl []models.Tasklist
	data.Where("FanganId=?", iid).Order("starttime").Find(&tl)

	c.HTML(200, "tasklist/index.html", gin.H{"tt": arr[0], "tasklist": tl, "ww": id, "msg": msg})
}

func (t *Tlist) Modlist(c *gin.Context) {
	dd := time.Now().String()
	arr := strings.Split(dd, " ")
	id, _ := c.Params.Get("id")
	listid, _ := c.Params.Get("listid")
	var tl models.Tasklist
	lid, _ := strconv.Atoi(listid)
	data.Where("id=?", lid).Find(&tl)

	c.HTML(200, "tasklist/modpage.html", gin.H{"tt": arr[0], "id": id, "onelist": tl})
}

func (t *Tlist) Createlist(c *gin.Context) {
	dd := time.Now().String()
	arr := strings.Split(dd, " ")
	msg, e := c.Params.Get("msg")
	if !e {
		msg = ""
	}
	id, _ := c.Params.Get("id")
	// 查找所有的铃声模板
	var al []models.Alltask
	//data.Find(&al)
	c.HTML(200, "tasklist/createpage.html", gin.H{"tt": arr[0], "msg": msg, "id": id, "alllings": al})
}

func (t *Tlist) Changelb(c *gin.Context) {

	tag, _ := c.Params.Get("tag")
	fmt.Println("==========>", tag)

	all := []models.Alltask{}
	data.Where("tag=?", tag).Find(&all)
	j, err := json.Marshal(all)
	if err != nil {
		c.Redirect(302, "/")
		return
	}
	c.JSON(200, string(j))
}

func (t *Tlist) Dellist(c *gin.Context) {
	id, e := c.Params.Get("id")
	if !e {
		c.Redirect(302, "/")
		return
	}

	listid, e := c.Params.Get("listid")
	if !e {
		msg := "没有找到相应的记录！"
		c.Redirect(302, "/tasklist/delerr/"+id+"/"+msg)
	}
	//idd, _ := strconv.Atoi(id)
	lid, _ := strconv.Atoi(listid)

	var tl models.Tasklist
	data.Delete(&tl, lid)
	c.Redirect(302, "/tasklist/"+id)
}
