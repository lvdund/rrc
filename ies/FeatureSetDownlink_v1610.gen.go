// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetDownlink-v1610 (line 19509).

var featureSetDownlinkV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "cbgPDSCH-ProcessingType1-DifferentTB-PerSlot-r16", Optional: true},
		{Name: "cbgPDSCH-ProcessingType2-DifferentTB-PerSlot-r16", Optional: true},
		{Name: "intraFreqDAPS-r16", Optional: true},
		{Name: "intraBandFreqSeparationDL-v1620", Optional: true},
		{Name: "intraBandFreqSeparationDL-Only-r16", Optional: true},
		{Name: "pdcch-Monitoring-r16", Optional: true},
		{Name: "pdcch-MonitoringMixed-r16", Optional: true},
		{Name: "crossCarrierSchedulingProcessing-DiffSCS-r16", Optional: true},
		{Name: "singleDCI-SDM-scheme-r16", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1610_Pdcch_MonitoringMixed_r16_Supported = 0
)

var featureSetDownlinkV1610PdcchMonitoringMixedR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1610_Pdcch_MonitoringMixed_r16_Supported},
}

const (
	FeatureSetDownlink_v1610_SingleDCI_SDM_Scheme_r16_Supported = 0
)

var featureSetDownlinkV1610SingleDCISDMSchemeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1610_SingleDCI_SDM_Scheme_r16_Supported},
}

var featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r16", Optional: true},
		{Name: "scs-30kHz-r16", Optional: true},
		{Name: "scs-60kHz-r16", Optional: true},
		{Name: "scs-120kHz-r16", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_15kHz_r16_One   = 0
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto2 = 1
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto4 = 2
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto7 = 3
)

var featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Scs15kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_15kHz_r16_One, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto2, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto4, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto7},
}

const (
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_30kHz_r16_One   = 0
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto2 = 1
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto4 = 2
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto7 = 3
)

var featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Scs30kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_30kHz_r16_One, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto2, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto4, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto7},
}

const (
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_60kHz_r16_One   = 0
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto2 = 1
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto4 = 2
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto7 = 3
)

var featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Scs60kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_60kHz_r16_One, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto2, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto4, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto7},
}

const (
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_120kHz_r16_One   = 0
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto2 = 1
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto4 = 2
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto7 = 3
)

var featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Scs120kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_120kHz_r16_One, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto2, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto4, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto7},
}

var featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r16", Optional: true},
		{Name: "scs-30kHz-r16", Optional: true},
		{Name: "scs-60kHz-r16", Optional: true},
		{Name: "scs-120kHz-r16", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_15kHz_r16_One   = 0
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto2 = 1
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto4 = 2
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto7 = 3
)

var featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Scs15kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_15kHz_r16_One, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto2, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto4, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_15kHz_r16_Upto7},
}

const (
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_30kHz_r16_One   = 0
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto2 = 1
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto4 = 2
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto7 = 3
)

var featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Scs30kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_30kHz_r16_One, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto2, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto4, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_30kHz_r16_Upto7},
}

const (
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_60kHz_r16_One   = 0
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto2 = 1
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto4 = 2
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto7 = 3
)

var featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Scs60kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_60kHz_r16_One, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto2, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto4, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_60kHz_r16_Upto7},
}

const (
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_120kHz_r16_One   = 0
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto2 = 1
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto4 = 2
	FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto7 = 3
)

var featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Scs120kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_120kHz_r16_One, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto2, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto4, FeatureSetDownlink_v1610_CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16_Scs_120kHz_r16_Upto7},
}

var featureSetDownlinkV1610IntraFreqDAPSR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "intraFreqDiffSCS-DAPS-r16", Optional: true},
		{Name: "intraFreqAsyncDAPS-r16", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1610_IntraFreqDAPS_r16_IntraFreqDiffSCS_DAPS_r16_Supported = 0
)

var featureSetDownlinkV1610IntraFreqDAPSR16IntraFreqDiffSCSDAPSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1610_IntraFreqDAPS_r16_IntraFreqDiffSCS_DAPS_r16_Supported},
}

const (
	FeatureSetDownlink_v1610_IntraFreqDAPS_r16_IntraFreqAsyncDAPS_r16_Supported = 0
)

var featureSetDownlinkV1610IntraFreqDAPSR16IntraFreqAsyncDAPSR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1610_IntraFreqDAPS_r16_IntraFreqAsyncDAPS_r16_Supported},
}

var featureSetDownlinkV1610PdcchMonitoringR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdsch-ProcessingType1-r16", Optional: true},
		{Name: "pdsch-ProcessingType2-r16", Optional: true},
	},
}

var featureSetDownlinkV1610PdcchMonitoringR16PdschProcessingType1R16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r16", Optional: true},
		{Name: "scs-30kHz-r16", Optional: true},
	},
}

var featureSetDownlinkV1610PdcchMonitoringR16PdschProcessingType2R16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-r16", Optional: true},
		{Name: "scs-30kHz-r16", Optional: true},
	},
}

var featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-15kHz-120kHz-r16", Optional: true},
		{Name: "scs-15kHz-60kHz-r16", Optional: true},
		{Name: "scs-30kHz-120kHz-r16", Optional: true},
		{Name: "scs-15kHz-30kHz-r16", Optional: true},
		{Name: "scs-30kHz-60kHz-r16", Optional: true},
		{Name: "scs-60kHz-120kHz-r16", Optional: true},
	},
}

const (
	FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_120kHz_r16_N1 = 0
	FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_120kHz_r16_N2 = 1
	FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_120kHz_r16_N4 = 2
)

var featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs15kHz120kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_120kHz_r16_N1, FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_120kHz_r16_N2, FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_120kHz_r16_N4},
}

const (
	FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_60kHz_r16_N1 = 0
	FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_60kHz_r16_N2 = 1
	FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_60kHz_r16_N4 = 2
)

var featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs15kHz60kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_60kHz_r16_N1, FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_60kHz_r16_N2, FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_60kHz_r16_N4},
}

const (
	FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_30kHz_120kHz_r16_N1 = 0
	FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_30kHz_120kHz_r16_N2 = 1
	FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_30kHz_120kHz_r16_N4 = 2
)

var featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs30kHz120kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_30kHz_120kHz_r16_N1, FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_30kHz_120kHz_r16_N2, FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_30kHz_120kHz_r16_N4},
}

const (
	FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_30kHz_r16_N2 = 0
)

var featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs15kHz30kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_15kHz_30kHz_r16_N2},
}

const (
	FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_30kHz_60kHz_r16_N2 = 0
)

var featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs30kHz60kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_30kHz_60kHz_r16_N2},
}

const (
	FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_60kHz_120kHz_r16_N2 = 0
)

var featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs60kHz120kHzR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetDownlink_v1610_CrossCarrierSchedulingProcessing_DiffSCS_r16_Scs_60kHz_120kHz_r16_N2},
}

type FeatureSetDownlink_v1610 struct {
	CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16 *struct {
		Scs_15kHz_r16  *int64
		Scs_30kHz_r16  *int64
		Scs_60kHz_r16  *int64
		Scs_120kHz_r16 *int64
	}
	CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16 *struct {
		Scs_15kHz_r16  *int64
		Scs_30kHz_r16  *int64
		Scs_60kHz_r16  *int64
		Scs_120kHz_r16 *int64
	}
	IntraFreqDAPS_r16 *struct {
		IntraFreqDiffSCS_DAPS_r16 *int64
		IntraFreqAsyncDAPS_r16    *int64
	}
	IntraBandFreqSeparationDL_v1620    *FreqSeparationClassDL_v1620
	IntraBandFreqSeparationDL_Only_r16 *FreqSeparationClassDL_Only_r16
	Pdcch_Monitoring_r16               *struct {
		Pdsch_ProcessingType1_r16 *struct {
			Scs_15kHz_r16 *PDCCH_MonitoringOccasions_r16
			Scs_30kHz_r16 *PDCCH_MonitoringOccasions_r16
		}
		Pdsch_ProcessingType2_r16 *struct {
			Scs_15kHz_r16 *PDCCH_MonitoringOccasions_r16
			Scs_30kHz_r16 *PDCCH_MonitoringOccasions_r16
		}
	}
	Pdcch_MonitoringMixed_r16                    *int64
	CrossCarrierSchedulingProcessing_DiffSCS_r16 *struct {
		Scs_15kHz_120kHz_r16 *int64
		Scs_15kHz_60kHz_r16  *int64
		Scs_30kHz_120kHz_r16 *int64
		Scs_15kHz_30kHz_r16  *int64
		Scs_30kHz_60kHz_r16  *int64
		Scs_60kHz_120kHz_r16 *int64
	}
	SingleDCI_SDM_Scheme_r16 *int64
}

func (ie *FeatureSetDownlink_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetDownlinkV1610Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16 != nil, ie.CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16 != nil, ie.IntraFreqDAPS_r16 != nil, ie.IntraBandFreqSeparationDL_v1620 != nil, ie.IntraBandFreqSeparationDL_Only_r16 != nil, ie.Pdcch_Monitoring_r16 != nil, ie.Pdcch_MonitoringMixed_r16 != nil, ie.CrossCarrierSchedulingProcessing_DiffSCS_r16 != nil, ie.SingleDCI_SDM_Scheme_r16 != nil}); err != nil {
		return err
	}

	// 2. cbgPDSCH-ProcessingType1-DifferentTB-PerSlot-r16: sequence
	{
		if ie.CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16 != nil {
			c := ie.CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16
			featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Seq := e.NewSequenceEncoder(featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Constraints)
			if err := featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Seq.EncodePreamble([]bool{c.Scs_15kHz_r16 != nil, c.Scs_30kHz_r16 != nil, c.Scs_60kHz_r16 != nil, c.Scs_120kHz_r16 != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_15kHz_r16), featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Scs15kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_30kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_30kHz_r16), featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Scs30kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_60kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_60kHz_r16), featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Scs60kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_120kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_120kHz_r16), featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Scs120kHzR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 3. cbgPDSCH-ProcessingType2-DifferentTB-PerSlot-r16: sequence
	{
		if ie.CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16 != nil {
			c := ie.CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16
			featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Seq := e.NewSequenceEncoder(featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Constraints)
			if err := featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Seq.EncodePreamble([]bool{c.Scs_15kHz_r16 != nil, c.Scs_30kHz_r16 != nil, c.Scs_60kHz_r16 != nil, c.Scs_120kHz_r16 != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_15kHz_r16), featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Scs15kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_30kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_30kHz_r16), featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Scs30kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_60kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_60kHz_r16), featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Scs60kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_120kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_120kHz_r16), featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Scs120kHzR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 4. intraFreqDAPS-r16: sequence
	{
		if ie.IntraFreqDAPS_r16 != nil {
			c := ie.IntraFreqDAPS_r16
			featureSetDownlinkV1610IntraFreqDAPSR16Seq := e.NewSequenceEncoder(featureSetDownlinkV1610IntraFreqDAPSR16Constraints)
			if err := featureSetDownlinkV1610IntraFreqDAPSR16Seq.EncodePreamble([]bool{c.IntraFreqDiffSCS_DAPS_r16 != nil, c.IntraFreqAsyncDAPS_r16 != nil}); err != nil {
				return err
			}
			if c.IntraFreqDiffSCS_DAPS_r16 != nil {
				if err := e.EncodeEnumerated((*c.IntraFreqDiffSCS_DAPS_r16), featureSetDownlinkV1610IntraFreqDAPSR16IntraFreqDiffSCSDAPSR16Constraints); err != nil {
					return err
				}
			}
			if c.IntraFreqAsyncDAPS_r16 != nil {
				if err := e.EncodeEnumerated((*c.IntraFreqAsyncDAPS_r16), featureSetDownlinkV1610IntraFreqDAPSR16IntraFreqAsyncDAPSR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 5. intraBandFreqSeparationDL-v1620: ref
	{
		if ie.IntraBandFreqSeparationDL_v1620 != nil {
			if err := ie.IntraBandFreqSeparationDL_v1620.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. intraBandFreqSeparationDL-Only-r16: ref
	{
		if ie.IntraBandFreqSeparationDL_Only_r16 != nil {
			if err := ie.IntraBandFreqSeparationDL_Only_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. pdcch-Monitoring-r16: sequence
	{
		if ie.Pdcch_Monitoring_r16 != nil {
			c := ie.Pdcch_Monitoring_r16
			featureSetDownlinkV1610PdcchMonitoringR16Seq := e.NewSequenceEncoder(featureSetDownlinkV1610PdcchMonitoringR16Constraints)
			if err := featureSetDownlinkV1610PdcchMonitoringR16Seq.EncodePreamble([]bool{c.Pdsch_ProcessingType1_r16 != nil, c.Pdsch_ProcessingType2_r16 != nil}); err != nil {
				return err
			}
			if c.Pdsch_ProcessingType1_r16 != nil {
				c := (*c.Pdsch_ProcessingType1_r16)
				featureSetDownlinkV1610PdcchMonitoringR16PdschProcessingType1R16Seq := e.NewSequenceEncoder(featureSetDownlinkV1610PdcchMonitoringR16PdschProcessingType1R16Constraints)
				if err := featureSetDownlinkV1610PdcchMonitoringR16PdschProcessingType1R16Seq.EncodePreamble([]bool{c.Scs_15kHz_r16 != nil, c.Scs_30kHz_r16 != nil}); err != nil {
					return err
				}
				if c.Scs_15kHz_r16 != nil {
					if err := c.Scs_15kHz_r16.Encode(e); err != nil {
						return err
					}
				}
				if c.Scs_30kHz_r16 != nil {
					if err := c.Scs_30kHz_r16.Encode(e); err != nil {
						return err
					}
				}
			}
			if c.Pdsch_ProcessingType2_r16 != nil {
				c := (*c.Pdsch_ProcessingType2_r16)
				featureSetDownlinkV1610PdcchMonitoringR16PdschProcessingType2R16Seq := e.NewSequenceEncoder(featureSetDownlinkV1610PdcchMonitoringR16PdschProcessingType2R16Constraints)
				if err := featureSetDownlinkV1610PdcchMonitoringR16PdschProcessingType2R16Seq.EncodePreamble([]bool{c.Scs_15kHz_r16 != nil, c.Scs_30kHz_r16 != nil}); err != nil {
					return err
				}
				if c.Scs_15kHz_r16 != nil {
					if err := c.Scs_15kHz_r16.Encode(e); err != nil {
						return err
					}
				}
				if c.Scs_30kHz_r16 != nil {
					if err := c.Scs_30kHz_r16.Encode(e); err != nil {
						return err
					}
				}
			}
		}
	}

	// 8. pdcch-MonitoringMixed-r16: enumerated
	{
		if ie.Pdcch_MonitoringMixed_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Pdcch_MonitoringMixed_r16, featureSetDownlinkV1610PdcchMonitoringMixedR16Constraints); err != nil {
				return err
			}
		}
	}

	// 9. crossCarrierSchedulingProcessing-DiffSCS-r16: sequence
	{
		if ie.CrossCarrierSchedulingProcessing_DiffSCS_r16 != nil {
			c := ie.CrossCarrierSchedulingProcessing_DiffSCS_r16
			featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Seq := e.NewSequenceEncoder(featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Constraints)
			if err := featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Seq.EncodePreamble([]bool{c.Scs_15kHz_120kHz_r16 != nil, c.Scs_15kHz_60kHz_r16 != nil, c.Scs_30kHz_120kHz_r16 != nil, c.Scs_15kHz_30kHz_r16 != nil, c.Scs_30kHz_60kHz_r16 != nil, c.Scs_60kHz_120kHz_r16 != nil}); err != nil {
				return err
			}
			if c.Scs_15kHz_120kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_15kHz_120kHz_r16), featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs15kHz120kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_15kHz_60kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_15kHz_60kHz_r16), featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs15kHz60kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_30kHz_120kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_30kHz_120kHz_r16), featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs30kHz120kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_15kHz_30kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_15kHz_30kHz_r16), featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs15kHz30kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_30kHz_60kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_30kHz_60kHz_r16), featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs30kHz60kHzR16Constraints); err != nil {
					return err
				}
			}
			if c.Scs_60kHz_120kHz_r16 != nil {
				if err := e.EncodeEnumerated((*c.Scs_60kHz_120kHz_r16), featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs60kHz120kHzR16Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 10. singleDCI-SDM-scheme-r16: enumerated
	{
		if ie.SingleDCI_SDM_Scheme_r16 != nil {
			if err := e.EncodeEnumerated(*ie.SingleDCI_SDM_Scheme_r16, featureSetDownlinkV1610SingleDCISDMSchemeR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetDownlink_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetDownlinkV1610Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. cbgPDSCH-ProcessingType1-DifferentTB-PerSlot-r16: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16 = &struct {
				Scs_15kHz_r16  *int64
				Scs_30kHz_r16  *int64
				Scs_60kHz_r16  *int64
				Scs_120kHz_r16 *int64
			}{}
			c := ie.CbgPDSCH_ProcessingType1_DifferentTB_PerSlot_r16
			featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Seq := d.NewSequenceDecoder(featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Constraints)
			if err := featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Seq.IsComponentPresent(0) {
				c.Scs_15kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Scs15kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_15kHz_r16) = v
			}
			if featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Seq.IsComponentPresent(1) {
				c.Scs_30kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Scs30kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_30kHz_r16) = v
			}
			if featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Seq.IsComponentPresent(2) {
				c.Scs_60kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Scs60kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_60kHz_r16) = v
			}
			if featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Seq.IsComponentPresent(3) {
				c.Scs_120kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1610CbgPDSCHProcessingType1DifferentTBPerSlotR16Scs120kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_120kHz_r16) = v
			}
		}
	}

	// 3. cbgPDSCH-ProcessingType2-DifferentTB-PerSlot-r16: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16 = &struct {
				Scs_15kHz_r16  *int64
				Scs_30kHz_r16  *int64
				Scs_60kHz_r16  *int64
				Scs_120kHz_r16 *int64
			}{}
			c := ie.CbgPDSCH_ProcessingType2_DifferentTB_PerSlot_r16
			featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Seq := d.NewSequenceDecoder(featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Constraints)
			if err := featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Seq.IsComponentPresent(0) {
				c.Scs_15kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Scs15kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_15kHz_r16) = v
			}
			if featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Seq.IsComponentPresent(1) {
				c.Scs_30kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Scs30kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_30kHz_r16) = v
			}
			if featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Seq.IsComponentPresent(2) {
				c.Scs_60kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Scs60kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_60kHz_r16) = v
			}
			if featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Seq.IsComponentPresent(3) {
				c.Scs_120kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1610CbgPDSCHProcessingType2DifferentTBPerSlotR16Scs120kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_120kHz_r16) = v
			}
		}
	}

	// 4. intraFreqDAPS-r16: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.IntraFreqDAPS_r16 = &struct {
				IntraFreqDiffSCS_DAPS_r16 *int64
				IntraFreqAsyncDAPS_r16    *int64
			}{}
			c := ie.IntraFreqDAPS_r16
			featureSetDownlinkV1610IntraFreqDAPSR16Seq := d.NewSequenceDecoder(featureSetDownlinkV1610IntraFreqDAPSR16Constraints)
			if err := featureSetDownlinkV1610IntraFreqDAPSR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetDownlinkV1610IntraFreqDAPSR16Seq.IsComponentPresent(0) {
				c.IntraFreqDiffSCS_DAPS_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1610IntraFreqDAPSR16IntraFreqDiffSCSDAPSR16Constraints)
				if err != nil {
					return err
				}
				(*c.IntraFreqDiffSCS_DAPS_r16) = v
			}
			if featureSetDownlinkV1610IntraFreqDAPSR16Seq.IsComponentPresent(1) {
				c.IntraFreqAsyncDAPS_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1610IntraFreqDAPSR16IntraFreqAsyncDAPSR16Constraints)
				if err != nil {
					return err
				}
				(*c.IntraFreqAsyncDAPS_r16) = v
			}
		}
	}

	// 5. intraBandFreqSeparationDL-v1620: ref
	{
		if seq.IsComponentPresent(3) {
			ie.IntraBandFreqSeparationDL_v1620 = new(FreqSeparationClassDL_v1620)
			if err := ie.IntraBandFreqSeparationDL_v1620.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. intraBandFreqSeparationDL-Only-r16: ref
	{
		if seq.IsComponentPresent(4) {
			ie.IntraBandFreqSeparationDL_Only_r16 = new(FreqSeparationClassDL_Only_r16)
			if err := ie.IntraBandFreqSeparationDL_Only_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. pdcch-Monitoring-r16: sequence
	{
		if seq.IsComponentPresent(5) {
			ie.Pdcch_Monitoring_r16 = &struct {
				Pdsch_ProcessingType1_r16 *struct {
					Scs_15kHz_r16 *PDCCH_MonitoringOccasions_r16
					Scs_30kHz_r16 *PDCCH_MonitoringOccasions_r16
				}
				Pdsch_ProcessingType2_r16 *struct {
					Scs_15kHz_r16 *PDCCH_MonitoringOccasions_r16
					Scs_30kHz_r16 *PDCCH_MonitoringOccasions_r16
				}
			}{}
			c := ie.Pdcch_Monitoring_r16
			featureSetDownlinkV1610PdcchMonitoringR16Seq := d.NewSequenceDecoder(featureSetDownlinkV1610PdcchMonitoringR16Constraints)
			if err := featureSetDownlinkV1610PdcchMonitoringR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetDownlinkV1610PdcchMonitoringR16Seq.IsComponentPresent(0) {
				c.Pdsch_ProcessingType1_r16 = &struct {
					Scs_15kHz_r16 *PDCCH_MonitoringOccasions_r16
					Scs_30kHz_r16 *PDCCH_MonitoringOccasions_r16
				}{}
				c := (*c.Pdsch_ProcessingType1_r16)
				featureSetDownlinkV1610PdcchMonitoringR16PdschProcessingType1R16Seq := d.NewSequenceDecoder(featureSetDownlinkV1610PdcchMonitoringR16PdschProcessingType1R16Constraints)
				if err := featureSetDownlinkV1610PdcchMonitoringR16PdschProcessingType1R16Seq.DecodePreamble(); err != nil {
					return err
				}
				if featureSetDownlinkV1610PdcchMonitoringR16PdschProcessingType1R16Seq.IsComponentPresent(0) {
					c.Scs_15kHz_r16 = new(PDCCH_MonitoringOccasions_r16)
					if err := (*c.Scs_15kHz_r16).Decode(d); err != nil {
						return err
					}
				}
				if featureSetDownlinkV1610PdcchMonitoringR16PdschProcessingType1R16Seq.IsComponentPresent(1) {
					c.Scs_30kHz_r16 = new(PDCCH_MonitoringOccasions_r16)
					if err := (*c.Scs_30kHz_r16).Decode(d); err != nil {
						return err
					}
				}
			}
			if featureSetDownlinkV1610PdcchMonitoringR16Seq.IsComponentPresent(1) {
				c.Pdsch_ProcessingType2_r16 = &struct {
					Scs_15kHz_r16 *PDCCH_MonitoringOccasions_r16
					Scs_30kHz_r16 *PDCCH_MonitoringOccasions_r16
				}{}
				c := (*c.Pdsch_ProcessingType2_r16)
				featureSetDownlinkV1610PdcchMonitoringR16PdschProcessingType2R16Seq := d.NewSequenceDecoder(featureSetDownlinkV1610PdcchMonitoringR16PdschProcessingType2R16Constraints)
				if err := featureSetDownlinkV1610PdcchMonitoringR16PdschProcessingType2R16Seq.DecodePreamble(); err != nil {
					return err
				}
				if featureSetDownlinkV1610PdcchMonitoringR16PdschProcessingType2R16Seq.IsComponentPresent(0) {
					c.Scs_15kHz_r16 = new(PDCCH_MonitoringOccasions_r16)
					if err := (*c.Scs_15kHz_r16).Decode(d); err != nil {
						return err
					}
				}
				if featureSetDownlinkV1610PdcchMonitoringR16PdschProcessingType2R16Seq.IsComponentPresent(1) {
					c.Scs_30kHz_r16 = new(PDCCH_MonitoringOccasions_r16)
					if err := (*c.Scs_30kHz_r16).Decode(d); err != nil {
						return err
					}
				}
			}
		}
	}

	// 8. pdcch-MonitoringMixed-r16: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1610PdcchMonitoringMixedR16Constraints)
			if err != nil {
				return err
			}
			ie.Pdcch_MonitoringMixed_r16 = &idx
		}
	}

	// 9. crossCarrierSchedulingProcessing-DiffSCS-r16: sequence
	{
		if seq.IsComponentPresent(7) {
			ie.CrossCarrierSchedulingProcessing_DiffSCS_r16 = &struct {
				Scs_15kHz_120kHz_r16 *int64
				Scs_15kHz_60kHz_r16  *int64
				Scs_30kHz_120kHz_r16 *int64
				Scs_15kHz_30kHz_r16  *int64
				Scs_30kHz_60kHz_r16  *int64
				Scs_60kHz_120kHz_r16 *int64
			}{}
			c := ie.CrossCarrierSchedulingProcessing_DiffSCS_r16
			featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Seq := d.NewSequenceDecoder(featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Constraints)
			if err := featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Seq.DecodePreamble(); err != nil {
				return err
			}
			if featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Seq.IsComponentPresent(0) {
				c.Scs_15kHz_120kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs15kHz120kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_15kHz_120kHz_r16) = v
			}
			if featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Seq.IsComponentPresent(1) {
				c.Scs_15kHz_60kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs15kHz60kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_15kHz_60kHz_r16) = v
			}
			if featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Seq.IsComponentPresent(2) {
				c.Scs_30kHz_120kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs30kHz120kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_30kHz_120kHz_r16) = v
			}
			if featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Seq.IsComponentPresent(3) {
				c.Scs_15kHz_30kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs15kHz30kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_15kHz_30kHz_r16) = v
			}
			if featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Seq.IsComponentPresent(4) {
				c.Scs_30kHz_60kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs30kHz60kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_30kHz_60kHz_r16) = v
			}
			if featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Seq.IsComponentPresent(5) {
				c.Scs_60kHz_120kHz_r16 = new(int64)
				v, err := d.DecodeEnumerated(featureSetDownlinkV1610CrossCarrierSchedulingProcessingDiffSCSR16Scs60kHz120kHzR16Constraints)
				if err != nil {
					return err
				}
				(*c.Scs_60kHz_120kHz_r16) = v
			}
		}
	}

	// 10. singleDCI-SDM-scheme-r16: enumerated
	{
		if seq.IsComponentPresent(8) {
			idx, err := d.DecodeEnumerated(featureSetDownlinkV1610SingleDCISDMSchemeR16Constraints)
			if err != nil {
				return err
			}
			ie.SingleDCI_SDM_Scheme_r16 = &idx
		}
	}

	return nil
}
