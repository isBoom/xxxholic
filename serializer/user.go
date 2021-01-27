package serializer

import "xxxholic/model"

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	Email  string `json:"email"`
	UserName  string `json:"userName"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
	Signature *string `json:"signature"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		Email:  user.Email,
		UserName:  user.UserName,
		Status:    user.Status,
		Avatar:    user.AvatarUrl(),
		CreatedAt: user.CreatedAt.Unix(),
		Signature:user.Signature,
	}
}
func BuildUsers(item []model.User) (users []User) {
	for _, value := range item {
		user := BuildUser(value)
		users = append(users, user)
	}
	return
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
// BuildUsersResponse 序列化用户响应
func BuildUsersResponse(users []model.User,count uint64) Response {
	return Response{
		Data: BuildUsers(users),
		Count:count,
	}
}
