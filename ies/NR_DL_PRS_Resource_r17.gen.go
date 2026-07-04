// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NR-DL-PRS-Resource-r17 (line 10702).

var nRDLPRSResourceR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "nr-DL-PRS-ResourceID-r17"},
		{Name: "dl-PRS-SequenceID-r17"},
		{Name: "dl-PRS-CombSizeN-AndReOffset-r17"},
		{Name: "dl-PRS-ResourceSlotOffset-r17"},
		{Name: "dl-PRS-ResourceSymbolOffset-r17"},
		{Name: "dl-PRS-QCL-Info-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var nRDLPRSResourceR17DlPRSSequenceIDR17Constraints = per.Constrained(0, 4095)

var nR_DL_PRS_Resource_r17DlPRSCombSizeNAndReOffsetR17Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "n2-r17"},
		{Name: "n4-r17"},
		{Name: "n6-r17"},
		{Name: "n12-r17"},
	},
}

const (
	NR_DL_PRS_Resource_r17_Dl_PRS_CombSizeN_AndReOffset_r17_N2_r17  = 0
	NR_DL_PRS_Resource_r17_Dl_PRS_CombSizeN_AndReOffset_r17_N4_r17  = 1
	NR_DL_PRS_Resource_r17_Dl_PRS_CombSizeN_AndReOffset_r17_N6_r17  = 2
	NR_DL_PRS_Resource_r17_Dl_PRS_CombSizeN_AndReOffset_r17_N12_r17 = 3
)

var nRDLPRSResourceR17DlPRSResourceSlotOffsetR17Constraints = per.Constrained(0, common.MaxNrofPRS_ResourceOffsetValue_1_r17)

var nRDLPRSResourceR17DlPRSResourceSymbolOffsetR17Constraints = per.Constrained(0, 12)

var nRDLPRSResourceR17DlPRSResourceSymbolOffsetV1800Constraints = per.Constrained(13, 13)

type NR_DL_PRS_Resource_r17 struct {
	Nr_DL_PRS_ResourceID_r17         NR_DL_PRS_ResourceID_r17
	Dl_PRS_SequenceID_r17            int64
	Dl_PRS_CombSizeN_AndReOffset_r17 struct {
		Choice  int
		N2_r17  *int64
		N4_r17  *int64
		N6_r17  *int64
		N12_r17 *int64
	}
	Dl_PRS_ResourceSlotOffset_r17     int64
	Dl_PRS_ResourceSymbolOffset_r17   int64
	Dl_PRS_QCL_Info_r17               *DL_PRS_QCL_Info_r17
	Dl_PRS_ResourceSymbolOffset_v1800 *int64
}

func (ie *NR_DL_PRS_Resource_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nRDLPRSResourceR17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Dl_PRS_ResourceSymbolOffset_v1800 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Dl_PRS_QCL_Info_r17 != nil}); err != nil {
		return err
	}

	// 3. nr-DL-PRS-ResourceID-r17: ref
	{
		if err := ie.Nr_DL_PRS_ResourceID_r17.Encode(e); err != nil {
			return err
		}
	}

	// 4. dl-PRS-SequenceID-r17: integer
	{
		if err := e.EncodeInteger(ie.Dl_PRS_SequenceID_r17, nRDLPRSResourceR17DlPRSSequenceIDR17Constraints); err != nil {
			return err
		}
	}

	// 5. dl-PRS-CombSizeN-AndReOffset-r17: choice
	{
		choiceEnc := e.NewChoiceEncoder(nR_DL_PRS_Resource_r17DlPRSCombSizeNAndReOffsetR17Constraints)
		if err := choiceEnc.EncodeChoice(int64(ie.Dl_PRS_CombSizeN_AndReOffset_r17.Choice), false, nil); err != nil {
			return err
		}
		switch ie.Dl_PRS_CombSizeN_AndReOffset_r17.Choice {
		case NR_DL_PRS_Resource_r17_Dl_PRS_CombSizeN_AndReOffset_r17_N2_r17:
			if err := e.EncodeInteger((*ie.Dl_PRS_CombSizeN_AndReOffset_r17.N2_r17), per.Constrained(0, 1)); err != nil {
				return err
			}
		case NR_DL_PRS_Resource_r17_Dl_PRS_CombSizeN_AndReOffset_r17_N4_r17:
			if err := e.EncodeInteger((*ie.Dl_PRS_CombSizeN_AndReOffset_r17.N4_r17), per.Constrained(0, 3)); err != nil {
				return err
			}
		case NR_DL_PRS_Resource_r17_Dl_PRS_CombSizeN_AndReOffset_r17_N6_r17:
			if err := e.EncodeInteger((*ie.Dl_PRS_CombSizeN_AndReOffset_r17.N6_r17), per.Constrained(0, 5)); err != nil {
				return err
			}
		case NR_DL_PRS_Resource_r17_Dl_PRS_CombSizeN_AndReOffset_r17_N12_r17:
			if err := e.EncodeInteger((*ie.Dl_PRS_CombSizeN_AndReOffset_r17.N12_r17), per.Constrained(0, 11)); err != nil {
				return err
			}
		default:
			return &per.ConstraintViolationError{TypeName: "choice", Value: int64(ie.Dl_PRS_CombSizeN_AndReOffset_r17.Choice), Constraint: "index out of range"}
		}
	}

	// 6. dl-PRS-ResourceSlotOffset-r17: integer
	{
		if err := e.EncodeInteger(ie.Dl_PRS_ResourceSlotOffset_r17, nRDLPRSResourceR17DlPRSResourceSlotOffsetR17Constraints); err != nil {
			return err
		}
	}

	// 7. dl-PRS-ResourceSymbolOffset-r17: integer
	{
		if err := e.EncodeInteger(ie.Dl_PRS_ResourceSymbolOffset_r17, nRDLPRSResourceR17DlPRSResourceSymbolOffsetR17Constraints); err != nil {
			return err
		}
	}

	// 8. dl-PRS-QCL-Info-r17: ref
	{
		if ie.Dl_PRS_QCL_Info_r17 != nil {
			if err := ie.Dl_PRS_QCL_Info_r17.Encode(e); err != nil {
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
					{Name: "dl-PRS-ResourceSymbolOffset-v1800", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Dl_PRS_ResourceSymbolOffset_v1800 != nil}); err != nil {
				return err
			}

			if ie.Dl_PRS_ResourceSymbolOffset_v1800 != nil {
				if err := ex.EncodeInteger(*ie.Dl_PRS_ResourceSymbolOffset_v1800, nRDLPRSResourceR17DlPRSResourceSymbolOffsetV1800Constraints); err != nil {
					return err
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

func (ie *NR_DL_PRS_Resource_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nRDLPRSResourceR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. nr-DL-PRS-ResourceID-r17: ref
	{
		if err := ie.Nr_DL_PRS_ResourceID_r17.Decode(d); err != nil {
			return err
		}
	}

	// 4. dl-PRS-SequenceID-r17: integer
	{
		v1, err := d.DecodeInteger(nRDLPRSResourceR17DlPRSSequenceIDR17Constraints)
		if err != nil {
			return err
		}
		ie.Dl_PRS_SequenceID_r17 = v1
	}

	// 5. dl-PRS-CombSizeN-AndReOffset-r17: choice
	{
		choiceDec := d.NewChoiceDecoder(nR_DL_PRS_Resource_r17DlPRSCombSizeNAndReOffsetR17Constraints)
		index, isExt, _, err := choiceDec.DecodeChoice()
		if err != nil {
			return err
		}
		ie.Dl_PRS_CombSizeN_AndReOffset_r17.Choice = int(index)
		if isExt {
			return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
		}
		switch index {
		case NR_DL_PRS_Resource_r17_Dl_PRS_CombSizeN_AndReOffset_r17_N2_r17:
			ie.Dl_PRS_CombSizeN_AndReOffset_r17.N2_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 1))
			if err != nil {
				return err
			}
			(*ie.Dl_PRS_CombSizeN_AndReOffset_r17.N2_r17) = v
		case NR_DL_PRS_Resource_r17_Dl_PRS_CombSizeN_AndReOffset_r17_N4_r17:
			ie.Dl_PRS_CombSizeN_AndReOffset_r17.N4_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 3))
			if err != nil {
				return err
			}
			(*ie.Dl_PRS_CombSizeN_AndReOffset_r17.N4_r17) = v
		case NR_DL_PRS_Resource_r17_Dl_PRS_CombSizeN_AndReOffset_r17_N6_r17:
			ie.Dl_PRS_CombSizeN_AndReOffset_r17.N6_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 5))
			if err != nil {
				return err
			}
			(*ie.Dl_PRS_CombSizeN_AndReOffset_r17.N6_r17) = v
		case NR_DL_PRS_Resource_r17_Dl_PRS_CombSizeN_AndReOffset_r17_N12_r17:
			ie.Dl_PRS_CombSizeN_AndReOffset_r17.N12_r17 = new(int64)
			v, err := d.DecodeInteger(per.Constrained(0, 11))
			if err != nil {
				return err
			}
			(*ie.Dl_PRS_CombSizeN_AndReOffset_r17.N12_r17) = v
		}
	}

	// 6. dl-PRS-ResourceSlotOffset-r17: integer
	{
		v3, err := d.DecodeInteger(nRDLPRSResourceR17DlPRSResourceSlotOffsetR17Constraints)
		if err != nil {
			return err
		}
		ie.Dl_PRS_ResourceSlotOffset_r17 = v3
	}

	// 7. dl-PRS-ResourceSymbolOffset-r17: integer
	{
		v4, err := d.DecodeInteger(nRDLPRSResourceR17DlPRSResourceSymbolOffsetR17Constraints)
		if err != nil {
			return err
		}
		ie.Dl_PRS_ResourceSymbolOffset_r17 = v4
	}

	// 8. dl-PRS-QCL-Info-r17: ref
	{
		if seq.IsComponentPresent(5) {
			ie.Dl_PRS_QCL_Info_r17 = new(DL_PRS_QCL_Info_r17)
			if err := ie.Dl_PRS_QCL_Info_r17.Decode(d); err != nil {
				return err
			}
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
				{Name: "dl-PRS-ResourceSymbolOffset-v1800", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(nRDLPRSResourceR17DlPRSResourceSymbolOffsetV1800Constraints)
			if err != nil {
				return err
			}
			ie.Dl_PRS_ResourceSymbolOffset_v1800 = &v
		}
	}

	return nil
}
