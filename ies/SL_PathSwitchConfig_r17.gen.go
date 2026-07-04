// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PathSwitchConfig-r17 (line 5768).

var sLPathSwitchConfigR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "targetRelayUE-Identity-r17"},
		{Name: "t420-r17"},
	},
}

const (
	SL_PathSwitchConfig_r17_T420_r17_Ms50    = 0
	SL_PathSwitchConfig_r17_T420_r17_Ms100   = 1
	SL_PathSwitchConfig_r17_T420_r17_Ms150   = 2
	SL_PathSwitchConfig_r17_T420_r17_Ms200   = 3
	SL_PathSwitchConfig_r17_T420_r17_Ms500   = 4
	SL_PathSwitchConfig_r17_T420_r17_Ms1000  = 5
	SL_PathSwitchConfig_r17_T420_r17_Ms2000  = 6
	SL_PathSwitchConfig_r17_T420_r17_Ms10000 = 7
)

var sLPathSwitchConfigR17T420R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PathSwitchConfig_r17_T420_r17_Ms50, SL_PathSwitchConfig_r17_T420_r17_Ms100, SL_PathSwitchConfig_r17_T420_r17_Ms150, SL_PathSwitchConfig_r17_T420_r17_Ms200, SL_PathSwitchConfig_r17_T420_r17_Ms500, SL_PathSwitchConfig_r17_T420_r17_Ms1000, SL_PathSwitchConfig_r17_T420_r17_Ms2000, SL_PathSwitchConfig_r17_T420_r17_Ms10000},
}

type SL_PathSwitchConfig_r17 struct {
	TargetRelayUE_Identity_r17 SL_SourceIdentity_r17
	T420_r17                   int64
}

func (ie *SL_PathSwitchConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPathSwitchConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. targetRelayUE-Identity-r17: ref
	{
		if err := ie.TargetRelayUE_Identity_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. t420-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.T420_r17, sLPathSwitchConfigR17T420R17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_PathSwitchConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPathSwitchConfigR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. targetRelayUE-Identity-r17: ref
	{
		if err := ie.TargetRelayUE_Identity_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. t420-r17: enumerated
	{
		v1, err := d.DecodeEnumerated(sLPathSwitchConfigR17T420R17Constraints)
		if err != nil {
			return err
		}
		ie.T420_r17 = v1
	}

	return nil
}
