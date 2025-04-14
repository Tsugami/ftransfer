package protocol

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
	if c.Region == "" {
		return ErrEmptyRegion
	}
	if c.Bucket == "" {
		return ErrEmptyBucket
	}
	if c.AccessKeyID == "" {
		return ErrEmptyAccessKeyID
	}
	if c.SecretAccessKey == "" {
		return ErrEmptySecretAccessKey
	}
	return nil
}
