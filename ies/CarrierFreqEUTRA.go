package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CarrierFreqEUTRA struct {
	CarrierFreq                ARFCN_ValueEUTRA            `madatory`
	Eutra_multiBandInfoList    *EUTRA_MultiBandInfoList    `optional`
	Eutra_FreqNeighCellList    *EUTRA_FreqNeighCellList    `optional`
	Eutra_ExcludedCellList     *EUTRA_FreqExcludedCellList `optional`
	AllowedMeasBandwidth       EUTRA_AllowedMeasBandwidth  `madatory`
	PresenceAntennaPort1       EUTRA_PresenceAntennaPort1  `madatory`
	CellReselectionPriority    *CellReselectionPriority    `optional`
	CellReselectionSubPriority *CellReselectionSubPriority `optional`
	ThreshX_High               ReselectionThreshold        `madatory`
	ThreshX_Low                ReselectionThreshold        `madatory`
	Q_RxLevMin                 int64                       `lb:-70,ub:-22,madatory`
	Q_QualMin                  int64                       `lb:-34,ub:-3,madatory`
	P_MaxEUTRA                 int64                       `lb:-30,ub:33,madatory`
	ThreshX_Q                  *ThreshX_Q                  `optional`
}

func (ie *CarrierFreqEUTRA) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Eutra_multiBandInfoList != nil, ie.Eutra_FreqNeighCellList != nil, ie.Eutra_ExcludedCellList != nil, ie.CellReselectionPriority != nil, ie.CellReselectionSubPriority != nil, ie.ThreshX_Q != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.CarrierFreq.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreq", err)
	}
	if ie.Eutra_multiBandInfoList != nil {
		if err = ie.Eutra_multiBandInfoList.Encode(w); err != nil {
			return utils.WrapError("Encode Eutra_multiBandInfoList", err)
		}
	}
	if ie.Eutra_FreqNeighCellList != nil {
		if err = ie.Eutra_FreqNeighCellList.Encode(w); err != nil {
			return utils.WrapError("Encode Eutra_FreqNeighCellList", err)
		}
	}
	if ie.Eutra_ExcludedCellList != nil {
		if err = ie.Eutra_ExcludedCellList.Encode(w); err != nil {
			return utils.WrapError("Encode Eutra_ExcludedCellList", err)
		}
	}
	if err = ie.AllowedMeasBandwidth.Encode(w); err != nil {
		return utils.WrapError("Encode AllowedMeasBandwidth", err)
	}
	if err = ie.PresenceAntennaPort1.Encode(w); err != nil {
		return utils.WrapError("Encode PresenceAntennaPort1", err)
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
	if err = ie.ThreshX_High.Encode(w); err != nil {
		return utils.WrapError("Encode ThreshX_High", err)
	}
	if err = ie.ThreshX_Low.Encode(w); err != nil {
		return utils.WrapError("Encode ThreshX_Low", err)
	}
	if err = w.WriteInteger(ie.Q_RxLevMin, &aper.Constraint{Lb: -70, Ub: -22}, false); err != nil {
		return utils.WrapError("WriteInteger Q_RxLevMin", err)
	}
	if err = w.WriteInteger(ie.Q_QualMin, &aper.Constraint{Lb: -34, Ub: -3}, false); err != nil {
		return utils.WrapError("WriteInteger Q_QualMin", err)
	}
	if err = w.WriteInteger(ie.P_MaxEUTRA, &aper.Constraint{Lb: -30, Ub: 33}, false); err != nil {
		return utils.WrapError("WriteInteger P_MaxEUTRA", err)
	}
	if ie.ThreshX_Q != nil {
		if err = ie.ThreshX_Q.Encode(w); err != nil {
			return utils.WrapError("Encode ThreshX_Q", err)
		}
	}
	return nil
}

func (ie *CarrierFreqEUTRA) Decode(r *aper.AperReader) error {
	var err error
	var Eutra_multiBandInfoListPresent bool
	if Eutra_multiBandInfoListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Eutra_FreqNeighCellListPresent bool
	if Eutra_FreqNeighCellListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Eutra_ExcludedCellListPresent bool
	if Eutra_ExcludedCellListPresent, err = r.ReadBool(); err != nil {
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
	var ThreshX_QPresent bool
	if ThreshX_QPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.CarrierFreq.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierFreq", err)
	}
	if Eutra_multiBandInfoListPresent {
		ie.Eutra_multiBandInfoList = new(EUTRA_MultiBandInfoList)
		if err = ie.Eutra_multiBandInfoList.Decode(r); err != nil {
			return utils.WrapError("Decode Eutra_multiBandInfoList", err)
		}
	}
	if Eutra_FreqNeighCellListPresent {
		ie.Eutra_FreqNeighCellList = new(EUTRA_FreqNeighCellList)
		if err = ie.Eutra_FreqNeighCellList.Decode(r); err != nil {
			return utils.WrapError("Decode Eutra_FreqNeighCellList", err)
		}
	}
	if Eutra_ExcludedCellListPresent {
		ie.Eutra_ExcludedCellList = new(EUTRA_FreqExcludedCellList)
		if err = ie.Eutra_ExcludedCellList.Decode(r); err != nil {
			return utils.WrapError("Decode Eutra_ExcludedCellList", err)
		}
	}
	if err = ie.AllowedMeasBandwidth.Decode(r); err != nil {
		return utils.WrapError("Decode AllowedMeasBandwidth", err)
	}
	if err = ie.PresenceAntennaPort1.Decode(r); err != nil {
		return utils.WrapError("Decode PresenceAntennaPort1", err)
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
	if err = ie.ThreshX_High.Decode(r); err != nil {
		return utils.WrapError("Decode ThreshX_High", err)
	}
	if err = ie.ThreshX_Low.Decode(r); err != nil {
		return utils.WrapError("Decode ThreshX_Low", err)
	}
	var tmp_int_Q_RxLevMin int64
	if tmp_int_Q_RxLevMin, err = r.ReadInteger(&aper.Constraint{Lb: -70, Ub: -22}, false); err != nil {
		return utils.WrapError("ReadInteger Q_RxLevMin", err)
	}
	ie.Q_RxLevMin = tmp_int_Q_RxLevMin
	var tmp_int_Q_QualMin int64
	if tmp_int_Q_QualMin, err = r.ReadInteger(&aper.Constraint{Lb: -34, Ub: -3}, false); err != nil {
		return utils.WrapError("ReadInteger Q_QualMin", err)
	}
	ie.Q_QualMin = tmp_int_Q_QualMin
	var tmp_int_P_MaxEUTRA int64
	if tmp_int_P_MaxEUTRA, err = r.ReadInteger(&aper.Constraint{Lb: -30, Ub: 33}, false); err != nil {
		return utils.WrapError("ReadInteger P_MaxEUTRA", err)
	}
	ie.P_MaxEUTRA = tmp_int_P_MaxEUTRA
	if ThreshX_QPresent {
		ie.ThreshX_Q = new(ThreshX_Q)
		if err = ie.ThreshX_Q.Decode(r); err != nil {
			return utils.WrapError("Decode ThreshX_Q", err)
		}
	}
	return nil
}
