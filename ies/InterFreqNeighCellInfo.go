package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqNeighCellInfo struct {
	PhysCellId              PhysCellId    `madatory`
	Q_OffsetCell            Q_OffsetRange `madatory`
	Q_RxLevMinOffsetCell    *int64        `lb:1,ub:8,optional`
	Q_RxLevMinOffsetCellSUL *int64        `lb:1,ub:8,optional`
	Q_QualMinOffsetCell     *int64        `lb:1,ub:8,optional`
}

func (ie *InterFreqNeighCellInfo) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Q_RxLevMinOffsetCell != nil, ie.Q_RxLevMinOffsetCellSUL != nil, ie.Q_QualMinOffsetCell != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.PhysCellId.Encode(w); err != nil {
		return utils.WrapError("Encode PhysCellId", err)
	}
	if err = ie.Q_OffsetCell.Encode(w); err != nil {
		return utils.WrapError("Encode Q_OffsetCell", err)
	}
	if ie.Q_RxLevMinOffsetCell != nil {
		if err = w.WriteInteger(*ie.Q_RxLevMinOffsetCell, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Q_RxLevMinOffsetCell", err)
		}
	}
	if ie.Q_RxLevMinOffsetCellSUL != nil {
		if err = w.WriteInteger(*ie.Q_RxLevMinOffsetCellSUL, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Q_RxLevMinOffsetCellSUL", err)
		}
	}
	if ie.Q_QualMinOffsetCell != nil {
		if err = w.WriteInteger(*ie.Q_QualMinOffsetCell, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Q_QualMinOffsetCell", err)
		}
	}
	return nil
}

func (ie *InterFreqNeighCellInfo) Decode(r *aper.AperReader) error {
	var err error
	var Q_RxLevMinOffsetCellPresent bool
	if Q_RxLevMinOffsetCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Q_RxLevMinOffsetCellSULPresent bool
	if Q_RxLevMinOffsetCellSULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Q_QualMinOffsetCellPresent bool
	if Q_QualMinOffsetCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.PhysCellId.Decode(r); err != nil {
		return utils.WrapError("Decode PhysCellId", err)
	}
	if err = ie.Q_OffsetCell.Decode(r); err != nil {
		return utils.WrapError("Decode Q_OffsetCell", err)
	}
	if Q_RxLevMinOffsetCellPresent {
		var tmp_int_Q_RxLevMinOffsetCell int64
		if tmp_int_Q_RxLevMinOffsetCell, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Q_RxLevMinOffsetCell", err)
		}
		ie.Q_RxLevMinOffsetCell = &tmp_int_Q_RxLevMinOffsetCell
	}
	if Q_RxLevMinOffsetCellSULPresent {
		var tmp_int_Q_RxLevMinOffsetCellSUL int64
		if tmp_int_Q_RxLevMinOffsetCellSUL, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Q_RxLevMinOffsetCellSUL", err)
		}
		ie.Q_RxLevMinOffsetCellSUL = &tmp_int_Q_RxLevMinOffsetCellSUL
	}
	if Q_QualMinOffsetCellPresent {
		var tmp_int_Q_QualMinOffsetCell int64
		if tmp_int_Q_QualMinOffsetCell, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Q_QualMinOffsetCell", err)
		}
		ie.Q_QualMinOffsetCell = &tmp_int_Q_QualMinOffsetCell
	}
	return nil
}
