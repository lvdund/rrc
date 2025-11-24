package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RedirectedCarrierInfo_EUTRA struct {
	EutraFrequency ARFCN_ValueEUTRA                    `madatory`
	CnType         *RedirectedCarrierInfo_EUTRA_cnType `optional`
}

func (ie *RedirectedCarrierInfo_EUTRA) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.CnType != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.EutraFrequency.Encode(w); err != nil {
		return utils.WrapError("Encode EutraFrequency", err)
	}
	if ie.CnType != nil {
		if err = ie.CnType.Encode(w); err != nil {
			return utils.WrapError("Encode CnType", err)
		}
	}
	return nil
}

func (ie *RedirectedCarrierInfo_EUTRA) Decode(r *uper.UperReader) error {
	var err error
	var CnTypePresent bool
	if CnTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.EutraFrequency.Decode(r); err != nil {
		return utils.WrapError("Decode EutraFrequency", err)
	}
	if CnTypePresent {
		ie.CnType = new(RedirectedCarrierInfo_EUTRA_cnType)
		if err = ie.CnType.Decode(r); err != nil {
			return utils.WrapError("Decode CnType", err)
		}
	}
	return nil
}
