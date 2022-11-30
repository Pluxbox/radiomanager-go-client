/*
RadioManager

Testing BlockApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package radiomanagerclient

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_radiomanagerclient_BlockApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test BlockApiService GetBlockById", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var id int64

        resp, httpRes, err := apiClient.BlockApi.GetBlockById(context.Background(), id).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BlockApiService GetCurrentBlock", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.BlockApi.GetCurrentBlock(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BlockApiService GetNextBlock", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.BlockApi.GetNextBlock(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test BlockApiService ListBlocks", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.BlockApi.ListBlocks(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
