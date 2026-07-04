// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-ResourceExt-v1610 (line 12086).

var pUCCHResourceExtV1610Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "interlaceAllocation-r16", Optional: true},
		{Name: "format-v1610", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var pUCCH_ResourceExt_v1610FormatV1610Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "interlace1-v1610"},
		{Name: "occ-v1610"},
	},
}

const (
	PUCCH_ResourceExt_v1610_Format_v1610_Interlace1_v1610 = 0
	PUCCH_ResourceExt_v1610_Format_v1610_Occ_v1610        = 1
)

const (
	PUCCH_ResourceExt_v1610_Ext_Pucch_RepetitionNrofSlots_r17_N1 = 0
	PUCCH_ResourceExt_v1610_Ext_Pucch_RepetitionNrofSlots_r17_N2 = 1
	PUCCH_ResourceExt_v1610_Ext_Pucch_RepetitionNrofSlots_r17_N4 = 2
	PUCCH_ResourceExt_v1610_Ext_Pucch_RepetitionNrofSlots_r17_N8 = 3
)

var pUCCHResourceExtV1610ExtPucchRepetitionNrofSlotsR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_ResourceExt_v1610_Ext_Pucch_RepetitionNrofSlots_r17_N1, PUCCH_ResourceExt_v1610_Ext_Pucch_RepetitionNrofSlots_r17_N2, PUCCH_ResourceExt_v1610_Ext_Pucch_RepetitionNrofSlots_r17_N4, PUCCH_ResourceExt_v1610_Ext_Pucch_RepetitionNrofSlots_r17_N8},
}

const (
	PUCCH_ResourceExt_v1610_Ext_ApplyIndicatedTCI_State_r18_First  = 0
	PUCCH_ResourceExt_v1610_Ext_ApplyIndicatedTCI_State_r18_Second = 1
	PUCCH_ResourceExt_v1610_Ext_ApplyIndicatedTCI_State_r18_Both   = 2
	PUCCH_ResourceExt_v1610_Ext_ApplyIndicatedTCI_State_r18_Spare1 = 3
)

var pUCCHResourceExtV1610ExtApplyIndicatedTCIStateR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_ResourceExt_v1610_Ext_ApplyIndicatedTCI_State_r18_First, PUCCH_ResourceExt_v1610_Ext_ApplyIndicatedTCI_State_r18_Second, PUCCH_ResourceExt_v1610_Ext_ApplyIndicatedTCI_State_r18_Both, PUCCH_ResourceExt_v1610_Ext_ApplyIndicatedTCI_State_r18_Spare1},
}

const (
	PUCCH_ResourceExt_v1610_Ext_MultipanelSFN_Scheme_r18_Enabled = 0
)

var pUCCHResourceExtV1610ExtMultipanelSFNSchemeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_ResourceExt_v1610_Ext_MultipanelSFN_Scheme_r18_Enabled},
}

var pUCCHResourceExtV1610ExtDlDataToULACKR18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_ResourceExt_v1610_Ext_Dl_DataToUL_ACK_r18_Release = 0
	PUCCH_ResourceExt_v1610_Ext_Dl_DataToUL_ACK_r18_Setup   = 1
)

var pUCCHResourceExtV1610ExtDlDataToULACKDCI12R18Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "release"},
		{Name: "setup"},
	},
}

const (
	PUCCH_ResourceExt_v1610_Ext_Dl_DataToUL_ACK_DCI_1_2_r18_Release = 0
	PUCCH_ResourceExt_v1610_Ext_Dl_DataToUL_ACK_DCI_1_2_r18_Setup   = 1
)

var pUCCHResourceExtV1610InterlaceAllocationR16Interlace0R16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "scs15"},
		{Name: "scs30"},
	},
}

const (
	PUCCH_ResourceExt_v1610_InterlaceAllocation_r16_Interlace0_r16_Scs15 = 0
	PUCCH_ResourceExt_v1610_InterlaceAllocation_r16_Interlace0_r16_Scs30 = 1
)

var pUCCHResourceExtV1610FormatV1610OccV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "occ-Length-v1610", Optional: true},
		{Name: "occ-Index-v1610", Optional: true},
	},
}

const (
	PUCCH_ResourceExt_v1610_Format_v1610_Occ_v1610_Occ_Length_v1610_N2 = 0
	PUCCH_ResourceExt_v1610_Format_v1610_Occ_v1610_Occ_Length_v1610_N4 = 1
)

var pUCCHResourceExtV1610FormatV1610OccV1610OccLengthV1610Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_ResourceExt_v1610_Format_v1610_Occ_v1610_Occ_Length_v1610_N2, PUCCH_ResourceExt_v1610_Format_v1610_Occ_v1610_Occ_Length_v1610_N4},
}

const (
	PUCCH_ResourceExt_v1610_Format_v1610_Occ_v1610_Occ_Index_v1610_N0 = 0
	PUCCH_ResourceExt_v1610_Format_v1610_Occ_v1610_Occ_Index_v1610_N1 = 1
	PUCCH_ResourceExt_v1610_Format_v1610_Occ_v1610_Occ_Index_v1610_N2 = 2
	PUCCH_ResourceExt_v1610_Format_v1610_Occ_v1610_Occ_Index_v1610_N3 = 3
)

var pUCCHResourceExtV1610FormatV1610OccV1610OccIndexV1610Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_ResourceExt_v1610_Format_v1610_Occ_v1610_Occ_Index_v1610_N0, PUCCH_ResourceExt_v1610_Format_v1610_Occ_v1610_Occ_Index_v1610_N1, PUCCH_ResourceExt_v1610_Format_v1610_Occ_v1610_Occ_Index_v1610_N2, PUCCH_ResourceExt_v1610_Format_v1610_Occ_v1610_Occ_Index_v1610_N3},
}

type PUCCH_ResourceExt_v1610 struct {
	InterlaceAllocation_r16 *struct {
		Rb_SetIndex_r16 int64
		Interlace0_r16  struct {
			Choice int
			Scs15  *int64
			Scs30  *int64
		}
	}
	Format_v1610 *struct {
		Choice           int
		Interlace1_v1610 *int64
		Occ_v1610        *struct {
			Occ_Length_v1610 *int64
			Occ_Index_v1610  *int64
		}
	}
	Format_v1700                  *struct{ NrofPRBs_r17 int64 }
	Pucch_RepetitionNrofSlots_r17 *int64
	ApplyIndicatedTCI_State_r18   *int64
	MultipanelSFN_Scheme_r18      *int64
	Dl_DataToUL_ACK_r18           *struct {
		Choice  int
		Release *struct{}
		Setup   *DL_DataToUL_ACK_r18
	}
	Dl_DataToUL_ACK_DCI_1_2_r18 *struct {
		Choice  int
		Release *struct{}
		Setup   *DL_DataToUL_ACK_DCI_1_2_r18
	}
}

func (ie *PUCCH_ResourceExt_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHResourceExtV1610Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Format_v1700 != nil || ie.Pucch_RepetitionNrofSlots_r17 != nil
	hasExtGroup1 := ie.ApplyIndicatedTCI_State_r18 != nil || ie.MultipanelSFN_Scheme_r18 != nil || ie.Dl_DataToUL_ACK_r18 != nil || ie.Dl_DataToUL_ACK_DCI_1_2_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.InterlaceAllocation_r16 != nil, ie.Format_v1610 != nil}); err != nil {
		return err
	}

	// 3. interlaceAllocation-r16: sequence
	{
		if ie.InterlaceAllocation_r16 != nil {
			c := ie.InterlaceAllocation_r16
			if err := e.EncodeInteger(c.Rb_SetIndex_r16, per.Constrained(0, 4)); err != nil {
				return err
			}
			{
				choiceEnc := e.NewChoiceEncoder(pUCCHResourceExtV1610InterlaceAllocationR16Interlace0R16Constraints)
				if err := choiceEnc.EncodeChoice(int64(c.Interlace0_r16.Choice), false, nil); err != nil {
					return err
				}
				switch c.Interlace0_r16.Choice {
				case PUCCH_ResourceExt_v1610_InterlaceAllocation_r16_Interlace0_r16_Scs15:
					if err := e.EncodeInteger((*c.Interlace0_r16.Scs15), per.Constrained(0, 9)); err != nil {
						return err
					}
				case PUCCH_ResourceExt_v1610_InterlaceAllocation_r16_Interlace0_r16_Scs30:
					if err := e.EncodeInteger((*c.Interlace0_r16.Scs30), per.Constrained(0, 4)); err != nil {
						return err
					}
				}
			}
		}
	}

	// 4. format-v1610: choice
	{
		if ie.Format_v1610 != nil {
			choiceEnc := e.NewChoiceEncoder(pUCCH_ResourceExt_v1610FormatV1610Constraints)
			if err := choiceEnc.EncodeChoice(int64((*ie.Format_v1610).Choice), false, nil); err != nil {
				return err
			}
			switch (*ie.Format_v1610).Choice {
			case PUCCH_ResourceExt_v1610_Format_v1610_Interlace1_v1610:
				if err := e.EncodeInteger((*(*ie.Format_v1610).Interlace1_v1610), per.Constrained(0, 9)); err != nil {
					return err
				}
			case PUCCH_ResourceExt_v1610_Format_v1610_Occ_v1610:
				c := (*(*ie.Format_v1610).Occ_v1610)
				pUCCHResourceExtV1610FormatV1610OccV1610Seq := e.NewSequenceEncoder(pUCCHResourceExtV1610FormatV1610OccV1610Constraints)
				if err := pUCCHResourceExtV1610FormatV1610OccV1610Seq.EncodePreamble([]bool{c.Occ_Length_v1610 != nil, c.Occ_Index_v1610 != nil}); err != nil {
					return err
				}
				if c.Occ_Length_v1610 != nil {
					if err := e.EncodeEnumerated((*c.Occ_Length_v1610), pUCCHResourceExtV1610FormatV1610OccV1610OccLengthV1610Constraints); err != nil {
						return err
					}
				}
				if c.Occ_Index_v1610 != nil {
					if err := e.EncodeEnumerated((*c.Occ_Index_v1610), pUCCHResourceExtV1610FormatV1610OccV1610OccIndexV1610Constraints); err != nil {
						return err
					}
				}
			default:
				return &per.ConstraintViolationError{TypeName: "choice", Value: int64((*ie.Format_v1610).Choice), Constraint: "index out of range"}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "format-v1700", Optional: true},
					{Name: "pucch-RepetitionNrofSlots-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Format_v1700 != nil, ie.Pucch_RepetitionNrofSlots_r17 != nil}); err != nil {
				return err
			}

			if ie.Format_v1700 != nil {
				c := ie.Format_v1700
				if err := ex.EncodeInteger(c.NrofPRBs_r17, per.Constrained(1, 16)); err != nil {
					return err
				}
			}

			if ie.Pucch_RepetitionNrofSlots_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.Pucch_RepetitionNrofSlots_r17, pUCCHResourceExtV1610ExtPucchRepetitionNrofSlotsR17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "applyIndicatedTCI-State-r18", Optional: true},
					{Name: "multipanelSFN-Scheme-r18", Optional: true},
					{Name: "dl-DataToUL-ACK-r18", Optional: true},
					{Name: "dl-DataToUL-ACK-DCI-1-2-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ApplyIndicatedTCI_State_r18 != nil, ie.MultipanelSFN_Scheme_r18 != nil, ie.Dl_DataToUL_ACK_r18 != nil, ie.Dl_DataToUL_ACK_DCI_1_2_r18 != nil}); err != nil {
				return err
			}

			if ie.ApplyIndicatedTCI_State_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.ApplyIndicatedTCI_State_r18, pUCCHResourceExtV1610ExtApplyIndicatedTCIStateR18Constraints); err != nil {
					return err
				}
			}

			if ie.MultipanelSFN_Scheme_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.MultipanelSFN_Scheme_r18, pUCCHResourceExtV1610ExtMultipanelSFNSchemeR18Constraints); err != nil {
					return err
				}
			}

			if ie.Dl_DataToUL_ACK_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUCCHResourceExtV1610ExtDlDataToULACKR18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Dl_DataToUL_ACK_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Dl_DataToUL_ACK_r18).Choice {
				case PUCCH_ResourceExt_v1610_Ext_Dl_DataToUL_ACK_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUCCH_ResourceExt_v1610_Ext_Dl_DataToUL_ACK_r18_Setup:
					if err := (*ie.Dl_DataToUL_ACK_r18).Setup.Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.Dl_DataToUL_ACK_DCI_1_2_r18 != nil {
				choiceEnc := ex.NewChoiceEncoder(pUCCHResourceExtV1610ExtDlDataToULACKDCI12R18Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.Dl_DataToUL_ACK_DCI_1_2_r18).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.Dl_DataToUL_ACK_DCI_1_2_r18).Choice {
				case PUCCH_ResourceExt_v1610_Ext_Dl_DataToUL_ACK_DCI_1_2_r18_Release:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case PUCCH_ResourceExt_v1610_Ext_Dl_DataToUL_ACK_DCI_1_2_r18_Setup:
					if err := (*ie.Dl_DataToUL_ACK_DCI_1_2_r18).Setup.Encode(ex); err != nil {
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

func (ie *PUCCH_ResourceExt_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHResourceExtV1610Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. interlaceAllocation-r16: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.InterlaceAllocation_r16 = &struct {
				Rb_SetIndex_r16 int64
				Interlace0_r16  struct {
					Choice int
					Scs15  *int64
					Scs30  *int64
				}
			}{}
			c := ie.InterlaceAllocation_r16
			{
				v, err := d.DecodeInteger(per.Constrained(0, 4))
				if err != nil {
					return err
				}
				c.Rb_SetIndex_r16 = v
			}
			{
				choiceDec := d.NewChoiceDecoder(pUCCHResourceExtV1610InterlaceAllocationR16Interlace0R16Constraints)
				index, _, _, err := choiceDec.DecodeChoice()
				if err != nil {
					return err
				}
				c.Interlace0_r16.Choice = int(index)
				switch index {
				case PUCCH_ResourceExt_v1610_InterlaceAllocation_r16_Interlace0_r16_Scs15:
					c.Interlace0_r16.Scs15 = new(int64)
					v, err := d.DecodeInteger(per.Constrained(0, 9))
					if err != nil {
						return err
					}
					(*c.Interlace0_r16.Scs15) = v
				case PUCCH_ResourceExt_v1610_InterlaceAllocation_r16_Interlace0_r16_Scs30:
					c.Interlace0_r16.Scs30 = new(int64)
					v, err := d.DecodeInteger(per.Constrained(0, 4))
					if err != nil {
						return err
					}
					(*c.Interlace0_r16.Scs30) = v
				}
			}
		}
	}

	// 4. format-v1610: choice
	{
		if seq.IsComponentPresent(1) {
			ie.Format_v1610 = &struct {
				Choice           int
				Interlace1_v1610 *int64
				Occ_v1610        *struct {
					Occ_Length_v1610 *int64
					Occ_Index_v1610  *int64
				}
			}{}
			choiceDec := d.NewChoiceDecoder(pUCCH_ResourceExt_v1610FormatV1610Constraints)
			index, isExt, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Format_v1610).Choice = int(index)
			if isExt {
				return &per.DecodeError{TypeName: "choice", Reason: "extension alternative not supported"}
			}
			switch index {
			case PUCCH_ResourceExt_v1610_Format_v1610_Interlace1_v1610:
				(*ie.Format_v1610).Interlace1_v1610 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(0, 9))
				if err != nil {
					return err
				}
				(*(*ie.Format_v1610).Interlace1_v1610) = v
			case PUCCH_ResourceExt_v1610_Format_v1610_Occ_v1610:
				(*ie.Format_v1610).Occ_v1610 = &struct {
					Occ_Length_v1610 *int64
					Occ_Index_v1610  *int64
				}{}
				c := (*(*ie.Format_v1610).Occ_v1610)
				pUCCHResourceExtV1610FormatV1610OccV1610Seq := d.NewSequenceDecoder(pUCCHResourceExtV1610FormatV1610OccV1610Constraints)
				if err := pUCCHResourceExtV1610FormatV1610OccV1610Seq.DecodePreamble(); err != nil {
					return err
				}
				if pUCCHResourceExtV1610FormatV1610OccV1610Seq.IsComponentPresent(0) {
					c.Occ_Length_v1610 = new(int64)
					v, err := d.DecodeEnumerated(pUCCHResourceExtV1610FormatV1610OccV1610OccLengthV1610Constraints)
					if err != nil {
						return err
					}
					(*c.Occ_Length_v1610) = v
				}
				if pUCCHResourceExtV1610FormatV1610OccV1610Seq.IsComponentPresent(1) {
					c.Occ_Index_v1610 = new(int64)
					v, err := d.DecodeEnumerated(pUCCHResourceExtV1610FormatV1610OccV1610OccIndexV1610Constraints)
					if err != nil {
						return err
					}
					(*c.Occ_Index_v1610) = v
				}
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
				{Name: "format-v1700", Optional: true},
				{Name: "pucch-RepetitionNrofSlots-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.Format_v1700 = &struct{ NrofPRBs_r17 int64 }{}
			c := ie.Format_v1700
			{
				v, err := dx.DecodeInteger(per.Constrained(1, 16))
				if err != nil {
					return err
				}
				c.NrofPRBs_r17 = v
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(pUCCHResourceExtV1610ExtPucchRepetitionNrofSlotsR17Constraints)
			if err != nil {
				return err
			}
			ie.Pucch_RepetitionNrofSlots_r17 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "applyIndicatedTCI-State-r18", Optional: true},
				{Name: "multipanelSFN-Scheme-r18", Optional: true},
				{Name: "dl-DataToUL-ACK-r18", Optional: true},
				{Name: "dl-DataToUL-ACK-DCI-1-2-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pUCCHResourceExtV1610ExtApplyIndicatedTCIStateR18Constraints)
			if err != nil {
				return err
			}
			ie.ApplyIndicatedTCI_State_r18 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(pUCCHResourceExtV1610ExtMultipanelSFNSchemeR18Constraints)
			if err != nil {
				return err
			}
			ie.MultipanelSFN_Scheme_r18 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			ie.Dl_DataToUL_ACK_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DL_DataToUL_ACK_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(pUCCHResourceExtV1610ExtDlDataToULACKR18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dl_DataToUL_ACK_r18).Choice = int(index)
			switch index {
			case PUCCH_ResourceExt_v1610_Ext_Dl_DataToUL_ACK_r18_Release:
				(*ie.Dl_DataToUL_ACK_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_ResourceExt_v1610_Ext_Dl_DataToUL_ACK_r18_Setup:
				(*ie.Dl_DataToUL_ACK_r18).Setup = new(DL_DataToUL_ACK_r18)
				if err := (*ie.Dl_DataToUL_ACK_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(3) {
			ie.Dl_DataToUL_ACK_DCI_1_2_r18 = &struct {
				Choice  int
				Release *struct{}
				Setup   *DL_DataToUL_ACK_DCI_1_2_r18
			}{}
			choiceDec := dx.NewChoiceDecoder(pUCCHResourceExtV1610ExtDlDataToULACKDCI12R18Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.Dl_DataToUL_ACK_DCI_1_2_r18).Choice = int(index)
			switch index {
			case PUCCH_ResourceExt_v1610_Ext_Dl_DataToUL_ACK_DCI_1_2_r18_Release:
				(*ie.Dl_DataToUL_ACK_DCI_1_2_r18).Release = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case PUCCH_ResourceExt_v1610_Ext_Dl_DataToUL_ACK_DCI_1_2_r18_Setup:
				(*ie.Dl_DataToUL_ACK_DCI_1_2_r18).Setup = new(DL_DataToUL_ACK_DCI_1_2_r18)
				if err := (*ie.Dl_DataToUL_ACK_DCI_1_2_r18).Setup.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
