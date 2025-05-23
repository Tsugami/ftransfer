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

	// TODO Region not required to self-hosted minio
	// if c.Region == "" {
	// 	return ErrEmptyRegion
	// }

	if c.Bucket == "" {
		return ErrEmptyS3Bucket
	}

	if c.AccessKeyID == "" {
		return ErrEmptyS3AccessKeyID
	}

	if c.SecretAccessKey == "" {
		return ErrEmptyS3SecretAccessKey
	}
	return nil
}

func (c *S3Connection) GetProtocol() Protocol {
	return ProtocolS3
}

func (c *S3Connection) GetJson() map[string]interface{} {
	return map[string]interface{}{
		"protocol":          c.GetProtocol(),
		"region":            c.Region,
		"bucket":            c.Bucket,
		"access_key_id":     c.AccessKeyID,
		"secret_access_key": c.SecretAccessKey,
		"endpoint":          c.Endpoint,
		"use_ssl":           c.UseSSL,
	}
}
