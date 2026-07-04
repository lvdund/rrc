// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: PLMN-IdentityInfo (line 11901).

var pLMNIdentityInfoConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-IdentityList"},
		{Name: "trackingAreaCode", Optional: true},
		{Name: "ranac", Optional: true},
		{Name: "cellIdentity"},
		{Name: "cellReservedForOperatorUse"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

var pLMNIdentityInfoPlmnIdentityListConstraints = per.SizeRange(1, common.MaxPLMN)

const (
	PLMN_IdentityInfo_CellReservedForOperatorUse_Reserved    = 0
	PLMN_IdentityInfo_CellReservedForOperatorUse_NotReserved = 1
)

var pLMNIdentityInfoCellReservedForOperatorUseConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PLMN_IdentityInfo_CellReservedForOperatorUse_Reserved, PLMN_IdentityInfo_CellReservedForOperatorUse_NotReserved},
}

const (
	PLMN_IdentityInfo_Ext_Iab_Support_r16_True = 0
)

var pLMNIdentityInfoExtIabSupportR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PLMN_IdentityInfo_Ext_Iab_Support_r16_True},
}

var pLMNIdentityInfoExtTrackingAreaListR17Constraints = per.SizeRange(1, common.MaxTAC_r17)

var pLMNIdentityInfoGNBIDLengthR17Constraints = per.Constrained(22, 32)

const (
	PLMN_IdentityInfo_Ext_MobileIAB_Support_r18_True = 0
)

var pLMNIdentityInfoExtMobileIABSupportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PLMN_IdentityInfo_Ext_MobileIAB_Support_r18_True},
}

type PLMN_IdentityInfo struct {
	Plmn_IdentityList          []PLMN_Identity
	TrackingAreaCode           *TrackingAreaCode
	Ranac                      *RAN_AreaCode
	CellIdentity               CellIdentity
	CellReservedForOperatorUse int64
	Iab_Support_r16            *int64
	TrackingAreaList_r17       []TrackingAreaCode
	GNB_ID_Length_r17          *int64
	MobileIAB_Support_r18      *int64
}

func (ie *PLMN_IdentityInfo) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pLMNIdentityInfoConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Iab_Support_r16 != nil
	hasExtGroup1 := ie.TrackingAreaList_r17 != nil || ie.GNB_ID_Length_r17 != nil
	hasExtGroup2 := ie.MobileIAB_Support_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.TrackingAreaCode != nil, ie.Ranac != nil}); err != nil {
		return err
	}

	// 3. plmn-IdentityList: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(pLMNIdentityInfoPlmnIdentityListConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.Plmn_IdentityList))); err != nil {
			return err
		}
		for i := range ie.Plmn_IdentityList {
			if err := ie.Plmn_IdentityList[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. trackingAreaCode: ref
	{
		if ie.TrackingAreaCode != nil {
			if err := ie.TrackingAreaCode.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. ranac: ref
	{
		if ie.Ranac != nil {
			if err := ie.Ranac.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. cellIdentity: ref
	{
		if err := ie.CellIdentity.Encode(e); err != nil {
			return err
		}
	}

	// 7. cellReservedForOperatorUse: enumerated
	{
		if err := e.EncodeEnumerated(ie.CellReservedForOperatorUse, pLMNIdentityInfoCellReservedForOperatorUseConstraints); err != nil {
			return err
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "iab-Support-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Iab_Support_r16 != nil}); err != nil {
				return err
			}

			if ie.Iab_Support_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Iab_Support_r16, pLMNIdentityInfoExtIabSupportR16Constraints); err != nil {
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
					{Name: "trackingAreaList-r17", Optional: true},
					{Name: "gNB-ID-Length-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.TrackingAreaList_r17 != nil, ie.GNB_ID_Length_r17 != nil}); err != nil {
				return err
			}

			if ie.TrackingAreaList_r17 != nil {
				seqOf := ex.NewSequenceOfEncoder(pLMNIdentityInfoExtTrackingAreaListR17Constraints)
				if err := seqOf.EncodeLength(int64(len(ie.TrackingAreaList_r17))); err != nil {
					return err
				}
				for i := range ie.TrackingAreaList_r17 {
					if err := ie.TrackingAreaList_r17[i].Encode(ex); err != nil {
						return err
					}
				}
			}

			if ie.GNB_ID_Length_r17 != nil {
				if err := ex.EncodeInteger(*ie.GNB_ID_Length_r17, pLMNIdentityInfoGNBIDLengthR17Constraints); err != nil {
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
					{Name: "mobileIAB-Support-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MobileIAB_Support_r18 != nil}); err != nil {
				return err
			}

			if ie.MobileIAB_Support_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.MobileIAB_Support_r18, pLMNIdentityInfoExtMobileIABSupportR18Constraints); err != nil {
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

func (ie *PLMN_IdentityInfo) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pLMNIdentityInfoConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. plmn-IdentityList: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(pLMNIdentityInfoPlmnIdentityListConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Plmn_IdentityList = make([]PLMN_Identity, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Plmn_IdentityList[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. trackingAreaCode: ref
	{
		if seq.IsComponentPresent(1) {
			ie.TrackingAreaCode = new(TrackingAreaCode)
			if err := ie.TrackingAreaCode.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. ranac: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Ranac = new(RAN_AreaCode)
			if err := ie.Ranac.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. cellIdentity: ref
	{
		if err := ie.CellIdentity.Decode(d); err != nil {
			return err
		}
	}

	// 7. cellReservedForOperatorUse: enumerated
	{
		v4, err := d.DecodeEnumerated(pLMNIdentityInfoCellReservedForOperatorUseConstraints)
		if err != nil {
			return err
		}
		ie.CellReservedForOperatorUse = v4
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
				{Name: "iab-Support-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pLMNIdentityInfoExtIabSupportR16Constraints)
			if err != nil {
				return err
			}
			ie.Iab_Support_r16 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "trackingAreaList-r17", Optional: true},
				{Name: "gNB-ID-Length-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			seqOf := dx.NewSequenceOfDecoder(pLMNIdentityInfoExtTrackingAreaListR17Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.TrackingAreaList_r17 = make([]TrackingAreaCode, n)
			for i := int64(0); i < n; i++ {
				if err := ie.TrackingAreaList_r17[i].Decode(dx); err != nil {
					return err
				}
			}
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeInteger(pLMNIdentityInfoGNBIDLengthR17Constraints)
			if err != nil {
				return err
			}
			ie.GNB_ID_Length_r17 = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "mobileIAB-Support-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(pLMNIdentityInfoExtMobileIABSupportR18Constraints)
			if err != nil {
				return err
			}
			ie.MobileIAB_Support_r18 = &v
		}
	}

	return nil
}
