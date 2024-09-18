package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/1206yaya/go-echo-jwt-noteapp-api/model"
	"github.com/1206yaya/go-echo-jwt-noteapp-api/usecase"
	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	usecase usecase.IUserUsecase
}

func NewUserController(usecase usecase.IUserUsecase) IUserController {
	return &userController{usecase}
}

func (controller *userController) SignUp(c echo.Context) error {
	user := model.User{}
	// バインドしたデータをuserに格納
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := controller.usecase.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) LogIn(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.usecase.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	// フロントエンド、バックエンドのドメインが異なる、クロスドメイン間のCookieの送受信を許可する
	cookie.SameSite = http.SameSiteNoneMode
	// postmanでテストする場合は、SameSiteDefaultModeを指定する
	// cookie.SameSite = http.SameSiteDefaultMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	// Nowとすることで、すぐに失効する
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	// cookie.SameSite = http.SameSiteDefaultMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}
func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
