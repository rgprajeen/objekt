package main

import (
	"fmt"

	"github.com/pulumi/pulumi-oci/sdk/go/oci/identity"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func setupOCI(ctx *pulumi.Context) error {
	conf := config.New(ctx, "objekt")
	ociConf := config.New(ctx, "oci")
	tenancyOcid := pulumi.StringPtr(ociConf.Get("tenancyOcid"))

	cmptName := "objekt-" + ctx.Stack() + "-cmpt"
	cmpt, err := identity.NewCompartment(ctx, "objekt-oci-compartment", &identity.CompartmentArgs{
		Name:          pulumi.StringPtr(cmptName),
		Description:   pulumi.String("Compartment that owns resources used by Objekt"),
		CompartmentId: pulumi.StringPtr(conf.Get("ociCompartmentOcid")),
		EnableDelete:  pulumi.BoolPtr(true),
		FreeformTags: pulumi.Map{
			"Product": pulumi.String("Objekt"),
			"Stack":   pulumi.String(ctx.Stack()),
		},
	})
	if err != nil {
		return err
	}
	ctx.Export("OCI Compartment OCID", cmpt.ID())

	user, err := identity.NewUser(ctx, "objekt-oci-user", &identity.UserArgs{
		Name:          pulumi.StringPtr("objekt-" + ctx.Stack()),
		Description:   pulumi.String("User used by objekt to access OCI resources"),
		Email:         pulumi.StringPtr(fmt.Sprintf("objekt%d@attoleap.com", conf.GetInt("ociRunVersion"))),
		CompartmentId: tenancyOcid,
		FreeformTags: pulumi.Map{
			"Product": pulumi.String("Objekt"),
			"Stack":   pulumi.String(ctx.Stack()),
		},
	})
	if err != nil {
		return err
	}
	ctx.Export("OCI User OCID", user.ID())

	apiKey, err := identity.NewApiKey(ctx, "objekt-oci-api-key", &identity.ApiKeyArgs{
		UserId:   user.ID(),
		KeyValue: pulumi.String(conf.Get("ociUserPubKey")),
	})
	if err != nil {
		return err
	}
	ctx.Export("OCI User API Key Fingerprint", apiKey.Fingerprint)

	groupName := "objekt-" + ctx.Stack() + "-admins"
	group, err := identity.NewGroup(ctx, "objekt-oci-group", &identity.GroupArgs{
		Name:          pulumi.String(groupName),
		Description:   pulumi.String("Group to grant access for Objekt admins"),
		CompartmentId: tenancyOcid,
		FreeformTags: pulumi.Map{
			"Product": pulumi.String("Objekt"),
			"Stack":   pulumi.String(ctx.Stack()),
		},
	})
	if err != nil {
		return err
	}
	ctx.Export("OCI Group OCID", group.ID())

	_, err = identity.NewUserGroupMembership(ctx, "objekt-oci-group-policy-attachment", &identity.UserGroupMembershipArgs{
		CompartmentId: tenancyOcid,
		GroupId:       group.ID(),
		UserId:        user.ID(),
	})
	if err != nil {
		return err
	}

	statements := []string{
		fmt.Sprintf("Allow group %s to manage buckets in compartment %s where all { target.bucket.name=/%s_*/ }", groupName, cmptName, "objekt"),
		fmt.Sprintf("Allow group %s to manage objects in compartment %s where all { target.bucket.name=/%s_*/ }", groupName, cmptName, "objekt"),
	}

	_, err = identity.NewPolicy(ctx, "objekt-oci-ob-policy", &identity.PolicyArgs{
		Name:          pulumi.StringPtr("objekt-" + ctx.Stack() + "-ob-policy"),
		Description:   pulumi.String("Policy statements to provide objekt admins access to ObjectStorage"),
		CompartmentId: pulumi.String(conf.Get("ociCompartmentOcid")),
		Statements:    pulumi.ToStringArray(statements),
		FreeformTags: pulumi.Map{
			"Product": pulumi.String("Objekt"),
			"Stack":   pulumi.String(ctx.Stack()),
		},
	})
	if err != nil {
		return err
	}

	return nil
}
