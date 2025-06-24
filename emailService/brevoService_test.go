// File: brevoService_test.go
package sharedServices

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	ctv "github.com/sty-holdings/sharedServices/v2025/constantsTypesVars"
	errs "github.com/sty-holdings/sharedServices/v2025/errorServices"
)

var (
	address           = "helpdesk@daveknows.ai"
	badConfigFilename = "local-email-config-bad.yaml"
	configFilename    = "/Users/syacko/workspace/sty-holdings/GriesPikeThomp/servers/connect-server/config/STARGATE/run-on-mac/local-email-config.yaml"
	htmlContent       = "<h1>Test</h1>"
	name              = "Local Dave"
	plainText         = "Test"
	subject           = "Test Email"
	toCCBCCList       = []EmailToCCBCC{
		{
			Name:    name,
			Address: address,
		},
	}
)

func TestNewBrevoServer(t *testing.T) {
	tests := []struct {
		name             string
		configFilename   string
		environment      string
		mockLoadConfig   func(string) (EmailConfig, errs.ErrorInfo)
		mockValidate     func(EmailConfig, string) errs.ErrorInfo
		expectedError    bool
		expectedErrorMsg string
	}{
		{
			name:           "ValidInput",
			configFilename: configFilename,
			environment:    "production",
			mockLoadConfig: func(string) (EmailConfig, errs.ErrorInfo) {
				return EmailConfig{
					DebugModeOn:          false,
					DefaultSenderAddress: "noreply@example.com",
					DefaultSenderName:    "Example",
					Brevo:                BrevoConfig{APIKey: "valid-api-key"},
				}, errs.ErrorInfo{}
			},
			mockValidate: func(EmailConfig, string) errs.ErrorInfo {
				return errs.ErrorInfo{}
			},
			expectedError: false,
		},
		{
			name:           "EmptyConfigFilename",
			configFilename: "",
			environment:    "production",
			mockLoadConfig: func(string) (EmailConfig, errs.ErrorInfo) {
				return EmailConfig{}, errs.ErrorInfo{}
			},
			mockValidate: func(EmailConfig, string) errs.ErrorInfo {
				return errs.ErrorInfo{}
			},
			expectedError: true,
		},
		{
			name:           "LoadConfigError",
			configFilename: badConfigFilename,
			environment:    "production",
			mockLoadConfig: func(string) (EmailConfig, errs.ErrorInfo) {
				return EmailConfig{}, errs.ErrorInfo{Error: errors.New("failed to load config")}
			},
			mockValidate: func(EmailConfig, string) errs.ErrorInfo {
				return errs.ErrorInfo{}
			},
			expectedError: true,
		},
		{
			name:           "InvalidConfig",
			configFilename: badConfigFilename,
			environment:    "production",
			mockLoadConfig: func(string) (EmailConfig, errs.ErrorInfo) {
				return EmailConfig{
					DebugModeOn:          false,
					DefaultSenderAddress: "",
					DefaultSenderName:    "",
					Brevo:                BrevoConfig{APIKey: ""},
				}, errs.ErrorInfo{}
			},
			mockValidate: func(EmailConfig, string) errs.ErrorInfo {
				return errs.ErrorInfo{Error: errors.New("invalid config")}
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				service, errInfo := NewBrevoServer(tt.configFilename, tt.environment)
				if tt.expectedError {
					assert.Nil(t, service)
					assert.NotNil(t, errInfo.Error)
				} else {
					assert.NotNil(t, service)
					assert.Nil(t, errInfo.Error)
					assert.Equal(t, "helpdesk@daveknows.ai", service.defaultSenderAddress)
					assert.Equal(t, "Local Dave", service.defaultSenderName)
					assert.NotNil(t, service.brevoClient.clientPtr)
					assert.NotNil(t, service.brevoClient.transactionalEmailsApiPtr)
				}
			},
		)
	}
}

func TestBuildEmailParams(t *testing.T) {

	tests := []struct {
		name          string
		emailParams   EmailParams
		expectedError bool
	}{
		{
			name: "ValidInput To List",
			emailParams: EmailParams{
				Attachments:    nil,
				BCCList:        nil,
				CCList:         nil,
				Sender:         EmailSender{},
				HTML:           htmlContent,
				PlainText:      plainText,
				Subject:        subject,
				TemplateID:     nil,
				TemplateParams: nil,
				ToList:         toCCBCCList,
			},
			expectedError: false,
		},
		{
			name: "ValidInput CC List",
			emailParams: EmailParams{
				Attachments:    nil,
				BCCList:        nil,
				CCList:         toCCBCCList,
				Sender:         EmailSender{},
				HTML:           htmlContent,
				PlainText:      plainText,
				Subject:        subject,
				TemplateID:     nil,
				TemplateParams: nil,
				ToList:         toCCBCCList,
			},
			expectedError: false,
		},
		{
			name: "ValidInput BCC List",
			emailParams: EmailParams{
				Attachments:    nil,
				BCCList:        toCCBCCList,
				CCList:         toCCBCCList,
				Sender:         EmailSender{},
				HTML:           htmlContent,
				PlainText:      plainText,
				Subject:        subject,
				TemplateID:     nil,
				TemplateParams: nil,
				ToList:         toCCBCCList,
			},
			expectedError: false,
		},
		{
			name: "EmptySubject",
			emailParams: EmailParams{
				Attachments:    nil,
				BCCList:        toCCBCCList,
				CCList:         toCCBCCList,
				Sender:         EmailSender{},
				HTML:           htmlContent,
				PlainText:      plainText,
				Subject:        ctv.VAL_EMPTY,
				TemplateID:     nil,
				TemplateParams: nil,
				ToList:         toCCBCCList,
			},
			expectedError: true,
		},
		{
			name: "EmptyToList",
			emailParams: EmailParams{
				Attachments:    nil,
				BCCList:        toCCBCCList,
				CCList:         toCCBCCList,
				Sender:         EmailSender{},
				HTML:           htmlContent,
				PlainText:      plainText,
				Subject:        subject,
				TemplateID:     nil,
				TemplateParams: nil,
				ToList:         nil,
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				service, _ := NewBrevoServer(configFilename, ctv.VAL_ENVIRONMENT_PRODUCTION)
				sendSMTPEmail := service.buildEmailParams(tt.emailParams)
				if tt.expectedError == false {
					assert.NotNil(t, sendSMTPEmail)
					assert.Equal(t, htmlContent, sendSMTPEmail.HtmlContent)
					assert.Equal(t, plainText, sendSMTPEmail.TextContent)
					assert.Equal(t, subject, sendSMTPEmail.Subject)
					assert.Equal(t, address, sendSMTPEmail.To[0].Email)
					assert.Equal(t, name, sendSMTPEmail.To[0].Name)
				}
			},
		)
	}
}

func TestSendEmail(t *testing.T) {

	tests := []struct {
		name          string
		emailParams   EmailParams
		expectedError bool
	}{
		{
			name: "ValidInput To List",
			emailParams: EmailParams{
				Attachments:    nil,
				BCCList:        nil,
				CCList:         nil,
				Sender:         EmailSender{},
				HTML:           htmlContent,
				PlainText:      plainText,
				Subject:        subject,
				TemplateID:     nil,
				TemplateParams: nil,
				ToList:         toCCBCCList,
			},
			expectedError: false,
		},
		{
			name: "ValidInput CC List with Template ID",
			emailParams: EmailParams{
				Attachments:    nil,
				BCCList:        nil,
				CCList:         toCCBCCList,
				Sender:         EmailSender{},
				HTML:           htmlContent,
				PlainText:      plainText,
				Subject:        subject,
				TemplateID:     4,
				TemplateParams: map[string]interface{}{"user_first_name": "Dave"},
				ToList:         toCCBCCList,
			},
			expectedError: false,
		},
		{
			name: "ValidInput BCC List",
			emailParams: EmailParams{
				Attachments:    nil,
				BCCList:        toCCBCCList,
				CCList:         toCCBCCList,
				Sender:         EmailSender{},
				HTML:           htmlContent,
				PlainText:      plainText,
				Subject:        subject,
				TemplateID:     nil,
				TemplateParams: nil,
				ToList:         toCCBCCList,
			},
			expectedError: false,
		},
		{
			name: "EmptySubject",
			emailParams: EmailParams{
				Attachments:    nil,
				BCCList:        toCCBCCList,
				CCList:         toCCBCCList,
				Sender:         EmailSender{},
				HTML:           htmlContent,
				PlainText:      plainText,
				Subject:        ctv.VAL_EMPTY,
				TemplateID:     nil,
				TemplateParams: nil,
				ToList:         toCCBCCList,
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				service, _ := NewBrevoServer(configFilename, ctv.VAL_ENVIRONMENT_PRODUCTION)
				service.SendEmail(tt.emailParams)
			},
		)
	}
}
