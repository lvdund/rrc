// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: LowBandCA-Switching-r19 (line 5843).

var lowBandCASwitchingR19Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "switchingPattern-r19", Optional: true},
		{Name: "gapDurationPCelltoSCell-r19", Optional: true},
		{Name: "gapDurationSCelltoPCell-r19", Optional: true},
	},
}

var lowBandCASwitchingR19SwitchingPatternR19Constraints = per.FixedSize(40)

var lowBandCASwitchingR19GapDurationPCelltoSCellR19Constraints = per.Constrained(1, 3)

var lowBandCASwitchingR19GapDurationSCelltoPCellR19Constraints = per.Constrained(1, 31)

type LowBandCA_Switching_r19 struct {
	SwitchingPattern_r19        *per.BitString
	GapDurationPCelltoSCell_r19 *int64
	GapDurationSCelltoPCell_r19 *int64
}

func (ie *LowBandCA_Switching_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(lowBandCASwitchingR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SwitchingPattern_r19 != nil, ie.GapDurationPCelltoSCell_r19 != nil, ie.GapDurationSCelltoPCell_r19 != nil}); err != nil {
		return err
	}

	// 3. switchingPattern-r19: bit-string
	{
		if ie.SwitchingPattern_r19 != nil {
			if err := e.EncodeBitString(*ie.SwitchingPattern_r19, lowBandCASwitchingR19SwitchingPatternR19Constraints); err != nil {
				return err
			}
		}
	}

	// 4. gapDurationPCelltoSCell-r19: integer
	{
		if ie.GapDurationPCelltoSCell_r19 != nil {
			if err := e.EncodeInteger(*ie.GapDurationPCelltoSCell_r19, lowBandCASwitchingR19GapDurationPCelltoSCellR19Constraints); err != nil {
				return err
			}
		}
	}

	// 5. gapDurationSCelltoPCell-r19: integer
	{
		if ie.GapDurationSCelltoPCell_r19 != nil {
			if err := e.EncodeInteger(*ie.GapDurationSCelltoPCell_r19, lowBandCASwitchingR19GapDurationSCelltoPCellR19Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *LowBandCA_Switching_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(lowBandCASwitchingR19Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. switchingPattern-r19: bit-string
	{
		if seq.IsComponentPresent(0) {
			v, err := d.DecodeBitString(lowBandCASwitchingR19SwitchingPatternR19Constraints)
			if err != nil {
				return err
			}
			ie.SwitchingPattern_r19 = &v
		}
	}

	// 4. gapDurationPCelltoSCell-r19: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(lowBandCASwitchingR19GapDurationPCelltoSCellR19Constraints)
			if err != nil {
				return err
			}
			ie.GapDurationPCelltoSCell_r19 = &v
		}
	}

	// 5. gapDurationSCelltoPCell-r19: integer
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeInteger(lowBandCASwitchingR19GapDurationSCelltoPCellR19Constraints)
			if err != nil {
				return err
			}
			ie.GapDurationSCelltoPCell_r19 = &v
		}
	}

	return nil
}
