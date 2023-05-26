package schema

type BusOperatorProfileRequestSchema struct {
	gorm.Model
	UserID           int             `json:"user_id"`
	User             User            `json:"user"`
	BusinessName     string          `json:"business_name"`
	OfficeAddress    string          `json:"office_address"`
	Ratings          int8            `json:"ratings"`
	ApprovalStatus   APPROVAL_STATUS `json:"approval_status"`
	RejectionComment string          `json:"rejection_comment"`
	BusinessLogoId   uint            `json:"business_logo_id"`
	BusinessLogo     Media           `json:"business_logo"`
}