package routers

import (
	"encoding/json"
	"github.com/SaenkoDmitry/advertisement-app/db"
	"github.com/SaenkoDmitry/advertisement-app/models"
	"github.com/SaenkoDmitry/advertisement-app/utils"
	"github.com/SaenkoDmitry/advertisement-app/validations"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
)

func init() {
	binding.Validator = new(validations.DefaultValidator)
}

func AdvCreate(c *gin.Context) {
	var adv models.AdvCreateReq
	if ok := utils.ValidateAndBind(c, &adv); !ok {
		return
	}

	jsonPhotos, _ := json.Marshal(adv.Photos)
	newAdv := models.Adv{
		Name:        adv.Name,
		Price:       adv.Price,
		Avatar:      adv.Avatar,
		Photos:      postgres.Jsonb{RawMessage: jsonPhotos},
		Description: adv.Description,
	}
	db.GetDB().Create(&newAdv)
	c.JSON(200, gin.H{
		"id": newAdv.ID,
	})
}

func AdvGet(c *gin.Context) {
	id := c.Param("id")
	var adv models.Adv
	if result := db.GetDB().First(&adv, id); result.Error != nil {
		c.Status(http.StatusNotFound)
		return
	}
	var advDto models.AdvDto
	advDto.Id = adv.ID
	advDto.Name = adv.Name
	advDto.Price = adv.Price
	advDto.Avatar = adv.Avatar
	_, ok := c.GetQuery("fields")
	if ok {
		advDto.Photos = adv.Photos.RawMessage
		advDto.Description = adv.Description
	}
	c.JSON(200, advDto)
}

func AdvGetAll(c *gin.Context) {
	var params models.AdvSearchReq
	if ok := utils.ValidateAndBind(c, &params); !ok {
		return
	}

	var advs []models.Adv
	pagination.Paging(&pagination.Param{
		DB:    db.GetDB(),
		Page:  utils.GetIntOrDefault(params.Page, 1),
		Limit: utils.GetIntOrDefault(params.Limit, 10),
		OrderBy: []string{utils.GetStringOrDefault(params.OrderBy, "price") +
			" " + utils.GetStringOrDefault(params.Direction, "asc")},
		ShowSQL: true,
	}, &advs)

	var result []models.AdvDto
	for _, adv := range advs {
		advDto := models.AdvDto{
			Id:     adv.ID,
			Name:   adv.Name,
			Price:  adv.Price,
			Avatar: adv.Avatar,
		}
		result = append(result, advDto)
	}

	c.JSON(200, result)
}
