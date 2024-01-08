package utils

var (
	UserVerify     = Rules{"Nickname": {NotEmpty()}, "Avatar": {NotEmpty()}}
	EditRoleVerify = Rules{"Title": {NotEmpty()}}
	LoginVerify    = Rules{"Nickname": {NotEmpty()}, "Password": {NotEmpty()}}
	EditUserVerify = Rules{"ID": {NotEmpty()}, "Nickname": {NotEmpty()}, "Phone": {NotEmpty()}, "Email": {NotEmpty()}, "Avatar": {NotEmpty()}}
)
