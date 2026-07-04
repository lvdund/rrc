// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SIB3 (line 3869).

var sIB3Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "intraFreqNeighCellList", Optional: true},
		{Name: "intraFreqExcludedCellList", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
	},
}

var sIB3LateNonCriticalExtensionConstraints = per.SizeConstraints{}

var sIB3ExtIntraFreqCAGCellListR16Constraints = per.SizeRange(1, common.MaxPLMN)

const (
	SIB3_Ext_ChannelAccessMode2_r17_Enabled = 0
)

var sIB3ExtChannelAccessMode2R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB3_Ext_ChannelAccessMode2_r17_Enabled},
}

var sIB3ExtIntraFreqODSIB1ExcludedCellListR19Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "emptyList-r19"},
		{Name: "excludedCells-r19"},
	},
}

const (
	SIB3_Ext_IntraFreqOD_SIB1_ExcludedCellList_r19_EmptyList_r19     = 0
	SIB3_Ext_IntraFreqOD_SIB1_ExcludedCellList_r19_ExcludedCells_r19 = 1
)

type SIB3 struct {
	IntraFreqNeighCellList                *IntraFreqNeighCellList
	IntraFreqExcludedCellList             *IntraFreqExcludedCellList
	LateNonCriticalExtension              []byte
	IntraFreqNeighCellList_v1610          *IntraFreqNeighCellList_v1610
	IntraFreqAllowedCellList_r16          *IntraFreqAllowedCellList_r16
	IntraFreqCAG_CellList_r16             []IntraFreqCAG_CellListPerPLMN_r16
	IntraFreqNeighHSDN_CellList_r17       *IntraFreqNeighHSDN_CellList_r17
	IntraFreqNeighCellList_v1710          *IntraFreqNeighCellList_v1710
	ChannelAccessMode2_r17                *int64
	IntraFreqOD_SIB1_ExcludedCellList_r19 *struct {
		Choice            int
		EmptyList_r19     *struct{}
		ExcludedCells_r19 *IntraFreqExcludedCellList
	}
}

func (ie *SIB3) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB3Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.IntraFreqNeighCellList_v1610 != nil || ie.IntraFreqAllowedCellList_r16 != nil || ie.IntraFreqCAG_CellList_r16 != nil
	hasExtGroup1 := ie.IntraFreqNeighHSDN_CellList_r17 != nil || ie.IntraFreqNeighCellList_v1710 != nil
	hasExtGroup2 := ie.ChannelAccessMode2_r17 != nil
	hasExtGroup3 := ie.IntraFreqOD_SIB1_ExcludedCellList_r19 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IntraFreqNeighCellList != nil, ie.IntraFreqExcludedCellList != nil, ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. intraFreqNeighCellList: ref
	{
		if ie.IntraFreqNeighCellList != nil {
			if err := ie.IntraFreqNeighCellList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. intraFreqExcludedCellList: ref
	{
		if ie.IntraFreqExcludedCellList != nil {
			if err := ie.IntraFreqExcludedCellList.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB3LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "intraFreqNeighCellList-v1610", Optional: true},
					{Name: "intraFreqAllowedCellList-r16", Optional: true},
					{Name: "intraFreqCAG-CellList-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.IntraFreqNeighCellList_v1610 != nil, ie.IntraFreqAllowedCellList_r16 != nil, ie.IntraFreqCAG_CellList_r16 != nil}); err != nil {
				return err
			}

			if ie.IntraFreqNeighCellList_v1610 != nil {
				if err := ie.IntraFreqNeighCellList_v1610.Encode(ex); err != nil {
					return err
				}
			}

			if ie.IntraFreqAllowedCellList_r16 != nil {
				if err := ie.IntraFreqAllowedCellList_r16.Encode(ex); err != nil {
					return err
				}
			}

			if ie.IntraFreqCAG_CellList_r16 != nil {
				seqOf := ex.NewSequenceOfEncoder(sIB3ExtIntraFreqCAGCellListR16Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.IntraFreqCAG_CellList_r16))); err != nil {
					return err
				}
				for i := range ie.IntraFreqCAG_CellList_r16 {
					if err := ie.IntraFreqCAG_CellList_r16[i].Encode(ex); err != nil {
						return err
					}
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
					{Name: "intraFreqNeighHSDN-CellList-r17", Optional: true},
					{Name: "intraFreqNeighCellList-v1710", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.IntraFreqNeighHSDN_CellList_r17 != nil, ie.IntraFreqNeighCellList_v1710 != nil}); err != nil {
				return err
			}

			if ie.IntraFreqNeighHSDN_CellList_r17 != nil {
				if err := ie.IntraFreqNeighHSDN_CellList_r17.Encode(ex); err != nil {
					return err
				}
			}

			if ie.IntraFreqNeighCellList_v1710 != nil {
				if err := ie.IntraFreqNeighCellList_v1710.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "channelAccessMode2-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ChannelAccessMode2_r17 != nil}); err != nil {
				return err
			}

			if ie.ChannelAccessMode2_r17 != nil {
				if err := ex.EncodeEnumerated(*ie.ChannelAccessMode2_r17, sIB3ExtChannelAccessMode2R17Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "intraFreqOD-SIB1-ExcludedCellList-r19", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.IntraFreqOD_SIB1_ExcludedCellList_r19 != nil}); err != nil {
				return err
			}

			if ie.IntraFreqOD_SIB1_ExcludedCellList_r19 != nil {
				choiceEnc := ex.NewChoiceEncoder(sIB3ExtIntraFreqODSIB1ExcludedCellListR19Constraints)
				if err := choiceEnc.EncodeChoice(int64((*ie.IntraFreqOD_SIB1_ExcludedCellList_r19).Choice), false, nil); err != nil {
					return err
				}
				switch (*ie.IntraFreqOD_SIB1_ExcludedCellList_r19).Choice {
				case SIB3_Ext_IntraFreqOD_SIB1_ExcludedCellList_r19_EmptyList_r19:
					if err := ex.EncodeNull(); err != nil {
						return err
					}
				case SIB3_Ext_IntraFreqOD_SIB1_ExcludedCellList_r19_ExcludedCells_r19:
					if err := (*ie.IntraFreqOD_SIB1_ExcludedCellList_r19).ExcludedCells_r19.Encode(ex); err != nil {
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

func (ie *SIB3) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB3Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. intraFreqNeighCellList: ref
	{
		if seq.IsComponentPresent(0) {
			ie.IntraFreqNeighCellList = new(IntraFreqNeighCellList)
			if err := ie.IntraFreqNeighCellList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. intraFreqExcludedCellList: ref
	{
		if seq.IsComponentPresent(1) {
			ie.IntraFreqExcludedCellList = new(IntraFreqExcludedCellList)
			if err := ie.IntraFreqExcludedCellList.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(sIB3LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
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
				{Name: "intraFreqNeighCellList-v1610", Optional: true},
				{Name: "intraFreqAllowedCellList-r16", Optional: true},
				{Name: "intraFreqCAG-CellList-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.IntraFreqNeighCellList_v1610 = new(IntraFreqNeighCellList_v1610)
			if err := ie.IntraFreqNeighCellList_v1610.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.IntraFreqAllowedCellList_r16 = new(IntraFreqAllowedCellList_r16)
			if err := ie.IntraFreqAllowedCellList_r16.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(2) {
			seqOf := dx.NewSequenceOfDecoder(sIB3ExtIntraFreqCAGCellListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.IntraFreqCAG_CellList_r16 = make([]IntraFreqCAG_CellListPerPLMN_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.IntraFreqCAG_CellList_r16[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "intraFreqNeighHSDN-CellList-r17", Optional: true},
				{Name: "intraFreqNeighCellList-v1710", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.IntraFreqNeighHSDN_CellList_r17 = new(IntraFreqNeighHSDN_CellList_r17)
			if err := ie.IntraFreqNeighHSDN_CellList_r17.Decode(dx); err != nil {
				return err
			}
		}

		if groupSeq.IsComponentPresent(1) {
			ie.IntraFreqNeighCellList_v1710 = new(IntraFreqNeighCellList_v1710)
			if err := ie.IntraFreqNeighCellList_v1710.Decode(dx); err != nil {
				return err
			}
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "channelAccessMode2-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(sIB3ExtChannelAccessMode2R17Constraints)
			if err != nil {
				return err
			}
			ie.ChannelAccessMode2_r17 = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "intraFreqOD-SIB1-ExcludedCellList-r19", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.IntraFreqOD_SIB1_ExcludedCellList_r19 = &struct {
				Choice            int
				EmptyList_r19     *struct{}
				ExcludedCells_r19 *IntraFreqExcludedCellList
			}{}
			choiceDec := dx.NewChoiceDecoder(sIB3ExtIntraFreqODSIB1ExcludedCellListR19Constraints)
			index, _, _, err := choiceDec.DecodeChoice()
			if err != nil {
				return err
			}
			(*ie.IntraFreqOD_SIB1_ExcludedCellList_r19).Choice = int(index)
			switch index {
			case SIB3_Ext_IntraFreqOD_SIB1_ExcludedCellList_r19_EmptyList_r19:
				(*ie.IntraFreqOD_SIB1_ExcludedCellList_r19).EmptyList_r19 = &struct{}{}
				if err := dx.DecodeNull(); err != nil {
					return err
				}
			case SIB3_Ext_IntraFreqOD_SIB1_ExcludedCellList_r19_ExcludedCells_r19:
				(*ie.IntraFreqOD_SIB1_ExcludedCellList_r19).ExcludedCells_r19 = new(IntraFreqExcludedCellList)
				if err := (*ie.IntraFreqOD_SIB1_ExcludedCellList_r19).ExcludedCells_r19.Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
