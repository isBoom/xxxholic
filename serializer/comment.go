package serializer

import (
	"xxxholic/model"
)

type Comment struct {
	Id         uint   `json:"id" form:"id"`
	User       User   `json:"user" json:"user"`
	ParentUser User   `json:"parentUser" form:"parentUser"`
	VideoId    uint   `json:"videoId" form:"videoId"`
	FirstId    uint   `json:"firstId" form:"firstId"`
	ParentId   uint   `json:"parentId" form:"parentId"`
	Content    string `json:"content" form:"content"`
	CreatedAt  int64  `json:"createdAt" form:"createdAt"`
}

func BuildComment(i model.Comment) Comment {
	user, _ := model.GetUser(i.UserId)
	parentUser, _ := model.GetUser(i.ParentUserId)
	return Comment{
		Id:         i.ID,
		User:       BuildUser(user),
		ParentUser: BuildUser(parentUser),
		VideoId:    i.VideoId,
		ParentId:   i.ParentId,
		Content:    i.Content,
		FirstId:    i.FirstId,
		CreatedAt:  i.CreatedAt.Unix(),
	}
}

//func BuildComments(items []model.Comment) (Comments []Comment) {
//	for _, item := range items {
//		comment := BuildComment(item)
//		Comments = append(Comments, comment)
//	}
//	return Comments
//}
