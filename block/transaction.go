package block

import (
	"encoding/json"
	"fmt"
	"strings"
)

type BlockTransaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(sender string, recipient string, value float32) *BlockTransaction {
	return &BlockTransaction{
		senderBlockchainAddress:    sender,
		recipientBlockchainAddress: recipient,
		value:                      value,
	}
}

func (bt *BlockTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockchainAddress    string  `json:"sender_blockchain_address"`
		RecipientBlockchainAddress string  `json:"recipient_blockchain_address"`
		Value                      float32 `json:"value"`
	}{
		SenderBlockchainAddress:    bt.senderBlockchainAddress,
		RecipientBlockchainAddress: bt.recipientBlockchainAddress,
		Value:                      bt.value,
	})
}

func (bt *BlockTransaction) UnmarshalJSON(data []byte) error {
	v := struct {
		SenderBlockchainAddress    *string  `json:"sender_blockchain_address"`
		RecipientBlockchainAddress *string  `json:"recipient_blockchain_address"`
		Value                      *float32 `json:"value"`
	}{
		SenderBlockchainAddress:    &bt.senderBlockchainAddress,
		RecipientBlockchainAddress: &bt.recipientBlockchainAddress,
		Value:                      &bt.value,
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	return nil
}

func (bt *BlockTransaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address       %s\n", bt.senderBlockchainAddress)
	fmt.Printf(" recipient_blockchain_address    %s\n", bt.recipientBlockchainAddress)
	fmt.Printf(" value:                          %.1f\n", bt.value)
}

type TransactionRequest struct {
	SenderBlockchainAddress    *string  `json:"sender_blockchain_address"`
	RecipientBlockchainAddress *string  `json:"recipient_blockchain_address"`
	SenderPublicKey            *string  `json:"sender_public_key"`
	Value                      *float32 `json:"value"`
	Signature                  *string  `json:"signature"`
}

func (tr *TransactionRequest) Validate() bool {
	if tr.SenderBlockchainAddress == nil || tr.RecipientBlockchainAddress == nil || tr.SenderPublicKey == nil || tr.Value == nil || tr.Signature == nil {
		return false
	}

	return true
}
