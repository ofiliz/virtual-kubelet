package aci

import (
	"fmt"
	"net/http"

	azure "github.com/virtual-kubelet/virtual-kubelet/providers/azure/client"
)

const (
	// BaseURI is the default URI used for compute services.
	BaseURI    = "https://management.azure.com"
	userAgent  = "virtual-kubelet/azure-arm-aci/2018-02-01"
	apiVersion = "2018-02-01-preview"

	containerGroupURLPath                    = "subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.ContainerInstance/containerGroups/{{.containerGroupName}}"
	containerGroupListURLPath                = "subscriptions/{{.subscriptionId}}/providers/Microsoft.ContainerInstance/containerGroups"
	containerGroupListByResourceGroupURLPath = "subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.ContainerInstance/containerGroups"
	containerLogsURLPath                     = containerGroupURLPath + "/containers/{{.containerName}}/logs"
)

// Client is a client for interacting with Azure Container Instances.
//
// Clients should be reused instead of created as needed.
// The methods of Client are safe for concurrent use by multiple goroutines.
type Client struct {
	hc   *http.Client
	auth *azure.Authentication
}

// NewClient creates a new Azure Container Instances client.
func NewClient(auth *azure.Authentication) (*Client, error) {
	if auth == nil {
		return nil, fmt.Errorf("Authentication is not supplied for the Azure client")
	}

	client, err := azure.NewClient(auth, BaseURI, userAgent)
	if err != nil {
		return nil, fmt.Errorf("Creating Azure client failed: %v", err)
	}

	return &Client{hc: client.HTTPClient, auth: auth}, nil
}
