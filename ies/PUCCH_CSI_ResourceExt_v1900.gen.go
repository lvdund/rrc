// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PUCCH-CSI-ResourceExt-v1900 (line 12218).

var pUCCHCSIResourceExtV1900Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "symbolType-r19", Optional: true},
	},
}

const (
	PUCCH_CSI_ResourceExt_v1900_SymbolType_r19_Sbfd     = 0
	PUCCH_CSI_ResourceExt_v1900_SymbolType_r19_Non_Sbfd = 1
)

var pUCCHCSIResourceExtV1900SymbolTypeR19Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PUCCH_CSI_ResourceExt_v1900_SymbolType_r19_Sbfd, PUCCH_CSI_ResourceExt_v1900_SymbolType_r19_Non_Sbfd},
}

type PUCCH_CSI_ResourceExt_v1900 struct {
	SymbolType_r19 *int64
}

func (ie *PUCCH_CSI_ResourceExt_v1900) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pUCCHCSIResourceExtV1900Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SymbolType_r19 != nil}); err != nil {
		return err
	}

	// 3. symbolType-r19: enumerated
	{
		if ie.SymbolType_r19 != nil {
			if err := e.EncodeEnumerated(*ie.SymbolType_r19, pUCCHCSIResourceExtV1900SymbolTypeR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PUCCH_CSI_ResourceExt_v1900) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pUCCHCSIResourceExtV1900Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. symbolType-r19: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(pUCCHCSIResourceExtV1900SymbolTypeR19Constraints)
			if err != nil {
				return err
			}
			ie.SymbolType_r19 = &idx
		}
	}

	return nil
}
