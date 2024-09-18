package router

import (
	"net/http"
	"os"

	"github.com/1206yaya/go-echo-jwt-noteapp-api/controller"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, tc controller.INoteController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// FE_URL: Frontendの本番環境のドメイン
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	// CSRFの設定
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		// CookieSameSite: http.SameSiteNoneMode,
		// SameSiteNoneModeを指定すると、自動的に secure が true になるため、
		// Postmanなどでテストする場合は、SameSiteDefaultModeを指定する
		CookieSameSite: http.SameSiteNoneMode,

		//CookieMaxAge:   60,
	}))

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

	t := e.Group("/notes")

	// /notes 以下のエンドポイントにアクセスする際にJWTの認証を行う
	// Use()メソッドでミドルウェアを登録
	// echojwt.WithConfig()でJWTの設定を行う
	t.Use(echojwt.WithConfig(echojwt.Config{
		// JWT の署名に使用される秘密鍵を設定
		SigningKey: []byte(os.Getenv("SECRET")),
		// JWT トークンがどこから取得されるかを指定
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllNotes)
	t.GET("/:noteId", tc.GetNoteById)
	t.POST("", tc.CreateNote)
	t.PUT("/:noteId", tc.UpdateNote)
	t.DELETE("/:noteId", tc.DeleteNote)

	return e
}
