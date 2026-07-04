// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IAB-IP-AddressNumReq-r16 (line 460).

var iABIPAddressNumReqR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "all-Traffic-NumReq-r16", Optional: true},
		{Name: "f1-C-Traffic-NumReq-r16", Optional: true},
		{Name: "f1-U-Traffic-NumReq-r16", Optional: true},
		{Name: "non-F1-Traffic-NumReq-r16", Optional: true},
	},
}

var iABIPAddressNumReqR16AllTrafficNumReqR16Constraints = per.Constrained(1, 8)

var iABIPAddressNumReqR16F1CTrafficNumReqR16Constraints = per.Constrained(1, 8)

var iABIPAddressNumReqR16F1UTrafficNumReqR16Constraints = per.Constrained(1, 8)

var iABIPAddressNumReqR16NonF1TrafficNumReqR16Constraints = per.Constrained(1, 8)

type IAB_IP_AddressNumReq_r16 struct {
	All_Traffic_NumReq_r16    *int64
	F1_C_Traffic_NumReq_r16   *int64
	F1_U_Traffic_NumReq_r16   *int64
	Non_F1_Traffic_NumReq_r16 *int64
}

func (ie *IAB_IP_AddressNumReq_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(iABIPAddressNumReqR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.All_Traffic_NumReq_r16 != nil, ie.F1_C_Traffic_NumReq_r16 != nil, ie.F1_U_Traffic_NumReq_r16 != nil, ie.Non_F1_Traffic_NumReq_r16 != nil}); err != nil {
		return err
	}

	// 3. all-Traffic-NumReq-r16: integer
	{
		if ie.All_Traffic_NumReq_r16 != nil {
			if err := e.EncodeInteger(*ie.All_Traffic_NumReq_r16, iABIPAddressNumReqR16AllTrafficNumReqR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. f1-C-Traffic-NumReq-r16: integer
	{
		if ie.F1_C_Traffic_NumReq_r16 != nil {
			if err := e.EncodeInteger(*ie.F1_C_Traffic_NumReq_r16, iABIPAddressNumReqR16F1CTrafficNumReqR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. f1-U-Traffic-NumReq-r16: integer
	{
		if ie.F1_U_Traffic_NumReq_r16 != nil {
			if err := e.EncodeInteger(*ie.F1_U_Traffic_NumReq_r16, iABIPAddressNumReqR16F1UTrafficNumReqR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. non-F1-Traffic-NumReq-r16: integer
	{
		if ie.Non_F1_Traffic_NumReq_r16 != nil {
			if err := e.EncodeInteger(*ie.Non_F1_Traffic_NumReq_r16, iABIPAddressNumReqR16NonF1TrafficNumReqR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *IAB_IP_AddressNumReq_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(iABIPAddressNumReqR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. all-Traffic-NumReq-r16: integer
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeInteger(iABIPAddressNumReqR16AllTrafficNumReqR16Constraints)
			if err != nil {
				return err
			}
			ie.All_Traffic_NumReq_r16 = &v
		}
	}

	// 4. f1-C-Traffic-NumReq-r16: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(iABIPAddressNumReqR16F1CTrafficNumReqR16Constraints)
			if err != nil {
				return err
			}
			ie.F1_C_Traffic_NumReq_r16 = &v
		}
	}

	// 5. f1-U-Traffic-NumReq-r16: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(iABIPAddressNumReqR16F1UTrafficNumReqR16Constraints)
			if err != nil {
				return err
			}
			ie.F1_U_Traffic_NumReq_r16 = &v
		}
	}

	// 6. non-F1-Traffic-NumReq-r16: integer
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeInteger(iABIPAddressNumReqR16NonF1TrafficNumReqR16Constraints)
			if err != nil {
				return err
			}
			ie.Non_F1_Traffic_NumReq_r16 = &v
		}
	}

	return nil
}
