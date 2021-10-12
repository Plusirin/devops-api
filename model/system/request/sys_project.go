package request

// Project create structure
type CreateProject struct {
	ProjectName string `json:"projectName"`
	NickName    string `json:"nickName"`
	Description string `json:"description"`
}


// Modify projectName structure
type ChangeProjectName struct {
	ID        uint           `json:"id"`
	NewProjectName string `json:"newProjectName"`
}

// Modify project nickName structure
type ChangeProjectNickName struct {
	ID        uint           `json:"id"`
	NewNickName string `json:"newNickName"`
}

// Modify project description structure
type ChangeProjectDescription struct {
	ID        uint           `json:"id"`
	NewDescription string `json:"newDescription"`
}