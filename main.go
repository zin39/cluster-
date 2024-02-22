package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)
		bucket, err := s3.NewBucket(ctx, "my-bucket", nil)
		if err != nil {
			return err
		}

		// Export the name of the bucket
		ctx.Export("bucketName", bucket.ID())
		return nil
	})
}

// before this comment is the basic file that pulumi created.
// Afer this it is the file which will take required input to create 


package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/redise/enterprise"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Define the IAM role for the Redis Enterprise cluster nodes
		role, err := iam.NewRole(ctx, "redisRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(`{
				"Version": "2012-10-17",
				"Statement": [{
					"Action": "sts:AssumeRole",
					"Principal": {
						"Service": "ec2.amazonaws.com"
					},
					"Effect": "Allow",
					"Sid": ""
				}]
			}`),
		})
		if err != nil {
			return err
		}

		// Define the Redis Enterprise cluster
		redisCluster, err := enterprise.NewCluster(ctx, "redisCluster", &enterprise.ClusterArgs{
			NodeType:             pulumi.String("cache.t3.small"),
			NumReplicasPerShard:  pulumi.Int(1), 
			ShardCount:           pulumi.Int(1),
			EngineVersion:        pulumi.String("6.0.0"),
			SecurityGroups:       pulumi.StringArray{}, // Define security groups
			SubnetIds:            pulumi.StringArray{}, // Define subnets
			IamRole:              role.Arn,
			ClusterName:          pulumi.String("my-redis-cluster"),
			EnableAutomaticFailover: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		// Export outputs
		ctx.Export("clusterID", redisCluster.ID())
		ctx.Export("clusterEndpoint", redisCluster.ClusterEndpoint)

		return nil
	})
}

