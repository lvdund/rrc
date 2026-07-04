// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultsPerCellIdleEUTRA-r16 (line 9925).

var measResultsPerCellIdleEUTRAR16Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "eutra-PhysCellId-r16"},
		{Name: "measIdleResultEUTRA-r16"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

var measResultsPerCellIdleEUTRAR16MeasIdleResultEUTRAR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "rsrp-ResultEUTRA-r16", Optional: true},
		{Name: "rsrq-ResultEUTRA-r16", Optional: true},
	},
}

type MeasResultsPerCellIdleEUTRA_r16 struct {
	Eutra_PhysCellId_r16    EUTRA_PhysCellId
	MeasIdleResultEUTRA_r16 struct {
		Rsrp_ResultEUTRA_r16 *RSRP_RangeEUTRA
		Rsrq_ResultEUTRA_r16 *RSRQ_RangeEUTRA_r16
	}
	ValidityStatus_r18 *MeasurementValidityDuration_r18
}

func (ie *MeasResultsPerCellIdleEUTRA_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultsPerCellIdleEUTRAR16Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.ValidityStatus_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. eutra-PhysCellId-r16: ref
	{
		if err := ie.Eutra_PhysCellId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 3. measIdleResultEUTRA-r16: sequence
	{
		{
			c := &ie.MeasIdleResultEUTRA_r16
			measResultsPerCellIdleEUTRAR16MeasIdleResultEUTRAR16Seq := e.NewSequenceEncoder(measResultsPerCellIdleEUTRAR16MeasIdleResultEUTRAR16Constraints)
			if err := measResultsPerCellIdleEUTRAR16MeasIdleResultEUTRAR16Seq.EncodePreamble([]bool{c.Rsrp_ResultEUTRA_r16 != nil, c.Rsrq_ResultEUTRA_r16 != nil}); err != nil {
				return err
			}
			if c.Rsrp_ResultEUTRA_r16 != nil {
				if err := c.Rsrp_ResultEUTRA_r16.Encode(e); err != nil {
					return err
				}
			}
			if c.Rsrq_ResultEUTRA_r16 != nil {
				if err := c.Rsrq_ResultEUTRA_r16.Encode(e); err != nil {
					return err
				}
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "validityStatus-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.ValidityStatus_r18 != nil}); err != nil {
				return err
			}

			if ie.ValidityStatus_r18 != nil {
				if err := ie.ValidityStatus_r18.Encode(ex); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		if err := seq.EncodeExtensionAdditions(extPresence, extValues); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasResultsPerCellIdleEUTRA_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultsPerCellIdleEUTRAR16Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. eutra-PhysCellId-r16: ref
	{
		if err := ie.Eutra_PhysCellId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 3. measIdleResultEUTRA-r16: sequence
	{
		{
			c := &ie.MeasIdleResultEUTRA_r16
			measResultsPerCellIdleEUTRAR16MeasIdleResultEUTRAR16Seq := d.NewSequenceDecoder(measResultsPerCellIdleEUTRAR16MeasIdleResultEUTRAR16Constraints)
			if err := measResultsPerCellIdleEUTRAR16MeasIdleResultEUTRAR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if measResultsPerCellIdleEUTRAR16MeasIdleResultEUTRAR16Seq.IsComponentPresent(0) {
				c.Rsrp_ResultEUTRA_r16 = new(RSRP_RangeEUTRA)
				if err := (*c.Rsrp_ResultEUTRA_r16).Decode(d); err != nil {
					return err
				}
			}
			if measResultsPerCellIdleEUTRAR16MeasIdleResultEUTRAR16Seq.IsComponentPresent(1) {
				c.Rsrq_ResultEUTRA_r16 = new(RSRQ_RangeEUTRA_r16)
				if err := (*c.Rsrq_ResultEUTRA_r16).Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// Extension additions as open types.
	extValues, err := seq.DecodeExtensionAdditions()
	if err != nil {
		return err
	}
	extValueIndex := 0

	// Extension group 0:
	if seq.IsExtensionPresent(0) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "validityStatus-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			ie.ValidityStatus_r18 = new(MeasurementValidityDuration_r18)
			if err := ie.ValidityStatus_r18.Decode(dx); err != nil {
				return err
			}
		}
	}

	return nil
}
