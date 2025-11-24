package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersFRX_Diff struct {
	Ss_SINR_Meas                                   *MeasAndMobParametersFRX_Diff_ss_SINR_Meas                                   `optional`
	Csi_RSRP_AndRSRQ_MeasWithSSB                   *MeasAndMobParametersFRX_Diff_csi_RSRP_AndRSRQ_MeasWithSSB                   `optional`
	Csi_RSRP_AndRSRQ_MeasWithoutSSB                *MeasAndMobParametersFRX_Diff_csi_RSRP_AndRSRQ_MeasWithoutSSB                `optional`
	Csi_SINR_Meas                                  *MeasAndMobParametersFRX_Diff_csi_SINR_Meas                                  `optional`
	Csi_RS_RLM                                     *MeasAndMobParametersFRX_Diff_csi_RS_RLM                                     `optional`
	HandoverInterF                                 *MeasAndMobParametersFRX_Diff_handoverInterF                                 `optional,ext-1`
	HandoverLTE_EPC                                *MeasAndMobParametersFRX_Diff_handoverLTE_EPC                                `optional,ext-1`
	HandoverLTE_5GC                                *MeasAndMobParametersFRX_Diff_handoverLTE_5GC                                `optional,ext-1`
	MaxNumberResource_CSI_RS_RLM                   *MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM                   `optional,ext-2`
	SimultaneousRxDataSSB_DiffNumerology           *MeasAndMobParametersFRX_Diff_simultaneousRxDataSSB_DiffNumerology           `optional,ext-3`
	Nr_AutonomousGaps_r16                          *MeasAndMobParametersFRX_Diff_nr_AutonomousGaps_r16                          `optional,ext-4`
	Nr_AutonomousGaps_ENDC_r16                     *MeasAndMobParametersFRX_Diff_nr_AutonomousGaps_ENDC_r16                     `optional,ext-4`
	Nr_AutonomousGaps_NEDC_r16                     *MeasAndMobParametersFRX_Diff_nr_AutonomousGaps_NEDC_r16                     `optional,ext-4`
	Nr_AutonomousGaps_NRDC_r16                     *MeasAndMobParametersFRX_Diff_nr_AutonomousGaps_NRDC_r16                     `optional,ext-4`
	Dummy                                          *MeasAndMobParametersFRX_Diff_dummy                                          `optional,ext-4`
	Cli_RSSI_Meas_r16                              *MeasAndMobParametersFRX_Diff_cli_RSSI_Meas_r16                              `optional,ext-4`
	Cli_SRS_RSRP_Meas_r16                          *MeasAndMobParametersFRX_Diff_cli_SRS_RSRP_Meas_r16                          `optional,ext-4`
	InterFrequencyMeas_NoGap_r16                   *MeasAndMobParametersFRX_Diff_interFrequencyMeas_NoGap_r16                   `optional,ext-4`
	SimultaneousRxDataSSB_DiffNumerology_Inter_r16 *MeasAndMobParametersFRX_Diff_simultaneousRxDataSSB_DiffNumerology_Inter_r16 `optional,ext-4`
	IdleInactiveNR_MeasReport_r16                  *MeasAndMobParametersFRX_Diff_idleInactiveNR_MeasReport_r16                  `optional,ext-4`
	IdleInactiveNR_MeasBeamReport_r16              *MeasAndMobParametersFRX_Diff_idleInactiveNR_MeasBeamReport_r16              `optional,ext-4`
	IncreasedNumberofCSIRSPerMO_r16                *MeasAndMobParametersFRX_Diff_increasedNumberofCSIRSPerMO_r16                `optional,ext-5`
}

func (ie *MeasAndMobParametersFRX_Diff) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.HandoverInterF != nil || ie.HandoverLTE_EPC != nil || ie.HandoverLTE_5GC != nil || ie.MaxNumberResource_CSI_RS_RLM != nil || ie.SimultaneousRxDataSSB_DiffNumerology != nil || ie.Nr_AutonomousGaps_r16 != nil || ie.Nr_AutonomousGaps_ENDC_r16 != nil || ie.Nr_AutonomousGaps_NEDC_r16 != nil || ie.Nr_AutonomousGaps_NRDC_r16 != nil || ie.Dummy != nil || ie.Cli_RSSI_Meas_r16 != nil || ie.Cli_SRS_RSRP_Meas_r16 != nil || ie.InterFrequencyMeas_NoGap_r16 != nil || ie.SimultaneousRxDataSSB_DiffNumerology_Inter_r16 != nil || ie.IdleInactiveNR_MeasReport_r16 != nil || ie.IdleInactiveNR_MeasBeamReport_r16 != nil || ie.IncreasedNumberofCSIRSPerMO_r16 != nil
	preambleBits := []bool{hasExtensions, ie.Ss_SINR_Meas != nil, ie.Csi_RSRP_AndRSRQ_MeasWithSSB != nil, ie.Csi_RSRP_AndRSRQ_MeasWithoutSSB != nil, ie.Csi_SINR_Meas != nil, ie.Csi_RS_RLM != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ss_SINR_Meas != nil {
		if err = ie.Ss_SINR_Meas.Encode(w); err != nil {
			return utils.WrapError("Encode Ss_SINR_Meas", err)
		}
	}
	if ie.Csi_RSRP_AndRSRQ_MeasWithSSB != nil {
		if err = ie.Csi_RSRP_AndRSRQ_MeasWithSSB.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_RSRP_AndRSRQ_MeasWithSSB", err)
		}
	}
	if ie.Csi_RSRP_AndRSRQ_MeasWithoutSSB != nil {
		if err = ie.Csi_RSRP_AndRSRQ_MeasWithoutSSB.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_RSRP_AndRSRQ_MeasWithoutSSB", err)
		}
	}
	if ie.Csi_SINR_Meas != nil {
		if err = ie.Csi_SINR_Meas.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_SINR_Meas", err)
		}
	}
	if ie.Csi_RS_RLM != nil {
		if err = ie.Csi_RS_RLM.Encode(w); err != nil {
			return utils.WrapError("Encode Csi_RS_RLM", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 5 bits for 5 extension groups
		extBitmap := []bool{ie.HandoverInterF != nil || ie.HandoverLTE_EPC != nil || ie.HandoverLTE_5GC != nil, ie.MaxNumberResource_CSI_RS_RLM != nil, ie.SimultaneousRxDataSSB_DiffNumerology != nil, ie.Nr_AutonomousGaps_r16 != nil || ie.Nr_AutonomousGaps_ENDC_r16 != nil || ie.Nr_AutonomousGaps_NEDC_r16 != nil || ie.Nr_AutonomousGaps_NRDC_r16 != nil || ie.Dummy != nil || ie.Cli_RSSI_Meas_r16 != nil || ie.Cli_SRS_RSRP_Meas_r16 != nil || ie.InterFrequencyMeas_NoGap_r16 != nil || ie.SimultaneousRxDataSSB_DiffNumerology_Inter_r16 != nil || ie.IdleInactiveNR_MeasReport_r16 != nil || ie.IdleInactiveNR_MeasBeamReport_r16 != nil, ie.IncreasedNumberofCSIRSPerMO_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasAndMobParametersFRX_Diff", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.HandoverInterF != nil, ie.HandoverLTE_EPC != nil, ie.HandoverLTE_5GC != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode HandoverInterF optional
			if ie.HandoverInterF != nil {
				if err = ie.HandoverInterF.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HandoverInterF", err)
				}
			}
			// encode HandoverLTE_EPC optional
			if ie.HandoverLTE_EPC != nil {
				if err = ie.HandoverLTE_EPC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HandoverLTE_EPC", err)
				}
			}
			// encode HandoverLTE_5GC optional
			if ie.HandoverLTE_5GC != nil {
				if err = ie.HandoverLTE_5GC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode HandoverLTE_5GC", err)
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
			optionals_ext_2 := []bool{ie.MaxNumberResource_CSI_RS_RLM != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode MaxNumberResource_CSI_RS_RLM optional
			if ie.MaxNumberResource_CSI_RS_RLM != nil {
				if err = ie.MaxNumberResource_CSI_RS_RLM.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxNumberResource_CSI_RS_RLM", err)
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
			optionals_ext_3 := []bool{ie.SimultaneousRxDataSSB_DiffNumerology != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SimultaneousRxDataSSB_DiffNumerology optional
			if ie.SimultaneousRxDataSSB_DiffNumerology != nil {
				if err = ie.SimultaneousRxDataSSB_DiffNumerology.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SimultaneousRxDataSSB_DiffNumerology", err)
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
			optionals_ext_4 := []bool{ie.Nr_AutonomousGaps_r16 != nil, ie.Nr_AutonomousGaps_ENDC_r16 != nil, ie.Nr_AutonomousGaps_NEDC_r16 != nil, ie.Nr_AutonomousGaps_NRDC_r16 != nil, ie.Dummy != nil, ie.Cli_RSSI_Meas_r16 != nil, ie.Cli_SRS_RSRP_Meas_r16 != nil, ie.InterFrequencyMeas_NoGap_r16 != nil, ie.SimultaneousRxDataSSB_DiffNumerology_Inter_r16 != nil, ie.IdleInactiveNR_MeasReport_r16 != nil, ie.IdleInactiveNR_MeasBeamReport_r16 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Nr_AutonomousGaps_r16 optional
			if ie.Nr_AutonomousGaps_r16 != nil {
				if err = ie.Nr_AutonomousGaps_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Nr_AutonomousGaps_r16", err)
				}
			}
			// encode Nr_AutonomousGaps_ENDC_r16 optional
			if ie.Nr_AutonomousGaps_ENDC_r16 != nil {
				if err = ie.Nr_AutonomousGaps_ENDC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Nr_AutonomousGaps_ENDC_r16", err)
				}
			}
			// encode Nr_AutonomousGaps_NEDC_r16 optional
			if ie.Nr_AutonomousGaps_NEDC_r16 != nil {
				if err = ie.Nr_AutonomousGaps_NEDC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Nr_AutonomousGaps_NEDC_r16", err)
				}
			}
			// encode Nr_AutonomousGaps_NRDC_r16 optional
			if ie.Nr_AutonomousGaps_NRDC_r16 != nil {
				if err = ie.Nr_AutonomousGaps_NRDC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Nr_AutonomousGaps_NRDC_r16", err)
				}
			}
			// encode Dummy optional
			if ie.Dummy != nil {
				if err = ie.Dummy.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dummy", err)
				}
			}
			// encode Cli_RSSI_Meas_r16 optional
			if ie.Cli_RSSI_Meas_r16 != nil {
				if err = ie.Cli_RSSI_Meas_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cli_RSSI_Meas_r16", err)
				}
			}
			// encode Cli_SRS_RSRP_Meas_r16 optional
			if ie.Cli_SRS_RSRP_Meas_r16 != nil {
				if err = ie.Cli_SRS_RSRP_Meas_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cli_SRS_RSRP_Meas_r16", err)
				}
			}
			// encode InterFrequencyMeas_NoGap_r16 optional
			if ie.InterFrequencyMeas_NoGap_r16 != nil {
				if err = ie.InterFrequencyMeas_NoGap_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode InterFrequencyMeas_NoGap_r16", err)
				}
			}
			// encode SimultaneousRxDataSSB_DiffNumerology_Inter_r16 optional
			if ie.SimultaneousRxDataSSB_DiffNumerology_Inter_r16 != nil {
				if err = ie.SimultaneousRxDataSSB_DiffNumerology_Inter_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SimultaneousRxDataSSB_DiffNumerology_Inter_r16", err)
				}
			}
			// encode IdleInactiveNR_MeasReport_r16 optional
			if ie.IdleInactiveNR_MeasReport_r16 != nil {
				if err = ie.IdleInactiveNR_MeasReport_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IdleInactiveNR_MeasReport_r16", err)
				}
			}
			// encode IdleInactiveNR_MeasBeamReport_r16 optional
			if ie.IdleInactiveNR_MeasBeamReport_r16 != nil {
				if err = ie.IdleInactiveNR_MeasBeamReport_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IdleInactiveNR_MeasBeamReport_r16", err)
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
			optionals_ext_5 := []bool{ie.IncreasedNumberofCSIRSPerMO_r16 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode IncreasedNumberofCSIRSPerMO_r16 optional
			if ie.IncreasedNumberofCSIRSPerMO_r16 != nil {
				if err = ie.IncreasedNumberofCSIRSPerMO_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode IncreasedNumberofCSIRSPerMO_r16", err)
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

func (ie *MeasAndMobParametersFRX_Diff) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Ss_SINR_MeasPresent bool
	if Ss_SINR_MeasPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_RSRP_AndRSRQ_MeasWithSSBPresent bool
	if Csi_RSRP_AndRSRQ_MeasWithSSBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_RSRP_AndRSRQ_MeasWithoutSSBPresent bool
	if Csi_RSRP_AndRSRQ_MeasWithoutSSBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_SINR_MeasPresent bool
	if Csi_SINR_MeasPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_RS_RLMPresent bool
	if Csi_RS_RLMPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Ss_SINR_MeasPresent {
		ie.Ss_SINR_Meas = new(MeasAndMobParametersFRX_Diff_ss_SINR_Meas)
		if err = ie.Ss_SINR_Meas.Decode(r); err != nil {
			return utils.WrapError("Decode Ss_SINR_Meas", err)
		}
	}
	if Csi_RSRP_AndRSRQ_MeasWithSSBPresent {
		ie.Csi_RSRP_AndRSRQ_MeasWithSSB = new(MeasAndMobParametersFRX_Diff_csi_RSRP_AndRSRQ_MeasWithSSB)
		if err = ie.Csi_RSRP_AndRSRQ_MeasWithSSB.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RSRP_AndRSRQ_MeasWithSSB", err)
		}
	}
	if Csi_RSRP_AndRSRQ_MeasWithoutSSBPresent {
		ie.Csi_RSRP_AndRSRQ_MeasWithoutSSB = new(MeasAndMobParametersFRX_Diff_csi_RSRP_AndRSRQ_MeasWithoutSSB)
		if err = ie.Csi_RSRP_AndRSRQ_MeasWithoutSSB.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RSRP_AndRSRQ_MeasWithoutSSB", err)
		}
	}
	if Csi_SINR_MeasPresent {
		ie.Csi_SINR_Meas = new(MeasAndMobParametersFRX_Diff_csi_SINR_Meas)
		if err = ie.Csi_SINR_Meas.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_SINR_Meas", err)
		}
	}
	if Csi_RS_RLMPresent {
		ie.Csi_RS_RLM = new(MeasAndMobParametersFRX_Diff_csi_RS_RLM)
		if err = ie.Csi_RS_RLM.Decode(r); err != nil {
			return utils.WrapError("Decode Csi_RS_RLM", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 5 bits for 5 extension groups
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

			HandoverInterFPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			HandoverLTE_EPCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			HandoverLTE_5GCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode HandoverInterF optional
			if HandoverInterFPresent {
				ie.HandoverInterF = new(MeasAndMobParametersFRX_Diff_handoverInterF)
				if err = ie.HandoverInterF.Decode(extReader); err != nil {
					return utils.WrapError("Decode HandoverInterF", err)
				}
			}
			// decode HandoverLTE_EPC optional
			if HandoverLTE_EPCPresent {
				ie.HandoverLTE_EPC = new(MeasAndMobParametersFRX_Diff_handoverLTE_EPC)
				if err = ie.HandoverLTE_EPC.Decode(extReader); err != nil {
					return utils.WrapError("Decode HandoverLTE_EPC", err)
				}
			}
			// decode HandoverLTE_5GC optional
			if HandoverLTE_5GCPresent {
				ie.HandoverLTE_5GC = new(MeasAndMobParametersFRX_Diff_handoverLTE_5GC)
				if err = ie.HandoverLTE_5GC.Decode(extReader); err != nil {
					return utils.WrapError("Decode HandoverLTE_5GC", err)
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

			MaxNumberResource_CSI_RS_RLMPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode MaxNumberResource_CSI_RS_RLM optional
			if MaxNumberResource_CSI_RS_RLMPresent {
				ie.MaxNumberResource_CSI_RS_RLM = new(MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM)
				if err = ie.MaxNumberResource_CSI_RS_RLM.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxNumberResource_CSI_RS_RLM", err)
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

			SimultaneousRxDataSSB_DiffNumerologyPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SimultaneousRxDataSSB_DiffNumerology optional
			if SimultaneousRxDataSSB_DiffNumerologyPresent {
				ie.SimultaneousRxDataSSB_DiffNumerology = new(MeasAndMobParametersFRX_Diff_simultaneousRxDataSSB_DiffNumerology)
				if err = ie.SimultaneousRxDataSSB_DiffNumerology.Decode(extReader); err != nil {
					return utils.WrapError("Decode SimultaneousRxDataSSB_DiffNumerology", err)
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

			Nr_AutonomousGaps_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Nr_AutonomousGaps_ENDC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Nr_AutonomousGaps_NEDC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Nr_AutonomousGaps_NRDC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			DummyPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cli_RSSI_Meas_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Cli_SRS_RSRP_Meas_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			InterFrequencyMeas_NoGap_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SimultaneousRxDataSSB_DiffNumerology_Inter_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			IdleInactiveNR_MeasReport_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			IdleInactiveNR_MeasBeamReport_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Nr_AutonomousGaps_r16 optional
			if Nr_AutonomousGaps_r16Present {
				ie.Nr_AutonomousGaps_r16 = new(MeasAndMobParametersFRX_Diff_nr_AutonomousGaps_r16)
				if err = ie.Nr_AutonomousGaps_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Nr_AutonomousGaps_r16", err)
				}
			}
			// decode Nr_AutonomousGaps_ENDC_r16 optional
			if Nr_AutonomousGaps_ENDC_r16Present {
				ie.Nr_AutonomousGaps_ENDC_r16 = new(MeasAndMobParametersFRX_Diff_nr_AutonomousGaps_ENDC_r16)
				if err = ie.Nr_AutonomousGaps_ENDC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Nr_AutonomousGaps_ENDC_r16", err)
				}
			}
			// decode Nr_AutonomousGaps_NEDC_r16 optional
			if Nr_AutonomousGaps_NEDC_r16Present {
				ie.Nr_AutonomousGaps_NEDC_r16 = new(MeasAndMobParametersFRX_Diff_nr_AutonomousGaps_NEDC_r16)
				if err = ie.Nr_AutonomousGaps_NEDC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Nr_AutonomousGaps_NEDC_r16", err)
				}
			}
			// decode Nr_AutonomousGaps_NRDC_r16 optional
			if Nr_AutonomousGaps_NRDC_r16Present {
				ie.Nr_AutonomousGaps_NRDC_r16 = new(MeasAndMobParametersFRX_Diff_nr_AutonomousGaps_NRDC_r16)
				if err = ie.Nr_AutonomousGaps_NRDC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Nr_AutonomousGaps_NRDC_r16", err)
				}
			}
			// decode Dummy optional
			if DummyPresent {
				ie.Dummy = new(MeasAndMobParametersFRX_Diff_dummy)
				if err = ie.Dummy.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dummy", err)
				}
			}
			// decode Cli_RSSI_Meas_r16 optional
			if Cli_RSSI_Meas_r16Present {
				ie.Cli_RSSI_Meas_r16 = new(MeasAndMobParametersFRX_Diff_cli_RSSI_Meas_r16)
				if err = ie.Cli_RSSI_Meas_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cli_RSSI_Meas_r16", err)
				}
			}
			// decode Cli_SRS_RSRP_Meas_r16 optional
			if Cli_SRS_RSRP_Meas_r16Present {
				ie.Cli_SRS_RSRP_Meas_r16 = new(MeasAndMobParametersFRX_Diff_cli_SRS_RSRP_Meas_r16)
				if err = ie.Cli_SRS_RSRP_Meas_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cli_SRS_RSRP_Meas_r16", err)
				}
			}
			// decode InterFrequencyMeas_NoGap_r16 optional
			if InterFrequencyMeas_NoGap_r16Present {
				ie.InterFrequencyMeas_NoGap_r16 = new(MeasAndMobParametersFRX_Diff_interFrequencyMeas_NoGap_r16)
				if err = ie.InterFrequencyMeas_NoGap_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode InterFrequencyMeas_NoGap_r16", err)
				}
			}
			// decode SimultaneousRxDataSSB_DiffNumerology_Inter_r16 optional
			if SimultaneousRxDataSSB_DiffNumerology_Inter_r16Present {
				ie.SimultaneousRxDataSSB_DiffNumerology_Inter_r16 = new(MeasAndMobParametersFRX_Diff_simultaneousRxDataSSB_DiffNumerology_Inter_r16)
				if err = ie.SimultaneousRxDataSSB_DiffNumerology_Inter_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SimultaneousRxDataSSB_DiffNumerology_Inter_r16", err)
				}
			}
			// decode IdleInactiveNR_MeasReport_r16 optional
			if IdleInactiveNR_MeasReport_r16Present {
				ie.IdleInactiveNR_MeasReport_r16 = new(MeasAndMobParametersFRX_Diff_idleInactiveNR_MeasReport_r16)
				if err = ie.IdleInactiveNR_MeasReport_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode IdleInactiveNR_MeasReport_r16", err)
				}
			}
			// decode IdleInactiveNR_MeasBeamReport_r16 optional
			if IdleInactiveNR_MeasBeamReport_r16Present {
				ie.IdleInactiveNR_MeasBeamReport_r16 = new(MeasAndMobParametersFRX_Diff_idleInactiveNR_MeasBeamReport_r16)
				if err = ie.IdleInactiveNR_MeasBeamReport_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode IdleInactiveNR_MeasBeamReport_r16", err)
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

			IncreasedNumberofCSIRSPerMO_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode IncreasedNumberofCSIRSPerMO_r16 optional
			if IncreasedNumberofCSIRSPerMO_r16Present {
				ie.IncreasedNumberofCSIRSPerMO_r16 = new(MeasAndMobParametersFRX_Diff_increasedNumberofCSIRSPerMO_r16)
				if err = ie.IncreasedNumberofCSIRSPerMO_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode IncreasedNumberofCSIRSPerMO_r16", err)
				}
			}
		}
	}
	return nil
}
