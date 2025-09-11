package dto

type CreateShiftRequest struct {
	UserName    string   `json:"userName" binding:"required"`
	StartTime   string `json:"startTime" binding:"required"`
	EndTime     string `json:"endTime" binding:"required"`
	WorkContent string `json:"workContent" binding:"required"`
	Issues      string `json:"issues" binding:"required"`
}
