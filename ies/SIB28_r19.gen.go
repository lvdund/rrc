// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB28-r19 (line 4786).

var sIB28R19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ntn-NeighCellConfigListBLCE-r19", Optional: true},
		{Name: "ntn-NeighCellConfigListNB-r19", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
}

var sIB28R19LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB28_r19 struct {
	Ntn_NeighCellConfigListBLCE_r19 *NTN_NeighCellConfigListIoT_r19
	Ntn_NeighCellConfigListNB_r19   *NTN_NeighCellConfigListIoT_r19
	LateNonCriticalExtension        []byte
}

func (ie *SIB28_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB28R19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ntn_NeighCellConfigListBLCE_r19 != nil, ie.Ntn_NeighCellConfigListNB_r19 != nil, ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. ntn-NeighCellConfigListBLCE-r19: ref
	{
		if ie.Ntn_NeighCellConfigListBLCE_r19 != nil {
			if err := ie.Ntn_NeighCellConfigListBLCE_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. ntn-NeighCellConfigListNB-r19: ref
	{
		if ie.Ntn_NeighCellConfigListNB_r19 != nil {
			if err := ie.Ntn_NeighCellConfigListNB_r19.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB28R19LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB28_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB28R19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ntn-NeighCellConfigListBLCE-r19: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ntn_NeighCellConfigListBLCE_r19 = new(NTN_NeighCellConfigListIoT_r19)
			if err := ie.Ntn_NeighCellConfigListBLCE_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. ntn-NeighCellConfigListNB-r19: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ntn_NeighCellConfigListNB_r19 = new(NTN_NeighCellConfigListIoT_r19)
			if err := ie.Ntn_NeighCellConfigListNB_r19.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeOctetString(sIB28R19LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	return nil
}
