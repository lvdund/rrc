// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB23-IEs-r18 (line 4638).

var sIB23IEsR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PosConfigCommonNR-r18"},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
}

var sIB23IEsR18LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB23_IEs_r18 struct {
	Sl_PosConfigCommonNR_r18 SL_PosConfigCommonNR_r18
	LateNonCriticalExtension []byte
}

func (ie *SIB23_IEs_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB23IEsR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. sl-PosConfigCommonNR-r18: ref
	{
		if err := ie.Sl_PosConfigCommonNR_r18.Encode(e); err != nil {
			return err
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB23IEsR18LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB23_IEs_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB23IEsR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. sl-PosConfigCommonNR-r18: ref
	{
		if err := ie.Sl_PosConfigCommonNR_r18.Decode(d); err != nil {
			return err
		}
	}

	// 4. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeOctetString(sIB23IEsR18LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	return nil
}
