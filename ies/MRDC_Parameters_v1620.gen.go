// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MRDC-Parameters-v1620 (line 22577).

var mRDCParametersV1620Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxUplinkDutyCycle-interBandENDC-TDD-PC2-r16", Optional: true},
		{Name: "tdm-restrictionTDD-endc-r16", Optional: true},
		{Name: "tdm-restrictionFDD-endc-r16", Optional: true},
		{Name: "singleUL-HARQ-offsetTDD-PCell-r16", Optional: true},
		{Name: "tdm-restrictionDualTX-FDD-endc-r16", Optional: true},
	},
}

const (
	MRDC_Parameters_v1620_Tdm_RestrictionTDD_Endc_r16_Supported = 0
)

var mRDCParametersV1620TdmRestrictionTDDEndcR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1620_Tdm_RestrictionTDD_Endc_r16_Supported},
}

const (
	MRDC_Parameters_v1620_Tdm_RestrictionFDD_Endc_r16_Supported = 0
)

var mRDCParametersV1620TdmRestrictionFDDEndcR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1620_Tdm_RestrictionFDD_Endc_r16_Supported},
}

const (
	MRDC_Parameters_v1620_SingleUL_HARQ_OffsetTDD_PCell_r16_Supported = 0
)

var mRDCParametersV1620SingleULHARQOffsetTDDPCellR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1620_SingleUL_HARQ_OffsetTDD_PCell_r16_Supported},
}

const (
	MRDC_Parameters_v1620_Tdm_RestrictionDualTX_FDD_Endc_r16_Supported = 0
)

var mRDCParametersV1620TdmRestrictionDualTXFDDEndcR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1620_Tdm_RestrictionDualTX_FDD_Endc_r16_Supported},
}

var mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "eutra-TDD-Config0-r16", Optional: true},
		{Name: "eutra-TDD-Config1-r16", Optional: true},
		{Name: "eutra-TDD-Config2-r16", Optional: true},
		{Name: "eutra-TDD-Config3-r16", Optional: true},
		{Name: "eutra-TDD-Config4-r16", Optional: true},
		{Name: "eutra-TDD-Config5-r16", Optional: true},
		{Name: "eutra-TDD-Config6-r16", Optional: true},
	},
}

const (
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config0_r16_N20  = 0
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config0_r16_N40  = 1
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config0_r16_N50  = 2
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config0_r16_N60  = 3
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config0_r16_N70  = 4
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config0_r16_N80  = 5
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config0_r16_N90  = 6
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config0_r16_N100 = 7
)

var mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig0R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config0_r16_N20, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config0_r16_N40, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config0_r16_N50, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config0_r16_N60, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config0_r16_N70, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config0_r16_N80, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config0_r16_N90, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config0_r16_N100},
}

const (
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config1_r16_N20  = 0
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config1_r16_N40  = 1
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config1_r16_N50  = 2
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config1_r16_N60  = 3
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config1_r16_N70  = 4
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config1_r16_N80  = 5
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config1_r16_N90  = 6
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config1_r16_N100 = 7
)

var mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig1R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config1_r16_N20, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config1_r16_N40, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config1_r16_N50, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config1_r16_N60, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config1_r16_N70, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config1_r16_N80, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config1_r16_N90, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config1_r16_N100},
}

const (
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config2_r16_N20  = 0
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config2_r16_N40  = 1
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config2_r16_N50  = 2
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config2_r16_N60  = 3
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config2_r16_N70  = 4
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config2_r16_N80  = 5
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config2_r16_N90  = 6
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config2_r16_N100 = 7
)

var mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig2R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config2_r16_N20, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config2_r16_N40, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config2_r16_N50, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config2_r16_N60, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config2_r16_N70, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config2_r16_N80, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config2_r16_N90, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config2_r16_N100},
}

const (
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config3_r16_N20  = 0
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config3_r16_N40  = 1
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config3_r16_N50  = 2
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config3_r16_N60  = 3
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config3_r16_N70  = 4
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config3_r16_N80  = 5
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config3_r16_N90  = 6
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config3_r16_N100 = 7
)

var mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig3R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config3_r16_N20, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config3_r16_N40, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config3_r16_N50, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config3_r16_N60, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config3_r16_N70, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config3_r16_N80, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config3_r16_N90, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config3_r16_N100},
}

const (
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config4_r16_N20  = 0
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config4_r16_N40  = 1
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config4_r16_N50  = 2
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config4_r16_N60  = 3
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config4_r16_N70  = 4
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config4_r16_N80  = 5
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config4_r16_N90  = 6
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config4_r16_N100 = 7
)

var mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig4R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config4_r16_N20, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config4_r16_N40, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config4_r16_N50, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config4_r16_N60, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config4_r16_N70, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config4_r16_N80, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config4_r16_N90, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config4_r16_N100},
}

const (
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config5_r16_N20  = 0
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config5_r16_N40  = 1
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config5_r16_N50  = 2
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config5_r16_N60  = 3
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config5_r16_N70  = 4
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config5_r16_N80  = 5
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config5_r16_N90  = 6
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config5_r16_N100 = 7
)

var mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig5R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config5_r16_N20, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config5_r16_N40, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config5_r16_N50, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config5_r16_N60, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config5_r16_N70, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config5_r16_N80, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config5_r16_N90, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config5_r16_N100},
}

const (
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config6_r16_N20  = 0
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config6_r16_N40  = 1
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config6_r16_N50  = 2
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config6_r16_N60  = 3
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config6_r16_N70  = 4
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config6_r16_N80  = 5
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config6_r16_N90  = 6
	MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config6_r16_N100 = 7
)

var mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig6R16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config6_r16_N20, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config6_r16_N40, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config6_r16_N50, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config6_r16_N60, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config6_r16_N70, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config6_r16_N80, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config6_r16_N90, MRDC_Parameters_v1620_MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16_Eutra_TDD_Config6_r16_N100},
}

type MRDC_Parameters_v1620 struct {
	MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16 *struct {
		Eutra_TDD_Config0_r16 *int64
		Eutra_TDD_Config1_r16 *int64
		Eutra_TDD_Config2_r16 *int64
		Eutra_TDD_Config3_r16 *int64
		Eutra_TDD_Config4_r16 *int64
		Eutra_TDD_Config5_r16 *int64
		Eutra_TDD_Config6_r16 *int64
	}
	Tdm_RestrictionTDD_Endc_r16        *int64
	Tdm_RestrictionFDD_Endc_r16        *int64
	SingleUL_HARQ_OffsetTDD_PCell_r16  *int64
	Tdm_RestrictionDualTX_FDD_Endc_r16 *int64
}

func (ie *MRDC_Parameters_v1620) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mRDCParametersV1620Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16 != nil, ie.Tdm_RestrictionTDD_Endc_r16 != nil, ie.Tdm_RestrictionFDD_Endc_r16 != nil, ie.SingleUL_HARQ_OffsetTDD_PCell_r16 != nil, ie.Tdm_RestrictionDualTX_FDD_Endc_r16 != nil}); err != nil {
		return err
	}

	// 2. maxUplinkDutyCycle-interBandENDC-TDD-PC2-r16: sequence
	{
		if ie.MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16 != nil {
			c := ie.MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16
			mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16Seq := e.NewSequenceEncoder(mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16Constraints)
			if err := mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16Seq.EncodePreamble([]bool{c.Eutra_TDD_Config0_r16 != nil, c.Eutra_TDD_Config1_r16 != nil, c.Eutra_TDD_Config2_r16 != nil, c.Eutra_TDD_Config3_r16 != nil, c.Eutra_TDD_Config4_r16 != nil, c.Eutra_TDD_Config5_r16 != nil, c.Eutra_TDD_Config6_r16 != nil}); err != nil {
				return err
			}
			if c.Eutra_TDD_Config0_r16 != nil {
				if err := e.EncodeEnumerated((*c.Eutra_TDD_Config0_r16), mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig0R16Constraints); err != nil {
					return err
				}
			}
			if c.Eutra_TDD_Config1_r16 != nil {
				if err := e.EncodeEnumerated((*c.Eutra_TDD_Config1_r16), mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig1R16Constraints); err != nil {
					return err
				}
			}
			if c.Eutra_TDD_Config2_r16 != nil {
				if err := e.EncodeEnumerated((*c.Eutra_TDD_Config2_r16), mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig2R16Constraints); err != nil {
					return err
				}
			}
			if c.Eutra_TDD_Config3_r16 != nil {
				if err := e.EncodeEnumerated((*c.Eutra_TDD_Config3_r16), mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig3R16Constraints); err != nil {
					return err
				}
			}
			if c.Eutra_TDD_Config4_r16 != nil {
				if err := e.EncodeEnumerated((*c.Eutra_TDD_Config4_r16), mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig4R16Constraints); err != nil {
					return err
				}
			}
			if c.Eutra_TDD_Config5_r16 != nil {
				if err := e.EncodeEnumerated((*c.Eutra_TDD_Config5_r16), mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig5R16Constraints); err != nil {
					return err
				}
			}
			if c.Eutra_TDD_Config6_r16 != nil {
				if err := e.EncodeEnumerated((*c.Eutra_TDD_Config6_r16), mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig6R16Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 3. tdm-restrictionTDD-endc-r16: enumerated
	{
		if ie.Tdm_RestrictionTDD_Endc_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Tdm_RestrictionTDD_Endc_r16, mRDCParametersV1620TdmRestrictionTDDEndcR16Constraints); err != nil {
				return err
			}
		}
	}

	// 4. tdm-restrictionFDD-endc-r16: enumerated
	{
		if ie.Tdm_RestrictionFDD_Endc_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Tdm_RestrictionFDD_Endc_r16, mRDCParametersV1620TdmRestrictionFDDEndcR16Constraints); err != nil {
				return err
			}
		}
	}

	// 5. singleUL-HARQ-offsetTDD-PCell-r16: enumerated
	{
		if ie.SingleUL_HARQ_OffsetTDD_PCell_r16 != nil {
			if err := e.EncodeEnumerated(*ie.SingleUL_HARQ_OffsetTDD_PCell_r16, mRDCParametersV1620SingleULHARQOffsetTDDPCellR16Constraints); err != nil {
				return err
			}
		}
	}

	// 6. tdm-restrictionDualTX-FDD-endc-r16: enumerated
	{
		if ie.Tdm_RestrictionDualTX_FDD_Endc_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Tdm_RestrictionDualTX_FDD_Endc_r16, mRDCParametersV1620TdmRestrictionDualTXFDDEndcR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MRDC_Parameters_v1620) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mRDCParametersV1620Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. maxUplinkDutyCycle-interBandENDC-TDD-PC2-r16: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16 = &struct {
				Eutra_TDD_Config0_r16 *int64
				Eutra_TDD_Config1_r16 *int64
				Eutra_TDD_Config2_r16 *int64
				Eutra_TDD_Config3_r16 *int64
				Eutra_TDD_Config4_r16 *int64
				Eutra_TDD_Config5_r16 *int64
				Eutra_TDD_Config6_r16 *int64
			}{}
			c := ie.MaxUplinkDutyCycle_InterBandENDC_TDD_PC2_r16
			mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16Seq := d.NewSequenceDecoder(mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16Constraints)
			if err := mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16Seq.DecodePreamble(); err != nil {
				return err
			}
			if mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16Seq.IsComponentPresent(0) {
				c.Eutra_TDD_Config0_r16 = new(int64)
				v, err := d.DecodeEnumerated(mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig0R16Constraints)
				if err != nil {
					return err
				}
				(*c.Eutra_TDD_Config0_r16) = v
			}
			if mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16Seq.IsComponentPresent(1) {
				c.Eutra_TDD_Config1_r16 = new(int64)
				v, err := d.DecodeEnumerated(mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig1R16Constraints)
				if err != nil {
					return err
				}
				(*c.Eutra_TDD_Config1_r16) = v
			}
			if mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16Seq.IsComponentPresent(2) {
				c.Eutra_TDD_Config2_r16 = new(int64)
				v, err := d.DecodeEnumerated(mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig2R16Constraints)
				if err != nil {
					return err
				}
				(*c.Eutra_TDD_Config2_r16) = v
			}
			if mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16Seq.IsComponentPresent(3) {
				c.Eutra_TDD_Config3_r16 = new(int64)
				v, err := d.DecodeEnumerated(mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig3R16Constraints)
				if err != nil {
					return err
				}
				(*c.Eutra_TDD_Config3_r16) = v
			}
			if mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16Seq.IsComponentPresent(4) {
				c.Eutra_TDD_Config4_r16 = new(int64)
				v, err := d.DecodeEnumerated(mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig4R16Constraints)
				if err != nil {
					return err
				}
				(*c.Eutra_TDD_Config4_r16) = v
			}
			if mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16Seq.IsComponentPresent(5) {
				c.Eutra_TDD_Config5_r16 = new(int64)
				v, err := d.DecodeEnumerated(mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig5R16Constraints)
				if err != nil {
					return err
				}
				(*c.Eutra_TDD_Config5_r16) = v
			}
			if mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16Seq.IsComponentPresent(6) {
				c.Eutra_TDD_Config6_r16 = new(int64)
				v, err := d.DecodeEnumerated(mRDCParametersV1620MaxUplinkDutyCycleInterBandENDCTDDPC2R16EutraTDDConfig6R16Constraints)
				if err != nil {
					return err
				}
				(*c.Eutra_TDD_Config6_r16) = v
			}
		}
	}

	// 3. tdm-restrictionTDD-endc-r16: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(mRDCParametersV1620TdmRestrictionTDDEndcR16Constraints)
			if err != nil {
				return err
			}
			ie.Tdm_RestrictionTDD_Endc_r16 = &idx
		}
	}

	// 4. tdm-restrictionFDD-endc-r16: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(mRDCParametersV1620TdmRestrictionFDDEndcR16Constraints)
			if err != nil {
				return err
			}
			ie.Tdm_RestrictionFDD_Endc_r16 = &idx
		}
	}

	// 5. singleUL-HARQ-offsetTDD-PCell-r16: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(mRDCParametersV1620SingleULHARQOffsetTDDPCellR16Constraints)
			if err != nil {
				return err
			}
			ie.SingleUL_HARQ_OffsetTDD_PCell_r16 = &idx
		}
	}

	// 6. tdm-restrictionDualTX-FDD-endc-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(mRDCParametersV1620TdmRestrictionDualTXFDDEndcR16Constraints)
			if err != nil {
				return err
			}
			ie.Tdm_RestrictionDualTX_FDD_Endc_r16 = &idx
		}
	}

	return nil
}
