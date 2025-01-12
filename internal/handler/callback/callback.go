package callback

import "smallBot/internal/sdk/gewe"

type CallbackHandler struct {
	sdk *gewe.Gewe
}

const (
	TYPENAME_SYSTEM_MESSAGES = "SystemMessages"
	TYPENAME_SYSTEM_SUCCES   = "Long_Serve_START_SUCCESS"
	TYPENAME_SYSTEM_CLOSE    = "Long_Serve_Close"
	TYPENAME_SYSTEM_FAILURE  = "Long_Serve_START_FAILURE"
	TYPENAME_SYSTEM_TIMEOUT  = "Long_Serve_START_Timeout"
	TYPENAME_MSG_ADDMSG      = "AddMsg"
	TYPENAME_MSG_DELCONTACTS = "DelContacts"
	TYPENAME_MSG_MODCONTACTS = "ModContacts"
	TYPENAME_MSG_FINDERMSG   = "FinderMsg"
	TYPENAME_MSG_OFFLINE     = "Offline"
)

func NewCallbackHandler(sdk *gewe.Gewe) *CallbackHandler {
	return &CallbackHandler{
		sdk: sdk,
	}
}
