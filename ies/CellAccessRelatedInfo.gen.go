// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CellAccessRelatedInfo (line 5501).

var cellAccessRelatedInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-IdentityInfoList"},
		{Name: "cellReservedForOtherUse", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

const (
	CellAccessRelatedInfo_CellReservedForOtherUse_True = 0
)

var cellAccessRelatedInfoCellReservedForOtherUseConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellAccessRelatedInfo_CellReservedForOtherUse_True},
}

const (
	CellAccessRelatedInfo_Ext_CellReservedForFutureUse_r16_True = 0
)

var cellAccessRelatedInfoExtCellReservedForFutureUseR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CellAccessRelatedInfo_Ext_CellReservedForFutureUse_r16_True},
}

var cellAccessRelatedInfoExtSnpnAccessInfoListR17Constraints = per.SizeRange(1, common.MaxNPN_r16)

type CellAccessRelatedInfo struct {
	Plmn_IdentityInfoList        PLMN_IdentityInfoList
	CellReservedForOtherUse      *int64
	CellReservedForFutureUse_r16 *int64
	Npn_IdentityInfoList_r16     *NPN_IdentityInfoList_r16
	Snpn_AccessInfoList_r17      []SNPN_AccessInfo_r17
}

func (ie *CellAccessRelatedInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cellAccessRelatedInfoConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.CellReservedForFutureUse_r16 != nil || ie.Npn_IdentityInfoList_r16 != nil
	hasExtGroup1 := ie.Snpn_AccessInfoList_r17 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CellReservedForOtherUse != nil}); err != nil {
		return err
	}

	// 3. plmn-IdentityInfoList: ref
	{
		if err := ie.Plmn_IdentityInfoList.Encode(e); err != nil {
			return err
		}
	}

	// 4. cellReservedForOtherUse: enumerated
	{
		if ie.CellReservedForOtherUse != nil {
			if err := e.EncodeEnumerated(*ie.CellReservedForOtherUse, cellAccessRelatedInfoCellReservedForOtherUseConstraints); err != nil {
				return err
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
					{Name: "cellReservedForFutureUse-r16", Optional: true},
					{Name: "npn-IdentityInfoList-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.CellReservedForFutureUse_r16 != nil, ie.Npn_IdentityInfoList_r16 != nil}); err != nil {
				return err
			}

			if ie.CellReservedForFutureUse_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.CellReservedForFutureUse_r16, cellAccessRelatedInfoExtCellReservedForFutureUseR16Constraints); err != nil {
					return err
				}
			}

			if ie.Npn_IdentityInfoList_r16 != nil {
				if err := ie.Npn_IdentityInfoList_r16.Encode(ex); err != nil {
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
					{Name: "snpn-AccessInfoList-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Snpn_AccessInfoList_r17 != nil}); err != nil {
				return err
			}

			if ie.Snpn_AccessInfoList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(cellAccessRelatedInfoExtSnpnAccessInfoListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.Snpn_AccessInfoList_r17))); err != nil {
					return err
				}
				for i := range ie.Snpn_AccessInfoList_r17 {
					if err := ie.Snpn_AccessInfoList_r17[i].Encode(ex); err != nil {
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

func (ie *CellAccessRelatedInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cellAccessRelatedInfoConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. plmn-IdentityInfoList: ref
	{
		if err := ie.Plmn_IdentityInfoList.Decode(d); err != nil {
			return err
		}
	}

	// 4. cellReservedForOtherUse: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(cellAccessRelatedInfoCellReservedForOtherUseConstraints)
			if err != nil {
				return err
			}
			ie.CellReservedForOtherUse = &idx
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
				{Name: "cellReservedForFutureUse-r16", Optional: true},
				{Name: "npn-IdentityInfoList-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(cellAccessRelatedInfoExtCellReservedForFutureUseR16Constraints)
			if err != nil {
				return err
			}
			ie.CellReservedForFutureUse_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			ie.Npn_IdentityInfoList_r16 = new(NPN_IdentityInfoList_r16)
			if err := ie.Npn_IdentityInfoList_r16.Decode(dx); err != nil {
				return err
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
				{Name: "snpn-AccessInfoList-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(cellAccessRelatedInfoExtSnpnAccessInfoListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Snpn_AccessInfoList_r17 = make([]SNPN_AccessInfo_r17, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Snpn_AccessInfoList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
