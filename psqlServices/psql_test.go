package sharedServices

import (
	"testing"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
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
