package main

import (
	"log"
	"os"

	"github.com/aster-void/openai-go-template/app/env"
	"github.com/aster-void/openai-go-template/lib/server"
	"github.com/aster-void/openai-go-template/router"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// public ディレクトリ下のファイルに適切なパスでアクセスできるようにする
	e.Static("/", "./public")

	if env.DEV_MODE {
		router.Test(e.Group("/test"))
	}
	router.Chat(e.Group("/chat"))

	// 使用するホスティングサービス (Render など) によってはリクエストを受け付けるポートが指定されている場合がある。
	// たいていの場合は PORT という名前の環境変数を通して参照できる。
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err := server.StartWithProperShutdown(e, ":"+port)
	if err != nil {
		log.Fatalln(err)
	}
}
