package v1

type TransferBackupItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type TransferResponse struct {
	Data []TransferBackupItem `json:"list"`
}

type TransferImportRequest struct {
	Backup string `json:"backup"`
}

type TransferImportResponse struct {
	Success int `json:"success"`
	Fail    int `json:"fail"`
}
