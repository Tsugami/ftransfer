package protocol

import "github.com/Tsugami/ftransfer/internal/domain/errs"

// S3Connection represents an S3 connection
type S3Connection struct {
	Region          string
	Bucket          string
	AccessKeyID     string
	SecretAccessKey string
	Endpoint        string
	UseSSL          bool
}

// Validate performs validation on the S3 connection
func (c *S3Connection) Validate() error {
	// TODO Region not required to self-hosted minio
	// if c.Region == "" {
	// 	return ErrEmptyRegion
	// }

	if c.Bucket == "" {
		return errs.ErrEmptyBucket
	}

	if c.AccessKeyID == "" {
		return errs.ErrEmptyAccessKeyID
	}

	if c.SecretAccessKey == "" {
		return errs.ErrEmptySecretAccessKey
	}
	return nil
}

func (c *S3Connection) GetProtocol() Protocol {
	return ProtocolS3
}
