package param

type LoginParam struct {
	// "password" or "code"
	Type     string
	Username string
	Password string
	Code     int32
}

type RegisterParam struct {
	Username string
	Email    string
	Password string
}
