package watcher

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

// Controller is the interface used by other components to communicate with the
// watcher.
type Controller interface {
	// Start allows the Controller to begin accepting watch requests.
	Start() error

	// Stop safely stops any ongoing requests within the Controller.
	Stop()

	// WatchAccountConf watches a new account on-chain for its confirmation. Only
	// one conf watcher per account can be used at any time.
	//
	// NOTE: If there is a previous conf watcher for the given account that has not
	// finished yet, it will be canceled!
	WatchAccountConf(traderKey *btcec.PublicKey,
		txHash chainhash.Hash, script []byte, numConfs, heightHint uint32) error

	// CancelAccountConf cancels the conf watcher of the given account, if one is
	// active.
	CancelAccountConf(traderKey *btcec.PublicKey)

	// WatchAccountSpend watches for the spend of an account. Only one spend watcher
	// per account can be used at any time.
	//
	// NOTE: If there is a previous spend watcher for the given account that has not
	// finished yet, it will be canceled!
	WatchAccountSpend(traderKey *btcec.PublicKey,
		accountPoint wire.OutPoint, script []byte, heightHint uint32) error

	// CancelAccountSpend cancels the spend watcher of the given account, if one is
	// active.
	CancelAccountSpend(traderKey *btcec.PublicKey)

	// WatchAccountExpiration watches for the expiration of an account on-chain.
	// Successive calls for the same account will cancel any previous expiration
	// watch requests and the new expiration will be tracked instead.
	WatchAccountExpiration(traderKey *btcec.PublicKey, expiry uint32) error
}
