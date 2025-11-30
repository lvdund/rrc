package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RLF_Report_r16_nr_RLF_Report_r16 struct {
	MeasResultLastServCell_r16     MeasResultRLFNR_r16                                        `madatory`
	MeasResultNeighCells_r16       *RLF_Report_r16_nr_RLF_Report_r16_measResultNeighCells_r16 `optional`
	C_RNTI_r16                     RNTI_Value                                                 `madatory`
	PreviousPCellId_r16            *RLF_Report_r16_nr_RLF_Report_r16_previousPCellId_r16      `optional`
	FailedPCellId_r16              RLF_Report_r16_nr_RLF_Report_r16_failedPCellId_r16         `madatory`
	ReconnectCellId_r16            *RLF_Report_r16_nr_RLF_Report_r16_reconnectCellId_r16      `optional`
	TimeUntilReconnection_r16      *TimeUntilReconnection_r16                                 `optional`
	ReestablishmentCellId_r16      *CGI_Info_Logging_r16                                      `optional`
	TimeConnFailure_r16            *int64                                                     `lb:0,ub:1023,optional`
	TimeSinceFailure_r16           TimeSinceFailure_r16                                       `madatory`
	ConnectionFailureType_r16      RLF_Report_r16_nr_RLF_Report_r16_connectionFailureType_r16 `madatory`
	Rlf_Cause_r16                  RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16             `madatory`
	LocationInfo_r16               *LocationInfo_r16                                          `optional`
	NoSuitableCellFound_r16        *RLF_Report_r16_nr_RLF_Report_r16_noSuitableCellFound_r16  `optional`
	Ra_InformationCommon_r16       *RA_InformationCommon_r16                                  `optional`
	Csi_rsRLMConfigBitmap_v1650    *aper.BitString                                            `lb:96,ub:96,optional`
	LastHO_Type_r17                *RLF_Report_r16_nr_RLF_Report_r16_lastHO_Type_r17          `optional`
	TimeConnSourceDAPS_Failure_r17 *TimeConnSourceDAPS_Failure_r17                            `optional`
	TimeSinceCHO_Reconfig_r17      *TimeSinceCHO_Reconfig_r17                                 `optional`
	ChoCellId_r17                  *RLF_Report_r16_nr_RLF_Report_r16_choCellId_r17            `optional`
	ChoCandidateCellList_r17       *ChoCandidateCellList_r17                                  `optional`
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasResultNeighCells_r16 != nil, ie.PreviousPCellId_r16 != nil, ie.ReconnectCellId_r16 != nil, ie.TimeUntilReconnection_r16 != nil, ie.ReestablishmentCellId_r16 != nil, ie.TimeConnFailure_r16 != nil, ie.LocationInfo_r16 != nil, ie.NoSuitableCellFound_r16 != nil, ie.Ra_InformationCommon_r16 != nil, ie.Csi_rsRLMConfigBitmap_v1650 != nil, ie.LastHO_Type_r17 != nil, ie.TimeConnSourceDAPS_Failure_r17 != nil, ie.TimeSinceCHO_Reconfig_r17 != nil, ie.ChoCellId_r17 != nil, ie.ChoCandidateCellList_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.MeasResultLastServCell_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultLastServCell_r16", err)
	}
	if ie.MeasResultNeighCells_r16 != nil {
		if err = ie.MeasResultNeighCells_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultNeighCells_r16", err)
		}
	}
	if err = ie.C_RNTI_r16.Encode(w); err != nil {
		return utils.WrapError("Encode C_RNTI_r16", err)
	}
	if ie.PreviousPCellId_r16 != nil {
		if err = ie.PreviousPCellId_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PreviousPCellId_r16", err)
		}
	}
	if err = ie.FailedPCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode FailedPCellId_r16", err)
	}
	if ie.ReconnectCellId_r16 != nil {
		if err = ie.ReconnectCellId_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ReconnectCellId_r16", err)
		}
	}
	if ie.TimeUntilReconnection_r16 != nil {
		if err = ie.TimeUntilReconnection_r16.Encode(w); err != nil {
			return utils.WrapError("Encode TimeUntilReconnection_r16", err)
		}
	}
	if ie.ReestablishmentCellId_r16 != nil {
		if err = ie.ReestablishmentCellId_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ReestablishmentCellId_r16", err)
		}
	}
	if ie.TimeConnFailure_r16 != nil {
		if err = w.WriteInteger(*ie.TimeConnFailure_r16, &aper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode TimeConnFailure_r16", err)
		}
	}
	if err = ie.TimeSinceFailure_r16.Encode(w); err != nil {
		return utils.WrapError("Encode TimeSinceFailure_r16", err)
	}
	if err = ie.ConnectionFailureType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ConnectionFailureType_r16", err)
	}
	if err = ie.Rlf_Cause_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Rlf_Cause_r16", err)
	}
	if ie.LocationInfo_r16 != nil {
		if err = ie.LocationInfo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode LocationInfo_r16", err)
		}
	}
	if ie.NoSuitableCellFound_r16 != nil {
		if err = ie.NoSuitableCellFound_r16.Encode(w); err != nil {
			return utils.WrapError("Encode NoSuitableCellFound_r16", err)
		}
	}
	if ie.Ra_InformationCommon_r16 != nil {
		if err = ie.Ra_InformationCommon_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ra_InformationCommon_r16", err)
		}
	}
	if ie.Csi_rsRLMConfigBitmap_v1650 != nil {
		if err = w.WriteBitString(ie.Csi_rsRLMConfigBitmap_v1650.Bytes, uint(ie.Csi_rsRLMConfigBitmap_v1650.NumBits), &aper.Constraint{Lb: 96, Ub: 96}, false); err != nil {
			return utils.WrapError("Encode Csi_rsRLMConfigBitmap_v1650", err)
		}
	}
	if ie.LastHO_Type_r17 != nil {
		if err = ie.LastHO_Type_r17.Encode(w); err != nil {
			return utils.WrapError("Encode LastHO_Type_r17", err)
		}
	}
	if ie.TimeConnSourceDAPS_Failure_r17 != nil {
		if err = ie.TimeConnSourceDAPS_Failure_r17.Encode(w); err != nil {
			return utils.WrapError("Encode TimeConnSourceDAPS_Failure_r17", err)
		}
	}
	if ie.TimeSinceCHO_Reconfig_r17 != nil {
		if err = ie.TimeSinceCHO_Reconfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode TimeSinceCHO_Reconfig_r17", err)
		}
	}
	if ie.ChoCellId_r17 != nil {
		if err = ie.ChoCellId_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ChoCellId_r17", err)
		}
	}
	if ie.ChoCandidateCellList_r17 != nil {
		if err = ie.ChoCandidateCellList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ChoCandidateCellList_r17", err)
		}
	}
	return nil
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16) Decode(r *aper.AperReader) error {
	var err error
	var MeasResultNeighCells_r16Present bool
	if MeasResultNeighCells_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PreviousPCellId_r16Present bool
	if PreviousPCellId_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReconnectCellId_r16Present bool
	if ReconnectCellId_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TimeUntilReconnection_r16Present bool
	if TimeUntilReconnection_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReestablishmentCellId_r16Present bool
	if ReestablishmentCellId_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TimeConnFailure_r16Present bool
	if TimeConnFailure_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LocationInfo_r16Present bool
	if LocationInfo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NoSuitableCellFound_r16Present bool
	if NoSuitableCellFound_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ra_InformationCommon_r16Present bool
	if Ra_InformationCommon_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Csi_rsRLMConfigBitmap_v1650Present bool
	if Csi_rsRLMConfigBitmap_v1650Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LastHO_Type_r17Present bool
	if LastHO_Type_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TimeConnSourceDAPS_Failure_r17Present bool
	if TimeConnSourceDAPS_Failure_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TimeSinceCHO_Reconfig_r17Present bool
	if TimeSinceCHO_Reconfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ChoCellId_r17Present bool
	if ChoCellId_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ChoCandidateCellList_r17Present bool
	if ChoCandidateCellList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.MeasResultLastServCell_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MeasResultLastServCell_r16", err)
	}
	if MeasResultNeighCells_r16Present {
		ie.MeasResultNeighCells_r16 = new(RLF_Report_r16_nr_RLF_Report_r16_measResultNeighCells_r16)
		if err = ie.MeasResultNeighCells_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultNeighCells_r16", err)
		}
	}
	if err = ie.C_RNTI_r16.Decode(r); err != nil {
		return utils.WrapError("Decode C_RNTI_r16", err)
	}
	if PreviousPCellId_r16Present {
		ie.PreviousPCellId_r16 = new(RLF_Report_r16_nr_RLF_Report_r16_previousPCellId_r16)
		if err = ie.PreviousPCellId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PreviousPCellId_r16", err)
		}
	}
	if err = ie.FailedPCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode FailedPCellId_r16", err)
	}
	if ReconnectCellId_r16Present {
		ie.ReconnectCellId_r16 = new(RLF_Report_r16_nr_RLF_Report_r16_reconnectCellId_r16)
		if err = ie.ReconnectCellId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ReconnectCellId_r16", err)
		}
	}
	if TimeUntilReconnection_r16Present {
		ie.TimeUntilReconnection_r16 = new(TimeUntilReconnection_r16)
		if err = ie.TimeUntilReconnection_r16.Decode(r); err != nil {
			return utils.WrapError("Decode TimeUntilReconnection_r16", err)
		}
	}
	if ReestablishmentCellId_r16Present {
		ie.ReestablishmentCellId_r16 = new(CGI_Info_Logging_r16)
		if err = ie.ReestablishmentCellId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ReestablishmentCellId_r16", err)
		}
	}
	if TimeConnFailure_r16Present {
		var tmp_int_TimeConnFailure_r16 int64
		if tmp_int_TimeConnFailure_r16, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode TimeConnFailure_r16", err)
		}
		ie.TimeConnFailure_r16 = &tmp_int_TimeConnFailure_r16
	}
	if err = ie.TimeSinceFailure_r16.Decode(r); err != nil {
		return utils.WrapError("Decode TimeSinceFailure_r16", err)
	}
	if err = ie.ConnectionFailureType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ConnectionFailureType_r16", err)
	}
	if err = ie.Rlf_Cause_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Rlf_Cause_r16", err)
	}
	if LocationInfo_r16Present {
		ie.LocationInfo_r16 = new(LocationInfo_r16)
		if err = ie.LocationInfo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode LocationInfo_r16", err)
		}
	}
	if NoSuitableCellFound_r16Present {
		ie.NoSuitableCellFound_r16 = new(RLF_Report_r16_nr_RLF_Report_r16_noSuitableCellFound_r16)
		if err = ie.NoSuitableCellFound_r16.Decode(r); err != nil {
			return utils.WrapError("Decode NoSuitableCellFound_r16", err)
		}
	}
	if Ra_InformationCommon_r16Present {
		ie.Ra_InformationCommon_r16 = new(RA_InformationCommon_r16)
		if err = ie.Ra_InformationCommon_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ra_InformationCommon_r16", err)
		}
	}
	if Csi_rsRLMConfigBitmap_v1650Present {
		var tmp_bs_Csi_rsRLMConfigBitmap_v1650 []byte
		var n_Csi_rsRLMConfigBitmap_v1650 uint
		if tmp_bs_Csi_rsRLMConfigBitmap_v1650, n_Csi_rsRLMConfigBitmap_v1650, err = r.ReadBitString(&aper.Constraint{Lb: 96, Ub: 96}, false); err != nil {
			return utils.WrapError("Decode Csi_rsRLMConfigBitmap_v1650", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_Csi_rsRLMConfigBitmap_v1650,
			NumBits: uint64(n_Csi_rsRLMConfigBitmap_v1650),
		}
		ie.Csi_rsRLMConfigBitmap_v1650 = &tmp_bitstring
	}
	if LastHO_Type_r17Present {
		ie.LastHO_Type_r17 = new(RLF_Report_r16_nr_RLF_Report_r16_lastHO_Type_r17)
		if err = ie.LastHO_Type_r17.Decode(r); err != nil {
			return utils.WrapError("Decode LastHO_Type_r17", err)
		}
	}
	if TimeConnSourceDAPS_Failure_r17Present {
		ie.TimeConnSourceDAPS_Failure_r17 = new(TimeConnSourceDAPS_Failure_r17)
		if err = ie.TimeConnSourceDAPS_Failure_r17.Decode(r); err != nil {
			return utils.WrapError("Decode TimeConnSourceDAPS_Failure_r17", err)
		}
	}
	if TimeSinceCHO_Reconfig_r17Present {
		ie.TimeSinceCHO_Reconfig_r17 = new(TimeSinceCHO_Reconfig_r17)
		if err = ie.TimeSinceCHO_Reconfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode TimeSinceCHO_Reconfig_r17", err)
		}
	}
	if ChoCellId_r17Present {
		ie.ChoCellId_r17 = new(RLF_Report_r16_nr_RLF_Report_r16_choCellId_r17)
		if err = ie.ChoCellId_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ChoCellId_r17", err)
		}
	}
	if ChoCandidateCellList_r17Present {
		ie.ChoCandidateCellList_r17 = new(ChoCandidateCellList_r17)
		if err = ie.ChoCandidateCellList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ChoCandidateCellList_r17", err)
		}
	}
	return nil
}
