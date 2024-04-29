package routers

import (
	"gexcel/methods"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Newrouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("templates/**/*")
	r.StaticFS("/static", http.Dir("./static"))

	/////////////////////////////////////////
	tasklist := methods.NewTlist() // struck tlis{}

	fangan := methods.NewFF()

	alltask := methods.NewAlllings()

	groups := methods.NewGg()
	api := r.Group("/")
	{
		/***********************************************************
		方案表
		删除 新建  创建  跟新，并且在方案中进一步创建生成列表，
		放置在index模块下面


		************************************************************/
		// 方案表
		api.GET("/", fangan.Index)
		api.GET("/delerrmsg/:msg", fangan.Index)

		api.GET("/cfangan/", fangan.CFangan)
		api.GET("/errmsg/:msg", fangan.CFangan) // 方案创建页面

		api.POST("/cfangan/sub", fangan.Subcreate) // 提交页面的呈现
		api.GET("/compute", fangan.Compute)
		api.POST("/computers", fangan.Computers)
		api.GET("/crs/:sum", fangan.Crs)
		api.GET("/cfangan/mod/:id", fangan.Modpage)
		api.POST("/cfangan/submod", fangan.Submod)
		api.GET("/cfangan/del/:id", fangan.Delfangan)

		///删除相应的方案

		/**********************************************************
		tasklist模块部分的问题


		************************************************************/
		/// tasklist 表的操作
		api.GET("/tasklist/:id", tasklist.Index)

		/**********************************************************
		alltask模块部分的问题
		是负责各种铃声和 音频的关联

		************************************************************/
		/// alltask 多个铃声任务的操作
		api.GET("/alltask/", alltask.Index)
		api.GET("/alltask/newtask", alltask.Newling)
		api.POST("/alltask/createtask", alltask.Createling)
		api.GET("/alltask/deltask/:id", alltask.Deltask)
		api.GET("/alltask/modtask/:id", alltask.Modtask)
		api.POST("/alltask/submod", alltask.Submod)

		/**********************************************************
		groups模块部分的问题


		************************************************************/
		/// groups 设置和添加
		api.GET("/groups/", groups.Index)
		api.GET("/groups/mod/:id", groups.Modgroups)
		api.GET("/groups/del/:id", groups.Delgroups)
		api.POST("/groups/submod", groups.Submod)

	}
	return r
}
