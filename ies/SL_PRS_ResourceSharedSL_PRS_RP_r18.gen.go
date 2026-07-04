// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-PRS-ResourceSharedSL-PRS-RP-r18 (line 28122).

var sLPRSResourceSharedSLPRSRPR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PRS-ResourceID-r18"},
		{Name: "mNumberOfSymbols-r18"},
		{Name: "sl-PRS-CombSizeN-AndReOffset-r18", Optional: true},
	},
}

var sLPRSResourceSharedSLPRSRPR18SlPRSResourceIDR18Constraints = per.Constrained(0, 16)

var sLPRSResourceSharedSLPRSRPR18MNumberOfSymbolsR18Constraints = per.Constrained(1, 9)

var sL_PRS_ResourceSharedSL_PRS_RP_r18SlPRSCombSizeNAndReOffsetR18Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "n2-r18"},
		{Name: "n4-r18"},
		{Name: "dummy1"},
	},
}

const (
	SL_PRS_ResourceSharedSL_PRS_RP_r18_Sl_PRS_CombSizeN_AndReOffset_r18_N2_r18 = 0
	SL_PRS_ResourceSharedSL_PRS_RP_r18_Sl_PRS_CombSizeN_AndReOffset_r18_N4_r18 = 1
	SL_PRS_ResourceSharedSL_PRS_RP_r18_Sl_PRS_CombSizeN_AndReOffset_r18_Dummy1 = 2
)

type SL_PRS_ResourceSharedSL_PRS_RP_r18 struct {
	Sl_PRS_ResourceID_r18            int64
	MNumberOfSymbols_r18             int64
	Sl_PRS_CombSizeN_AndReOffset_r18 *struct {
		Choice int
		N2_r18 *int64
		N4_r18 *int64
		Dummy1 *int64
	}
}

func (ie *SL_PRS_ResourceSharedSL_PRS_RP_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLPRSResourceSharedSLPRSRPR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Sl_PRS_CombSizeN_AndReOffset_r18 != nil}); err != nil {
		return err
	}

	// 2. sl-PRS-ResourceID-r18: integer
	{
		if err := e.EncodeInteger(ie.Sl_PRS_ResourceID_r18, sLPRSResourceSharedSLPRSRPR18SlPRSResourceIDR18Constraints); err != nil {
			return err
		}
	}

	// 3. mNumberOfSymbols-r18: integer
	{
		if err := e.EncodeInteger(ie.MNumberOfSymbols_r18, sLPRSResourceSharedSLPRSRPR18MNumberOfSymbolsR18Constraints); err != nil {
			return err
		}
	}

	// 4. sl-PRS-CombSizeN-AndReOffset-r18: choice
	{
		if ie.Sl_PRS_CombSizeN_AndReOffset_r18 != nil {
			choiceEnc := e.NewChoiceEncoder(sL_PRS_ResourceSharedSL_PRS_RP_r18SlPRSCombSizeNAndReOffsetR18Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Sl_PRS_CombSizeN_AndReOffset_r18).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Sl_PRS_CombSizeN_AndReOffset_r18).Choice {
			case SL_PRS_ResourceSharedSL_PRS_RP_r18_Sl_PRS_CombSizeN_AndReOffset_r18_N2_r18:
				if err := e.EncodeInteger((*(*ie.Sl_PRS_CombSizeN_AndReOffset_r18).N2_r18), per.Constrained(0, 1)); err != nil {
					return err
				}
			case SL_PRS_ResourceSharedSL_PRS_RP_r18_Sl_PRS_CombSizeN_AndReOffset_r18_N4_r18:
				if err := e.EncodeInteger((*(*ie.Sl_PRS_CombSizeN_AndReOffset_r18).N4_r18), per.Constrained(0, 3)); err != nil {
					return err
				}
			case SL_PRS_ResourceSharedSL_PRS_RP_r18_Sl_PRS_CombSizeN_AndReOffset_r18_Dummy1:
				if err := e.EncodeInteger((*(*ie.Sl_PRS_CombSizeN_AndReOffset_r18).Dummy1), per.Constrained(0, 5)); err != nil {
					return err
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Sl_PRS_CombSizeN_AndReOffset_r18).Choice), Constraint: "index out of range"}
			}
		}
	}

	return nil
}

func (ie *SL_PRS_ResourceSharedSL_PRS_RP_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLPRSResourceSharedSLPRSRPR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. sl-PRS-ResourceID-r18: integer
	{
		v0, err := d.DecodeInteger(sLPRSResourceSharedSLPRSRPR18SlPRSResourceIDR18Constraints)
		if err != nil {
			return err
		}
		ie.Sl_PRS_ResourceID_r18 = v0
	}

	// 3. mNumberOfSymbols-r18: integer
	{
		v1, err := d.DecodeInteger(sLPRSResourceSharedSLPRSRPR18MNumberOfSymbolsR18Constraints)
		if err != nil {
			return err
		}
		ie.MNumberOfSymbols_r18 = v1
	}

	// 4. sl-PRS-CombSizeN-AndReOffset-r18: choice
	{
		if seq.IsComponentPresent(2) {
			ie.Sl_PRS_CombSizeN_AndReOffset_r18 = &struct {
				Choice int
				N2_r18 *int64
				N4_r18 *int64
				Dummy1 *int64
			}{}
			choiceDec := d.NewChoiceDecoder(sL_PRS_ResourceSharedSL_PRS_RP_r18SlPRSCombSizeNAndReOffsetR18Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Sl_PRS_CombSizeN_AndReOffset_r18).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case SL_PRS_ResourceSharedSL_PRS_RP_r18_Sl_PRS_CombSizeN_AndReOffset_r18_N2_r18:
				(*ie.Sl_PRS_CombSizeN_AndReOffset_r18).N2_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 1))
				if err != nil {
					return err
				}
				(*(*ie.Sl_PRS_CombSizeN_AndReOffset_r18).N2_r18) = v
			case SL_PRS_ResourceSharedSL_PRS_RP_r18_Sl_PRS_CombSizeN_AndReOffset_r18_N4_r18:
				(*ie.Sl_PRS_CombSizeN_AndReOffset_r18).N4_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 3))
				if err != nil {
					return err
				}
				(*(*ie.Sl_PRS_CombSizeN_AndReOffset_r18).N4_r18) = v
			case SL_PRS_ResourceSharedSL_PRS_RP_r18_Sl_PRS_CombSizeN_AndReOffset_r18_Dummy1:
				(*ie.Sl_PRS_CombSizeN_AndReOffset_r18).Dummy1 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 5))
				if err != nil {
					return err
				}
				(*(*ie.Sl_PRS_CombSizeN_AndReOffset_r18).Dummy1) = v
			}
		}
	}

	return nil
}
