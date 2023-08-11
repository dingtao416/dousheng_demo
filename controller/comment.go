package controller

import (
	"github.com/abuziming/dousheng_demo/dao"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	//token := c.Query("token")
	//获取参数数据
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	text := c.Query("comment_text")
	commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	//转换int64
	newVideoid, _ := strconv.ParseInt(videoId, 10, 64)
	userid, _ := dao.GetUserIDByVideoID(newVideoid)
	user, _ := dao.GetUserByID(c, userid)
	//添加数据
	if actionType == "1" {
		err := dao.NewComment(c, &dao.Comment{
			Id:      commentId,
			UserId:  userid,
			VideoId: newVideoid,
			User: dao.UserComment{
				Id:            user.Id,
				Name:          user.Name,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
			},
			Content:    text,
			CreateDate: time.Now().Format("2006-01-02 15:04:05 -0700 MST"),
		})
		if err != nil {
			return
		}
		//响应数据
		c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
			Comment: Comment{
				Id: commentId,
				User: User{
					Id:            user.Id,
					Name:          user.Name,
					FollowCount:   user.FollowCount,
					FollowerCount: user.FollowerCount,
				},
				Content:    text,
				CreateDate: time.Now().Format("2006-01-02 15:04:05 -0700 MST"),
			}})
		return
	}
	//删除数据
	if actionType == "2" {
		err := dao.DelComment(c, commentId, newVideoid)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	}

}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	//token := c.Query("token")
	//检验不存在，则返回 error “token有误”
	//检验过期，返回 error “token过期”
	//检验正确
	id := c.Query("video_id")

	comments, _ := dao.GetVideoComments(c, id)
	// 创建一个空的 []Comment 切片
	newComments := make([]Comment, len(comments))

	// 遍历原始切片，并复制元素到新切片
	for i, comment := range comments {
		newComments[i].Content = comment.Content
		user, _ := dao.GetUserByID(c, comment.UserId)
		newuser := User{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
		}
		newComments[i].User = newuser
		newComments[i].Id = comment.Id
		//日期格式转换 time转换为string

		newComments[i].CreateDate = comment.CreatedAt.Format("2006-01-02 15:04:05 -0700 MST")
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: newComments,
	})
}
