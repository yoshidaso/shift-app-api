package dto

type CreateShiftRequest struct {
	UserID      uint   `json:"userId" binding:"required"`
	StartTime   string `json:"startTime" binding:"required"`
	EndTime     string `json:"endTime" binding:"required"`
	WorkContent string `json:"workContent" binding:"required"`
	Issues      string `json:"issues" binding:"required"`
}
