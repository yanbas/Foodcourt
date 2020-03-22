package app

import (
	"os"
	"strconv"

	"github.com/foodCourt/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type App struct {
	DB       *gorm.DB
	DBConfig *DBConfig
	Gin      *gin.Engine
	Log      *os.File
}

func (a *App) Running() {
	a.DB.AutoMigrate(&model.Menu{}, &model.Order{}, &model.Payment{})
	a.DB.Model(&model.Order{}).AddForeignKey("menu_id", "menus(id)", "RESTRICT", "RESTRICT")
	a.DB.Model(&model.Payment{}).AddForeignKey("order_id", "orders(id)", "RESTRICT", "RESTRICT")

	a.Gin = gin.Default()
	a.Gin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	a.Gin.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	// Food Feature Routes
	a.Gin.GET("/menu", a.GetAllFood)
	a.Gin.GET("/menu/:id", a.GetFood)
	a.Gin.POST("/menu/", a.CreateFood)
	a.Gin.PUT("/menu/:id", a.UpdateFood)
	a.Gin.DELETE("/menu/:id", a.DeleteFood)

	// Order Feature Routes
	a.Gin.GET("/order/", a.GetAllOrders)
	a.Gin.GET("/order/:id", a.GetOrder)
	a.Gin.POST("/order/", a.CreateOrder)
	a.Gin.PUT("/order/:id", a.UpdateOrder)
	a.Gin.DELETE("/order/:id", a.DeleteOrder)

	// Payment Feature Routes
	a.Gin.GET("/payment", a.GetAllPayments)
	a.Gin.GET("/payment/:id", a.GetPayment)
	a.Gin.POST("/payment/", a.CreatePayment)
	a.Gin.PUT("/payment/:id", a.UpdatePayment)
	a.Gin.DELETE("/payment/:id", a.DeletePayment)

	a.Gin.Run()
}

func (a *App) GetAllFood(c *gin.Context) {
	menu := []model.Menu{}

	a.Logger("Start process get all menu")

	if err := a.DB.Find(&menu).Error; err != nil {
		a.Logger(err.Error())
		c.JSON(501, RESULTS_ERROR)
	}

	c.JSON(200, ResultData{
		Code:    100,
		Success: true,
		Data:    menu,
	})
	a.Logger("Success Process \n")
}

func (a *App) CreateFood(c *gin.Context) {
	input := &Menu{}
	err := c.ShouldBindJSON(input)
	if err != nil {
		a.Logger(err.Error())
		c.JSON(501, BINDING_JSON_ERROR)
		return
	}

	menu := model.Menu{
		Name:   input.Name,
		Price:  input.Price,
		Status: input.Status,
	}

	if err := a.DB.Create(&menu).Error; err != nil {
		a.Logger(err.Error())
		c.JSON(501, INSERT_DATA_ERROR)
		return
	}

	c.JSON(201, ResultData{
		Code:    100,
		Success: true,
		Data:    "",
	})
}

func (a *App) GetFood(c *gin.Context) {
	menu := model.Menu{}

	if err := a.DB.First(&menu, c.Param("id")).Error; err != nil {
		a.Logger(err.Error())
		c.JSON(500, RESULT_ERROR)
		return
	}

	c.JSON(200, ResultData{
		Code:    100,
		Success: true,
		Data:    menu,
	})
}

func (a *App) UpdateFood(c *gin.Context) {
	input := &Menu{}
	err := c.ShouldBindJSON(input)
	if err != nil {
		a.Logger(err.Error())
		c.JSON(500, BINDING_JSON_ERROR)
		return
	}

	menu := model.Menu{}
	id, _ := strconv.Atoi(c.Param("id"))

	if err := a.DB.First(&menu, c.Param("id")).Error; err != nil {
		a.Logger(err.Error())
		c.JSON(500, RESULT_ERROR)
	}

	menu = model.Menu{
		Name:   input.Name,
		Price:  input.Price,
		Status: input.Status,
	}

	if err := a.DB.Model(&menu).Where("id = ?", id).Updates(&menu).Error; err != nil {
		a.Logger(err.Error())
		c.JSON(500, UPDATE_DATA_ERROR)
	}

	c.JSON(204, ResultData{
		Code:    100,
		Success: true,
		Data:    "",
	})
}

func (a *App) DeleteFood(c *gin.Context) {
	menu := model.Menu{}

	if err := a.DB.Delete(&menu, c.Param("id")).Error; err != nil {
		a.Logger(err.Error())
		c.JSON(500, DELET_DATA_ERROR)
		return
	}

	c.JSON(204, ResultData{
		Code:    100,
		Success: true,
		Data:    menu,
	})
}

//Order
func (a *App) GetAllOrders(c *gin.Context) {
	orders := []model.Order{}

	a.Logger("Start process get all orders")

	if err := a.DB.Find(&orders).Error; err != nil {
		a.Logger(err.Error())
		c.JSON(501, RESULTS_ERROR)
		return
	}

	c.JSON(200, ResultData{
		Code:    100,
		Success: true,
		Data:    orders,
	})
	a.Logger("Success Process Orders \n")

}

func (a *App) GetOrder(c *gin.Context) {
	order := model.Order{}

	if err := a.DB.First(&order, c.Param("id")).Error; err != nil {
		a.Logger(err.Error())
		c.JSON(500, RESULT_ERROR)
		return
	}

	c.JSON(200, ResultData{
		Code:    100,
		Success: true,
		Data:    order,
	})
}

func (a *App) CreateOrder(c *gin.Context) {
	input := &Order{}
	err := c.ShouldBindJSON(input)
	if err != nil {
		a.Logger(err.Error())
		c.JSON(501, BINDING_JSON_ERROR)
		return
	}

	order := model.Order{
		CustomerName: input.CustomerName,
		MenuId:       input.Menu,
		Qty:          input.Qty,
		Status:       input.Status,
		TableNumber:  input.TableNumber,
	}

	if err := a.DB.Create(&order).Error; err != nil {
		a.Logger(err.Error())
		c.JSON(501, INSERT_DATA_ERROR)
		return
	}

	c.JSON(201, ResultData{
		Code:    100,
		Success: true,
		Data:    "",
	})
}

func (a *App) UpdateOrder(c *gin.Context) {
	input := &Order{}
	err := c.ShouldBindJSON(input)
	if err != nil {
		a.Logger(err.Error())
		c.JSON(500, BINDING_JSON_ERROR)
		return
	}

	order := model.Order{}
	id, _ := strconv.Atoi(c.Param("id"))

	if err := a.DB.First(&order, c.Param("id")).Error; err != nil {
		a.Logger(err.Error())
		c.JSON(500, RESULT_ERROR)
		return
	}

	order = model.Order{
		CustomerName: input.CustomerName,
		Qty:          input.Qty,
		Status:       input.Status,
		TableNumber:  input.TableNumber,
		MenuId:       input.Menu,
	}

	if err := a.DB.Model(&order).Where("id = ?", id).Updates(&order).Error; err != nil {
		a.Logger(err.Error())
		c.JSON(500, UPDATE_DATA_ERROR)
		return
	}

	c.JSON(204, ResultData{
		Code:    100,
		Success: true,
		Data:    "",
	})
}

func (a *App) DeleteOrder(c *gin.Context) {
	order := model.Order{}

	if err := a.DB.Delete(&order, c.Param("id")).Error; err != nil {
		a.Logger(err.Error())
		c.JSON(500, DELET_DATA_ERROR)
	}

	c.JSON(204, ResultData{
		Code:    100,
		Success: true,
		Data:    order,
	})
}

//Payment
func (a *App) GetAllPayments(c *gin.Context) {
	payments := []model.Payment{}

	a.Logger("Start process get all payments")

	if err := a.DB.Find(&payments).Error; err != nil {
		a.Logger(err.Error())
		c.JSON(501, RESULTS_ERROR)
	}

	c.JSON(200, ResultData{
		Code:    100,
		Success: true,
		Data:    payments,
	})
	a.Logger("Success Process Orders \n")
}

func (a *App) GetPayment(c *gin.Context) {
	payment := model.Payment{}

	if err := a.DB.First(&payment, c.Param("id")).Error; err != nil {
		a.Logger(err.Error())
		c.JSON(500, RESULT_ERROR)
		return
	}

	c.JSON(200, ResultData{
		Code:    100,
		Success: true,
		Data:    payment,
	})
}

func (a *App) CreatePayment(c *gin.Context) {
	input := &Payment{}
	err := c.ShouldBindJSON(input)
	if err != nil {
		a.Logger(err.Error())
		c.JSON(501, BINDING_JSON_ERROR)
		return
	}

	payment := model.Payment{
		Amount:          input.Amount,
		OrderId:         input.OrderId,
		Status:          input.Status,
		ReferenceNumber: input.ReferenceNumber,
		Type:            input.Type,
	}

	if err := a.DB.Create(&payment).Error; err != nil {
		a.Logger(err.Error())
		c.JSON(501, INSERT_DATA_ERROR)
		return
	}

	c.JSON(201, ResultData{
		Code:    100,
		Success: true,
		Data:    "",
	})
}

func (a *App) UpdatePayment(c *gin.Context) {
	input := &Payment{}
	err := c.ShouldBindJSON(input)
	if err != nil {
		a.Logger(err.Error())
		c.JSON(500, BINDING_JSON_ERROR)
		return
	}

	payment := model.Payment{}
	id, _ := strconv.Atoi(c.Param("id"))

	if err := a.DB.First(&payment, c.Param("id")).Error; err != nil {
		a.Logger(err.Error())
		c.JSON(500, RESULT_ERROR)
		return
	}

	payment = model.Payment{
		Amount:          input.Amount,
		OrderId:         input.OrderId,
		Status:          input.Status,
		ReferenceNumber: input.ReferenceNumber,
		Type:            input.Type,
	}

	if err := a.DB.Model(&payment).Where("id = ?", id).Updates(&payment).Error; err != nil {
		a.Logger(err.Error())
		c.JSON(500, UPDATE_DATA_ERROR)
		return
	}

	c.JSON(204, ResultData{
		Code:    100,
		Success: true,
		Data:    "",
	})
}

func (a *App) DeletePayment(c *gin.Context) {
	payment := model.Payment{}

	if err := a.DB.Delete(&payment, c.Param("id")).Error; err != nil {
		a.Logger(err.Error())
		c.JSON(500, DELET_DATA_ERROR)
	}

	c.JSON(204, ResultData{
		Code:    100,
		Success: true,
		Data:    payment,
	})
}
