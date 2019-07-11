package bitint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetGetValue(t *testing.T) {
	var bitInt BitInt
	bitInt.SetValue(2, 3, 5)                       // 5==2#101
	assert.Equal(t, 20, int(bitInt))               // // 20=2#10100
	assert.Equal(t, 5, int(bitInt.GetValue(2, 3))) // // 20=2#10100
	bitInt.SetValue(5, 3, 5)                       //
	assert.Equal(t, 180, int(bitInt))              // // 180=2#10110100
	assert.Equal(t, 5, int(bitInt.GetValue(5, 3))) // // 2#10110100
	assert.Equal(t, 5, int(bitInt.GetValue(2, 3))) // // 20=2#10100

}

func TestSetFlagByPos(t *testing.T) {
	var bitInt BitInt
	var pos uint64 = 0

	pos = 1
	assert.True(t, bitInt.Get(pos) == 0)
	bitInt.Set(pos)
	assert.True(t, bitInt == 2) // 2#10
	assert.True(t, bitInt.Get(pos) == 1)

	var pos5 uint64 = 5
	assert.True(t, bitInt.Get(pos5) == 0)
	bitInt.Set(pos5)
	assert.True(t, bitInt == 34) // 2#100010
	assert.True(t, bitInt.Get(pos5) == 1)

	bitInt.Unset(pos5)
	assert.True(t, bitInt == 2) // 2#10
	assert.True(t, bitInt.Get(pos5) == 0)
	assert.False(t, bitInt.IsZero())

}

func TestBitIntGetFlagCount(t *testing.T) {
	var bitInt BitInt = 0
	assert.True(t, 0 == bitInt.GetFlagCount())
	bitInt = 1 // 2#1
	assert.True(t, 1 == bitInt.GetFlagCount())
	bitInt = 3 // 2#11
	assert.True(t, 2 == bitInt.GetFlagCount())
	assert.True(t, 1 == bitInt.GetFlagCount(1))

	bitInt = 5 // 2#101
	assert.True(t, 2 == bitInt.GetFlagCount())
	assert.True(t, 1 == bitInt.GetFlagCount(2))

}

func TestBitIntSetByte(t *testing.T) {
	var bitInt BitInt = 0
	bitInt.SetByte(0, 255)
	assert.True(t, 255 == bitInt.GetByte(0))
	assert.True(t, 255 == uint64(bitInt))
	bitInt.SetByte(0, 0)
	assert.True(t, 0 == bitInt.GetByte(0))

	bitInt.SetByte(1, 255)
	assert.True(t, 255 == bitInt.GetByte(1))
	bitInt.SetByte(1, 0)
	assert.True(t, 0 == bitInt.GetByte(1))

	bitInt.SetByte(1, 128)
	assert.True(t, 128 == bitInt.GetByte(1))

	bitInt.SetByte(2, 255)
	assert.True(t, 255 == bitInt.GetByte(2))
	bitInt.SetByte(2, 0)
	assert.True(t, 0 == bitInt.GetByte(2))

	bitInt.SetByte(3, 255)
	assert.True(t, 255 == bitInt.GetByte(3))
	bitInt.SetByte(3, 0)
	assert.True(t, 0 == bitInt.GetByte(3))

	bitInt.SetByte(7, 255)
	assert.True(t, 255 == bitInt.GetByte(7))
	bitInt.SetByte(7, 0)
	assert.True(t, 0 == bitInt.GetByte(7))

}

func TestBitIntSetU16(t *testing.T) {
	var bitInt BitInt = 0
	bitInt.SetUint16(0, 0xffff)
	assert.True(t, 0xffff == bitInt.GetUint16(0))
	assert.True(t, 0xffff == uint64(bitInt))
	bitInt.SetUint16(0, 0)
	assert.True(t, 0 == bitInt.GetUint16(0))

	bitInt.SetUint16(1, 0xffff)
	assert.True(t, 0xffff == bitInt.GetUint16(1))
	bitInt.SetUint16(1, 0)
	assert.True(t, 0 == bitInt.GetUint16(1))

	bitInt.SetUint16(2, 0xffff)
	assert.True(t, 0xffff == bitInt.GetUint16(2))
	bitInt.SetUint16(2, 0)
	assert.True(t, 0 == bitInt.GetUint16(2))

	bitInt.SetUint16(3, 0xffff)
	assert.True(t, 0xffff == bitInt.GetUint16(3), bitInt.GetUint16(3))
	bitInt.SetUint16(3, 0)
	assert.True(t, 0 == bitInt.GetUint16(3))

}

func TestBitIntSetU32(t *testing.T) {
	var bitInt BitInt = 0
	bitInt.SetUint32(0, 0xffffffff)
	assert.True(t, 0xffffffff == bitInt.GetUint32(0))
	assert.True(t, 0xffffffff == uint64(bitInt))
	bitInt.SetUint32(0, 0)
	assert.True(t, 0 == bitInt.GetUint32(0))

	bitInt.SetUint32(1, 0xffffffff)
	assert.True(t, 0xffffffff == bitInt.GetUint32(1))
	bitInt.SetUint32(1, 0)
	assert.True(t, 0 == bitInt.GetUint32(1))

}
