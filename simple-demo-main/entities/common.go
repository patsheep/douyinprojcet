package entities

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type Video2 struct {
	Id            int64  `form:"id,omitempty"`
	AuthorId      int64  `form:"author_id"`
	PlayUrl       string `form:"play_url"`
	CoverUrl      string `form:"cover_url,omitempty"`
	FavoriteCount int64  `form:"favorite_count,omitempty"`
	CommentCount  int64  `form:"comment_count,omitempty"`
	IsFavorite    bool   `form:"is_favorite,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}
type User2 struct {
	Id            int64  `form:"id,omitempty"`
	UserId        string `form:"user_id,omitempty"`
	Password      string `form:"password,omitempty"`
	Name          string `form:"name,omitempty"`
	FollowCount   int64  `form:"follow_count,omitempty"`
	FollowerCount int64  `form:"follower_count,omitempty"`
	IsFollow      bool   `form:"is_follow,omitempty"`
}
