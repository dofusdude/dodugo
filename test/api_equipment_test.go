/*
Dofusdude

Testing EquipmentApiService

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

func Test_dodugo_EquipmentApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EquipmentApiService GetAllItemsEquipmentList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var language string
		var game string

		resp, httpRes, err := apiClient.EquipmentApi.GetAllItemsEquipmentList(context.Background(), language, game).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EquipmentApiService GetItemsEquipmentList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var language string
		var game string

		resp, httpRes, err := apiClient.EquipmentApi.GetItemsEquipmentList(context.Background(), language, game).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EquipmentApiService GetItemsEquipmentSearch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var language string
		var game string

		resp, httpRes, err := apiClient.EquipmentApi.GetItemsEquipmentSearch(context.Background(), language, game).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EquipmentApiService GetItemsEquipmentSingle", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var language string
		var ankamaId int32
		var game string

		resp, httpRes, err := apiClient.EquipmentApi.GetItemsEquipmentSingle(context.Background(), language, ankamaId, game).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
