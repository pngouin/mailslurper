// Copyright 2013-2018 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package mailslurper

/*
ServiceSettings represents the necessary settings to connect to
and talk to the MailSlurper service tier.
*/
type ServiceSettings struct {
	AuthenticationScheme string `json:"authenticationScheme"`
	IsPublicSSL          bool   `json:"isPublicSSL"`
	ServiceAddress       string `json:"serviceAddress"`
	ServicePort          int    `json:"servicePort"`
	ServicePublicAddress string `json:"servicePublicAddress"`
	ServicePublicPort    int    `json:"servicePublicPort"`
	Version              string `json:"version"`
}
