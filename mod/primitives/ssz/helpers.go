// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package ssz

import (
	"encoding/binary"
	"fmt"
	"reflect"

	"github.com/berachain/beacon-kit/mod/primitives/constants"
	"github.com/prysmaticlabs/gohashtree"
)

// SizeOfBasic returns the size of a basic type.
func SizeOfBasic[RootT ~[32]byte, B Basic[SpecT, RootT], SpecT any](
	b B,
) uint64 {
	// TODO: Boolean maybe this doesnt work.
	return uint64(reflect.TypeOf(b).Size())
}

// SizeOfComposite returns the size of a composite type.
func SizeOfComposite[RootT ~[32]byte, C Composite[SpecT, RootT], SpecT any](
	c C,
) uint64 {
	//#nosec:G701 // This is a safe operation.
	return uint64(c.SizeSSZ())
}

// SizeOfContainer returns the size of a container type.
func SizeOfContainer[RootT ~[32]byte, C Container[SpecT, RootT], SpecT any](
	c C,
) int {
	size := 0
	rValue := reflect.ValueOf(c)
	if rValue.Kind() == reflect.Ptr {
		rValue = rValue.Elem()
	}
	for i := range rValue.NumField() {
		fieldValue := rValue.Field(i)
		if !fieldValue.CanInterface() {
			return -1
		}

		// TODO: handle different types.
		field, ok := fieldValue.Interface().(Basic[SpecT, RootT])
		if !ok {
			return -1
		}
		size += field.SizeSSZ()

		// TODO: handle the offset calculation.
	}

	// TODO: This doesn't yet handle anything to do with offset calculation.
	return size
}

// ChunkCount returns the number of chunks required to store a value.
func ChunkCountBasic[RootT ~[32]byte, B Basic[SpecT, RootT], SpecT any](
	B,
) uint64 {
	return 1
}

// ChunkCountBitListVec returns the number of chunks required to store a bitlist
// or bitvector.
func ChunkCountBitListVec[T any](t []T) uint64 {
	//nolint:mnd // 256 is okay.
	return (uint64(len(t)) + 255) / 256
}

// ChunkCountBasicList returns the number of chunks required to store a list
// or vector of basic types.
func ChunkCountBasicList[SpecT any, RootT ~[32]byte, B Basic[SpecT, RootT]](
	b []B,
	maxCapacity uint64,
) uint64 {
	numItems := uint64(len(b))
	if numItems == 0 {
		return 1
	}
	size := SizeOfBasic[RootT, B, SpecT](b[0])
	//nolint:mnd // 32 is okay.
	limit := (maxCapacity*size + 31) / 32
	if limit != 0 {
		return limit
	}

	return numItems
}

// ChunkCountCompositeList returns the number of chunks required to store a
// list or vector of composite types.
func ChunkCountCompositeList[
	SpecT any, RootT ~[32]byte, C Composite[SpecT, RootT],
](
	c []C,
	limit uint64,
) uint64 {
	return max(uint64(len(c)), limit)
}

// ChunkCountContainer returns the number of chunks required to store a
// container.
func ChunkCountContainer[SpecT any, RootT ~[32]byte, C Container[SpecT, RootT]](
	c C,
) uint64 {
	//#nosec:G701 // This is a safe operation.
	return uint64(reflect.ValueOf(c).NumField())
}

// PadTo function to pad the chunks to the effective limit with zeroed chunks.
func PadTo[U64T ~uint64, ChunkT ~[32]byte](
	chunks []ChunkT,
	effectiveLimit U64T,
) []ChunkT {
	paddedChunks := make([]ChunkT, effectiveLimit)
	copy(paddedChunks, chunks)
	//#nosec:G701 // This is a safe operation.
	for i := uint64(len(chunks)); i < uint64(effectiveLimit); i++ {
		paddedChunks[i] = ChunkT{}
	}
	return paddedChunks
}

// Pack packs a list of SSZ-marshallable elements into a single byte slice.
func Pack[
	U64T U64[U64T],
	U256L U256LT,
	SpecT any,
	RootT ~[32]byte,
	B Basic[SpecT, RootT],
](b []B) ([]RootT, error) {
	// Pack each element into separate buffers.
	var packed []byte
	for _, el := range b {
		fieldValue := reflect.ValueOf(el)
		if fieldValue.Kind() == reflect.Ptr {
			fieldValue = fieldValue.Elem()
		}

		if !fieldValue.CanInterface() {
			return nil, fmt.Errorf("cannot interface with field %v", fieldValue)
		}

		// TODO: Do we need a safety check for Basic only here?
		// TODO: use a real interface instead of hood inline.
		el, ok := reflect.ValueOf(el).
			Interface().(interface{ MarshalSSZ() ([]byte, error) })
		if !ok {
			return nil, fmt.Errorf("unsupported type %T", el)
		}

		// TODO: Do we need a safety check for Basic only here?
		buf, err := el.MarshalSSZ()
		if err != nil {
			return nil, err
		}
		packed = append(packed, buf...)
	}

	root, _, err := PartitionBytes[RootT](packed)
	return root, err
}

// PartitionBytes partitions a byte slice into chunks of a given length.
func PartitionBytes[RootT ~[32]byte](input []byte) ([]RootT, uint64, error) {
	//nolint:mnd // we add 31 in order to round up the division.
	numChunks := max((uint64(len(input))+31)/constants.RootLength, 1)
	chunks := make([]RootT, numChunks)
	for i := range chunks {
		copy(chunks[i][:], input[32*i:])
	}
	return chunks, numChunks, nil
}

// MerkleizeByteSlice hashes a byteslice by chunkifying it and returning the
// corresponding HTR as if it were a fixed vector of bytes of the given length.
func MerkleizeByteSlice[U64T U64[U64T], RootT ~[32]byte](
	input []byte,
) (RootT, error) {
	chunks, numChunks, err := PartitionBytes[RootT](input)
	if err != nil {
		return RootT{}, err
	}
	return Merkleize[U64T, RootT, RootT](
		chunks,
		numChunks,
	)
}

// MixinLength takes a root element and mixes in the length of the elements
// that were hashed to produce it.
func MixinLength[RootT ~[32]byte](element RootT, length uint64) RootT {
	//nolint:mnd // 2 is okay.
	chunks := make([][32]byte, 2)
	chunks[0] = element
	binary.LittleEndian.PutUint64(chunks[1][:], length)
	if err := gohashtree.Hash(chunks, chunks); err != nil {
		return RootT{}
	}
	return chunks[0]
}

// IsBasicType returns true if the type is a basic type
// i.e. UInt or Boolean types.
func IsBasicType(t Type) bool {
	k := t.Kind()
	return k == KindUInt || k == KindBool
}

// IsVariableSize returns true if the object is variable-size.
// A variable-size types to be lists, (unions, Bitlist are not supported yet)
// and all types that contain a variable-size type.
// All other types are said to be fixed-size.
func IsVariableSize(t Type) bool {
	switch t.Kind() {
	case KindUInt, KindBool:
		return false
	case KindList:
		return true
	case KindVector:
		return IsVariableSize(t.(VectorType).ElemType())
	case KindContainer:
		for _, ft := range t.(ContainerType).FieldTypes() {
			if IsVariableSize(ft) {
				return true
			}
		}
		return false
	default:
		return false
	}
}

// IsFixedSize returns true if the object is fixed-size.
func IsFixedSize(t Type) bool {
	return !IsVariableSize(t)
}

// IsList returns true if the type is a list type.
func IsList(t Type) bool {
	return t.Kind() == KindList
}
