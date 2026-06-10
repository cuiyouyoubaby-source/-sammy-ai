package main

import (
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func main() {
	app := pocketbase.New()

	// 注册自定义路由
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// 静态文件服务
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./frontend"), true))

		// Sammy 成长状态API
		e.Router.GET("/api/sammy/growth/:userId", func(c echo.Context) error {
			userId := c.PathParam("userId")
			
			// 获取 Sammy 成长数据
			return c.JSON(200, map[string]interface{}{
				"userId": userId,
				"age": "3岁",
				"mood": "开心",
				"growthPoints": 1250,
				"lastInteraction": time.Now().Format("2006-01-02 15:04:05"),
				"personality": "活泼开朗",
				"nextMilestone": "4岁生日",
				"daysUntilMilestone": 15,
			})
		})

		// Sammy 互动API
		e.Router.POST("/api/sammy/interact", func(c echo.Context) error {
			var data map[string]interface{}
			if err := c.Bind(&data); err != nil {
				return c.JSON(400, map[string]string{"error": "invalid request"})
			}
			
			// 处理互动逻辑
			return c.JSON(200, map[string]interface{}{
				"status": "success",
				"response": "Sammy 收到了你的互动！",
				"moodChange": "+5",
				"growthPoints": 10,
				"timestamp": time.Now().Format("2006-01-02 15:04:05"),
			})
		})

		// 获取 Sammy 列表
		e.Router.GET("/api/sammy/list", func(c echo.Context) error {
			// 返回 Sammy 列表
			return c.JSON(200, map[string]interface{}{
				"items": []map[string]interface{}{
					{
						"id": "sammy_001",
						"name": "Sammy",
						"age": "3岁",
						"mood": "开心",
						"personality": "活泼开朗",
						"growthPoints": 1250,
					},
				},
			})
		})

		// 创建 Sammy
		e.Router.POST("/api/sammy/create", func(c echo.Context) error {
			var data map[string]interface{}
			if err := c.Bind(&data); err != nil {
				return c.JSON(400, map[string]string{"error": "invalid request"})
			}
			
			// 创建 Sammy 逻辑
			return c.JSON(200, map[string]interface{}{
				"status": "success",
				"message": "Sammy 创建成功！",
				"sammyId": "sammy_001",
				"timestamp": time.Now().Format("2006-01-02 15:04:05"),
			})
		})

		// 更新 Sammy 状态
		e.Router.PUT("/api/sammy/update/:sammyId", func(c echo.Context) error {
			sammyId := c.PathParam("sammyId")
			
			var data map[string]interface{}
			if err := c.Bind(&data); err != nil {
				return c.JSON(400, map[string]string{"error": "invalid request"})
			}
			
			// 更新 Sammy 逻辑
			return c.JSON(200, map[string]interface{}{
				"status": "success",
				"message": "Sammy 状态更新成功！",
				"sammyId": sammyId,
				"timestamp": time.Now().Format("2006-01-02 15:04:05"),
			})
		})

		// 获取互动历史
		e.Router.GET("/api/sammy/interactions/:sammyId", func(c echo.Context) error {
			sammyId := c.PathParam("sammyId")
			
			// 返回互动历史
			return c.JSON(200, map[string]interface{}{
				"sammyId": sammyId,
				"interactions": []map[string]interface{}{
					{
						"id": "int_001",
						"type": "chat",
						"content": "你好呀！我是 Sammy！",
						"timestamp": time.Now().Add(-24 * time.Hour).Format("2006-01-02 15:04:05"),
						"moodChange": 5,
					},
					{
						"id": "int_002",
						"type": "play",
						"content": "Sammy 玩得好开心！",
						"timestamp": time.Now().Add(-12 * time.Hour).Format("2006-01-02 15:04:05"),
						"moodChange": 10,
					},
				},
			})
		})

		return nil
	})

	// 启动应用
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
