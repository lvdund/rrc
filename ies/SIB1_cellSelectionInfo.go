package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB1_cellSelectionInfo struct {
	Q_RxLevMin       Q_RxLevMin  `madatory`
	Q_RxLevMinOffset *int64      `lb:1,ub:8,optional`
	Q_RxLevMinSUL    *Q_RxLevMin `optional`
	Q_QualMin        *Q_QualMin  `optional`
	Q_QualMinOffset  *int64      `lb:1,ub:8,optional`
}

func (ie *SIB1_cellSelectionInfo) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Q_RxLevMinOffset != nil, ie.Q_RxLevMinSUL != nil, ie.Q_QualMin != nil, ie.Q_QualMinOffset != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Q_RxLevMin.Encode(w); err != nil {
		return utils.WrapError("Encode Q_RxLevMin", err)
	}
	if ie.Q_RxLevMinOffset != nil {
		if err = w.WriteInteger(*ie.Q_RxLevMinOffset, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Q_RxLevMinOffset", err)
		}
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
	if ie.Q_QualMinOffset != nil {
		if err = w.WriteInteger(*ie.Q_QualMinOffset, &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Q_QualMinOffset", err)
		}
	}
	return nil
}

func (ie *SIB1_cellSelectionInfo) Decode(r *aper.AperReader) error {
	var err error
	var Q_RxLevMinOffsetPresent bool
	if Q_RxLevMinOffsetPresent, err = r.ReadBool(); err != nil {
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
	var Q_QualMinOffsetPresent bool
	if Q_QualMinOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Q_RxLevMin.Decode(r); err != nil {
		return utils.WrapError("Decode Q_RxLevMin", err)
	}
	if Q_RxLevMinOffsetPresent {
		var tmp_int_Q_RxLevMinOffset int64
		if tmp_int_Q_RxLevMinOffset, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Q_RxLevMinOffset", err)
		}
		ie.Q_RxLevMinOffset = &tmp_int_Q_RxLevMinOffset
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
	if Q_QualMinOffsetPresent {
		var tmp_int_Q_QualMinOffset int64
		if tmp_int_Q_QualMinOffset, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Q_QualMinOffset", err)
		}
		ie.Q_QualMinOffset = &tmp_int_Q_QualMinOffset
	}
	return nil
}
