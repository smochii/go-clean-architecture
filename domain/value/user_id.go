package value

type UserId struct {
	PrimaryIdBase
}

func NewUserId() UserId {
	return UserId{PrimaryIdBase: newPrimaryIdBase()}
}

func NewUserIdFromString(id string) UserId {
	return UserId{PrimaryIdBase: newPrimaryIdBaseFromString(id)}
}
