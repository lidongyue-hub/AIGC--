package v1

import (
	"qa/api"
	"qa/serializer"
	v1 "qa/service/v1/question"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 发表问题
func QuestionAdd(c *gin.Context) {
	var service v1.QesAddService
	user := api.CurrentUser(c)
	if err := c.ShouldBind(&service); err == nil {
		res := service.QuestionAdd(user)
		c.JSON(200, res)
	} else {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
	}
}

// 查看问题
func FindOneQuestion(c *gin.Context) {
	qid, err := strconv.Atoi(c.Param("qid"))
	if err != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
		return
	}
	user := api.CurrentUser(c)
	var uid uint
	if user == nil {
		uid = 0
	} else {
		uid = user.ID
	}
	res := v1.FindOneQuestion(uint(qid), uid)
	c.JSON(200, res)
}

// 修改问题
func EditQuestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("qid"))
	user := api.CurrentUser(c)
	var service v1.EditQuestionService
	err = c.ShouldBind(&service)
	if err != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
		return
	}
	res := service.EditQuestion(user, uint(id))
	c.JSON(200, res)
}

// 删除问题
func DeleteQuestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("qid"))
	user := api.CurrentUser(c)
	if err != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
		return
	}
	res := v1.DeleteQuestion(user, uint(id))
	c.JSON(200, res)
}

// 获取首页问题列表
func FindQuestions(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(200, serializer.ErrorResponse(serializer.CodeParamError))
		return
	}
	res := v1.FindQuestions(limit, offset)
	c.JSON(200, res)
}

// 获取问题热榜
func FindHotQuestions(c *gin.Context) {
	res := v1.FindHotQuestions()
	c.JSON(200, res)
}
