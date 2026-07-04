// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IAB-IP-AddressConfiguration-r16 (line 1106).

var iABIPAddressConfigurationR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "iab-IP-AddressIndex-r16"},
		{Name: "iab-IP-Address-r16", Optional: true},
		{Name: "iab-IP-Usage-r16", Optional: true},
		{Name: "iab-donor-DU-BAP-Address-r16", Optional: true},
	},
}

var iABIPAddressConfigurationR16IabDonorDUBAPAddressR16Constraints = per.FixedSize(10)

type IAB_IP_AddressConfiguration_r16 struct {
	Iab_IP_AddressIndex_r16      IAB_IP_AddressIndex_r16
	Iab_IP_Address_r16           *IAB_IP_Address_r16
	Iab_IP_Usage_r16             *IAB_IP_Usage_r16
	Iab_Donor_DU_BAP_Address_r16 *per.BitString
}

func (ie *IAB_IP_AddressConfiguration_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iABIPAddressConfigurationR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Iab_IP_Address_r16 != nil, ie.Iab_IP_Usage_r16 != nil, ie.Iab_Donor_DU_BAP_Address_r16 != nil}); err != nil {
		return err
	}

	// 3. iab-IP-AddressIndex-r16: ref
	{
		if err := ie.Iab_IP_AddressIndex_r16.Encode(e); err != nil {
			return err
		}
	}

	// 4. iab-IP-Address-r16: ref
	{
		if ie.Iab_IP_Address_r16 != nil {
			if err := ie.Iab_IP_Address_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. iab-IP-Usage-r16: ref
	{
		if ie.Iab_IP_Usage_r16 != nil {
			if err := ie.Iab_IP_Usage_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. iab-donor-DU-BAP-Address-r16: bit-string
	{
		if ie.Iab_Donor_DU_BAP_Address_r16 != nil {
			if err := e.EncodeBitString(*ie.Iab_Donor_DU_BAP_Address_r16, iABIPAddressConfigurationR16IabDonorDUBAPAddressR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *IAB_IP_AddressConfiguration_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iABIPAddressConfigurationR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. iab-IP-AddressIndex-r16: ref
	{
		if err := ie.Iab_IP_AddressIndex_r16.Decode(d); err != nil {
			return err
		}
	}

	// 4. iab-IP-Address-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Iab_IP_Address_r16 = new(IAB_IP_Address_r16)
			if err := ie.Iab_IP_Address_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. iab-IP-Usage-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Iab_IP_Usage_r16 = new(IAB_IP_Usage_r16)
			if err := ie.Iab_IP_Usage_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. iab-donor-DU-BAP-Address-r16: bit-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeBitString(iABIPAddressConfigurationR16IabDonorDUBAPAddressR16Constraints)
			if err != nil {
				return err
			}
			ie.Iab_Donor_DU_BAP_Address_r16 = &v
		}
	}

	return nil
}
