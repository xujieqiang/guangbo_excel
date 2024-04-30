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
type Grr struct {
	Xiaozu []string `form:"gsp[]"`
}

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
	var fa models.Fangan
	data.Where("id=?", iid).Find(&fa)
	// 获取tasklist
	var tl []models.Tasklist
	data.Where("FanganId=?", iid).Order("starttime").Find(&tl)

	c.HTML(200, "tasklist/index.html", gin.H{"tt": arr[0], "tasklist": tl, "ww": id, "msg": msg, "fanganname": fa.Listname})
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

	var gp []models.Groups
	data.Find(&gp)
	//data.Find(&al)
	c.HTML(200, "tasklist/createpage.html", gin.H{"tt": arr[0], "msg": msg, "id": id, "alllings": al, "groups": gp})
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
func (t *Tlist) Subcreate(c *gin.Context) {
	fanganid := c.PostForm("fanganid")
	listname := c.PostForm("listname")
	tag := c.PostForm("leibie")
	lingstemp := c.PostForm("lingstemp") //获取的是alltask的id
	starttime := c.PostForm("starttime")
	jobdata := c.PostForm("jobdata")
	repeatnum := c.PostForm("repeatnum")
	playmode := c.PostForm("playmode")
	var g Grr
	c.ShouldBind(&g)

	gup := ""
	for _, vv := range g.Xiaozu {
		gup += vv
	}
	fmt.Println("========>", gup)
	var tl models.Tasklist
	fid, _ := strconv.Atoi(fanganid)
	tl.Fanganid = fid
	if listname == "" {
		msg := "listname 是空的，请检查！"
		c.Redirect(302, "/tasklist/createerr/"+fanganid+"/"+msg)
		return
	}
	tl.Name = listname
	ts, _ := strconv.Atoi(tag)
	fmt.Println(ts)
	// 从alltask 模板中获取相关数据
	var lt models.Alltask
	lings, _ := strconv.Atoi(lingstemp)
	data.Where("id=?", lings).Find(&lt)
	tl.Taskid = lt.Id
	tl.Jobtype = lt.Jobtype
	tl.Jobmask = lt.Jobmask
	tl.Medias = lt.Medias
	var year models.Setting
	data.Where("id=?", 1).Find(&year)
	y := year.Dayd
	tl.Starttime = y + " " + starttime
	tl.Stoptime = y
	jd, _ := strconv.Atoi(jobdata)
	fmt.Println("============<<<<>>>>>>", jd)
	if jd != 65663 {
		msg := "jobdata 错误，请检查！"
		c.Redirect(302, "/tasklist/createerr/"+fanganid+"/"+msg)
		return
	}
	tl.Jobdata = jd
	rp, _ := strconv.Atoi(repeatnum)
	if rp >= 3 {
		msg := "repeatnum重复次数是否过多，请检查！"
		c.Redirect(302, "/tasklist/createerr/"+fanganid+"/"+msg)
		return
	}
	tl.Repeatnum = rp
	pm, _ := strconv.Atoi(playmode)
	if pm != 0 {
		msg := "playmod设置错误，可能被禁用，请检查！"
		c.Redirect(302, "/tasklist/createerr/"+fanganid+"/"+msg)
		return
	}
	tl.Playmode = pm
	tl.Groups = gup
	fmt.Println(tl)
	data.Create(&tl)
	c.Redirect(302, "/tasklist/"+fanganid)

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
