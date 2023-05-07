package payload

import "time"

type CreateCampaignRequest struct {
	Name             string    `json:"name" form:"name" binding:"required"`
	ShortDescription string    `json:"short_description" form:"short_description" binding:"required"`
	Description      string    `json:"description" form:"description" binding:"required"`
	GoalAmount       int       `json:"goal_amount" form:"goal_amount" binding:"required"`
	EndDate          time.Time `json:"end_date"`
}
