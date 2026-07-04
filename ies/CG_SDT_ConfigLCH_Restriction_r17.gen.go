// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CG-SDT-ConfigLCH-Restriction-r17 (line 1418).

var cGSDTConfigLCHRestrictionR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "logicalChannelIdentity-r17"},
		{Name: "configuredGrantType1Allowed-r17", Optional: true},
		{Name: "allowedCG-List-r17", Optional: true},
	},
}

const (
	CG_SDT_ConfigLCH_Restriction_r17_ConfiguredGrantType1Allowed_r17_True = 0
)

var cGSDTConfigLCHRestrictionR17ConfiguredGrantType1AllowedR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CG_SDT_ConfigLCH_Restriction_r17_ConfiguredGrantType1Allowed_r17_True},
}

var cGSDTConfigLCHRestrictionR17AllowedCGListR17Constraints = per.SizeRange(0, common.MaxNrofConfiguredGrantConfigMAC_1_r16)

type CG_SDT_ConfigLCH_Restriction_r17 struct {
	LogicalChannelIdentity_r17      LogicalChannelIdentity
	ConfiguredGrantType1Allowed_r17 *int64
	AllowedCG_List_r17              []ConfiguredGrantConfigIndexMAC_r16
}

func (ie *CG_SDT_ConfigLCH_Restriction_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cGSDTConfigLCHRestrictionR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ConfiguredGrantType1Allowed_r17 != nil, ie.AllowedCG_List_r17 != nil}); err != nil {
		return err
	}

	// 2. logicalChannelIdentity-r17: ref
	{
		if err := ie.LogicalChannelIdentity_r17.Encode(e); err != nil {
			return err
		}
	}

	// 3. configuredGrantType1Allowed-r17: enumerated
	{
		if ie.ConfiguredGrantType1Allowed_r17 != nil {
			if err := e.EncodeEnumerated(*ie.ConfiguredGrantType1Allowed_r17, cGSDTConfigLCHRestrictionR17ConfiguredGrantType1AllowedR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. allowedCG-List-r17: sequence-of
	{
		if ie.AllowedCG_List_r17 != nil {
			seqOf := e.NewSequenceOfEncoder(cGSDTConfigLCHRestrictionR17AllowedCGListR17Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.AllowedCG_List_r17))); err != nil {
				return err
			}
			for i := range ie.AllowedCG_List_r17 {
				if err := ie.AllowedCG_List_r17[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *CG_SDT_ConfigLCH_Restriction_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cGSDTConfigLCHRestrictionR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. logicalChannelIdentity-r17: ref
	{
		if err := ie.LogicalChannelIdentity_r17.Decode(d); err != nil {
			return err
		}
	}

	// 3. configuredGrantType1Allowed-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(cGSDTConfigLCHRestrictionR17ConfiguredGrantType1AllowedR17Constraints)
			if err != nil {
				return err
			}
			ie.ConfiguredGrantType1Allowed_r17 = &idx
		}
	}

	// 4. allowedCG-List-r17: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(cGSDTConfigLCHRestrictionR17AllowedCGListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.AllowedCG_List_r17 = make([]ConfiguredGrantConfigIndexMAC_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.AllowedCG_List_r17[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
