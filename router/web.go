package router

import (
	"blog/app/controller"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"html/template"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//注册自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"formatTime":utils.FormatTime,
	})
	//静态文件访问路径
	r.Static("/static","./static")
	//配置视图访问路径
	r.LoadHTMLGlob("app/view/**/*")
	r.GET("/index",controller.Index)
	r.GET("/bill/index",controller.BillIndex)
	r.GET("/bill/add",controller.BillAdd)
	r.GET("/bill/edit/:id",controller.BillEdit)
	r.GET("/bill/detail/:id",controller.BillDetail)
	r.GET("/supplier/index",controller.SupplierIndex)
	r.GET("/supplier/add",controller.SupplierAdd)
	r.POST("/supplier/add",controller.SupplierCreate)
	r.GET("/supplier/detail/:id",controller.SupplierDetail)
	r.GET("/supplier/edit/:id",controller.SupplierEdit)
	r.POST("/supplier/update",controller.SupplierUpdate)
	r.GET("/supplier/delete/:id",controller.SupplierDelete)
	r.POST("/bill/add",controller.BillCreate)
	r.POST("/bill/edit",controller.BillUpdate)
	r.GET("/bill/delete/:id",controller.BillDelete)
	return r
}
