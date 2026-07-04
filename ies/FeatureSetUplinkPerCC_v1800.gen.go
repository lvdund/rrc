// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetUplinkPerCC-v1800 (line 20575).

var featureSetUplinkPerCCV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "twoPUSCH-MultiDCI-STx2P-TwoTA-r18", Optional: true},
		{Name: "pusch-CB-SingleDCI-STx2P-SDM-r18", Optional: true},
		{Name: "pusch-NonCB-SingleDCI-STx2P-SDM-r18", Optional: true},
		{Name: "pusch-CB-SingleDCI-STx2P-SFN-r18", Optional: true},
		{Name: "pusch-NonCB-SingleDCI-STx2P-SFN-r18", Optional: true},
		{Name: "twoPUSCH-CB-MultiDCI-STx2P-DG-DG-r18", Optional: true},
		{Name: "twoPUSCH-NonCB-MultiDCI-STx2P-DG-DG-r18", Optional: true},
		{Name: "twoPUSCH-MultiDCI-STx2P-OutOfOrder-r18", Optional: true},
		{Name: "codebookParameter8TxPUSCH-r18", Optional: true},
		{Name: "nonCodebook-8TxPUSCH-r18", Optional: true},
		{Name: "nonCodebook-CSI-RS-SRS-r18", Optional: true},
		{Name: "cgb-2CW-PUSCH-r18", Optional: true},
	},
}

const (
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_MultiDCI_STx2P_TwoTA_r18_Supported = 0
)

var featureSetUplinkPerCCV1800TwoPUSCHMultiDCISTx2PTwoTAR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_TwoPUSCH_MultiDCI_STx2P_TwoTA_r18_Supported},
}

const (
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_MultiDCI_STx2P_OutOfOrder_r18_Supported = 0
)

var featureSetUplinkPerCCV1800TwoPUSCHMultiDCISTx2POutOfOrderR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_TwoPUSCH_MultiDCI_STx2P_OutOfOrder_r18_Supported},
}

const (
	FeatureSetUplinkPerCC_v1800_NonCodebook_CSI_RS_SRS_r18_Supported = 0
)

var featureSetUplinkPerCCV1800NonCodebookCSIRSSRSR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_NonCodebook_CSI_RS_SRS_r18_Supported},
}

const (
	FeatureSetUplinkPerCC_v1800_Cgb_2CW_PUSCH_r18_Supported = 0
)

var featureSetUplinkPerCCV1800Cgb2CWPUSCHR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_Cgb_2CW_PUSCH_r18_Supported},
}

const (
	FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SDM_r18_MaxNumberSRS_ResourcePerSet_r18_N1 = 0
	FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SDM_r18_MaxNumberSRS_ResourcePerSet_r18_N2 = 1
	FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SDM_r18_MaxNumberSRS_ResourcePerSet_r18_N4 = 2
)

var featureSetUplinkPerCCV1800PuschCBSingleDCISTx2PSDMR18MaxNumberSRSResourcePerSetR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SDM_r18_MaxNumberSRS_ResourcePerSet_r18_N1, FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SDM_r18_MaxNumberSRS_ResourcePerSet_r18_N2, FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SDM_r18_MaxNumberSRS_ResourcePerSet_r18_N4},
}

const (
	FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SDM_r18_MaxNumberNZP_PUSCH_PortsPerSet_r18_N1 = 0
	FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SDM_r18_MaxNumberNZP_PUSCH_PortsPerSet_r18_N2 = 1
	FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SDM_r18_MaxNumberNZP_PUSCH_PortsPerSet_r18_N4 = 2
)

var featureSetUplinkPerCCV1800PuschCBSingleDCISTx2PSDMR18MaxNumberNZPPUSCHPortsPerSetR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SDM_r18_MaxNumberNZP_PUSCH_PortsPerSet_r18_N1, FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SDM_r18_MaxNumberNZP_PUSCH_PortsPerSet_r18_N2, FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SDM_r18_MaxNumberNZP_PUSCH_PortsPerSet_r18_N4},
}

const (
	FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SDM_r18_MaxNumberSRS_AntennaPortsPerSet_r18_N1 = 0
	FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SDM_r18_MaxNumberSRS_AntennaPortsPerSet_r18_N2 = 1
	FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SDM_r18_MaxNumberSRS_AntennaPortsPerSet_r18_N4 = 2
)

var featureSetUplinkPerCCV1800PuschCBSingleDCISTx2PSDMR18MaxNumberSRSAntennaPortsPerSetR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SDM_r18_MaxNumberSRS_AntennaPortsPerSet_r18_N1, FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SDM_r18_MaxNumberSRS_AntennaPortsPerSet_r18_N2, FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SDM_r18_MaxNumberSRS_AntennaPortsPerSet_r18_N4},
}

const (
	FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SFN_r18_MaxNumberSRS_ResourcePerSet_r18_N1 = 0
	FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SFN_r18_MaxNumberSRS_ResourcePerSet_r18_N2 = 1
	FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SFN_r18_MaxNumberSRS_ResourcePerSet_r18_N4 = 2
)

var featureSetUplinkPerCCV1800PuschCBSingleDCISTx2PSFNR18MaxNumberSRSResourcePerSetR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SFN_r18_MaxNumberSRS_ResourcePerSet_r18_N1, FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SFN_r18_MaxNumberSRS_ResourcePerSet_r18_N2, FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SFN_r18_MaxNumberSRS_ResourcePerSet_r18_N4},
}

const (
	FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SFN_r18_MaxNumberSRS_AntennaPortsPerSet_r18_N1 = 0
	FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SFN_r18_MaxNumberSRS_AntennaPortsPerSet_r18_N2 = 1
	FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SFN_r18_MaxNumberSRS_AntennaPortsPerSet_r18_N4 = 2
)

var featureSetUplinkPerCCV1800PuschCBSingleDCISTx2PSFNR18MaxNumberSRSAntennaPortsPerSetR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SFN_r18_MaxNumberSRS_AntennaPortsPerSet_r18_N1, FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SFN_r18_MaxNumberSRS_AntennaPortsPerSet_r18_N2, FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SFN_r18_MaxNumberSRS_AntennaPortsPerSet_r18_N4},
}

const (
	FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SFN_r18_MaxNumberNZP_PUSCH_PortsPerSet_r18_N1 = 0
	FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SFN_r18_MaxNumberNZP_PUSCH_PortsPerSet_r18_N2 = 1
	FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SFN_r18_MaxNumberNZP_PUSCH_PortsPerSet_r18_N4 = 2
)

var featureSetUplinkPerCCV1800PuschCBSingleDCISTx2PSFNR18MaxNumberNZPPUSCHPortsPerSetR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SFN_r18_MaxNumberNZP_PUSCH_PortsPerSet_r18_N1, FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SFN_r18_MaxNumberNZP_PUSCH_PortsPerSet_r18_N2, FeatureSetUplinkPerCC_v1800_Pusch_CB_SingleDCI_STx2P_SFN_r18_MaxNumberNZP_PUSCH_PortsPerSet_r18_N4},
}

var featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberSRS-ResourcePerSet-r18"},
		{Name: "maxNumberLayerOverlapping-r18"},
		{Name: "maxNumberNZP-PUSCH-Overlapping-r18"},
		{Name: "maxNumberPUSCH-PerCORESET-PerSlot-r18", Optional: true},
		{Name: "maxNumberTotalLayerOverlapping-r18"},
		{Name: "maxNumberSRS-AntennaPortsPerSet-r18"},
	},
}

const (
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberSRS_ResourcePerSet_r18_N1 = 0
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberSRS_ResourcePerSet_r18_N2 = 1
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberSRS_ResourcePerSet_r18_N4 = 2
)

var featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberSRSResourcePerSetR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberSRS_ResourcePerSet_r18_N1, FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberSRS_ResourcePerSet_r18_N2, FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberSRS_ResourcePerSet_r18_N4},
}

const (
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberNZP_PUSCH_Overlapping_r18_N1 = 0
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberNZP_PUSCH_Overlapping_r18_N2 = 1
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberNZP_PUSCH_Overlapping_r18_N4 = 2
)

var featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberNZPPUSCHOverlappingR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberNZP_PUSCH_Overlapping_r18_N1, FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberNZP_PUSCH_Overlapping_r18_N2, FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberNZP_PUSCH_Overlapping_r18_N4},
}

var featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz-r18", Optional: true},
		{Name: "scs-120kHz-r18", Optional: true},
	},
}

const (
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_60kHz_r18_N1 = 0
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_60kHz_r18_N2 = 1
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_60kHz_r18_N3 = 2
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_60kHz_r18_N4 = 3
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_60kHz_r18_N7 = 4
)

var featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Scs60kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_60kHz_r18_N1, FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_60kHz_r18_N2, FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_60kHz_r18_N3, FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_60kHz_r18_N4, FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_60kHz_r18_N7},
}

const (
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_120kHz_r18_N1 = 0
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_120kHz_r18_N2 = 1
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_120kHz_r18_N3 = 2
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_120kHz_r18_N4 = 3
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_120kHz_r18_N7 = 4
)

var featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Scs120kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_120kHz_r18_N1, FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_120kHz_r18_N2, FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_120kHz_r18_N3, FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_120kHz_r18_N4, FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_120kHz_r18_N7},
}

const (
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberSRS_AntennaPortsPerSet_r18_N1 = 0
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberSRS_AntennaPortsPerSet_r18_N2 = 1
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberSRS_AntennaPortsPerSet_r18_N4 = 2
)

var featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberSRSAntennaPortsPerSetR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberSRS_AntennaPortsPerSet_r18_N1, FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberSRS_AntennaPortsPerSet_r18_N2, FeatureSetUplinkPerCC_v1800_TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18_MaxNumberSRS_AntennaPortsPerSet_r18_N4},
}

var featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxNumberSRS-ResourcePerSet-r18"},
		{Name: "maxNumberLayerOverlapping-r18"},
		{Name: "maxNumberSimulSRS-ResourcePerSet-r18"},
		{Name: "maxNumberPUSCH-PerCORESET-PerSlot-r18", Optional: true},
		{Name: "maxNumberTotalLayerOverlapping-r18"},
	},
}

var featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scs-60kHz-r18", Optional: true},
		{Name: "scs-120kHz-r18", Optional: true},
	},
}

const (
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_60kHz_r18_N1 = 0
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_60kHz_r18_N2 = 1
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_60kHz_r18_N3 = 2
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_60kHz_r18_N4 = 3
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_60kHz_r18_N7 = 4
)

var featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Scs60kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_60kHz_r18_N1, FeatureSetUplinkPerCC_v1800_TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_60kHz_r18_N2, FeatureSetUplinkPerCC_v1800_TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_60kHz_r18_N3, FeatureSetUplinkPerCC_v1800_TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_60kHz_r18_N4, FeatureSetUplinkPerCC_v1800_TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_60kHz_r18_N7},
}

const (
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_120kHz_r18_N1 = 0
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_120kHz_r18_N2 = 1
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_120kHz_r18_N3 = 2
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_120kHz_r18_N4 = 3
	FeatureSetUplinkPerCC_v1800_TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_120kHz_r18_N7 = 4
)

var featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Scs120kHzR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_120kHz_r18_N1, FeatureSetUplinkPerCC_v1800_TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_120kHz_r18_N2, FeatureSetUplinkPerCC_v1800_TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_120kHz_r18_N3, FeatureSetUplinkPerCC_v1800_TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_120kHz_r18_N4, FeatureSetUplinkPerCC_v1800_TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18_MaxNumberPUSCH_PerCORESET_PerSlot_r18_Scs_120kHz_r18_N7},
}

var featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "codebook-8TxBasic-r18"},
		{Name: "codebook1-8TxPUSCH-r18"},
		{Name: "codebook2-8TxPUSCH-r18", Optional: true},
		{Name: "codebook3-8TxPUSCH-r18", Optional: true},
		{Name: "codebook4-8TxPUSCH-r18", Optional: true},
		{Name: "ul-FullPwrTransMode0-r18", Optional: true},
		{Name: "ul-FullPwrTransMode1-r18", Optional: true},
		{Name: "ul-FullPwrTransMode2-r18", Optional: true},
		{Name: "ul-SRS-TransMode2-r18", Optional: true},
		{Name: "tpmi-FullPwrCodebook2-r18", Optional: true},
	},
}

const (
	FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Codebook_8TxBasic_r18_Srs_8TxPorts_r18_NoTDM = 0
	FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Codebook_8TxBasic_r18_Srs_8TxPorts_r18_Both  = 1
)

var featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook8TxBasicR18Srs8TxPortsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Codebook_8TxBasic_r18_Srs_8TxPorts_r18_NoTDM, FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Codebook_8TxBasic_r18_Srs_8TxPorts_r18_Both},
}

var featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook18TxPUSCHR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "codebookN1N4-r18", Optional: true},
		{Name: "srs-8TxPorts-r18"},
	},
}

const (
	FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Codebook1_8TxPUSCH_r18_CodebookN1N4_r18_Ng1n4n1 = 0
	FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Codebook1_8TxPUSCH_r18_CodebookN1N4_r18_Ng1n2n2 = 1
	FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Codebook1_8TxPUSCH_r18_CodebookN1N4_r18_Both    = 2
)

var featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook18TxPUSCHR18CodebookN1N4R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Codebook1_8TxPUSCH_r18_CodebookN1N4_r18_Ng1n4n1, FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Codebook1_8TxPUSCH_r18_CodebookN1N4_r18_Ng1n2n2, FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Codebook1_8TxPUSCH_r18_CodebookN1N4_r18_Both},
}

const (
	FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Codebook1_8TxPUSCH_r18_Srs_8TxPorts_r18_NoTDM = 0
	FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Codebook1_8TxPUSCH_r18_Srs_8TxPorts_r18_Both  = 1
)

var featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook18TxPUSCHR18Srs8TxPortsR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Codebook1_8TxPUSCH_r18_Srs_8TxPorts_r18_NoTDM, FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Codebook1_8TxPUSCH_r18_Srs_8TxPorts_r18_Both},
}

const (
	FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Codebook2_8TxPUSCH_r18_Supported = 0
)

var featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook28TxPUSCHR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Codebook2_8TxPUSCH_r18_Supported},
}

const (
	FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Codebook3_8TxPUSCH_r18_Supported = 0
)

var featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook38TxPUSCHR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Codebook3_8TxPUSCH_r18_Supported},
}

const (
	FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Codebook4_8TxPUSCH_r18_Supported = 0
)

var featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook48TxPUSCHR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Codebook4_8TxPUSCH_r18_Supported},
}

const (
	FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Ul_FullPwrTransMode0_r18_Supported = 0
)

var featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18UlFullPwrTransMode0R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Ul_FullPwrTransMode0_r18_Supported},
}

const (
	FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Ul_FullPwrTransMode1_r18_Supported = 0
)

var featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18UlFullPwrTransMode1R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Ul_FullPwrTransMode1_r18_Supported},
}

const (
	FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Ul_FullPwrTransMode2_r18_N1 = 0
	FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Ul_FullPwrTransMode2_r18_N2 = 1
	FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Ul_FullPwrTransMode2_r18_N4 = 2
)

var featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18UlFullPwrTransMode2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Ul_FullPwrTransMode2_r18_N1, FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Ul_FullPwrTransMode2_r18_N2, FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Ul_FullPwrTransMode2_r18_N4},
}

const (
	FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Tpmi_FullPwrCodebook2_r18_First  = 0
	FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Tpmi_FullPwrCodebook2_r18_Second = 1
)

var featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18TpmiFullPwrCodebook2R18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Tpmi_FullPwrCodebook2_r18_First, FeatureSetUplinkPerCC_v1800_CodebookParameter8TxPUSCH_r18_Tpmi_FullPwrCodebook2_r18_Second},
}

type FeatureSetUplinkPerCC_v1800 struct {
	TwoPUSCH_MultiDCI_STx2P_TwoTA_r18 *int64
	Pusch_CB_SingleDCI_STx2P_SDM_r18  *struct {
		MaxNumberSRS_ResourcePerSet_r18     int64
		MaxNumberLayerPerPanel_r18          int64
		MaxNumberNZP_PUSCH_PortsPerSet_r18  int64
		MaxNumberSRS_AntennaPortsPerSet_r18 int64
	}
	Pusch_NonCB_SingleDCI_STx2P_SDM_r18 *struct {
		MaxNumberSRS_ResourcePerSet_r18         int64
		MaxNumberLayerPerPanel_r18              int64
		MaxNumberSimulSRS_OneResourcePerSet_r18 int64
		MaxNumberSimulSRS_TwoResourcePerSet_r18 int64
	}
	Pusch_CB_SingleDCI_STx2P_SFN_r18 *struct {
		MaxNumberSRS_ResourcePerSet_r18     int64
		MaxNumberLayerPerSet_r18            int64
		MaxNumberSRS_AntennaPortsPerSet_r18 int64
		MaxNumberNZP_PUSCH_PortsPerSet_r18  int64
	}
	Pusch_NonCB_SingleDCI_STx2P_SFN_r18 *struct {
		MaxNumberSRS_ResourcePerSet_r18         int64
		MaxNumberLayerPerSet_r18                int64
		MaxNumberSimulSRS_OneResourcePerSet_r18 int64
		MaxNumberSimulSRS_TwoResourcePerSet_r18 int64
	}
	TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18 *struct {
		MaxNumberSRS_ResourcePerSet_r18       int64
		MaxNumberLayerOverlapping_r18         int64
		MaxNumberNZP_PUSCH_Overlapping_r18    int64
		MaxNumberPUSCH_PerCORESET_PerSlot_r18 *struct {
			Scs_60kHz_r18  *int64
			Scs_120kHz_r18 *int64
		}
		MaxNumberTotalLayerOverlapping_r18  int64
		MaxNumberSRS_AntennaPortsPerSet_r18 int64
	}
	TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18 *struct {
		MaxNumberSRS_ResourcePerSet_r18       int64
		MaxNumberLayerOverlapping_r18         int64
		MaxNumberSimulSRS_ResourcePerSet_r18  int64
		MaxNumberPUSCH_PerCORESET_PerSlot_r18 *struct {
			Scs_60kHz_r18  *int64
			Scs_120kHz_r18 *int64
		}
		MaxNumberTotalLayerOverlapping_r18 int64
	}
	TwoPUSCH_MultiDCI_STx2P_OutOfOrder_r18 *int64
	CodebookParameter8TxPUSCH_r18          *struct {
		Codebook_8TxBasic_r18 struct {
			MaxNumberPUSCH_MIMO_Layer_r18 int64
			MaxNumberSRS_Resource_r18     int64
			Srs_8TxPorts_r18              int64
		}
		Codebook1_8TxPUSCH_r18 struct {
			CodebookN1N4_r18 *int64
			Srs_8TxPorts_r18 int64
		}
		Codebook2_8TxPUSCH_r18    *int64
		Codebook3_8TxPUSCH_r18    *int64
		Codebook4_8TxPUSCH_r18    *int64
		Ul_FullPwrTransMode0_r18  *int64
		Ul_FullPwrTransMode1_r18  *int64
		Ul_FullPwrTransMode2_r18  *int64
		Ul_SRS_TransMode2_r18     *per.BitString
		Tpmi_FullPwrCodebook2_r18 *int64
	}
	NonCodebook_8TxPUSCH_r18 *struct {
		MaxNumberPUSCH_MIMO_Layer_r18 int64
		MaxNumberSRS_Resource_r18     int64
		MaxNumberSimultaneousSRS_r18  int64
	}
	NonCodebook_CSI_RS_SRS_r18 *int64
	Cgb_2CW_PUSCH_r18          *int64
}

func (ie *FeatureSetUplinkPerCC_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetUplinkPerCCV1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.TwoPUSCH_MultiDCI_STx2P_TwoTA_r18 != nil, ie.Pusch_CB_SingleDCI_STx2P_SDM_r18 != nil, ie.Pusch_NonCB_SingleDCI_STx2P_SDM_r18 != nil, ie.Pusch_CB_SingleDCI_STx2P_SFN_r18 != nil, ie.Pusch_NonCB_SingleDCI_STx2P_SFN_r18 != nil, ie.TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18 != nil, ie.TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18 != nil, ie.TwoPUSCH_MultiDCI_STx2P_OutOfOrder_r18 != nil, ie.CodebookParameter8TxPUSCH_r18 != nil, ie.NonCodebook_8TxPUSCH_r18 != nil, ie.NonCodebook_CSI_RS_SRS_r18 != nil, ie.Cgb_2CW_PUSCH_r18 != nil}); err != nil {
		return err
	}

	// 2. twoPUSCH-MultiDCI-STx2P-TwoTA-r18: enumerated
	{
		if ie.TwoPUSCH_MultiDCI_STx2P_TwoTA_r18 != nil {
			if err := e.EncodeEnumerated(*ie.TwoPUSCH_MultiDCI_STx2P_TwoTA_r18, featureSetUplinkPerCCV1800TwoPUSCHMultiDCISTx2PTwoTAR18Constraints); err != nil {
				return err
			}
		}
	}

	// 3. pusch-CB-SingleDCI-STx2P-SDM-r18: sequence
	{
		if ie.Pusch_CB_SingleDCI_STx2P_SDM_r18 != nil {
			c := ie.Pusch_CB_SingleDCI_STx2P_SDM_r18
			if err := e.EncodeEnumerated(c.MaxNumberSRS_ResourcePerSet_r18, featureSetUplinkPerCCV1800PuschCBSingleDCISTx2PSDMR18MaxNumberSRSResourcePerSetR18Constraints); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberLayerPerPanel_r18, per.Constrained(1, 2)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MaxNumberNZP_PUSCH_PortsPerSet_r18, featureSetUplinkPerCCV1800PuschCBSingleDCISTx2PSDMR18MaxNumberNZPPUSCHPortsPerSetR18Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MaxNumberSRS_AntennaPortsPerSet_r18, featureSetUplinkPerCCV1800PuschCBSingleDCISTx2PSDMR18MaxNumberSRSAntennaPortsPerSetR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. pusch-NonCB-SingleDCI-STx2P-SDM-r18: sequence
	{
		if ie.Pusch_NonCB_SingleDCI_STx2P_SDM_r18 != nil {
			c := ie.Pusch_NonCB_SingleDCI_STx2P_SDM_r18
			if err := e.EncodeInteger(c.MaxNumberSRS_ResourcePerSet_r18, per.Constrained(1, 4)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberLayerPerPanel_r18, per.Constrained(1, 2)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberSimulSRS_OneResourcePerSet_r18, per.Constrained(1, 4)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberSimulSRS_TwoResourcePerSet_r18, per.Constrained(1, 8)); err != nil {
				return err
			}
		}
	}

	// 5. pusch-CB-SingleDCI-STx2P-SFN-r18: sequence
	{
		if ie.Pusch_CB_SingleDCI_STx2P_SFN_r18 != nil {
			c := ie.Pusch_CB_SingleDCI_STx2P_SFN_r18
			if err := e.EncodeEnumerated(c.MaxNumberSRS_ResourcePerSet_r18, featureSetUplinkPerCCV1800PuschCBSingleDCISTx2PSFNR18MaxNumberSRSResourcePerSetR18Constraints); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberLayerPerSet_r18, per.Constrained(1, 2)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MaxNumberSRS_AntennaPortsPerSet_r18, featureSetUplinkPerCCV1800PuschCBSingleDCISTx2PSFNR18MaxNumberSRSAntennaPortsPerSetR18Constraints); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MaxNumberNZP_PUSCH_PortsPerSet_r18, featureSetUplinkPerCCV1800PuschCBSingleDCISTx2PSFNR18MaxNumberNZPPUSCHPortsPerSetR18Constraints); err != nil {
				return err
			}
		}
	}

	// 6. pusch-NonCB-SingleDCI-STx2P-SFN-r18: sequence
	{
		if ie.Pusch_NonCB_SingleDCI_STx2P_SFN_r18 != nil {
			c := ie.Pusch_NonCB_SingleDCI_STx2P_SFN_r18
			if err := e.EncodeInteger(c.MaxNumberSRS_ResourcePerSet_r18, per.Constrained(1, 4)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberLayerPerSet_r18, per.Constrained(1, 2)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberSimulSRS_OneResourcePerSet_r18, per.Constrained(1, 4)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberSimulSRS_TwoResourcePerSet_r18, per.Constrained(1, 8)); err != nil {
				return err
			}
		}
	}

	// 7. twoPUSCH-CB-MultiDCI-STx2P-DG-DG-r18: sequence
	{
		if ie.TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18 != nil {
			c := ie.TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18
			featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18Seq := e.NewSequenceEncoder(featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18Constraints)
			if err := featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18Seq.EncodePreamble([]bool{c.MaxNumberPUSCH_PerCORESET_PerSlot_r18 != nil}); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MaxNumberSRS_ResourcePerSet_r18, featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberSRSResourcePerSetR18Constraints); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberLayerOverlapping_r18, per.Constrained(1, 2)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MaxNumberNZP_PUSCH_Overlapping_r18, featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberNZPPUSCHOverlappingR18Constraints); err != nil {
				return err
			}
			if c.MaxNumberPUSCH_PerCORESET_PerSlot_r18 != nil {
				c := (*c.MaxNumberPUSCH_PerCORESET_PerSlot_r18)
				featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Seq := e.NewSequenceEncoder(featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Constraints)
				if err := featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Seq.EncodePreamble([]bool{c.Scs_60kHz_r18 != nil, c.Scs_120kHz_r18 != nil}); err != nil {
					return err
				}
				if c.Scs_60kHz_r18 != nil {
					if err := e.EncodeEnumerated((*c.Scs_60kHz_r18), featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Scs60kHzR18Constraints); err != nil {
						return err
					}
				}
				if c.Scs_120kHz_r18 != nil {
					if err := e.EncodeEnumerated((*c.Scs_120kHz_r18), featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Scs120kHzR18Constraints); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeInteger(c.MaxNumberTotalLayerOverlapping_r18, per.Constrained(2, 4)); err != nil {
				return err
			}
			if err := e.EncodeEnumerated(c.MaxNumberSRS_AntennaPortsPerSet_r18, featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberSRSAntennaPortsPerSetR18Constraints); err != nil {
				return err
			}
		}
	}

	// 8. twoPUSCH-NonCB-MultiDCI-STx2P-DG-DG-r18: sequence
	{
		if ie.TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18 != nil {
			c := ie.TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18
			featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18Seq := e.NewSequenceEncoder(featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18Constraints)
			if err := featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18Seq.EncodePreamble([]bool{c.MaxNumberPUSCH_PerCORESET_PerSlot_r18 != nil}); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberSRS_ResourcePerSet_r18, per.Constrained(1, 4)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberLayerOverlapping_r18, per.Constrained(1, 2)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberSimulSRS_ResourcePerSet_r18, per.Constrained(1, 4)); err != nil {
				return err
			}
			if c.MaxNumberPUSCH_PerCORESET_PerSlot_r18 != nil {
				c := (*c.MaxNumberPUSCH_PerCORESET_PerSlot_r18)
				featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Seq := e.NewSequenceEncoder(featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Constraints)
				if err := featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Seq.EncodePreamble([]bool{c.Scs_60kHz_r18 != nil, c.Scs_120kHz_r18 != nil}); err != nil {
					return err
				}
				if c.Scs_60kHz_r18 != nil {
					if err := e.EncodeEnumerated((*c.Scs_60kHz_r18), featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Scs60kHzR18Constraints); err != nil {
						return err
					}
				}
				if c.Scs_120kHz_r18 != nil {
					if err := e.EncodeEnumerated((*c.Scs_120kHz_r18), featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Scs120kHzR18Constraints); err != nil {
						return err
					}
				}
			}
			if err := e.EncodeInteger(c.MaxNumberTotalLayerOverlapping_r18, per.Constrained(2, 4)); err != nil {
				return err
			}
		}
	}

	// 9. twoPUSCH-MultiDCI-STx2P-OutOfOrder-r18: enumerated
	{
		if ie.TwoPUSCH_MultiDCI_STx2P_OutOfOrder_r18 != nil {
			if err := e.EncodeEnumerated(*ie.TwoPUSCH_MultiDCI_STx2P_OutOfOrder_r18, featureSetUplinkPerCCV1800TwoPUSCHMultiDCISTx2POutOfOrderR18Constraints); err != nil {
				return err
			}
		}
	}

	// 10. codebookParameter8TxPUSCH-r18: sequence
	{
		if ie.CodebookParameter8TxPUSCH_r18 != nil {
			c := ie.CodebookParameter8TxPUSCH_r18
			featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Seq := e.NewSequenceEncoder(featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Constraints)
			if err := featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Seq.EncodePreamble([]bool{c.Codebook2_8TxPUSCH_r18 != nil, c.Codebook3_8TxPUSCH_r18 != nil, c.Codebook4_8TxPUSCH_r18 != nil, c.Ul_FullPwrTransMode0_r18 != nil, c.Ul_FullPwrTransMode1_r18 != nil, c.Ul_FullPwrTransMode2_r18 != nil, c.Ul_SRS_TransMode2_r18 != nil, c.Tpmi_FullPwrCodebook2_r18 != nil}); err != nil {
				return err
			}
			{
				c := &c.Codebook_8TxBasic_r18
				if err := e.EncodeInteger(c.MaxNumberPUSCH_MIMO_Layer_r18, per.Constrained(1, 8)); err != nil {
					return err
				}
				if err := e.EncodeInteger(c.MaxNumberSRS_Resource_r18, per.Constrained(1, 2)); err != nil {
					return err
				}
				if err := e.EncodeEnumerated(c.Srs_8TxPorts_r18, featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook8TxBasicR18Srs8TxPortsR18Constraints); err != nil {
					return err
				}
			}
			{
				c := &c.Codebook1_8TxPUSCH_r18
				featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook18TxPUSCHR18Seq := e.NewSequenceEncoder(featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook18TxPUSCHR18Constraints)
				if err := featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook18TxPUSCHR18Seq.EncodePreamble([]bool{c.CodebookN1N4_r18 != nil}); err != nil {
					return err
				}
				if c.CodebookN1N4_r18 != nil {
					if err := e.EncodeEnumerated((*c.CodebookN1N4_r18), featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook18TxPUSCHR18CodebookN1N4R18Constraints); err != nil {
						return err
					}
				}
				if err := e.EncodeEnumerated(c.Srs_8TxPorts_r18, featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook18TxPUSCHR18Srs8TxPortsR18Constraints); err != nil {
					return err
				}
			}
			if c.Codebook2_8TxPUSCH_r18 != nil {
				if err := e.EncodeEnumerated((*c.Codebook2_8TxPUSCH_r18), featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook28TxPUSCHR18Constraints); err != nil {
					return err
				}
			}
			if c.Codebook3_8TxPUSCH_r18 != nil {
				if err := e.EncodeEnumerated((*c.Codebook3_8TxPUSCH_r18), featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook38TxPUSCHR18Constraints); err != nil {
					return err
				}
			}
			if c.Codebook4_8TxPUSCH_r18 != nil {
				if err := e.EncodeEnumerated((*c.Codebook4_8TxPUSCH_r18), featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook48TxPUSCHR18Constraints); err != nil {
					return err
				}
			}
			if c.Ul_FullPwrTransMode0_r18 != nil {
				if err := e.EncodeEnumerated((*c.Ul_FullPwrTransMode0_r18), featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18UlFullPwrTransMode0R18Constraints); err != nil {
					return err
				}
			}
			if c.Ul_FullPwrTransMode1_r18 != nil {
				if err := e.EncodeEnumerated((*c.Ul_FullPwrTransMode1_r18), featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18UlFullPwrTransMode1R18Constraints); err != nil {
					return err
				}
			}
			if c.Ul_FullPwrTransMode2_r18 != nil {
				if err := e.EncodeEnumerated((*c.Ul_FullPwrTransMode2_r18), featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18UlFullPwrTransMode2R18Constraints); err != nil {
					return err
				}
			}
			if c.Ul_SRS_TransMode2_r18 != nil {
				if err := e.EncodeBitString((*c.Ul_SRS_TransMode2_r18), per.FixedSize(3)); err != nil {
					return err
				}
			}
			if c.Tpmi_FullPwrCodebook2_r18 != nil {
				if err := e.EncodeEnumerated((*c.Tpmi_FullPwrCodebook2_r18), featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18TpmiFullPwrCodebook2R18Constraints); err != nil {
					return err
				}
			}
		}
	}

	// 11. nonCodebook-8TxPUSCH-r18: sequence
	{
		if ie.NonCodebook_8TxPUSCH_r18 != nil {
			c := ie.NonCodebook_8TxPUSCH_r18
			if err := e.EncodeInteger(c.MaxNumberPUSCH_MIMO_Layer_r18, per.Constrained(1, 8)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberSRS_Resource_r18, per.Constrained(1, 8)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberSimultaneousSRS_r18, per.Constrained(1, 8)); err != nil {
				return err
			}
		}
	}

	// 12. nonCodebook-CSI-RS-SRS-r18: enumerated
	{
		if ie.NonCodebook_CSI_RS_SRS_r18 != nil {
			if err := e.EncodeEnumerated(*ie.NonCodebook_CSI_RS_SRS_r18, featureSetUplinkPerCCV1800NonCodebookCSIRSSRSR18Constraints); err != nil {
				return err
			}
		}
	}

	// 13. cgb-2CW-PUSCH-r18: enumerated
	{
		if ie.Cgb_2CW_PUSCH_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Cgb_2CW_PUSCH_r18, featureSetUplinkPerCCV1800Cgb2CWPUSCHR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetUplinkPerCC_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetUplinkPerCCV1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. twoPUSCH-MultiDCI-STx2P-TwoTA-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800TwoPUSCHMultiDCISTx2PTwoTAR18Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUSCH_MultiDCI_STx2P_TwoTA_r18 = &idx
		}
	}

	// 3. pusch-CB-SingleDCI-STx2P-SDM-r18: sequence
	{
		if seq.IsComponentPresent(1) {
			ie.Pusch_CB_SingleDCI_STx2P_SDM_r18 = &struct {
				MaxNumberSRS_ResourcePerSet_r18     int64
				MaxNumberLayerPerPanel_r18          int64
				MaxNumberNZP_PUSCH_PortsPerSet_r18  int64
				MaxNumberSRS_AntennaPortsPerSet_r18 int64
			}{}
			c := ie.Pusch_CB_SingleDCI_STx2P_SDM_r18
			{
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800PuschCBSingleDCISTx2PSDMR18MaxNumberSRSResourcePerSetR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberSRS_ResourcePerSet_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.MaxNumberLayerPerPanel_r18 = v
			}
			{
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800PuschCBSingleDCISTx2PSDMR18MaxNumberNZPPUSCHPortsPerSetR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberNZP_PUSCH_PortsPerSet_r18 = v
			}
			{
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800PuschCBSingleDCISTx2PSDMR18MaxNumberSRSAntennaPortsPerSetR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberSRS_AntennaPortsPerSet_r18 = v
			}
		}
	}

	// 4. pusch-NonCB-SingleDCI-STx2P-SDM-r18: sequence
	{
		if seq.IsComponentPresent(2) {
			ie.Pusch_NonCB_SingleDCI_STx2P_SDM_r18 = &struct {
				MaxNumberSRS_ResourcePerSet_r18         int64
				MaxNumberLayerPerPanel_r18              int64
				MaxNumberSimulSRS_OneResourcePerSet_r18 int64
				MaxNumberSimulSRS_TwoResourcePerSet_r18 int64
			}{}
			c := ie.Pusch_NonCB_SingleDCI_STx2P_SDM_r18
			{
				v, err := d.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumberSRS_ResourcePerSet_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.MaxNumberLayerPerPanel_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumberSimulSRS_OneResourcePerSet_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumberSimulSRS_TwoResourcePerSet_r18 = v
			}
		}
	}

	// 5. pusch-CB-SingleDCI-STx2P-SFN-r18: sequence
	{
		if seq.IsComponentPresent(3) {
			ie.Pusch_CB_SingleDCI_STx2P_SFN_r18 = &struct {
				MaxNumberSRS_ResourcePerSet_r18     int64
				MaxNumberLayerPerSet_r18            int64
				MaxNumberSRS_AntennaPortsPerSet_r18 int64
				MaxNumberNZP_PUSCH_PortsPerSet_r18  int64
			}{}
			c := ie.Pusch_CB_SingleDCI_STx2P_SFN_r18
			{
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800PuschCBSingleDCISTx2PSFNR18MaxNumberSRSResourcePerSetR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberSRS_ResourcePerSet_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.MaxNumberLayerPerSet_r18 = v
			}
			{
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800PuschCBSingleDCISTx2PSFNR18MaxNumberSRSAntennaPortsPerSetR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberSRS_AntennaPortsPerSet_r18 = v
			}
			{
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800PuschCBSingleDCISTx2PSFNR18MaxNumberNZPPUSCHPortsPerSetR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberNZP_PUSCH_PortsPerSet_r18 = v
			}
		}
	}

	// 6. pusch-NonCB-SingleDCI-STx2P-SFN-r18: sequence
	{
		if seq.IsComponentPresent(4) {
			ie.Pusch_NonCB_SingleDCI_STx2P_SFN_r18 = &struct {
				MaxNumberSRS_ResourcePerSet_r18         int64
				MaxNumberLayerPerSet_r18                int64
				MaxNumberSimulSRS_OneResourcePerSet_r18 int64
				MaxNumberSimulSRS_TwoResourcePerSet_r18 int64
			}{}
			c := ie.Pusch_NonCB_SingleDCI_STx2P_SFN_r18
			{
				v, err := d.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumberSRS_ResourcePerSet_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.MaxNumberLayerPerSet_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumberSimulSRS_OneResourcePerSet_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumberSimulSRS_TwoResourcePerSet_r18 = v
			}
		}
	}

	// 7. twoPUSCH-CB-MultiDCI-STx2P-DG-DG-r18: sequence
	{
		if seq.IsComponentPresent(5) {
			ie.TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18 = &struct {
				MaxNumberSRS_ResourcePerSet_r18       int64
				MaxNumberLayerOverlapping_r18         int64
				MaxNumberNZP_PUSCH_Overlapping_r18    int64
				MaxNumberPUSCH_PerCORESET_PerSlot_r18 *struct {
					Scs_60kHz_r18  *int64
					Scs_120kHz_r18 *int64
				}
				MaxNumberTotalLayerOverlapping_r18  int64
				MaxNumberSRS_AntennaPortsPerSet_r18 int64
			}{}
			c := ie.TwoPUSCH_CB_MultiDCI_STx2P_DG_DG_r18
			featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18Seq := d.NewSequenceDecoder(featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18Constraints)
			if err := featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberSRSResourcePerSetR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberSRS_ResourcePerSet_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.MaxNumberLayerOverlapping_r18 = v
			}
			{
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberNZPPUSCHOverlappingR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberNZP_PUSCH_Overlapping_r18 = v
			}
			if featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18Seq.IsComponentPresent(3) {
				c.MaxNumberPUSCH_PerCORESET_PerSlot_r18 = &struct {
					Scs_60kHz_r18  *int64
					Scs_120kHz_r18 *int64
				}{}
				c := (*c.MaxNumberPUSCH_PerCORESET_PerSlot_r18)
				featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Seq := d.NewSequenceDecoder(featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Constraints)
				if err := featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Seq.DecodePreamble(); err != nil {
					return err
				}
				if featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Seq.IsComponentPresent(0) {
					c.Scs_60kHz_r18 = new(int64)
					v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Scs60kHzR18Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_60kHz_r18) = v
				}
				if featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Seq.IsComponentPresent(1) {
					c.Scs_120kHz_r18 = new(int64)
					v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Scs120kHzR18Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_120kHz_r18) = v
				}
			}
			{
				v, err := d.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				c.MaxNumberTotalLayerOverlapping_r18 = v
			}
			{
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800TwoPUSCHCBMultiDCISTx2PDGDGR18MaxNumberSRSAntennaPortsPerSetR18Constraints)
				if err != nil {
					return err
				}
				c.MaxNumberSRS_AntennaPortsPerSet_r18 = v
			}
		}
	}

	// 8. twoPUSCH-NonCB-MultiDCI-STx2P-DG-DG-r18: sequence
	{
		if seq.IsComponentPresent(6) {
			ie.TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18 = &struct {
				MaxNumberSRS_ResourcePerSet_r18       int64
				MaxNumberLayerOverlapping_r18         int64
				MaxNumberSimulSRS_ResourcePerSet_r18  int64
				MaxNumberPUSCH_PerCORESET_PerSlot_r18 *struct {
					Scs_60kHz_r18  *int64
					Scs_120kHz_r18 *int64
				}
				MaxNumberTotalLayerOverlapping_r18 int64
			}{}
			c := ie.TwoPUSCH_NonCB_MultiDCI_STx2P_DG_DG_r18
			featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18Seq := d.NewSequenceDecoder(featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18Constraints)
			if err := featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumberSRS_ResourcePerSet_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 2))
				if err != nil {
					return err
				}
				c.MaxNumberLayerOverlapping_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumberSimulSRS_ResourcePerSet_r18 = v
			}
			if featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18Seq.IsComponentPresent(3) {
				c.MaxNumberPUSCH_PerCORESET_PerSlot_r18 = &struct {
					Scs_60kHz_r18  *int64
					Scs_120kHz_r18 *int64
				}{}
				c := (*c.MaxNumberPUSCH_PerCORESET_PerSlot_r18)
				featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Seq := d.NewSequenceDecoder(featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Constraints)
				if err := featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Seq.DecodePreamble(); err != nil {
					return err
				}
				if featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Seq.IsComponentPresent(0) {
					c.Scs_60kHz_r18 = new(int64)
					v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Scs60kHzR18Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_60kHz_r18) = v
				}
				if featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Seq.IsComponentPresent(1) {
					c.Scs_120kHz_r18 = new(int64)
					v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800TwoPUSCHNonCBMultiDCISTx2PDGDGR18MaxNumberPUSCHPerCORESETPerSlotR18Scs120kHzR18Constraints)
					if err != nil {
						return err
					}
					(*c.Scs_120kHz_r18) = v
				}
			}
			{
				v, err := d.DecodeInteger(per.Constrained(2, 4))
				if err != nil {
					return err
				}
				c.MaxNumberTotalLayerOverlapping_r18 = v
			}
		}
	}

	// 9. twoPUSCH-MultiDCI-STx2P-OutOfOrder-r18: enumerated
	{
		if seq.IsComponentPresent(7) {
			idx, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800TwoPUSCHMultiDCISTx2POutOfOrderR18Constraints)
			if err != nil {
				return err
			}
			ie.TwoPUSCH_MultiDCI_STx2P_OutOfOrder_r18 = &idx
		}
	}

	// 10. codebookParameter8TxPUSCH-r18: sequence
	{
		if seq.IsComponentPresent(8) {
			ie.CodebookParameter8TxPUSCH_r18 = &struct {
				Codebook_8TxBasic_r18 struct {
					MaxNumberPUSCH_MIMO_Layer_r18 int64
					MaxNumberSRS_Resource_r18     int64
					Srs_8TxPorts_r18              int64
				}
				Codebook1_8TxPUSCH_r18 struct {
					CodebookN1N4_r18 *int64
					Srs_8TxPorts_r18 int64
				}
				Codebook2_8TxPUSCH_r18    *int64
				Codebook3_8TxPUSCH_r18    *int64
				Codebook4_8TxPUSCH_r18    *int64
				Ul_FullPwrTransMode0_r18  *int64
				Ul_FullPwrTransMode1_r18  *int64
				Ul_FullPwrTransMode2_r18  *int64
				Ul_SRS_TransMode2_r18     *per.BitString
				Tpmi_FullPwrCodebook2_r18 *int64
			}{}
			c := ie.CodebookParameter8TxPUSCH_r18
			featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Seq := d.NewSequenceDecoder(featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Constraints)
			if err := featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Seq.DecodePreamble(); err != nil {
				return err
			}
			{
				c := &c.Codebook_8TxBasic_r18
				{
					v, err := d.DecodeInteger(per.Constrained(1, 8))
					if err != nil {
						return err
					}
					c.MaxNumberPUSCH_MIMO_Layer_r18 = v
				}
				{
					v, err := d.DecodeInteger(per.Constrained(1, 2))
					if err != nil {
						return err
					}
					c.MaxNumberSRS_Resource_r18 = v
				}
				{
					v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook8TxBasicR18Srs8TxPortsR18Constraints)
					if err != nil {
						return err
					}
					c.Srs_8TxPorts_r18 = v
				}
			}
			{
				c := &c.Codebook1_8TxPUSCH_r18
				featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook18TxPUSCHR18Seq := d.NewSequenceDecoder(featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook18TxPUSCHR18Constraints)
				if err := featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook18TxPUSCHR18Seq.DecodePreamble(); err != nil {
					return err
				}
				if featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook18TxPUSCHR18Seq.IsComponentPresent(0) {
					c.CodebookN1N4_r18 = new(int64)
					v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook18TxPUSCHR18CodebookN1N4R18Constraints)
					if err != nil {
						return err
					}
					(*c.CodebookN1N4_r18) = v
				}
				{
					v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook18TxPUSCHR18Srs8TxPortsR18Constraints)
					if err != nil {
						return err
					}
					c.Srs_8TxPorts_r18 = v
				}
			}
			if featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Seq.IsComponentPresent(2) {
				c.Codebook2_8TxPUSCH_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook28TxPUSCHR18Constraints)
				if err != nil {
					return err
				}
				(*c.Codebook2_8TxPUSCH_r18) = v
			}
			if featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Seq.IsComponentPresent(3) {
				c.Codebook3_8TxPUSCH_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook38TxPUSCHR18Constraints)
				if err != nil {
					return err
				}
				(*c.Codebook3_8TxPUSCH_r18) = v
			}
			if featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Seq.IsComponentPresent(4) {
				c.Codebook4_8TxPUSCH_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Codebook48TxPUSCHR18Constraints)
				if err != nil {
					return err
				}
				(*c.Codebook4_8TxPUSCH_r18) = v
			}
			if featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Seq.IsComponentPresent(5) {
				c.Ul_FullPwrTransMode0_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18UlFullPwrTransMode0R18Constraints)
				if err != nil {
					return err
				}
				(*c.Ul_FullPwrTransMode0_r18) = v
			}
			if featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Seq.IsComponentPresent(6) {
				c.Ul_FullPwrTransMode1_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18UlFullPwrTransMode1R18Constraints)
				if err != nil {
					return err
				}
				(*c.Ul_FullPwrTransMode1_r18) = v
			}
			if featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Seq.IsComponentPresent(7) {
				c.Ul_FullPwrTransMode2_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18UlFullPwrTransMode2R18Constraints)
				if err != nil {
					return err
				}
				(*c.Ul_FullPwrTransMode2_r18) = v
			}
			if featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Seq.IsComponentPresent(8) {
				c.Ul_SRS_TransMode2_r18 = new(per.BitString)
				v, err := d.DecodeBitString(per.FixedSize(3))
				if err != nil {
					return err
				}
				(*c.Ul_SRS_TransMode2_r18) = v
			}
			if featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18Seq.IsComponentPresent(9) {
				c.Tpmi_FullPwrCodebook2_r18 = new(int64)
				v, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800CodebookParameter8TxPUSCHR18TpmiFullPwrCodebook2R18Constraints)
				if err != nil {
					return err
				}
				(*c.Tpmi_FullPwrCodebook2_r18) = v
			}
		}
	}

	// 11. nonCodebook-8TxPUSCH-r18: sequence
	{
		if seq.IsComponentPresent(9) {
			ie.NonCodebook_8TxPUSCH_r18 = &struct {
				MaxNumberPUSCH_MIMO_Layer_r18 int64
				MaxNumberSRS_Resource_r18     int64
				MaxNumberSimultaneousSRS_r18  int64
			}{}
			c := ie.NonCodebook_8TxPUSCH_r18
			{
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumberPUSCH_MIMO_Layer_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumberSRS_Resource_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 8))
				if err != nil {
					return err
				}
				c.MaxNumberSimultaneousSRS_r18 = v
			}
		}
	}

	// 12. nonCodebook-CSI-RS-SRS-r18: enumerated
	{
		if seq.IsComponentPresent(10) {
			idx, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800NonCodebookCSIRSSRSR18Constraints)
			if err != nil {
				return err
			}
			ie.NonCodebook_CSI_RS_SRS_r18 = &idx
		}
	}

	// 13. cgb-2CW-PUSCH-r18: enumerated
	{
		if seq.IsComponentPresent(11) {
			idx, err := d.DecodeEnumerated(featureSetUplinkPerCCV1800Cgb2CWPUSCHR18Constraints)
			if err != nil {
				return err
			}
			ie.Cgb_2CW_PUSCH_r18 = &idx
		}
	}

	return nil
}
