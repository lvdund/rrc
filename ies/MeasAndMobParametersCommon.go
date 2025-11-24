package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersCommon struct {
	SupportedGapPattern                          *uper.BitString                                                          `lb:22,ub:22,optional`
	Ssb_RLM                                      *MeasAndMobParametersCommon_ssb_RLM                                      `optional`
	Ssb_AndCSI_RS_RLM                            *MeasAndMobParametersCommon_ssb_AndCSI_RS_RLM                            `optional`
	EventB_MeasAndReport                         *MeasAndMobParametersCommon_eventB_MeasAndReport                         `optional,ext-1`
	HandoverFDD_TDD                              *MeasAndMobParametersCommon_handoverFDD_TDD                              `optional,ext-1`
	Eutra_CGI_Reporting                          *MeasAndMobParametersCommon_eutra_CGI_Reporting                          `optional,ext-1`
	Nr_CGI_Reporting                             *MeasAndMobParametersCommon_nr_CGI_Reporting                             `optional,ext-1`
	IndependentGapConfig                         *MeasAndMobParametersCommon_independentGapConfig                         `optional,ext-2`
	PeriodicEUTRA_MeasAndReport                  *MeasAndMobParametersCommon_periodicEUTRA_MeasAndReport                  `optional,ext-2`
	HandoverFR1_FR2                              *MeasAndMobParametersCommon_handoverFR1_FR2                              `optional,ext-2`
	MaxNumberCSI_RS_RRM_RS_SINR                  *MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR                  `optional,ext-2`
	Nr_CGI_Reporting_ENDC                        *MeasAndMobParametersCommon_nr_CGI_Reporting_ENDC                        `optional,ext-3`
	Eutra_CGI_Reporting_NEDC                     *MeasAndMobParametersCommon_eutra_CGI_Reporting_NEDC                     `optional,ext-4`
	Eutra_CGI_Reporting_NRDC                     *MeasAndMobParametersCommon_eutra_CGI_Reporting_NRDC                     `optional,ext-4`
	Nr_CGI_Reporting_NEDC                        *MeasAndMobParametersCommon_nr_CGI_Reporting_NEDC                        `optional,ext-4`
	Nr_CGI_Reporting_NRDC                        *MeasAndMobParametersCommon_nr_CGI_Reporting_NRDC                        `optional,ext-4`
	ReportAddNeighMeasForPeriodic_r16            *MeasAndMobParametersCommon_reportAddNeighMeasForPeriodic_r16            `optional,ext-5`
	CondHandoverParametersCommon_r16             *MeasAndMobParametersCommon_condHandoverParametersCommon_r16             `optional,ext-5`
	Nr_NeedForGap_Reporting_r16                  *MeasAndMobParametersCommon_nr_NeedForGap_Reporting_r16                  `optional,ext-5`
	SupportedGapPattern_NRonly_r16               *uper.BitString                                                          `lb:10,ub:10,optional,ext-5`
	SupportedGapPattern_NRonly_NEDC_r16          *MeasAndMobParametersCommon_supportedGapPattern_NRonly_NEDC_r16          `optional,ext-5`
	MaxNumberCLI_RSSI_r16                        *MeasAndMobParametersCommon_maxNumberCLI_RSSI_r16                        `optional,ext-5`
	MaxNumberCLI_SRS_RSRP_r16                    *MeasAndMobParametersCommon_maxNumberCLI_SRS_RSRP_r16                    `optional,ext-5`
	MaxNumberPerSlotCLI_SRS_RSRP_r16             *MeasAndMobParametersCommon_maxNumberPerSlotCLI_SRS_RSRP_r16             `optional,ext-5`
	Mfbi_IAB_r16                                 *MeasAndMobParametersCommon_mfbi_IAB_r16                                 `optional,ext-5`
	Dummy                                        *MeasAndMobParametersCommon_dummy                                        `optional,ext-5`
	Nr_CGI_Reporting_NPN_r16                     *MeasAndMobParametersCommon_nr_CGI_Reporting_NPN_r16                     `optional,ext-5`
	IdleInactiveEUTRA_MeasReport_r16             *MeasAndMobParametersCommon_idleInactiveEUTRA_MeasReport_r16             `optional,ext-5`
	IdleInactive_ValidityArea_r16                *MeasAndMobParametersCommon_idleInactive_ValidityArea_r16                `optional,ext-5`
	Eutra_AutonomousGaps_r16                     *MeasAndMobParametersCommon_eutra_AutonomousGaps_r16                     `optional,ext-5`
	Eutra_AutonomousGaps_NEDC_r16                *MeasAndMobParametersCommon_eutra_AutonomousGaps_NEDC_r16                `optional,ext-5`
	Eutra_AutonomousGaps_NRDC_r16                *MeasAndMobParametersCommon_eutra_AutonomousGaps_NRDC_r16                `optional,ext-5`
	PcellT312_r16                                *MeasAndMobParametersCommon_pcellT312_r16                                `optional,ext-5`
	SupportedGapPattern_r16                      *uper.BitString                                                          `lb:2,ub:2,optional,ext-5`
	ConcurrentMeasGap_r17                        *MeasAndMobParametersCommon_concurrentMeasGap_r17                        `optional,ext-6`
	Nr_NeedForGapNCSG_Reporting_r17              *MeasAndMobParametersCommon_nr_NeedForGapNCSG_Reporting_r17              `optional,ext-6`
	Eutra_NeedForGapNCSG_Reporting_r17           *MeasAndMobParametersCommon_eutra_NeedForGapNCSG_Reporting_r17           `optional,ext-6`
	Ncsg_MeasGapPerFR_r17                        *MeasAndMobParametersCommon_ncsg_MeasGapPerFR_r17                        `optional,ext-6`
	Ncsg_MeasGapPatterns_r17                     *uper.BitString                                                          `lb:24,ub:24,optional,ext-6`
	Ncsg_MeasGapNR_Patterns_r17                  *uper.BitString                                                          `lb:24,ub:24,optional,ext-6`
	PreconfiguredUE_AutonomousMeasGap_r17        *MeasAndMobParametersCommon_preconfiguredUE_AutonomousMeasGap_r17        `optional,ext-6`
	PreconfiguredNW_ControlledMeasGap_r17        *MeasAndMobParametersCommon_preconfiguredNW_ControlledMeasGap_r17        `optional,ext-6`
	HandoverFR1_FR2_2_r17                        *MeasAndMobParametersCommon_handoverFR1_FR2_2_r17                        `optional,ext-6`
	HandoverFR2_1_FR2_2_r17                      *MeasAndMobParametersCommon_handoverFR2_1_FR2_2_r17                      `optional,ext-6`
	IndependentGapConfigPRS_r17                  *MeasAndMobParametersCommon_independentGapConfigPRS_r17                  `optional,ext-6`
	Rrm_RelaxationRRC_ConnectedRedCap_r17        *MeasAndMobParametersCommon_rrm_RelaxationRRC_ConnectedRedCap_r17        `optional,ext-6`
	ParallelMeasurementGap_r17                   *MeasAndMobParametersCommon_parallelMeasurementGap_r17                   `optional,ext-6`
	CondHandoverWithSCG_NRDC_r17                 *MeasAndMobParametersCommon_condHandoverWithSCG_NRDC_r17                 `optional,ext-6`
	GNB_ID_LengthReporting_r17                   *MeasAndMobParametersCommon_gNB_ID_LengthReporting_r17                   `optional,ext-6`
	GNB_ID_LengthReporting_ENDC_r17              *MeasAndMobParametersCommon_gNB_ID_LengthReporting_ENDC_r17              `optional,ext-6`
	GNB_ID_LengthReporting_NEDC_r17              *MeasAndMobParametersCommon_gNB_ID_LengthReporting_NEDC_r17              `optional,ext-6`
	GNB_ID_LengthReporting_NRDC_r17              *MeasAndMobParametersCommon_gNB_ID_LengthReporting_NRDC_r17              `optional,ext-6`
	GNB_ID_LengthReporting_NPN_r17               *MeasAndMobParametersCommon_gNB_ID_LengthReporting_NPN_r17               `optional,ext-6`
	ParallelSMTC_r17                             *MeasAndMobParametersCommon_parallelSMTC_r17                             `optional,ext-7`
	ConcurrentMeasGapEUTRA_r17                   *MeasAndMobParametersCommon_concurrentMeasGapEUTRA_r17                   `optional,ext-7`
	ServiceLinkPropDelayDiffReporting_r17        *MeasAndMobParametersCommon_serviceLinkPropDelayDiffReporting_r17        `optional,ext-7`
	Ncsg_SymbolLevelScheduleRestrictionInter_r17 *MeasAndMobParametersCommon_ncsg_SymbolLevelScheduleRestrictionInter_r17 `optional,ext-7`
	EventD1_MeasReportTrigger_r17                *MeasAndMobParametersCommon_eventD1_MeasReportTrigger_r17                `optional,ext-8`
	IndependentGapConfig_maxCC_r17               *MeasAndMobParametersCommon_independentGapConfig_maxCC_r17               `lb:1,ub:32,optional,ext-8`
}

func (ie *MeasAndMobParametersCommon) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.EventB_MeasAndReport != nil || ie.HandoverFDD_TDD != nil || ie.Eutra_CGI_Reporting != nil || ie.Nr_CGI_Reporting != nil || ie.IndependentGapConfig != nil || ie.PeriodicEUTRA_MeasAndReport != nil || ie.HandoverFR1_FR2 != nil || ie.MaxNumberCSI_RS_RRM_RS_SINR != nil || ie.Nr_CGI_Reporting_ENDC != nil || ie.Eutra_CGI_Reporting_NEDC != nil || ie.Eutra_CGI_Reporting_NRDC != nil || ie.Nr_CGI_Reporting_NEDC != nil || ie.Nr_CGI_Reporting_NRDC != nil || ie.ReportAddNeighMeasForPeriodic_r16 != nil || ie.CondHandoverParametersCommon_r16 != nil || ie.Nr_NeedForGap_Reporting_r16 != nil || ie.SupportedGapPattern_NRonly_r16 != nil || ie.SupportedGapPattern_NRonly_NEDC_r16 != nil || ie.MaxNumberCLI_RSSI_r16 != nil || ie.MaxNumberCLI_SRS_RSRP_r16 != nil || ie.MaxNumberPerSlotCLI_SRS_RSRP_r16 != nil || ie.Mfbi_IAB_r16 != nil || ie.Dummy != nil || ie.Nr_CGI_Reporting_NPN_r16 != nil || ie.IdleInactiveEUTRA_MeasReport_r16 != nil || ie.IdleInactive_ValidityArea_r16 != nil || ie.Eutra_AutonomousGaps_r16 != nil || ie.Eutra_AutonomousGaps_NEDC_r16 != nil || ie.Eutra_AutonomousGaps_NRDC_r16 != nil || ie.PcellT312_r16 != nil || ie.SupportedGapPattern_r16 != nil || ie.ConcurrentMeasGap_r17 != nil || ie.Nr_NeedForGapNCSG_Reporting_r17 != nil || ie.Eutra_NeedForGapNCSG_Reporting_r17 != nil || ie.Ncsg_MeasGapPerFR_r17 != nil || ie.Ncsg_MeasGapPatterns_r17 != nil || ie.Ncsg_MeasGapNR_Patterns_r17 != nil || ie.PreconfiguredUE_AutonomousMeasGap_r17 != nil || ie.PreconfiguredNW_ControlledMeasGap_r17 != nil || ie.HandoverFR1_FR2_2_r17 != nil || ie.HandoverFR2_1_FR2_2_r17 != nil || ie.IndependentGapConfigPRS_r17 != nil || ie.Rrm_RelaxationRRC_ConnectedRedCap_r17 != nil || ie.ParallelMeasurementGap_r17 != nil || ie.CondHandoverWithSCG_NRDC_r17 != nil || ie.GNB_ID_LengthReporting_r17 != nil || ie.GNB_ID_LengthReporting_ENDC_r17 != nil || ie.GNB_ID_LengthReporting_NEDC_r17 != nil || ie.GNB_ID_LengthReporting_NRDC_r17 != nil || ie.GNB_ID_LengthReporting_NPN_r17 != nil || ie.ParallelSMTC_r17 != nil || ie.ConcurrentMeasGapEUTRA_r17 != nil || ie.ServiceLinkPropDelayDiffReporting_r17 != nil || ie.Ncsg_SymbolLevelScheduleRestrictionInter_r17 != nil || ie.EventD1_MeasReportTrigger_r17 != nil || ie.IndependentGapConfig_maxCC_r17 != nil
	preambleBits := []bool{hasExtensions, ie.SupportedGapPattern != nil, ie.Ssb_RLM != nil, ie.Ssb_AndCSI_RS_RLM != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SupportedGapPattern != nil {
		if err = w.WriteBitString(ie.SupportedGapPattern.Bytes, uint(ie.SupportedGapPattern.NumBits), &uper.Constraint{Lb: 22, Ub: 22}, false); err != nil {
			return utils.WrapError("Encode SupportedGapPattern", err)
		}
	}
	if ie.Ssb_RLM != nil {
		if err = ie.Ssb_RLM.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_RLM", err)
		}
	}
	if ie.Ssb_AndCSI_RS_RLM != nil {
		if err = ie.Ssb_AndCSI_RS_RLM.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_AndCSI_RS_RLM", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 8 bits for 8 extension groups
		extBitmap := []bool{ie.EventB_MeasAndReport != nil || ie.HandoverFDD_TDD != nil || ie.Eutra_CGI_Reporting != nil || ie.Nr_CGI_Reporting != nil, ie.IndependentGapConfig != nil || ie.PeriodicEUTRA_MeasAndReport != nil || ie.HandoverFR1_FR2 != nil || ie.MaxNumberCSI_RS_RRM_RS_SINR != nil, ie.Nr_CGI_Reporting_ENDC != nil, ie.Eutra_CGI_Reporting_NEDC != nil || ie.Eutra_CGI_Reporting_NRDC != nil || ie.Nr_CGI_Reporting_NEDC != nil || ie.Nr_CGI_Reporting_NRDC != nil, ie.ReportAddNeighMeasForPeriodic_r16 != nil || ie.CondHandoverParametersCommon_r16 != nil || ie.Nr_NeedForGap_Reporting_r16 != nil || ie.SupportedGapPattern_NRonly_r16 != nil || ie.SupportedGapPattern_NRonly_NEDC_r16 != nil || ie.MaxNumberCLI_RSSI_r16 != nil || ie.MaxNumberCLI_SRS_RSRP_r16 != nil || ie.MaxNumberPerSlotCLI_SRS_RSRP_r16 != nil || ie.Mfbi_IAB_r16 != nil || ie.Dummy != nil || ie.Nr_CGI_Reporting_NPN_r16 != nil || ie.IdleInactiveEUTRA_MeasReport_r16 != nil || ie.IdleInactive_ValidityArea_r16 != nil || ie.Eutra_AutonomousGaps_r16 != nil || ie.Eutra_AutonomousGaps_NEDC_r16 != nil || ie.Eutra_AutonomousGaps_NRDC_r16 != nil || ie.PcellT312_r16 != nil || ie.SupportedGapPattern_r16 != nil, ie.ConcurrentMeasGap_r17 != nil || ie.Nr_NeedForGapNCSG_Reporting_r17 != nil || ie.Eutra_NeedForGapNCSG_Reporting_r17 != nil || ie.Ncsg_MeasGapPerFR_r17 != nil || ie.Ncsg_MeasGapPatterns_r17 != nil || ie.Ncsg_MeasGapNR_Patterns_r17 != nil || ie.PreconfiguredUE_AutonomousMeasGap_r17 != nil || ie.PreconfiguredNW_ControlledMeasGap_r17 != nil || ie.HandoverFR1_FR2_2_r17 != nil || ie.HandoverFR2_1_FR2_2_r17 != nil || ie.IndependentGapConfigPRS_r17 != nil || ie.Rrm_RelaxationRRC_ConnectedRedCap_r17 != nil || ie.ParallelMeasurementGap_r17 != nil || ie.CondHandoverWithSCG_NRDC_r17 != nil || ie.GNB_ID_LengthReporting_r17 != nil || ie.GNB_ID_LengthReporting_ENDC_r17 != nil || ie.GNB_ID_LengthReporting_NEDC_r17 != nil || ie.GNB_ID_LengthReporting_NRDC_r17 != nil || ie.GNB_ID_LengthReporting_NPN_r17 != nil, ie.ParallelSMTC_r17 != nil || ie.ConcurrentMeasGapEUTRA_r17 != nil || ie.ServiceLinkPropDelayDiffReporting_r17 != nil || ie.Ncsg_SymbolLevelScheduleRestrictionInter_r17 != nil, ie.EventD1_MeasReportTrigger_r17 != nil || ie.IndependentGapConfig_maxCC_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasAndMobParametersCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.EventB_MeasAndReport != nil, ie.HandoverFDD_TDD != nil, ie.Eutra_CGI_Reporting != nil, ie.Nr_CGI_Reporting != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode EventB_MeasAndReport optional
			if ie.EventB_MeasAndReport != nil {
				if err = ie.EventB_MeasAndReport.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EventB_MeasAndReport", err)
				}
			}
			// encode HandoverFDD_TDD optional
			if ie.HandoverFDD_TDD != nil {
				if err = ie.HandoverFDD_TDD.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HandoverFDD_TDD", err)
				}
			}
			// encode Eutra_CGI_Reporting optional
			if ie.Eutra_CGI_Reporting != nil {
				if err = ie.Eutra_CGI_Reporting.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Eutra_CGI_Reporting", err)
				}
			}
			// encode Nr_CGI_Reporting optional
			if ie.Nr_CGI_Reporting != nil {
				if err = ie.Nr_CGI_Reporting.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Nr_CGI_Reporting", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.IndependentGapConfig != nil, ie.PeriodicEUTRA_MeasAndReport != nil, ie.HandoverFR1_FR2 != nil, ie.MaxNumberCSI_RS_RRM_RS_SINR != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode IndependentGapConfig optional
			if ie.IndependentGapConfig != nil {
				if err = ie.IndependentGapConfig.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IndependentGapConfig", err)
				}
			}
			// encode PeriodicEUTRA_MeasAndReport optional
			if ie.PeriodicEUTRA_MeasAndReport != nil {
				if err = ie.PeriodicEUTRA_MeasAndReport.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PeriodicEUTRA_MeasAndReport", err)
				}
			}
			// encode HandoverFR1_FR2 optional
			if ie.HandoverFR1_FR2 != nil {
				if err = ie.HandoverFR1_FR2.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HandoverFR1_FR2", err)
				}
			}
			// encode MaxNumberCSI_RS_RRM_RS_SINR optional
			if ie.MaxNumberCSI_RS_RRM_RS_SINR != nil {
				if err = ie.MaxNumberCSI_RS_RRM_RS_SINR.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxNumberCSI_RS_RRM_RS_SINR", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.Nr_CGI_Reporting_ENDC != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Nr_CGI_Reporting_ENDC optional
			if ie.Nr_CGI_Reporting_ENDC != nil {
				if err = ie.Nr_CGI_Reporting_ENDC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Nr_CGI_Reporting_ENDC", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 4
		if extBitmap[3] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{ie.Eutra_CGI_Reporting_NEDC != nil, ie.Eutra_CGI_Reporting_NRDC != nil, ie.Nr_CGI_Reporting_NEDC != nil, ie.Nr_CGI_Reporting_NRDC != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Eutra_CGI_Reporting_NEDC optional
			if ie.Eutra_CGI_Reporting_NEDC != nil {
				if err = ie.Eutra_CGI_Reporting_NEDC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Eutra_CGI_Reporting_NEDC", err)
				}
			}
			// encode Eutra_CGI_Reporting_NRDC optional
			if ie.Eutra_CGI_Reporting_NRDC != nil {
				if err = ie.Eutra_CGI_Reporting_NRDC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Eutra_CGI_Reporting_NRDC", err)
				}
			}
			// encode Nr_CGI_Reporting_NEDC optional
			if ie.Nr_CGI_Reporting_NEDC != nil {
				if err = ie.Nr_CGI_Reporting_NEDC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Nr_CGI_Reporting_NEDC", err)
				}
			}
			// encode Nr_CGI_Reporting_NRDC optional
			if ie.Nr_CGI_Reporting_NRDC != nil {
				if err = ie.Nr_CGI_Reporting_NRDC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Nr_CGI_Reporting_NRDC", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 5
		if extBitmap[4] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 5
			optionals_ext_5 := []bool{ie.ReportAddNeighMeasForPeriodic_r16 != nil, ie.CondHandoverParametersCommon_r16 != nil, ie.Nr_NeedForGap_Reporting_r16 != nil, ie.SupportedGapPattern_NRonly_r16 != nil, ie.SupportedGapPattern_NRonly_NEDC_r16 != nil, ie.MaxNumberCLI_RSSI_r16 != nil, ie.MaxNumberCLI_SRS_RSRP_r16 != nil, ie.MaxNumberPerSlotCLI_SRS_RSRP_r16 != nil, ie.Mfbi_IAB_r16 != nil, ie.Dummy != nil, ie.Nr_CGI_Reporting_NPN_r16 != nil, ie.IdleInactiveEUTRA_MeasReport_r16 != nil, ie.IdleInactive_ValidityArea_r16 != nil, ie.Eutra_AutonomousGaps_r16 != nil, ie.Eutra_AutonomousGaps_NEDC_r16 != nil, ie.Eutra_AutonomousGaps_NRDC_r16 != nil, ie.PcellT312_r16 != nil, ie.SupportedGapPattern_r16 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ReportAddNeighMeasForPeriodic_r16 optional
			if ie.ReportAddNeighMeasForPeriodic_r16 != nil {
				if err = ie.ReportAddNeighMeasForPeriodic_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ReportAddNeighMeasForPeriodic_r16", err)
				}
			}
			// encode CondHandoverParametersCommon_r16 optional
			if ie.CondHandoverParametersCommon_r16 != nil {
				if err = ie.CondHandoverParametersCommon_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CondHandoverParametersCommon_r16", err)
				}
			}
			// encode Nr_NeedForGap_Reporting_r16 optional
			if ie.Nr_NeedForGap_Reporting_r16 != nil {
				if err = ie.Nr_NeedForGap_Reporting_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Nr_NeedForGap_Reporting_r16", err)
				}
			}
			// encode SupportedGapPattern_NRonly_r16 optional
			if ie.SupportedGapPattern_NRonly_r16 != nil {
				if err = extWriter.WriteBitString(ie.SupportedGapPattern_NRonly_r16.Bytes, uint(ie.SupportedGapPattern_NRonly_r16.NumBits), &uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
					return utils.WrapError("Encode SupportedGapPattern_NRonly_r16", err)
				}
			}
			// encode SupportedGapPattern_NRonly_NEDC_r16 optional
			if ie.SupportedGapPattern_NRonly_NEDC_r16 != nil {
				if err = ie.SupportedGapPattern_NRonly_NEDC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SupportedGapPattern_NRonly_NEDC_r16", err)
				}
			}
			// encode MaxNumberCLI_RSSI_r16 optional
			if ie.MaxNumberCLI_RSSI_r16 != nil {
				if err = ie.MaxNumberCLI_RSSI_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxNumberCLI_RSSI_r16", err)
				}
			}
			// encode MaxNumberCLI_SRS_RSRP_r16 optional
			if ie.MaxNumberCLI_SRS_RSRP_r16 != nil {
				if err = ie.MaxNumberCLI_SRS_RSRP_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxNumberCLI_SRS_RSRP_r16", err)
				}
			}
			// encode MaxNumberPerSlotCLI_SRS_RSRP_r16 optional
			if ie.MaxNumberPerSlotCLI_SRS_RSRP_r16 != nil {
				if err = ie.MaxNumberPerSlotCLI_SRS_RSRP_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxNumberPerSlotCLI_SRS_RSRP_r16", err)
				}
			}
			// encode Mfbi_IAB_r16 optional
			if ie.Mfbi_IAB_r16 != nil {
				if err = ie.Mfbi_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Mfbi_IAB_r16", err)
				}
			}
			// encode Dummy optional
			if ie.Dummy != nil {
				if err = ie.Dummy.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dummy", err)
				}
			}
			// encode Nr_CGI_Reporting_NPN_r16 optional
			if ie.Nr_CGI_Reporting_NPN_r16 != nil {
				if err = ie.Nr_CGI_Reporting_NPN_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Nr_CGI_Reporting_NPN_r16", err)
				}
			}
			// encode IdleInactiveEUTRA_MeasReport_r16 optional
			if ie.IdleInactiveEUTRA_MeasReport_r16 != nil {
				if err = ie.IdleInactiveEUTRA_MeasReport_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IdleInactiveEUTRA_MeasReport_r16", err)
				}
			}
			// encode IdleInactive_ValidityArea_r16 optional
			if ie.IdleInactive_ValidityArea_r16 != nil {
				if err = ie.IdleInactive_ValidityArea_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IdleInactive_ValidityArea_r16", err)
				}
			}
			// encode Eutra_AutonomousGaps_r16 optional
			if ie.Eutra_AutonomousGaps_r16 != nil {
				if err = ie.Eutra_AutonomousGaps_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Eutra_AutonomousGaps_r16", err)
				}
			}
			// encode Eutra_AutonomousGaps_NEDC_r16 optional
			if ie.Eutra_AutonomousGaps_NEDC_r16 != nil {
				if err = ie.Eutra_AutonomousGaps_NEDC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Eutra_AutonomousGaps_NEDC_r16", err)
				}
			}
			// encode Eutra_AutonomousGaps_NRDC_r16 optional
			if ie.Eutra_AutonomousGaps_NRDC_r16 != nil {
				if err = ie.Eutra_AutonomousGaps_NRDC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Eutra_AutonomousGaps_NRDC_r16", err)
				}
			}
			// encode PcellT312_r16 optional
			if ie.PcellT312_r16 != nil {
				if err = ie.PcellT312_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PcellT312_r16", err)
				}
			}
			// encode SupportedGapPattern_r16 optional
			if ie.SupportedGapPattern_r16 != nil {
				if err = extWriter.WriteBitString(ie.SupportedGapPattern_r16.Bytes, uint(ie.SupportedGapPattern_r16.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
					return utils.WrapError("Encode SupportedGapPattern_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 6
		if extBitmap[5] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 6
			optionals_ext_6 := []bool{ie.ConcurrentMeasGap_r17 != nil, ie.Nr_NeedForGapNCSG_Reporting_r17 != nil, ie.Eutra_NeedForGapNCSG_Reporting_r17 != nil, ie.Ncsg_MeasGapPerFR_r17 != nil, ie.Ncsg_MeasGapPatterns_r17 != nil, ie.Ncsg_MeasGapNR_Patterns_r17 != nil, ie.PreconfiguredUE_AutonomousMeasGap_r17 != nil, ie.PreconfiguredNW_ControlledMeasGap_r17 != nil, ie.HandoverFR1_FR2_2_r17 != nil, ie.HandoverFR2_1_FR2_2_r17 != nil, ie.IndependentGapConfigPRS_r17 != nil, ie.Rrm_RelaxationRRC_ConnectedRedCap_r17 != nil, ie.ParallelMeasurementGap_r17 != nil, ie.CondHandoverWithSCG_NRDC_r17 != nil, ie.GNB_ID_LengthReporting_r17 != nil, ie.GNB_ID_LengthReporting_ENDC_r17 != nil, ie.GNB_ID_LengthReporting_NEDC_r17 != nil, ie.GNB_ID_LengthReporting_NRDC_r17 != nil, ie.GNB_ID_LengthReporting_NPN_r17 != nil}
			for _, bit := range optionals_ext_6 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ConcurrentMeasGap_r17 optional
			if ie.ConcurrentMeasGap_r17 != nil {
				if err = ie.ConcurrentMeasGap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ConcurrentMeasGap_r17", err)
				}
			}
			// encode Nr_NeedForGapNCSG_Reporting_r17 optional
			if ie.Nr_NeedForGapNCSG_Reporting_r17 != nil {
				if err = ie.Nr_NeedForGapNCSG_Reporting_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Nr_NeedForGapNCSG_Reporting_r17", err)
				}
			}
			// encode Eutra_NeedForGapNCSG_Reporting_r17 optional
			if ie.Eutra_NeedForGapNCSG_Reporting_r17 != nil {
				if err = ie.Eutra_NeedForGapNCSG_Reporting_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Eutra_NeedForGapNCSG_Reporting_r17", err)
				}
			}
			// encode Ncsg_MeasGapPerFR_r17 optional
			if ie.Ncsg_MeasGapPerFR_r17 != nil {
				if err = ie.Ncsg_MeasGapPerFR_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ncsg_MeasGapPerFR_r17", err)
				}
			}
			// encode Ncsg_MeasGapPatterns_r17 optional
			if ie.Ncsg_MeasGapPatterns_r17 != nil {
				if err = extWriter.WriteBitString(ie.Ncsg_MeasGapPatterns_r17.Bytes, uint(ie.Ncsg_MeasGapPatterns_r17.NumBits), &uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
					return utils.WrapError("Encode Ncsg_MeasGapPatterns_r17", err)
				}
			}
			// encode Ncsg_MeasGapNR_Patterns_r17 optional
			if ie.Ncsg_MeasGapNR_Patterns_r17 != nil {
				if err = extWriter.WriteBitString(ie.Ncsg_MeasGapNR_Patterns_r17.Bytes, uint(ie.Ncsg_MeasGapNR_Patterns_r17.NumBits), &uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
					return utils.WrapError("Encode Ncsg_MeasGapNR_Patterns_r17", err)
				}
			}
			// encode PreconfiguredUE_AutonomousMeasGap_r17 optional
			if ie.PreconfiguredUE_AutonomousMeasGap_r17 != nil {
				if err = ie.PreconfiguredUE_AutonomousMeasGap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PreconfiguredUE_AutonomousMeasGap_r17", err)
				}
			}
			// encode PreconfiguredNW_ControlledMeasGap_r17 optional
			if ie.PreconfiguredNW_ControlledMeasGap_r17 != nil {
				if err = ie.PreconfiguredNW_ControlledMeasGap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PreconfiguredNW_ControlledMeasGap_r17", err)
				}
			}
			// encode HandoverFR1_FR2_2_r17 optional
			if ie.HandoverFR1_FR2_2_r17 != nil {
				if err = ie.HandoverFR1_FR2_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HandoverFR1_FR2_2_r17", err)
				}
			}
			// encode HandoverFR2_1_FR2_2_r17 optional
			if ie.HandoverFR2_1_FR2_2_r17 != nil {
				if err = ie.HandoverFR2_1_FR2_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HandoverFR2_1_FR2_2_r17", err)
				}
			}
			// encode IndependentGapConfigPRS_r17 optional
			if ie.IndependentGapConfigPRS_r17 != nil {
				if err = ie.IndependentGapConfigPRS_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IndependentGapConfigPRS_r17", err)
				}
			}
			// encode Rrm_RelaxationRRC_ConnectedRedCap_r17 optional
			if ie.Rrm_RelaxationRRC_ConnectedRedCap_r17 != nil {
				if err = ie.Rrm_RelaxationRRC_ConnectedRedCap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Rrm_RelaxationRRC_ConnectedRedCap_r17", err)
				}
			}
			// encode ParallelMeasurementGap_r17 optional
			if ie.ParallelMeasurementGap_r17 != nil {
				if err = ie.ParallelMeasurementGap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ParallelMeasurementGap_r17", err)
				}
			}
			// encode CondHandoverWithSCG_NRDC_r17 optional
			if ie.CondHandoverWithSCG_NRDC_r17 != nil {
				if err = ie.CondHandoverWithSCG_NRDC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CondHandoverWithSCG_NRDC_r17", err)
				}
			}
			// encode GNB_ID_LengthReporting_r17 optional
			if ie.GNB_ID_LengthReporting_r17 != nil {
				if err = ie.GNB_ID_LengthReporting_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode GNB_ID_LengthReporting_r17", err)
				}
			}
			// encode GNB_ID_LengthReporting_ENDC_r17 optional
			if ie.GNB_ID_LengthReporting_ENDC_r17 != nil {
				if err = ie.GNB_ID_LengthReporting_ENDC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode GNB_ID_LengthReporting_ENDC_r17", err)
				}
			}
			// encode GNB_ID_LengthReporting_NEDC_r17 optional
			if ie.GNB_ID_LengthReporting_NEDC_r17 != nil {
				if err = ie.GNB_ID_LengthReporting_NEDC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode GNB_ID_LengthReporting_NEDC_r17", err)
				}
			}
			// encode GNB_ID_LengthReporting_NRDC_r17 optional
			if ie.GNB_ID_LengthReporting_NRDC_r17 != nil {
				if err = ie.GNB_ID_LengthReporting_NRDC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode GNB_ID_LengthReporting_NRDC_r17", err)
				}
			}
			// encode GNB_ID_LengthReporting_NPN_r17 optional
			if ie.GNB_ID_LengthReporting_NPN_r17 != nil {
				if err = ie.GNB_ID_LengthReporting_NPN_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode GNB_ID_LengthReporting_NPN_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 7
		if extBitmap[6] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 7
			optionals_ext_7 := []bool{ie.ParallelSMTC_r17 != nil, ie.ConcurrentMeasGapEUTRA_r17 != nil, ie.ServiceLinkPropDelayDiffReporting_r17 != nil, ie.Ncsg_SymbolLevelScheduleRestrictionInter_r17 != nil}
			for _, bit := range optionals_ext_7 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ParallelSMTC_r17 optional
			if ie.ParallelSMTC_r17 != nil {
				if err = ie.ParallelSMTC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ParallelSMTC_r17", err)
				}
			}
			// encode ConcurrentMeasGapEUTRA_r17 optional
			if ie.ConcurrentMeasGapEUTRA_r17 != nil {
				if err = ie.ConcurrentMeasGapEUTRA_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ConcurrentMeasGapEUTRA_r17", err)
				}
			}
			// encode ServiceLinkPropDelayDiffReporting_r17 optional
			if ie.ServiceLinkPropDelayDiffReporting_r17 != nil {
				if err = ie.ServiceLinkPropDelayDiffReporting_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ServiceLinkPropDelayDiffReporting_r17", err)
				}
			}
			// encode Ncsg_SymbolLevelScheduleRestrictionInter_r17 optional
			if ie.Ncsg_SymbolLevelScheduleRestrictionInter_r17 != nil {
				if err = ie.Ncsg_SymbolLevelScheduleRestrictionInter_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ncsg_SymbolLevelScheduleRestrictionInter_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 8
		if extBitmap[7] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 8
			optionals_ext_8 := []bool{ie.EventD1_MeasReportTrigger_r17 != nil, ie.IndependentGapConfig_maxCC_r17 != nil}
			for _, bit := range optionals_ext_8 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode EventD1_MeasReportTrigger_r17 optional
			if ie.EventD1_MeasReportTrigger_r17 != nil {
				if err = ie.EventD1_MeasReportTrigger_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode EventD1_MeasReportTrigger_r17", err)
				}
			}
			// encode IndependentGapConfig_maxCC_r17 optional
			if ie.IndependentGapConfig_maxCC_r17 != nil {
				if err = ie.IndependentGapConfig_maxCC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IndependentGapConfig_maxCC_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *MeasAndMobParametersCommon) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedGapPatternPresent bool
	if SupportedGapPatternPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_RLMPresent bool
	if Ssb_RLMPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_AndCSI_RS_RLMPresent bool
	if Ssb_AndCSI_RS_RLMPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SupportedGapPatternPresent {
		var tmp_bs_SupportedGapPattern []byte
		var n_SupportedGapPattern uint
		if tmp_bs_SupportedGapPattern, n_SupportedGapPattern, err = r.ReadBitString(&uper.Constraint{Lb: 22, Ub: 22}, false); err != nil {
			return utils.WrapError("Decode SupportedGapPattern", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_SupportedGapPattern,
			NumBits: uint64(n_SupportedGapPattern),
		}
		ie.SupportedGapPattern = &tmp_bitstring
	}
	if Ssb_RLMPresent {
		ie.Ssb_RLM = new(MeasAndMobParametersCommon_ssb_RLM)
		if err = ie.Ssb_RLM.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_RLM", err)
		}
	}
	if Ssb_AndCSI_RS_RLMPresent {
		ie.Ssb_AndCSI_RS_RLM = new(MeasAndMobParametersCommon_ssb_AndCSI_RS_RLM)
		if err = ie.Ssb_AndCSI_RS_RLM.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_AndCSI_RS_RLM", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 8 bits for 8 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			EventB_MeasAndReportPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			HandoverFDD_TDDPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Eutra_CGI_ReportingPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Nr_CGI_ReportingPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode EventB_MeasAndReport optional
			if EventB_MeasAndReportPresent {
				ie.EventB_MeasAndReport = new(MeasAndMobParametersCommon_eventB_MeasAndReport)
				if err = ie.EventB_MeasAndReport.Decode(extReader); err != nil {
					return utils.WrapError("Decode EventB_MeasAndReport", err)
				}
			}
			// decode HandoverFDD_TDD optional
			if HandoverFDD_TDDPresent {
				ie.HandoverFDD_TDD = new(MeasAndMobParametersCommon_handoverFDD_TDD)
				if err = ie.HandoverFDD_TDD.Decode(extReader); err != nil {
					return utils.WrapError("Decode HandoverFDD_TDD", err)
				}
			}
			// decode Eutra_CGI_Reporting optional
			if Eutra_CGI_ReportingPresent {
				ie.Eutra_CGI_Reporting = new(MeasAndMobParametersCommon_eutra_CGI_Reporting)
				if err = ie.Eutra_CGI_Reporting.Decode(extReader); err != nil {
					return utils.WrapError("Decode Eutra_CGI_Reporting", err)
				}
			}
			// decode Nr_CGI_Reporting optional
			if Nr_CGI_ReportingPresent {
				ie.Nr_CGI_Reporting = new(MeasAndMobParametersCommon_nr_CGI_Reporting)
				if err = ie.Nr_CGI_Reporting.Decode(extReader); err != nil {
					return utils.WrapError("Decode Nr_CGI_Reporting", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			IndependentGapConfigPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PeriodicEUTRA_MeasAndReportPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			HandoverFR1_FR2Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxNumberCSI_RS_RRM_RS_SINRPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode IndependentGapConfig optional
			if IndependentGapConfigPresent {
				ie.IndependentGapConfig = new(MeasAndMobParametersCommon_independentGapConfig)
				if err = ie.IndependentGapConfig.Decode(extReader); err != nil {
					return utils.WrapError("Decode IndependentGapConfig", err)
				}
			}
			// decode PeriodicEUTRA_MeasAndReport optional
			if PeriodicEUTRA_MeasAndReportPresent {
				ie.PeriodicEUTRA_MeasAndReport = new(MeasAndMobParametersCommon_periodicEUTRA_MeasAndReport)
				if err = ie.PeriodicEUTRA_MeasAndReport.Decode(extReader); err != nil {
					return utils.WrapError("Decode PeriodicEUTRA_MeasAndReport", err)
				}
			}
			// decode HandoverFR1_FR2 optional
			if HandoverFR1_FR2Present {
				ie.HandoverFR1_FR2 = new(MeasAndMobParametersCommon_handoverFR1_FR2)
				if err = ie.HandoverFR1_FR2.Decode(extReader); err != nil {
					return utils.WrapError("Decode HandoverFR1_FR2", err)
				}
			}
			// decode MaxNumberCSI_RS_RRM_RS_SINR optional
			if MaxNumberCSI_RS_RRM_RS_SINRPresent {
				ie.MaxNumberCSI_RS_RRM_RS_SINR = new(MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR)
				if err = ie.MaxNumberCSI_RS_RRM_RS_SINR.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxNumberCSI_RS_RRM_RS_SINR", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Nr_CGI_Reporting_ENDCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Nr_CGI_Reporting_ENDC optional
			if Nr_CGI_Reporting_ENDCPresent {
				ie.Nr_CGI_Reporting_ENDC = new(MeasAndMobParametersCommon_nr_CGI_Reporting_ENDC)
				if err = ie.Nr_CGI_Reporting_ENDC.Decode(extReader); err != nil {
					return utils.WrapError("Decode Nr_CGI_Reporting_ENDC", err)
				}
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Eutra_CGI_Reporting_NEDCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Eutra_CGI_Reporting_NRDCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Nr_CGI_Reporting_NEDCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Nr_CGI_Reporting_NRDCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Eutra_CGI_Reporting_NEDC optional
			if Eutra_CGI_Reporting_NEDCPresent {
				ie.Eutra_CGI_Reporting_NEDC = new(MeasAndMobParametersCommon_eutra_CGI_Reporting_NEDC)
				if err = ie.Eutra_CGI_Reporting_NEDC.Decode(extReader); err != nil {
					return utils.WrapError("Decode Eutra_CGI_Reporting_NEDC", err)
				}
			}
			// decode Eutra_CGI_Reporting_NRDC optional
			if Eutra_CGI_Reporting_NRDCPresent {
				ie.Eutra_CGI_Reporting_NRDC = new(MeasAndMobParametersCommon_eutra_CGI_Reporting_NRDC)
				if err = ie.Eutra_CGI_Reporting_NRDC.Decode(extReader); err != nil {
					return utils.WrapError("Decode Eutra_CGI_Reporting_NRDC", err)
				}
			}
			// decode Nr_CGI_Reporting_NEDC optional
			if Nr_CGI_Reporting_NEDCPresent {
				ie.Nr_CGI_Reporting_NEDC = new(MeasAndMobParametersCommon_nr_CGI_Reporting_NEDC)
				if err = ie.Nr_CGI_Reporting_NEDC.Decode(extReader); err != nil {
					return utils.WrapError("Decode Nr_CGI_Reporting_NEDC", err)
				}
			}
			// decode Nr_CGI_Reporting_NRDC optional
			if Nr_CGI_Reporting_NRDCPresent {
				ie.Nr_CGI_Reporting_NRDC = new(MeasAndMobParametersCommon_nr_CGI_Reporting_NRDC)
				if err = ie.Nr_CGI_Reporting_NRDC.Decode(extReader); err != nil {
					return utils.WrapError("Decode Nr_CGI_Reporting_NRDC", err)
				}
			}
		}
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ReportAddNeighMeasForPeriodic_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CondHandoverParametersCommon_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Nr_NeedForGap_Reporting_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportedGapPattern_NRonly_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportedGapPattern_NRonly_NEDC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxNumberCLI_RSSI_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxNumberCLI_SRS_RSRP_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxNumberPerSlotCLI_SRS_RSRP_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Mfbi_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DummyPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Nr_CGI_Reporting_NPN_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			IdleInactiveEUTRA_MeasReport_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			IdleInactive_ValidityArea_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Eutra_AutonomousGaps_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Eutra_AutonomousGaps_NEDC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Eutra_AutonomousGaps_NRDC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PcellT312_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SupportedGapPattern_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ReportAddNeighMeasForPeriodic_r16 optional
			if ReportAddNeighMeasForPeriodic_r16Present {
				ie.ReportAddNeighMeasForPeriodic_r16 = new(MeasAndMobParametersCommon_reportAddNeighMeasForPeriodic_r16)
				if err = ie.ReportAddNeighMeasForPeriodic_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ReportAddNeighMeasForPeriodic_r16", err)
				}
			}
			// decode CondHandoverParametersCommon_r16 optional
			if CondHandoverParametersCommon_r16Present {
				ie.CondHandoverParametersCommon_r16 = new(MeasAndMobParametersCommon_condHandoverParametersCommon_r16)
				if err = ie.CondHandoverParametersCommon_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode CondHandoverParametersCommon_r16", err)
				}
			}
			// decode Nr_NeedForGap_Reporting_r16 optional
			if Nr_NeedForGap_Reporting_r16Present {
				ie.Nr_NeedForGap_Reporting_r16 = new(MeasAndMobParametersCommon_nr_NeedForGap_Reporting_r16)
				if err = ie.Nr_NeedForGap_Reporting_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Nr_NeedForGap_Reporting_r16", err)
				}
			}
			// decode SupportedGapPattern_NRonly_r16 optional
			if SupportedGapPattern_NRonly_r16Present {
				var tmp_bs_SupportedGapPattern_NRonly_r16 []byte
				var n_SupportedGapPattern_NRonly_r16 uint
				if tmp_bs_SupportedGapPattern_NRonly_r16, n_SupportedGapPattern_NRonly_r16, err = extReader.ReadBitString(&uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
					return utils.WrapError("Decode SupportedGapPattern_NRonly_r16", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_SupportedGapPattern_NRonly_r16,
					NumBits: uint64(n_SupportedGapPattern_NRonly_r16),
				}
				ie.SupportedGapPattern_NRonly_r16 = &tmp_bitstring
			}
			// decode SupportedGapPattern_NRonly_NEDC_r16 optional
			if SupportedGapPattern_NRonly_NEDC_r16Present {
				ie.SupportedGapPattern_NRonly_NEDC_r16 = new(MeasAndMobParametersCommon_supportedGapPattern_NRonly_NEDC_r16)
				if err = ie.SupportedGapPattern_NRonly_NEDC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SupportedGapPattern_NRonly_NEDC_r16", err)
				}
			}
			// decode MaxNumberCLI_RSSI_r16 optional
			if MaxNumberCLI_RSSI_r16Present {
				ie.MaxNumberCLI_RSSI_r16 = new(MeasAndMobParametersCommon_maxNumberCLI_RSSI_r16)
				if err = ie.MaxNumberCLI_RSSI_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxNumberCLI_RSSI_r16", err)
				}
			}
			// decode MaxNumberCLI_SRS_RSRP_r16 optional
			if MaxNumberCLI_SRS_RSRP_r16Present {
				ie.MaxNumberCLI_SRS_RSRP_r16 = new(MeasAndMobParametersCommon_maxNumberCLI_SRS_RSRP_r16)
				if err = ie.MaxNumberCLI_SRS_RSRP_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxNumberCLI_SRS_RSRP_r16", err)
				}
			}
			// decode MaxNumberPerSlotCLI_SRS_RSRP_r16 optional
			if MaxNumberPerSlotCLI_SRS_RSRP_r16Present {
				ie.MaxNumberPerSlotCLI_SRS_RSRP_r16 = new(MeasAndMobParametersCommon_maxNumberPerSlotCLI_SRS_RSRP_r16)
				if err = ie.MaxNumberPerSlotCLI_SRS_RSRP_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxNumberPerSlotCLI_SRS_RSRP_r16", err)
				}
			}
			// decode Mfbi_IAB_r16 optional
			if Mfbi_IAB_r16Present {
				ie.Mfbi_IAB_r16 = new(MeasAndMobParametersCommon_mfbi_IAB_r16)
				if err = ie.Mfbi_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Mfbi_IAB_r16", err)
				}
			}
			// decode Dummy optional
			if DummyPresent {
				ie.Dummy = new(MeasAndMobParametersCommon_dummy)
				if err = ie.Dummy.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dummy", err)
				}
			}
			// decode Nr_CGI_Reporting_NPN_r16 optional
			if Nr_CGI_Reporting_NPN_r16Present {
				ie.Nr_CGI_Reporting_NPN_r16 = new(MeasAndMobParametersCommon_nr_CGI_Reporting_NPN_r16)
				if err = ie.Nr_CGI_Reporting_NPN_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Nr_CGI_Reporting_NPN_r16", err)
				}
			}
			// decode IdleInactiveEUTRA_MeasReport_r16 optional
			if IdleInactiveEUTRA_MeasReport_r16Present {
				ie.IdleInactiveEUTRA_MeasReport_r16 = new(MeasAndMobParametersCommon_idleInactiveEUTRA_MeasReport_r16)
				if err = ie.IdleInactiveEUTRA_MeasReport_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode IdleInactiveEUTRA_MeasReport_r16", err)
				}
			}
			// decode IdleInactive_ValidityArea_r16 optional
			if IdleInactive_ValidityArea_r16Present {
				ie.IdleInactive_ValidityArea_r16 = new(MeasAndMobParametersCommon_idleInactive_ValidityArea_r16)
				if err = ie.IdleInactive_ValidityArea_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode IdleInactive_ValidityArea_r16", err)
				}
			}
			// decode Eutra_AutonomousGaps_r16 optional
			if Eutra_AutonomousGaps_r16Present {
				ie.Eutra_AutonomousGaps_r16 = new(MeasAndMobParametersCommon_eutra_AutonomousGaps_r16)
				if err = ie.Eutra_AutonomousGaps_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Eutra_AutonomousGaps_r16", err)
				}
			}
			// decode Eutra_AutonomousGaps_NEDC_r16 optional
			if Eutra_AutonomousGaps_NEDC_r16Present {
				ie.Eutra_AutonomousGaps_NEDC_r16 = new(MeasAndMobParametersCommon_eutra_AutonomousGaps_NEDC_r16)
				if err = ie.Eutra_AutonomousGaps_NEDC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Eutra_AutonomousGaps_NEDC_r16", err)
				}
			}
			// decode Eutra_AutonomousGaps_NRDC_r16 optional
			if Eutra_AutonomousGaps_NRDC_r16Present {
				ie.Eutra_AutonomousGaps_NRDC_r16 = new(MeasAndMobParametersCommon_eutra_AutonomousGaps_NRDC_r16)
				if err = ie.Eutra_AutonomousGaps_NRDC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Eutra_AutonomousGaps_NRDC_r16", err)
				}
			}
			// decode PcellT312_r16 optional
			if PcellT312_r16Present {
				ie.PcellT312_r16 = new(MeasAndMobParametersCommon_pcellT312_r16)
				if err = ie.PcellT312_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode PcellT312_r16", err)
				}
			}
			// decode SupportedGapPattern_r16 optional
			if SupportedGapPattern_r16Present {
				var tmp_bs_SupportedGapPattern_r16 []byte
				var n_SupportedGapPattern_r16 uint
				if tmp_bs_SupportedGapPattern_r16, n_SupportedGapPattern_r16, err = extReader.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
					return utils.WrapError("Decode SupportedGapPattern_r16", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_SupportedGapPattern_r16,
					NumBits: uint64(n_SupportedGapPattern_r16),
				}
				ie.SupportedGapPattern_r16 = &tmp_bitstring
			}
		}
		// decode extension group 6
		if len(extBitmap) > 5 && extBitmap[5] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ConcurrentMeasGap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Nr_NeedForGapNCSG_Reporting_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Eutra_NeedForGapNCSG_Reporting_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ncsg_MeasGapPerFR_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ncsg_MeasGapPatterns_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ncsg_MeasGapNR_Patterns_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PreconfiguredUE_AutonomousMeasGap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			PreconfiguredNW_ControlledMeasGap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			HandoverFR1_FR2_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			HandoverFR2_1_FR2_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			IndependentGapConfigPRS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Rrm_RelaxationRRC_ConnectedRedCap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ParallelMeasurementGap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CondHandoverWithSCG_NRDC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			GNB_ID_LengthReporting_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			GNB_ID_LengthReporting_ENDC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			GNB_ID_LengthReporting_NEDC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			GNB_ID_LengthReporting_NRDC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			GNB_ID_LengthReporting_NPN_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ConcurrentMeasGap_r17 optional
			if ConcurrentMeasGap_r17Present {
				ie.ConcurrentMeasGap_r17 = new(MeasAndMobParametersCommon_concurrentMeasGap_r17)
				if err = ie.ConcurrentMeasGap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ConcurrentMeasGap_r17", err)
				}
			}
			// decode Nr_NeedForGapNCSG_Reporting_r17 optional
			if Nr_NeedForGapNCSG_Reporting_r17Present {
				ie.Nr_NeedForGapNCSG_Reporting_r17 = new(MeasAndMobParametersCommon_nr_NeedForGapNCSG_Reporting_r17)
				if err = ie.Nr_NeedForGapNCSG_Reporting_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Nr_NeedForGapNCSG_Reporting_r17", err)
				}
			}
			// decode Eutra_NeedForGapNCSG_Reporting_r17 optional
			if Eutra_NeedForGapNCSG_Reporting_r17Present {
				ie.Eutra_NeedForGapNCSG_Reporting_r17 = new(MeasAndMobParametersCommon_eutra_NeedForGapNCSG_Reporting_r17)
				if err = ie.Eutra_NeedForGapNCSG_Reporting_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Eutra_NeedForGapNCSG_Reporting_r17", err)
				}
			}
			// decode Ncsg_MeasGapPerFR_r17 optional
			if Ncsg_MeasGapPerFR_r17Present {
				ie.Ncsg_MeasGapPerFR_r17 = new(MeasAndMobParametersCommon_ncsg_MeasGapPerFR_r17)
				if err = ie.Ncsg_MeasGapPerFR_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ncsg_MeasGapPerFR_r17", err)
				}
			}
			// decode Ncsg_MeasGapPatterns_r17 optional
			if Ncsg_MeasGapPatterns_r17Present {
				var tmp_bs_Ncsg_MeasGapPatterns_r17 []byte
				var n_Ncsg_MeasGapPatterns_r17 uint
				if tmp_bs_Ncsg_MeasGapPatterns_r17, n_Ncsg_MeasGapPatterns_r17, err = extReader.ReadBitString(&uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
					return utils.WrapError("Decode Ncsg_MeasGapPatterns_r17", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_Ncsg_MeasGapPatterns_r17,
					NumBits: uint64(n_Ncsg_MeasGapPatterns_r17),
				}
				ie.Ncsg_MeasGapPatterns_r17 = &tmp_bitstring
			}
			// decode Ncsg_MeasGapNR_Patterns_r17 optional
			if Ncsg_MeasGapNR_Patterns_r17Present {
				var tmp_bs_Ncsg_MeasGapNR_Patterns_r17 []byte
				var n_Ncsg_MeasGapNR_Patterns_r17 uint
				if tmp_bs_Ncsg_MeasGapNR_Patterns_r17, n_Ncsg_MeasGapNR_Patterns_r17, err = extReader.ReadBitString(&uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
					return utils.WrapError("Decode Ncsg_MeasGapNR_Patterns_r17", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_Ncsg_MeasGapNR_Patterns_r17,
					NumBits: uint64(n_Ncsg_MeasGapNR_Patterns_r17),
				}
				ie.Ncsg_MeasGapNR_Patterns_r17 = &tmp_bitstring
			}
			// decode PreconfiguredUE_AutonomousMeasGap_r17 optional
			if PreconfiguredUE_AutonomousMeasGap_r17Present {
				ie.PreconfiguredUE_AutonomousMeasGap_r17 = new(MeasAndMobParametersCommon_preconfiguredUE_AutonomousMeasGap_r17)
				if err = ie.PreconfiguredUE_AutonomousMeasGap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode PreconfiguredUE_AutonomousMeasGap_r17", err)
				}
			}
			// decode PreconfiguredNW_ControlledMeasGap_r17 optional
			if PreconfiguredNW_ControlledMeasGap_r17Present {
				ie.PreconfiguredNW_ControlledMeasGap_r17 = new(MeasAndMobParametersCommon_preconfiguredNW_ControlledMeasGap_r17)
				if err = ie.PreconfiguredNW_ControlledMeasGap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode PreconfiguredNW_ControlledMeasGap_r17", err)
				}
			}
			// decode HandoverFR1_FR2_2_r17 optional
			if HandoverFR1_FR2_2_r17Present {
				ie.HandoverFR1_FR2_2_r17 = new(MeasAndMobParametersCommon_handoverFR1_FR2_2_r17)
				if err = ie.HandoverFR1_FR2_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode HandoverFR1_FR2_2_r17", err)
				}
			}
			// decode HandoverFR2_1_FR2_2_r17 optional
			if HandoverFR2_1_FR2_2_r17Present {
				ie.HandoverFR2_1_FR2_2_r17 = new(MeasAndMobParametersCommon_handoverFR2_1_FR2_2_r17)
				if err = ie.HandoverFR2_1_FR2_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode HandoverFR2_1_FR2_2_r17", err)
				}
			}
			// decode IndependentGapConfigPRS_r17 optional
			if IndependentGapConfigPRS_r17Present {
				ie.IndependentGapConfigPRS_r17 = new(MeasAndMobParametersCommon_independentGapConfigPRS_r17)
				if err = ie.IndependentGapConfigPRS_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode IndependentGapConfigPRS_r17", err)
				}
			}
			// decode Rrm_RelaxationRRC_ConnectedRedCap_r17 optional
			if Rrm_RelaxationRRC_ConnectedRedCap_r17Present {
				ie.Rrm_RelaxationRRC_ConnectedRedCap_r17 = new(MeasAndMobParametersCommon_rrm_RelaxationRRC_ConnectedRedCap_r17)
				if err = ie.Rrm_RelaxationRRC_ConnectedRedCap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Rrm_RelaxationRRC_ConnectedRedCap_r17", err)
				}
			}
			// decode ParallelMeasurementGap_r17 optional
			if ParallelMeasurementGap_r17Present {
				ie.ParallelMeasurementGap_r17 = new(MeasAndMobParametersCommon_parallelMeasurementGap_r17)
				if err = ie.ParallelMeasurementGap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ParallelMeasurementGap_r17", err)
				}
			}
			// decode CondHandoverWithSCG_NRDC_r17 optional
			if CondHandoverWithSCG_NRDC_r17Present {
				ie.CondHandoverWithSCG_NRDC_r17 = new(MeasAndMobParametersCommon_condHandoverWithSCG_NRDC_r17)
				if err = ie.CondHandoverWithSCG_NRDC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode CondHandoverWithSCG_NRDC_r17", err)
				}
			}
			// decode GNB_ID_LengthReporting_r17 optional
			if GNB_ID_LengthReporting_r17Present {
				ie.GNB_ID_LengthReporting_r17 = new(MeasAndMobParametersCommon_gNB_ID_LengthReporting_r17)
				if err = ie.GNB_ID_LengthReporting_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode GNB_ID_LengthReporting_r17", err)
				}
			}
			// decode GNB_ID_LengthReporting_ENDC_r17 optional
			if GNB_ID_LengthReporting_ENDC_r17Present {
				ie.GNB_ID_LengthReporting_ENDC_r17 = new(MeasAndMobParametersCommon_gNB_ID_LengthReporting_ENDC_r17)
				if err = ie.GNB_ID_LengthReporting_ENDC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode GNB_ID_LengthReporting_ENDC_r17", err)
				}
			}
			// decode GNB_ID_LengthReporting_NEDC_r17 optional
			if GNB_ID_LengthReporting_NEDC_r17Present {
				ie.GNB_ID_LengthReporting_NEDC_r17 = new(MeasAndMobParametersCommon_gNB_ID_LengthReporting_NEDC_r17)
				if err = ie.GNB_ID_LengthReporting_NEDC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode GNB_ID_LengthReporting_NEDC_r17", err)
				}
			}
			// decode GNB_ID_LengthReporting_NRDC_r17 optional
			if GNB_ID_LengthReporting_NRDC_r17Present {
				ie.GNB_ID_LengthReporting_NRDC_r17 = new(MeasAndMobParametersCommon_gNB_ID_LengthReporting_NRDC_r17)
				if err = ie.GNB_ID_LengthReporting_NRDC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode GNB_ID_LengthReporting_NRDC_r17", err)
				}
			}
			// decode GNB_ID_LengthReporting_NPN_r17 optional
			if GNB_ID_LengthReporting_NPN_r17Present {
				ie.GNB_ID_LengthReporting_NPN_r17 = new(MeasAndMobParametersCommon_gNB_ID_LengthReporting_NPN_r17)
				if err = ie.GNB_ID_LengthReporting_NPN_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode GNB_ID_LengthReporting_NPN_r17", err)
				}
			}
		}
		// decode extension group 7
		if len(extBitmap) > 6 && extBitmap[6] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ParallelSMTC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ConcurrentMeasGapEUTRA_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ServiceLinkPropDelayDiffReporting_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ncsg_SymbolLevelScheduleRestrictionInter_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ParallelSMTC_r17 optional
			if ParallelSMTC_r17Present {
				ie.ParallelSMTC_r17 = new(MeasAndMobParametersCommon_parallelSMTC_r17)
				if err = ie.ParallelSMTC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ParallelSMTC_r17", err)
				}
			}
			// decode ConcurrentMeasGapEUTRA_r17 optional
			if ConcurrentMeasGapEUTRA_r17Present {
				ie.ConcurrentMeasGapEUTRA_r17 = new(MeasAndMobParametersCommon_concurrentMeasGapEUTRA_r17)
				if err = ie.ConcurrentMeasGapEUTRA_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ConcurrentMeasGapEUTRA_r17", err)
				}
			}
			// decode ServiceLinkPropDelayDiffReporting_r17 optional
			if ServiceLinkPropDelayDiffReporting_r17Present {
				ie.ServiceLinkPropDelayDiffReporting_r17 = new(MeasAndMobParametersCommon_serviceLinkPropDelayDiffReporting_r17)
				if err = ie.ServiceLinkPropDelayDiffReporting_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ServiceLinkPropDelayDiffReporting_r17", err)
				}
			}
			// decode Ncsg_SymbolLevelScheduleRestrictionInter_r17 optional
			if Ncsg_SymbolLevelScheduleRestrictionInter_r17Present {
				ie.Ncsg_SymbolLevelScheduleRestrictionInter_r17 = new(MeasAndMobParametersCommon_ncsg_SymbolLevelScheduleRestrictionInter_r17)
				if err = ie.Ncsg_SymbolLevelScheduleRestrictionInter_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ncsg_SymbolLevelScheduleRestrictionInter_r17", err)
				}
			}
		}
		// decode extension group 8
		if len(extBitmap) > 7 && extBitmap[7] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			EventD1_MeasReportTrigger_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			IndependentGapConfig_maxCC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode EventD1_MeasReportTrigger_r17 optional
			if EventD1_MeasReportTrigger_r17Present {
				ie.EventD1_MeasReportTrigger_r17 = new(MeasAndMobParametersCommon_eventD1_MeasReportTrigger_r17)
				if err = ie.EventD1_MeasReportTrigger_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode EventD1_MeasReportTrigger_r17", err)
				}
			}
			// decode IndependentGapConfig_maxCC_r17 optional
			if IndependentGapConfig_maxCC_r17Present {
				ie.IndependentGapConfig_maxCC_r17 = new(MeasAndMobParametersCommon_independentGapConfig_maxCC_r17)
				if err = ie.IndependentGapConfig_maxCC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode IndependentGapConfig_maxCC_r17", err)
				}
			}
		}
	}
	return nil
}
