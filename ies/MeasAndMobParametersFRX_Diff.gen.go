// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasAndMobParametersFRX-Diff (line 21323).

var measAndMobParametersFRXDiffConstraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "ss-SINR-Meas", Optional: true},
		{Name: "csi-RSRP-AndRSRQ-MeasWithSSB", Optional: true},
		{Name: "csi-RSRP-AndRSRQ-MeasWithoutSSB", Optional: true},
		{Name: "csi-SINR-Meas", Optional: true},
		{Name: "csi-RS-RLM", Optional: true},
	},
	ExtComponents: []per.ComponentInfo{
		{Name: "extension-addition-0", Optional: true},
		{Name: "extension-addition-1", Optional: true},
		{Name: "extension-addition-2", Optional: true},
		{Name: "extension-addition-3", Optional: true},
		{Name: "extension-addition-4", Optional: true},
	},
}

const (
	MeasAndMobParametersFRX_Diff_Ss_SINR_Meas_Supported = 0
)

var measAndMobParametersFRXDiffSsSINRMeasConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Ss_SINR_Meas_Supported},
}

const (
	MeasAndMobParametersFRX_Diff_Csi_RSRP_AndRSRQ_MeasWithSSB_Supported = 0
)

var measAndMobParametersFRXDiffCsiRSRPAndRSRQMeasWithSSBConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Csi_RSRP_AndRSRQ_MeasWithSSB_Supported},
}

const (
	MeasAndMobParametersFRX_Diff_Csi_RSRP_AndRSRQ_MeasWithoutSSB_Supported = 0
)

var measAndMobParametersFRXDiffCsiRSRPAndRSRQMeasWithoutSSBConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Csi_RSRP_AndRSRQ_MeasWithoutSSB_Supported},
}

const (
	MeasAndMobParametersFRX_Diff_Csi_SINR_Meas_Supported = 0
)

var measAndMobParametersFRXDiffCsiSINRMeasConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Csi_SINR_Meas_Supported},
}

const (
	MeasAndMobParametersFRX_Diff_Csi_RS_RLM_Supported = 0
)

var measAndMobParametersFRXDiffCsiRSRLMConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Csi_RS_RLM_Supported},
}

const (
	MeasAndMobParametersFRX_Diff_Ext_HandoverInterF_Supported = 0
)

var measAndMobParametersFRXDiffExtHandoverInterFConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Ext_HandoverInterF_Supported},
}

const (
	MeasAndMobParametersFRX_Diff_Ext_HandoverLTE_EPC_Supported = 0
)

var measAndMobParametersFRXDiffExtHandoverLTEEPCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Ext_HandoverLTE_EPC_Supported},
}

const (
	MeasAndMobParametersFRX_Diff_Ext_HandoverLTE_5GC_Supported = 0
)

var measAndMobParametersFRXDiffExtHandoverLTE5GCConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Ext_HandoverLTE_5GC_Supported},
}

const (
	MeasAndMobParametersFRX_Diff_Ext_MaxNumberResource_CSI_RS_RLM_N2 = 0
	MeasAndMobParametersFRX_Diff_Ext_MaxNumberResource_CSI_RS_RLM_N4 = 1
	MeasAndMobParametersFRX_Diff_Ext_MaxNumberResource_CSI_RS_RLM_N6 = 2
	MeasAndMobParametersFRX_Diff_Ext_MaxNumberResource_CSI_RS_RLM_N8 = 3
)

var measAndMobParametersFRXDiffExtMaxNumberResourceCSIRSRLMConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Ext_MaxNumberResource_CSI_RS_RLM_N2, MeasAndMobParametersFRX_Diff_Ext_MaxNumberResource_CSI_RS_RLM_N4, MeasAndMobParametersFRX_Diff_Ext_MaxNumberResource_CSI_RS_RLM_N6, MeasAndMobParametersFRX_Diff_Ext_MaxNumberResource_CSI_RS_RLM_N8},
}

const (
	MeasAndMobParametersFRX_Diff_Ext_SimultaneousRxDataSSB_DiffNumerology_Supported = 0
)

var measAndMobParametersFRXDiffExtSimultaneousRxDataSSBDiffNumerologyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Ext_SimultaneousRxDataSSB_DiffNumerology_Supported},
}

const (
	MeasAndMobParametersFRX_Diff_Ext_Nr_AutonomousGaps_r16_Supported = 0
)

var measAndMobParametersFRXDiffExtNrAutonomousGapsR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Ext_Nr_AutonomousGaps_r16_Supported},
}

const (
	MeasAndMobParametersFRX_Diff_Ext_Nr_AutonomousGaps_ENDC_r16_Supported = 0
)

var measAndMobParametersFRXDiffExtNrAutonomousGapsENDCR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Ext_Nr_AutonomousGaps_ENDC_r16_Supported},
}

const (
	MeasAndMobParametersFRX_Diff_Ext_Nr_AutonomousGaps_NEDC_r16_Supported = 0
)

var measAndMobParametersFRXDiffExtNrAutonomousGapsNEDCR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Ext_Nr_AutonomousGaps_NEDC_r16_Supported},
}

const (
	MeasAndMobParametersFRX_Diff_Ext_Nr_AutonomousGaps_NRDC_r16_Supported = 0
)

var measAndMobParametersFRXDiffExtNrAutonomousGapsNRDCR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Ext_Nr_AutonomousGaps_NRDC_r16_Supported},
}

const (
	MeasAndMobParametersFRX_Diff_Ext_Dummy_Supported = 0
)

var measAndMobParametersFRXDiffExtDummyConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Ext_Dummy_Supported},
}

const (
	MeasAndMobParametersFRX_Diff_Ext_Cli_RSSI_Meas_r16_Supported = 0
)

var measAndMobParametersFRXDiffExtCliRSSIMeasR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Ext_Cli_RSSI_Meas_r16_Supported},
}

const (
	MeasAndMobParametersFRX_Diff_Ext_Cli_SRS_RSRP_Meas_r16_Supported = 0
)

var measAndMobParametersFRXDiffExtCliSRSRSRPMeasR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Ext_Cli_SRS_RSRP_Meas_r16_Supported},
}

const (
	MeasAndMobParametersFRX_Diff_Ext_InterFrequencyMeas_NoGap_r16_Supported = 0
)

var measAndMobParametersFRXDiffExtInterFrequencyMeasNoGapR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Ext_InterFrequencyMeas_NoGap_r16_Supported},
}

const (
	MeasAndMobParametersFRX_Diff_Ext_SimultaneousRxDataSSB_DiffNumerology_Inter_r16_Supported = 0
)

var measAndMobParametersFRXDiffExtSimultaneousRxDataSSBDiffNumerologyInterR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Ext_SimultaneousRxDataSSB_DiffNumerology_Inter_r16_Supported},
}

const (
	MeasAndMobParametersFRX_Diff_Ext_IdleInactiveNR_MeasReport_r16_Supported = 0
)

var measAndMobParametersFRXDiffExtIdleInactiveNRMeasReportR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Ext_IdleInactiveNR_MeasReport_r16_Supported},
}

const (
	MeasAndMobParametersFRX_Diff_Ext_IdleInactiveNR_MeasBeamReport_r16_Supported = 0
)

var measAndMobParametersFRXDiffExtIdleInactiveNRMeasBeamReportR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Ext_IdleInactiveNR_MeasBeamReport_r16_Supported},
}

const (
	MeasAndMobParametersFRX_Diff_Ext_IncreasedNumberofCSIRSPerMO_r16_Supported = 0
)

var measAndMobParametersFRXDiffExtIncreasedNumberofCSIRSPerMOR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MeasAndMobParametersFRX_Diff_Ext_IncreasedNumberofCSIRSPerMO_r16_Supported},
}

type MeasAndMobParametersFRX_Diff struct {
	Ss_SINR_Meas                                   *int64
	Csi_RSRP_AndRSRQ_MeasWithSSB                   *int64
	Csi_RSRP_AndRSRQ_MeasWithoutSSB                *int64
	Csi_SINR_Meas                                  *int64
	Csi_RS_RLM                                     *int64
	HandoverInterF                                 *int64
	HandoverLTE_EPC                                *int64
	HandoverLTE_5GC                                *int64
	MaxNumberResource_CSI_RS_RLM                   *int64
	SimultaneousRxDataSSB_DiffNumerology           *int64
	Nr_AutonomousGaps_r16                          *int64
	Nr_AutonomousGaps_ENDC_r16                     *int64
	Nr_AutonomousGaps_NEDC_r16                     *int64
	Nr_AutonomousGaps_NRDC_r16                     *int64
	Dummy                                          *int64
	Cli_RSSI_Meas_r16                              *int64
	Cli_SRS_RSRP_Meas_r16                          *int64
	InterFrequencyMeas_NoGap_r16                   *int64
	SimultaneousRxDataSSB_DiffNumerology_Inter_r16 *int64
	IdleInactiveNR_MeasReport_r16                  *int64
	IdleInactiveNR_MeasBeamReport_r16              *int64
	IncreasedNumberofCSIRSPerMO_r16                *int64
}

func (ie *MeasAndMobParametersFRX_Diff) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measAndMobParametersFRXDiffConstraints)

	// Extension-group presence (one var per [[...]] group).
	hasExtGroup0 := ie.HandoverInterF != nil || ie.HandoverLTE_EPC != nil || ie.HandoverLTE_5GC != nil
	hasExtGroup1 := ie.MaxNumberResource_CSI_RS_RLM != nil
	hasExtGroup2 := ie.SimultaneousRxDataSSB_DiffNumerology != nil
	hasExtGroup3 := ie.Nr_AutonomousGaps_r16 != nil || ie.Nr_AutonomousGaps_ENDC_r16 != nil || ie.Nr_AutonomousGaps_NEDC_r16 != nil || ie.Nr_AutonomousGaps_NRDC_r16 != nil || ie.Dummy != nil || ie.Cli_RSSI_Meas_r16 != nil || ie.Cli_SRS_RSRP_Meas_r16 != nil || ie.InterFrequencyMeas_NoGap_r16 != nil || ie.SimultaneousRxDataSSB_DiffNumerology_Inter_r16 != nil || ie.IdleInactiveNR_MeasReport_r16 != nil || ie.IdleInactiveNR_MeasBeamReport_r16 != nil
	hasExtGroup4 := ie.IncreasedNumberofCSIRSPerMO_r16 != nil
	hasExtensions := hasExtGroup0 || hasExtGroup1 || hasExtGroup2 || hasExtGroup3 || hasExtGroup4

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(hasExtensions); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ss_SINR_Meas != nil, ie.Csi_RSRP_AndRSRQ_MeasWithSSB != nil, ie.Csi_RSRP_AndRSRQ_MeasWithoutSSB != nil, ie.Csi_SINR_Meas != nil, ie.Csi_RS_RLM != nil}); err != nil {
		return err
	}

	// 3. ss-SINR-Meas: enumerated
	{
		if ie.Ss_SINR_Meas != nil {
			if err := e.EncodeEnumerated(*ie.Ss_SINR_Meas, measAndMobParametersFRXDiffSsSINRMeasConstraints); err != nil {
				return err
			}
		}
	}

	// 4. csi-RSRP-AndRSRQ-MeasWithSSB: enumerated
	{
		if ie.Csi_RSRP_AndRSRQ_MeasWithSSB != nil {
			if err := e.EncodeEnumerated(*ie.Csi_RSRP_AndRSRQ_MeasWithSSB, measAndMobParametersFRXDiffCsiRSRPAndRSRQMeasWithSSBConstraints); err != nil {
				return err
			}
		}
	}

	// 5. csi-RSRP-AndRSRQ-MeasWithoutSSB: enumerated
	{
		if ie.Csi_RSRP_AndRSRQ_MeasWithoutSSB != nil {
			if err := e.EncodeEnumerated(*ie.Csi_RSRP_AndRSRQ_MeasWithoutSSB, measAndMobParametersFRXDiffCsiRSRPAndRSRQMeasWithoutSSBConstraints); err != nil {
				return err
			}
		}
	}

	// 6. csi-SINR-Meas: enumerated
	{
		if ie.Csi_SINR_Meas != nil {
			if err := e.EncodeEnumerated(*ie.Csi_SINR_Meas, measAndMobParametersFRXDiffCsiSINRMeasConstraints); err != nil {
				return err
			}
		}
	}

	// 7. csi-RS-RLM: enumerated
	{
		if ie.Csi_RS_RLM != nil {
			if err := e.EncodeEnumerated(*ie.Csi_RS_RLM, measAndMobParametersFRXDiffCsiRSRLMConstraints); err != nil {
				return err
			}
		}
	}

	if hasExtensions {
		extPresence := []bool{hasExtGroup0, hasExtGroup1, hasExtGroup2, hasExtGroup3, hasExtGroup4}
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
				if err := ex.EncodeEnumerated(*ie.HandoverInterF, measAndMobParametersFRXDiffExtHandoverInterFConstraints); err != nil {
					return err
				}
			}

			if ie.HandoverLTE_EPC != nil {
				if err := ex.EncodeEnumerated(*ie.HandoverLTE_EPC, measAndMobParametersFRXDiffExtHandoverLTEEPCConstraints); err != nil {
					return err
				}
			}

			if ie.HandoverLTE_5GC != nil {
				if err := ex.EncodeEnumerated(*ie.HandoverLTE_5GC, measAndMobParametersFRXDiffExtHandoverLTE5GCConstraints); err != nil {
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
					{Name: "maxNumberResource-CSI-RS-RLM", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.MaxNumberResource_CSI_RS_RLM != nil}); err != nil {
				return err
			}

			if ie.MaxNumberResource_CSI_RS_RLM != nil {
				if err := ex.EncodeEnumerated(*ie.MaxNumberResource_CSI_RS_RLM, measAndMobParametersFRXDiffExtMaxNumberResourceCSIRSRLMConstraints); err != nil {
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
					{Name: "simultaneousRxDataSSB-DiffNumerology", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.SimultaneousRxDataSSB_DiffNumerology != nil}); err != nil {
				return err
			}

			if ie.SimultaneousRxDataSSB_DiffNumerology != nil {
				if err := ex.EncodeEnumerated(*ie.SimultaneousRxDataSSB_DiffNumerology, measAndMobParametersFRXDiffExtSimultaneousRxDataSSBDiffNumerologyConstraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 3:
		if hasExtGroup3 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "nr-AutonomousGaps-r16", Optional: true},
					{Name: "nr-AutonomousGaps-ENDC-r16", Optional: true},
					{Name: "nr-AutonomousGaps-NEDC-r16", Optional: true},
					{Name: "nr-AutonomousGaps-NRDC-r16", Optional: true},
					{Name: "dummy", Optional: true},
					{Name: "cli-RSSI-Meas-r16", Optional: true},
					{Name: "cli-SRS-RSRP-Meas-r16", Optional: true},
					{Name: "interFrequencyMeas-NoGap-r16", Optional: true},
					{Name: "simultaneousRxDataSSB-DiffNumerology-Inter-r16", Optional: true},
					{Name: "idleInactiveNR-MeasReport-r16", Optional: true},
					{Name: "idleInactiveNR-MeasBeamReport-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.Nr_AutonomousGaps_r16 != nil, ie.Nr_AutonomousGaps_ENDC_r16 != nil, ie.Nr_AutonomousGaps_NEDC_r16 != nil, ie.Nr_AutonomousGaps_NRDC_r16 != nil, ie.Dummy != nil, ie.Cli_RSSI_Meas_r16 != nil, ie.Cli_SRS_RSRP_Meas_r16 != nil, ie.InterFrequencyMeas_NoGap_r16 != nil, ie.SimultaneousRxDataSSB_DiffNumerology_Inter_r16 != nil, ie.IdleInactiveNR_MeasReport_r16 != nil, ie.IdleInactiveNR_MeasBeamReport_r16 != nil}); err != nil {
				return err
			}

			if ie.Nr_AutonomousGaps_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Nr_AutonomousGaps_r16, measAndMobParametersFRXDiffExtNrAutonomousGapsR16Constraints); err != nil {
					return err
				}
			}

			if ie.Nr_AutonomousGaps_ENDC_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Nr_AutonomousGaps_ENDC_r16, measAndMobParametersFRXDiffExtNrAutonomousGapsENDCR16Constraints); err != nil {
					return err
				}
			}

			if ie.Nr_AutonomousGaps_NEDC_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Nr_AutonomousGaps_NEDC_r16, measAndMobParametersFRXDiffExtNrAutonomousGapsNEDCR16Constraints); err != nil {
					return err
				}
			}

			if ie.Nr_AutonomousGaps_NRDC_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Nr_AutonomousGaps_NRDC_r16, measAndMobParametersFRXDiffExtNrAutonomousGapsNRDCR16Constraints); err != nil {
					return err
				}
			}

			if ie.Dummy != nil {
				if err := ex.EncodeEnumerated(*ie.Dummy, measAndMobParametersFRXDiffExtDummyConstraints); err != nil {
					return err
				}
			}

			if ie.Cli_RSSI_Meas_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Cli_RSSI_Meas_r16, measAndMobParametersFRXDiffExtCliRSSIMeasR16Constraints); err != nil {
					return err
				}
			}

			if ie.Cli_SRS_RSRP_Meas_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.Cli_SRS_RSRP_Meas_r16, measAndMobParametersFRXDiffExtCliSRSRSRPMeasR16Constraints); err != nil {
					return err
				}
			}

			if ie.InterFrequencyMeas_NoGap_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.InterFrequencyMeas_NoGap_r16, measAndMobParametersFRXDiffExtInterFrequencyMeasNoGapR16Constraints); err != nil {
					return err
				}
			}

			if ie.SimultaneousRxDataSSB_DiffNumerology_Inter_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.SimultaneousRxDataSSB_DiffNumerology_Inter_r16, measAndMobParametersFRXDiffExtSimultaneousRxDataSSBDiffNumerologyInterR16Constraints); err != nil {
					return err
				}
			}

			if ie.IdleInactiveNR_MeasReport_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.IdleInactiveNR_MeasReport_r16, measAndMobParametersFRXDiffExtIdleInactiveNRMeasReportR16Constraints); err != nil {
					return err
				}
			}

			if ie.IdleInactiveNR_MeasBeamReport_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.IdleInactiveNR_MeasBeamReport_r16, measAndMobParametersFRXDiffExtIdleInactiveNRMeasBeamReportR16Constraints); err != nil {
					return err
				}
			}
			extValues = append(extValues, ex.Bytes())
		}

		// Extension group 4:
		if hasExtGroup4 {
			ex := per.NewEncoder(e.Variant())
			groupSeq := ex.NewSequenceEncoder(per.SequenceConstraints{
				Extensible: false,
				RootComponents: []per.ComponentInfo{
					{Name: "increasedNumberofCSIRSPerMO-r16", Optional: true},
				},
			})
			if err := groupSeq.EncodePreamble([]bool{ie.IncreasedNumberofCSIRSPerMO_r16 != nil}); err != nil {
				return err
			}

			if ie.IncreasedNumberofCSIRSPerMO_r16 != nil {
				if err := ex.EncodeEnumerated(*ie.IncreasedNumberofCSIRSPerMO_r16, measAndMobParametersFRXDiffExtIncreasedNumberofCSIRSPerMOR16Constraints); err != nil {
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

func (ie *MeasAndMobParametersFRX_Diff) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measAndMobParametersFRXDiffConstraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. ss-SINR-Meas: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(measAndMobParametersFRXDiffSsSINRMeasConstraints)
			if err != nil {
				return err
			}
			ie.Ss_SINR_Meas = &idx
		}
	}

	// 4. csi-RSRP-AndRSRQ-MeasWithSSB: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(measAndMobParametersFRXDiffCsiRSRPAndRSRQMeasWithSSBConstraints)
			if err != nil {
				return err
			}
			ie.Csi_RSRP_AndRSRQ_MeasWithSSB = &idx
		}
	}

	// 5. csi-RSRP-AndRSRQ-MeasWithoutSSB: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(measAndMobParametersFRXDiffCsiRSRPAndRSRQMeasWithoutSSBConstraints)
			if err != nil {
				return err
			}
			ie.Csi_RSRP_AndRSRQ_MeasWithoutSSB = &idx
		}
	}

	// 6. csi-SINR-Meas: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(measAndMobParametersFRXDiffCsiSINRMeasConstraints)
			if err != nil {
				return err
			}
			ie.Csi_SINR_Meas = &idx
		}
	}

	// 7. csi-RS-RLM: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(measAndMobParametersFRXDiffCsiRSRLMConstraints)
			if err != nil {
				return err
			}
			ie.Csi_RS_RLM = &idx
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
			v, err := dx.DecodeEnumerated(measAndMobParametersFRXDiffExtHandoverInterFConstraints)
			if err != nil {
				return err
			}
			ie.HandoverInterF = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(measAndMobParametersFRXDiffExtHandoverLTEEPCConstraints)
			if err != nil {
				return err
			}
			ie.HandoverLTE_EPC = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(measAndMobParametersFRXDiffExtHandoverLTE5GCConstraints)
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
				{Name: "maxNumberResource-CSI-RS-RLM", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measAndMobParametersFRXDiffExtMaxNumberResourceCSIRSRLMConstraints)
			if err != nil {
				return err
			}
			ie.MaxNumberResource_CSI_RS_RLM = &v
		}
	}

	// Extension group 2:
	if seq.IsExtensionPresent(2) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "simultaneousRxDataSSB-DiffNumerology", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measAndMobParametersFRXDiffExtSimultaneousRxDataSSBDiffNumerologyConstraints)
			if err != nil {
				return err
			}
			ie.SimultaneousRxDataSSB_DiffNumerology = &v
		}
	}

	// Extension group 3:
	if seq.IsExtensionPresent(3) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "nr-AutonomousGaps-r16", Optional: true},
				{Name: "nr-AutonomousGaps-ENDC-r16", Optional: true},
				{Name: "nr-AutonomousGaps-NEDC-r16", Optional: true},
				{Name: "nr-AutonomousGaps-NRDC-r16", Optional: true},
				{Name: "dummy", Optional: true},
				{Name: "cli-RSSI-Meas-r16", Optional: true},
				{Name: "cli-SRS-RSRP-Meas-r16", Optional: true},
				{Name: "interFrequencyMeas-NoGap-r16", Optional: true},
				{Name: "simultaneousRxDataSSB-DiffNumerology-Inter-r16", Optional: true},
				{Name: "idleInactiveNR-MeasReport-r16", Optional: true},
				{Name: "idleInactiveNR-MeasBeamReport-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measAndMobParametersFRXDiffExtNrAutonomousGapsR16Constraints)
			if err != nil {
				return err
			}
			ie.Nr_AutonomousGaps_r16 = &v
		}

		if groupSeq.IsComponentPresent(1) {
			v, err := dx.DecodeEnumerated(measAndMobParametersFRXDiffExtNrAutonomousGapsENDCR16Constraints)
			if err != nil {
				return err
			}
			ie.Nr_AutonomousGaps_ENDC_r16 = &v
		}

		if groupSeq.IsComponentPresent(2) {
			v, err := dx.DecodeEnumerated(measAndMobParametersFRXDiffExtNrAutonomousGapsNEDCR16Constraints)
			if err != nil {
				return err
			}
			ie.Nr_AutonomousGaps_NEDC_r16 = &v
		}

		if groupSeq.IsComponentPresent(3) {
			v, err := dx.DecodeEnumerated(measAndMobParametersFRXDiffExtNrAutonomousGapsNRDCR16Constraints)
			if err != nil {
				return err
			}
			ie.Nr_AutonomousGaps_NRDC_r16 = &v
		}

		if groupSeq.IsComponentPresent(4) {
			v, err := dx.DecodeEnumerated(measAndMobParametersFRXDiffExtDummyConstraints)
			if err != nil {
				return err
			}
			ie.Dummy = &v
		}

		if groupSeq.IsComponentPresent(5) {
			v, err := dx.DecodeEnumerated(measAndMobParametersFRXDiffExtCliRSSIMeasR16Constraints)
			if err != nil {
				return err
			}
			ie.Cli_RSSI_Meas_r16 = &v
		}

		if groupSeq.IsComponentPresent(6) {
			v, err := dx.DecodeEnumerated(measAndMobParametersFRXDiffExtCliSRSRSRPMeasR16Constraints)
			if err != nil {
				return err
			}
			ie.Cli_SRS_RSRP_Meas_r16 = &v
		}

		if groupSeq.IsComponentPresent(7) {
			v, err := dx.DecodeEnumerated(measAndMobParametersFRXDiffExtInterFrequencyMeasNoGapR16Constraints)
			if err != nil {
				return err
			}
			ie.InterFrequencyMeas_NoGap_r16 = &v
		}

		if groupSeq.IsComponentPresent(8) {
			v, err := dx.DecodeEnumerated(measAndMobParametersFRXDiffExtSimultaneousRxDataSSBDiffNumerologyInterR16Constraints)
			if err != nil {
				return err
			}
			ie.SimultaneousRxDataSSB_DiffNumerology_Inter_r16 = &v
		}

		if groupSeq.IsComponentPresent(9) {
			v, err := dx.DecodeEnumerated(measAndMobParametersFRXDiffExtIdleInactiveNRMeasReportR16Constraints)
			if err != nil {
				return err
			}
			ie.IdleInactiveNR_MeasReport_r16 = &v
		}

		if groupSeq.IsComponentPresent(10) {
			v, err := dx.DecodeEnumerated(measAndMobParametersFRXDiffExtIdleInactiveNRMeasBeamReportR16Constraints)
			if err != nil {
				return err
			}
			ie.IdleInactiveNR_MeasBeamReport_r16 = &v
		}
	}

	// Extension group 4:
	if seq.IsExtensionPresent(4) {
		dx := per.NewDecoder(extValues[extValueIndex], d.Variant())
		extValueIndex++
		groupSeq := dx.NewSequenceDecoder(per.SequenceConstraints{
			Extensible: false,
			RootComponents: []per.ComponentInfo{
				{Name: "increasedNumberofCSIRSPerMO-r16", Optional: true},
			},
		})
		if err := groupSeq.DecodePreamble(); err != nil {
			return err
		}

		if groupSeq.IsComponentPresent(0) {
			v, err := dx.DecodeEnumerated(measAndMobParametersFRXDiffExtIncreasedNumberofCSIRSPerMOR16Constraints)
			if err != nil {
				return err
			}
			ie.IncreasedNumberofCSIRSPerMO_r16 = &v
		}
	}

	return nil
}
