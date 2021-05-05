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

	gin.SetMode(gin.ReleaseMode)
	//setup router
	router := gin.Default()
	v1 := router.Group("v1")
	{
		portal := v1.Group("portal")
		{
			portal.GET("/groups/:groupID/restaurants", bean.PortalService.GetRestaurantsByRestaurantGroupID)
			portal.GET("/groups/:groupID/categories", bean.PortalService.GetRestaurantGroupCategories)
			portal.GET("/groups/:groupID/products/:productID/comments", bean.PortalService.GetRestaurantGroupCommentsByProductID)
			portal.GET("/restaurants", bean.PortalService.GetRestaurantByCode)
			portal.GET("/restaurants/:restaurantID/areas", bean.PortalService.GetRestaurantAreas)
			portal.GET("/restaurants/:restaurantID/areas/:areaID/tables", bean.PortalService.GetRestaurantTablesByAreaID)
			portal.GET("/restaurants/:restaurantID/areas/:areaID/orders", bean.PortalService.GetRestaurantOrdersByAreaIDAndGroupByOrderNumberID)
			portal.GET("/restaurants/:restaurantID/categories/:categoryID/products", bean.PortalService.GetRestaurantProductsByCategoryID)
			portal.GET("/restaurants/:restaurantID/orders", bean.PortalService.GetRestaurantOrderDetails)
			portal.PUT("/restaurants/:restaurantID/orders", bean.PortalService.PutRestaurantOrders)
		}

	}
	router.Run(":8081")
}
