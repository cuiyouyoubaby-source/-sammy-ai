package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

// SammyResponse 定义 Sammy 的回复结构
type SammyResponse struct {
	Text     string `json:"text"`
	Emotion  string `json:"emotion"`
	Action   string `json:"action"`
	AudioURL string `json:"audio_url,omitempty"`
}

// InteractionRequest 定义互动请求
type InteractionRequest struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func main() {
	app := pocketbase.New()

	// Sammy 的回复模板
	sammyResponses := map[string][]string{
		"chat": {
			"真的吗？好有趣呀！",
			"Sammy 也这么想！",
			"哇，好厉害！",
			"Sammy 想和你一起玩～",
			"今天好开心呀！",
			"Sammy 最喜欢你了！",
			"我们一起唱歌吧！",
			"Sammy 给你讲个故事吧！",
			"你觉得呢？Sammy 觉得好棒！",
			"Sammy 学到了新东西！",
		},
		"feed": {
			"Sammy 吃得好香呀～",
			"这个好好吃！",
			"Sammy 吃饱了，谢谢你！",
			"Sammy 最喜欢这个了！",
			"Sammy 要长得高高的！",
		},
		"play": {
			"Sammy 玩得好开心！",
			"再来一次！",
			"Sammy 好喜欢这个玩具！",
			"我们玩什么好呢？",
			"Sammy 要赢啦！",
		},
		"sleep": {
			"Sammy 要睡觉觉了，晚安～",
			"Sammy 好困呀...",
			"明天见！",
			"Sammy 会做个好梦的！",
			"晚安，好梦～",
		},
	}

	// Sammy 的情绪反应
	sammyEmotions := map[string][]string{
		"chat":   {"happy", "excited", "curious"},
		"feed":   {"happy", "satisfied", "grateful"},
		"play":   {"excited", "joyful", "playful"},
		"sleep":  {"sleepy", "peaceful", "tired"},
	}

	// 注册自定义路由
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// 静态文件服务
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./frontend"), true))

		// Sammy 互动 API
		e.Router.POST("/api/sammy/interact", func(c echo.Context) error {
			var req InteractionRequest
			if err := c.Bind(&req); err != nil {
				return c.JSON(400, map[string]string{"error": "invalid request"})
			}

			// 获取随机回复
			responses := sammyResponses[req.Type]
			if responses == nil {
				responses = sammyResponses["chat"]
			}
			
			emotions := sammyEmotions[req.Type]
			if emotions == nil {
				emotions = sammyEmotions["chat"]
			}

			response := SammyResponse{
				Text:    responses[rand.Intn(len(responses))],
				Emotion: emotions[rand.Intn(len(emotions))],
				Action:  req.Type,
			}

			return c.JSON(200, response)
		})

		// Sammy 聊天 API
		e.Router.POST("/api/sammy/chat", func(c echo.Context) error {
			var data map[string]interface{}
			if err := c.Bind(&data); err != nil {
				return c.JSON(400, map[string]string{"error": "invalid request"})
			}

			userMessage := data["message"].(string)
			
			// 根据用户消息生成回复
			responses := []string{
				"真的吗？好有趣呀！",
				"Sammy 也这么想！",
				"哇，好厉害！",
				"Sammy 想和你一起玩～",
				"今天好开心呀！",
				"Sammy 最喜欢你了！",
				"我们一起唱歌吧！",
				"Sammy 给你讲个故事吧！",
			}
			
			response := SammyResponse{
				Text:    responses[rand.Intn(len(responses))],
				Emotion: "happy",
				Action:  "chat",
			}

			return c.JSON(200, response)
		})

		// Sammy 成长状态 API
		e.Router.GET("/api/sammy/growth/:userId", func(c echo.Context) error {
			userId := c.PathParam("userId")
			
			return c.JSON(200, map[string]interface{}{
				"userId":            userId,
				"age":              "3岁",
				"mood":             "开心",
				"growthPoints":     1250,
				"lastInteraction":  time.Now().Format("2006-01-02 15:04:05"),
				"personality":      "活泼开朗",
				"nextMilestone":    "4岁生日",
				"daysUntilMilestone": 15,
			})
		})

		// Sammy 列表 API
		e.Router.GET("/api/sammy/list", func(c echo.Context) error {
			return c.JSON(200, map[string]interface{}{
				"items": []map[string]interface{}{
					{
						"id":            "sammy_001",
						"name":         "Sammy",
						"age":          "3岁",
						"mood":         "开心",
						"personality":  "活泼开朗",
						"growthPoints": 1250,
					},
				},
			})
		})

		// 创建 Sammy API
		e.Router.POST("/api/sammy/create", func(c echo.Context) error {
			var data map[string]interface{}
			if err := c.Bind(&data); err != nil {
				return c.JSON(400, map[string]string{"error": "invalid request"})
			}
			
			return c.JSON(200, map[string]interface{}{
				"status":    "success",
				"message":   "Sammy 创建成功！",
				"sammyId":   "sammy_001",
				"timestamp": time.Now().Format("2006-01-02 15:04:05"),
			})
		})

		// 更新 Sammy 状态 API
		e.Router.PUT("/api/sammy/update/:sammyId", func(c echo.Context) error {
			sammyId := c.PathParam("sammyId")
			
			var data map[string]interface{}
			if err := c.Bind(&data); err != nil {
				return c.JSON(400, map[string]string{"error": "invalid request"})
			}
			
			return c.JSON(200, map[string]interface{}{
				"status":    "success",
				"message":   "Sammy 状态更新成功！",
				"sammyId":   sammyId,
				"timestamp": time.Now().Format("2006-01-02 15:04:05"),
			})
		})

		// 获取互动历史 API
		e.Router.GET("/api/sammy/interactions/:sammyId", func(c echo.Context) error {
			sammyId := c.PathParam("sammyId")
			
			return c.JSON(200, map[string]interface{}{
				"sammyId": sammyId,
				"interactions": []map[string]interface{}{
					{
						"id":          "int_001",
						"type":       "chat",
						"content":    "你好呀！我是 Sammy，今天也要开心哦～",
						"timestamp":  time.Now().Add(-24 * time.Hour).Format("2006-01-02 15:04:05"),
						"moodChange": 5,
					},
					{
						"id":          "int_002",
						"type":       "play",
						"content":    "Sammy 玩得好开心！",
						"timestamp":  time.Now().Add(-12 * time.Hour).Format("2006-01-02 15:04:05"),
						"moodChange": 10,
					},
				},
			})
		})

		// 语音合成 API (TTS)
	e.Router.POST("/api/sammy/tts", func(c echo.Context) error {
			var data map[string]interface{}
			if err := c.Bind(&data); err != nil {
				return c.JSON(400, map[string]string{"error": "invalid request"})
			}

			text := data["text"].(string)
			
			// 返回 TTS 配置信息（前端使用 Web Speech API）
			return c.JSON(200, map[string]interface{}{
				"status":  "success",
				"text":    text,
				"voice":   "zh-CN",
				"rate":    1.2,
				"pitch":   1.5,
				"message": "请使用浏览器 Web Speech API 进行语音合成",
			})
		})

		return nil
	})

	// 启动应用
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
