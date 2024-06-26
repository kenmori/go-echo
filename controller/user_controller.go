package controller

import (
	"go-echo/usecase"
	"net/http"
	"os"
	"time"

	"go-echo/model"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) Login(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true                      // https通信のみでcookieを送信する
	cookie.HttpOnly = true                    // client-side scriptからcookieにアクセスできないようにする
	cookie.SameSite = http.SameSiteStrictMode // 今回はフロントとバックが違うドメイン間の通信を行うため、StrictModeを設定
	c.SetCookie(cookie)                       // httpレスポンスヘッダにcookieをセット
	return c.NoContent(http.StatusOK)         // okStatusをクライアントに返す
}

func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true                      // https通信のみでcookieを送信する
	cookie.HttpOnly = true                    // client-side scriptからcookieにアクセスできないようにする
	cookie.SameSite = http.SameSiteStrictMode // 今回はフロントとバックが違うドメイン間の通信を行うため、StrictModeを設定
	c.SetCookie(cookie)                       // httpレスポンスヘッダにcookieをセット
	return c.NoContent(http.StatusOK)
}

func (uc *userController) CsrfToken(c echo.Context) error {
	// cookie := new(http.Cookie)
	// cookie.Name = "csrfToken"
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{"csrf_token": token})
}
