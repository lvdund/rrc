// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PRS-ResourceDedicatedSL-PRS-RP-r18 (line 27664).

var sLPRSResourceDedicatedSLPRSRPR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PRS-ResourceID-r18", Optional: true},
		{Name: "sl-NumberOfSymbols-r18", Optional: true},
		{Name: "sl-CombSize-r18", Optional: true},
		{Name: "sl-PRS-starting-symbol-r18", Optional: true},
		{Name: "sl-PRS-comb-offset-r18", Optional: true},
	},
}

var sLPRSResourceDedicatedSLPRSRPR18SlPRSResourceIDR18Constraints = per.Constrained(0, 11)

var sLPRSResourceDedicatedSLPRSRPR18SlNumberOfSymbolsR18Constraints = per.Constrained(1, 9)

const (
	SL_PRS_ResourceDedicatedSL_PRS_RP_r18_Sl_CombSize_r18_N2 = 0
	SL_PRS_ResourceDedicatedSL_PRS_RP_r18_Sl_CombSize_r18_N4 = 1
	SL_PRS_ResourceDedicatedSL_PRS_RP_r18_Sl_CombSize_r18_N6 = 2
)

var sLPRSResourceDedicatedSLPRSRPR18SlCombSizeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SL_PRS_ResourceDedicatedSL_PRS_RP_r18_Sl_CombSize_r18_N2, SL_PRS_ResourceDedicatedSL_PRS_RP_r18_Sl_CombSize_r18_N4, SL_PRS_ResourceDedicatedSL_PRS_RP_r18_Sl_CombSize_r18_N6},
}

var sLPRSResourceDedicatedSLPRSRPR18SlPRSStartingSymbolR18Constraints = per.Constrained(4, 12)

var sLPRSResourceDedicatedSLPRSRPR18SlPRSCombOffsetR18Constraints = per.Constrained(1, 5)

type SL_PRS_ResourceDedicatedSL_PRS_RP_r18 struct {
	Sl_PRS_ResourceID_r18      *int64
	Sl_NumberOfSymbols_r18     *int64
	Sl_CombSize_r18            *int64
	Sl_PRS_Starting_Symbol_r18 *int64
	Sl_PRS_Comb_Offset_r18     *int64
}

func (ie *SL_PRS_ResourceDedicatedSL_PRS_RP_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPRSResourceDedicatedSLPRSRPR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PRS_ResourceID_r18 != nil, ie.Sl_NumberOfSymbols_r18 != nil, ie.Sl_CombSize_r18 != nil, ie.Sl_PRS_Starting_Symbol_r18 != nil, ie.Sl_PRS_Comb_Offset_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-PRS-ResourceID-r18: integer
	{
		if ie.Sl_PRS_ResourceID_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_PRS_ResourceID_r18, sLPRSResourceDedicatedSLPRSRPR18SlPRSResourceIDR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. sl-NumberOfSymbols-r18: integer
	{
		if ie.Sl_NumberOfSymbols_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_NumberOfSymbols_r18, sLPRSResourceDedicatedSLPRSRPR18SlNumberOfSymbolsR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. sl-CombSize-r18: enumerated
	{
		if ie.Sl_CombSize_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Sl_CombSize_r18, sLPRSResourceDedicatedSLPRSRPR18SlCombSizeR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. sl-PRS-starting-symbol-r18: integer
	{
		if ie.Sl_PRS_Starting_Symbol_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_PRS_Starting_Symbol_r18, sLPRSResourceDedicatedSLPRSRPR18SlPRSStartingSymbolR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. sl-PRS-comb-offset-r18: integer
	{
		if ie.Sl_PRS_Comb_Offset_r18 != nil {
			if err := e.EncodeInteger(*ie.Sl_PRS_Comb_Offset_r18, sLPRSResourceDedicatedSLPRSRPR18SlPRSCombOffsetR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SL_PRS_ResourceDedicatedSL_PRS_RP_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPRSResourceDedicatedSLPRSRPR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-PRS-ResourceID-r18: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(sLPRSResourceDedicatedSLPRSRPR18SlPRSResourceIDR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PRS_ResourceID_r18 = &v
		}
	}

	// 3. sl-NumberOfSymbols-r18: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(sLPRSResourceDedicatedSLPRSRPR18SlNumberOfSymbolsR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_NumberOfSymbols_r18 = &v
		}
	}

	// 4. sl-CombSize-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sLPRSResourceDedicatedSLPRSRPR18SlCombSizeR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_CombSize_r18 = &idx
		}
	}

	// 5. sl-PRS-starting-symbol-r18: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(sLPRSResourceDedicatedSLPRSRPR18SlPRSStartingSymbolR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PRS_Starting_Symbol_r18 = &v
		}
	}

	// 6. sl-PRS-comb-offset-r18: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(sLPRSResourceDedicatedSLPRSRPR18SlPRSCombOffsetR18Constraints)
			if err != nil {
				return err
			}
			ie.Sl_PRS_Comb_Offset_r18 = &v
		}
	}

	return nil
}
