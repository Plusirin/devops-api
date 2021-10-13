package request

// Tool create structure
type CreateTool struct {
	ToolName    string `json:"toolName"`
	URL         string `json:"url"`
	Username    string `json:"userName"`
	Token       string `json:"token"`
	Description string `json:"description"`
}

// Modify toolName structure
type ChangeToolName struct {
	ID          uint   `json:"id"`
	NewToolName string `json:"newToolName"`
}

// Modify tool URL structure
type ChangeURL struct {
	ID     uint   `json:"id"`
	NewURL string `json:"newURL"`
}

// Modify tool username structure
type ChangeUsername struct {
	ID          uint   `json:"id"`
	NewUsername string `json:"newUsername"`
}

// Modify tool token structure
type ChangeToken struct {
	ID       uint   `json:"id"`
	NewToken string `json:"newToken"`
}

// Modify tool description structure
type ChangeToolDescription struct {
	ID             uint   `json:"id"`
	NewDescription string `json:"newDescription"`
}

// Get job LastNum structure
type GetJobLastNum struct {
	ID      uint   `json:"id"`
	JobName string `json:"jobName"`
}
