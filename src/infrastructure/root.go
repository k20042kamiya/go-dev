package instructure

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Root() {
	// インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルートを設定
	e.GET("/", Hello) // ローカル環境の場合、http://localhost:1323/ にGETアクセスされるとhelloハンドラーを実行する

	// サーバーをポート番号1323で起動
	e.Logger.Fatal(e.Start(":1323"))
}

// ハンドラーを定義
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
