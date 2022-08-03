package grpc

import (
	"crypto/tls"
	"errors"
	"fmt"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"riasc.eu/wice/pkg/signaling"
)

type BackendConfig struct {
	signaling.BackendConfig

	Target string

	Options []grpc.DialOption
}

func (c *BackendConfig) Parse(cfg *signaling.BackendConfig) error {
	c.BackendConfig = *cfg

	options := c.URI.Query()

	isInsecure := false
	if options.Has("insecure") {
		var err error
		if isInsecure, err = strconv.ParseBool(options.Get("insecure")); err != nil {
			return fmt.Errorf("failed to parse 'insecure' option: %w", err)
		}
	}

	skipVerify := false
	if options.Has("skip_verify") {
		var err error
		if skipVerify, err = strconv.ParseBool(options.Get("skip_verify")); err != nil {
			return fmt.Errorf("failed to parse 'skip_verify' option: %w", err)
		}
	}

	var creds credentials.TransportCredentials
	if isInsecure {
		creds = insecure.NewCredentials()
	} else {
		// Use system certificate store
		cfg := &tls.Config{
			//#nosec G402 -- Users should have the freedom to disable verification for self-signed certificates
			InsecureSkipVerify: skipVerify,
		}
		creds = credentials.NewTLS(cfg)
	}

	c.Options = append(c.Options, grpc.WithTransportCredentials(creds))

	if c.URI.Host == "" {
		return errors.New("missing gRPC server url")
	}

	c.Target = c.URI.Host

	return nil
}
