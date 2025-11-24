package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FreqPrioritySlicing_r17 struct {
	Dl_ImplicitCarrierFreq_r17 int64              `lb:0,ub:maxFreq,madatory`
	SliceInfoList_r17          *SliceInfoList_r17 `optional`
}

func (ie *FreqPrioritySlicing_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.SliceInfoList_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.Dl_ImplicitCarrierFreq_r17, &uper.Constraint{Lb: 0, Ub: maxFreq}, false); err != nil {
		return utils.WrapError("WriteInteger Dl_ImplicitCarrierFreq_r17", err)
	}
	if ie.SliceInfoList_r17 != nil {
		if err = ie.SliceInfoList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SliceInfoList_r17", err)
		}
	}
	return nil
}

func (ie *FreqPrioritySlicing_r17) Decode(r *uper.UperReader) error {
	var err error
	var SliceInfoList_r17Present bool
	if SliceInfoList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_Dl_ImplicitCarrierFreq_r17 int64
	if tmp_int_Dl_ImplicitCarrierFreq_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxFreq}, false); err != nil {
		return utils.WrapError("ReadInteger Dl_ImplicitCarrierFreq_r17", err)
	}
	ie.Dl_ImplicitCarrierFreq_r17 = tmp_int_Dl_ImplicitCarrierFreq_r17
	if SliceInfoList_r17Present {
		ie.SliceInfoList_r17 = new(SliceInfoList_r17)
		if err = ie.SliceInfoList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SliceInfoList_r17", err)
		}
	}
	return nil
}
