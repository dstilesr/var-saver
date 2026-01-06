package confighandle

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (sv *SavedVariables) HasBackupCfg() bool {
	return sv.BackupCfg != nil && sv.BackupCfg.Bucket != "" && sv.BackupCfg.Prefix != ""
}

// SaveBackup saves the variables file to S3.
func (sv *SavedVariables) SaveBackup() error {
	if !sv.HasBackupCfg() {
		return errors.New("Backup configuration has not been set up")
	}
	c, err := GetS3Client(sv.BackupCfg.Profile)
	if err != nil {
		return err
	}

	// Compare updated times. If backup updated after local file, do not save.
	luS3, err := sv.backupLastUpdated(c)
	if err != nil {
		return err
	}
	luLocal, err := time.Parse(time.RFC3339, sv.Meta.LastUpdated)
	if err != nil {
		return err
	}
	if luS3.After(luLocal) {
		return errors.New("Backup updated after local file!")
	}

	// Save backup
	cfgPath, err := ConfigPath()
	if err != nil {
		return err
	}
	file, err := os.Open(cfgPath)
	if err != nil {
		return err
	}
	defer file.Close()

	key := sv.getS3Key()
	_, err = c.PutObject(
		context.TODO(),
		&s3.PutObjectInput{
			Key:    &key,
			Bucket: &sv.BackupCfg.Bucket,
			Body:   file,
		},
	)
	return err
}

// BackupLastUpdated Gets the time when the backup was last updated on S3.
func (sv *SavedVariables) backupLastUpdated(c *s3.Client) (time.Time, error) {
	key := sv.getS3Key()
	objInfo, err := c.HeadObject(
		context.TODO(),
		&s3.HeadObjectInput{
			Bucket: &sv.BackupCfg.Bucket,
			Key:    &key,
		},
	)
	if err != nil {
		return time.Time{}, err
	}

	return *objInfo.LastModified, nil
}

// GetS3Client returns an S3 client with the given profile. If no profile is provided,
// it loads credentials from the environment variables or the default profile.
func GetS3Client(prof string) (*s3.Client, error) {
	var err error = nil
	var cfg aws.Config
	if prof == "" {
		// Load credentials from env vars / default profile
		cfg, err = config.LoadDefaultConfig(context.TODO())
	} else {
		cfg, err = config.LoadDefaultConfig(
			context.TODO(),
			config.WithSharedConfigProfile(prof),
		)
	}
	if err != nil {
		return nil, err
	}
	return s3.NewFromConfig(cfg), nil
}

// getS3Key returns the S3 key for the saved variables.
func (sv *SavedVariables) getS3Key() string {
	return fmt.Sprintf("%s/vars.toml", sv.BackupCfg.Prefix)
}
