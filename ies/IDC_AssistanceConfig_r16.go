package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IDC_AssistanceConfig_r16 struct {
	CandidateServingFreqListNR_r16 *CandidateServingFreqListNR_r16 `optional`
}

func (ie *IDC_AssistanceConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.CandidateServingFreqListNR_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CandidateServingFreqListNR_r16 != nil {
		if err = ie.CandidateServingFreqListNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CandidateServingFreqListNR_r16", err)
		}
	}
	return nil
}

func (ie *IDC_AssistanceConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var CandidateServingFreqListNR_r16Present bool
	if CandidateServingFreqListNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if CandidateServingFreqListNR_r16Present {
		ie.CandidateServingFreqListNR_r16 = new(CandidateServingFreqListNR_r16)
		if err = ie.CandidateServingFreqListNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CandidateServingFreqListNR_r16", err)
		}
	}
	return nil
}
