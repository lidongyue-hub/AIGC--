package v1

import (
	"fmt"
	"math/rand"
	"qa/model"
	"qa/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	UserName        string `form:"username" json:"username" binding:"required,min=3,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=6,max=18"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=6,max=18"`
}

// Valid 验证表单
func (service *UserRegisterService) Valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return serializer.ErrorResponse(serializer.CodePasswordConfirmError)
	}

	res := model.DB.Where("username = ?", service.UserName).First(&model.User{})
	if res.RowsAffected > 0 {
		return serializer.ErrorResponse(serializer.CodeUserExistError)
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() *serializer.Response {
	// 生成随机初始头像
	avatar := fmt.Sprintf("http://images.nowcoder.com/head/%dt.png", rand.Intn(1000))

	user := model.User{
		Username: service.UserName,
		UserProfile: model.UserProfile{
			Nickname: service.UserName,
			Avatar:   avatar,
		},
	}

	// 表单验证
	if err := service.Valid(); err != nil {
		return err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.ErrorResponse(serializer.CodeUnknownError)
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.ErrorResponse(serializer.CodeDatabaseError)
	}
	return serializer.OkResponse(serializer.BuildUserResponse(&user))
}
