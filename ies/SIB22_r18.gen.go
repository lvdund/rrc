// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB22-r18 (line 4612).

var sIB22R18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "atg-Config-r18", Optional: true},
		{Name: "hs-ATG-CellReselectionSet-r18", Optional: true},
		{Name: "atg-NeighCellConfigList-r18", Optional: true},
		{Name: "lateNonCriticalExtension", Optional: true},
	},
}

const (
	SIB22_r18_Hs_ATG_CellReselectionSet_r18_True = 0
)

var sIB22R18HsATGCellReselectionSetR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB22_r18_Hs_ATG_CellReselectionSet_r18_True},
}

var sIB22R18LateNonCriticalExtensionConstraints = per.SizeConstraints{}

type SIB22_r18 struct {
	Atg_Config_r18                *ATG_Config_r18
	Hs_ATG_CellReselectionSet_r18 *int64
	Atg_NeighCellConfigList_r18   *ATG_NeighCellConfigList_r18
	LateNonCriticalExtension      []byte
}

func (ie *SIB22_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB22R18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Atg_Config_r18 != nil, ie.Hs_ATG_CellReselectionSet_r18 != nil, ie.Atg_NeighCellConfigList_r18 != nil, ie.LateNonCriticalExtension != nil}); err != nil {
		return err
	}

	// 3. atg-Config-r18: ref
	{
		if ie.Atg_Config_r18 != nil {
			if err := ie.Atg_Config_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. hs-ATG-CellReselectionSet-r18: enumerated
	{
		if ie.Hs_ATG_CellReselectionSet_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Hs_ATG_CellReselectionSet_r18, sIB22R18HsATGCellReselectionSetR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. atg-NeighCellConfigList-r18: ref
	{
		if ie.Atg_NeighCellConfigList_r18 != nil {
			if err := ie.Atg_NeighCellConfigList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, sIB22R18LateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SIB22_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB22R18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. atg-Config-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Atg_Config_r18 = new(ATG_Config_r18)
			if err := ie.Atg_Config_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. hs-ATG-CellReselectionSet-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sIB22R18HsATGCellReselectionSetR18Constraints)
			if err != nil {
				return err
			}
			ie.Hs_ATG_CellReselectionSet_r18 = &idx
		}
	}

	// 5. atg-NeighCellConfigList-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Atg_NeighCellConfigList_r18 = new(ATG_NeighCellConfigList_r18)
			if err := ie.Atg_NeighCellConfigList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeOctetString(sIB22R18LateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	return nil
}
