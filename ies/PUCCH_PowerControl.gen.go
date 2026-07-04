// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PUCCH-PowerControl (line 12234).

var pUCCHPowerControlConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "deltaF-PUCCH-f0", Optional: true},
		{Name: "deltaF-PUCCH-f1", Optional: true},
		{Name: "deltaF-PUCCH-f2", Optional: true},
		{Name: "deltaF-PUCCH-f3", Optional: true},
		{Name: "deltaF-PUCCH-f4", Optional: true},
		{Name: "p0-Set", Optional: true},
		{Name: "pathlossReferenceRSs", Optional: true},
		{Name: "twoPUCCH-PC-AdjustmentStates", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var pUCCHPowerControlDeltaFPUCCHF0Constraints = per.Constrained(-16, 15)

var pUCCHPowerControlDeltaFPUCCHF1Constraints = per.Constrained(-16, 15)

var pUCCHPowerControlDeltaFPUCCHF2Constraints = per.Constrained(-16, 15)

var pUCCHPowerControlDeltaFPUCCHF3Constraints = per.Constrained(-16, 15)

var pUCCHPowerControlDeltaFPUCCHF4Constraints = per.Constrained(-16, 15)

var pUCCHPowerControlP0SetConstraints = per.SizeRange(1, common.MaxNrofPUCCH_P0_PerSet)

var pUCCHPowerControlPathlossReferenceRSsConstraints = per.SizeRange(1, common.MaxNrofPUCCH_PathlossReferenceRSs)

const (
	PUCCH_PowerControl_TwoPUCCH_PC_AdjustmentStates_TwoStates = 0
)

var pUCCHPowerControlTwoPUCCHPCAdjustmentStatesConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_PowerControl_TwoPUCCH_PC_AdjustmentStates_TwoStates},
}

var pUCCHPowerControlExtPathlossReferenceRSsV1610Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_PowerControl_Ext_PathlossReferenceRSs_v1610_Release = 0
	PUCCH_PowerControl_Ext_PathlossReferenceRSs_v1610_Setup   = 1
)

type PUCCH_PowerControl struct {
	DeltaF_PUCCH_F0              *int64
	DeltaF_PUCCH_F1              *int64
	DeltaF_PUCCH_F2              *int64
	DeltaF_PUCCH_F3              *int64
	DeltaF_PUCCH_F4              *int64
	P0_Set                       []P0_PUCCH
	PathlossReferenceRSs         []PUCCH_PathlossReferenceRS
	TwoPUCCH_PC_AdjustmentStates *int64
	PathlossReferenceRSs_v1610   *struct {
		Choice  int
		Release *struct{}
		Setup   *PathlossReferenceRSs_v1610
	}
}

func (ie *PUCCH_PowerControl) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHPowerControlConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.PathlossReferenceRSs_v1610 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.DeltaF_PUCCH_F0 != nil, ie.DeltaF_PUCCH_F1 != nil, ie.DeltaF_PUCCH_F2 != nil, ie.DeltaF_PUCCH_F3 != nil, ie.DeltaF_PUCCH_F4 != nil, ie.P0_Set != nil, ie.PathlossReferenceRSs != nil, ie.TwoPUCCH_PC_AdjustmentStates != nil}); err != nil {
		return err
	}

	// 3. deltaF-PUCCH-f0: integer
	{
		if ie.DeltaF_PUCCH_F0 != nil {
			if err := e.EncodeInteger(*ie.DeltaF_PUCCH_F0, pUCCHPowerControlDeltaFPUCCHF0Constraints); err != nil {
				return err
			}
		}
	}

	// 4. deltaF-PUCCH-f1: integer
	{
		if ie.DeltaF_PUCCH_F1 != nil {
			if err := e.EncodeInteger(*ie.DeltaF_PUCCH_F1, pUCCHPowerControlDeltaFPUCCHF1Constraints); err != nil {
				return err
			}
		}
	}

	// 5. deltaF-PUCCH-f2: integer
	{
		if ie.DeltaF_PUCCH_F2 != nil {
			if err := e.EncodeInteger(*ie.DeltaF_PUCCH_F2, pUCCHPowerControlDeltaFPUCCHF2Constraints); err != nil {
				return err
			}
		}
	}

	// 6. deltaF-PUCCH-f3: integer
	{
		if ie.DeltaF_PUCCH_F3 != nil {
			if err := e.EncodeInteger(*ie.DeltaF_PUCCH_F3, pUCCHPowerControlDeltaFPUCCHF3Constraints); err != nil {
				return err
			}
		}
	}

	// 7. deltaF-PUCCH-f4: integer
	{
		if ie.DeltaF_PUCCH_F4 != nil {
			if err := e.EncodeInteger(*ie.DeltaF_PUCCH_F4, pUCCHPowerControlDeltaFPUCCHF4Constraints); err != nil {
				return err
			}
		}
	}

	// 8. p0-Set: sequence-of
	{
		if ie.P0_Set != nil {
			seqOf := e.NewSequenceOfEncoder(pUCCHPowerControlP0SetConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.P0_Set))); err != nil {
				return err
			}
			for i := range ie.P0_Set {
				if err := ie.P0_Set[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 9. pathlossReferenceRSs: sequence-of
	{
		if ie.PathlossReferenceRSs != nil {
			seqOf := e.NewSequenceOfEncoder(pUCCHPowerControlPathlossReferenceRSsConstraints)
			if err := seqOf.EncodeLength(int64(len(ie.PathlossReferenceRSs))); err != nil {
				return err
			}
			for i := range ie.PathlossReferenceRSs {
				if err := ie.PathlossReferenceRSs[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 10. twoPUCCH-PC-AdjustmentStates: enumerated
	{
		if ie.TwoPUCCH_PC_AdjustmentStates != nil {
			if err := e.EncodeEnumerated(*ie.TwoPUCCH_PC_AdjustmentStates, pUCCHPowerControlTwoPUCCHPCAdjustmentStatesConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "pathlossReferenceRSs-v1610", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.PathlossReferenceRSs_v1610 != nil}); err != nil {
				return err
			}

			if ie.PathlossReferenceRSs_v1610 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUCCHPowerControlExtPathlossReferenceRSsV1610Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.PathlossReferenceRSs_v1610).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.PathlossReferenceRSs_v1610).Choice {
				case PUCCH_PowerControl_Ext_PathlossReferenceRSs_v1610_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUCCH_PowerControl_Ext_PathlossReferenceRSs_v1610_Setup:
					if err := (*ie.PathlossReferenceRSs_v1610).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *PUCCH_PowerControl) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHPowerControlConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. deltaF-PUCCH-f0: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(pUCCHPowerControlDeltaFPUCCHF0Constraints)
			if err != nil {
				return err
			}
			ie.DeltaF_PUCCH_F0 = &v
		}
	}

	// 4. deltaF-PUCCH-f1: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(pUCCHPowerControlDeltaFPUCCHF1Constraints)
			if err != nil {
				return err
			}
			ie.DeltaF_PUCCH_F1 = &v
		}
	}

	// 5. deltaF-PUCCH-f2: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(pUCCHPowerControlDeltaFPUCCHF2Constraints)
			if err != nil {
				return err
			}
			ie.DeltaF_PUCCH_F2 = &v
		}
	}

	// 6. deltaF-PUCCH-f3: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(pUCCHPowerControlDeltaFPUCCHF3Constraints)
			if err != nil {
				return err
			}
			ie.DeltaF_PUCCH_F3 = &v
		}
	}

	// 7. deltaF-PUCCH-f4: integer
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeInteger(pUCCHPowerControlDeltaFPUCCHF4Constraints)
			if err != nil {
				return err
			}
			ie.DeltaF_PUCCH_F4 = &v
		}
	}

	// 8. p0-Set: sequence-of
	{
		if seq.IsComponentPresent(5) {
			seqOf := d.NewSequenceOfDecoder(pUCCHPowerControlP0SetConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.P0_Set = make([]P0_PUCCH, n)
			for i := int64(0); i < n; i++ {
				if err := ie.P0_Set[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 9. pathlossReferenceRSs: sequence-of
	{
		if seq.IsComponentPresent(6) {
			seqOf := d.NewSequenceOfDecoder(pUCCHPowerControlPathlossReferenceRSsConstraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.PathlossReferenceRSs = make([]PUCCH_PathlossReferenceRS, n)
			for i := int64(0); i < n; i++ {
				if err := ie.PathlossReferenceRSs[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 10. twoPUCCH-PC-AdjustmentStates: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(pUCCHPowerControlTwoPUCCHPCAdjustmentStatesConstraints)
			if err != nil {
				return err
			}
			ie.TwoPUCCH_PC_AdjustmentStates = &idx
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "pathlossReferenceRSs-v1610", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.PathlossReferenceRSs_v1610 = &struct {
				Choice  int
				Release *struct{}
				Setup   *PathlossReferenceRSs_v1610
			}{}
			choiceDec := dx.NewChoiceDecoder(pUCCHPowerControlExtPathlossReferenceRSsV1610Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.PathlossReferenceRSs_v1610).Choice = int(index)
			switch index {
			case PUCCH_PowerControl_Ext_PathlossReferenceRSs_v1610_Release:
				(*ie.PathlossReferenceRSs_v1610).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_PowerControl_Ext_PathlossReferenceRSs_v1610_Setup:
				(*ie.PathlossReferenceRSs_v1610).Setup = new(PathlossReferenceRSs_v1610)
				if err := (*ie.PathlossReferenceRSs_v1610).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
