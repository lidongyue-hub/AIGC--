package v1

import (
	"net/http"
	"qa/api"
	"qa/cache"
	"qa/model"
	"qa/serializer"
	v1 "qa/service/v1/user"

	"github.com/gin-gonic/gin"
	//"github.com/go-delve/delve/service"
)

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	var service v1.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
	}
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	var service v1.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login()
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := api.CurrentUser(c)
	c.JSON(http.StatusOK, serializer.OkResponse(serializer.BuildUserResponse(user)))
}

// 查看用户发布的问题
func GetUserQuestions(c *gin.Context) {
	user := api.CurrentUser(c)
	questions, err := model.GetUserQuestions(user.ID)
	if err != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeDatabaseError))
	} else {
		res := serializer.BuildUserQuestionsResponse(questions)
		c.JSON(200, serializer.OkResponse(res))
	}
}

// 查看用户发布的回答
func GetUserAnswers(c *gin.Context) {
	user := api.CurrentUser(c)
	answers, err := model.GetUserAnswers(user.ID)
	if err != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeDatabaseError))
	} else {
		res := serializer.BuildUserAnswersResponse(answers)
		c.JSON(200, serializer.OkResponse(res))
	}
}

// Logout 用户退出登录
func Logout(c *gin.Context) {
	token, _ := c.Get("token")
	tokenString := token.(string)

	cache.RedisClient.SAdd("jwt:baned", tokenString)
	c.JSON(http.StatusOK, serializer.OkResponse(nil))
}
