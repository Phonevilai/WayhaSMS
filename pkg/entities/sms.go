package entities

type SMS struct {
	Trans_ID string `json:"ticket_id" binding:"required"`
	MsisDN   string `json:"phone_no" binding:"required"`
	Message  string `json:"message" binding:"required"`
}
