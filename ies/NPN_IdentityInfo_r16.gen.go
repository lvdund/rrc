// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NPN-IdentityInfo-r16 (line 10584).

var nPNIdentityInfoR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "npn-IdentityList-r16"},
		{Name: "trackingAreaCode-r16"},
		{Name: "ranac-r16", Optional: true},
		{Name: "cellIdentity-r16"},
		{Name: "cellReservedForOperatorUse-r16"},
		{Name: "iab-Support-r16", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
	},
}

var nPNIdentityInfoR16NpnIdentityListR16Constraints = per.SizeRange(1, common.MaxNPN_r16)

const (
	NPN_IdentityInfo_r16_CellReservedForOperatorUse_r16_Reserved    = 0
	NPN_IdentityInfo_r16_CellReservedForOperatorUse_r16_NotReserved = 1
)

var nPNIdentityInfoR16CellReservedForOperatorUseR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NPN_IdentityInfo_r16_CellReservedForOperatorUse_r16_Reserved, NPN_IdentityInfo_r16_CellReservedForOperatorUse_r16_NotReserved},
}

const (
	NPN_IdentityInfo_r16_Iab_Support_r16_True = 0
)

var nPNIdentityInfoR16IabSupportR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NPN_IdentityInfo_r16_Iab_Support_r16_True},
}

var nPNIdentityInfoR16GNBIDLengthR17Constraints = per.Constrained(22, 32)

const (
	NPN_IdentityInfo_r16_Ext_MobileIAB_Support_r18_True = 0
)

var nPNIdentityInfoR16ExtMobileIABSupportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{NPN_IdentityInfo_r16_Ext_MobileIAB_Support_r18_True},
}

type NPN_IdentityInfo_r16 struct {
	Npn_IdentityList_r16           []NPN_Identity_r16
	TrackingAreaCode_r16           TrackingAreaCode
	Ranac_r16                      *RAN_AreaCode
	CellIdentity_r16               CellIdentity
	CellReservedForOperatorUse_r16 int64
	Iab_Support_r16                *int64
	GNB_ID_Length_r17              *int64
	MobileIAB_Support_r18          *int64
}

func (ie *NPN_IdentityInfo_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nPNIdentityInfoR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.GNB_ID_Length_r17 != nil
	hasExtGroup1 := ie.MobileIAB_Support_r18 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ranac_r16 != nil, ie.Iab_Support_r16 != nil}); err != nil {
		return err
	}

	// 3. npn-IdentityList-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(nPNIdentityInfoR16NpnIdentityListR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Npn_IdentityList_r16))); err != nil {
			return err
		}
		for i := range ie.Npn_IdentityList_r16 {
			if err := ie.Npn_IdentityList_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. trackingAreaCode-r16: ref
	{
		if err := ie.TrackingAreaCode_r16.Encode(e); err != nil {
			return err
		}
	}

	// 5. ranac-r16: ref
	{
		if ie.Ranac_r16 != nil {
			if err := ie.Ranac_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. cellIdentity-r16: ref
	{
		if err := ie.CellIdentity_r16.Encode(e); err != nil {
			return err
		}
	}

	// 7. cellReservedForOperatorUse-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.CellReservedForOperatorUse_r16, nPNIdentityInfoR16CellReservedForOperatorUseR16Constraints); err != nil {
			return err
		}
	}

	// 8. iab-Support-r16: enumerated
	{
		if ie.Iab_Support_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Iab_Support_r16, nPNIdentityInfoR16IabSupportR16Constraints); err != nil {
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
					{Name: "gNB-ID-Length-r17", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.GNB_ID_Length_r17 != nil}); err != nil {
				return err
			}

			if ie.GNB_ID_Length_r17 != nil {
				if err := ex.EncodeInteger(*ie.GNB_ID_Length_r17, nPNIdentityInfoR16GNBIDLengthR17Constraints); err != nil {
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
					{Name: "mobileIAB-Support-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MobileIAB_Support_r18 != nil}); err != nil {
				return err
			}

			if ie.MobileIAB_Support_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.MobileIAB_Support_r18, nPNIdentityInfoR16ExtMobileIABSupportR18Constraints); err != nil {
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

func (ie *NPN_IdentityInfo_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nPNIdentityInfoR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. npn-IdentityList-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(nPNIdentityInfoR16NpnIdentityListR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Npn_IdentityList_r16 = make([]NPN_Identity_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Npn_IdentityList_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. trackingAreaCode-r16: ref
	{
		if err := ie.TrackingAreaCode_r16.Decode(d); err != nil {
			return err
		}
	}

	// 5. ranac-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Ranac_r16 = new(RAN_AreaCode)
			if err := ie.Ranac_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. cellIdentity-r16: ref
	{
		if err := ie.CellIdentity_r16.Decode(d); err != nil {
			return err
		}
	}

	// 7. cellReservedForOperatorUse-r16: enumerated
	{
		v4, err := d.DecodeEnumerated(nPNIdentityInfoR16CellReservedForOperatorUseR16Constraints)
		if err != nil {
			return err
		}
		ie.CellReservedForOperatorUse_r16 = v4
	}

	// 8. iab-Support-r16: enumerated
	{
		if seq.IsComponentPresent(5) {
			idx, err := d.DecodeEnumerated(nPNIdentityInfoR16IabSupportR16Constraints)
			if err != nil {
				return err
			}
			ie.Iab_Support_r16 = &idx
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
				{Name: "gNB-ID-Length-r17", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeInteger(nPNIdentityInfoR16GNBIDLengthR17Constraints)
			if err != nil {
				return err
			}
			ie.GNB_ID_Length_r17 = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
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
			v, err := dx.DecodeEnumerated(nPNIdentityInfoR16ExtMobileIABSupportR18Constraints)
			if err != nil {
				return err
			}
			ie.MobileIAB_Support_r18 = &v
		}
	}

	return nil
}
