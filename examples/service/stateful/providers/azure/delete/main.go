package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
	"log"
)

func main() {
	// All clients require a Session. The Session provides the client with
	// shared configuration such as account and credentials.
	// A Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.
	sess := session.New()

	// Create a new instance of the service's client with a Session.
	// Optional spotinst.Config values can also be provided as variadic
	// arguments to the New function. This option allows you to provide
	// service specific configuration.
	svc := azure.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Delete stateful node.
	_, err := svc.Delete(ctx, &azure.DeleteStatefulNodeInput{
		ID: spotinst.String("ssn-01234567"),
		DeallocationConfig: &azure.DeallocationConfig{
			ShouldTerminateVM: spotinst.Bool(true),
			NetworkDeallocationConfig: &azure.ResourceDeallocationConfig{
				ShouldDeallocate: spotinst.Bool(true),
				TtlInHours:       spotinst.Int(0),
			},
			DiskDeallocationConfig: &azure.ResourceDeallocationConfig{
				ShouldDeallocate: spotinst.Bool(true),
				TtlInHours:       spotinst.Int(0),
			},
			SnapshotDeallocationConfig: &azure.ResourceDeallocationConfig{
				ShouldDeallocate: spotinst.Bool(true),
				TtlInHours:       spotinst.Int(0),
			},
			PublicIpDeallocationConfig: &azure.ResourceDeallocationConfig{
				ShouldDeallocate: spotinst.Bool(true),
				TtlInHours:       spotinst.Int(0),
			},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to delete stateful node: %v", err)
	}
}
