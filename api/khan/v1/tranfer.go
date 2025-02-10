package v1

type TransferBackupItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type TransferResponse struct {
	Data []TransferBackupItem `json:"list"`
	Tip  string               `json:"tip"`
}

type TransferImportRequest struct {
	Data []TransferBackupItem `json:"list"`
}

type TransferImportResponse struct {
	Success int `json:"success"`
	Fail    int `json:"fail"`
}
