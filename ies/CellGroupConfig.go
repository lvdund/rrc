package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellGroupConfig struct {
	CellGroupId                                CellGroupId                                                `madatory`
	Rlc_BearerToAddModList                     []RLC_BearerConfig                                         `lb:1,ub:maxLC_ID,optional`
	Rlc_BearerToReleaseList                    []LogicalChannelIdentity                                   `lb:1,ub:maxLC_ID,optional`
	Mac_CellGroupConfig                        *MAC_CellGroupConfig                                       `optional`
	PhysicalCellGroupConfig                    *PhysicalCellGroupConfig                                   `optional`
	SpCellConfig                               *SpCellConfig                                              `optional`
	SCellToAddModList                          []SCellConfig                                              `lb:1,ub:maxNrofSCells,optional`
	SCellToReleaseList                         []SCellIndex                                               `lb:1,ub:maxNrofSCells,optional`
	ReportUplinkTxDirectCurrent                *CellGroupConfig_reportUplinkTxDirectCurrent               `optional,ext-1`
	Bap_Address_r16                            *uper.BitString                                            `lb:10,ub:10,optional,ext-2`
	Bh_RLC_ChannelToAddModList_r16             []BH_RLC_ChannelConfig_r16                                 `lb:1,ub:maxBH_RLC_ChannelID_r16,optional,ext-2`
	Bh_RLC_ChannelToReleaseList_r16            []BH_RLC_ChannelID_r16                                     `lb:1,ub:maxBH_RLC_ChannelID_r16,optional,ext-2`
	F1c_TransferPath_r16                       *CellGroupConfig_f1c_TransferPath_r16                      `optional,ext-2`
	SimultaneousTCI_UpdateList1_r16            []ServCellIndex                                            `lb:1,ub:maxNrofServingCellsTCI_r16,optional,ext-2`
	SimultaneousTCI_UpdateList2_r16            []ServCellIndex                                            `lb:1,ub:maxNrofServingCellsTCI_r16,optional,ext-2`
	SimultaneousSpatial_UpdatedList1_r16       []ServCellIndex                                            `lb:1,ub:maxNrofServingCellsTCI_r16,optional,ext-2`
	SimultaneousSpatial_UpdatedList2_r16       []ServCellIndex                                            `lb:1,ub:maxNrofServingCellsTCI_r16,optional,ext-2`
	UplinkTxSwitchingOption_r16                *CellGroupConfig_uplinkTxSwitchingOption_r16               `optional,ext-2`
	UplinkTxSwitchingPowerBoosting_r16         *CellGroupConfig_uplinkTxSwitchingPowerBoosting_r16        `optional,ext-2`
	ReportUplinkTxDirectCurrentTwoCarrier_r16  *CellGroupConfig_reportUplinkTxDirectCurrentTwoCarrier_r16 `optional,ext-3`
	F1c_TransferPathNRDC_r17                   *CellGroupConfig_f1c_TransferPathNRDC_r17                  `optional,ext-4`
	UplinkTxSwitching_2T_Mode_r17              *CellGroupConfig_uplinkTxSwitching_2T_Mode_r17             `optional,ext-4`
	UplinkTxSwitching_DualUL_TxState_r17       *CellGroupConfig_uplinkTxSwitching_DualUL_TxState_r17      `optional,ext-4`
	Uu_RelayRLC_ChannelToAddModList_r17        []Uu_RelayRLC_ChannelConfig_r17                            `lb:1,ub:maxUu_RelayRLC_ChannelID_r17,optional,ext-4`
	Uu_RelayRLC_ChannelToReleaseList_r17       []Uu_RelayRLC_ChannelID_r17                                `lb:1,ub:maxUu_RelayRLC_ChannelID_r17,optional,ext-4`
	SimultaneousU_TCI_UpdateList1_r17          []ServCellIndex                                            `lb:1,ub:maxNrofServingCellsTCI_r16,optional,ext-4`
	SimultaneousU_TCI_UpdateList2_r17          []ServCellIndex                                            `lb:1,ub:maxNrofServingCellsTCI_r16,optional,ext-4`
	SimultaneousU_TCI_UpdateList3_r17          []ServCellIndex                                            `lb:1,ub:maxNrofServingCellsTCI_r16,optional,ext-4`
	SimultaneousU_TCI_UpdateList4_r17          []ServCellIndex                                            `lb:1,ub:maxNrofServingCellsTCI_r16,optional,ext-4`
	Rlc_BearerToReleaseListExt_r17             []LogicalChannelIdentityExt_r17                            `lb:1,ub:maxLC_ID,optional,ext-4`
	Iab_ResourceConfigToAddModList_r17         []IAB_ResourceConfig_r17                                   `lb:1,ub:maxNrofIABResourceConfig_r17,optional,ext-4`
	Iab_ResourceConfigToReleaseList_r17        []IAB_ResourceConfigID_r17                                 `lb:1,ub:maxNrofIABResourceConfig_r17,optional,ext-4`
	ReportUplinkTxDirectCurrentMoreCarrier_r17 *ReportUplinkTxDirectCurrentMoreCarrier_r17                `optional,ext-5`
}

func (ie *CellGroupConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.ReportUplinkTxDirectCurrent != nil || ie.Bap_Address_r16 != nil || len(ie.Bh_RLC_ChannelToAddModList_r16) > 0 || len(ie.Bh_RLC_ChannelToReleaseList_r16) > 0 || ie.F1c_TransferPath_r16 != nil || len(ie.SimultaneousTCI_UpdateList1_r16) > 0 || len(ie.SimultaneousTCI_UpdateList2_r16) > 0 || len(ie.SimultaneousSpatial_UpdatedList1_r16) > 0 || len(ie.SimultaneousSpatial_UpdatedList2_r16) > 0 || ie.UplinkTxSwitchingOption_r16 != nil || ie.UplinkTxSwitchingPowerBoosting_r16 != nil || ie.ReportUplinkTxDirectCurrentTwoCarrier_r16 != nil || ie.F1c_TransferPathNRDC_r17 != nil || ie.UplinkTxSwitching_2T_Mode_r17 != nil || ie.UplinkTxSwitching_DualUL_TxState_r17 != nil || len(ie.Uu_RelayRLC_ChannelToAddModList_r17) > 0 || len(ie.Uu_RelayRLC_ChannelToReleaseList_r17) > 0 || len(ie.SimultaneousU_TCI_UpdateList1_r17) > 0 || len(ie.SimultaneousU_TCI_UpdateList2_r17) > 0 || len(ie.SimultaneousU_TCI_UpdateList3_r17) > 0 || len(ie.SimultaneousU_TCI_UpdateList4_r17) > 0 || len(ie.Rlc_BearerToReleaseListExt_r17) > 0 || len(ie.Iab_ResourceConfigToAddModList_r17) > 0 || len(ie.Iab_ResourceConfigToReleaseList_r17) > 0 || ie.ReportUplinkTxDirectCurrentMoreCarrier_r17 != nil
	preambleBits := []bool{hasExtensions, len(ie.Rlc_BearerToAddModList) > 0, len(ie.Rlc_BearerToReleaseList) > 0, ie.Mac_CellGroupConfig != nil, ie.PhysicalCellGroupConfig != nil, ie.SpCellConfig != nil, len(ie.SCellToAddModList) > 0, len(ie.SCellToReleaseList) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.CellGroupId.Encode(w); err != nil {
		return utils.WrapError("Encode CellGroupId", err)
	}
	if len(ie.Rlc_BearerToAddModList) > 0 {
		tmp_Rlc_BearerToAddModList := utils.NewSequence[*RLC_BearerConfig]([]*RLC_BearerConfig{}, uper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
		for _, i := range ie.Rlc_BearerToAddModList {
			tmp_Rlc_BearerToAddModList.Value = append(tmp_Rlc_BearerToAddModList.Value, &i)
		}
		if err = tmp_Rlc_BearerToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode Rlc_BearerToAddModList", err)
		}
	}
	if len(ie.Rlc_BearerToReleaseList) > 0 {
		tmp_Rlc_BearerToReleaseList := utils.NewSequence[*LogicalChannelIdentity]([]*LogicalChannelIdentity{}, uper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
		for _, i := range ie.Rlc_BearerToReleaseList {
			tmp_Rlc_BearerToReleaseList.Value = append(tmp_Rlc_BearerToReleaseList.Value, &i)
		}
		if err = tmp_Rlc_BearerToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode Rlc_BearerToReleaseList", err)
		}
	}
	if ie.Mac_CellGroupConfig != nil {
		if err = ie.Mac_CellGroupConfig.Encode(w); err != nil {
			return utils.WrapError("Encode Mac_CellGroupConfig", err)
		}
	}
	if ie.PhysicalCellGroupConfig != nil {
		if err = ie.PhysicalCellGroupConfig.Encode(w); err != nil {
			return utils.WrapError("Encode PhysicalCellGroupConfig", err)
		}
	}
	if ie.SpCellConfig != nil {
		if err = ie.SpCellConfig.Encode(w); err != nil {
			return utils.WrapError("Encode SpCellConfig", err)
		}
	}
	if len(ie.SCellToAddModList) > 0 {
		tmp_SCellToAddModList := utils.NewSequence[*SCellConfig]([]*SCellConfig{}, uper.Constraint{Lb: 1, Ub: maxNrofSCells}, false)
		for _, i := range ie.SCellToAddModList {
			tmp_SCellToAddModList.Value = append(tmp_SCellToAddModList.Value, &i)
		}
		if err = tmp_SCellToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode SCellToAddModList", err)
		}
	}
	if len(ie.SCellToReleaseList) > 0 {
		tmp_SCellToReleaseList := utils.NewSequence[*SCellIndex]([]*SCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofSCells}, false)
		for _, i := range ie.SCellToReleaseList {
			tmp_SCellToReleaseList.Value = append(tmp_SCellToReleaseList.Value, &i)
		}
		if err = tmp_SCellToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode SCellToReleaseList", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 5 bits for 5 extension groups
		extBitmap := []bool{ie.ReportUplinkTxDirectCurrent != nil, ie.Bap_Address_r16 != nil || len(ie.Bh_RLC_ChannelToAddModList_r16) > 0 || len(ie.Bh_RLC_ChannelToReleaseList_r16) > 0 || ie.F1c_TransferPath_r16 != nil || len(ie.SimultaneousTCI_UpdateList1_r16) > 0 || len(ie.SimultaneousTCI_UpdateList2_r16) > 0 || len(ie.SimultaneousSpatial_UpdatedList1_r16) > 0 || len(ie.SimultaneousSpatial_UpdatedList2_r16) > 0 || ie.UplinkTxSwitchingOption_r16 != nil || ie.UplinkTxSwitchingPowerBoosting_r16 != nil, ie.ReportUplinkTxDirectCurrentTwoCarrier_r16 != nil, ie.F1c_TransferPathNRDC_r17 != nil || ie.UplinkTxSwitching_2T_Mode_r17 != nil || ie.UplinkTxSwitching_DualUL_TxState_r17 != nil || len(ie.Uu_RelayRLC_ChannelToAddModList_r17) > 0 || len(ie.Uu_RelayRLC_ChannelToReleaseList_r17) > 0 || len(ie.SimultaneousU_TCI_UpdateList1_r17) > 0 || len(ie.SimultaneousU_TCI_UpdateList2_r17) > 0 || len(ie.SimultaneousU_TCI_UpdateList3_r17) > 0 || len(ie.SimultaneousU_TCI_UpdateList4_r17) > 0 || len(ie.Rlc_BearerToReleaseListExt_r17) > 0 || len(ie.Iab_ResourceConfigToAddModList_r17) > 0 || len(ie.Iab_ResourceConfigToReleaseList_r17) > 0, ie.ReportUplinkTxDirectCurrentMoreCarrier_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CellGroupConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.ReportUplinkTxDirectCurrent != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ReportUplinkTxDirectCurrent optional
			if ie.ReportUplinkTxDirectCurrent != nil {
				if err = ie.ReportUplinkTxDirectCurrent.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ReportUplinkTxDirectCurrent", err)
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
			optionals_ext_2 := []bool{ie.Bap_Address_r16 != nil, len(ie.Bh_RLC_ChannelToAddModList_r16) > 0, len(ie.Bh_RLC_ChannelToReleaseList_r16) > 0, ie.F1c_TransferPath_r16 != nil, len(ie.SimultaneousTCI_UpdateList1_r16) > 0, len(ie.SimultaneousTCI_UpdateList2_r16) > 0, len(ie.SimultaneousSpatial_UpdatedList1_r16) > 0, len(ie.SimultaneousSpatial_UpdatedList2_r16) > 0, ie.UplinkTxSwitchingOption_r16 != nil, ie.UplinkTxSwitchingPowerBoosting_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Bap_Address_r16 optional
			if ie.Bap_Address_r16 != nil {
				if err = extWriter.WriteBitString(ie.Bap_Address_r16.Bytes, uint(ie.Bap_Address_r16.NumBits), &uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
					return utils.WrapError("Encode Bap_Address_r16", err)
				}
			}
			// encode Bh_RLC_ChannelToAddModList_r16 optional
			if len(ie.Bh_RLC_ChannelToAddModList_r16) > 0 {
				tmp_Bh_RLC_ChannelToAddModList_r16 := utils.NewSequence[*BH_RLC_ChannelConfig_r16]([]*BH_RLC_ChannelConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxBH_RLC_ChannelID_r16}, false)
				for _, i := range ie.Bh_RLC_ChannelToAddModList_r16 {
					tmp_Bh_RLC_ChannelToAddModList_r16.Value = append(tmp_Bh_RLC_ChannelToAddModList_r16.Value, &i)
				}
				if err = tmp_Bh_RLC_ChannelToAddModList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Bh_RLC_ChannelToAddModList_r16", err)
				}
			}
			// encode Bh_RLC_ChannelToReleaseList_r16 optional
			if len(ie.Bh_RLC_ChannelToReleaseList_r16) > 0 {
				tmp_Bh_RLC_ChannelToReleaseList_r16 := utils.NewSequence[*BH_RLC_ChannelID_r16]([]*BH_RLC_ChannelID_r16{}, uper.Constraint{Lb: 1, Ub: maxBH_RLC_ChannelID_r16}, false)
				for _, i := range ie.Bh_RLC_ChannelToReleaseList_r16 {
					tmp_Bh_RLC_ChannelToReleaseList_r16.Value = append(tmp_Bh_RLC_ChannelToReleaseList_r16.Value, &i)
				}
				if err = tmp_Bh_RLC_ChannelToReleaseList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Bh_RLC_ChannelToReleaseList_r16", err)
				}
			}
			// encode F1c_TransferPath_r16 optional
			if ie.F1c_TransferPath_r16 != nil {
				if err = ie.F1c_TransferPath_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode F1c_TransferPath_r16", err)
				}
			}
			// encode SimultaneousTCI_UpdateList1_r16 optional
			if len(ie.SimultaneousTCI_UpdateList1_r16) > 0 {
				tmp_SimultaneousTCI_UpdateList1_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				for _, i := range ie.SimultaneousTCI_UpdateList1_r16 {
					tmp_SimultaneousTCI_UpdateList1_r16.Value = append(tmp_SimultaneousTCI_UpdateList1_r16.Value, &i)
				}
				if err = tmp_SimultaneousTCI_UpdateList1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SimultaneousTCI_UpdateList1_r16", err)
				}
			}
			// encode SimultaneousTCI_UpdateList2_r16 optional
			if len(ie.SimultaneousTCI_UpdateList2_r16) > 0 {
				tmp_SimultaneousTCI_UpdateList2_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				for _, i := range ie.SimultaneousTCI_UpdateList2_r16 {
					tmp_SimultaneousTCI_UpdateList2_r16.Value = append(tmp_SimultaneousTCI_UpdateList2_r16.Value, &i)
				}
				if err = tmp_SimultaneousTCI_UpdateList2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SimultaneousTCI_UpdateList2_r16", err)
				}
			}
			// encode SimultaneousSpatial_UpdatedList1_r16 optional
			if len(ie.SimultaneousSpatial_UpdatedList1_r16) > 0 {
				tmp_SimultaneousSpatial_UpdatedList1_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				for _, i := range ie.SimultaneousSpatial_UpdatedList1_r16 {
					tmp_SimultaneousSpatial_UpdatedList1_r16.Value = append(tmp_SimultaneousSpatial_UpdatedList1_r16.Value, &i)
				}
				if err = tmp_SimultaneousSpatial_UpdatedList1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SimultaneousSpatial_UpdatedList1_r16", err)
				}
			}
			// encode SimultaneousSpatial_UpdatedList2_r16 optional
			if len(ie.SimultaneousSpatial_UpdatedList2_r16) > 0 {
				tmp_SimultaneousSpatial_UpdatedList2_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				for _, i := range ie.SimultaneousSpatial_UpdatedList2_r16 {
					tmp_SimultaneousSpatial_UpdatedList2_r16.Value = append(tmp_SimultaneousSpatial_UpdatedList2_r16.Value, &i)
				}
				if err = tmp_SimultaneousSpatial_UpdatedList2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SimultaneousSpatial_UpdatedList2_r16", err)
				}
			}
			// encode UplinkTxSwitchingOption_r16 optional
			if ie.UplinkTxSwitchingOption_r16 != nil {
				if err = ie.UplinkTxSwitchingOption_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UplinkTxSwitchingOption_r16", err)
				}
			}
			// encode UplinkTxSwitchingPowerBoosting_r16 optional
			if ie.UplinkTxSwitchingPowerBoosting_r16 != nil {
				if err = ie.UplinkTxSwitchingPowerBoosting_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UplinkTxSwitchingPowerBoosting_r16", err)
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
			optionals_ext_3 := []bool{ie.ReportUplinkTxDirectCurrentTwoCarrier_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ReportUplinkTxDirectCurrentTwoCarrier_r16 optional
			if ie.ReportUplinkTxDirectCurrentTwoCarrier_r16 != nil {
				if err = ie.ReportUplinkTxDirectCurrentTwoCarrier_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ReportUplinkTxDirectCurrentTwoCarrier_r16", err)
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
			optionals_ext_4 := []bool{ie.F1c_TransferPathNRDC_r17 != nil, ie.UplinkTxSwitching_2T_Mode_r17 != nil, ie.UplinkTxSwitching_DualUL_TxState_r17 != nil, len(ie.Uu_RelayRLC_ChannelToAddModList_r17) > 0, len(ie.Uu_RelayRLC_ChannelToReleaseList_r17) > 0, len(ie.SimultaneousU_TCI_UpdateList1_r17) > 0, len(ie.SimultaneousU_TCI_UpdateList2_r17) > 0, len(ie.SimultaneousU_TCI_UpdateList3_r17) > 0, len(ie.SimultaneousU_TCI_UpdateList4_r17) > 0, len(ie.Rlc_BearerToReleaseListExt_r17) > 0, len(ie.Iab_ResourceConfigToAddModList_r17) > 0, len(ie.Iab_ResourceConfigToReleaseList_r17) > 0}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode F1c_TransferPathNRDC_r17 optional
			if ie.F1c_TransferPathNRDC_r17 != nil {
				if err = ie.F1c_TransferPathNRDC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode F1c_TransferPathNRDC_r17", err)
				}
			}
			// encode UplinkTxSwitching_2T_Mode_r17 optional
			if ie.UplinkTxSwitching_2T_Mode_r17 != nil {
				if err = ie.UplinkTxSwitching_2T_Mode_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UplinkTxSwitching_2T_Mode_r17", err)
				}
			}
			// encode UplinkTxSwitching_DualUL_TxState_r17 optional
			if ie.UplinkTxSwitching_DualUL_TxState_r17 != nil {
				if err = ie.UplinkTxSwitching_DualUL_TxState_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UplinkTxSwitching_DualUL_TxState_r17", err)
				}
			}
			// encode Uu_RelayRLC_ChannelToAddModList_r17 optional
			if len(ie.Uu_RelayRLC_ChannelToAddModList_r17) > 0 {
				tmp_Uu_RelayRLC_ChannelToAddModList_r17 := utils.NewSequence[*Uu_RelayRLC_ChannelConfig_r17]([]*Uu_RelayRLC_ChannelConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxUu_RelayRLC_ChannelID_r17}, false)
				for _, i := range ie.Uu_RelayRLC_ChannelToAddModList_r17 {
					tmp_Uu_RelayRLC_ChannelToAddModList_r17.Value = append(tmp_Uu_RelayRLC_ChannelToAddModList_r17.Value, &i)
				}
				if err = tmp_Uu_RelayRLC_ChannelToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Uu_RelayRLC_ChannelToAddModList_r17", err)
				}
			}
			// encode Uu_RelayRLC_ChannelToReleaseList_r17 optional
			if len(ie.Uu_RelayRLC_ChannelToReleaseList_r17) > 0 {
				tmp_Uu_RelayRLC_ChannelToReleaseList_r17 := utils.NewSequence[*Uu_RelayRLC_ChannelID_r17]([]*Uu_RelayRLC_ChannelID_r17{}, uper.Constraint{Lb: 1, Ub: maxUu_RelayRLC_ChannelID_r17}, false)
				for _, i := range ie.Uu_RelayRLC_ChannelToReleaseList_r17 {
					tmp_Uu_RelayRLC_ChannelToReleaseList_r17.Value = append(tmp_Uu_RelayRLC_ChannelToReleaseList_r17.Value, &i)
				}
				if err = tmp_Uu_RelayRLC_ChannelToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Uu_RelayRLC_ChannelToReleaseList_r17", err)
				}
			}
			// encode SimultaneousU_TCI_UpdateList1_r17 optional
			if len(ie.SimultaneousU_TCI_UpdateList1_r17) > 0 {
				tmp_SimultaneousU_TCI_UpdateList1_r17 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				for _, i := range ie.SimultaneousU_TCI_UpdateList1_r17 {
					tmp_SimultaneousU_TCI_UpdateList1_r17.Value = append(tmp_SimultaneousU_TCI_UpdateList1_r17.Value, &i)
				}
				if err = tmp_SimultaneousU_TCI_UpdateList1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SimultaneousU_TCI_UpdateList1_r17", err)
				}
			}
			// encode SimultaneousU_TCI_UpdateList2_r17 optional
			if len(ie.SimultaneousU_TCI_UpdateList2_r17) > 0 {
				tmp_SimultaneousU_TCI_UpdateList2_r17 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				for _, i := range ie.SimultaneousU_TCI_UpdateList2_r17 {
					tmp_SimultaneousU_TCI_UpdateList2_r17.Value = append(tmp_SimultaneousU_TCI_UpdateList2_r17.Value, &i)
				}
				if err = tmp_SimultaneousU_TCI_UpdateList2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SimultaneousU_TCI_UpdateList2_r17", err)
				}
			}
			// encode SimultaneousU_TCI_UpdateList3_r17 optional
			if len(ie.SimultaneousU_TCI_UpdateList3_r17) > 0 {
				tmp_SimultaneousU_TCI_UpdateList3_r17 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				for _, i := range ie.SimultaneousU_TCI_UpdateList3_r17 {
					tmp_SimultaneousU_TCI_UpdateList3_r17.Value = append(tmp_SimultaneousU_TCI_UpdateList3_r17.Value, &i)
				}
				if err = tmp_SimultaneousU_TCI_UpdateList3_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SimultaneousU_TCI_UpdateList3_r17", err)
				}
			}
			// encode SimultaneousU_TCI_UpdateList4_r17 optional
			if len(ie.SimultaneousU_TCI_UpdateList4_r17) > 0 {
				tmp_SimultaneousU_TCI_UpdateList4_r17 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				for _, i := range ie.SimultaneousU_TCI_UpdateList4_r17 {
					tmp_SimultaneousU_TCI_UpdateList4_r17.Value = append(tmp_SimultaneousU_TCI_UpdateList4_r17.Value, &i)
				}
				if err = tmp_SimultaneousU_TCI_UpdateList4_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SimultaneousU_TCI_UpdateList4_r17", err)
				}
			}
			// encode Rlc_BearerToReleaseListExt_r17 optional
			if len(ie.Rlc_BearerToReleaseListExt_r17) > 0 {
				tmp_Rlc_BearerToReleaseListExt_r17 := utils.NewSequence[*LogicalChannelIdentityExt_r17]([]*LogicalChannelIdentityExt_r17{}, uper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
				for _, i := range ie.Rlc_BearerToReleaseListExt_r17 {
					tmp_Rlc_BearerToReleaseListExt_r17.Value = append(tmp_Rlc_BearerToReleaseListExt_r17.Value, &i)
				}
				if err = tmp_Rlc_BearerToReleaseListExt_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Rlc_BearerToReleaseListExt_r17", err)
				}
			}
			// encode Iab_ResourceConfigToAddModList_r17 optional
			if len(ie.Iab_ResourceConfigToAddModList_r17) > 0 {
				tmp_Iab_ResourceConfigToAddModList_r17 := utils.NewSequence[*IAB_ResourceConfig_r17]([]*IAB_ResourceConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofIABResourceConfig_r17}, false)
				for _, i := range ie.Iab_ResourceConfigToAddModList_r17 {
					tmp_Iab_ResourceConfigToAddModList_r17.Value = append(tmp_Iab_ResourceConfigToAddModList_r17.Value, &i)
				}
				if err = tmp_Iab_ResourceConfigToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Iab_ResourceConfigToAddModList_r17", err)
				}
			}
			// encode Iab_ResourceConfigToReleaseList_r17 optional
			if len(ie.Iab_ResourceConfigToReleaseList_r17) > 0 {
				tmp_Iab_ResourceConfigToReleaseList_r17 := utils.NewSequence[*IAB_ResourceConfigID_r17]([]*IAB_ResourceConfigID_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofIABResourceConfig_r17}, false)
				for _, i := range ie.Iab_ResourceConfigToReleaseList_r17 {
					tmp_Iab_ResourceConfigToReleaseList_r17.Value = append(tmp_Iab_ResourceConfigToReleaseList_r17.Value, &i)
				}
				if err = tmp_Iab_ResourceConfigToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Iab_ResourceConfigToReleaseList_r17", err)
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
			optionals_ext_5 := []bool{ie.ReportUplinkTxDirectCurrentMoreCarrier_r17 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ReportUplinkTxDirectCurrentMoreCarrier_r17 optional
			if ie.ReportUplinkTxDirectCurrentMoreCarrier_r17 != nil {
				if err = ie.ReportUplinkTxDirectCurrentMoreCarrier_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ReportUplinkTxDirectCurrentMoreCarrier_r17", err)
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

func (ie *CellGroupConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Rlc_BearerToAddModListPresent bool
	if Rlc_BearerToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Rlc_BearerToReleaseListPresent bool
	if Rlc_BearerToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Mac_CellGroupConfigPresent bool
	if Mac_CellGroupConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var PhysicalCellGroupConfigPresent bool
	if PhysicalCellGroupConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SpCellConfigPresent bool
	if SpCellConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SCellToAddModListPresent bool
	if SCellToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SCellToReleaseListPresent bool
	if SCellToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.CellGroupId.Decode(r); err != nil {
		return utils.WrapError("Decode CellGroupId", err)
	}
	if Rlc_BearerToAddModListPresent {
		tmp_Rlc_BearerToAddModList := utils.NewSequence[*RLC_BearerConfig]([]*RLC_BearerConfig{}, uper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
		fn_Rlc_BearerToAddModList := func() *RLC_BearerConfig {
			return new(RLC_BearerConfig)
		}
		if err = tmp_Rlc_BearerToAddModList.Decode(r, fn_Rlc_BearerToAddModList); err != nil {
			return utils.WrapError("Decode Rlc_BearerToAddModList", err)
		}
		ie.Rlc_BearerToAddModList = []RLC_BearerConfig{}
		for _, i := range tmp_Rlc_BearerToAddModList.Value {
			ie.Rlc_BearerToAddModList = append(ie.Rlc_BearerToAddModList, *i)
		}
	}
	if Rlc_BearerToReleaseListPresent {
		tmp_Rlc_BearerToReleaseList := utils.NewSequence[*LogicalChannelIdentity]([]*LogicalChannelIdentity{}, uper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
		fn_Rlc_BearerToReleaseList := func() *LogicalChannelIdentity {
			return new(LogicalChannelIdentity)
		}
		if err = tmp_Rlc_BearerToReleaseList.Decode(r, fn_Rlc_BearerToReleaseList); err != nil {
			return utils.WrapError("Decode Rlc_BearerToReleaseList", err)
		}
		ie.Rlc_BearerToReleaseList = []LogicalChannelIdentity{}
		for _, i := range tmp_Rlc_BearerToReleaseList.Value {
			ie.Rlc_BearerToReleaseList = append(ie.Rlc_BearerToReleaseList, *i)
		}
	}
	if Mac_CellGroupConfigPresent {
		ie.Mac_CellGroupConfig = new(MAC_CellGroupConfig)
		if err = ie.Mac_CellGroupConfig.Decode(r); err != nil {
			return utils.WrapError("Decode Mac_CellGroupConfig", err)
		}
	}
	if PhysicalCellGroupConfigPresent {
		ie.PhysicalCellGroupConfig = new(PhysicalCellGroupConfig)
		if err = ie.PhysicalCellGroupConfig.Decode(r); err != nil {
			return utils.WrapError("Decode PhysicalCellGroupConfig", err)
		}
	}
	if SpCellConfigPresent {
		ie.SpCellConfig = new(SpCellConfig)
		if err = ie.SpCellConfig.Decode(r); err != nil {
			return utils.WrapError("Decode SpCellConfig", err)
		}
	}
	if SCellToAddModListPresent {
		tmp_SCellToAddModList := utils.NewSequence[*SCellConfig]([]*SCellConfig{}, uper.Constraint{Lb: 1, Ub: maxNrofSCells}, false)
		fn_SCellToAddModList := func() *SCellConfig {
			return new(SCellConfig)
		}
		if err = tmp_SCellToAddModList.Decode(r, fn_SCellToAddModList); err != nil {
			return utils.WrapError("Decode SCellToAddModList", err)
		}
		ie.SCellToAddModList = []SCellConfig{}
		for _, i := range tmp_SCellToAddModList.Value {
			ie.SCellToAddModList = append(ie.SCellToAddModList, *i)
		}
	}
	if SCellToReleaseListPresent {
		tmp_SCellToReleaseList := utils.NewSequence[*SCellIndex]([]*SCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofSCells}, false)
		fn_SCellToReleaseList := func() *SCellIndex {
			return new(SCellIndex)
		}
		if err = tmp_SCellToReleaseList.Decode(r, fn_SCellToReleaseList); err != nil {
			return utils.WrapError("Decode SCellToReleaseList", err)
		}
		ie.SCellToReleaseList = []SCellIndex{}
		for _, i := range tmp_SCellToReleaseList.Value {
			ie.SCellToReleaseList = append(ie.SCellToReleaseList, *i)
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

			ReportUplinkTxDirectCurrentPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ReportUplinkTxDirectCurrent optional
			if ReportUplinkTxDirectCurrentPresent {
				ie.ReportUplinkTxDirectCurrent = new(CellGroupConfig_reportUplinkTxDirectCurrent)
				if err = ie.ReportUplinkTxDirectCurrent.Decode(extReader); err != nil {
					return utils.WrapError("Decode ReportUplinkTxDirectCurrent", err)
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

			Bap_Address_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Bh_RLC_ChannelToAddModList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Bh_RLC_ChannelToReleaseList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			F1c_TransferPath_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SimultaneousTCI_UpdateList1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SimultaneousTCI_UpdateList2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SimultaneousSpatial_UpdatedList1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SimultaneousSpatial_UpdatedList2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UplinkTxSwitchingOption_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UplinkTxSwitchingPowerBoosting_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Bap_Address_r16 optional
			if Bap_Address_r16Present {
				var tmp_bs_Bap_Address_r16 []byte
				var n_Bap_Address_r16 uint
				if tmp_bs_Bap_Address_r16, n_Bap_Address_r16, err = extReader.ReadBitString(&uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
					return utils.WrapError("Decode Bap_Address_r16", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_Bap_Address_r16,
					NumBits: uint64(n_Bap_Address_r16),
				}
				ie.Bap_Address_r16 = &tmp_bitstring
			}
			// decode Bh_RLC_ChannelToAddModList_r16 optional
			if Bh_RLC_ChannelToAddModList_r16Present {
				tmp_Bh_RLC_ChannelToAddModList_r16 := utils.NewSequence[*BH_RLC_ChannelConfig_r16]([]*BH_RLC_ChannelConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxBH_RLC_ChannelID_r16}, false)
				fn_Bh_RLC_ChannelToAddModList_r16 := func() *BH_RLC_ChannelConfig_r16 {
					return new(BH_RLC_ChannelConfig_r16)
				}
				if err = tmp_Bh_RLC_ChannelToAddModList_r16.Decode(extReader, fn_Bh_RLC_ChannelToAddModList_r16); err != nil {
					return utils.WrapError("Decode Bh_RLC_ChannelToAddModList_r16", err)
				}
				ie.Bh_RLC_ChannelToAddModList_r16 = []BH_RLC_ChannelConfig_r16{}
				for _, i := range tmp_Bh_RLC_ChannelToAddModList_r16.Value {
					ie.Bh_RLC_ChannelToAddModList_r16 = append(ie.Bh_RLC_ChannelToAddModList_r16, *i)
				}
			}
			// decode Bh_RLC_ChannelToReleaseList_r16 optional
			if Bh_RLC_ChannelToReleaseList_r16Present {
				tmp_Bh_RLC_ChannelToReleaseList_r16 := utils.NewSequence[*BH_RLC_ChannelID_r16]([]*BH_RLC_ChannelID_r16{}, uper.Constraint{Lb: 1, Ub: maxBH_RLC_ChannelID_r16}, false)
				fn_Bh_RLC_ChannelToReleaseList_r16 := func() *BH_RLC_ChannelID_r16 {
					return new(BH_RLC_ChannelID_r16)
				}
				if err = tmp_Bh_RLC_ChannelToReleaseList_r16.Decode(extReader, fn_Bh_RLC_ChannelToReleaseList_r16); err != nil {
					return utils.WrapError("Decode Bh_RLC_ChannelToReleaseList_r16", err)
				}
				ie.Bh_RLC_ChannelToReleaseList_r16 = []BH_RLC_ChannelID_r16{}
				for _, i := range tmp_Bh_RLC_ChannelToReleaseList_r16.Value {
					ie.Bh_RLC_ChannelToReleaseList_r16 = append(ie.Bh_RLC_ChannelToReleaseList_r16, *i)
				}
			}
			// decode F1c_TransferPath_r16 optional
			if F1c_TransferPath_r16Present {
				ie.F1c_TransferPath_r16 = new(CellGroupConfig_f1c_TransferPath_r16)
				if err = ie.F1c_TransferPath_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode F1c_TransferPath_r16", err)
				}
			}
			// decode SimultaneousTCI_UpdateList1_r16 optional
			if SimultaneousTCI_UpdateList1_r16Present {
				tmp_SimultaneousTCI_UpdateList1_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				fn_SimultaneousTCI_UpdateList1_r16 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_SimultaneousTCI_UpdateList1_r16.Decode(extReader, fn_SimultaneousTCI_UpdateList1_r16); err != nil {
					return utils.WrapError("Decode SimultaneousTCI_UpdateList1_r16", err)
				}
				ie.SimultaneousTCI_UpdateList1_r16 = []ServCellIndex{}
				for _, i := range tmp_SimultaneousTCI_UpdateList1_r16.Value {
					ie.SimultaneousTCI_UpdateList1_r16 = append(ie.SimultaneousTCI_UpdateList1_r16, *i)
				}
			}
			// decode SimultaneousTCI_UpdateList2_r16 optional
			if SimultaneousTCI_UpdateList2_r16Present {
				tmp_SimultaneousTCI_UpdateList2_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				fn_SimultaneousTCI_UpdateList2_r16 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_SimultaneousTCI_UpdateList2_r16.Decode(extReader, fn_SimultaneousTCI_UpdateList2_r16); err != nil {
					return utils.WrapError("Decode SimultaneousTCI_UpdateList2_r16", err)
				}
				ie.SimultaneousTCI_UpdateList2_r16 = []ServCellIndex{}
				for _, i := range tmp_SimultaneousTCI_UpdateList2_r16.Value {
					ie.SimultaneousTCI_UpdateList2_r16 = append(ie.SimultaneousTCI_UpdateList2_r16, *i)
				}
			}
			// decode SimultaneousSpatial_UpdatedList1_r16 optional
			if SimultaneousSpatial_UpdatedList1_r16Present {
				tmp_SimultaneousSpatial_UpdatedList1_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				fn_SimultaneousSpatial_UpdatedList1_r16 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_SimultaneousSpatial_UpdatedList1_r16.Decode(extReader, fn_SimultaneousSpatial_UpdatedList1_r16); err != nil {
					return utils.WrapError("Decode SimultaneousSpatial_UpdatedList1_r16", err)
				}
				ie.SimultaneousSpatial_UpdatedList1_r16 = []ServCellIndex{}
				for _, i := range tmp_SimultaneousSpatial_UpdatedList1_r16.Value {
					ie.SimultaneousSpatial_UpdatedList1_r16 = append(ie.SimultaneousSpatial_UpdatedList1_r16, *i)
				}
			}
			// decode SimultaneousSpatial_UpdatedList2_r16 optional
			if SimultaneousSpatial_UpdatedList2_r16Present {
				tmp_SimultaneousSpatial_UpdatedList2_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				fn_SimultaneousSpatial_UpdatedList2_r16 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_SimultaneousSpatial_UpdatedList2_r16.Decode(extReader, fn_SimultaneousSpatial_UpdatedList2_r16); err != nil {
					return utils.WrapError("Decode SimultaneousSpatial_UpdatedList2_r16", err)
				}
				ie.SimultaneousSpatial_UpdatedList2_r16 = []ServCellIndex{}
				for _, i := range tmp_SimultaneousSpatial_UpdatedList2_r16.Value {
					ie.SimultaneousSpatial_UpdatedList2_r16 = append(ie.SimultaneousSpatial_UpdatedList2_r16, *i)
				}
			}
			// decode UplinkTxSwitchingOption_r16 optional
			if UplinkTxSwitchingOption_r16Present {
				ie.UplinkTxSwitchingOption_r16 = new(CellGroupConfig_uplinkTxSwitchingOption_r16)
				if err = ie.UplinkTxSwitchingOption_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode UplinkTxSwitchingOption_r16", err)
				}
			}
			// decode UplinkTxSwitchingPowerBoosting_r16 optional
			if UplinkTxSwitchingPowerBoosting_r16Present {
				ie.UplinkTxSwitchingPowerBoosting_r16 = new(CellGroupConfig_uplinkTxSwitchingPowerBoosting_r16)
				if err = ie.UplinkTxSwitchingPowerBoosting_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode UplinkTxSwitchingPowerBoosting_r16", err)
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

			ReportUplinkTxDirectCurrentTwoCarrier_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ReportUplinkTxDirectCurrentTwoCarrier_r16 optional
			if ReportUplinkTxDirectCurrentTwoCarrier_r16Present {
				ie.ReportUplinkTxDirectCurrentTwoCarrier_r16 = new(CellGroupConfig_reportUplinkTxDirectCurrentTwoCarrier_r16)
				if err = ie.ReportUplinkTxDirectCurrentTwoCarrier_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ReportUplinkTxDirectCurrentTwoCarrier_r16", err)
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

			F1c_TransferPathNRDC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UplinkTxSwitching_2T_Mode_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UplinkTxSwitching_DualUL_TxState_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Uu_RelayRLC_ChannelToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Uu_RelayRLC_ChannelToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SimultaneousU_TCI_UpdateList1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SimultaneousU_TCI_UpdateList2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SimultaneousU_TCI_UpdateList3_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SimultaneousU_TCI_UpdateList4_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Rlc_BearerToReleaseListExt_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Iab_ResourceConfigToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Iab_ResourceConfigToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode F1c_TransferPathNRDC_r17 optional
			if F1c_TransferPathNRDC_r17Present {
				ie.F1c_TransferPathNRDC_r17 = new(CellGroupConfig_f1c_TransferPathNRDC_r17)
				if err = ie.F1c_TransferPathNRDC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode F1c_TransferPathNRDC_r17", err)
				}
			}
			// decode UplinkTxSwitching_2T_Mode_r17 optional
			if UplinkTxSwitching_2T_Mode_r17Present {
				ie.UplinkTxSwitching_2T_Mode_r17 = new(CellGroupConfig_uplinkTxSwitching_2T_Mode_r17)
				if err = ie.UplinkTxSwitching_2T_Mode_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UplinkTxSwitching_2T_Mode_r17", err)
				}
			}
			// decode UplinkTxSwitching_DualUL_TxState_r17 optional
			if UplinkTxSwitching_DualUL_TxState_r17Present {
				ie.UplinkTxSwitching_DualUL_TxState_r17 = new(CellGroupConfig_uplinkTxSwitching_DualUL_TxState_r17)
				if err = ie.UplinkTxSwitching_DualUL_TxState_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode UplinkTxSwitching_DualUL_TxState_r17", err)
				}
			}
			// decode Uu_RelayRLC_ChannelToAddModList_r17 optional
			if Uu_RelayRLC_ChannelToAddModList_r17Present {
				tmp_Uu_RelayRLC_ChannelToAddModList_r17 := utils.NewSequence[*Uu_RelayRLC_ChannelConfig_r17]([]*Uu_RelayRLC_ChannelConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxUu_RelayRLC_ChannelID_r17}, false)
				fn_Uu_RelayRLC_ChannelToAddModList_r17 := func() *Uu_RelayRLC_ChannelConfig_r17 {
					return new(Uu_RelayRLC_ChannelConfig_r17)
				}
				if err = tmp_Uu_RelayRLC_ChannelToAddModList_r17.Decode(extReader, fn_Uu_RelayRLC_ChannelToAddModList_r17); err != nil {
					return utils.WrapError("Decode Uu_RelayRLC_ChannelToAddModList_r17", err)
				}
				ie.Uu_RelayRLC_ChannelToAddModList_r17 = []Uu_RelayRLC_ChannelConfig_r17{}
				for _, i := range tmp_Uu_RelayRLC_ChannelToAddModList_r17.Value {
					ie.Uu_RelayRLC_ChannelToAddModList_r17 = append(ie.Uu_RelayRLC_ChannelToAddModList_r17, *i)
				}
			}
			// decode Uu_RelayRLC_ChannelToReleaseList_r17 optional
			if Uu_RelayRLC_ChannelToReleaseList_r17Present {
				tmp_Uu_RelayRLC_ChannelToReleaseList_r17 := utils.NewSequence[*Uu_RelayRLC_ChannelID_r17]([]*Uu_RelayRLC_ChannelID_r17{}, uper.Constraint{Lb: 1, Ub: maxUu_RelayRLC_ChannelID_r17}, false)
				fn_Uu_RelayRLC_ChannelToReleaseList_r17 := func() *Uu_RelayRLC_ChannelID_r17 {
					return new(Uu_RelayRLC_ChannelID_r17)
				}
				if err = tmp_Uu_RelayRLC_ChannelToReleaseList_r17.Decode(extReader, fn_Uu_RelayRLC_ChannelToReleaseList_r17); err != nil {
					return utils.WrapError("Decode Uu_RelayRLC_ChannelToReleaseList_r17", err)
				}
				ie.Uu_RelayRLC_ChannelToReleaseList_r17 = []Uu_RelayRLC_ChannelID_r17{}
				for _, i := range tmp_Uu_RelayRLC_ChannelToReleaseList_r17.Value {
					ie.Uu_RelayRLC_ChannelToReleaseList_r17 = append(ie.Uu_RelayRLC_ChannelToReleaseList_r17, *i)
				}
			}
			// decode SimultaneousU_TCI_UpdateList1_r17 optional
			if SimultaneousU_TCI_UpdateList1_r17Present {
				tmp_SimultaneousU_TCI_UpdateList1_r17 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				fn_SimultaneousU_TCI_UpdateList1_r17 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_SimultaneousU_TCI_UpdateList1_r17.Decode(extReader, fn_SimultaneousU_TCI_UpdateList1_r17); err != nil {
					return utils.WrapError("Decode SimultaneousU_TCI_UpdateList1_r17", err)
				}
				ie.SimultaneousU_TCI_UpdateList1_r17 = []ServCellIndex{}
				for _, i := range tmp_SimultaneousU_TCI_UpdateList1_r17.Value {
					ie.SimultaneousU_TCI_UpdateList1_r17 = append(ie.SimultaneousU_TCI_UpdateList1_r17, *i)
				}
			}
			// decode SimultaneousU_TCI_UpdateList2_r17 optional
			if SimultaneousU_TCI_UpdateList2_r17Present {
				tmp_SimultaneousU_TCI_UpdateList2_r17 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				fn_SimultaneousU_TCI_UpdateList2_r17 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_SimultaneousU_TCI_UpdateList2_r17.Decode(extReader, fn_SimultaneousU_TCI_UpdateList2_r17); err != nil {
					return utils.WrapError("Decode SimultaneousU_TCI_UpdateList2_r17", err)
				}
				ie.SimultaneousU_TCI_UpdateList2_r17 = []ServCellIndex{}
				for _, i := range tmp_SimultaneousU_TCI_UpdateList2_r17.Value {
					ie.SimultaneousU_TCI_UpdateList2_r17 = append(ie.SimultaneousU_TCI_UpdateList2_r17, *i)
				}
			}
			// decode SimultaneousU_TCI_UpdateList3_r17 optional
			if SimultaneousU_TCI_UpdateList3_r17Present {
				tmp_SimultaneousU_TCI_UpdateList3_r17 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				fn_SimultaneousU_TCI_UpdateList3_r17 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_SimultaneousU_TCI_UpdateList3_r17.Decode(extReader, fn_SimultaneousU_TCI_UpdateList3_r17); err != nil {
					return utils.WrapError("Decode SimultaneousU_TCI_UpdateList3_r17", err)
				}
				ie.SimultaneousU_TCI_UpdateList3_r17 = []ServCellIndex{}
				for _, i := range tmp_SimultaneousU_TCI_UpdateList3_r17.Value {
					ie.SimultaneousU_TCI_UpdateList3_r17 = append(ie.SimultaneousU_TCI_UpdateList3_r17, *i)
				}
			}
			// decode SimultaneousU_TCI_UpdateList4_r17 optional
			if SimultaneousU_TCI_UpdateList4_r17Present {
				tmp_SimultaneousU_TCI_UpdateList4_r17 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, uper.Constraint{Lb: 1, Ub: maxNrofServingCellsTCI_r16}, false)
				fn_SimultaneousU_TCI_UpdateList4_r17 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_SimultaneousU_TCI_UpdateList4_r17.Decode(extReader, fn_SimultaneousU_TCI_UpdateList4_r17); err != nil {
					return utils.WrapError("Decode SimultaneousU_TCI_UpdateList4_r17", err)
				}
				ie.SimultaneousU_TCI_UpdateList4_r17 = []ServCellIndex{}
				for _, i := range tmp_SimultaneousU_TCI_UpdateList4_r17.Value {
					ie.SimultaneousU_TCI_UpdateList4_r17 = append(ie.SimultaneousU_TCI_UpdateList4_r17, *i)
				}
			}
			// decode Rlc_BearerToReleaseListExt_r17 optional
			if Rlc_BearerToReleaseListExt_r17Present {
				tmp_Rlc_BearerToReleaseListExt_r17 := utils.NewSequence[*LogicalChannelIdentityExt_r17]([]*LogicalChannelIdentityExt_r17{}, uper.Constraint{Lb: 1, Ub: maxLC_ID}, false)
				fn_Rlc_BearerToReleaseListExt_r17 := func() *LogicalChannelIdentityExt_r17 {
					return new(LogicalChannelIdentityExt_r17)
				}
				if err = tmp_Rlc_BearerToReleaseListExt_r17.Decode(extReader, fn_Rlc_BearerToReleaseListExt_r17); err != nil {
					return utils.WrapError("Decode Rlc_BearerToReleaseListExt_r17", err)
				}
				ie.Rlc_BearerToReleaseListExt_r17 = []LogicalChannelIdentityExt_r17{}
				for _, i := range tmp_Rlc_BearerToReleaseListExt_r17.Value {
					ie.Rlc_BearerToReleaseListExt_r17 = append(ie.Rlc_BearerToReleaseListExt_r17, *i)
				}
			}
			// decode Iab_ResourceConfigToAddModList_r17 optional
			if Iab_ResourceConfigToAddModList_r17Present {
				tmp_Iab_ResourceConfigToAddModList_r17 := utils.NewSequence[*IAB_ResourceConfig_r17]([]*IAB_ResourceConfig_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofIABResourceConfig_r17}, false)
				fn_Iab_ResourceConfigToAddModList_r17 := func() *IAB_ResourceConfig_r17 {
					return new(IAB_ResourceConfig_r17)
				}
				if err = tmp_Iab_ResourceConfigToAddModList_r17.Decode(extReader, fn_Iab_ResourceConfigToAddModList_r17); err != nil {
					return utils.WrapError("Decode Iab_ResourceConfigToAddModList_r17", err)
				}
				ie.Iab_ResourceConfigToAddModList_r17 = []IAB_ResourceConfig_r17{}
				for _, i := range tmp_Iab_ResourceConfigToAddModList_r17.Value {
					ie.Iab_ResourceConfigToAddModList_r17 = append(ie.Iab_ResourceConfigToAddModList_r17, *i)
				}
			}
			// decode Iab_ResourceConfigToReleaseList_r17 optional
			if Iab_ResourceConfigToReleaseList_r17Present {
				tmp_Iab_ResourceConfigToReleaseList_r17 := utils.NewSequence[*IAB_ResourceConfigID_r17]([]*IAB_ResourceConfigID_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofIABResourceConfig_r17}, false)
				fn_Iab_ResourceConfigToReleaseList_r17 := func() *IAB_ResourceConfigID_r17 {
					return new(IAB_ResourceConfigID_r17)
				}
				if err = tmp_Iab_ResourceConfigToReleaseList_r17.Decode(extReader, fn_Iab_ResourceConfigToReleaseList_r17); err != nil {
					return utils.WrapError("Decode Iab_ResourceConfigToReleaseList_r17", err)
				}
				ie.Iab_ResourceConfigToReleaseList_r17 = []IAB_ResourceConfigID_r17{}
				for _, i := range tmp_Iab_ResourceConfigToReleaseList_r17.Value {
					ie.Iab_ResourceConfigToReleaseList_r17 = append(ie.Iab_ResourceConfigToReleaseList_r17, *i)
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

			ReportUplinkTxDirectCurrentMoreCarrier_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ReportUplinkTxDirectCurrentMoreCarrier_r17 optional
			if ReportUplinkTxDirectCurrentMoreCarrier_r17Present {
				ie.ReportUplinkTxDirectCurrentMoreCarrier_r17 = new(ReportUplinkTxDirectCurrentMoreCarrier_r17)
				if err = ie.ReportUplinkTxDirectCurrentMoreCarrier_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ReportUplinkTxDirectCurrentMoreCarrier_r17", err)
				}
			}
		}
	}
	return nil
}
