package types

type BlockHeader struct {
	Version           uint64          `json:"version"`
	Height            uint64          `json:"height"`
	PreviousBlockHash string          `json:"previous_block_hash"`
	Timestamp         uint64          `json:"timestamp"`
	Nonce             uint64          `json:"nonce"`
	Bits              uint64          `json:"bits"`
	BlockCommitment   BlockCommitment `json:"block_commitment"`
}

type BlockCommitment struct {
	TransactionMerkleRoot string `json:"transaction_merkle_root"`
	TransactionStatusHash string `json:"transaction_status_hash"`
}
