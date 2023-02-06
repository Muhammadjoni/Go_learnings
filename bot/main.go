package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/infracloudio/msbotbuilder-go/core"
	"github.com/infracloudio/msbotbuilder-go/core/activity"
	"github.com/infracloudio/msbotbuilder-go/schema"
	"github.com/spf13/viper"
)

type TeamsConnection struct {
	AppId       string `mapstructure:"APP_ID"`
	AppPassword string `mapstructure:"APP_PASSWORD"`
	AppTenantId string `mapstructure:"APP_TENANTID"`
}

// Card content
// Visit: https://adaptivecards.io/explorer to build your own card format
/*
var cardJSON = []byte(`{
	"$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
	"type": "AdaptiveCard",
	"version": "1.0",
	"body": [
	  {
		"type": "Container",
		"items": [
		  {
			"type": "TextBlock",
			"text": "Publish Adaptive Card schema",
			"weight": "bolder",
			"size": "medium"
		  },
		  {
			"type": "ColumnSet",
			"columns": [
			  {
				"type": "Column",
				"width": "auto",
				"items": [
				  {
					"type": "Image",
					"url": "https://pbs.twimg.com/profile_images/3647943215/d7f12830b3c17a5a9e4afcc370e3a37e_400x400.jpeg",
					"altText": "Matt Hidinger",
					"size": "small",
					"style": "person"
				  }
				]
			  },
			  {
				"type": "Column",
				"width": "stretch",
				"items": [
				  {
					"type": "TextBlock",
					"text": "Matt Hidinger",
					"weight": "bolder",
					"wrap": true
				  },
				  {
					"type": "TextBlock",
					"spacing": "none",
					"text": "Created {{DATE(2017-02-14T06:08:39Z, SHORT)}}",
					"isSubtle": true,
					"wrap": true
				  }
				]
			  }
			]
		  }
		]
	  },
	  {
		"type": "Container",
		"items": [
		  {
			"type": "TextBlock",
			"text": "Now that we have defined the main rules and features of the format, we need to produce a schema and publish it to GitHub. The schema will be the starting point of our reference documentation.",
			"wrap": true
		  },
		  {
			"type": "FactSet",
			"facts": [
			  {
				"title": "Board:",
				"value": "Adaptive Card"
			  },
			  {
				"title": "List:",
				"value": "Backlog"
			  },
			  {
				"title": "Assigned to:",
				"value": "Matt Hidinger"
			  },
			  {
				"title": "Due date:",
				"value": "Not set"
			  }
			]
		  }
		]
	  }
	],
	"actions": [
	  {
		"type": "Action.ShowCard",
		"title": "Comment",
		"card": {
		  "type": "AdaptiveCard",
		  "body": [
			{
			  "type": "Input.Text",
			  "id": "comment",
			  "isMultiline": true,
			  "placeholder": "Enter your comment"
			}
		  ],
		  "actions": [
			{
			  "type": "Action.Submit",
			  "title": "OK"
			}
		  ]
		}
	  },
	  {
		"type": "Action.OpenUrl",
		"title": "View",
		"url": "https://adaptivecards.io"
	  }
	]
  }`)
*/
// conversationRef to store conversation reference to which proactive messages will be sent

var customHandler = activity.HandlerFuncs{
	OnMessageFunc: func(turn *activity.TurnContext) (schema.Activity, error) {
		return turn.SendActivity(activity.MsgOptionText("Echo: " + turn.Activity.Text))
	},
}

/*
var conversationRef schema.ConversationReference
var welcomeHandler = activity.HandlerFuncs{
	OnMessageFunc: func(turn *activity.TurnContext) (schema.Activity, error) {
		return turn.SendActivity(activity.MsgOptionText("Welcome"))
	},
}

var attachHandler = activity.HandlerFuncs{
	OnMessageFunc: func(turn *activity.TurnContext) (schema.Activity, error) {
		var obj map[string]interface{}
		err := json.Unmarshal(cardJSON, &obj)
		if err != nil {
			return schema.Activity{}, err
		}
		attachments := []schema.Attachment{
			{
				ContentType: "application/vnd.microsoft.card.adaptive",
				Content:     obj,
			},
		}
		return turn.SendActivity(activity.MsgOptionText("Sample attachment"), activity.MsgOptionAttachments(attachments))
	},
}
*/

// HTTPHandler handles the HTTP requests from then connector service
type HTTPHandler struct {
	core.Adapter
}

func (ht *HTTPHandler) processMessage(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	act, err := ht.Adapter.ParseRequest(ctx, req)

	// // Set conversation reference
	// conversationRef = activity.GetCoversationReference(act)

	if err != nil {
		fmt.Println("Failed to parse request.", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ht.ProcessActivity(ctx, act, customHandler)
	if err != nil {
		fmt.Println("Failed to process request.", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Request processed successfully.")
	// Send proactive message
	// ht.ProactiveMessage(context.TODO(), conversationRef, welcomeHandler)
	// ht.welcome()
}

// func (ht *HTTPHandler) welcome() {
// 	err := ht.Adapter.ProactiveMessage(context.TODO(), conversationRef, attachHandler)
// 	if err != nil {
// 		fmt.Println("Failed to send proactive message.", err)
// 		return
// 	}
// 	fmt.Println("Proactive message sent successfully.")
// }

func LoadConfig(path string) (config TeamsConnection, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func main() {

	config, err := LoadConfig(".")
	if err != nil {
		log.Fatalf("couldnt load config")
	}

	setting := core.AdapterSetting{
		AppID:             config.AppId,
		AppPassword:       config.AppPassword,
		ChannelAuthTenant: config.AppTenantId,
	}

	adapter, err := core.NewBotAdapter(setting)
	if err != nil {
		log.Fatal("Error creating adapter: ", err)
	}

	httpHandler := &HTTPHandler{adapter}

	http.HandleFunc("/api/messages", httpHandler.processMessage)
	fmt.Println("Starting server on port:8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error in adapter: ", err)
	}
}
