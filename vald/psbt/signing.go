package psbt

import (
	"bytes"
	"context"
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/scalarorg/scalar-core/utils/clog"
	"github.com/scalarorg/scalar-core/utils/log"
	"github.com/scalarorg/scalar-core/x/chains/types"
	covenantTypes "github.com/scalarorg/scalar-core/x/covenant/types"
)

// TODO: Validate psbt inputs whether they are available in the btc chain

// ProcessSigningPsbtStarted handles event signing psbt started
func (mgr *Mgr) ProcessSigningPsbtStarted(event *covenantTypes.SigningPsbtStarted) error {
	if !types.IsBitcoinChain(event.Chain) {
		return nil
	}

	clog.Redf("ProcessSigningPsbtStarted, event: %+v", event)

	clog.Redf("PSBT: %+x", event.Psbt)

	mgrParticipant := mgr.valAddr.String()

	clog.Yellowf("mgrParticipant: %s", mgrParticipant)

	pubKey, ok := event.PubKeys[mgrParticipant]
	if !ok {
		return nil
	}

	clog.Yellowf("pubKey: %v", pubKey.String())
	clog.Yellowf("pubKey: %v", pubKey)
	clog.Yellow("pubKey: ", pubKey)

	if !mgr.validatePubKey(pubKey) {
		return fmt.Errorf("invalid pubKey")
	}

	// TODO: validate the psbt inputs whether they are available in the btc chain

	keyUID := fmt.Sprintf("%s_%d", event.GetKeyID().String(), 0)
	partyUID := mgr.valAddr.String()

	clog.Yellowf("keyUID: %s", keyUID)
	clog.Yellowf("partyUID: %s", partyUID)

	listOfTapScriptSig, err := mgr.sign(keyUID, event.Psbt)
	if err != nil {
		return err
	}

	for i, tapScriptSig := range listOfTapScriptSig.TapScriptSigs {
		clog.Yellowf("tapScriptSig[%d]: %+v", i, tapScriptSig)
	}

	log.Infof("operator %s sending signature for signing %d", partyUID, event.GetSigID())

	msg := covenantTypes.NewSubmitTapScriptSigsRequest(mgr.ctx.FromAddress, event.GetSigID(), listOfTapScriptSig)
	if _, err := mgr.b.Broadcast(context.Background(), msg); err != nil {
		return sdkerrors.Wrap(err, "handler goroutine: failure to broadcast outgoing submit signature message")
	}

	return nil
}

func (mgr *Mgr) validatePubKey(pubKey []byte) bool {
	if len(pubKey) != 33 {
		return false
	}

	if !bytes.Equal(mgr.privKey.Serialize(), pubKey) {
		return false
	}

	return true
}
