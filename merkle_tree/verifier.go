package merkle_tree

import (
	"bytes"
	"crypto/sha256"
	"errors"
)

type MerkleProof struct {
	Proof     [][]byte
	Direction []bool
}

type MerkleTreeVerifier struct{}

func (t *MerkleTreeVerifier) ValidateFileByProof(file []byte, proof *MerkleProof, rootHash []byte) (bool, error) {
	if len(proof.Proof) != len(proof.Direction) {
		return false, errors.New("invalid proof provided")
	}

	hasher := sha256.New()

	tempBytes := file

	for i, proofHeight := range proof.Proof {
		hasher.Reset()
		if proof.Direction[i] {
			hasher.Write(tempBytes)
			hasher.Write(proofHeight)
		} else {
			hasher.Write(proofHeight)
			hasher.Write(tempBytes)
		}
		tempBytes = hasher.Sum(nil)
	}

	return bytes.Equal(tempBytes, rootHash), nil
}
