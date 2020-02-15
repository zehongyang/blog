package controller

import (
	"blog/app/model"
	"blog/conf"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SupplierIndex(c *gin.Context)  {
	var suppliers []model.Supplier
	err := conf.DB.Find(&suppliers).Error
	if err != nil {
		log.Println(err)
		return
	}
	c.HTML(http.StatusOK,"supplier/index.html",gin.H{"suppliers":suppliers})
	return
}

func SupplierAdd(c *gin.Context)  {
	c.HTML(http.StatusOK,"supplier/add.html",nil)
	return
}

func SupplierCreate(c *gin.Context)  {
	var supplier model.Supplier
	if err := c.ShouldBind(&supplier);err != nil{
		log.Println(err)
		return
	}
	err := conf.DB.Create(&supplier).Error
	if err != nil {
		log.Println(err)
		return
	}
	c.Redirect(http.StatusFound,"/supplier/index")
}

func SupplierDetail(c *gin.Context)  {
	id := c.Param("id")
	var supplier model.Supplier
	err := conf.DB.First(&supplier, id).Error
	if err != nil {
		log.Println(err)
		return
	}
	c.HTML(http.StatusOK,"/supplier/detail.html",gin.H{"supplier":supplier})
	return
}

func SupplierEdit(c *gin.Context)  {
	id := c.Param("id")
	var supplier model.Supplier
	err := conf.DB.First(&supplier, id).Error
	if err != nil {
		log.Println(err)
		return
	}
	c.HTML(http.StatusOK,"supplier/edit.html",gin.H{"supplier":supplier})
	return
}

func SupplierUpdate(c *gin.Context)  {
	var supplier model.Supplier
	if err := c.ShouldBind(&supplier);err != nil{
		log.Println(err)
		return
	}
	if err := conf.DB.Save(&supplier).Error;err != nil{
		log.Println(err)
		return
	}
	c.Redirect(http.StatusFound,"/supplier/index")
}

func SupplierDelete(c *gin.Context)  {
	id := c.Param("id")
	err := conf.DB.Where("id = ?", id).Delete(model.Supplier{}).Error
	if err != nil {
		log.Println(err)
		return
	}
	c.Redirect(http.StatusFound,"/supplier/index")
}