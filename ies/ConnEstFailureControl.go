package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ConnEstFailureControl struct {
	ConnEstFailCount          ConnEstFailureControl_connEstFailCount          `madatory`
	ConnEstFailOffsetValidity ConnEstFailureControl_connEstFailOffsetValidity `madatory`
	ConnEstFailOffset         *int64                                          `lb:0,ub:15,optional`
}

func (ie *ConnEstFailureControl) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ConnEstFailOffset != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ConnEstFailCount.Encode(w); err != nil {
		return utils.WrapError("Encode ConnEstFailCount", err)
	}
	if err = ie.ConnEstFailOffsetValidity.Encode(w); err != nil {
		return utils.WrapError("Encode ConnEstFailOffsetValidity", err)
	}
	if ie.ConnEstFailOffset != nil {
		if err = w.WriteInteger(*ie.ConnEstFailOffset, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode ConnEstFailOffset", err)
		}
	}
	return nil
}

func (ie *ConnEstFailureControl) Decode(r *uper.UperReader) error {
	var err error
	var ConnEstFailOffsetPresent bool
	if ConnEstFailOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ConnEstFailCount.Decode(r); err != nil {
		return utils.WrapError("Decode ConnEstFailCount", err)
	}
	if err = ie.ConnEstFailOffsetValidity.Decode(r); err != nil {
		return utils.WrapError("Decode ConnEstFailOffsetValidity", err)
	}
	if ConnEstFailOffsetPresent {
		var tmp_int_ConnEstFailOffset int64
		if tmp_int_ConnEstFailOffset, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode ConnEstFailOffset", err)
		}
		ie.ConnEstFailOffset = &tmp_int_ConnEstFailOffset
	}
	return nil
}
