package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"goblockchain/utils"
)

type WalletTransaction struct {
	senderPrivateKey           *ecdsa.PrivateKey
	senderPublicKey            *ecdsa.PublicKey
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey, sender string, recipient string, value float32) *WalletTransaction {
	return &WalletTransaction{privateKey, publicKey, sender, recipient, value}
}

func (wt *WalletTransaction) GenerateSignature() *utils.Signature {
	m, _ := json.Marshal(wt)
	h := sha256.Sum256([]byte(m))
	r, s, _ := ecdsa.Sign(rand.Reader, wt.senderPrivateKey, h[:])
	return &utils.Signature{R: r, S: s}
}

func (wt *WalletTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    wt.senderBlockchainAddress,
		Recipient: wt.recipientBlockchainAddress,
		Value:     wt.value,
	})
}

type TransactionRequest struct {
	SenderPrivateKey           *string `json:"sender_private_key"`
	SenderPublicKey            *string `json:"sender_public_key"`
	SenderBlockchainAddress    *string `json:"sender_blockchain_address"`
	RecipientBlockchainAddress *string `json:"recipient_blockchain_address"`
	Value                      *string `json:"value"`
}

func (tr *TransactionRequest) Validate() bool {
	if tr.SenderPrivateKey == nil || tr.SenderPublicKey == nil || tr.SenderBlockchainAddress == nil || tr.RecipientBlockchainAddress == nil || tr.Value == nil {
		return false
	}

	return true
}
