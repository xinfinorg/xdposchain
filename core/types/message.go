package types

import (
	"github.com/XinFinOrg/XDPoSChain/common"
	"github.com/XinFinOrg/XDPoSChain/common/math"
	"math/big"
)

// Message is a fully derived transaction and implements core.Message
//
// NOTE: In a future PR this will be removed.
type Message struct {
	to              *common.Address
	from            common.Address
	nonce           uint64
	amount          *big.Int
	gasLimit        uint64
	gasPrice        *big.Int
	feeCap          *big.Int
	tip             *big.Int
	data            []byte
	accessList      AccessList
	checkNonce      bool
	balanceTokenFee *big.Int
}

func NewMessage(from common.Address, to *common.Address, nonce uint64, amount *big.Int, gasLimit uint64, gasPrice, feeCap, tip *big.Int, accessList AccessList, data []byte, checkNonce bool, balanceTokenFee *big.Int) Message {
	if balanceTokenFee != nil {
		gasPrice = common.TRC21GasPrice
	}
	return Message{
		from:            from,
		to:              to,
		nonce:           nonce,
		amount:          amount,
		gasLimit:        gasLimit,
		gasPrice:        gasPrice,
		feeCap:          feeCap,
		tip:             tip,
		data:            data,
		accessList:      accessList,
		checkNonce:      checkNonce,
		balanceTokenFee: balanceTokenFee,
	}
}

// AsMessage returns the transaction as a core.Message.
func (tx *Transaction) AsMessage(s Signer, balanceFee *big.Int, number *big.Int, baseFee *big.Int) (Message, error) {
	msg := Message{
		nonce:           tx.Nonce(),
		gasLimit:        tx.Gas(),
		gasPrice:        new(big.Int).Set(tx.GasPrice()),
		feeCap:          new(big.Int).Set(tx.FeeCap()),
		tip:             new(big.Int).Set(tx.Tip()),
		to:              tx.To(),
		amount:          tx.Value(),
		data:            tx.Data(),
		accessList:      tx.AccessList(),
		balanceTokenFee: balanceFee,
		checkNonce:      true,
	}
	if balanceFee != nil {
		if number.Cmp(common.TIPTRC21Fee) > 0 {
			msg.gasPrice = common.TRC21GasPrice
		} else {
			msg.gasPrice = common.TRC21GasPriceBefore
		}
	}
	var err error
	msg.from, err = Sender(s, tx)

	// If baseFee provided, set gasPrice to effectiveGasPrice.
	if baseFee != nil {
		msg.gasPrice = math.BigMin(msg.gasPrice.Add(msg.tip, baseFee), msg.feeCap)
	}
	if tx.FeeCap() != nil {
		msg.feeCap.Set(tx.FeeCap())
	}
	if tx.Tip() != nil {
		msg.tip.Set(tx.Tip())
	}

	return msg, err
}

func (m Message) From() common.Address      { return m.from }
func (m Message) BalanceTokenFee() *big.Int { return m.balanceTokenFee }
func (m Message) To() *common.Address       { return m.to }
func (m Message) GasPrice() *big.Int        { return m.gasPrice }
func (m Message) FeeCap() *big.Int          { return new(big.Int).Set(m.feeCap) }
func (m Message) Tip() *big.Int             { return new(big.Int).Set(m.tip) }
func (m Message) Value() *big.Int           { return m.amount }
func (m Message) Gas() uint64               { return m.gasLimit }
func (m Message) Nonce() uint64             { return m.nonce }
func (m Message) Data() []byte              { return m.data }
func (m Message) CheckNonce() bool          { return m.checkNonce }

func (m *Message) SetNonce(nonce uint64) { m.nonce = nonce }
