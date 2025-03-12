package v1

type LicenseUploadRequest struct {
	File string `form:"file" binding:"required"`
}
