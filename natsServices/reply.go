package sharedServices

import (
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
)

//==============================
// DK Generic
//==============================

type DKReply struct {
	Reply     []byte         `json:"reply"`
	ErrorInfo errs.ErrorInfo `json:"error,omitempty"`
}
