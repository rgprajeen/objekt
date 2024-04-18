package main

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		group, err := iam.NewGroup(ctx, "objekt-aws-group", &iam.GroupArgs{
			Name: pulumi.String(fmt.Sprintf("objekt-%s-admins", ctx.Stack())),
			Path: pulumi.String(fmt.Sprintf("/group/objekt/%s/", ctx.Stack())),
		})
		if err != nil {
			return err
		}

		policyJson, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				{
					"Sid":    fmt.Sprintf("Objekt%sS3Policy", ctx.Stack()),
					"Effect": "Allow",
					"Action": []string{
						"s3:*",
					},
					"Resource": []string{
						"arn:aws:s3:::objekt-*",
					},
				},
			},
		})
		if err != nil {
			return err
		}

		policy, err := iam.NewPolicy(ctx, "objekt-aws-s3-policy", &iam.PolicyArgs{
			Name:        pulumi.String(fmt.Sprintf("objekt-%s-s3-policy", ctx.Stack())),
			Description: pulumi.String("Policy for Objekt S3 access"),
			Path:        pulumi.String(fmt.Sprintf("/policy/objekt/%s/", ctx.Stack())),
			Policy:      pulumi.String(policyJson),
			Tags: pulumi.StringMap{
				"Product": pulumi.String("Objekt"),
				"Stack":   pulumi.String(ctx.Stack()),
			},
		})
		if err != nil {
			return err
		}

		_, err = iam.NewGroupPolicyAttachment(ctx, "objekt-aws-group-policy-attachment", &iam.GroupPolicyAttachmentArgs{
			Group:     group.Name,
			PolicyArn: policy.Arn,
		})
		if err != nil {
			return err
		}

		user, err := iam.NewUser(ctx, "objekt-aws-user", &iam.UserArgs{
			Name: pulumi.String(fmt.Sprintf("objekt-%s", ctx.Stack())),
			Path: pulumi.String(fmt.Sprintf("/user/objekt/%s/", ctx.Stack())),
			Tags: pulumi.StringMap{
				"Product": pulumi.String("Objekt"),
				"Stack":   pulumi.String(ctx.Stack()),
			},
		})
		if err != nil {
			return err
		}

		_, err = iam.NewGroupMembership(ctx, "objekt-aws-group-user-attachment", &iam.GroupMembershipArgs{
			Name:  pulumi.String(fmt.Sprintf("objekt-%s-admins-membership", ctx.Stack())),
			Group: group.Name,
			Users: pulumi.StringArray{user.Name},
		})
		if err != nil {
			return err
		}

		accessKey, err := iam.NewAccessKey(ctx, "objekt-aws-access-key", &iam.AccessKeyArgs{
			User: user.Name,
		})
		if err != nil {
			return err
		}
		ctx.Export("AWS Access Key ID", accessKey.ID())
		ctx.Export("AWS Secret Key", accessKey.Secret)

		return nil
	})
}
