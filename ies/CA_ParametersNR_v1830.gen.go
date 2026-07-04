// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNR-v1830 (line 17849).

var cAParametersNRV1830Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "intraFreqL1-MeasConfig-r18", Optional: true},
		{Name: "interFreqL1-MeasConfig-r18", Optional: true},
		{Name: "currentSpCellInclL1-Report-r18", Optional: true},
		{Name: "multiCellL1-measRTD-greaterThan-CP-r18", Optional: true},
		{Name: "interFreqSSB-L1-MeasWithoutGaps-r18", Optional: true},
		{Name: "maxFreqLayersL1-Meas-r18", Optional: true},
		{Name: "maxNeighCellsPerFreqLayerL1-Meas-r18", Optional: true},
		{Name: "supportedMaxCellsWithoutGapsL1-Meas-r18", Optional: true},
		{Name: "supportedMaxSSB-WithinSlotL1-Meas-r18", Optional: true},
		{Name: "dummy", Optional: true},
		{Name: "supportedMaxSSB-L1-Meas-r18", Optional: true},
		{Name: "qcl-MultiCellDCI-1-3-r18", Optional: true},
		{Name: "bwp-SwitchingDCI-0-3-And-1-3-r18", Optional: true},
	},
}

const (
	CA_ParametersNR_v1830_CurrentSpCellInclL1_Report_r18_Supported = 0
)

var cAParametersNRV1830CurrentSpCellInclL1ReportR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1830_CurrentSpCellInclL1_Report_r18_Supported},
}

const (
	CA_ParametersNR_v1830_MultiCellL1_MeasRTD_GreaterThan_CP_r18_Supported = 0
)

var cAParametersNRV1830MultiCellL1MeasRTDGreaterThanCPR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1830_MultiCellL1_MeasRTD_GreaterThan_CP_r18_Supported},
}

const (
	CA_ParametersNR_v1830_InterFreqSSB_L1_MeasWithoutGaps_r18_Supported = 0
)

var cAParametersNRV1830InterFreqSSBL1MeasWithoutGapsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1830_InterFreqSSB_L1_MeasWithoutGaps_r18_Supported},
}

var cAParametersNRV1830SupportedMaxCellsWithoutGapsL1MeasR18Constraints = per.Constrained(1, 24)

const (
	CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N1  = 0
	CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N2  = 1
	CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N3  = 2
	CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N4  = 3
	CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N5  = 4
	CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N6  = 5
	CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N7  = 6
	CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N8  = 7
	CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N16 = 8
	CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N32 = 9
	CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N48 = 10
	CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N64 = 11
)

var cAParametersNRV1830SupportedMaxSSBWithinSlotL1MeasR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N1, CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N2, CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N3, CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N4, CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N5, CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N6, CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N7, CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N8, CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N16, CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N32, CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N48, CA_ParametersNR_v1830_SupportedMaxSSB_WithinSlotL1_Meas_r18_N64},
}

const (
	CA_ParametersNR_v1830_SupportedMaxSSB_L1_Meas_r18_N2  = 0
	CA_ParametersNR_v1830_SupportedMaxSSB_L1_Meas_r18_N4  = 1
	CA_ParametersNR_v1830_SupportedMaxSSB_L1_Meas_r18_N8  = 2
	CA_ParametersNR_v1830_SupportedMaxSSB_L1_Meas_r18_N12 = 3
	CA_ParametersNR_v1830_SupportedMaxSSB_L1_Meas_r18_N16 = 4
	CA_ParametersNR_v1830_SupportedMaxSSB_L1_Meas_r18_N32 = 5
	CA_ParametersNR_v1830_SupportedMaxSSB_L1_Meas_r18_N64 = 6
)

var cAParametersNRV1830SupportedMaxSSBL1MeasR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1830_SupportedMaxSSB_L1_Meas_r18_N2, CA_ParametersNR_v1830_SupportedMaxSSB_L1_Meas_r18_N4, CA_ParametersNR_v1830_SupportedMaxSSB_L1_Meas_r18_N8, CA_ParametersNR_v1830_SupportedMaxSSB_L1_Meas_r18_N12, CA_ParametersNR_v1830_SupportedMaxSSB_L1_Meas_r18_N16, CA_ParametersNR_v1830_SupportedMaxSSB_L1_Meas_r18_N32, CA_ParametersNR_v1830_SupportedMaxSSB_L1_Meas_r18_N64},
}

const (
	CA_ParametersNR_v1830_Qcl_MultiCellDCI_1_3_r18_Diff = 0
	CA_ParametersNR_v1830_Qcl_MultiCellDCI_1_3_r18_Both = 1
)

var cAParametersNRV1830QclMultiCellDCI13R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1830_Qcl_MultiCellDCI_1_3_r18_Diff, CA_ParametersNR_v1830_Qcl_MultiCellDCI_1_3_r18_Both},
}

const (
	CA_ParametersNR_v1830_Bwp_SwitchingDCI_0_3_And_1_3_r18_Supported = 0
)

var cAParametersNRV1830BwpSwitchingDCI03And13R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1830_Bwp_SwitchingDCI_0_3_And_1_3_r18_Supported},
}

const (
	CA_ParametersNR_v1830_IntraFreqL1_MeasConfig_r18_SupportedMaxReportBeamsReports_r18_N1  = 0
	CA_ParametersNR_v1830_IntraFreqL1_MeasConfig_r18_SupportedMaxReportBeamsReports_r18_N2  = 1
	CA_ParametersNR_v1830_IntraFreqL1_MeasConfig_r18_SupportedMaxReportBeamsReports_r18_N3  = 2
	CA_ParametersNR_v1830_IntraFreqL1_MeasConfig_r18_SupportedMaxReportBeamsReports_r18_N4  = 3
	CA_ParametersNR_v1830_IntraFreqL1_MeasConfig_r18_SupportedMaxReportBeamsReports_r18_N6  = 4
	CA_ParametersNR_v1830_IntraFreqL1_MeasConfig_r18_SupportedMaxReportBeamsReports_r18_N8  = 5
	CA_ParametersNR_v1830_IntraFreqL1_MeasConfig_r18_SupportedMaxReportBeamsReports_r18_N9  = 6
	CA_ParametersNR_v1830_IntraFreqL1_MeasConfig_r18_SupportedMaxReportBeamsReports_r18_N12 = 7
	CA_ParametersNR_v1830_IntraFreqL1_MeasConfig_r18_SupportedMaxReportBeamsReports_r18_N16 = 8
)

var cAParametersNRV1830IntraFreqL1MeasConfigR18SupportedMaxReportBeamsReportsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1830_IntraFreqL1_MeasConfig_r18_SupportedMaxReportBeamsReports_r18_N1, CA_ParametersNR_v1830_IntraFreqL1_MeasConfig_r18_SupportedMaxReportBeamsReports_r18_N2, CA_ParametersNR_v1830_IntraFreqL1_MeasConfig_r18_SupportedMaxReportBeamsReports_r18_N3, CA_ParametersNR_v1830_IntraFreqL1_MeasConfig_r18_SupportedMaxReportBeamsReports_r18_N4, CA_ParametersNR_v1830_IntraFreqL1_MeasConfig_r18_SupportedMaxReportBeamsReports_r18_N6, CA_ParametersNR_v1830_IntraFreqL1_MeasConfig_r18_SupportedMaxReportBeamsReports_r18_N8, CA_ParametersNR_v1830_IntraFreqL1_MeasConfig_r18_SupportedMaxReportBeamsReports_r18_N9, CA_ParametersNR_v1830_IntraFreqL1_MeasConfig_r18_SupportedMaxReportBeamsReports_r18_N12, CA_ParametersNR_v1830_IntraFreqL1_MeasConfig_r18_SupportedMaxReportBeamsReports_r18_N16},
}

const (
	CA_ParametersNR_v1830_InterFreqL1_MeasConfig_r18_SupportedMaxIntraInterFreqBeamsReports_r18_N1  = 0
	CA_ParametersNR_v1830_InterFreqL1_MeasConfig_r18_SupportedMaxIntraInterFreqBeamsReports_r18_N2  = 1
	CA_ParametersNR_v1830_InterFreqL1_MeasConfig_r18_SupportedMaxIntraInterFreqBeamsReports_r18_N3  = 2
	CA_ParametersNR_v1830_InterFreqL1_MeasConfig_r18_SupportedMaxIntraInterFreqBeamsReports_r18_N4  = 3
	CA_ParametersNR_v1830_InterFreqL1_MeasConfig_r18_SupportedMaxIntraInterFreqBeamsReports_r18_N6  = 4
	CA_ParametersNR_v1830_InterFreqL1_MeasConfig_r18_SupportedMaxIntraInterFreqBeamsReports_r18_N8  = 5
	CA_ParametersNR_v1830_InterFreqL1_MeasConfig_r18_SupportedMaxIntraInterFreqBeamsReports_r18_N9  = 6
	CA_ParametersNR_v1830_InterFreqL1_MeasConfig_r18_SupportedMaxIntraInterFreqBeamsReports_r18_N12 = 7
	CA_ParametersNR_v1830_InterFreqL1_MeasConfig_r18_SupportedMaxIntraInterFreqBeamsReports_r18_N16 = 8
)

var cAParametersNRV1830InterFreqL1MeasConfigR18SupportedMaxIntraInterFreqBeamsReportsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1830_InterFreqL1_MeasConfig_r18_SupportedMaxIntraInterFreqBeamsReports_r18_N1, CA_ParametersNR_v1830_InterFreqL1_MeasConfig_r18_SupportedMaxIntraInterFreqBeamsReports_r18_N2, CA_ParametersNR_v1830_InterFreqL1_MeasConfig_r18_SupportedMaxIntraInterFreqBeamsReports_r18_N3, CA_ParametersNR_v1830_InterFreqL1_MeasConfig_r18_SupportedMaxIntraInterFreqBeamsReports_r18_N4, CA_ParametersNR_v1830_InterFreqL1_MeasConfig_r18_SupportedMaxIntraInterFreqBeamsReports_r18_N6, CA_ParametersNR_v1830_InterFreqL1_MeasConfig_r18_SupportedMaxIntraInterFreqBeamsReports_r18_N8, CA_ParametersNR_v1830_InterFreqL1_MeasConfig_r18_SupportedMaxIntraInterFreqBeamsReports_r18_N9, CA_ParametersNR_v1830_InterFreqL1_MeasConfig_r18_SupportedMaxIntraInterFreqBeamsReports_r18_N12, CA_ParametersNR_v1830_InterFreqL1_MeasConfig_r18_SupportedMaxIntraInterFreqBeamsReports_r18_N16},
}

var cAParametersNRV1830MaxFreqLayersL1MeasR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedMaxIntraInterFreqLayersWithoutGaps-r18", Optional: true},
		{Name: "supportedMaxInterFreqLayersWithGaps-r18", Optional: true},
	},
}

var cAParametersNRV1830MaxNeighCellsPerFreqLayerL1MeasR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedMaxNeighCellsPerFreqLayersWithoutGaps-r18", Optional: true},
		{Name: "supportedMaxNeighCellsPerFreqLayersWithGaps-r18", Optional: true},
	},
}

var cAParametersNRV1830DummyConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedMaxSSB-PerFreqLayerWithoutGaps-r18", Optional: true},
		{Name: "supportedMaxSSB-PerFreqLayerWithGaps-r18", Optional: true},
	},
}

type CA_ParametersNR_v1830 struct {
	IntraFreqL1_MeasConfig_r18 *struct {
		SupportedMaxIntraFreqCellsConfig_r18                int64
		SupportedMaxIntraFreqCellsPerReport_r18             int64
		SupportedMaxReportBeamsPerReportedCell_r18          int64
		SupportedMaxReportBeamsReports_r18                  int64
		SupportedMaxAperiodic_LTM_CSI_ReportConfig_r18      int64
		SupportedMaxPeriodic_LTM_CSI_ReportConfig_r18       int64
		SupportedMaxSemiPersistent_LTM_CSI_ReportConfig_r18 int64
	}
	InterFreqL1_MeasConfig_r18 *struct {
		SupportedMaxIntraInterFreqCellsConfig_r18         int64
		SupportedMaxIntraInterFreqCellsPerReport_r18      int64
		SupportedMaxIntraInterFreqBeamsPerCellReports_r18 int64
		SupportedMaxIntraInterFreqBeamsReports_r18        int64
	}
	CurrentSpCellInclL1_Report_r18         *int64
	MultiCellL1_MeasRTD_GreaterThan_CP_r18 *int64
	InterFreqSSB_L1_MeasWithoutGaps_r18    *int64
	MaxFreqLayersL1_Meas_r18               *struct {
		SupportedMaxIntraInterFreqLayersWithoutGaps_r18 *int64
		SupportedMaxInterFreqLayersWithGaps_r18         *int64
	}
	MaxNeighCellsPerFreqLayerL1_Meas_r18 *struct {
		SupportedMaxNeighCellsPerFreqLayersWithoutGaps_r18 *int64
		SupportedMaxNeighCellsPerFreqLayersWithGaps_r18    *int64
	}
	SupportedMaxCellsWithoutGapsL1_Meas_r18 *int64
	SupportedMaxSSB_WithinSlotL1_Meas_r18   *int64
	Dummy                                   *struct {
		SupportedMaxSSB_PerFreqLayerWithoutGaps_r18 *int64
		SupportedMaxSSB_PerFreqLayerWithGaps_r18    *int64
	}
	SupportedMaxSSB_L1_Meas_r18      *int64
	Qcl_MultiCellDCI_1_3_r18         *int64
	Bwp_SwitchingDCI_0_3_And_1_3_r18 *int64
}

func (ie *CA_ParametersNR_v1830) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRV1830Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IntraFreqL1_MeasConfig_r18 != nil, ie.InterFreqL1_MeasConfig_r18 != nil, ie.CurrentSpCellInclL1_Report_r18 != nil, ie.MultiCellL1_MeasRTD_GreaterThan_CP_r18 != nil, ie.InterFreqSSB_L1_MeasWithoutGaps_r18 != nil, ie.MaxFreqLayersL1_Meas_r18 != nil, ie.MaxNeighCellsPerFreqLayerL1_Meas_r18 != nil, ie.SupportedMaxCellsWithoutGapsL1_Meas_r18 != nil, ie.SupportedMaxSSB_WithinSlotL1_Meas_r18 != nil, ie.Dummy != nil, ie.SupportedMaxSSB_L1_Meas_r18 != nil, ie.Qcl_MultiCellDCI_1_3_r18 != nil, ie.Bwp_SwitchingDCI_0_3_And_1_3_r18 != nil}); err != nil {
		return err
	}

	// 2. intraFreqL1-MeasConfig-r18: sequence
	{
		if ie.IntraFreqL1_MeasConfig_r18 != nil {
			c := ie.IntraFreqL1_MeasConfig_r18
			if err := e.EncodeInteger(c.SupportedMaxIntraFreqCellsConfig_r18, per.Constrained(1, 8)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.SupportedMaxIntraFreqCellsPerReport_r18, per.Constrained(1, 4)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.SupportedMaxReportBeamsPerReportedCell_r18, per.Constrained(1, 4)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.SupportedMaxReportBeamsReports_r18, cAParametersNRV1830IntraFreqL1MeasConfigR18SupportedMaxReportBeamsReportsR18Constraints); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.SupportedMaxAperiodic_LTM_CSI_ReportConfig_r18, per.Constrained(0, 4)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.SupportedMaxPeriodic_LTM_CSI_ReportConfig_r18, per.Constrained(1, 4)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.SupportedMaxSemiPersistent_LTM_CSI_ReportConfig_r18, per.Constrained(0, 4)); err != nil {
				return err
			}
		}
	}

	// 3. interFreqL1-MeasConfig-r18: sequence
	{
		if ie.InterFreqL1_MeasConfig_r18 != nil {
			c := ie.InterFreqL1_MeasConfig_r18
			if err := e.EncodeInteger(c.SupportedMaxIntraInterFreqCellsConfig_r18, per.Constrained(1, 8)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.SupportedMaxIntraInterFreqCellsPerReport_r18, per.Constrained(1, 4)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.SupportedMaxIntraInterFreqBeamsPerCellReports_r18, per.Constrained(1, 4)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.SupportedMaxIntraInterFreqBeamsReports_r18, cAParametersNRV1830InterFreqL1MeasConfigR18SupportedMaxIntraInterFreqBeamsReportsR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. currentSpCellInclL1-Report-r18: enumerated
	{
		if ie.CurrentSpCellInclL1_Report_r18 != nil {
			if err := e.EncodeEnumerated(*ie.CurrentSpCellInclL1_Report_r18, cAParametersNRV1830CurrentSpCellInclL1ReportR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. multiCellL1-measRTD-greaterThan-CP-r18: enumerated
	{
		if ie.MultiCellL1_MeasRTD_GreaterThan_CP_r18 != nil {
			if err := e.EncodeEnumerated(*ie.MultiCellL1_MeasRTD_GreaterThan_CP_r18, cAParametersNRV1830MultiCellL1MeasRTDGreaterThanCPR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. interFreqSSB-L1-MeasWithoutGaps-r18: enumerated
	{
		if ie.InterFreqSSB_L1_MeasWithoutGaps_r18 != nil {
			if err := e.EncodeEnumerated(*ie.InterFreqSSB_L1_MeasWithoutGaps_r18, cAParametersNRV1830InterFreqSSBL1MeasWithoutGapsR18Constraints); err != nil {
				return err
			}
		}
	}

	// 7. maxFreqLayersL1-Meas-r18: sequence
	{
		if ie.MaxFreqLayersL1_Meas_r18 != nil {
			c := ie.MaxFreqLayersL1_Meas_r18
			cAParametersNRV1830MaxFreqLayersL1MeasR18Seq := e.NewSequenceEncoder(cAParametersNRV1830MaxFreqLayersL1MeasR18Constraints)
			if err := cAParametersNRV1830MaxFreqLayersL1MeasR18Seq.EncodePreamble([]bool{c.SupportedMaxIntraInterFreqLayersWithoutGaps_r18 != nil, c.SupportedMaxInterFreqLayersWithGaps_r18 != nil}); err != nil {
				return err
			}
			if c.SupportedMaxIntraInterFreqLayersWithoutGaps_r18 != nil {
				if err := e.EncodeInteger((*c.SupportedMaxIntraInterFreqLayersWithoutGaps_r18), per.Constrained(1, 8)); err != nil {
					return err
				}
			}
			if c.SupportedMaxInterFreqLayersWithGaps_r18 != nil {
				if err := e.EncodeInteger((*c.SupportedMaxInterFreqLayersWithGaps_r18), per.Constrained(1, 8)); err != nil {
					return err
				}
			}
		}
	}

	// 8. maxNeighCellsPerFreqLayerL1-Meas-r18: sequence
	{
		if ie.MaxNeighCellsPerFreqLayerL1_Meas_r18 != nil {
			c := ie.MaxNeighCellsPerFreqLayerL1_Meas_r18
			cAParametersNRV1830MaxNeighCellsPerFreqLayerL1MeasR18Seq := e.NewSequenceEncoder(cAParametersNRV1830MaxNeighCellsPerFreqLayerL1MeasR18Constraints)
			if err := cAParametersNRV1830MaxNeighCellsPerFreqLayerL1MeasR18Seq.EncodePreamble([]bool{c.SupportedMaxNeighCellsPerFreqLayersWithoutGaps_r18 != nil, c.SupportedMaxNeighCellsPerFreqLayersWithGaps_r18 != nil}); err != nil {
				return err
			}
			if c.SupportedMaxNeighCellsPerFreqLayersWithoutGaps_r18 != nil {
				if err := e.EncodeInteger((*c.SupportedMaxNeighCellsPerFreqLayersWithoutGaps_r18), per.Constrained(1, 8)); err != nil {
					return err
				}
			}
			if c.SupportedMaxNeighCellsPerFreqLayersWithGaps_r18 != nil {
				if err := e.EncodeInteger((*c.SupportedMaxNeighCellsPerFreqLayersWithGaps_r18), per.Constrained(1, 8)); err != nil {
					return err
				}
			}
		}
	}

	// 9. supportedMaxCellsWithoutGapsL1-Meas-r18: integer
	{
		if ie.SupportedMaxCellsWithoutGapsL1_Meas_r18 != nil {
			if err := e.EncodeInteger(*ie.SupportedMaxCellsWithoutGapsL1_Meas_r18, cAParametersNRV1830SupportedMaxCellsWithoutGapsL1MeasR18Constraints); err != nil {
				return err
			}
		}
	}

	// 10. supportedMaxSSB-WithinSlotL1-Meas-r18: enumerated
	{
		if ie.SupportedMaxSSB_WithinSlotL1_Meas_r18 != nil {
			if err := e.EncodeEnumerated(*ie.SupportedMaxSSB_WithinSlotL1_Meas_r18, cAParametersNRV1830SupportedMaxSSBWithinSlotL1MeasR18Constraints); err != nil {
				return err
			}
		}
	}

	// 11. dummy: sequence
	{
		if ie.Dummy != nil {
			c := ie.Dummy
			cAParametersNRV1830DummySeq := e.NewSequenceEncoder(cAParametersNRV1830DummyConstraints)
			if err := cAParametersNRV1830DummySeq.EncodePreamble([]bool{c.SupportedMaxSSB_PerFreqLayerWithoutGaps_r18 != nil, c.SupportedMaxSSB_PerFreqLayerWithGaps_r18 != nil}); err != nil {
				return err
			}
			if c.SupportedMaxSSB_PerFreqLayerWithoutGaps_r18 != nil {
				if err := e.EncodeInteger((*c.SupportedMaxSSB_PerFreqLayerWithoutGaps_r18), per.Constrained(1, 8)); err != nil {
					return err
				}
			}
			if c.SupportedMaxSSB_PerFreqLayerWithGaps_r18 != nil {
				if err := e.EncodeInteger((*c.SupportedMaxSSB_PerFreqLayerWithGaps_r18), per.Constrained(1, 8)); err != nil {
					return err
				}
			}
		}
	}

	// 12. supportedMaxSSB-L1-Meas-r18: enumerated
	{
		if ie.SupportedMaxSSB_L1_Meas_r18 != nil {
			if err := e.EncodeEnumerated(*ie.SupportedMaxSSB_L1_Meas_r18, cAParametersNRV1830SupportedMaxSSBL1MeasR18Constraints); err != nil {
				return err
			}
		}
	}

	// 13. qcl-MultiCellDCI-1-3-r18: enumerated
	{
		if ie.Qcl_MultiCellDCI_1_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Qcl_MultiCellDCI_1_3_r18, cAParametersNRV1830QclMultiCellDCI13R18Constraints); err != nil {
				return err
			}
		}
	}

	// 14. bwp-SwitchingDCI-0-3-And-1-3-r18: enumerated
	{
		if ie.Bwp_SwitchingDCI_0_3_And_1_3_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Bwp_SwitchingDCI_0_3_And_1_3_r18, cAParametersNRV1830BwpSwitchingDCI03And13R18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNR_v1830) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRV1830Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. intraFreqL1-MeasConfig-r18: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.IntraFreqL1_MeasConfig_r18 = &struct {
				SupportedMaxIntraFreqCellsConfig_r18                int64
				SupportedMaxIntraFreqCellsPerReport_r18             int64
				SupportedMaxReportBeamsPerReportedCell_r18          int64
				SupportedMaxReportBeamsReports_r18                  int64
				SupportedMaxAperiodic_LTM_CSI_ReportConfig_r18      int64
				SupportedMaxPeriodic_LTM_CSI_ReportConfig_r18       int64
				SupportedMaxSemiPersistent_LTM_CSI_ReportConfig_r18 int64
			}{}
			c := ie.IntraFreqL1_MeasConfig_r18
			{
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.SupportedMaxIntraFreqCellsConfig_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.SupportedMaxIntraFreqCellsPerReport_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.SupportedMaxReportBeamsPerReportedCell_r18 = v
			}
			{
				v, err := d.DecodeEnumerated(cAParametersNRV1830IntraFreqL1MeasConfigR18SupportedMaxReportBeamsReportsR18Constraints)
				if err != nil {
					return err
				}
				c.SupportedMaxReportBeamsReports_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 4))
				if err != nil {
					return err
				}
				c.SupportedMaxAperiodic_LTM_CSI_ReportConfig_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.SupportedMaxPeriodic_LTM_CSI_ReportConfig_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(0, 4))
				if err != nil {
					return err
				}
				c.SupportedMaxSemiPersistent_LTM_CSI_ReportConfig_r18 = v
			}
		}
	}

	// 3. interFreqL1-MeasConfig-r18: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.InterFreqL1_MeasConfig_r18 = &struct {
				SupportedMaxIntraInterFreqCellsConfig_r18         int64
				SupportedMaxIntraInterFreqCellsPerReport_r18      int64
				SupportedMaxIntraInterFreqBeamsPerCellReports_r18 int64
				SupportedMaxIntraInterFreqBeamsReports_r18        int64
			}{}
			c := ie.InterFreqL1_MeasConfig_r18
			{
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.SupportedMaxIntraInterFreqCellsConfig_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.SupportedMaxIntraInterFreqCellsPerReport_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.SupportedMaxIntraInterFreqBeamsPerCellReports_r18 = v
			}
			{
				v, err := d.DecodeEnumerated(cAParametersNRV1830InterFreqL1MeasConfigR18SupportedMaxIntraInterFreqBeamsReportsR18Constraints)
				if err != nil {
					return err
				}
				c.SupportedMaxIntraInterFreqBeamsReports_r18 = v
			}
		}
	}

	// 4. currentSpCellInclL1-Report-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1830CurrentSpCellInclL1ReportR18Constraints)
			if err != nil {
				return err
			}
			ie.CurrentSpCellInclL1_Report_r18 = &idx
		}
	}

	// 5. multiCellL1-measRTD-greaterThan-CP-r18: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1830MultiCellL1MeasRTDGreaterThanCPR18Constraints)
			if err != nil {
				return err
			}
			ie.MultiCellL1_MeasRTD_GreaterThan_CP_r18 = &idx
		}
	}

	// 6. interFreqSSB-L1-MeasWithoutGaps-r18: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1830InterFreqSSBL1MeasWithoutGapsR18Constraints)
			if err != nil {
				return err
			}
			ie.InterFreqSSB_L1_MeasWithoutGaps_r18 = &idx
		}
	}

	// 7. maxFreqLayersL1-Meas-r18: sequence
	{
		if seq.IsComponentPresent(5) {
			ie.MaxFreqLayersL1_Meas_r18 = &struct {
				SupportedMaxIntraInterFreqLayersWithoutGaps_r18 *int64
				SupportedMaxInterFreqLayersWithGaps_r18         *int64
			}{}
			c := ie.MaxFreqLayersL1_Meas_r18
			cAParametersNRV1830MaxFreqLayersL1MeasR18Seq := d.NewSequenceDecoder(cAParametersNRV1830MaxFreqLayersL1MeasR18Constraints)
			if err := cAParametersNRV1830MaxFreqLayersL1MeasR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if cAParametersNRV1830MaxFreqLayersL1MeasR18Seq.IsComponentPresent(0) {
				c.SupportedMaxIntraInterFreqLayersWithoutGaps_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				(*c.SupportedMaxIntraInterFreqLayersWithoutGaps_r18) = v
			}
			if cAParametersNRV1830MaxFreqLayersL1MeasR18Seq.IsComponentPresent(1) {
				c.SupportedMaxInterFreqLayersWithGaps_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				(*c.SupportedMaxInterFreqLayersWithGaps_r18) = v
			}
		}
	}

	// 8. maxNeighCellsPerFreqLayerL1-Meas-r18: sequence
	{
		if seq.IsComponentPresent(6) {
			ie.MaxNeighCellsPerFreqLayerL1_Meas_r18 = &struct {
				SupportedMaxNeighCellsPerFreqLayersWithoutGaps_r18 *int64
				SupportedMaxNeighCellsPerFreqLayersWithGaps_r18    *int64
			}{}
			c := ie.MaxNeighCellsPerFreqLayerL1_Meas_r18
			cAParametersNRV1830MaxNeighCellsPerFreqLayerL1MeasR18Seq := d.NewSequenceDecoder(cAParametersNRV1830MaxNeighCellsPerFreqLayerL1MeasR18Constraints)
			if err := cAParametersNRV1830MaxNeighCellsPerFreqLayerL1MeasR18Seq.DecodePreamble(); err != nil {
				return err
			}
			if cAParametersNRV1830MaxNeighCellsPerFreqLayerL1MeasR18Seq.IsComponentPresent(0) {
				c.SupportedMaxNeighCellsPerFreqLayersWithoutGaps_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				(*c.SupportedMaxNeighCellsPerFreqLayersWithoutGaps_r18) = v
			}
			if cAParametersNRV1830MaxNeighCellsPerFreqLayerL1MeasR18Seq.IsComponentPresent(1) {
				c.SupportedMaxNeighCellsPerFreqLayersWithGaps_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				(*c.SupportedMaxNeighCellsPerFreqLayersWithGaps_r18) = v
			}
		}
	}

	// 9. supportedMaxCellsWithoutGapsL1-Meas-r18: integer
	{
		if seq.IsComponentPresent(7) {
			v, err := d.DecodeInteger(cAParametersNRV1830SupportedMaxCellsWithoutGapsL1MeasR18Constraints)
			if err != nil {
				return err
			}
			ie.SupportedMaxCellsWithoutGapsL1_Meas_r18 = &v
		}
	}

	// 10. supportedMaxSSB-WithinSlotL1-Meas-r18: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1830SupportedMaxSSBWithinSlotL1MeasR18Constraints)
			if err != nil {
				return err
			}
			ie.SupportedMaxSSB_WithinSlotL1_Meas_r18 = &idx
		}
	}

	// 11. dummy: sequence
	{
		if seq.IsComponentPresent(9) {
			ie.Dummy = &struct {
				SupportedMaxSSB_PerFreqLayerWithoutGaps_r18 *int64
				SupportedMaxSSB_PerFreqLayerWithGaps_r18    *int64
			}{}
			c := ie.Dummy
			cAParametersNRV1830DummySeq := d.NewSequenceDecoder(cAParametersNRV1830DummyConstraints)
			if err := cAParametersNRV1830DummySeq.DecodePreamble(); err != nil {
				return err
			}
			if cAParametersNRV1830DummySeq.IsComponentPresent(0) {
				c.SupportedMaxSSB_PerFreqLayerWithoutGaps_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				(*c.SupportedMaxSSB_PerFreqLayerWithoutGaps_r18) = v
			}
			if cAParametersNRV1830DummySeq.IsComponentPresent(1) {
				c.SupportedMaxSSB_PerFreqLayerWithGaps_r18 = new(int64)
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				(*c.SupportedMaxSSB_PerFreqLayerWithGaps_r18) = v
			}
		}
	}

	// 12. supportedMaxSSB-L1-Meas-r18: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1830SupportedMaxSSBL1MeasR18Constraints)
			if err != nil {
				return err
			}
			ie.SupportedMaxSSB_L1_Meas_r18 = &idx
		}
	}

	// 13. qcl-MultiCellDCI-1-3-r18: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1830QclMultiCellDCI13R18Constraints)
			if err != nil {
				return err
			}
			ie.Qcl_MultiCellDCI_1_3_r18 = &idx
		}
	}

	// 14. bwp-SwitchingDCI-0-3-And-1-3-r18: enumerated
	{
		if seq.IsComponentPresent(12) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1830BwpSwitchingDCI03And13R18Constraints)
			if err != nil {
				return err
			}
			ie.Bwp_SwitchingDCI_0_3_And_1_3_r18 = &idx
		}
	}

	return nil
}
