package merkle_tree

import (
	"crypto/sha256"
	"math"
)

type MerkleTree struct {
	tree   [][][]byte
	height int
}

func (t *MerkleTree) Init(hashes [][]byte) *MerkleTree {
	height := calcHeight(len(hashes))

	*t = MerkleTree{tree: make([][][]byte, height), height: height}

	if t.height == 0 {
		return t
	}

	t.tree[0] = hashes

	for i := 1; i < t.height; i++ {
		t.tree[i] = calcLevel(t.tree[i-1])
	}

	return t
}

func (t *MerkleTree) GetProof(file uint64) *MerkleProof {
	proof := make([][]byte, t.height-1)
	direction := make([]bool, t.height-1)
	var provenData []byte
	for level := 0; level < t.height-1; level++ {
		hashes := modifyOddLengthHashes(t.tree[level])
		if provenData == nil {
			provenData = hashes[file]
		}
		if file%2 == 0 {
			direction[level] = true
			proof[level] = hashes[file+1]
		} else {
			direction[level] = false
			proof[level] = hashes[file-1]
		}
		file /= 2
	}

	return &MerkleProof{
		Proof:      proof,
		Direction:  direction,
		ProvenData: provenData,
	}
}

func (t *MerkleTree) GetRoot() []byte {
	return t.tree[len(t.tree)-1][0]
}

func modifyOddLengthHashes(hashes [][]byte) [][]byte {
	if len(hashes)%2 == 1 {
		hashes = append(hashes, hashes[len(hashes)-1])
	}
	return hashes
}

func calcLevel(hashes [][]byte) [][]byte {
	if len(hashes)%2 == 1 {
		hashes = append(hashes, hashes[len(hashes)-1])
	}

	result := make([][]byte, len(hashes)/2)
	hash := sha256.New()

	for i := 0; i < len(hashes); i += 2 {
		hash.Reset()
		hash.Write(hashes[i])
		hash.Write(hashes[i+1])
		result[i/2] = hash.Sum(nil)
	}

	return result
}

func calcHeight(amount int) int {
	if amount == 0 {
		return 0
	}
	if amount == 1 {
		return 2
	}
	return int(math.Ceil(math.Log10(float64(amount))/math.Log10(2))) + 1
}
