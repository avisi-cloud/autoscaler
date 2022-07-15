// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package fsx_test

import (
	"fmt"
	"strings"
	"time"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/awserr"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/session"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/fsx"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To copy a backup
//
// This operation copies an Amazon FSx backup.
func ExampleFSx_CopyBackup_shared00() {
	svc := fsx.New(session.New())
	input := &fsx.CopyBackupInput{
		SourceBackupId: aws.String("backup-03e3c82e0183b7b6b"),
		SourceRegion:   aws.String("us-east-2"),
	}

	result, err := svc.CopyBackup(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeBackupNotFound:
				fmt.Println(fsx.ErrCodeBackupNotFound, aerr.Error())
			case fsx.ErrCodeServiceLimitExceeded:
				fmt.Println(fsx.ErrCodeServiceLimitExceeded, aerr.Error())
			case fsx.ErrCodeUnsupportedOperation:
				fmt.Println(fsx.ErrCodeUnsupportedOperation, aerr.Error())
			case fsx.ErrCodeIncompatibleParameterError:
				fmt.Println(fsx.ErrCodeIncompatibleParameterError, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			case fsx.ErrCodeInvalidSourceKmsKey:
				fmt.Println(fsx.ErrCodeInvalidSourceKmsKey, aerr.Error())
			case fsx.ErrCodeInvalidDestinationKmsKey:
				fmt.Println(fsx.ErrCodeInvalidDestinationKmsKey, aerr.Error())
			case fsx.ErrCodeInvalidRegion:
				fmt.Println(fsx.ErrCodeInvalidRegion, aerr.Error())
			case fsx.ErrCodeSourceBackupUnavailable:
				fmt.Println(fsx.ErrCodeSourceBackupUnavailable, aerr.Error())
			case fsx.ErrCodeIncompatibleRegionForMultiAZ:
				fmt.Println(fsx.ErrCodeIncompatibleRegionForMultiAZ, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a new backup
//
// This operation creates a new backup.
func ExampleFSx_CreateBackup_shared00() {
	svc := fsx.New(session.New())
	input := &fsx.CreateBackupInput{
		FileSystemId: aws.String("fs-0498eed5fe91001ec"),
		Tags: []*fsx.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String("MyBackup"),
			},
		},
	}

	result, err := svc.CreateBackup(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeUnsupportedOperation:
				fmt.Println(fsx.ErrCodeUnsupportedOperation, aerr.Error())
			case fsx.ErrCodeFileSystemNotFound:
				fmt.Println(fsx.ErrCodeFileSystemNotFound, aerr.Error())
			case fsx.ErrCodeVolumeNotFound:
				fmt.Println(fsx.ErrCodeVolumeNotFound, aerr.Error())
			case fsx.ErrCodeBackupInProgress:
				fmt.Println(fsx.ErrCodeBackupInProgress, aerr.Error())
			case fsx.ErrCodeIncompatibleParameterError:
				fmt.Println(fsx.ErrCodeIncompatibleParameterError, aerr.Error())
			case fsx.ErrCodeServiceLimitExceeded:
				fmt.Println(fsx.ErrCodeServiceLimitExceeded, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a new file system
//
// This operation creates a new Amazon FSx for Windows File Server file system.
func ExampleFSx_CreateFileSystem_shared00() {
	svc := fsx.New(session.New())
	input := &fsx.CreateFileSystemInput{
		ClientRequestToken: aws.String("a8ca07e4-61ec-4399-99f4-19853801bcd5"),
		FileSystemType:     aws.String("WINDOWS"),
		KmsKeyId:           aws.String("arn:aws:kms:us-east-1:012345678912:key/1111abcd-2222-3333-4444-55556666eeff"),
		SecurityGroupIds: []*string{
			aws.String("sg-edcd9784"),
		},
		StorageCapacity: aws.Int64(3200),
		StorageType:     aws.String("HDD"),
		SubnetIds: []*string{
			aws.String("subnet-1234abcd"),
		},
		Tags: []*fsx.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String("MyFileSystem"),
			},
		},
		WindowsConfiguration: &fsx.CreateFileSystemWindowsConfiguration{
			ActiveDirectoryId: aws.String("d-1234abcd12"),
			Aliases: []*string{
				aws.String("accounting.corp.example.com"),
			},
			AutomaticBackupRetentionDays:  aws.Int64(30),
			DailyAutomaticBackupStartTime: aws.String("05:00"),
			ThroughputCapacity:            aws.Int64(32),
			WeeklyMaintenanceStartTime:    aws.String("1:05:00"),
		},
	}

	result, err := svc.CreateFileSystem(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeActiveDirectoryError:
				fmt.Println(fsx.ErrCodeActiveDirectoryError, aerr.Error())
			case fsx.ErrCodeIncompatibleParameterError:
				fmt.Println(fsx.ErrCodeIncompatibleParameterError, aerr.Error())
			case fsx.ErrCodeInvalidImportPath:
				fmt.Println(fsx.ErrCodeInvalidImportPath, aerr.Error())
			case fsx.ErrCodeInvalidExportPath:
				fmt.Println(fsx.ErrCodeInvalidExportPath, aerr.Error())
			case fsx.ErrCodeInvalidNetworkSettings:
				fmt.Println(fsx.ErrCodeInvalidNetworkSettings, aerr.Error())
			case fsx.ErrCodeInvalidPerUnitStorageThroughput:
				fmt.Println(fsx.ErrCodeInvalidPerUnitStorageThroughput, aerr.Error())
			case fsx.ErrCodeServiceLimitExceeded:
				fmt.Println(fsx.ErrCodeServiceLimitExceeded, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			case fsx.ErrCodeMissingFileSystemConfiguration:
				fmt.Println(fsx.ErrCodeMissingFileSystemConfiguration, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a new file system from backup
//
// This operation creates a new file system from backup.
func ExampleFSx_CreateFileSystemFromBackup_shared00() {
	svc := fsx.New(session.New())
	input := &fsx.CreateFileSystemFromBackupInput{
		BackupId:           aws.String("backup-03e3c82e0183b7b6b"),
		ClientRequestToken: aws.String("f4c94ed7-238d-4c46-93db-48cd62ec33b7"),
		SecurityGroupIds: []*string{
			aws.String("sg-edcd9784"),
		},
		SubnetIds: []*string{
			aws.String("subnet-1234abcd"),
		},
		Tags: []*fsx.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String("MyFileSystem"),
			},
		},
		WindowsConfiguration: &fsx.CreateFileSystemWindowsConfiguration{
			ThroughputCapacity: aws.Int64(8),
		},
	}

	result, err := svc.CreateFileSystemFromBackup(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeActiveDirectoryError:
				fmt.Println(fsx.ErrCodeActiveDirectoryError, aerr.Error())
			case fsx.ErrCodeIncompatibleParameterError:
				fmt.Println(fsx.ErrCodeIncompatibleParameterError, aerr.Error())
			case fsx.ErrCodeInvalidNetworkSettings:
				fmt.Println(fsx.ErrCodeInvalidNetworkSettings, aerr.Error())
			case fsx.ErrCodeInvalidPerUnitStorageThroughput:
				fmt.Println(fsx.ErrCodeInvalidPerUnitStorageThroughput, aerr.Error())
			case fsx.ErrCodeServiceLimitExceeded:
				fmt.Println(fsx.ErrCodeServiceLimitExceeded, aerr.Error())
			case fsx.ErrCodeBackupNotFound:
				fmt.Println(fsx.ErrCodeBackupNotFound, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			case fsx.ErrCodeMissingFileSystemConfiguration:
				fmt.Println(fsx.ErrCodeMissingFileSystemConfiguration, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete a backup
//
// This operation deletes an Amazon FSx file system backup.
func ExampleFSx_DeleteBackup_shared00() {
	svc := fsx.New(session.New())
	input := &fsx.DeleteBackupInput{
		BackupId: aws.String("backup-03e3c82e0183b7b6b"),
	}

	result, err := svc.DeleteBackup(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeBackupInProgress:
				fmt.Println(fsx.ErrCodeBackupInProgress, aerr.Error())
			case fsx.ErrCodeBackupNotFound:
				fmt.Println(fsx.ErrCodeBackupNotFound, aerr.Error())
			case fsx.ErrCodeBackupRestoring:
				fmt.Println(fsx.ErrCodeBackupRestoring, aerr.Error())
			case fsx.ErrCodeIncompatibleParameterError:
				fmt.Println(fsx.ErrCodeIncompatibleParameterError, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			case fsx.ErrCodeBackupBeingCopied:
				fmt.Println(fsx.ErrCodeBackupBeingCopied, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete a file system
//
// This operation deletes an Amazon FSx file system.
func ExampleFSx_DeleteFileSystem_shared00() {
	svc := fsx.New(session.New())
	input := &fsx.DeleteFileSystemInput{
		FileSystemId: aws.String("fs-0498eed5fe91001ec"),
	}

	result, err := svc.DeleteFileSystem(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeIncompatibleParameterError:
				fmt.Println(fsx.ErrCodeIncompatibleParameterError, aerr.Error())
			case fsx.ErrCodeFileSystemNotFound:
				fmt.Println(fsx.ErrCodeFileSystemNotFound, aerr.Error())
			case fsx.ErrCodeServiceLimitExceeded:
				fmt.Println(fsx.ErrCodeServiceLimitExceeded, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe Amazon FSx backups
//
// This operation describes all of the Amazon FSx backups in an account.
func ExampleFSx_DescribeBackups_shared00() {
	svc := fsx.New(session.New())
	input := &fsx.DescribeBackupsInput{}

	result, err := svc.DescribeBackups(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeFileSystemNotFound:
				fmt.Println(fsx.ErrCodeFileSystemNotFound, aerr.Error())
			case fsx.ErrCodeVolumeNotFound:
				fmt.Println(fsx.ErrCodeVolumeNotFound, aerr.Error())
			case fsx.ErrCodeBackupNotFound:
				fmt.Println(fsx.ErrCodeBackupNotFound, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe an Amazon FSx file system
//
// This operation describes all of the Amazon FSx file systems in an account.
func ExampleFSx_DescribeFileSystems_shared00() {
	svc := fsx.New(session.New())
	input := &fsx.DescribeFileSystemsInput{}

	result, err := svc.DescribeFileSystems(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeFileSystemNotFound:
				fmt.Println(fsx.ErrCodeFileSystemNotFound, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To list tags for a resource
//
// This operation lists tags for an Amazon FSx resource.
func ExampleFSx_ListTagsForResource_shared00() {
	svc := fsx.New(session.New())
	input := &fsx.ListTagsForResourceInput{
		ResourceARN: aws.String("arn:aws:fsx:us-east-1:012345678912:file-system/fs-0498eed5fe91001ec"),
	}

	result, err := svc.ListTagsForResource(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			case fsx.ErrCodeResourceNotFound:
				fmt.Println(fsx.ErrCodeResourceNotFound, aerr.Error())
			case fsx.ErrCodeNotServiceResourceError:
				fmt.Println(fsx.ErrCodeNotServiceResourceError, aerr.Error())
			case fsx.ErrCodeResourceDoesNotSupportTagging:
				fmt.Println(fsx.ErrCodeResourceDoesNotSupportTagging, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To tag a resource
//
// This operation tags an Amazon FSx resource.
func ExampleFSx_TagResource_shared00() {
	svc := fsx.New(session.New())
	input := &fsx.TagResourceInput{
		ResourceARN: aws.String("arn:aws:fsx:us-east-1:012345678912:file-system/fs-0498eed5fe91001ec"),
		Tags: []*fsx.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String("MyFileSystem"),
			},
		},
	}

	result, err := svc.TagResource(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			case fsx.ErrCodeResourceNotFound:
				fmt.Println(fsx.ErrCodeResourceNotFound, aerr.Error())
			case fsx.ErrCodeNotServiceResourceError:
				fmt.Println(fsx.ErrCodeNotServiceResourceError, aerr.Error())
			case fsx.ErrCodeResourceDoesNotSupportTagging:
				fmt.Println(fsx.ErrCodeResourceDoesNotSupportTagging, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To untag a resource
//
// This operation untags an Amazon FSx resource.
func ExampleFSx_UntagResource_shared00() {
	svc := fsx.New(session.New())
	input := &fsx.UntagResourceInput{
		ResourceARN: aws.String("arn:aws:fsx:us-east-1:012345678912:file-system/fs-0498eed5fe91001ec"),
		TagKeys: []*string{
			aws.String("Name"),
		},
	}

	result, err := svc.UntagResource(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			case fsx.ErrCodeResourceNotFound:
				fmt.Println(fsx.ErrCodeResourceNotFound, aerr.Error())
			case fsx.ErrCodeNotServiceResourceError:
				fmt.Println(fsx.ErrCodeNotServiceResourceError, aerr.Error())
			case fsx.ErrCodeResourceDoesNotSupportTagging:
				fmt.Println(fsx.ErrCodeResourceDoesNotSupportTagging, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To update an existing file system
//
// This operation updates an existing file system.
func ExampleFSx_UpdateFileSystem_shared00() {
	svc := fsx.New(session.New())
	input := &fsx.UpdateFileSystemInput{
		FileSystemId: aws.String("fs-0498eed5fe91001ec"),
		WindowsConfiguration: &fsx.UpdateFileSystemWindowsConfiguration{
			AutomaticBackupRetentionDays:  aws.Int64(10),
			DailyAutomaticBackupStartTime: aws.String("06:00"),
			WeeklyMaintenanceStartTime:    aws.String("3:06:00"),
		},
	}

	result, err := svc.UpdateFileSystem(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeUnsupportedOperation:
				fmt.Println(fsx.ErrCodeUnsupportedOperation, aerr.Error())
			case fsx.ErrCodeIncompatibleParameterError:
				fmt.Println(fsx.ErrCodeIncompatibleParameterError, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			case fsx.ErrCodeFileSystemNotFound:
				fmt.Println(fsx.ErrCodeFileSystemNotFound, aerr.Error())
			case fsx.ErrCodeMissingFileSystemConfiguration:
				fmt.Println(fsx.ErrCodeMissingFileSystemConfiguration, aerr.Error())
			case fsx.ErrCodeServiceLimitExceeded:
				fmt.Println(fsx.ErrCodeServiceLimitExceeded, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
