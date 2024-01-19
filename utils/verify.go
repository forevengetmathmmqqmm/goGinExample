package utils

var (
	UserVerify       = Rules{"Nickname": {NotEmpty()}, "Avatar": {NotEmpty()}}
	EditRoleVerify   = Rules{"Title": {NotEmpty()}}
	LoginVerify      = Rules{"Nickname": {NotEmpty()}, "Password": {NotEmpty()}}
	EditUserVerify   = Rules{"ID": {NotEmpty()}, "Nickname": {NotEmpty()}, "Password": {NotEmpty()}, "Phone": {NotEmpty()}, "Email": {NotEmpty()}, "Avatar": {NotEmpty()}}
	AddUserVerify    = Rules{"Nickname": {NotEmpty()}, "Phone": {NotEmpty()}, "Password": {NotEmpty()}, "Email": {NotEmpty()}, "Avatar": {NotEmpty()}}
	AddAccessVerify  = Rules{"Path": {NotEmpty()}, "Name": {NotEmpty()}}
	EditAccessVerify = Rules{"ID": {NotEmpty()}, "Path": {NotEmpty()}, "Name": {NotEmpty()}}
)
