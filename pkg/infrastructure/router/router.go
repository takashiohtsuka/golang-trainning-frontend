package router

import (
	"golang-trainning-frontend/pkg/adapter/controller"
	"golang-trainning-frontend/pkg/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: config.C.Server.AllowOrigins,
		AllowMethods: []string{echo.GET},
	}))

	g := e.Group("/frontend")

	g.GET("/stores/:id", func(ctx echo.Context) error { return c.Store.GetStoreDetail(ctx) })
	g.GET("/stores/:id/women", func(ctx echo.Context) error { return c.Woman.GetStoreWomanList(ctx) })

	g.GET("/women", func(ctx echo.Context) error { return c.Woman.GetWomanList(ctx) })
	g.GET("/women/:id", func(ctx echo.Context) error { return c.Woman.GetWomanDetail(ctx) })

	g.GET("/districts/:id/women", func(ctx echo.Context) error { return c.WomanDistrict.GetWomanDistrictList(ctx) })
	g.GET("/districts/:id/search-woman-count", func(ctx echo.Context) error { return c.WomanDistrictCount.GetWomanDistrictCount(ctx) })

	g.GET("/immediate_available_women", func(ctx echo.Context) error { return c.ImmediateAvailableWoman.GetImmediateAvailableWomanList(ctx) })

	g.GET("/prefectures", func(ctx echo.Context) error { return c.Prefecture.GetPrefectureList(ctx) })
	g.GET("/prefectures/:id/districts", func(ctx echo.Context) error { return c.District.GetDistrictList(ctx) })

	g.GET("/business_types", func(ctx echo.Context) error { return c.BusinessType.GetBusinessTypeList(ctx) })

	return e
}
