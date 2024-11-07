// Package constant_type_vars
/*
This file contains USA states and postal codes

RESTRICTIONS:
	- Do not edit this comment section.

NOTES:
    To improve code readability, the constant names do not follow camelCase.
	Do not remove IDE inspection directives

COPYRIGHT and WARRANTY:
	Copyright 2022
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.

*/
package sty_shared

//goland:noinspection ALL
const (
	POSTGRES_SSL_MODE_DISABLE  = "disable"
	POSTGRES_SSL_MODE_ALLOW    = "allow"
	POSTGRES_SSL_MODE_PREFER   = "prefer"
	POSTGRES_SSL_MODE_REQUIRED = "require"
	POSTGRES_CONN_STRING       = "dbname=%v user=%v password=%v host=%v port=%v connect_timeout=%v sslmode=%v pool_max_conns=%v"
)
