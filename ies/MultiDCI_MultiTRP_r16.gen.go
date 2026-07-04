// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MultiDCI-MultiTRP-r16 (line 19960).

var multiDCIMultiTRPR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberCORESET-r16"},
		{Name: "maxNumberCORESETPerPoolIndex-r16"},
		{Name: "maxNumberUnicastPDSCH-PerPool-r16"},
	},
}

const (
	MultiDCI_MultiTRP_r16_MaxNumberCORESET_r16_N2 = 0
	MultiDCI_MultiTRP_r16_MaxNumberCORESET_r16_N3 = 1
	MultiDCI_MultiTRP_r16_MaxNumberCORESET_r16_N4 = 2
	MultiDCI_MultiTRP_r16_MaxNumberCORESET_r16_N5 = 3
)

var multiDCIMultiTRPR16MaxNumberCORESETR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MultiDCI_MultiTRP_r16_MaxNumberCORESET_r16_N2, MultiDCI_MultiTRP_r16_MaxNumberCORESET_r16_N3, MultiDCI_MultiTRP_r16_MaxNumberCORESET_r16_N4, MultiDCI_MultiTRP_r16_MaxNumberCORESET_r16_N5},
}

var multiDCIMultiTRPR16MaxNumberCORESETPerPoolIndexR16Constraints = per.Constrained(1, 3)

const (
	MultiDCI_MultiTRP_r16_MaxNumberUnicastPDSCH_PerPool_r16_N1 = 0
	MultiDCI_MultiTRP_r16_MaxNumberUnicastPDSCH_PerPool_r16_N2 = 1
	MultiDCI_MultiTRP_r16_MaxNumberUnicastPDSCH_PerPool_r16_N3 = 2
	MultiDCI_MultiTRP_r16_MaxNumberUnicastPDSCH_PerPool_r16_N4 = 3
	MultiDCI_MultiTRP_r16_MaxNumberUnicastPDSCH_PerPool_r16_N7 = 4
)

var multiDCIMultiTRPR16MaxNumberUnicastPDSCHPerPoolR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MultiDCI_MultiTRP_r16_MaxNumberUnicastPDSCH_PerPool_r16_N1, MultiDCI_MultiTRP_r16_MaxNumberUnicastPDSCH_PerPool_r16_N2, MultiDCI_MultiTRP_r16_MaxNumberUnicastPDSCH_PerPool_r16_N3, MultiDCI_MultiTRP_r16_MaxNumberUnicastPDSCH_PerPool_r16_N4, MultiDCI_MultiTRP_r16_MaxNumberUnicastPDSCH_PerPool_r16_N7},
}

type MultiDCI_MultiTRP_r16 struct {
	MaxNumberCORESET_r16              int64
	MaxNumberCORESETPerPoolIndex_r16  int64
	MaxNumberUnicastPDSCH_PerPool_r16 int64
}

func (ie *MultiDCI_MultiTRP_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(multiDCIMultiTRPR16Constraints)
	_ = seq

	// 1. maxNumberCORESET-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberCORESET_r16, multiDCIMultiTRPR16MaxNumberCORESETR16Constraints); err != nil {
			return err
		}
	}

	// 2. maxNumberCORESETPerPoolIndex-r16: integer
	{
		if err := e.EncodeInteger(ie.MaxNumberCORESETPerPoolIndex_r16, multiDCIMultiTRPR16MaxNumberCORESETPerPoolIndexR16Constraints); err != nil {
			return err
		}
	}

	// 3. maxNumberUnicastPDSCH-PerPool-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.MaxNumberUnicastPDSCH_PerPool_r16, multiDCIMultiTRPR16MaxNumberUnicastPDSCHPerPoolR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MultiDCI_MultiTRP_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(multiDCIMultiTRPR16Constraints)
	_ = seq

	// 1. maxNumberCORESET-r16: enumerated
	{
		v0, err := d.DecodeEnumerated(multiDCIMultiTRPR16MaxNumberCORESETR16Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumberCORESET_r16 = v0
	}

	// 2. maxNumberCORESETPerPoolIndex-r16: integer
	{
		v1, err := d.DecodeInteger(multiDCIMultiTRPR16MaxNumberCORESETPerPoolIndexR16Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumberCORESETPerPoolIndex_r16 = v1
	}

	// 3. maxNumberUnicastPDSCH-PerPool-r16: enumerated
	{
		v2, err := d.DecodeEnumerated(multiDCIMultiTRPR16MaxNumberUnicastPDSCHPerPoolR16Constraints)
		if err != nil {
			return err
		}
		ie.MaxNumberUnicastPDSCH_PerPool_r16 = v2
	}

	return nil
}
