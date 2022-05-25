package entities

type SMS struct {
	Trans_ID string `json:"ticketID" binding:"required"`
	MsisDN   string `json:"phoneNo" binding:"required"`
	Message  string `json:"message" binding:"required"`
}
