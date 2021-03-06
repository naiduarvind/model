package main

import (
	"fmt"
	"os"

	"goa.design/model/eval"
	_ "goa.design/model/examples/basic/model" // DSL
	"goa.design/model/service"
)

// Executes the DSL and uploads the corresponding workspace to Structurizr.
func main() {
	// Run the model DSL
	w, err := eval.RunDSL()
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid model: %s", err.Error())
		os.Exit(1)
	}

	// Upload the model to the Structurizr service.
	// The API key and secret must be set in the STRUCTURIZR_KEY and
	// STRUCTURIZR_SECRET environment variables respectively. The
	// workspace ID must be set in STRUCTURIZR_WORKSPACE_ID.
	var (
		key    = os.Getenv("STRUCTURIZR_KEY")
		secret = os.Getenv("STRUCTURIZR_SECRET")
		wid    = os.Getenv("STRUCTURIZR_WORKSPACE_ID")
	)
	if key == "" || secret == "" || wid == "" {
		fmt.Fprintln(os.Stderr, "missing STRUCTURIZR_KEY, STRUCTURIZR_SECRET or STRUCTURIZR_WORKSPACE_ID environment variable.")
		os.Exit(1)
	}
	c := service.NewClient(key, secret)
	if err := c.Put(wid, w); err != nil {
		fmt.Fprintf(os.Stderr, "failed to store workspace: %s\n", err.Error())
		os.Exit(1)
	}
}
