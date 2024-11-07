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
	//
	//  HTTP Ports
	HTTP_PORT_SECURE     = 8443
	HTTP_PORT_NON_SECURE = 8080
	//
	// HTTP Methods
	HTTP_POST = "POST"
	HTTP_GET  = "GET"
	//
	//  HTTP Status codes
	HTTP_STATUS_200 = 200 // Success
	HTTP_STATUS_400 = 400 // Client Error
	HTTP_STATUS_401 = 401 // Unauthorized to access page
	HTTP_STATUS_404 = 404 // Page not found
	HTTP_STATUS_500 = 500 // Server Error
)
