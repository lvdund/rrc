// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IAB-IP-AddressPrefixReq-r16 (line 468).

var iABIPAddressPrefixReqR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "all-Traffic-PrefixReq-r16", Optional: true},
		{Name: "f1-C-Traffic-PrefixReq-r16", Optional: true},
		{Name: "f1-U-Traffic-PrefixReq-r16", Optional: true},
		{Name: "non-F1-Traffic-PrefixReq-r16", Optional: true},
	},
}

const (
	IAB_IP_AddressPrefixReq_r16_All_Traffic_PrefixReq_r16_True = 0
)

var iABIPAddressPrefixReqR16AllTrafficPrefixReqR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{IAB_IP_AddressPrefixReq_r16_All_Traffic_PrefixReq_r16_True},
}

const (
	IAB_IP_AddressPrefixReq_r16_F1_C_Traffic_PrefixReq_r16_True = 0
)

var iABIPAddressPrefixReqR16F1CTrafficPrefixReqR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{IAB_IP_AddressPrefixReq_r16_F1_C_Traffic_PrefixReq_r16_True},
}

const (
	IAB_IP_AddressPrefixReq_r16_F1_U_Traffic_PrefixReq_r16_True = 0
)

var iABIPAddressPrefixReqR16F1UTrafficPrefixReqR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{IAB_IP_AddressPrefixReq_r16_F1_U_Traffic_PrefixReq_r16_True},
}

const (
	IAB_IP_AddressPrefixReq_r16_Non_F1_Traffic_PrefixReq_r16_True = 0
)

var iABIPAddressPrefixReqR16NonF1TrafficPrefixReqR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{IAB_IP_AddressPrefixReq_r16_Non_F1_Traffic_PrefixReq_r16_True},
}

type IAB_IP_AddressPrefixReq_r16 struct {
	All_Traffic_PrefixReq_r16    *int64
	F1_C_Traffic_PrefixReq_r16   *int64
	F1_U_Traffic_PrefixReq_r16   *int64
	Non_F1_Traffic_PrefixReq_r16 *int64
}

func (ie *IAB_IP_AddressPrefixReq_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iABIPAddressPrefixReqR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.All_Traffic_PrefixReq_r16 != nil, ie.F1_C_Traffic_PrefixReq_r16 != nil, ie.F1_U_Traffic_PrefixReq_r16 != nil, ie.Non_F1_Traffic_PrefixReq_r16 != nil}); err != nil {
		return err
	}

	// 3. all-Traffic-PrefixReq-r16: enumerated
	{
		if ie.All_Traffic_PrefixReq_r16 != nil {
			if err := e.EncodeEnumerated(*ie.All_Traffic_PrefixReq_r16, iABIPAddressPrefixReqR16AllTrafficPrefixReqR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. f1-C-Traffic-PrefixReq-r16: enumerated
	{
		if ie.F1_C_Traffic_PrefixReq_r16 != nil {
			if err := e.EncodeEnumerated(*ie.F1_C_Traffic_PrefixReq_r16, iABIPAddressPrefixReqR16F1CTrafficPrefixReqR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. f1-U-Traffic-PrefixReq-r16: enumerated
	{
		if ie.F1_U_Traffic_PrefixReq_r16 != nil {
			if err := e.EncodeEnumerated(*ie.F1_U_Traffic_PrefixReq_r16, iABIPAddressPrefixReqR16F1UTrafficPrefixReqR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. non-F1-Traffic-PrefixReq-r16: enumerated
	{
		if ie.Non_F1_Traffic_PrefixReq_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Non_F1_Traffic_PrefixReq_r16, iABIPAddressPrefixReqR16NonF1TrafficPrefixReqR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *IAB_IP_AddressPrefixReq_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iABIPAddressPrefixReqR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. all-Traffic-PrefixReq-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(iABIPAddressPrefixReqR16AllTrafficPrefixReqR16Constraints)
			if err != nil {
				return err
			}
			ie.All_Traffic_PrefixReq_r16 = &idx
		}
	}

	// 4. f1-C-Traffic-PrefixReq-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(iABIPAddressPrefixReqR16F1CTrafficPrefixReqR16Constraints)
			if err != nil {
				return err
			}
			ie.F1_C_Traffic_PrefixReq_r16 = &idx
		}
	}

	// 5. f1-U-Traffic-PrefixReq-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(iABIPAddressPrefixReqR16F1UTrafficPrefixReqR16Constraints)
			if err != nil {
				return err
			}
			ie.F1_U_Traffic_PrefixReq_r16 = &idx
		}
	}

	// 6. non-F1-Traffic-PrefixReq-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(iABIPAddressPrefixReqR16NonF1TrafficPrefixReqR16Constraints)
			if err != nil {
				return err
			}
			ie.Non_F1_Traffic_PrefixReq_r16 = &idx
		}
	}

	return nil
}
