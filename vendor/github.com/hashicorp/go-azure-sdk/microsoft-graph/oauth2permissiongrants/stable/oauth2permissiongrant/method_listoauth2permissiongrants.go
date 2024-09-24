package oauth2permissiongrant

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOAuth2PermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.OAuth2PermissionGrant
}

type ListOAuth2PermissionGrantsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.OAuth2PermissionGrant
}

type ListOAuth2PermissionGrantsOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListOAuth2PermissionGrantsOperationOptions() ListOAuth2PermissionGrantsOperationOptions {
	return ListOAuth2PermissionGrantsOperationOptions{}
}

func (o ListOAuth2PermissionGrantsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOAuth2PermissionGrantsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListOAuth2PermissionGrantsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOAuth2PermissionGrantsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOAuth2PermissionGrantsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOAuth2PermissionGrants - List oauth2PermissionGrants (delegated permission grants). Retrieve a list of
// oAuth2PermissionGrant objects, representing delegated permissions which have been granted for client applications to
// access APIs on behalf of signed-in users.
func (c OAuth2PermissionGrantClient) ListOAuth2PermissionGrants(ctx context.Context, options ListOAuth2PermissionGrantsOperationOptions) (result ListOAuth2PermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOAuth2PermissionGrantsCustomPager{},
		Path:          "/oauth2PermissionGrants",
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]stable.OAuth2PermissionGrant `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOAuth2PermissionGrantsComplete retrieves all the results into a single object
func (c OAuth2PermissionGrantClient) ListOAuth2PermissionGrantsComplete(ctx context.Context, options ListOAuth2PermissionGrantsOperationOptions) (ListOAuth2PermissionGrantsCompleteResult, error) {
	return c.ListOAuth2PermissionGrantsCompleteMatchingPredicate(ctx, options, OAuth2PermissionGrantOperationPredicate{})
}

// ListOAuth2PermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OAuth2PermissionGrantClient) ListOAuth2PermissionGrantsCompleteMatchingPredicate(ctx context.Context, options ListOAuth2PermissionGrantsOperationOptions, predicate OAuth2PermissionGrantOperationPredicate) (result ListOAuth2PermissionGrantsCompleteResult, err error) {
	items := make([]stable.OAuth2PermissionGrant, 0)

	resp, err := c.ListOAuth2PermissionGrants(ctx, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListOAuth2PermissionGrantsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}