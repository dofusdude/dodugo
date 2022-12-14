/*
Dofusdude

Testing MetaApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package dodugo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/dofusdude/dodugo"
)

func Test_dodugo_MetaApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MetaApiService GetMetaAlmanaxBonuses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var language string

		resp, httpRes, err := apiClient.MetaApi.GetMetaAlmanaxBonuses(context.Background(), language).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetaApiService GetMetaElements", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MetaApi.GetMetaElements(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
