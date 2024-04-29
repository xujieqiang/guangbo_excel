package methods

import (
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
	iid, err := strconv.Atoi(id)
	if err != nil {
		c.Redirect(301, "/")
		return
	}
	// 获取tasklist
	var tl []models.Tasklist
	data.Where("FanganId=?", iid).Find(&tl)

	c.HTML(200, "tasklist/index.html", gin.H{"tt": arr[0], "tasklist": tl})
}
