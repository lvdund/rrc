package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UEAssistanceInformation_v1700_IEs_nonSDT_DataIndication_r17 struct {
	ResumeCause_r17 *ResumeCause `optional`
}

func (ie *UEAssistanceInformation_v1700_IEs_nonSDT_DataIndication_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ResumeCause_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ResumeCause_r17 != nil {
		if err = ie.ResumeCause_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ResumeCause_r17", err)
		}
	}
	return nil
}

func (ie *UEAssistanceInformation_v1700_IEs_nonSDT_DataIndication_r17) Decode(r *aper.AperReader) error {
	var err error
	var ResumeCause_r17Present bool
	if ResumeCause_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ResumeCause_r17Present {
		ie.ResumeCause_r17 = new(ResumeCause)
		if err = ie.ResumeCause_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ResumeCause_r17", err)
		}
	}
	return nil
}
