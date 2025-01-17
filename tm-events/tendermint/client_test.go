package tendermint_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/rpc/client"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/scalarorg/scalar-core/tm-events/tendermint"
	"github.com/scalarorg/scalar-core/tm-events/tendermint/mock"
	. "github.com/scalarorg/scalar-core/utils/test"
	"github.com/scalarorg/scalar-core/utils/test/rand"
)

func TestResettableClient(t *testing.T) {
	var (
		expectedBlockHeight = rand.PosI64()
		resettableClient    *tendermint.RobustClient
	)
	Given("a resettable client with an unreliable connection", func() {
		resettableClient = tendermint.NewRobustClient(func() (client.Client, error) {
			untilFailure := rand.I64Between(1, 100)
			calls := int64(0)
			return &mock.ClientMock{
				StatusFunc: func(context.Context) (*coretypes.ResultStatus, error) {
					calls++
					if calls > untilFailure {
						return nil, fmt.Errorf("post failed: some connection error")
					}

					return &coretypes.ResultStatus{
						SyncInfo: coretypes.SyncInfo{LatestBlockHeight: expectedBlockHeight},
					}, nil
				},
				StopFunc: func() error { return nil },
			}, nil
		})
	}).
		When("calling a tendermint function until the connection fails", func() {
			var err error
			for err == nil {
				var syncInfo *coretypes.SyncInfo
				syncInfo, err = resettableClient.LatestSyncInfo(context.Background())
				if err == nil {
					assert.Equal(t, expectedBlockHeight, syncInfo.LatestBlockHeight)
				}
			}
		}).
		Then("recover the connection", func(t *testing.T) {
			syncInfo, err := resettableClient.LatestSyncInfo(context.Background())
			assert.NoError(t, err)
			assert.Equal(t, expectedBlockHeight, syncInfo.LatestBlockHeight)
		}).Run(t, 20)
}
