package main

//struct for Profile Update Fields only contain the contents to be Updated and directly passed to bson set
type UserUpdate struct {
	Username       string `json:",omitempty"`
	ProfilePicture string `json:",omitempty"`
	Bio            string `json:",omitempty"`
}

type FollowStruct struct {
	Username        string `json:",omitempty"`
	FollowedAccount string `json:",omitempty"`
}

//use first letter as capital to recognize as json
type User struct {
	UserName       string
	ProfilePicture string
	Followerscount int64
	Followers      []string
	Followingcount int64
	Following      []string
	Bio            string
	Posts          []Post
}

type Post struct {
	PostId       string
	PostURL      string
	LikesCount   int64
	Likes        []Like
	CommentCount int64
	Comment      []Comment
}

type Like struct {
	Username string
	Reaction int64
}

type Comment struct {
	Username    string
	CommentText string
}

type AddLikeStruct struct {
	Postid   string
	Username string
	Reaction int
}

type AddCommentStruct struct {
	Postid      string
	Username    string
	CommentText string
}
