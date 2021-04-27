package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/smarest/smarest-api/application"
)

func main() {
	//connect to DB
	bean, err := application.InitBean()
	defer bean.DestroyBean()
	if err != nil {
		log.Fatalln("can not create bean", err)
	}
	//setup router
	router := gin.Default()
	v1 := router.Group("v1")
	{
		restaurant := v1.Group("restaurants")
		{
			restaurant.GET("/:id", bean.RestaurantService.GetByID)
			restaurant.GET("", bean.RestaurantService.GetAll)
			restaurant.GET("/:id/areas", bean.RestaurantService.GetAreas)
			restaurant.GET("/:id/products", bean.RestaurantService.GetProducts)
		}

		area := v1.Group("areas")
		{
			area.GET("", bean.AreaService.GetAll)
			area.GET("/:id", bean.AreaService.GetByID)
			area.GET("/:id/tables", bean.AreaService.GetTables)
			area.GET("/:id/orders", bean.AreaService.GetOrders)
		}

		table := v1.Group("tables")
		{
			table.GET("", bean.TableService.GetAll)
			table.GET("/:id", bean.TableService.GetByID)
		}

		product := v1.Group("products")
		{
			product.GET("/:id", bean.ProductService.GetByID)
			product.GET("", bean.ProductService.GetAll)
		}

		ingredient := v1.Group("ingredients")
		{
			ingredient.GET("/:id", bean.IngredientService.GetByID)
			ingredient.GET("", bean.IngredientService.GetAll)

		}

		category := v1.Group("categories")
		{
			category.GET("/:id", bean.CategoryService.GetByID)
			category.GET("", bean.CategoryService.GetAll)

		}

		order := v1.Group("orders")
		{
			order.GET("", bean.OrderService.GetOrderDetails)
			order.PUT("", bean.OrderService.PutOrders)

		}
		v1.GET("/comments/:id", bean.CommentService.GetByID)
		v1.GET("/comments", bean.CommentService.GetAll)

		/*  v1.GET("/instructions/:id", app.GetInstruction)
		    v1.POST("/instructions", app.PostInstruction)
		    v1.PUT("/instructions/:id", app.UpdateInstruction)
		    v1.DELETE("/instructions/:id", app.DeleteInstruction)*/
	}
	router.Run(":8080")
}
