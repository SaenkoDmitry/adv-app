package models

type AdvCreateReq struct {
	Name        string   `form:"name" binding:"required,min=1,max=200"`
	Price       float32  `form:"price" binding:"required"`
	Avatar      string   `form:"avatar" binding:"required"`
	Photos      []string `form:"photos" binding:"required,max=4"`
	Description string   `form:"description" binding:"max=1000"`
}
