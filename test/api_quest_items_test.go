/*
dofusdude

Testing QuestItemsAPIService

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

func Test_dodugo_QuestItemsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test QuestItemsAPIService GetAllItemsQuestList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var language string
		var game string

		resp, httpRes, err := apiClient.QuestItemsAPI.GetAllItemsQuestList(context.Background(), language, game).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QuestItemsAPIService GetItemQuestSingle", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var language string
		var ankamaId int32
		var game string

		resp, httpRes, err := apiClient.QuestItemsAPI.GetItemQuestSingle(context.Background(), language, ankamaId, game).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QuestItemsAPIService GetItemsQuestList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var language string
		var game string

		resp, httpRes, err := apiClient.QuestItemsAPI.GetItemsQuestList(context.Background(), language, game).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QuestItemsAPIService GetItemsQuestSearch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var language string
		var game string

		resp, httpRes, err := apiClient.QuestItemsAPI.GetItemsQuestSearch(context.Background(), language, game).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
