package manage

import "github.com/gin-gonic/gin"

type ManageAdminUserApi struct {
}

// CreateAdminUser 创建AdminUser
func (m *ManageAdminUserApi) CreateAdminUser(c *gin.Context) {

}

// UpdateAdminUserPassword 修改密码
func (m *ManageAdminUserApi) UpdateAdminUserPassword(c *gin.Context) {

}

// UpdateAdminUserName 更新用户名
func (m *ManageAdminUserApi) UpdateAdminUserName(c *gin.Context) {

}

// AdminUserProfile 用户id查询AdminUser
func (m *ManageAdminUserApi) AdminUserProfile(c *gin.Context) {

}

// AdminLogin 管理员登录
func (m *ManageAdminUserApi) AdminLogin(c *gin.Context) {

}

// AdminLoginOut 管理员退出登录
func (m *ManageAdminUserApi) AdminLoginOut(c *gin.Context) {

}

// UserList 商城注册用户列表
func (m *ManageAdminUserApi) UserList(c *gin.Context) {

}

// LockUser 用户禁用与解除禁用(0-未锁定 1-已锁定)
func (m *ManageAdminUserApi) LockUser(c *gin.Context) {

}

// UploadFile 上传图片，但是原前段项目的图片链接为服务器地址，如需要显示图片，需要修改前端指向的图片链接
func (m *ManageAdminUserApi) UploadFile(c *gin.Context) {

}
