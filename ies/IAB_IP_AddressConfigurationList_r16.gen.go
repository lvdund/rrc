// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: IAB-IP-AddressConfigurationList-r16 (line 1100).

var iABIPAddressConfigurationListR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "iab-IP-AddressToAddModList-r16", Optional: true},
		{Name: "iab-IP-AddressToReleaseList-r16", Optional: true},
	},
}

var iABIPAddressConfigurationListR16IabIPAddressToAddModListR16Constraints = per.SizeRange(1, common.MaxIAB_IP_Address_r16)

var iABIPAddressConfigurationListR16IabIPAddressToReleaseListR16Constraints = per.SizeRange(1, common.MaxIAB_IP_Address_r16)

type IAB_IP_AddressConfigurationList_r16 struct {
	Iab_IP_AddressToAddModList_r16  []IAB_IP_AddressConfiguration_r16
	Iab_IP_AddressToReleaseList_r16 []IAB_IP_AddressIndex_r16
}

func (ie *IAB_IP_AddressConfigurationList_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iABIPAddressConfigurationListR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Iab_IP_AddressToAddModList_r16 != nil, ie.Iab_IP_AddressToReleaseList_r16 != nil}); err != nil {
		return err
	}

	// 3. iab-IP-AddressToAddModList-r16: sequence-of
	{
		if ie.Iab_IP_AddressToAddModList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(iABIPAddressConfigurationListR16IabIPAddressToAddModListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Iab_IP_AddressToAddModList_r16))); err != nil {
				return err
			}
			for i := range ie.Iab_IP_AddressToAddModList_r16 {
				if err := ie.Iab_IP_AddressToAddModList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. iab-IP-AddressToReleaseList-r16: sequence-of
	{
		if ie.Iab_IP_AddressToReleaseList_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(iABIPAddressConfigurationListR16IabIPAddressToReleaseListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Iab_IP_AddressToReleaseList_r16))); err != nil {
				return err
			}
			for i := range ie.Iab_IP_AddressToReleaseList_r16 {
				if err := ie.Iab_IP_AddressToReleaseList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *IAB_IP_AddressConfigurationList_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iABIPAddressConfigurationListR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. iab-IP-AddressToAddModList-r16: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(iABIPAddressConfigurationListR16IabIPAddressToAddModListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Iab_IP_AddressToAddModList_r16 = make([]IAB_IP_AddressConfiguration_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Iab_IP_AddressToAddModList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. iab-IP-AddressToReleaseList-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(iABIPAddressConfigurationListR16IabIPAddressToReleaseListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Iab_IP_AddressToReleaseList_r16 = make([]IAB_IP_AddressIndex_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Iab_IP_AddressToReleaseList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
