// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MIMOParam-v1900 (line 14822).

var mIMOParamV1900Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "uplink-PowerControlToAddModListExt-v1910", Optional: true},
	},
}

var mIMOParamV1900UplinkPowerControlToAddModListExtV1910Constraints = per.SizeRange(1, common.MaxUL_TCI_r17)

type MIMOParam_v1900 struct {
	Uplink_PowerControlToAddModListExt_v1910 []Uplink_PowerControlExt_v1900
}

func (ie *MIMOParam_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mIMOParamV1900Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Uplink_PowerControlToAddModListExt_v1910 != nil}); err != nil {
		return err
	}

	// 3. uplink-PowerControlToAddModListExt-v1910: sequence-of
	{
		if ie.Uplink_PowerControlToAddModListExt_v1910 != nil {
			seqOf := e.NewSequenceOfEncoder(mIMOParamV1900UplinkPowerControlToAddModListExtV1910Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.Uplink_PowerControlToAddModListExt_v1910))); err != nil {
				return err
			}
			for i := range ie.Uplink_PowerControlToAddModListExt_v1910 {
				if err := ie.Uplink_PowerControlToAddModListExt_v1910[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *MIMOParam_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mIMOParamV1900Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. uplink-PowerControlToAddModListExt-v1910: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(mIMOParamV1900UplinkPowerControlToAddModListExtV1910Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.Uplink_PowerControlToAddModListExt_v1910 = make([]Uplink_PowerControlExt_v1900, n)
			for i := int64(0); i < n; i++ {
				if err := ie.Uplink_PowerControlToAddModListExt_v1910[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
