package sharedServices

import (
	"testing"
	"time"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
)

var (
	rowValues = map[int][]any{
		1: {123, "Search", "My Campaign 1", time.Date(2023, 10, 27, 0, 0, 0, 0, time.UTC), 100, int64(10000), 1.0, 2.50, 250.00, 25.00, 50.00, 0.02, 1000.00},
		2: {456, "Display", "Product Ads", time.Date(2023, 10, 28, 0, 0, 0, 0, time.UTC), 50, int64(5000), 0.5, 1.25, 62.50, 12.50, 25.00, 0.01, 500.00},
		3: {789, "Video", "Brand Awareness", time.Date(2023, 10, 29, 0, 0, 0, 0, time.UTC), 200, int64(20000), 2.0, 0.75, 150.00, 7.50, 75.00, 0.03, 1500.00},
		// Add more rows as needed...
	}
)

func TestPostgresql(tPtr *testing.T) {
	type args struct {
		configFilename string
	}

	var (
		errorInfo errs.ErrorInfo
		gotError  bool
	)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Connecting to Postgresql Server",
			args: args{
				configFilename: "test_connection.yaml",
			},
			wantErr: false,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Bad dbName",
			args: args{
				configFilename: "test-bad-config-dbname.yaml",
			},
			wantErr: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Bad Host",
			args: args{
				configFilename: "test-bad-config-host.yaml",
			},
			wantErr: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Bad Insecure = true",
			args: args{
				configFilename: "test-bad-config-insecure-true.yaml",
			},
			wantErr: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Bad Max Conn High",
			args: args{
				configFilename: "test-bad-config-max-conn-high.yaml",
			},
			wantErr: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Bad Max Conn Zero",
			args: args{
				configFilename: "test-bad-config-max-conn-zero.yaml",
			},
			wantErr: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Bad Max Conn Zero",
			args: args{
				configFilename: "test-bad-config-max-conn-zero.yaml",
			},
			wantErr: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Bad Password",
			args: args{
				configFilename: "test-bad-config-password.yaml",
			},
			wantErr: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Bad Port",
			args: args{
				configFilename: "test-bad-config-port.yaml",
			},
			wantErr: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Bad SSL Mode",
			args: args{
				configFilename: "test-bad-config-ssl-mode.yaml",
			},
			wantErr: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Bad Timeout High",
			args: args{
				configFilename: "test-bad-config-timeout-high.yaml",
			},
			wantErr: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Bad Timeout Zero",
			args: args{
				configFilename: "test-bad-config-timeout-zero.yaml",
			},
			wantErr: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Bad CA Root",
			args: args{
				configFilename: "test-bad-config-tls-ca-root.yaml",
			},
			wantErr: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Bad Cert",
			args: args{
				configFilename: "test-bad-config-tls-cert.yaml",
			},
			wantErr: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Bad Key",
			args: args{
				configFilename: "test-bad-config-tls-key.yaml",
			},
			wantErr: true,
		},
		{
			name: ctv.TEST_NEGATIVE_SUCCESS + "Bad Username",
			args: args{
				configFilename: "test-bad-config-username.yaml",
			},
			wantErr: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				if _, errorInfo = NewPSQLServer(ts.args.configFilename); errorInfo.Error != nil {
					gotError = true
				} else {
					gotError = false
				}
				if gotError != ts.wantErr {
					tPtr.Error(ts.name)
					tPtr.Error(errorInfo)
				}
			},
		)
	}
}

func TestBatchInsert(tPtr *testing.T) {

	type arguments struct {
		batchName       string
		insertStatement string
		insertValues    map[int][]any
	}

	var (
		errorInfo       errs.ErrorInfo
		gotError        bool
		tPSQLServicePtr *PSQLService
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: ctv.TEST_POSITIVE_SUCCESS + "Batch 1.",
			arguments: arguments{
				batchName:       "Batch 1",
				insertStatement: INSERT_DAILY_PERFORMANCE,
				insertValues:    rowValues,
			},
			wantError: false,
		},
	}

	if tPSQLServicePtr, errorInfo = NewPSQLServer("/Users/syacko/workspace/sty-holdings/GriesPikeThomp/servers/connect-server/config/digits/run-on-mac/local-psql-config.yaml"); errorInfo.Error != nil {
		tPtr.Error(errorInfo.Error)
		tPtr.FailNow()
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				if errorInfo = tPSQLServicePtr.BatchInsert(ts.arguments.batchName, ts.arguments.insertStatement, ts.arguments.insertValues); errorInfo.Error != nil {
					gotError = true
				} else {
					gotError = false
				}
				if gotError != ts.wantError {
					tPtr.Error(ts.name)
					tPtr.Error(errorInfo)
				}
			},
		)
	}
}
