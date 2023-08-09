package controller

import (
	"errors"
	"net/http"

	"github.com/abuziming/dousheng_demo/dao"
	"github.com/gin-gonic/gin"
)

const (
	maxLength = 32
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
// 这里不能删，message 的 demo 保留全局变量
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type loginResponse struct {
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type userLoginResponse struct {
	Response
	*loginResponse
}

type queryLoginProxy struct {
	username string
	password string
	data     loginResponse
}

type userResponse struct {
	Response
	User dao.User `json:"user"`
}

func RegisterHandler(c *gin.Context) {
	userHandler(c, register)
}

func LoginHandler(c *gin.Context) {
	userHandler(c, login)
}

func UserInfoHandler(c *gin.Context) {
	// user_id 为 jwt 上层解析的
	rawId, _ := c.Get("user_id")
	user := &dao.User{Id: rawId.(int64)}
	if err := dao.QueryUserInfo(user); err != nil {
		c.JSON(http.StatusOK, userLoginResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	}
	c.JSON(http.StatusOK, userResponse{
		Response: Response{StatusCode: 0},
		User:     *user,
	})
}

func userHandler(c *gin.Context, f func(string, string) (*loginResponse, error)) {
	username := c.Query("username")
	password := c.Query("password")

	response, err := f(username, password)

	if err != nil {
		c.JSON(http.StatusOK, userLoginResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	}

	c.JSON(http.StatusOK, userLoginResponse{
		Response:      Response{StatusCode: 0},
		loginResponse: response,
	})
}

func register(username, password string) (*loginResponse, error) {
	q := &queryLoginProxy{username: username, password: password}

	if err := q.check(); err != nil {
		return nil, err
	}

	if err := q.insert(); err != nil {
		return nil, err
	}
	return &q.data, nil
}

func login(username, password string) (*loginResponse, error) {
	q := &queryLoginProxy{username: username, password: password}

	if err := q.check(); err != nil {
		return nil, err
	}

	if err := q.queryLogin(); err != nil {
		return nil, err
	}
	return &q.data, nil
}

// 检验参数
func (q *queryLoginProxy) check() error {
	if len(q.username) > maxLength {
		return errors.New("用户名长度超过限制")
	}
	if len(q.password) > maxLength {
		return errors.New("密码长度超过限制")
	}
	return nil
}

// 插入数据到数据库，设置 userid，并发放 token
func (q *queryLoginProxy) insert() error {
	if err := dao.IsUserExist(q.username); err != nil {
		return err
	}

	userLogin := &dao.UserLogin{Username: q.username, Password: q.password}
	user := &dao.User{User: userLogin, Name: q.username}
	// 插入时，自动依据主键获得 id
	if err := dao.AddUser(user); err != nil {
		return err
	}

	token, err := GetToken(userLogin)
	if err != nil {
		return err
	}

	q.data.Token = token
	q.data.UserId = user.Id

	return nil
}

// 查询当前用户数据，设置 userid，发放 token
func (q *queryLoginProxy) queryLogin() error {
	userLogin := &dao.UserLogin{Username: q.username, Password: q.password}
	// 查询时获得 userid
	if err := dao.QueryUserLogin(userLogin); err != nil {
		return err
	}

	token, err := GetToken(userLogin)
	if err != nil {
		return err
	}

	q.data.Token = token
	q.data.UserId = userLogin.UserId
	return nil
}
