package main

import (
	"github.com/slackhq/nebula/cert"
)

func parseNebulaCertificate(raw []byte) (*cert.NebulaCertificate, error) {
	parsed, _, err := cert.UnmarshalNebulaCertificateFromPEM(raw)
	if err != nil {
		return nil, err
	}
	return parsed, nil
}
