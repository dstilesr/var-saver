/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	cfgh "var-saver/configHandle"

	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup {configure} [options]",
	Short: "Handle backups of the application data to S3",
	Long:  `Subcommands handle various operations related to saving config backups to S3.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("Please include a valid subcommand for 'backup'")
	},
}

var backupConfigureCmd = &cobra.Command{
	Use:   "configure --bucket '<s3-bucket-name>' --prefix '<directory>' [--profile '<aws-profile-to-use>']",
	Short: "Configure the S3 backup settings",
	Long: `Configure the S3  backup settings.
This is the location on S3 where a backup of the file will be stored on request. You must provide
the S3 bucket where the file will be stored and a prefix or directory within the bucket. The backup
file will be stored at 's3://{bucket}/{prefix}/vars.toml'.

Examples

configure --bucket my-bucket --prefix "a/folder/in/bucket"

configure -b my-bucket -p "a/folder" --profile an-aws-profile
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		flags := cmd.Flags()
		bkt, _ := flags.GetString("bucket")
		prof, _ := flags.GetString("profile")
		pref, _ := flags.GetString("prefix")

		sv := cfgh.ReadConfig()
		defer sv.SaveCfg()

		sv.SetS3Config(bkt, pref, prof)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)

	// Configure subcommand setup
	backupCmd.AddCommand(backupConfigureCmd)
	backupConfigureCmd.Flags().StringP("bucket", "b", "", "S3 bucket name for backups")
	backupConfigureCmd.Flags().StringP("prefix", "p", "", "Directory/prefix within the S3 bucket")
	backupConfigureCmd.Flags().String("profile", "", "AWS profile to use (optional)")
	backupConfigureCmd.MarkFlagRequired("bucket")
	backupConfigureCmd.MarkFlagRequired("prefix")
}
