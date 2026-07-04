// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IAB-IP-PrefixAndTraffic-r16 (line 483).

var iABIPPrefixAndTrafficR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "all-Traffic-IAB-IP-Address-r16", Optional: true},
		{Name: "f1-C-Traffic-IP-Address-r16", Optional: true},
		{Name: "f1-U-Traffic-IP-Address-r16", Optional: true},
		{Name: "non-F1-Traffic-IP-Address-r16", Optional: true},
	},
}

type IAB_IP_PrefixAndTraffic_r16 struct {
	All_Traffic_IAB_IP_Address_r16 *IAB_IP_Address_r16
	F1_C_Traffic_IP_Address_r16    *IAB_IP_Address_r16
	F1_U_Traffic_IP_Address_r16    *IAB_IP_Address_r16
	Non_F1_Traffic_IP_Address_r16  *IAB_IP_Address_r16
}

func (ie *IAB_IP_PrefixAndTraffic_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iABIPPrefixAndTrafficR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.All_Traffic_IAB_IP_Address_r16 != nil, ie.F1_C_Traffic_IP_Address_r16 != nil, ie.F1_U_Traffic_IP_Address_r16 != nil, ie.Non_F1_Traffic_IP_Address_r16 != nil}); err != nil {
		return err
	}

	// 2. all-Traffic-IAB-IP-Address-r16: ref
	{
		if ie.All_Traffic_IAB_IP_Address_r16 != nil {
			if err := ie.All_Traffic_IAB_IP_Address_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. f1-C-Traffic-IP-Address-r16: ref
	{
		if ie.F1_C_Traffic_IP_Address_r16 != nil {
			if err := ie.F1_C_Traffic_IP_Address_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. f1-U-Traffic-IP-Address-r16: ref
	{
		if ie.F1_U_Traffic_IP_Address_r16 != nil {
			if err := ie.F1_U_Traffic_IP_Address_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. non-F1-Traffic-IP-Address-r16: ref
	{
		if ie.Non_F1_Traffic_IP_Address_r16 != nil {
			if err := ie.Non_F1_Traffic_IP_Address_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *IAB_IP_PrefixAndTraffic_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iABIPPrefixAndTrafficR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. all-Traffic-IAB-IP-Address-r16: ref
	{
		if seq.IsComponentPresent(0) {
			ie.All_Traffic_IAB_IP_Address_r16 = new(IAB_IP_Address_r16)
			if err := ie.All_Traffic_IAB_IP_Address_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. f1-C-Traffic-IP-Address-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.F1_C_Traffic_IP_Address_r16 = new(IAB_IP_Address_r16)
			if err := ie.F1_C_Traffic_IP_Address_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. f1-U-Traffic-IP-Address-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.F1_U_Traffic_IP_Address_r16 = new(IAB_IP_Address_r16)
			if err := ie.F1_U_Traffic_IP_Address_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. non-F1-Traffic-IP-Address-r16: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Non_F1_Traffic_IP_Address_r16 = new(IAB_IP_Address_r16)
			if err := ie.Non_F1_Traffic_IP_Address_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
