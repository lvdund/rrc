// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ReportCGI (line 13543).

var reportCGIConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "cellForWhichToReportCGI"},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	ReportCGI_Ext_UseAutonomousGaps_r16_Setup = 0
)

var reportCGIExtUseAutonomousGapsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ReportCGI_Ext_UseAutonomousGaps_r16_Setup},
}

type ReportCGI struct {
	CellForWhichToReportCGI PhysCellId
	UseAutonomousGaps_r16   *int64
}

func (ie *ReportCGI) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(reportCGIConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.UseAutonomousGaps_r16 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. cellForWhichToReportCGI: ref
	{
		if err := ie.CellForWhichToReportCGI.Encode(e); err != nil {
			return err
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
					{Name: "useAutonomousGaps-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.UseAutonomousGaps_r16 != nil}); err != nil {
				return err
			}

			if ie.UseAutonomousGaps_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.UseAutonomousGaps_r16, reportCGIExtUseAutonomousGapsR16Constraints); err != nil {
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

func (ie *ReportCGI) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(reportCGIConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. cellForWhichToReportCGI: ref
	{
		if err := ie.CellForWhichToReportCGI.Decode(d); err != nil {
			return err
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
				{Name: "useAutonomousGaps-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(reportCGIExtUseAutonomousGapsR16Constraints)
			if err != nil {
				return err
			}
			ie.UseAutonomousGaps_r16 = &v
		}
	}

	return nil
}
