// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PosBWP-ConfigCommon-r18 (line 27607).

var sLPosBWPConfigCommonR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-BWP-Generic-r18", Optional: true},
		{Name: "sl-BWP-PRS-PoolConfigCommon-r18", Optional: true},
	},
}

type SL_PosBWP_ConfigCommon_r18 struct {
	Sl_BWP_Generic_r18              *SL_BWP_Generic_r16
	Sl_BWP_PRS_PoolConfigCommon_r18 *SL_BWP_PRS_PoolConfigCommon_r18
}

func (ie *SL_PosBWP_ConfigCommon_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPosBWPConfigCommonR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_BWP_Generic_r18 != nil, ie.Sl_BWP_PRS_PoolConfigCommon_r18 != nil}); err != nil {
		return err
	}

	// 3. sl-BWP-Generic-r18: ref
	{
		if ie.Sl_BWP_Generic_r18 != nil {
			if err := ie.Sl_BWP_Generic_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. sl-BWP-PRS-PoolConfigCommon-r18: ref
	{
		if ie.Sl_BWP_PRS_PoolConfigCommon_r18 != nil {
			if err := ie.Sl_BWP_PRS_PoolConfigCommon_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_PosBWP_ConfigCommon_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPosBWPConfigCommonR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-BWP-Generic-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Sl_BWP_Generic_r18 = new(SL_BWP_Generic_r16)
			if err := ie.Sl_BWP_Generic_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. sl-BWP-PRS-PoolConfigCommon-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Sl_BWP_PRS_PoolConfigCommon_r18 = new(SL_BWP_PRS_PoolConfigCommon_r18)
			if err := ie.Sl_BWP_PRS_PoolConfigCommon_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
