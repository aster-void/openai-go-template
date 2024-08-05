package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

var OPENAI_API_KEY string
var client *openai.Client

func init() {
	if envfile := os.Getenv("ENV_FILE"); envfile != "" {
		fmt.Println("loading env file", envfile, "...")
		err := godotenv.Load(envfile)
		if err != nil {
			log.Fatalln("Failed to load env file:", err)
		}
	}

	OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")
	if OPENAI_API_KEY == "" {
		log.Fatalln("You must provide an API key!")
	}

	client = openai.NewClient(OPENAI_API_KEY)
}

func main() {
	e := echo.New()

	// public ディレクトリ下のファイルに適切なパスでアクセスできるようにする
	e.Static("/", "./public")

	e.POST("/chat", func(c echo.Context) error {
		// リクエストボディを JSON として解釈して request.body に格納する
		// (クライアントから送られてきたデータは無条件で信用しない)
		var request struct {
			PromptText string `json:"promptText"`
		}
		err := json.NewDecoder(c.Request().Body).Decode(&request)
		if err != nil {
			return c.NoContent(400)
		}

		text, err := InvokeAI(request.PromptText, context.Background())
		if err != nil {
			fmt.Println(err)
			return c.NoContent(500)
		}
		fmt.Println(text)

		type response struct {
			Content string `json:"content"`
		}
		return c.JSON(200, response{
			Content: text,
		})
	})

	// 使用するホスティングサービス (Render など) によってはリクエストを受け付けるポートが指定されている場合がある。
	// たいていの場合は PORT という名前の環境変数を通して参照できる。
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := e.Start(":" + port); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalln(err)
	}
}

func InvokeAI(s string, ctx context.Context) (string, error) {
	// example copied from: https://github.com/sashabaranov/go-openai
	res, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return res.Choices[0].Message.Content, nil
}
