import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as awsx from "@pulumi/awsx";

const user = new aws.iam.User("objekt-aws-user", {
    name: `objekt-${pulumi.getStack()}`,
    path: `/user/objekt/${pulumi.getStack()}/`,
    tags: {
        "Product": "Objekt",
    }
})

export const accessKey = new aws.iam.AccessKey("objekt-aws-access-key", {
    user: user.name,
})

const group = new aws.iam.Group("objekt-aws-group", {
    name: `objekt-${pulumi.getStack()}-admins`,
    path: `/group/objekt/${pulumi.getStack()}/`,
})

const groupMembership = new aws.iam.GroupMembership("objekt-aws-group-membership", {
    name: `objekt-${pulumi.getStack()}-admins`,
    users: [user.name],
    group: group.name,
})

const s3Policy = new aws.iam.Policy("objekt-aws-s3-policy", {
    name: `objekt-${pulumi.getStack()}-s3-policy`,
    description: "Policy for Objekt S3 access",
    path: `/policy/objekt/${pulumi.getStack()}/`,
    policy: JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Sid: `Objekt${pulumi.getStack()}S3Policy`,
            Effect: "Allow",
            Action: [
                "s3:*"
            ],
            Resource: [
                "arn:aws:s3:::objekt_*"
            ],
        }],
    }),
})

const policyAttachment = new aws.iam.GroupPolicyAttachment("objekt-aws-policy-attachment", {
    group: group.name,
    policyArn: s3Policy.arn
})
