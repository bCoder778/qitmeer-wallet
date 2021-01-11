package util

import "github.com/Qitmeer/qitmeer/core/types"
// calcMinRequiredTxRelayFee returns the minimum transaction fee required for a
// transaction with the passed serialized size to be accepted into the memory
// pool and relayed.
func CalcMinRequiredTxRelayFee(serializedSize int64, minRelayTxFee types.Amount) int64 {
	// Calculate the minimum fee for a transaction to be allowed into the
	// mempool and relayed by scaling the base fee (which is the minimum
	// free transaction relay fee).  minTxRelayFee is in Atom/KB, so
	// multiply by serializedSize (which is in bytes) and divide by 1000 to
	// get minimum Atoms.
	minFee := (serializedSize * minRelayTxFee.Value) / 1000

	if minFee == 0 && minRelayTxFee.Value > 0 {
		minFee = minRelayTxFee.Value
	}

	// Set the minimum fee to the maximum possible value if the calculated
	// fee is not in the valid range for monetary amounts.
	if minFee < 0 || minFee > types.MaxAmount {
		minFee = types.MaxAmount
	}

	return minFee
}
