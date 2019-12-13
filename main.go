package main

import (
	"log"

	"go-blog/handler"
	"go-blog/repository"

	_ "github.com/go-sql-driver/mysql" // Using MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

// const tmplPath = "C:\\Users\\toru-ishikawa\\go\\src\\go-blog\\src\\template\\"

var db *sqlx.DB
var e = createMux()

func main() {
	db = connectDB()
	repository.SetDB(db)
	// パスと処理関数(ハンドラ)の紐付け
	/*
		e.GET("/", articleIndex)
		e.GET("/new", articleNew)
		// コロンから始まる部分はパスパラメータ となり、任意の文字列を受理
		e.GET("/:id", articleShow)
		e.GET("/:id/edit", articleEdit)
		// Webサーバーをポート番号 8080 で起動する

	*/
	e.GET("/", handler.ArticleIndex)
	e.GET("/new", handler.ArticleNew)
	e.GET("/:id", handler.ArticleShow1)
	e.GET("/:id/edit", handler.ArticleEdit)
	e.POST("/", handler.ArticleCreate)
	e.DELETE("/:id", handler.ArticleDelete)

	e.Logger.Fatal(e.Start("localhost:1111"))
}

func createMux() *echo.Echo {
	//インスタンス生成
	e := echo.New()

	//インスタンスに各種ミドルウェア設定
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.CSRF())

	e.Static("/css", "src/css")
	e.Static("/js", "src/js")

	e.Validator = &CustomValidator{validator: validator.New()}

	return e
}

// ハンドラ関数定義
// func articleIndex(c echo.Context) error {
// 	/*
// 	   HTTP リクエストの情報（リクエストの送信元や各種パラメータ等）は、
// 	   echo.Context という構造体でハンドラ関数に渡される。
// 	   リクエスト情報を構造体へ詰め替える作業や、ハンドラ関数に対して
// 	   データを引数で受け渡す部分は、echoが担当。
// 	*/
// 	// return c.String(http.StatusOK, "Hello World")

// 	data := map[string]interface{}{
// 		// "Message":  "Hello World",
// 		"Message":  "Article Index",
// 		"Now time": time.Now(),
// 	}
// 	return render(c, "article\\index.html", data)
// }

// func articleNew(c echo.Context) error {
// 	data := map[string]interface{}{
// 		"Message": "Article New",
// 		"Now":     time.Now,
// 	}
// 	return render(c, "article\\new.html", data)
// }

// func articleShow(c echo.Context) error {
// 	// パスパラメータの :id の部分を抽出
// 	// strconv.Atoiで文字列型→数値型へキャスト
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	data := map[string]interface{}{
// 		"Message": "Article Show",
// 		"Now":     time.Now(),
// 		"ID":      id,
// 	}
// 	return render(c, "article\\show.html", data)
// }

// func articleEdit(c echo.Context) error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	data := map[string]interface{}{
// 		"Message": "Article Edit",
// 		"Now":     time.Now(),
// 		"ID":      id,
// 	}
// 	return render(c, "article\\edit.html", data)
// }

// func htmlBlob(file string, data map[string]interface{}) ([]byte, error) {
// 	return pongo2.Must(pongo2.FromCache(tmplPath + file)).ExecuteBytes(data)
// }

// func render(c echo.Context, file string, data map[string]interface{}) error {
// 	// 生成された HTML をバイトデータとして受け取る
// 	b, err := htmlBlob(file, data)
// 	if err != nil {
// 		return c.NoContent(http.StatusInternalServerError)
// 	}
// 	// ステータスコード 200 で HTML データをレスポンス
// 	return c.HTMLBlob(http.StatusOK, b)
// }

func connectDB() *sqlx.DB {
	//環境変数から接続情報取得する場合もあり
	// dsn := "workuser:Passw0rd!@tcp(127.0.0.1:3306)/techblog?parseTime=true&autocommit=0&sql_mode=%27TRADITIONAL,NO_AUTO_VALUE_ON_ZERO,ONLY_FULL_GROUP_BY%27"
	dsn := "root:mysitepass@tcp(127.0.0.1:3306)/techblog?parseTime=true&autocommit=0&sql_mode=%27TRADITIONAL,NO_AUTO_VALUE_ON_ZERO,ONLY_FULL_GROUP_BY%27"

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded!!")
	return db
}

// CustomValidator ...
type CustomValidator struct {
	validator *validator.Validate
}

// Validate ...
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
