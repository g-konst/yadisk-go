package yadisk

type DiskError struct {
	Message     string `json:"message"`
	Description string `json:"description"`
	ErrorCode   string `json:"error"`

	// 413: Request entity too large
	Limit  uint   `json:"limit"`
	Reason string `json:"reason"`
}

func (e *DiskError) Error() string {
	return e.Message
}
