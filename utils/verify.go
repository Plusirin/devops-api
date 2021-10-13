package utils

var (
	IdVerify       = Rules{"ID": {NotEmpty()}}
	ApiVerify      = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	MenuVerify     = Rules{"Path": {NotEmpty()}, "ParentId": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify = Rules{"Title": {NotEmpty()}}
	//LoginVerify            = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	LoginVerify             = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify          = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	PageInfoVerify          = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	CustomerVerify          = Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	AuthorityVerify         = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}, "ParentId": {NotEmpty()}}
	AuthorityIdVerify       = Rules{"AuthorityId": {NotEmpty()}}
	OldAuthorityVerify      = Rules{"OldAuthorityId": {NotEmpty()}}
	ChangePasswordVerify    = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	ChangeProjectNameVerify = Rules{"ID": {NotEmpty()}, "NewProjectName": {NotEmpty()}}
	ChangeToolNameVerify    = Rules{"ID": {NotEmpty()}, "NewProjectName": {NotEmpty()}}
	GetJobLastNumVerify     = Rules{"ID": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}, "JobName": {NotEmpty()}}
	SetUserAuthorityVerify  = Rules{"AuthorityId": {NotEmpty()}}
	CreateProjectVerify     = Rules{"ProjectName": {NotEmpty()}, "NickName": {NotEmpty()}, "Description": {NotEmpty()}}
	CreateToolVerify        = Rules{"ToolName": {NotEmpty()}, "URL": {NotEmpty()}, "Username": {NotEmpty()}, "Token": {NotEmpty()}, "Description": {NotEmpty()}}
)
