package methods

import (
	"fmt"
	"gexcel/models"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tcolgate/mp3"
)

type Alllings struct{}
type Alltask struct {
	Taskname string `json:"taskname"`
	Jobtype  int    `json:"jobtype"`
	Jobmask  int    `json:jobmask`
	Medias   string `json:"medias"`
	Tag      int    `json:"tags"`
}

func NewAlllings() Alllings {
	return Alllings{}
}

func (a *Alllings) Index(c *gin.Context) {
	dd := time.Now().String()
	arr := strings.Split(dd, " ")

	var al []models.Alltask
	data.Where("tag=?", 1).Find(&al)
	for i, val := range al {
		_, err := os.Stat(val.Medias)
		if err != nil {
			al[i].Medias = "==>>路径错误，无法读取！"
		}
	}
	var sal []models.Alltask
	data.Where("tag=?", 0).Find(&sal)
	for i, val := range sal {
		_, err := os.Stat(val.Medias)
		if err != nil {
			sal[i].Medias = "==>>路径错误，无法读取！"
		}
	}
	var kaoshi []models.Alltask
	data.Where("tag=?", 2).Find(&kaoshi)
	for i, val := range kaoshi {
		_, err := os.Stat(val.Medias)
		if err != nil {
			kaoshi[i].Medias = "==>>路径错误，无法读取！"
		}
	}

	c.HTML(200, "alltask/index.html", gin.H{
		"tt":       arr[0],
		"alllings": al,
		"spe":      sal,
		"kaoshi":   kaoshi,
	})
}

func (a *Alllings) Refreshtime(c *gin.Context) {
	var aat []models.Alltask
	data.Find(&aat)
	for _, val := range aat {
		//var onetask models.Alltask
		music_path := val.Medias
		f, err := os.Open(music_path)
		if err != nil {
			c.Redirect(302, "/alltask")
			return
		}

		defer f.Close()

		d := mp3.NewDecoder(f)
		var frame mp3.Frame
		skipped := 0
		t := 0.0
		for {

			if err := d.Decode(&frame, &skipped); err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(err)
				return
			}

			t = t + frame.Duration().Seconds()
		}
		cxtime := int(t)
		val.Cxtime = cxtime
		data.Save(&val)
	}
	c.Redirect(302, "/alltask/")

}

func (a *Alllings) Newling(c *gin.Context) {
	dd := time.Now().String()
	arr := strings.Split(dd, " ")
	c.HTML(200, "alltask/newling.html", gin.H{"tt": arr[0]})
}

func (a *Alllings) Createling(c *gin.Context) {
	var f models.Alltask

	taskname := c.PostForm("taskname")
	jobtype := c.PostForm("jobtype")
	jobmask := c.PostForm("jobmask")
	medias := c.PostForm("medias")
	tag := c.PostForm("tags")
	f.Taskname = taskname
	f.Medias = medias
	jt, err := strconv.Atoi(jobtype)
	if err != nil || (jt != 1 && jt != 2 && jt != 3 && jt != 4) {
		fmt.Println(err, "err in jt", jt)
		c.Redirect(302, "/alltask/newtask")
		return
	}
	jm, err := strconv.Atoi(jobmask)
	if err != nil || (jm != 1 && jm != 0) {
		fmt.Println(err, "err in jm", jm)
		c.Redirect(302, "/alltask/newtask")
		return
	}
	ttag, err := strconv.Atoi(tag)
	if err != nil || (ttag != 1 && ttag != 0 && ttag != 2) {
		c.Redirect(302, "/alltask/newtask")
		return
	}
	f.Jobmask = jm
	f.Jobtype = jt
	f.Tag = ttag
	fmt.Println(f)
	data.Create(&f)

	c.Redirect(302, "/alltask/")
}

func (a *Alllings) Deltask(c *gin.Context) {
	id, e := c.Params.Get("id")

	if !e {
		c.Redirect(302, "/alltask/")
		return
	}
	idd, err := strconv.Atoi(id)
	if err != nil {
		c.Redirect(302, "/alltask/")
		return
	}
	var dtask models.Alltask
	data.Delete(&dtask, idd)
	c.Redirect(302, "/alltask/")
}

func (a *Alllings) Modtask(c *gin.Context) {
	id, e := c.Params.Get("id")

	if !e {
		c.Redirect(302, "/alltask/")
		return
	}
	fmt.Println(id)

	idd, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		c.Redirect(302, "/alltask/")
		return
	}

	// 查询相关id的task

	var newtask models.Alltask
	data.Where("id=?", idd).Find(&newtask)

	dd := time.Now().String()
	arr := strings.Split(dd, " ")
	c.HTML(200, "alltask/modpage.html", gin.H{"tt": arr[0], "onetask": newtask})
}

func (a *Alllings) Submod(c *gin.Context) {
	id := c.PostForm("id")
	taskname := c.PostForm("taskname")
	jobtype := c.PostForm("jobtype")
	jobmask := c.PostForm("jobmask")
	medias := c.PostForm("medias")
	tag := c.PostForm("tags")
	idd, _ := strconv.Atoi(id)

	var al models.Alltask

	jt, err := strconv.Atoi(jobtype)
	if err != nil || (jt != 1 && jt != 2 && jt != 3 && jt != 4) {
		fmt.Println(err, "err in jt", jt)
		c.Redirect(302, "/alltask/modtask/"+id)
		return
	}
	jm, err := strconv.Atoi(jobmask)
	if err != nil || (jm != 1 && jm != 0) {
		fmt.Println(err, "err in jm", jm)
		c.Redirect(302, "/alltask/modtask/"+id)
		return
	}
	ttag, err := strconv.Atoi(tag)
	if err != nil || (ttag != 1 && ttag != 0 && ttag != 2) {
		c.Redirect(302, "/alltask/modtask/"+id)
		return
	}
	al.Id = idd
	al.Taskname = taskname
	al.Medias = medias
	al.Jobtype = jt
	al.Jobmask = jm
	al.Tag = ttag

	data.Save(&al)
	c.Redirect(302, "/alltask/")
}
