package user

type UserModel struct {
	Email    string `json:"email" bson:"email"`
	Salt     string `json:"salt" bson:"salt"`
	Password string `json:"password" bson:"password"`
}
