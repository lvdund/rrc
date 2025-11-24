package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB21_r17 struct {
	Mbs_FSAI_IntraFreq_r17     *MBS_FSAI_List_r17          `optional`
	Mbs_FSAI_InterFreqList_r17 *MBS_FSAI_InterFreqList_r17 `optional`
	LateNonCriticalExtension   *[]byte                     `optional`
}

func (ie *SIB21_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Mbs_FSAI_IntraFreq_r17 != nil, ie.Mbs_FSAI_InterFreqList_r17 != nil, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Mbs_FSAI_IntraFreq_r17 != nil {
		if err = ie.Mbs_FSAI_IntraFreq_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mbs_FSAI_IntraFreq_r17", err)
		}
	}
	if ie.Mbs_FSAI_InterFreqList_r17 != nil {
		if err = ie.Mbs_FSAI_InterFreqList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mbs_FSAI_InterFreqList_r17", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB21_r17) Decode(r *uper.UperReader) error {
	var err error
	var Mbs_FSAI_IntraFreq_r17Present bool
	if Mbs_FSAI_IntraFreq_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mbs_FSAI_InterFreqList_r17Present bool
	if Mbs_FSAI_InterFreqList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Mbs_FSAI_IntraFreq_r17Present {
		ie.Mbs_FSAI_IntraFreq_r17 = new(MBS_FSAI_List_r17)
		if err = ie.Mbs_FSAI_IntraFreq_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mbs_FSAI_IntraFreq_r17", err)
		}
	}
	if Mbs_FSAI_InterFreqList_r17Present {
		ie.Mbs_FSAI_InterFreqList_r17 = new(MBS_FSAI_InterFreqList_r17)
		if err = ie.Mbs_FSAI_InterFreqList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mbs_FSAI_InterFreqList_r17", err)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
