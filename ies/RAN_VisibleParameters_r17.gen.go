// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RAN-VisibleParameters-r17 (line 26036).

var rANVisibleParametersR17Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ran-VisiblePeriodicity-r17", Optional: true},
		{Name: "numberOfBufferLevelEntries-r17", Optional: true},
		{Name: "reportPlayoutDelayForMediaStartup-r17", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
	},
}

const (
	RAN_VisibleParameters_r17_Ran_VisiblePeriodicity_r17_Ms120  = 0
	RAN_VisibleParameters_r17_Ran_VisiblePeriodicity_r17_Ms240  = 1
	RAN_VisibleParameters_r17_Ran_VisiblePeriodicity_r17_Ms480  = 2
	RAN_VisibleParameters_r17_Ran_VisiblePeriodicity_r17_Ms640  = 3
	RAN_VisibleParameters_r17_Ran_VisiblePeriodicity_r17_Ms1024 = 4
)

var rANVisibleParametersR17RanVisiblePeriodicityR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RAN_VisibleParameters_r17_Ran_VisiblePeriodicity_r17_Ms120, RAN_VisibleParameters_r17_Ran_VisiblePeriodicity_r17_Ms240, RAN_VisibleParameters_r17_Ran_VisiblePeriodicity_r17_Ms480, RAN_VisibleParameters_r17_Ran_VisiblePeriodicity_r17_Ms640, RAN_VisibleParameters_r17_Ran_VisiblePeriodicity_r17_Ms1024},
}

var rANVisibleParametersR17NumberOfBufferLevelEntriesR17Constraints = per.Constrained(1, 8)

const (
	RAN_VisibleParameters_r17_Ext_Ran_VisibleReportingSRB_r18_Srb4 = 0
	RAN_VisibleParameters_r17_Ext_Ran_VisibleReportingSRB_r18_Srb5 = 1
)

var rANVisibleParametersR17ExtRanVisibleReportingSRBR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RAN_VisibleParameters_r17_Ext_Ran_VisibleReportingSRB_r18_Srb4, RAN_VisibleParameters_r17_Ext_Ran_VisibleReportingSRB_r18_Srb5},
}

type RAN_VisibleParameters_r17 struct {
	Ran_VisiblePeriodicity_r17            *int64
	NumberOfBufferLevelEntries_r17        *int64
	ReportPlayoutDelayForMediaStartup_r17 *bool
	Ran_VisibleReportingSRB_r18           *int64
}

func (ie *RAN_VisibleParameters_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rANVisibleParametersR17Constraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.Ran_VisibleReportingSRB_r18 != nil
	hasExtensions := hasExtGroup0

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ran_VisiblePeriodicity_r17 != nil, ie.NumberOfBufferLevelEntries_r17 != nil, ie.ReportPlayoutDelayForMediaStartup_r17 != nil}); err != nil {
		return err
	}

	// 3. ran-VisiblePeriodicity-r17: enumerated
	{
		if ie.Ran_VisiblePeriodicity_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ran_VisiblePeriodicity_r17, rANVisibleParametersR17RanVisiblePeriodicityR17Constraints); err != nil {
				return err
			}
		}
	}

	// 4. numberOfBufferLevelEntries-r17: integer
	{
		if ie.NumberOfBufferLevelEntries_r17 != nil {
			if err := e.EncodeInteger(*ie.NumberOfBufferLevelEntries_r17, rANVisibleParametersR17NumberOfBufferLevelEntriesR17Constraints); err != nil {
				return err
			}
		}
	}

	// 5. reportPlayoutDelayForMediaStartup-r17: boolean
	{
		if ie.ReportPlayoutDelayForMediaStartup_r17 != nil {
			if err := e.EncodeBoolean(*ie.ReportPlayoutDelayForMediaStartup_r17); err != nil {
				return err
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
					{Name: "ran-VisibleReportingSRB-r18", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Ran_VisibleReportingSRB_r18 != nil}); err != nil {
				return err
			}

			if ie.Ran_VisibleReportingSRB_r18 != nil {
				if err := ex.EncodeEnumerated(*ie.Ran_VisibleReportingSRB_r18, rANVisibleParametersR17ExtRanVisibleReportingSRBR18Constraints); err != nil {
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

func (ie *RAN_VisibleParameters_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rANVisibleParametersR17Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ran-VisiblePeriodicity-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(rANVisibleParametersR17RanVisiblePeriodicityR17Constraints)
			if err != nil {
				return err
			}
			ie.Ran_VisiblePeriodicity_r17 = &idx
		}
	}

	// 4. numberOfBufferLevelEntries-r17: integer
	{
		if seq.IsComponentPresent(1) {
			v, err := d.DecodeInteger(rANVisibleParametersR17NumberOfBufferLevelEntriesR17Constraints)
			if err != nil {
				return err
			}
			ie.NumberOfBufferLevelEntries_r17 = &v
		}
	}

	// 5. reportPlayoutDelayForMediaStartup-r17: boolean
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeBoolean()
			if err != nil {
				return err
			}
			ie.ReportPlayoutDelayForMediaStartup_r17 = &v
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
				{Name: "ran-VisibleReportingSRB-r18", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(rANVisibleParametersR17ExtRanVisibleReportingSRBR18Constraints)
			if err != nil {
				return err
			}
			ie.Ran_VisibleReportingSRB_r18 = &v
		}
	}

	return nil
}
