package merkle_tree

import (
	"bytes"
	"crypto/sha256"
	"errors"

	proto "github.com/RedCuckoo/merkle-tree-verifier/proto/generated"
)

type MerkleProof struct {
	Proof      [][]byte
	Direction  []bool
	ProvenData []byte
}

func (m *MerkleProof) MarshalProto() *proto.MerkleProof {
	return &proto.MerkleProof{
		Proof:      m.Proof,
		Direction:  m.Direction,
		ProvenData: m.ProvenData,
	}
}

func (m *MerkleProof) UnmarshalProto(protoProof *proto.MerkleProof) *MerkleProof {
	*m = MerkleProof{
		Proof:      protoProof.GetProof(),
		Direction:  protoProof.GetDirection(),
		ProvenData: protoProof.GetProvenData(),
	}

	return m
}

type MerkleTreeVerifier struct{}

func (t *MerkleTreeVerifier) ValidateFileByProof(
	file []byte,
	proof *MerkleProof,
	rootHash []byte,
) (bool, error) {
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
