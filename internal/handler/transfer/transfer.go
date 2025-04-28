package transfer

import "maogou/khan/internal/sdk/khan"

type TransferHandler struct {
	sdk *khan.Khan
}

func NewTransferHandler(sdk *khan.Khan) *TransferHandler {
	return &TransferHandler{
		sdk: sdk,
	}
}
