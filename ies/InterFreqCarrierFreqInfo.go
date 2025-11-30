package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqCarrierFreqInfo struct {
	Dl_CarrierFreq                  ARFCN_ValueNR                 `madatory`
	FrequencyBandList               *MultiFrequencyBandListNR_SIB `optional`
	FrequencyBandListSUL            *MultiFrequencyBandListNR_SIB `optional`
	NrofSS_BlocksToAverage          *int64                        `lb:2,ub:maxNrofSS_BlocksToAverage,optional`
	AbsThreshSS_BlocksConsolidation *ThresholdNR                  `optional`
	Smtc                            *SSB_MTC                      `optional`
	SsbSubcarrierSpacing            SubcarrierSpacing             `madatory`
	Ssb_ToMeasure                   *SSB_ToMeasure                `optional`
	DeriveSSB_IndexFromCell         bool                          `madatory`
	Ss_RSSI_Measurement             *SS_RSSI_Measurement          `optional`
	Q_RxLevMin                      Q_RxLevMin                    `madatory`
	Q_RxLevMinSUL                   *Q_RxLevMin                   `optional`
	Q_QualMin                       *Q_QualMin                    `optional`
	P_Max                           *P_Max                        `optional`
	T_ReselectionNR                 T_Reselection                 `madatory`
	T_ReselectionNR_SF              *SpeedStateScaleFactors       `optional`
	ThreshX_HighP                   ReselectionThreshold          `madatory`
	ThreshX_LowP                    ReselectionThreshold          `madatory`
	ThreshX_Q                       *ThreshX_Q                    `optional`
	CellReselectionPriority         *CellReselectionPriority      `optional`
	CellReselectionSubPriority      *CellReselectionSubPriority   `optional`
	Q_OffsetFreq                    Q_OffsetRange                 `madatory`
	InterFreqNeighCellList          *InterFreqNeighCellList       `optional`
	InterFreqExcludedCellList       *InterFreqExcludedCellList    `optional`
}

func (ie *InterFreqCarrierFreqInfo) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.FrequencyBandList != nil, ie.FrequencyBandListSUL != nil, ie.NrofSS_BlocksToAverage != nil, ie.AbsThreshSS_BlocksConsolidation != nil, ie.Smtc != nil, ie.Ssb_ToMeasure != nil, ie.Ss_RSSI_Measurement != nil, ie.Q_RxLevMinSUL != nil, ie.Q_QualMin != nil, ie.P_Max != nil, ie.T_ReselectionNR_SF != nil, ie.ThreshX_Q != nil, ie.CellReselectionPriority != nil, ie.CellReselectionSubPriority != nil, ie.InterFreqNeighCellList != nil, ie.InterFreqExcludedCellList != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Dl_CarrierFreq.Encode(w); err != nil {
		return utils.WrapError("Encode Dl_CarrierFreq", err)
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
	if ie.NrofSS_BlocksToAverage != nil {
		if err = w.WriteInteger(*ie.NrofSS_BlocksToAverage, &aper.Constraint{Lb: 2, Ub: maxNrofSS_BlocksToAverage}, false); err != nil {
			return utils.WrapError("Encode NrofSS_BlocksToAverage", err)
		}
	}
	if ie.AbsThreshSS_BlocksConsolidation != nil {
		if err = ie.AbsThreshSS_BlocksConsolidation.Encode(w); err != nil {
			return utils.WrapError("Encode AbsThreshSS_BlocksConsolidation", err)
		}
	}
	if ie.Smtc != nil {
		if err = ie.Smtc.Encode(w); err != nil {
			return utils.WrapError("Encode Smtc", err)
		}
	}
	if err = ie.SsbSubcarrierSpacing.Encode(w); err != nil {
		return utils.WrapError("Encode SsbSubcarrierSpacing", err)
	}
	if ie.Ssb_ToMeasure != nil {
		if err = ie.Ssb_ToMeasure.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_ToMeasure", err)
		}
	}
	if err = w.WriteBoolean(ie.DeriveSSB_IndexFromCell); err != nil {
		return utils.WrapError("WriteBoolean DeriveSSB_IndexFromCell", err)
	}
	if ie.Ss_RSSI_Measurement != nil {
		if err = ie.Ss_RSSI_Measurement.Encode(w); err != nil {
			return utils.WrapError("Encode Ss_RSSI_Measurement", err)
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
	if ie.P_Max != nil {
		if err = ie.P_Max.Encode(w); err != nil {
			return utils.WrapError("Encode P_Max", err)
		}
	}
	if err = ie.T_ReselectionNR.Encode(w); err != nil {
		return utils.WrapError("Encode T_ReselectionNR", err)
	}
	if ie.T_ReselectionNR_SF != nil {
		if err = ie.T_ReselectionNR_SF.Encode(w); err != nil {
			return utils.WrapError("Encode T_ReselectionNR_SF", err)
		}
	}
	if err = ie.ThreshX_HighP.Encode(w); err != nil {
		return utils.WrapError("Encode ThreshX_HighP", err)
	}
	if err = ie.ThreshX_LowP.Encode(w); err != nil {
		return utils.WrapError("Encode ThreshX_LowP", err)
	}
	if ie.ThreshX_Q != nil {
		if err = ie.ThreshX_Q.Encode(w); err != nil {
			return utils.WrapError("Encode ThreshX_Q", err)
		}
	}
	if ie.CellReselectionPriority != nil {
		if err = ie.CellReselectionPriority.Encode(w); err != nil {
			return utils.WrapError("Encode CellReselectionPriority", err)
		}
	}
	if ie.CellReselectionSubPriority != nil {
		if err = ie.CellReselectionSubPriority.Encode(w); err != nil {
			return utils.WrapError("Encode CellReselectionSubPriority", err)
		}
	}
	if err = ie.Q_OffsetFreq.Encode(w); err != nil {
		return utils.WrapError("Encode Q_OffsetFreq", err)
	}
	if ie.InterFreqNeighCellList != nil {
		if err = ie.InterFreqNeighCellList.Encode(w); err != nil {
			return utils.WrapError("Encode InterFreqNeighCellList", err)
		}
	}
	if ie.InterFreqExcludedCellList != nil {
		if err = ie.InterFreqExcludedCellList.Encode(w); err != nil {
			return utils.WrapError("Encode InterFreqExcludedCellList", err)
		}
	}
	return nil
}

func (ie *InterFreqCarrierFreqInfo) Decode(r *aper.AperReader) error {
	var err error
	var FrequencyBandListPresent bool
	if FrequencyBandListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var FrequencyBandListSULPresent bool
	if FrequencyBandListSULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NrofSS_BlocksToAveragePresent bool
	if NrofSS_BlocksToAveragePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AbsThreshSS_BlocksConsolidationPresent bool
	if AbsThreshSS_BlocksConsolidationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SmtcPresent bool
	if SmtcPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_ToMeasurePresent bool
	if Ssb_ToMeasurePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ss_RSSI_MeasurementPresent bool
	if Ss_RSSI_MeasurementPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Q_RxLevMinSULPresent bool
	if Q_RxLevMinSULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Q_QualMinPresent bool
	if Q_QualMinPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var P_MaxPresent bool
	if P_MaxPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var T_ReselectionNR_SFPresent bool
	if T_ReselectionNR_SFPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ThreshX_QPresent bool
	if ThreshX_QPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CellReselectionPriorityPresent bool
	if CellReselectionPriorityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CellReselectionSubPriorityPresent bool
	if CellReselectionSubPriorityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var InterFreqNeighCellListPresent bool
	if InterFreqNeighCellListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var InterFreqExcludedCellListPresent bool
	if InterFreqExcludedCellListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Dl_CarrierFreq.Decode(r); err != nil {
		return utils.WrapError("Decode Dl_CarrierFreq", err)
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
	if NrofSS_BlocksToAveragePresent {
		var tmp_int_NrofSS_BlocksToAverage int64
		if tmp_int_NrofSS_BlocksToAverage, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: maxNrofSS_BlocksToAverage}, false); err != nil {
			return utils.WrapError("Decode NrofSS_BlocksToAverage", err)
		}
		ie.NrofSS_BlocksToAverage = &tmp_int_NrofSS_BlocksToAverage
	}
	if AbsThreshSS_BlocksConsolidationPresent {
		ie.AbsThreshSS_BlocksConsolidation = new(ThresholdNR)
		if err = ie.AbsThreshSS_BlocksConsolidation.Decode(r); err != nil {
			return utils.WrapError("Decode AbsThreshSS_BlocksConsolidation", err)
		}
	}
	if SmtcPresent {
		ie.Smtc = new(SSB_MTC)
		if err = ie.Smtc.Decode(r); err != nil {
			return utils.WrapError("Decode Smtc", err)
		}
	}
	if err = ie.SsbSubcarrierSpacing.Decode(r); err != nil {
		return utils.WrapError("Decode SsbSubcarrierSpacing", err)
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
	if Ss_RSSI_MeasurementPresent {
		ie.Ss_RSSI_Measurement = new(SS_RSSI_Measurement)
		if err = ie.Ss_RSSI_Measurement.Decode(r); err != nil {
			return utils.WrapError("Decode Ss_RSSI_Measurement", err)
		}
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
	if P_MaxPresent {
		ie.P_Max = new(P_Max)
		if err = ie.P_Max.Decode(r); err != nil {
			return utils.WrapError("Decode P_Max", err)
		}
	}
	if err = ie.T_ReselectionNR.Decode(r); err != nil {
		return utils.WrapError("Decode T_ReselectionNR", err)
	}
	if T_ReselectionNR_SFPresent {
		ie.T_ReselectionNR_SF = new(SpeedStateScaleFactors)
		if err = ie.T_ReselectionNR_SF.Decode(r); err != nil {
			return utils.WrapError("Decode T_ReselectionNR_SF", err)
		}
	}
	if err = ie.ThreshX_HighP.Decode(r); err != nil {
		return utils.WrapError("Decode ThreshX_HighP", err)
	}
	if err = ie.ThreshX_LowP.Decode(r); err != nil {
		return utils.WrapError("Decode ThreshX_LowP", err)
	}
	if ThreshX_QPresent {
		ie.ThreshX_Q = new(ThreshX_Q)
		if err = ie.ThreshX_Q.Decode(r); err != nil {
			return utils.WrapError("Decode ThreshX_Q", err)
		}
	}
	if CellReselectionPriorityPresent {
		ie.CellReselectionPriority = new(CellReselectionPriority)
		if err = ie.CellReselectionPriority.Decode(r); err != nil {
			return utils.WrapError("Decode CellReselectionPriority", err)
		}
	}
	if CellReselectionSubPriorityPresent {
		ie.CellReselectionSubPriority = new(CellReselectionSubPriority)
		if err = ie.CellReselectionSubPriority.Decode(r); err != nil {
			return utils.WrapError("Decode CellReselectionSubPriority", err)
		}
	}
	if err = ie.Q_OffsetFreq.Decode(r); err != nil {
		return utils.WrapError("Decode Q_OffsetFreq", err)
	}
	if InterFreqNeighCellListPresent {
		ie.InterFreqNeighCellList = new(InterFreqNeighCellList)
		if err = ie.InterFreqNeighCellList.Decode(r); err != nil {
			return utils.WrapError("Decode InterFreqNeighCellList", err)
		}
	}
	if InterFreqExcludedCellListPresent {
		ie.InterFreqExcludedCellList = new(InterFreqExcludedCellList)
		if err = ie.InterFreqExcludedCellList.Decode(r); err != nil {
			return utils.WrapError("Decode InterFreqExcludedCellList", err)
		}
	}
	return nil
}
