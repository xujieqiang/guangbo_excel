package methods

import (
	"encoding/json"
	"fmt"
	"gexcel/db"
	"gexcel/models"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type Tlist struct{}
type Grr struct {
	Xiaozu []string `form:"gsp[]"`
}
type Newterms struct {
	Zhongduan []string `form:"terms[]"`
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
	for i, val := range tl {
		s := val.Starttime
		sarr := strings.Split(s, " ")

		tl[i].Starttime = sarr[1]
		music := val.Medias
		_, err := os.Stat(music)
		if err != nil {
			tl[i].Medias = "==>>文件路径错误，无法读取！"
		} else {
			marr := strings.Split(music, "\\")
			tl[i].Medias = marr[len(marr)-1]
		}
		sg := val.Groups
		sg_arr := strings.Split(sg, ";")
		sg_arr = sg_arr[:len(sg_arr)-1]
		gname := ""

		for _, v := range sg_arr {
			var ggr models.Groups
			ng := v + ";"
			data.Where("val=?", ng).Find(&ggr)
			gname += ggr.Groupname
			gname += "; "
		}
		tl[i].Groups = gname

	}

	c.HTML(200, "tasklist/index.html", gin.H{"tt": arr[0], "tasklist": tl, "ww": id, "msg": msg, "fanganname": fa.Listname})
}

// func (t *Tlist) Modlist(c *gin.Context) {
// 	dd := time.Now().String()
// 	arr := strings.Split(dd, " ")
// 	id, _ := c.Params.Get("id")
// 	listid, _ := c.Params.Get("listid")
// 	var tl models.Tasklist
// 	lid, _ := strconv.Atoi(listid)
// 	data.Where("id=?", lid).Find(&tl)

// 	c.HTML(200, "tasklist/modpage.html", gin.H{"tt": arr[0], "id": id, "onelist": tl})
// }

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

	var allgp, gp, building_one, building_two, building_three, building_others []models.Groups
	data.Where("ty=?", 0).Find(&gp)
	data.Find(&allgp)
	data.Where("ty=?", 1).Where("groupname LIKE ?", "%一号楼%").Order("groupname").Find(&building_one)
	data.Where("ty=?", 1).Where("groupname LIKE ?", "%二号楼%").Order("groupname").Find(&building_two)
	data.Where("ty=?", 1).Where("groupname LIKE ?", "%三号楼%").Order("groupname").Find(&building_three)
	var others []int
	for _, vv := range allgp {
		tag := 0
		for _, v1 := range building_one {
			if v1.Id == vv.Id {
				tag = 1
				break
			}
		}
		if tag == 0 {
			for _, v2 := range building_two {
				if v2.Id == vv.Id {
					tag = 1
					break
				}
			}
		}
		if tag == 0 {
			for _, v3 := range building_three {
				if v3.Id == vv.Id {
					tag = 1
					break
				}
			}
		}

		if tag == 0 {
			others = append(others, vv.Id)
		}

	}
	data.Where("ty=?", 1).Where("id IN ?", others).Find(&building_others)
	//data.Find(&al)
	c.HTML(200, "tasklist/createpage.html", gin.H{
		"tt":              arr[0],
		"msg":             msg,
		"id":              id,
		"alllings":        al,
		"groups":          gp,
		"building_one":    building_one,
		"building_two":    building_two,
		"building_three":  building_three,
		"building_others": building_others,
	})
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

	var tt Newterms
	c.ShouldBind(&tt)
	allterms := ""
	for _, tv := range tt.Zhongduan {
		allterms += tv
	}

	termsarr := strings.Split(allterms, ";")
	area_len := len(termsarr)
	areamasks := ""
	for i := 0; i < area_len-1; i++ {
		areamasks += ";"
	}

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
	// if jd != 65663 {
	// 	msg := "jobdata 错误，请检查！"
	// 	c.Redirect(302, "/tasklist/createerr/"+fanganid+"/"+msg)
	// 	return
	// }
	tl.Jobdata = jd
	tl.Areamasks = areamasks
	tl.Terms = allterms
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

func (t *Tlist) Copytask(c *gin.Context) {
	id, e := c.Params.Get("id")
	if !e {
		c.Redirect(302, "/")
		return
	}
	listid, e := c.Params.Get("listid")
	if !e {
		msg := "没有找到相应的记录！"
		c.Redirect(302, "tasklist/delerr/"+id+"/"+msg)
		return
	}
	lid, _ := strconv.Atoi(listid)
	//idd, _ := strconv.Atoi(id)
	var onetask, tt models.Tasklist
	data.Where("id=?", lid).Find(&tt)
	onetask.Fanganid = tt.Fanganid
	onetask.Taskid = tt.Taskid
	onetask.Name = tt.Name
	onetask.Jobtype = tt.Jobtype
	onetask.Jobmask = tt.Jobmask
	onetask.Duration = tt.Duration
	onetask.Starttime = tt.Starttime
	onetask.Stoptime = tt.Stoptime
	onetask.Jobdata = tt.Jobdata
	onetask.Repeatnum = tt.Repeatnum
	onetask.Playmode = tt.Playmode
	onetask.Playvol = tt.Playvol
	onetask.Medias = tt.Medias
	onetask.Terms = tt.Terms
	onetask.Areamasks = tt.Areamasks
	onetask.Groups = tt.Groups
	onetask.Poweraheadplay = tt.Poweraheadplay
	data.Save(&onetask)
	c.Redirect(302, "/tasklist/"+id)
}
func (t *Tlist) Modlist(c *gin.Context) {
	dd := time.Now().String()
	arr := strings.Split(dd, " ")
	msg, e := c.Params.Get("msg")
	if !e {
		msg = ""
	}
	// 查找tasklist id对应的内容
	id, _ := c.Params.Get("id")
	//idd, _ := strconv.Atoi(id)
	lid, _ := c.Params.Get("listid")
	lidd, _ := strconv.Atoi(lid) //listid
	var tl models.Tasklist
	data.Where("id=?", lidd).Find(&tl)
	g := tl.Groups
	garr := strings.Split(g, ";")
	allgroups := ""
	garr = garr[:len(garr)-1]

	for _, val := range garr {
		var gsearch models.Groups
		newone := val + ";"
		data.Where("val=?", newone).Find(&gsearch)
		allgroups = allgroups + gsearch.Groupname + ";"
	}
	var gp []models.Groups
	data.Find(&gp)

	c.HTML(200, "tasklist/modpage.html", gin.H{"tt": arr[0], "msg": msg, "onetask": tl, "allgroups": allgroups, "id": id, "groups": gp})
}

func (t *Tlist) Submod(c *gin.Context) {
	id := c.PostForm("id") //tasklist  id
	idd, _ := strconv.Atoi(id)
	var tl models.Tasklist
	data.Where("id=?", idd).Find(&tl)
	fmt.Println(tl)
	var g Grr
	c.ShouldBind(&g)

	gup := ""
	for _, vv := range g.Xiaozu {
		gup += vv
	}

	listname := c.PostForm("listname")
	starttime := c.PostForm("starttime")
	jobdata := c.PostForm("jobdata")
	jd, _ := strconv.Atoi(jobdata)
	repeatnum := c.PostForm("repeatnum")
	rp, _ := strconv.Atoi(repeatnum)
	playmode := c.PostForm("playmode")
	pm, _ := strconv.Atoi(playmode)
	medias := c.PostForm("medias")
	tl.Name = listname
	tl.Starttime = starttime
	tl.Jobdata = jd
	tl.Repeatnum = rp
	tl.Playmode = pm
	tl.Medias = medias
	//tl.Groups = gup
	iidd := tl.Fanganid
	nid := strconv.Itoa(iidd)
	data.Save(&tl)
	c.Redirect(302, "/tasklist/"+nid)

}

func (t *Tlist) ExportExcel(c *gin.Context) {
	fid, _ := c.Params.Get("faid")
	faid, _ := strconv.Atoi(fid)

	var tl []models.Tasklist
	data.Where("fanganid=?", faid).Order("starttime").Find(&tl)
	var fangan models.Fangan
	data.Where("id=?", faid).Find(&fangan)
	fname := fangan.Listname

	ef := excelize.NewFile()
	ef.NewSheet(fname)
	ef.SetCellValue(fname, "A1", "Name")
	ef.SetCellValue(fname, "B1", "JobType")
	ef.SetCellValue(fname, "C1", "JobMask")
	ef.SetCellValue(fname, "D1", "Duration")
	ef.SetCellValue(fname, "E1", "StartTime")
	ef.SetCellValue(fname, "F1", "StopTime")
	ef.SetCellValue(fname, "G1", "JobData")
	ef.SetCellValue(fname, "H1", "RepeatNum")
	ef.SetCellValue(fname, "I1", "PlayMode")
	ef.SetCellValue(fname, "J1", "PlayVol")
	ef.SetCellValue(fname, "K1", "Medias")
	ef.SetCellValue(fname, "L1", "Terms")
	ef.SetCellValue(fname, "M1", "AreaMasks")
	ef.SetCellValue(fname, "N1", "Groups")
	ef.SetCellValue(fname, "O1", "PowerAheadPlay")

	line := 2
	for _, val := range tl {
		ef.SetCellValue(fname, fmt.Sprintf("A%d", line), val.Name)
		ef.SetCellValue(fname, fmt.Sprintf("B%d", line), val.Jobtype)
		ef.SetCellValue(fname, fmt.Sprintf("C%d", line), val.Jobmask)
		ef.SetCellValue(fname, fmt.Sprintf("D%d", line), val.Duration)
		ef.SetCellValue(fname, fmt.Sprintf("E%d", line), val.Starttime)
		ef.SetCellValue(fname, fmt.Sprintf("F%d", line), val.Stoptime)
		ef.SetCellValue(fname, fmt.Sprintf("G%d", line), val.Jobdata)
		ef.SetCellValue(fname, fmt.Sprintf("H%d", line), val.Repeatnum)
		ef.SetCellValue(fname, fmt.Sprintf("I%d", line), val.Playmode)
		ef.SetCellValue(fname, fmt.Sprintf("J%d", line), val.Playvol)
		ef.SetCellValue(fname, fmt.Sprintf("K%d", line), val.Medias)
		ef.SetCellValue(fname, fmt.Sprintf("L%d", line), val.Terms)
		ef.SetCellValue(fname, fmt.Sprintf("M%d", line), val.Areamasks)
		ef.SetCellValue(fname, fmt.Sprintf("N%d", line), val.Groups)
		ef.SetCellValue(fname, fmt.Sprintf("O%d", line), val.Poweraheadplay)
		line += 1
	}
	err := ef.SaveAs("data.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	c.Redirect(302, "/tasklist/"+fid)
}

func (t *Tlist) Importdata(c *gin.Context) {
	fanganid, _ := c.Params.Get("faid")
	faid, _ := strconv.Atoi(fanganid)
	fmt.Println(faid)
}

func (t *Tlist) ChangeFormat(c *gin.Context) {
	dd := time.Now().String()
	arr := strings.Split(dd, " ")

	c.HTML(200, "tasklist/cfmat.html", gin.H{"tt": arr[0]})
}
