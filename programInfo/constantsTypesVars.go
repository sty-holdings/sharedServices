package sharedServices

import (
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
)

type ProgramInfo struct {
	ErrorInfo        errs.ErrorInfo
	FileName         string       `json:"program_filename"`
	FunctionInfo     FunctionInfo `json:"function_info"`
	GoVersion        string       `json:"go_version"`
	NumberCPUs       int          `json:"number_cpus"`
	DebugModeOn      bool         `json:"debug_mode_on"`
	WorkingDirectory string       `json:"working_directory"`
}

type FunctionInfo struct {
	FileName   string `json:"function_filename"`
	Name       string `json:"function_name"`
	LineNumber int    `json:"function_line_number"`
}
