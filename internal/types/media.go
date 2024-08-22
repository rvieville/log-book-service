package types

type CreateMediaPayload struct {
	Key    string `json:"key" validate:"required"`
	Bucket string `json:"bucket" validate:"required"`
	DiveID uint   `json:"diveId" validate:"required"`
}
