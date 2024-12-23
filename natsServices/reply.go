package sharedServices

import (
	"time"

	errs "github.com/sty-holdings/sharedServices/v2024/errorServices"
)

//==============================
// DK Generic
//==============================

type DKReply struct {
	Reply     []byte         `json:"reply"`
	ErrorInfo errs.ErrorInfo `json:"error,omitempty"`
}

//==============================
// Analyze Question
//==============================

type AnalyzeQuestionReply struct {
	SentenceType string `json:"sentence_type"`
	Dissect      struct {
		Subject   string `json:"subject"`
		Verb      string `json:"verb"`
		Object    string `json:"object"`
		Adverbial string `json:"adverbial"`
	} `json:"dissect"`
	PrepositionalPhrase struct {
		Preposition string `json:"preposition"`
		Object      string `json:"object"`
	} `json:"prepositional_phrase"`
	PointInTime string      `json:"point_in_time"`
	TimeRange   interface{} `json:"time_range"`
	Category    interface{} `json:"category"`
	SubCategory interface{} `json:"sub_category"`
}

//==============================
// HAL
//==============================

type GetMyAnswerReply struct {
	Answer string `json:"answer"`
}

type GetPersonalAccessTokenReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Created time.Time `json:"created"`
		Id      string    `json:"id"`
		Name    string    `json:"name"`
	} `json:"response"`
}

//==============================
// DK System
//==============================

type GetSystemReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		ConnectionType          string    `json:"connection_type"`
		Created                 time.Time `json:"created"`
		HasManagedOperator      bool      `json:"has_managed_operator"`
		HasManagedSystemAccount bool      `json:"has_managed_system_account"`
		HostSystemId            string    `json:"host_system_id"`
		Id                      string    `json:"id"`
		IsTenant                bool      `json:"is_tenant"`
		JetstreamDomain         string    `json:"jetstream_domain"`
		JetstreamEnabled        bool      `json:"jetstream_enabled"`
		JetstreamTiers          []string  `json:"jetstream_tiers"`
		Name                    string    `json:"name"`
		State                   string    `json:"state"`
		Team                    struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"team"`
		UserJwtExpiresInSecs int `json:"user_jwt_expires_in_secs"`
	} `json:"response"`
}

type GetSystemLimitsReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Allocated struct {
			NumAccounts    int `json:"num_accounts"`
			ResourceLimits struct {
				Conn               int  `json:"conn"`
				Consumer           int  `json:"consumer"`
				Data               int  `json:"data"`
				DisallowBearer     bool `json:"disallow_bearer"`
				DiskMaxStreamBytes int  `json:"disk_max_stream_bytes"`
				DiskStorage        int  `json:"disk_storage"`
				Exports            int  `json:"exports"`
				Imports            int  `json:"imports"`
				Leaf               int  `json:"leaf"`
				MaxAckPending      int  `json:"max_ack_pending"`
				MaxBytesRequired   bool `json:"max_bytes_required"`
				MemMaxStreamBytes  int  `json:"mem_max_stream_bytes"`
				MemStorage         int  `json:"mem_storage"`
				Payload            int  `json:"payload"`
				Streams            int  `json:"streams"`
				Subs               int  `json:"subs"`
				TieredLimits       struct {
					R1 struct {
						Consumer           int   `json:"consumer"`
						DiskMaxStreamBytes int64 `json:"disk_max_stream_bytes"`
						DiskStorage        int64 `json:"disk_storage"`
						MaxAckPending      int   `json:"max_ack_pending"`
						MaxBytesRequired   bool  `json:"max_bytes_required"`
						MemMaxStreamBytes  int   `json:"mem_max_stream_bytes"`
						MemStorage         int   `json:"mem_storage"`
						Streams            int   `json:"streams"`
					} `json:"R1"`
				} `json:"tiered_limits"`
				Wildcards bool `json:"wildcards"`
			} `json:"resource_limits"`
		} `json:"allocated"`
		Available struct {
			NumAccounts    int `json:"num_accounts"`
			ResourceLimits struct {
				Conn               int  `json:"conn"`
				Consumer           int  `json:"consumer"`
				Data               int  `json:"data"`
				DisallowBearer     bool `json:"disallow_bearer"`
				DiskMaxStreamBytes int  `json:"disk_max_stream_bytes"`
				DiskStorage        int  `json:"disk_storage"`
				Exports            int  `json:"exports"`
				Imports            int  `json:"imports"`
				Leaf               int  `json:"leaf"`
				MaxAckPending      int  `json:"max_ack_pending"`
				MaxBytesRequired   bool `json:"max_bytes_required"`
				MemMaxStreamBytes  int  `json:"mem_max_stream_bytes"`
				MemStorage         int  `json:"mem_storage"`
				Payload            int  `json:"payload"`
				Streams            int  `json:"streams"`
				Subs               int  `json:"subs"`
				TieredLimits       struct {
					R1 struct {
						Consumer           int   `json:"consumer"`
						DiskMaxStreamBytes int64 `json:"disk_max_stream_bytes"`
						DiskStorage        int   `json:"disk_storage"`
						MaxAckPending      int   `json:"max_ack_pending"`
						MaxBytesRequired   bool  `json:"max_bytes_required"`
						MemMaxStreamBytes  int   `json:"mem_max_stream_bytes"`
						MemStorage         int   `json:"mem_storage"`
						Streams            int   `json:"streams"`
					} `json:"R1"`
				} `json:"tiered_limits"`
				Wildcards bool `json:"wildcards"`
			} `json:"resource_limits"`
		} `json:"available"`
		Total struct {
			NumAccounts    int `json:"num_accounts"`
			ResourceLimits struct {
				Conn               int  `json:"conn"`
				Consumer           int  `json:"consumer"`
				Data               int  `json:"data"`
				DisallowBearer     bool `json:"disallow_bearer"`
				DiskMaxStreamBytes int  `json:"disk_max_stream_bytes"`
				DiskStorage        int  `json:"disk_storage"`
				Exports            int  `json:"exports"`
				Imports            int  `json:"imports"`
				Leaf               int  `json:"leaf"`
				MaxAckPending      int  `json:"max_ack_pending"`
				MaxBytesRequired   bool `json:"max_bytes_required"`
				MemMaxStreamBytes  int  `json:"mem_max_stream_bytes"`
				MemStorage         int  `json:"mem_storage"`
				Payload            int  `json:"payload"`
				Streams            int  `json:"streams"`
				Subs               int  `json:"subs"`
				TieredLimits       struct {
					R1 struct {
						Consumer           int   `json:"consumer"`
						DiskMaxStreamBytes int64 `json:"disk_max_stream_bytes"`
						DiskStorage        int64 `json:"disk_storage"`
						MaxAckPending      int   `json:"max_ack_pending"`
						MaxBytesRequired   bool  `json:"max_bytes_required"`
						MemMaxStreamBytes  int   `json:"mem_max_stream_bytes"`
						MemStorage         int   `json:"mem_storage"`
						Streams            int   `json:"streams"`
					} `json:"R1"`
				} `json:"tiered_limits"`
				Wildcards bool `json:"wildcards"`
			} `json:"resource_limits"`
		} `json:"total"`
	} `json:"response"`
}

type GetTeamReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Created time.Time `json:"created"`
		Id      string    `json:"id"`
		Limits  struct {
			DisablePrometheusFederation bool `json:"disable_prometheus_federation"`
			NumUsers                    int  `json:"num_users"`
		} `json:"limits"`
		Name string `json:"name"`
	} `json:"response"`
}

type GetTeamLimitsReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Allocated struct {
			DisablePrometheusFederation bool `json:"disable_prometheus_federation"`
			NumUsers                    int  `json:"num_users"`
		} `json:"allocated"`
		Available struct {
			DisablePrometheusFederation bool `json:"disable_prometheus_federation"`
			NumUsers                    int  `json:"num_users"`
		} `json:"available"`
		Total struct {
			DisablePrometheusFederation bool `json:"disable_prometheus_federation"`
			NumUsers                    int  `json:"num_users"`
		} `json:"total"`
	} `json:"response"`
}

type GetVersionReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Commit  string    `json:"commit"`
		Date    time.Time `json:"date"`
		Version string    `json:"version"`
	} `json:"response"`
}

type ListAccountsReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Items []struct {
			AccountPublicKey string `json:"account_public_key"`
			Claims           struct {
				Iat  int    `json:"iat"`
				Iss  string `json:"iss"`
				Jti  string `json:"jti"`
				Name string `json:"name"`
				Nats struct {
					Authorization struct {
						AuthUsers interface{} `json:"auth_users"`
					} `json:"authorization"`
					DefaultPermissions struct {
						Pub struct {
						} `json:"pub"`
						Sub struct {
						} `json:"sub"`
					} `json:"default_permissions"`
					Description string `json:"description,omitempty"`
					Imports     []struct {
						Account      string `json:"account"`
						LocalSubject string `json:"local_subject"`
						Name         string `json:"name"`
						Subject      string `json:"subject"`
						Type         string `json:"type"`
					} `json:"imports"`
					Limits struct {
						Conn         int `json:"conn"`
						Data         int `json:"data"`
						Exports      int `json:"exports"`
						Imports      int `json:"imports"`
						Leaf         int `json:"leaf"`
						Payload      int `json:"payload"`
						Subs         int `json:"subs"`
						TieredLimits struct {
							R1 struct {
								Consumer           int   `json:"consumer"`
								DiskMaxStreamBytes int64 `json:"disk_max_stream_bytes"`
								DiskStorage        int64 `json:"disk_storage"`
								MaxBytesRequired   bool  `json:"max_bytes_required"`
								Streams            int   `json:"streams"`
							} `json:"R1"`
						} `json:"tiered_limits"`
						Wildcards bool `json:"wildcards"`
					} `json:"limits"`
					SigningKeys []string `json:"signing_keys"`
					Type        string   `json:"type"`
					Version     int      `json:"version"`
				} `json:"nats"`
				Sub string `json:"sub"`
			} `json:"claims"`
			ClaimsInfo struct {
				JsEnabled bool `json:"js_enabled"`
			} `json:"claims_info"`
			Created         time.Time `json:"created"`
			Id              string    `json:"id"`
			IsScpAccount    bool      `json:"is_scp_account"`
			IsSystemAccount bool      `json:"is_system_account"`
			Jwt             string    `json:"jwt"`
			JwtSettings     struct {
				Authorization struct {
					AuthUsers interface{} `json:"auth_users"`
				} `json:"authorization"`
				Description string `json:"description,omitempty"`
				Limits      struct {
					Conn         int `json:"conn"`
					Data         int `json:"data"`
					Exports      int `json:"exports"`
					Imports      int `json:"imports"`
					Leaf         int `json:"leaf"`
					Payload      int `json:"payload"`
					Subs         int `json:"subs"`
					TieredLimits struct {
						R1 struct {
							Consumer           int   `json:"consumer"`
							DiskMaxStreamBytes int64 `json:"disk_max_stream_bytes"`
							DiskStorage        int64 `json:"disk_storage"`
							MaxBytesRequired   bool  `json:"max_bytes_required"`
							Streams            int   `json:"streams"`
						} `json:"R1"`
					} `json:"tiered_limits"`
					Wildcards bool `json:"wildcards"`
				} `json:"limits"`
			} `json:"jwt_settings"`
			JwtSyncError  string `json:"jwt_sync_error"`
			JwtSyncStatus string `json:"jwt_sync_status"`
			Name          string `json:"name"`
			NscManaged    bool   `json:"nsc_managed"`
			System        struct {
				ConnectionType       string   `json:"connection_type"`
				HttpGatewayUrl       string   `json:"http_gateway_url"`
				Id                   string   `json:"id"`
				IsTenant             bool     `json:"is_tenant"`
				JetstreamDomain      string   `json:"jetstream_domain"`
				JetstreamEnabled     bool     `json:"jetstream_enabled"`
				JetstreamTiers       []string `json:"jetstream_tiers"`
				Name                 string   `json:"name"`
				ServerUrls           string   `json:"server_urls"`
				State                string   `json:"state"`
				UserJwtExpiresInSecs int      `json:"user_jwt_expires_in_secs"`
			} `json:"system"`
			Team struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"team"`
			UserJwtExpiresInSecs int `json:"user_jwt_expires_in_secs"`
		} `json:"items"`
	} `json:"response"`
}

type ListInfoAppUsersTeamReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Items []struct {
			AppUser struct {
				Id         string `json:"id"`
				Identifier string `json:"identifier"`
				Name       string `json:"name"`
			} `json:"app_user"`
			Created           time.Time `json:"created"`
			Id                string    `json:"id"`
			PendingInvitation bool      `json:"pending_invitation"`
			RoleId            string    `json:"role_id"`
			RoleName          string    `json:"role_name"`
			Team              struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"team"`
		} `json:"items"`
	} `json:"response"`
}

type ListNATSUsersReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Items []struct {
			Account struct {
				AccountPublicKey     string `json:"account_public_key"`
				Id                   string `json:"id"`
				IsScpAccount         bool   `json:"is_scp_account"`
				IsSystemAccount      bool   `json:"is_system_account"`
				Name                 string `json:"name"`
				UserJwtExpiresInSecs int    `json:"user_jwt_expires_in_secs"`
			} `json:"account"`
			Claims struct {
				Iat  int    `json:"iat"`
				Iss  string `json:"iss"`
				Jti  string `json:"jti"`
				Name string `json:"name"`
				Nats struct {
					Data          int    `json:"data"`
					IssuerAccount string `json:"issuer_account"`
					Payload       int    `json:"payload"`
					Pub           struct {
					} `json:"pub"`
					Sub struct {
					} `json:"sub"`
					Subs    int    `json:"subs"`
					Type    string `json:"type"`
					Version int    `json:"version"`
				} `json:"nats"`
				Sub string `json:"sub"`
			} `json:"claims"`
			Created          time.Time `json:"created"`
			Id               string    `json:"id"`
			Jwt              string    `json:"jwt"`
			JwtExpiresAtMax  int       `json:"jwt_expires_at_max"`
			JwtExpiresInSecs int       `json:"jwt_expires_in_secs"`
			JwtSettings      struct {
				Data    int `json:"data"`
				Payload int `json:"payload"`
				Pub     struct {
				} `json:"pub"`
				Sub struct {
				} `json:"sub"`
				Subs int `json:"subs"`
			} `json:"jwt_settings"`
			Name      string `json:"name"`
			SkGroupId string `json:"sk_group_id"`
			System    struct {
				ConnectionType       string   `json:"connection_type"`
				HttpGatewayUrl       string   `json:"http_gateway_url"`
				Id                   string   `json:"id"`
				IsTenant             bool     `json:"is_tenant"`
				JetstreamDomain      string   `json:"jetstream_domain"`
				JetstreamEnabled     bool     `json:"jetstream_enabled"`
				JetstreamTiers       []string `json:"jetstream_tiers"`
				Name                 string   `json:"name"`
				ServerUrls           string   `json:"server_urls"`
				State                string   `json:"state"`
				UserJwtExpiresInSecs int      `json:"user_jwt_expires_in_secs"`
			} `json:"system"`
			Team struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"team"`
			UserPublicKey string `json:"user_public_key"`
		} `json:"items"`
	} `json:"response"`
}

type ListPersonalAccessTokensReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Items []struct {
			Created    time.Time `json:"created"`
			Id         string    `json:"id"`
			Name       string    `json:"name"`
			Expires    time.Time `json:"expires,omitempty"`
			LastActive time.Time `json:"last_active,omitempty"`
		} `json:"items"`
	} `json:"response"`
}

type ListSystemAccountInfoReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Items []struct {
			AccountPublicKey string `json:"account_public_key"`
			Id               string `json:"id"`
			IsScpAccount     bool   `json:"is_scp_account"`
			IsSystemAccount  bool   `json:"is_system_account"`
			Name             string `json:"name"`
		} `json:"items"`
	} `json:"response"`
}

type ListSystemServerInfoReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Items []struct {
			Cluster   string    `json:"cluster"`
			Domain    string    `json:"domain"`
			Host      string    `json:"host"`
			Id        string    `json:"id"`
			Jetstream bool      `json:"jetstream"`
			Name      string    `json:"name"`
			Seq       int       `json:"seq"`
			Tags      []string  `json:"tags"`
			Time      time.Time `json:"time"`
			Ver       string    `json:"ver"`
		} `json:"items"`
	} `json:"response"`
}

type ListSystemsReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Items []struct {
			ConnectionType          string    `json:"connection_type"`
			Created                 time.Time `json:"created"`
			HasManagedOperator      bool      `json:"has_managed_operator"`
			HasManagedSystemAccount bool      `json:"has_managed_system_account"`
			HostSystemId            string    `json:"host_system_id"`
			Id                      string    `json:"id"`
			IsTenant                bool      `json:"is_tenant"`
			JetstreamDomain         string    `json:"jetstream_domain"`
			JetstreamEnabled        bool      `json:"jetstream_enabled"`
			JetstreamTiers          []string  `json:"jetstream_tiers"`
			Name                    string    `json:"name"`
			State                   string    `json:"state"`
			Team                    struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"team"`
			UserJwtExpiresInSecs int `json:"user_jwt_expires_in_secs"`
		} `json:"items"`
	} `json:"response"`
}

type ListTeamsReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Items []struct {
			Created time.Time `json:"created"`
			Id      string    `json:"id"`
			Limits  struct {
				DisablePrometheusFederation bool `json:"disable_prometheus_federation"`
				NumUsers                    int  `json:"num_users"`
			} `json:"limits"`
			Name string `json:"name"`
		} `json:"items"`
	} `json:"response"`
}

type ListTeamServerAccountsReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
}
