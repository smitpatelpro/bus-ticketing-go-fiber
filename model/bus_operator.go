package model

import "gorm.io/gorm"

type APPROVAL_STATUS string

const (
	STATUS_APPROVED         APPROVAL_STATUS = "APPROVED"
	STATUS_PENDING_APPROVAL APPROVAL_STATUS = "PENDING_APPROVAL"
	STATUS_REJECTED         APPROVAL_STATUS = "REJECTED"
)

// Model struct
type BusOperatorProfile struct {
	gorm.Model
	UserID           int             `json:"user_id"`
	User             User            `json:"user"`
	BusinessName     string          `json:"business_name"`
	OfficeAddress    string          `json:"office_address"`
	Ratings          int8            `json:"ratings"`
	ApprovalStatus   APPROVAL_STATUS `json:"approval_status"`
	RejectionComment string          `json:"rejection_comment"`
	BusinessLogoId   int             `json:"business_logo_id"`
	// BusinessLogo     Media           `json:"business_logo"`
}
