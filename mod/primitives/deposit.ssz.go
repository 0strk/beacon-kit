// Code generated by fastssz. DO NOT EDIT.
// Hash: 7ab730ffe3a73ffb790511943e7ed53934906cb741a5d360a98d21d633a56d60
// Version: 0.1.3
package primitives

import (
	"github.com/berachain/beacon-kit/mod/primitives/math"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Deposit object
func (d *Deposit) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(d)
}

// MarshalSSZTo ssz marshals the Deposit object to a target array
func (d *Deposit) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Pubkey'
	dst = append(dst, d.Pubkey[:]...)

	// Field (1) 'Credentials'
	dst = append(dst, d.Credentials[:]...)

	// Field (2) 'Amount'
	dst = ssz.MarshalUint64(dst, uint64(d.Amount))

	// Field (3) 'Signature'
	dst = append(dst, d.Signature[:]...)

	// Field (4) 'Index'
	dst = ssz.MarshalUint64(dst, d.Index)

	return
}

// UnmarshalSSZ ssz unmarshals the Deposit object
func (d *Deposit) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 192 {
		return ssz.ErrSize
	}

	// Field (0) 'Pubkey'
	copy(d.Pubkey[:], buf[0:48])

	// Field (1) 'Credentials'
	copy(d.Credentials[:], buf[48:80])

	// Field (2) 'Amount'
	d.Amount = math.Gwei(ssz.UnmarshallUint64(buf[80:88]))

	// Field (3) 'Signature'
	copy(d.Signature[:], buf[88:184])

	// Field (4) 'Index'
	d.Index = ssz.UnmarshallUint64(buf[184:192])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Deposit object
func (d *Deposit) SizeSSZ() (size int) {
	size = 192
	return
}

// HashTreeRoot ssz hashes the Deposit object
func (d *Deposit) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(d)
}

// HashTreeRootWith ssz hashes the Deposit object with a hasher
func (d *Deposit) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Pubkey'
	hh.PutBytes(d.Pubkey[:])

	// Field (1) 'Credentials'
	hh.PutBytes(d.Credentials[:])

	// Field (2) 'Amount'
	hh.PutUint64(uint64(d.Amount))

	// Field (3) 'Signature'
	hh.PutBytes(d.Signature[:])

	// Field (4) 'Index'
	hh.PutUint64(d.Index)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Deposit object
func (d *Deposit) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(d)
}
