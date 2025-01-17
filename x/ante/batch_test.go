package ante_test

// import (
// 	"testing"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/stretchr/testify/assert"

// 	appParams "github.com/scalarorg/scalar-core/app/params"
// 	"github.com/scalarorg/scalar-core/testutils/rand"
// 	auxiliarytypes "github.com/scalarorg/scalar-core/x/auxiliary/types"
// 	chainsTypes "github.com/scalarorg/scalar-core/x/chains/types"
// 	vote "github.com/scalarorg/scalar-core/x/vote/exported"
// 	votetypes "github.com/scalarorg/scalar-core/x/vote/types"
// 	. "github.com/scalarorg/scalar-core/utils/test"
// 	"github.com/scalarorg/scalar-core/x/ante"
// 	"github.com/scalarorg/scalar-core/x/ante/types/mock"
// 	nexus "github.com/scalarorg/scalar-core/x/nexus/exported"
// )

// func TestBatch(t *testing.T) {
// 	var (
// 		handler       sdk.AnteDecorator
// 		tx            *mock.FeeTxMock
// 		batchMsg      *auxiliarytypes.BatchRequest
// 		unwrappedMsgs []sdk.Msg
// 	)

// 	sender := rand.AccAddr()

// 	givenBatchAnteHandler := Given("the batch ante handler", func() {
// 		encCfg := appParams.MakeEncodingConfig()
// 		handler = ante.NewBatchDecorator(encCfg.Codec)
// 	})

// 	givenBatchAnteHandler.
// 		When("messages do not contain batch", func() {
// 			tx = &mock.FeeTxMock{
// 				GetMsgsFunc: func() []sdk.Msg {
// 					return []sdk.Msg{
// 						votetypes.NewVoteRequest(sender, vote.PollID(rand.PosI64()), chainsTypes.NewVoteEvents(nexus.ChainName(rand.NormalizedStr(3)))),
// 						votetypes.NewVoteRequest(sender, vote.PollID(rand.PosI64()), chainsTypes.NewVoteEvents(nexus.ChainName(rand.NormalizedStr(3)))),
// 						votetypes.NewVoteRequest(sender, vote.PollID(rand.PosI64()), chainsTypes.NewVoteEvents(nexus.ChainName(rand.NormalizedStr(3)))),
// 					}
// 				},
// 			}
// 		}).
// 		Then("should pass messages as it", func(t *testing.T) {
// 			_, err := handler.AnteHandle(sdk.Context{}, tx, false,
// 				func(_ sdk.Context, tx sdk.Tx, _ bool) (sdk.Context, error) {
// 					unwrappedMsgs = tx.GetMsgs()
// 					return sdk.Context{}, nil
// 				})

// 			assert.NoError(t, err)
// 			assert.Equal(t, 3, len(unwrappedMsgs))
// 		}).
// 		Run(t)

// 	givenBatchAnteHandler.
// 		When("a Batch Request is valid", func() {
// 			batchMsg = auxiliarytypes.NewBatchRequest(sender, []sdk.Msg{
// 				votetypes.NewVoteRequest(sender, vote.PollID(rand.PosI64()), chainsTypes.NewVoteEvents(nexus.ChainName(rand.NormalizedStr(3)))),
// 				votetypes.NewVoteRequest(sender, vote.PollID(rand.PosI64()), chainsTypes.NewVoteEvents(nexus.ChainName(rand.NormalizedStr(3)))),
// 			})
// 		}).
// 		Then("should unwrap inner message", func(t *testing.T) {
// 			tx = &mock.FeeTxMock{
// 				GetMsgsFunc: func() []sdk.Msg {
// 					return []sdk.Msg{
// 						votetypes.NewVoteRequest(sender, vote.PollID(rand.PosI64()), chainsTypes.NewVoteEvents(nexus.ChainName(rand.NormalizedStr(3)))),
// 						batchMsg,
// 						votetypes.NewVoteRequest(sender, vote.PollID(rand.PosI64()), chainsTypes.NewVoteEvents(nexus.ChainName(rand.NormalizedStr(3)))),
// 						batchMsg,
// 						votetypes.NewVoteRequest(sender, vote.PollID(rand.PosI64()), chainsTypes.NewVoteEvents(nexus.ChainName(rand.NormalizedStr(3)))),
// 					}
// 				},
// 			}

// 			_, err := handler.AnteHandle(sdk.Context{}, tx, false,
// 				func(_ sdk.Context, tx sdk.Tx, _ bool) (sdk.Context, error) {
// 					unwrappedMsgs = tx.GetMsgs()
// 					return sdk.Context{}, nil
// 				})

// 			assert.NoError(t, err)
// 			assert.Equal(t, 9, len(unwrappedMsgs))
// 		}).
// 		Run(t)
// }
