// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParametersXDD-Diff (line 21305).

var measAndMobParametersXDDDiffConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "intraAndInterF-MeasAndReport", Optional: true},
		{Name: "eventA-MeasAndReport", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
	},
}

const (
	MeasAndMobParametersXDD_Diff_IntraAndInterF_MeasAndReport_Supported = 0
)

var measAndMobParametersXDDDiffIntraAndInterFMeasAndReportConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersXDD_Diff_IntraAndInterF_MeasAndReport_Supported},
}

const (
	MeasAndMobParametersXDD_Diff_EventA_MeasAndReport_Supported = 0
)

var measAndMobParametersXDDDiffEventAMeasAndReportConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersXDD_Diff_EventA_MeasAndReport_Supported},
}

const (
	MeasAndMobParametersXDD_Diff_Ext_HandoverInterF_Supported = 0
)

var measAndMobParametersXDDDiffExtHandoverInterFConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersXDD_Diff_Ext_HandoverInterF_Supported},
}

const (
	MeasAndMobParametersXDD_Diff_Ext_HandoverLTE_EPC_Supported = 0
)

var measAndMobParametersXDDDiffExtHandoverLTEEPCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersXDD_Diff_Ext_HandoverLTE_EPC_Supported},
}

const (
	MeasAndMobParametersXDD_Diff_Ext_HandoverLTE_5GC_Supported = 0
)

var measAndMobParametersXDDDiffExtHandoverLTE5GCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersXDD_Diff_Ext_HandoverLTE_5GC_Supported},
}

const (
	MeasAndMobParametersXDD_Diff_Ext_Sftd_MeasNR_Neigh_Supported = 0
)

var measAndMobParametersXDDDiffExtSftdMeasNRNeighConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersXDD_Diff_Ext_Sftd_MeasNR_Neigh_Supported},
}

const (
	MeasAndMobParametersXDD_Diff_Ext_Sftd_MeasNR_Neigh_DRX_Supported = 0
)

var measAndMobParametersXDDDiffExtSftdMeasNRNeighDRXConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersXDD_Diff_Ext_Sftd_MeasNR_Neigh_DRX_Supported},
}

const (
	MeasAndMobParametersXDD_Diff_Ext_Dummy_Supported = 0
)

var measAndMobParametersXDDDiffExtDummyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersXDD_Diff_Ext_Dummy_Supported},
}

type MeasAndMobParametersXDD_Diff struct {
	IntraAndInterF_MeasAndReport *int64
	EventA_MeasAndReport         *int64
	HandoverInterF               *int64
	HandoverLTE_EPC              *int64
	HandoverLTE_5GC              *int64
	Sftd_MeasNR_Neigh            *int64
	Sftd_MeasNR_Neigh_DRX        *int64
	Dummy                        *int64
}

func (ie *MeasAndMobParametersXDD_Diff) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersXDDDiffConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.HandoverInterF != nil || ie.HandoverLTE_EPC != nil || ie.HandoverLTE_5GC != nil
	hasExtGroup1 := ie.Sftd_MeasNR_Neigh != nil || ie.Sftd_MeasNR_Neigh_DRX != nil
	hasExtGroup2 := ie.Dummy != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IntraAndInterF_MeasAndReport != nil, ie.EventA_MeasAndReport != nil}); err != nil {
		return err
	}

	// 3. intraAndInterF-MeasAndReport: enumerated
	{
		if ie.IntraAndInterF_MeasAndReport != nil {
			if err := e.EncodeEnumerated(*ie.IntraAndInterF_MeasAndReport, measAndMobParametersXDDDiffIntraAndInterFMeasAndReportConstraints); err != nil {
				return err
			}
		}
	}

	// 4. eventA-MeasAndReport: enumerated
	{
		if ie.EventA_MeasAndReport != nil {
			if err := e.EncodeEnumerated(*ie.EventA_MeasAndReport, measAndMobParametersXDDDiffEventAMeasAndReportConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2}
		var extValues [][]byte

		// Extension group 0:
		if hasExtGroup0 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "handoverInterF", Optional: true},
					{Name: "handoverLTE-EPC", Optional: true},
					{Name: "handoverLTE-5GC", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.HandoverInterF != nil, ie.HandoverLTE_EPC != nil, ie.HandoverLTE_5GC != nil}); err != nil {
				return err
			}

			if ie.HandoverInterF != nil {
				if err := ex.EncodeEnumerated(*ie.HandoverInterF, measAndMobParametersXDDDiffExtHandoverInterFConstraints); err != nil {
					return err
				}
			}

			if ie.HandoverLTE_EPC != nil {
				if err := ex.EncodeEnumerated(*ie.HandoverLTE_EPC, measAndMobParametersXDDDiffExtHandoverLTEEPCConstraints); err != nil {
					return err
				}
			}

			if ie.HandoverLTE_5GC != nil {
				if err := ex.EncodeEnumerated(*ie.HandoverLTE_5GC, measAndMobParametersXDDDiffExtHandoverLTE5GCConstraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 1:
		if hasExtGroup1 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "sftd-MeasNR-Neigh", Optional: true},
					{Name: "sftd-MeasNR-Neigh-DRX", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Sftd_MeasNR_Neigh != nil, ie.Sftd_MeasNR_Neigh_DRX != nil}); err != nil {
				return err
			}

			if ie.Sftd_MeasNR_Neigh != nil {
				if err := ex.EncodeEnumerated(*ie.Sftd_MeasNR_Neigh, measAndMobParametersXDDDiffExtSftdMeasNRNeighConstraints); err != nil {
					return err
				}
			}

			if ie.Sftd_MeasNR_Neigh_DRX != nil {
				if err := ex.EncodeEnumerated(*ie.Sftd_MeasNR_Neigh_DRX, measAndMobParametersXDDDiffExtSftdMeasNRNeighDRXConstraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 2:
		if hasExtGroup2 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "dummy", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Dummy != nil}); err != nil {
				return err
			}

			if ie.Dummy != nil {
				if err := ex.EncodeEnumerated(*ie.Dummy, measAndMobParametersXDDDiffExtDummyConstraints); err != nil {
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

func (ie *MeasAndMobParametersXDD_Diff) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersXDDDiffConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. intraAndInterF-MeasAndReport: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(measAndMobParametersXDDDiffIntraAndInterFMeasAndReportConstraints)
			if err != nil {
				return err
			}
			ie.IntraAndInterF_MeasAndReport = &idx
		}
	}

	// 4. eventA-MeasAndReport: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(measAndMobParametersXDDDiffEventAMeasAndReportConstraints)
			if err != nil {
				return err
			}
			ie.EventA_MeasAndReport = &idx
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
				{Name: "handoverInterF", Optional: true},
				{Name: "handoverLTE-EPC", Optional: true},
				{Name: "handoverLTE-5GC", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measAndMobParametersXDDDiffExtHandoverInterFConstraints)
			if err != nil {
				return err
			}
			ie.HandoverInterF = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(measAndMobParametersXDDDiffExtHandoverLTEEPCConstraints)
			if err != nil {
				return err
			}
			ie.HandoverLTE_EPC = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(measAndMobParametersXDDDiffExtHandoverLTE5GCConstraints)
			if err != nil {
				return err
			}
			ie.HandoverLTE_5GC = &v
		}
	}

	// Extension group 1:
	if seq.IsExtensionPresent(1) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "sftd-MeasNR-Neigh", Optional: true},
				{Name: "sftd-MeasNR-Neigh-DRX", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measAndMobParametersXDDDiffExtSftdMeasNRNeighConstraints)
			if err != nil {
				return err
			}
			ie.Sftd_MeasNR_Neigh = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(measAndMobParametersXDDDiffExtSftdMeasNRNeighDRXConstraints)
			if err != nil {
				return err
			}
			ie.Sftd_MeasNR_Neigh_DRX = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "dummy", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measAndMobParametersXDDDiffExtDummyConstraints)
			if err != nil {
				return err
			}
			ie.Dummy = &v
		}
	}

	return nil
}
