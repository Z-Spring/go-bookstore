package v1

import (
	"bookstore/data"
	"bookstore/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CartInput struct {
	Id       int     `json:"id,omitempty"`
	Pid      int     `json:"pid,omitempty"`
	Uid      string  `json:"uid,omitempty"`
	Count    int     `json:"count,omitempty"`
	SumPrice float64 `json:"sum_price,omitempty"`
}

// GetCart todo: 为什么这里的id是常量？应该改成变量啊 5.20
func GetCart(c *gin.Context) {
	var cart model.Cart
	const id = 1
	if err := data.DB.First(&cart, "id=?", id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "购物车为空！"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": cart})
	// select id from carts
}
func AddCart(c *gin.Context) {
	var cartInput CartInput
	//uid := "c-" + strconv.FormatInt(time.Now().Unix(), 10)
	//fmt.Println(uid)
	// 将参数储存在cartInput结构体中
	if err := c.ShouldBindJSON(&cartInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cart := model.Cart{
		Pid:      cartInput.Pid,
		Uid:      cartInput.Uid,
		Count:    cartInput.Count,
		SumPrice: cartInput.SumPrice,
	}
	data.DB.Create(&cart)
	// 将uid插入到数据库中
	//data.DB.Model(&cart).Update("uid", uid)
	c.JSON(http.StatusOK, gin.H{"data": cart})
}
