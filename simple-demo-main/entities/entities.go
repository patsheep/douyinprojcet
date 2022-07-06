package entities

import "time"

type Video struct {
	Id            int64  `form:"id,omitempty"`
	AuthorId      int64  `form:"author_id"`
	PlayUrl       string `form:"play_url"`
	CoverUrl      string `form:"cover_url,omitempty"`
	FavoriteCount int64  `form:"favorite_count,omitempty"`
	CommentCount  int64  `form:"comment_count,omitempty"`
	IsFavorite    bool   `form:"is_favorite,omitempty"`
}
type User struct {
	Id            int64  `form:"id,omitempty"`
	UserId        string `form:"user_id,omitempty"`
	Password      string `form:"password,omitempty"`
	Name          string `form:"name,omitempty"`
	FollowCount   int64  `form:"follow_count,omitempty"`
	FollowerCount int64  `form:"follower_count,omitempty"`
	IsFollow      bool   `form:"is_follow,omitempty"`
}
type Comment struct {
	ID       int64    `gorm:"comment:自增主键"`
	UserID  int64  `gorm:"type:BIGINT;not null;index:idx_user_id;评论用户ID" json:"user_id"`
	VideoID int64  `gorm:"type:BIGINT;not null;index:idx_video_id;comment:被评论视频ID" json:"video_id"`
	Content string `gorm:"type:varchar(300);not null;comment:评论内容" json:"content"`
	CreateTime time.Time `gorm:"type:timestamp;not null;default:current_timestamp()"`
	UpdateTime time.Time `gorm:"type:timestamp;not null;default:current_timestamp()"`

}