package methods

import (
	"fmt"
	"gexcel/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type FF struct {
}
type Xq struct {
	Val []string `form:"xingqi[]"`
}
type Father struct {
	Fa string `form:"fatherfa"`
}

func NewFF() FF {
	return FF{}
}

func (f *FF) Index(c *gin.Context) {

	dd := time.Now().String()
	arr := strings.Split(dd, " ")
	msg, e := c.Params.Get("msg")
	if !e {
		msg = ""
	}
	// 获取tasklist
	var fa []models.Fangan
	data.Order("createtime").Find(&fa)

	c.HTML(200, "index/index.html", gin.H{"tt": arr[0], "fangan": fa, "msg": msg})
}

func (f *FF) Delfangan(c *gin.Context) {

	s, _ := c.Params.Get("id")

	ns, err := strconv.Atoi(s)
	if err != nil {

		c.Redirect(301, "/")
		return
	}
	fmt.Println(ns)
	var ff models.Fangan

	//在删除之前，先判断在 tasklist 表中是否存在和方案相关的数据

	var tl []models.Tasklist
	rs := data.Where("fanganid=?", ns).Find(&tl)
	if rs.RowsAffected >= 1 {
		msg := "在生成的铃声列表里面有和方案相关的记录,请进入方案删除所有的数据，再删除方案！"
		c.Redirect(302, "/delerrmsg/"+msg)
		return
	}

	data.Delete(&ff, ns)

	//c.HTML(200, "index/fangdel.html", gin.H{"tt": ns})
	c.Redirect(302, "/")
}

func (f *FF) CFangan(c *gin.Context) {
	dd := time.Now().String()
	arr := strings.Split(dd, " ")
	msg, e := c.Params.Get("msg")
	if !e {
		msg = ""
	}
	c.HTML(200, "index/cfangan.html", gin.H{"tt": arr[0], "msg": msg})
}
func (f *FF) Subcreate(c *gin.Context) {
	// dd := time.Now().String()
	// arr := strings.Split(dd, " ")
	var father Father
	c.ShouldBind(&father)
	ff := father.Fa
	fatag, _ := strconv.Atoi(ff)

	var fan models.Fangan

	faname := c.PostForm("faname")

	discreption := c.PostForm("discreption")
	fan.Discreption = discreption
	fan.Listname = faname
	nn := time.Now().String()
	fan.Createtime = nn[:19]
	fan.Fatherfa = fatag

	var newf []models.Fangan
	rs := data.Where("listname=?", faname).Find(&newf)

	if rs.RowsAffected >= 1 {
		msg := "方案名称重复"

		//c.HTML(200, "index/errormsg.html", gin.H{"tt": arr[0], "errormsg": msg})
		c.Redirect(302, "/errmsg/"+msg)
		return
	}

	data.Create(&fan)
	c.Redirect(302, "/")
}
func (f *FF) Modpage(c *gin.Context) {
	dd := time.Now().String()
	arr := strings.Split(dd, " ")
	id, e := c.Params.Get("id")
	if !e {
		c.Redirect(302, "/")
		return
	}
	idd, _ := strconv.Atoi(id)
	var fa models.Fangan
	data.Where("id=?", idd).Find(&fa)

	c.HTML(200, "index/modpage.html", gin.H{"tt": arr[0], "onefa": fa})
}

func (f *FF) Submod(c *gin.Context) {
	var fal models.Fangan
	fa := c.PostForm("faname")
	discrep := c.PostForm("discreption")
	ctime := c.PostForm("ctime")
	id := c.PostForm("id")
	fal.Discreption = discrep
	fal.Listname = fa
	idd, _ := strconv.Atoi(id)
	fal.Id = idd
	fal.Createtime = ctime

	fa_fu := c.PostForm("fatherfa")
	ff, _ := strconv.Atoi(fa_fu)
	fal.Fatherfa = ff

	data.Save(&fal)
	c.Redirect(302, "/")
}
func (f *FF) Compute(c *gin.Context) {
	dd := time.Now().String()
	arr := strings.Split(dd, " ")

	c.HTML(200, "index/compute.html", gin.H{"tt": arr[0]})
}

func (f *FF) Computers(c *gin.Context) {

	var xx Xq
	c.ShouldBind(&xx)
	fmt.Println(xx)
	narr := xx.Val
	sum := 0
	for _, v := range narr {
		switch v {
		case "2":
			a := 0b10
			sum += a
		case "3":
			a := 0b100
			sum += a
		case "4":
			a := 0b1000
			sum += a
		case "5":
			a := 0b10000
			sum += a
		case "6":
			a := 0b100000
			sum += a
		case "7":
			a := 0b1000000
			sum += a
		default:
			a := 1
			sum += a
		}
	}
	x := 0b10000000000000000
	sum += x
	fmt.Println(sum)
	c.Redirect(302, "/crs/"+strconv.Itoa(sum))
}

func (f *FF) Crs(c *gin.Context) {
	sum, _ := c.Params.Get("sum")
	c.HTML(200, "index/crs.html", gin.H{"jj": sum})
}
