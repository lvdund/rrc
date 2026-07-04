// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SSB-InfoNcell-r16 (line 15640).

var sSBInfoNcellR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "physicalCellId-r16"},
		{Name: "ssb-IndexNcell-r16", Optional: true},
		{Name: "ssb-Configuration-r16", Optional: true},
	},
}

type SSB_InfoNcell_r16 struct {
	PhysicalCellId_r16    PhysCellId
	Ssb_IndexNcell_r16    *SSB_Index
	Ssb_Configuration_r16 *SSB_Configuration_r16
}

func (ie *SSB_InfoNcell_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sSBInfoNcellR16Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ssb_IndexNcell_r16 != nil, ie.Ssb_Configuration_r16 != nil}); err != nil {
		return err
	}

	// 2. physicalCellId-r16: ref
	{
		if err := ie.PhysicalCellId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. ssb-IndexNcell-r16: ref
	{
		if ie.Ssb_IndexNcell_r16 != nil {
			if err := ie.Ssb_IndexNcell_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. ssb-Configuration-r16: ref
	{
		if ie.Ssb_Configuration_r16 != nil {
			if err := ie.Ssb_Configuration_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SSB_InfoNcell_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sSBInfoNcellR16Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. physicalCellId-r16: ref
	{
		if err := ie.PhysicalCellId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. ssb-IndexNcell-r16: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ssb_IndexNcell_r16 = new(SSB_Index)
			if err := ie.Ssb_IndexNcell_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. ssb-Configuration-r16: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Ssb_Configuration_r16 = new(SSB_Configuration_r16)
			if err := ie.Ssb_Configuration_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
