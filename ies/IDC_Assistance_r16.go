package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IDC_Assistance_r16 struct {
	AffectedCarrierFreqList_r16     *AffectedCarrierFreqList_r16     `optional`
	AffectedCarrierFreqCombList_r16 *AffectedCarrierFreqCombList_r16 `optional`
}

func (ie *IDC_Assistance_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.AffectedCarrierFreqList_r16 != nil, ie.AffectedCarrierFreqCombList_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.AffectedCarrierFreqList_r16 != nil {
		if err = ie.AffectedCarrierFreqList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AffectedCarrierFreqList_r16", err)
		}
	}
	if ie.AffectedCarrierFreqCombList_r16 != nil {
		if err = ie.AffectedCarrierFreqCombList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AffectedCarrierFreqCombList_r16", err)
		}
	}
	return nil
}

func (ie *IDC_Assistance_r16) Decode(r *uper.UperReader) error {
	var err error
	var AffectedCarrierFreqList_r16Present bool
	if AffectedCarrierFreqList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AffectedCarrierFreqCombList_r16Present bool
	if AffectedCarrierFreqCombList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if AffectedCarrierFreqList_r16Present {
		ie.AffectedCarrierFreqList_r16 = new(AffectedCarrierFreqList_r16)
		if err = ie.AffectedCarrierFreqList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode AffectedCarrierFreqList_r16", err)
		}
	}
	if AffectedCarrierFreqCombList_r16Present {
		ie.AffectedCarrierFreqCombList_r16 = new(AffectedCarrierFreqCombList_r16)
		if err = ie.AffectedCarrierFreqCombList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode AffectedCarrierFreqCombList_r16", err)
		}
	}
	return nil
}
