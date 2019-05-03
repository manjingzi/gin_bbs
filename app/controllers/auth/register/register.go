package register

import (
	"gin_bbs/app/auth"
	"gin_bbs/app/controllers"
	userRequest "gin_bbs/app/requests/user"

	"github.com/gin-gonic/gin"
)

func ShowRegistrationForm(c *gin.Context) {
	controllers.Render(c, "auth/register", gin.H{})
}

func Register(c *gin.Context) {
	// 验证参数和创建用户
	userCreateForm := &userRequest.UserCreateForm{
		Name:                 c.PostForm("name"),
		Email:                c.PostForm("email"),
		Password:             c.PostForm("password"),
		PasswordConfirmation: c.PostForm("password_confirmation"),
	}
	ok, user := userCreateForm.ValidateAndSave(c)

	if !ok || user == nil {
		controllers.RedirectRouter(c, "register")
		return
	}

	auth.Login(c, user)
	controllers.RedirectRouter(c, "root")
}
