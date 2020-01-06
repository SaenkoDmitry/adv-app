package models

type AdvCreateReq struct {
	Name   string   `form:"name" binding:"required,min=1,max=1000"`
	Price  float32  `form:"price" binding:"required"`
	Photos []string `form:"photos" binding:"required,min=1,max=4"`
}
