package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EUTRA_NS_PmaxValue struct {
	AdditionalPmax             *int64 `lb:-30,ub:33,optional`
	AdditionalSpectrumEmission *int64 `lb:1,ub:288,optional`
}

func (ie *EUTRA_NS_PmaxValue) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.AdditionalPmax != nil, ie.AdditionalSpectrumEmission != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.AdditionalPmax != nil {
		if err = w.WriteInteger(*ie.AdditionalPmax, &uper.Constraint{Lb: -30, Ub: 33}, false); err != nil {
			return utils.WrapError("Encode AdditionalPmax", err)
		}
	}
	if ie.AdditionalSpectrumEmission != nil {
		if err = w.WriteInteger(*ie.AdditionalSpectrumEmission, &uper.Constraint{Lb: 1, Ub: 288}, false); err != nil {
			return utils.WrapError("Encode AdditionalSpectrumEmission", err)
		}
	}
	return nil
}

func (ie *EUTRA_NS_PmaxValue) Decode(r *uper.UperReader) error {
	var err error
	var AdditionalPmaxPresent bool
	if AdditionalPmaxPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AdditionalSpectrumEmissionPresent bool
	if AdditionalSpectrumEmissionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if AdditionalPmaxPresent {
		var tmp_int_AdditionalPmax int64
		if tmp_int_AdditionalPmax, err = r.ReadInteger(&uper.Constraint{Lb: -30, Ub: 33}, false); err != nil {
			return utils.WrapError("Decode AdditionalPmax", err)
		}
		ie.AdditionalPmax = &tmp_int_AdditionalPmax
	}
	if AdditionalSpectrumEmissionPresent {
		var tmp_int_AdditionalSpectrumEmission int64
		if tmp_int_AdditionalSpectrumEmission, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 288}, false); err != nil {
			return utils.WrapError("Decode AdditionalSpectrumEmission", err)
		}
		ie.AdditionalSpectrumEmission = &tmp_int_AdditionalSpectrumEmission
	}
	return nil
}
