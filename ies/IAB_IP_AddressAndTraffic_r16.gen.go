// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IAB-IP-AddressAndTraffic-r16 (line 476).

var iABIPAddressAndTrafficR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "all-Traffic-IAB-IP-Address-r16", Optional: true},
		{Name: "f1-C-Traffic-IP-Address-r16", Optional: true},
		{Name: "f1-U-Traffic-IP-Address-r16", Optional: true},
		{Name: "non-F1-Traffic-IP-Address-r16", Optional: true},
	},
}

var iABIPAddressAndTrafficR16AllTrafficIABIPAddressR16Constraints = per.SizeRange(1, 8)

var iABIPAddressAndTrafficR16F1CTrafficIPAddressR16Constraints = per.SizeRange(1, 8)

var iABIPAddressAndTrafficR16F1UTrafficIPAddressR16Constraints = per.SizeRange(1, 8)

var iABIPAddressAndTrafficR16NonF1TrafficIPAddressR16Constraints = per.SizeRange(1, 8)

type IAB_IP_AddressAndTraffic_r16 struct {
	All_Traffic_IAB_IP_Address_r16 []IAB_IP_Address_r16
	F1_C_Traffic_IP_Address_r16    []IAB_IP_Address_r16
	F1_U_Traffic_IP_Address_r16    []IAB_IP_Address_r16
	Non_F1_Traffic_IP_Address_r16  []IAB_IP_Address_r16
}

func (ie *IAB_IP_AddressAndTraffic_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iABIPAddressAndTrafficR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.All_Traffic_IAB_IP_Address_r16 != nil, ie.F1_C_Traffic_IP_Address_r16 != nil, ie.F1_U_Traffic_IP_Address_r16 != nil, ie.Non_F1_Traffic_IP_Address_r16 != nil}); err != nil {
		return err
	}

	// 2. all-Traffic-IAB-IP-Address-r16: sequence-of
	{
		if ie.All_Traffic_IAB_IP_Address_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(iABIPAddressAndTrafficR16AllTrafficIABIPAddressR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.All_Traffic_IAB_IP_Address_r16))); err != nil {
				return err
			}
			for i := range ie.All_Traffic_IAB_IP_Address_r16 {
				if err := ie.All_Traffic_IAB_IP_Address_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. f1-C-Traffic-IP-Address-r16: sequence-of
	{
		if ie.F1_C_Traffic_IP_Address_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(iABIPAddressAndTrafficR16F1CTrafficIPAddressR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.F1_C_Traffic_IP_Address_r16))); err != nil {
				return err
			}
			for i := range ie.F1_C_Traffic_IP_Address_r16 {
				if err := ie.F1_C_Traffic_IP_Address_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 4. f1-U-Traffic-IP-Address-r16: sequence-of
	{
		if ie.F1_U_Traffic_IP_Address_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(iABIPAddressAndTrafficR16F1UTrafficIPAddressR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.F1_U_Traffic_IP_Address_r16))); err != nil {
				return err
			}
			for i := range ie.F1_U_Traffic_IP_Address_r16 {
				if err := ie.F1_U_Traffic_IP_Address_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 5. non-F1-Traffic-IP-Address-r16: sequence-of
	{
		if ie.Non_F1_Traffic_IP_Address_r16 != nil {
			seqOf := e.NewSequenceOfEncoder(iABIPAddressAndTrafficR16NonF1TrafficIPAddressR16Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Non_F1_Traffic_IP_Address_r16))); err != nil {
				return err
			}
			for i := range ie.Non_F1_Traffic_IP_Address_r16 {
				if err := ie.Non_F1_Traffic_IP_Address_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *IAB_IP_AddressAndTraffic_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iABIPAddressAndTrafficR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. all-Traffic-IAB-IP-Address-r16: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(iABIPAddressAndTrafficR16AllTrafficIABIPAddressR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.All_Traffic_IAB_IP_Address_r16 = make([]IAB_IP_Address_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.All_Traffic_IAB_IP_Address_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. f1-C-Traffic-IP-Address-r16: sequence-of
	{
		if seq.IsComponentPresent(1) {
			seqOf := d.NewSequenceOfDecoder(iABIPAddressAndTrafficR16F1CTrafficIPAddressR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.F1_C_Traffic_IP_Address_r16 = make([]IAB_IP_Address_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.F1_C_Traffic_IP_Address_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 4. f1-U-Traffic-IP-Address-r16: sequence-of
	{
		if seq.IsComponentPresent(2) {
			seqOf := d.NewSequenceOfDecoder(iABIPAddressAndTrafficR16F1UTrafficIPAddressR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.F1_U_Traffic_IP_Address_r16 = make([]IAB_IP_Address_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.F1_U_Traffic_IP_Address_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 5. non-F1-Traffic-IP-Address-r16: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(iABIPAddressAndTrafficR16NonF1TrafficIPAddressR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Non_F1_Traffic_IP_Address_r16 = make([]IAB_IP_Address_r16, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Non_F1_Traffic_IP_Address_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
