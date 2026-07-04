// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UL-DataSplitThreshold (line 11276).

const (
	UL_DataSplitThreshold_B0       = 0
	UL_DataSplitThreshold_B100     = 1
	UL_DataSplitThreshold_B200     = 2
	UL_DataSplitThreshold_B400     = 3
	UL_DataSplitThreshold_B800     = 4
	UL_DataSplitThreshold_B1600    = 5
	UL_DataSplitThreshold_B3200    = 6
	UL_DataSplitThreshold_B6400    = 7
	UL_DataSplitThreshold_B12800   = 8
	UL_DataSplitThreshold_B25600   = 9
	UL_DataSplitThreshold_B51200   = 10
	UL_DataSplitThreshold_B102400  = 11
	UL_DataSplitThreshold_B204800  = 12
	UL_DataSplitThreshold_B409600  = 13
	UL_DataSplitThreshold_B819200  = 14
	UL_DataSplitThreshold_B1228800 = 15
	UL_DataSplitThreshold_B1638400 = 16
	UL_DataSplitThreshold_B2457600 = 17
	UL_DataSplitThreshold_B3276800 = 18
	UL_DataSplitThreshold_B4096000 = 19
	UL_DataSplitThreshold_B4915200 = 20
	UL_DataSplitThreshold_B5734400 = 21
	UL_DataSplitThreshold_B6553600 = 22
	UL_DataSplitThreshold_Infinity = 23
	UL_DataSplitThreshold_Spare8   = 24
	UL_DataSplitThreshold_Spare7   = 25
	UL_DataSplitThreshold_Spare6   = 26
	UL_DataSplitThreshold_Spare5   = 27
	UL_DataSplitThreshold_Spare4   = 28
	UL_DataSplitThreshold_Spare3   = 29
	UL_DataSplitThreshold_Spare2   = 30
	UL_DataSplitThreshold_Spare1   = 31
)

var uLDataSplitThresholdConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UL_DataSplitThreshold_B0, UL_DataSplitThreshold_B100, UL_DataSplitThreshold_B200, UL_DataSplitThreshold_B400, UL_DataSplitThreshold_B800, UL_DataSplitThreshold_B1600, UL_DataSplitThreshold_B3200, UL_DataSplitThreshold_B6400, UL_DataSplitThreshold_B12800, UL_DataSplitThreshold_B25600, UL_DataSplitThreshold_B51200, UL_DataSplitThreshold_B102400, UL_DataSplitThreshold_B204800, UL_DataSplitThreshold_B409600, UL_DataSplitThreshold_B819200, UL_DataSplitThreshold_B1228800, UL_DataSplitThreshold_B1638400, UL_DataSplitThreshold_B2457600, UL_DataSplitThreshold_B3276800, UL_DataSplitThreshold_B4096000, UL_DataSplitThreshold_B4915200, UL_DataSplitThreshold_B5734400, UL_DataSplitThreshold_B6553600, UL_DataSplitThreshold_Infinity, UL_DataSplitThreshold_Spare8, UL_DataSplitThreshold_Spare7, UL_DataSplitThreshold_Spare6, UL_DataSplitThreshold_Spare5, UL_DataSplitThreshold_Spare4, UL_DataSplitThreshold_Spare3, UL_DataSplitThreshold_Spare2, UL_DataSplitThreshold_Spare1},
}

type UL_DataSplitThreshold struct {
	Value int64
}

func (v *UL_DataSplitThreshold) Encode(e *per.Encoder) error {
	return e.EncodeEnumerated(v.Value, uLDataSplitThresholdConstraints)
}

func (v *UL_DataSplitThreshold) Decode(d *per.Decoder) error {
	x, err := d.DecodeEnumerated(uLDataSplitThresholdConstraints)
	if err != nil {
		return err
	}
	v.Value = x
	return nil
}
