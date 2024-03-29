/*
dofusdude

Testing AlmanaxAPIService

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

func Test_dodugo_AlmanaxAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AlmanaxAPIService GetAlmanaxDate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var language string
		var date string

		resp, httpRes, err := apiClient.AlmanaxAPI.GetAlmanaxDate(context.Background(), language, date).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AlmanaxAPIService GetAlmanaxRange", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var language string

		resp, httpRes, err := apiClient.AlmanaxAPI.GetAlmanaxRange(context.Background(), language).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
