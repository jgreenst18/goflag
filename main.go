package main

import (
	"fmt"
	"gopkg.in/launchdarkly/go-sdk-common.v1/ldvalue"
	ld "gopkg.in/launchdarkly/go-server-sdk.v4"
	"gopkg.in/launchdarkly/go-server-sdk.v4/ldlog"
	"os"
	"time"
)

// Set sdkKey to your LaunchDarkly SDK key before compiling
const sdkKey = "sdk-a78831d0-0634-4bf1-a9c4-afb4e719a9cb"

// Set featureFlagKey to the feature flag key you want to evaluate
const featureFlagKey = "new-button"

func main() {
	if sdkKey == "" {
		fmt.Println("Please edit main.go to set sdkKey to your LaunchDarkly SDK key first")
		os.Exit(1)
	}

	// The only custom configuration we are doing here is to reduce the amount of logging.
	config := ld.DefaultConfig
	config.Loggers.SetMinLevel(ldlog.Warn)

	client, err := ld.MakeCustomClient(sdkKey, config, 5*time.Second)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Set up the user properties. This user should appear on your LaunchDarkly users dashboard
	// soon after you run the demo.
	user := ld.NewUserBuilder("bob@example.com").
		FirstName("Bob").
		LastName("Loblaw").
		Custom("groups", ldvalue.ArrayBuild().Add(ldvalue.String("beta_testers")).Build()).
		Build()

	showFeature, err := client.BoolVariation(featureFlagKey, user, false)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Feature flag '%s' is %t for this user\n", featureFlagKey, showFeature)

	// Calling client.Close() ensures that the SDK shuts down cleanly before the program exits.
	// Unless you do this, the SDK may not have a chance to deliver analytics events to LaunchDarkly,
	// so the user properties and the flag usage statistics may not appear on your dashboard. In a
	// normal long-running application, events would be delivered automatically in the background
	// and you would not need to close the client.
	client.Close()
}
