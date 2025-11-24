package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_FreqNeighCellInfo struct {
	PhysCellId           EUTRA_PhysCellId    `madatory`
	Dummy                EUTRA_Q_OffsetRange `madatory`
	Q_RxLevMinOffsetCell *int64              `lb:1,ub:8,optional`
	Q_QualMinOffsetCell  *int64              `lb:1,ub:8,optional`
}

func (ie *EUTRA_FreqNeighCellInfo) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Q_RxLevMinOffsetCell != nil, ie.Q_QualMinOffsetCell != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.PhysCellId.Encode(w); err != nil {
		return utils.WrapError("Encode PhysCellId", err)
	}
	if err = ie.Dummy.Encode(w); err != nil {
		return utils.WrapError("Encode Dummy", err)
	}
	if ie.Q_RxLevMinOffsetCell != nil {
		if err = w.WriteInteger(*ie.Q_RxLevMinOffsetCell, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Q_RxLevMinOffsetCell", err)
		}
	}
	if ie.Q_QualMinOffsetCell != nil {
		if err = w.WriteInteger(*ie.Q_QualMinOffsetCell, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Q_QualMinOffsetCell", err)
		}
	}
	return nil
}

func (ie *EUTRA_FreqNeighCellInfo) Decode(r *uper.UperReader) error {
	var err error
	var Q_RxLevMinOffsetCellPresent bool
	if Q_RxLevMinOffsetCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Q_QualMinOffsetCellPresent bool
	if Q_QualMinOffsetCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.PhysCellId.Decode(r); err != nil {
		return utils.WrapError("Decode PhysCellId", err)
	}
	if err = ie.Dummy.Decode(r); err != nil {
		return utils.WrapError("Decode Dummy", err)
	}
	if Q_RxLevMinOffsetCellPresent {
		var tmp_int_Q_RxLevMinOffsetCell int64
		if tmp_int_Q_RxLevMinOffsetCell, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Q_RxLevMinOffsetCell", err)
		}
		ie.Q_RxLevMinOffsetCell = &tmp_int_Q_RxLevMinOffsetCell
	}
	if Q_QualMinOffsetCellPresent {
		var tmp_int_Q_QualMinOffsetCell int64
		if tmp_int_Q_QualMinOffsetCell, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Q_QualMinOffsetCell", err)
		}
		ie.Q_QualMinOffsetCell = &tmp_int_Q_QualMinOffsetCell
	}
	return nil
}
