package dao

import (
	"gorm.io/gorm"
	"time"
)

/*
此处分表是防止用户表过大，且经过考察业务，用户名与密码仅在登录注册时用到
后续均以 token 和 Id 的形式
*/
type UserLogin struct {
	Username string `gorm:"primaryKey"` // 登录信息以用户名为主键，用来查用户名是否重复
	Password string
	UserId   int64
}
type UserComment struct {
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	Id              int64  `json:"id"`               // 用户id
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Name            string `json:"name"`             // 用户名称
	Signature       string `json:"signature"`        // 个人简介
	TotalFavorited  string `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数
}

/*
token 存在内存中，由服务器动态分发
*/
type User struct {
	Id            int64      `json:"id"`
	Name          string     `json:"name"`
	FollowCount   int64      `json:"follow_count"`                          // 社交部分保留字段，用户关注总数，现默认返回 0
	FollowerCount int64      `json:"follower_count"`                        // 社交部分保留字段，用户粉丝总数，现默认返回 0
	IsFollow      bool       `json:"is_follow" gorm:"-"`                    // 社交部分保留字段，用户是否被当前询问用户关注，现默认返回 false
	User          *UserLogin `json:"-"`                                     // 用户与账号密码之间的一对一
	Videos        []*Video   `json:"-"`                                     // 用户与投稿视频的一对多
	FavorVideos   []*Video   `json:"-" gorm:"many2many:user_favor_videos;"` // 用户与点赞视频之间的多对多
	Comments      []*Comment `json:"-"`                                     // 用户与评论的一对多
	// Follows       []*User    `json:"-" gorm:"many2many:user_relations;"`    // 社交部分保留字段，用户之间的多对多
}

type Video struct {
	Id            int64      `json:"id"`
	UserId        int64      `json:"-"`
	Author        User       `json:"author" gorm:"-"`                       // 用户与投稿视频的一对多
	PlayUrl       string     `json:"play_url"`                              // 视频文件保存的地址路由
	CoverUrl      string     `json:"cover_url"`                             // 视频封面图片保存的地址路由
	FavoriteCount int64      `json:"favorite_count"`                        // 视频的点赞总数
	CommentCount  int64      `json:"comment_count"`                         // 视频的评论总数
	IsFavorite    bool       `json:"is_favorite" gorm:"-"`                  // 视频是否被当前询问用户点赞
	Title         string     `json:"title,omitempty"`                       // 视频标题
	Users         []*User    `json:"-" gorm:"many2many:user_favor_videos;"` // 用户与点赞视频之间的多对多
	Comments      []*Comment `json:"-"`                                     // 视频与评论的一对多
	CreatedAt     int64      `json:"-" gorm:"autoCreateTime:milli"`         // 视频创建时间戳毫秒，与前端传来的数据对应，省去转化的过程，使用 CreateAt 字段，创建时间会自动写入该字段
}

type Comment struct {
	Id         int64          `json:"id"`
	UserId     int64          `json:"-"`
	VideoId    int64          `json:"-"`                    // 视频与评论的一对多
	User       UserComment    `json:"user" gorm:"-"`        // 用户与评论的一对多
	Content    string         `json:"content"`              // 评论内容
	CreatedAt  time.Time      `json:"-"`                    // 创建的时间，以这种格式存储方便转为 json 的字符串格式
	CreateDate string         `json:"create_date" gorm:"-"` // json 传回去的字符串时间形式，数据库无需存储
	DeletedAt  gorm.DeletedAt `json:"-"`                    // 评论删除时间戳，仅为软删除，保留评论内容作为分析，此处组合了 gorm.DeletedAt，直接删除就是软删除
}
