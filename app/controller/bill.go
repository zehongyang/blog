package controller

import (
	"blog/app/model"
	"blog/conf"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func BillIndex(c *gin.Context)  {
	var bills []model.Bill
	if err := conf.DB.Preload("Supplier").Find(&bills).Error;err != nil{
		log.Println(err)
		return
	}
	c.HTML(http.StatusOK,"bill/index.html",gin.H{"bills":bills})
	return
}

func BillAdd(c *gin.Context)  {
	var suppliers []model.Supplier
	err := conf.DB.Find(&suppliers).Error
	if err != nil {
		log.Println(err)
		return
	}
	c.HTML(http.StatusOK,"bill/add.html",gin.H{"suppliers":suppliers})
	return
}

func BillEdit(c *gin.Context)  {
	id := c.Param("id")
	var bill model.Bill
	if err := conf.DB.Preload("Supplier").First(&bill, id).Error;err != nil{
		log.Println(err)
		return
	}
	c.HTML(http.StatusOK,"bill/edit.html",gin.H{"bill":bill})
	return
}

func BillDetail(c *gin.Context)  {
	id := c.Param("id")
	var bill model.Bill
	if err := conf.DB.Preload("Supplier").First(&bill, id).Error;err != nil{
		log.Println(err)
		return
	}
	c.HTML(http.StatusOK,"bill/detail.html",gin.H{"bill":bill})
	return
}

func BillCreate(c *gin.Context)  {
	var bill model.Bill
	if err := c.ShouldBind(&bill);err != nil{
		log.Println(err)
		return
	}
	if err := conf.DB.Create(&bill).Error;err != nil{
		log.Println(err)
		return
	}
	c.Redirect(http.StatusFound,"/bill/index")
}

func BillUpdate(c *gin.Context)  {
	var bill model.Bill
	if err := c.ShouldBind(&bill);err != nil{
		log.Println(err)
		return
	}
	if err := conf.DB.Model(&bill).Select("product_name","product_unit","total_account","product_num",
		"total_account","is_pay").Updates(map[string]interface{}{"product_name": bill.ProductName,
			"product_unit": bill.ProductUnit, "product_num": bill.ProductNum, "total_account": bill.TotalAccount,
			"is_pay": bill.IsPay}).Error;err != nil{
		log.Println(err)
		return
	}
	c.Redirect(http.StatusFound,"/bill/index")
}

func BillDelete(c *gin.Context)  {
	id := c.Param("id")
	if err := conf.DB.Where("id = ?", id).Delete(model.Bill{}).Error;err != nil{
		log.Println(err)
		return
	}
	c.Redirect(http.StatusFound,"/bill/index")
}