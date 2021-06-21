package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/skos-ninja/cf-cli-auth"

	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{
		Use:  "github-releaser",
		Args: cobra.ExactArgs(0),
		RunE: runE,
	}

	clientId     = ""
	clientSecret = ""
	appDomain    = ""
)

func init() {
	cmd.Flags().StringVar(&clientId, "client-id", clientId, "")
	cmd.Flags().StringVar(&clientSecret, "client-secret", clientSecret, "")
	cmd.Flags().StringVar(&appDomain, "app-domain", appDomain, "")
}

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

func runE(cmd *cobra.Command, args []string) error {
	tr := http.DefaultTransport
	var transport cf.Transport
	if clientId != "" && clientSecret != "" {
		transport = cf.NewServiceTokenClient(tr, clientId, clientSecret)
	} else {
		var err error
		transport, err = cf.NewAccessTokenClient(cmd.Context(), tr, appDomain)
		if err != nil {
			return err
		}
	}

	client := &http.Client{Transport: transport}
	req, err := http.NewRequestWithContext(cmd.Context(), "GET", appDomain, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Print(string(body))
	return nil
}
