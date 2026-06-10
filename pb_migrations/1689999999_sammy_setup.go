package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/dbx"
)

func init() {
	// 创建 Sammy 档案集合
	App.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// 检查集合是否已存在
		if App.Dao().FindCollectionByNameOrId("sammy_profiles") != nil {
			return nil
		}

		// 创建 sammy_profiles 集合
		collection := &core.Collection{
			Name:       "sammy_profiles",
			Type:       "base",
			ListRule:   "@request.auth.id = user",
			ViewRule:   "@request.auth.id = user",
			CreateRule: "@request.auth.id = user",
			UpdateRule: "@request.auth.id = user",
			DeleteRule: "@request.auth.id = user",
			Schema: []core.SchemaField{
				{
					Name:     "user",
					Type:     "relation",
					Required: true,
					Options: core.RelationOptions{
						CollectionId: "_pb_users_auth_",
						MaxSelect:    1,
					},
				},
				{
					Name:     "name",
					Type:     "text",
					Required: true,
				},
				{
					Name:     "gender",
					Type:     "select",
					Required: true,
					Options: core.SelectOptions{
						Values: []string{"boy", "girl"},
					},
				},
				{
					Name:     "birthDate",
					Type:     "date",
					Required: true,
				},
				{
					Name:     "personality",
					Type:     "select",
					Required: true,
					Options: core.SelectOptions{
						Values: []string{"活泼开朗", "温柔乖巧", "调皮可爱", "安静内敛", "治愈软萌"},
					},
				},
				{
					Name:     "appearance",
					Type:     "text",
					Required: false,
				},
				{
					Name:     "growthPoints",
					Type:     "number",
					Required: false,
					Options: core.NumberOptions{
						Min: dbx.NewFloat64(0),
					},
				},
				{
					Name:     "mood",
					Type:     "text",
					Required: false,
				},
				{
					Name:     "lastInteraction",
					Type:     "date",
					Required: false,
				},
			},
		}

		if err := App.Dao().SaveCollection(collection); err != nil {
			return err
		}

		// 创建 interactions 集合
		if App.Dao().FindCollectionByNameOrId("interactions") != nil {
			return nil
		}

		interactionCollection := &core.Collection{
			Name:       "interactions",
			Type:       "base",
			ListRule:   "@request.auth.id = user",
			ViewRule:   "@request.auth.id = user",
			CreateRule: "@request.auth.id = user",
			UpdateRule: "@request.auth.id = user",
			DeleteRule: "@request.auth.id = user",
			Schema: []core.SchemaField{
				{
					Name:     "user",
					Type:     "relation",
					Required: true,
					Options: core.RelationOptions{
						CollectionId: "_pb_users_auth_",
						MaxSelect:    1,
					},
				},
				{
					Name:     "sammy",
					Type:     "relation",
					Required: true,
					Options: core.RelationOptions{
						CollectionId: "sammy_profiles",
						MaxSelect:    1,
					},
				},
				{
					Name:     "type",
					Type:     "select",
					Required: true,
					Options: core.SelectOptions{
						Values: []string{"chat", "voice", "play", "sleep", "feed"},
					},
				},
				{
					Name:     "content",
					Type:     "text",
					Required: true,
				},
				{
					Name:     "moodChange",
					Type:     "number",
					Required: false,
				},
				{
					Name:     "timestamp",
					Type:     "date",
					Required: true,
				},
			},
		}

		return App.Dao().SaveCollection(interactionCollection)
	})
}
