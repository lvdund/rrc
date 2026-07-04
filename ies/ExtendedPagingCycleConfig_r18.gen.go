// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ExtendedPagingCycleConfig-r18 (line 1512).

var extendedPagingCycleConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "extendedPagingCycle-r18"},
		{Name: "pagingPTWLength-r18"},
	},
}

const (
	ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Hf2    = 0
	ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Hf4    = 1
	ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Hf8    = 2
	ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Hf16   = 3
	ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Hf32   = 4
	ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Hf64   = 5
	ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Hf128  = 6
	ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Hf256  = 7
	ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Hf512  = 8
	ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Hf1024 = 9
	ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Spare6 = 10
	ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Spare5 = 11
	ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Spare4 = 12
	ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Spare3 = 13
	ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Spare2 = 14
	ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Spare1 = 15
)

var extendedPagingCycleConfigR18ExtendedPagingCycleR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Hf2, ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Hf4, ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Hf8, ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Hf16, ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Hf32, ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Hf64, ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Hf128, ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Hf256, ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Hf512, ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Hf1024, ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Spare6, ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Spare5, ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Spare4, ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Spare3, ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Spare2, ExtendedPagingCycleConfig_r18_ExtendedPagingCycle_r18_Spare1},
}

const (
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms1280  = 0
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms2560  = 1
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms3840  = 2
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms5120  = 3
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms6400  = 4
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms7680  = 5
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms8960  = 6
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms10240 = 7
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms11520 = 8
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms12800 = 9
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms14080 = 10
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms15360 = 11
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms16640 = 12
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms17920 = 13
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms19200 = 14
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms20480 = 15
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms21760 = 16
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms23040 = 17
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms24320 = 18
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms25600 = 19
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms26880 = 20
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms28160 = 21
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms29440 = 22
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms30720 = 23
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms32000 = 24
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms33280 = 25
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms34560 = 26
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms35840 = 27
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms37120 = 28
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms38400 = 29
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms39680 = 30
	ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms40960 = 31
)

var extendedPagingCycleConfigR18PagingPTWLengthR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms1280, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms2560, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms3840, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms5120, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms6400, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms7680, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms8960, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms10240, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms11520, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms12800, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms14080, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms15360, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms16640, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms17920, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms19200, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms20480, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms21760, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms23040, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms24320, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms25600, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms26880, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms28160, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms29440, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms30720, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms32000, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms33280, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms34560, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms35840, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms37120, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms38400, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms39680, ExtendedPagingCycleConfig_r18_PagingPTWLength_r18_Ms40960},
}

type ExtendedPagingCycleConfig_r18 struct {
	ExtendedPagingCycle_r18 int64
	PagingPTWLength_r18     int64
}

func (ie *ExtendedPagingCycleConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(extendedPagingCycleConfigR18Constraints)
	_ = seq

	// 1. extendedPagingCycle-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.ExtendedPagingCycle_r18, extendedPagingCycleConfigR18ExtendedPagingCycleR18Constraints); err != nil {
			return err
		}
	}

	// 2. pagingPTWLength-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.PagingPTWLength_r18, extendedPagingCycleConfigR18PagingPTWLengthR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *ExtendedPagingCycleConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(extendedPagingCycleConfigR18Constraints)
	_ = seq

	// 1. extendedPagingCycle-r18: enumerated
	{
		v0, err := d.DecodeEnumerated(extendedPagingCycleConfigR18ExtendedPagingCycleR18Constraints)
		if err != nil {
			return err
		}
		ie.ExtendedPagingCycle_r18 = v0
	}

	// 2. pagingPTWLength-r18: enumerated
	{
		v1, err := d.DecodeEnumerated(extendedPagingCycleConfigR18PagingPTWLengthR18Constraints)
		if err != nil {
			return err
		}
		ie.PagingPTWLength_r18 = v1
	}

	return nil
}
