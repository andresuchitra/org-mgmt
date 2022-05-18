package models

type UserResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`

	TotalFollowers  uint32 `json:"total_followers"`
	TotalFollowings uint32 `json:"total_followings"`
}

func (u *User) ConvertToUserResponse() UserResponse {
	member := UserResponse{}
	member.ID = u.ID
	member.Email = u.Email
	member.Avatar = u.Avatar
	member.Name = u.Name

	return member
}

type CommentResponse struct {
	ID             uint   `json:"id"`
	Text           string `json:"name"`
	OrganizationID uint   `json:"organization_id"`
	AuthorID       uint   `json:"author_id"`
}

func (c *Comment) ConvertToCommentResponse() CommentResponse {
	cresp := CommentResponse{}
	cresp.ID = c.ID
	cresp.AuthorID = c.AuthorID
	cresp.OrganizationID = c.OrganizationID
	cresp.Text = c.Text

	return cresp
}
