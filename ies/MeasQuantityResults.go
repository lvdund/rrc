package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasQuantityResults struct {
	Rsrp *RSRP_Range `optional`
	Rsrq *RSRQ_Range `optional`
	Sinr *SINR_Range `optional`
}

func (ie *MeasQuantityResults) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Rsrp != nil, ie.Rsrq != nil, ie.Sinr != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Rsrp != nil {
		if err = ie.Rsrp.Encode(w); err != nil {
			return utils.WrapError("Encode Rsrp", err)
		}
	}
	if ie.Rsrq != nil {
		if err = ie.Rsrq.Encode(w); err != nil {
			return utils.WrapError("Encode Rsrq", err)
		}
	}
	if ie.Sinr != nil {
		if err = ie.Sinr.Encode(w); err != nil {
			return utils.WrapError("Encode Sinr", err)
		}
	}
	return nil
}

func (ie *MeasQuantityResults) Decode(r *uper.UperReader) error {
	var err error
	var RsrpPresent bool
	if RsrpPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RsrqPresent bool
	if RsrqPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SinrPresent bool
	if SinrPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if RsrpPresent {
		ie.Rsrp = new(RSRP_Range)
		if err = ie.Rsrp.Decode(r); err != nil {
			return utils.WrapError("Decode Rsrp", err)
		}
	}
	if RsrqPresent {
		ie.Rsrq = new(RSRQ_Range)
		if err = ie.Rsrq.Decode(r); err != nil {
			return utils.WrapError("Decode Rsrq", err)
		}
	}
	if SinrPresent {
		ie.Sinr = new(SINR_Range)
		if err = ie.Sinr.Decode(r); err != nil {
			return utils.WrapError("Decode Sinr", err)
		}
	}
	return nil
}
