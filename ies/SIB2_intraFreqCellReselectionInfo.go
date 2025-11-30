package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB2_intraFreqCellReselectionInfo struct {
	Q_RxLevMin                 Q_RxLevMin                    `madatory`
	Q_RxLevMinSUL              *Q_RxLevMin                   `optional`
	Q_QualMin                  *Q_QualMin                    `optional`
	S_IntraSearchP             ReselectionThreshold          `madatory`
	S_IntraSearchQ             *ReselectionThresholdQ        `optional`
	T_ReselectionNR            T_Reselection                 `madatory`
	FrequencyBandList          *MultiFrequencyBandListNR_SIB `optional`
	FrequencyBandListSUL       *MultiFrequencyBandListNR_SIB `optional`
	P_Max                      *P_Max                        `optional`
	Smtc                       *SSB_MTC                      `optional`
	Ss_RSSI_Measurement        *SS_RSSI_Measurement          `optional`
	Ssb_ToMeasure              *SSB_ToMeasure                `optional`
	DeriveSSB_IndexFromCell    bool                          `madatory`
	T_ReselectionNR_SF         *SpeedStateScaleFactors       `optional`
	Smtc2_LP_r16               *SSB_MTC2_LP_r16              `optional`
	Ssb_PositionQCL_Common_r16 *SSB_PositionQCL_Relation_r16 `optional`
	Ssb_PositionQCL_Common_r17 *SSB_PositionQCL_Relation_r17 `optional`
	Smtc4list_r17              *SSB_MTC4List_r17             `optional`
}

func (ie *SIB2_intraFreqCellReselectionInfo) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Q_RxLevMinSUL != nil, ie.Q_QualMin != nil, ie.S_IntraSearchQ != nil, ie.FrequencyBandList != nil, ie.FrequencyBandListSUL != nil, ie.P_Max != nil, ie.Smtc != nil, ie.Ss_RSSI_Measurement != nil, ie.Ssb_ToMeasure != nil, ie.T_ReselectionNR_SF != nil, ie.Smtc2_LP_r16 != nil, ie.Ssb_PositionQCL_Common_r16 != nil, ie.Ssb_PositionQCL_Common_r17 != nil, ie.Smtc4list_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Q_RxLevMin.Encode(w); err != nil {
		return utils.WrapError("Encode Q_RxLevMin", err)
	}
	if ie.Q_RxLevMinSUL != nil {
		if err = ie.Q_RxLevMinSUL.Encode(w); err != nil {
			return utils.WrapError("Encode Q_RxLevMinSUL", err)
		}
	}
	if ie.Q_QualMin != nil {
		if err = ie.Q_QualMin.Encode(w); err != nil {
			return utils.WrapError("Encode Q_QualMin", err)
		}
	}
	if err = ie.S_IntraSearchP.Encode(w); err != nil {
		return utils.WrapError("Encode S_IntraSearchP", err)
	}
	if ie.S_IntraSearchQ != nil {
		if err = ie.S_IntraSearchQ.Encode(w); err != nil {
			return utils.WrapError("Encode S_IntraSearchQ", err)
		}
	}
	if err = ie.T_ReselectionNR.Encode(w); err != nil {
		return utils.WrapError("Encode T_ReselectionNR", err)
	}
	if ie.FrequencyBandList != nil {
		if err = ie.FrequencyBandList.Encode(w); err != nil {
			return utils.WrapError("Encode FrequencyBandList", err)
		}
	}
	if ie.FrequencyBandListSUL != nil {
		if err = ie.FrequencyBandListSUL.Encode(w); err != nil {
			return utils.WrapError("Encode FrequencyBandListSUL", err)
		}
	}
	if ie.P_Max != nil {
		if err = ie.P_Max.Encode(w); err != nil {
			return utils.WrapError("Encode P_Max", err)
		}
	}
	if ie.Smtc != nil {
		if err = ie.Smtc.Encode(w); err != nil {
			return utils.WrapError("Encode Smtc", err)
		}
	}
	if ie.Ss_RSSI_Measurement != nil {
		if err = ie.Ss_RSSI_Measurement.Encode(w); err != nil {
			return utils.WrapError("Encode Ss_RSSI_Measurement", err)
		}
	}
	if ie.Ssb_ToMeasure != nil {
		if err = ie.Ssb_ToMeasure.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_ToMeasure", err)
		}
	}
	if err = w.WriteBoolean(ie.DeriveSSB_IndexFromCell); err != nil {
		return utils.WrapError("WriteBoolean DeriveSSB_IndexFromCell", err)
	}
	if ie.T_ReselectionNR_SF != nil {
		if err = ie.T_ReselectionNR_SF.Encode(w); err != nil {
			return utils.WrapError("Encode T_ReselectionNR_SF", err)
		}
	}
	if ie.Smtc2_LP_r16 != nil {
		if err = ie.Smtc2_LP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Smtc2_LP_r16", err)
		}
	}
	if ie.Ssb_PositionQCL_Common_r16 != nil {
		if err = ie.Ssb_PositionQCL_Common_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_PositionQCL_Common_r16", err)
		}
	}
	if ie.Ssb_PositionQCL_Common_r17 != nil {
		if err = ie.Ssb_PositionQCL_Common_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_PositionQCL_Common_r17", err)
		}
	}
	if ie.Smtc4list_r17 != nil {
		if err = ie.Smtc4list_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Smtc4list_r17", err)
		}
	}
	return nil
}

func (ie *SIB2_intraFreqCellReselectionInfo) Decode(r *aper.AperReader) error {
	var err error
	var Q_RxLevMinSULPresent bool
	if Q_RxLevMinSULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Q_QualMinPresent bool
	if Q_QualMinPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var S_IntraSearchQPresent bool
	if S_IntraSearchQPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FrequencyBandListPresent bool
	if FrequencyBandListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FrequencyBandListSULPresent bool
	if FrequencyBandListSULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var P_MaxPresent bool
	if P_MaxPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SmtcPresent bool
	if SmtcPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ss_RSSI_MeasurementPresent bool
	if Ss_RSSI_MeasurementPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_ToMeasurePresent bool
	if Ssb_ToMeasurePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var T_ReselectionNR_SFPresent bool
	if T_ReselectionNR_SFPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Smtc2_LP_r16Present bool
	if Smtc2_LP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_PositionQCL_Common_r16Present bool
	if Ssb_PositionQCL_Common_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_PositionQCL_Common_r17Present bool
	if Ssb_PositionQCL_Common_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Smtc4list_r17Present bool
	if Smtc4list_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Q_RxLevMin.Decode(r); err != nil {
		return utils.WrapError("Decode Q_RxLevMin", err)
	}
	if Q_RxLevMinSULPresent {
		ie.Q_RxLevMinSUL = new(Q_RxLevMin)
		if err = ie.Q_RxLevMinSUL.Decode(r); err != nil {
			return utils.WrapError("Decode Q_RxLevMinSUL", err)
		}
	}
	if Q_QualMinPresent {
		ie.Q_QualMin = new(Q_QualMin)
		if err = ie.Q_QualMin.Decode(r); err != nil {
			return utils.WrapError("Decode Q_QualMin", err)
		}
	}
	if err = ie.S_IntraSearchP.Decode(r); err != nil {
		return utils.WrapError("Decode S_IntraSearchP", err)
	}
	if S_IntraSearchQPresent {
		ie.S_IntraSearchQ = new(ReselectionThresholdQ)
		if err = ie.S_IntraSearchQ.Decode(r); err != nil {
			return utils.WrapError("Decode S_IntraSearchQ", err)
		}
	}
	if err = ie.T_ReselectionNR.Decode(r); err != nil {
		return utils.WrapError("Decode T_ReselectionNR", err)
	}
	if FrequencyBandListPresent {
		ie.FrequencyBandList = new(MultiFrequencyBandListNR_SIB)
		if err = ie.FrequencyBandList.Decode(r); err != nil {
			return utils.WrapError("Decode FrequencyBandList", err)
		}
	}
	if FrequencyBandListSULPresent {
		ie.FrequencyBandListSUL = new(MultiFrequencyBandListNR_SIB)
		if err = ie.FrequencyBandListSUL.Decode(r); err != nil {
			return utils.WrapError("Decode FrequencyBandListSUL", err)
		}
	}
	if P_MaxPresent {
		ie.P_Max = new(P_Max)
		if err = ie.P_Max.Decode(r); err != nil {
			return utils.WrapError("Decode P_Max", err)
		}
	}
	if SmtcPresent {
		ie.Smtc = new(SSB_MTC)
		if err = ie.Smtc.Decode(r); err != nil {
			return utils.WrapError("Decode Smtc", err)
		}
	}
	if Ss_RSSI_MeasurementPresent {
		ie.Ss_RSSI_Measurement = new(SS_RSSI_Measurement)
		if err = ie.Ss_RSSI_Measurement.Decode(r); err != nil {
			return utils.WrapError("Decode Ss_RSSI_Measurement", err)
		}
	}
	if Ssb_ToMeasurePresent {
		ie.Ssb_ToMeasure = new(SSB_ToMeasure)
		if err = ie.Ssb_ToMeasure.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_ToMeasure", err)
		}
	}
	var tmp_bool_DeriveSSB_IndexFromCell bool
	if tmp_bool_DeriveSSB_IndexFromCell, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean DeriveSSB_IndexFromCell", err)
	}
	ie.DeriveSSB_IndexFromCell = tmp_bool_DeriveSSB_IndexFromCell
	if T_ReselectionNR_SFPresent {
		ie.T_ReselectionNR_SF = new(SpeedStateScaleFactors)
		if err = ie.T_ReselectionNR_SF.Decode(r); err != nil {
			return utils.WrapError("Decode T_ReselectionNR_SF", err)
		}
	}
	if Smtc2_LP_r16Present {
		ie.Smtc2_LP_r16 = new(SSB_MTC2_LP_r16)
		if err = ie.Smtc2_LP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Smtc2_LP_r16", err)
		}
	}
	if Ssb_PositionQCL_Common_r16Present {
		ie.Ssb_PositionQCL_Common_r16 = new(SSB_PositionQCL_Relation_r16)
		if err = ie.Ssb_PositionQCL_Common_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_PositionQCL_Common_r16", err)
		}
	}
	if Ssb_PositionQCL_Common_r17Present {
		ie.Ssb_PositionQCL_Common_r17 = new(SSB_PositionQCL_Relation_r17)
		if err = ie.Ssb_PositionQCL_Common_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_PositionQCL_Common_r17", err)
		}
	}
	if Smtc4list_r17Present {
		ie.Smtc4list_r17 = new(SSB_MTC4List_r17)
		if err = ie.Smtc4list_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Smtc4list_r17", err)
		}
	}
	return nil
}
