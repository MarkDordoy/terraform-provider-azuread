package client

import (
	"github.com/manicminer/hamilton/environments"
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	ApplicationsClient         *msgraph.ApplicationsClient
	ApplicationTemplatesClient *msgraph.ApplicationTemplatesClient
	DirectoryObjectsClient     *msgraph.DirectoryObjectsClient
}

func NewClient(o *common.ClientOptions) *Client {
	applicationsClient := msgraph.NewApplicationsClient(o.TenantID)
	o.ConfigureClient(&applicationsClient.BaseClient)

	applicationTemplatesClient := msgraph.NewApplicationTemplatesClient(o.TenantID)

	o.ConfigureClient(&applicationTemplatesClient.BaseClient)

	directoryObjectsClient := msgraph.NewDirectoryObjectsClient(o.TenantID)
	o.ConfigureClient(&directoryObjectsClient.BaseClient)

	return &Client{
		ApplicationsClient:         applicationsClient,
		ApplicationTemplatesClient: applicationTemplatesClient,
		DirectoryObjectsClient:     directoryObjectsClient,
	}
}
