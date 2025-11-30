package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SidelinkParameters_r16 struct {
	SidelinkParametersNR_r16    *SidelinkParametersNR_r16    `optional`
	SidelinkParametersEUTRA_r16 *SidelinkParametersEUTRA_r16 `optional`
}

func (ie *SidelinkParameters_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SidelinkParametersNR_r16 != nil, ie.SidelinkParametersEUTRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SidelinkParametersNR_r16 != nil {
		if err = ie.SidelinkParametersNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SidelinkParametersNR_r16", err)
		}
	}
	if ie.SidelinkParametersEUTRA_r16 != nil {
		if err = ie.SidelinkParametersEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SidelinkParametersEUTRA_r16", err)
		}
	}
	return nil
}

func (ie *SidelinkParameters_r16) Decode(r *aper.AperReader) error {
	var err error
	var SidelinkParametersNR_r16Present bool
	if SidelinkParametersNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SidelinkParametersEUTRA_r16Present bool
	if SidelinkParametersEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SidelinkParametersNR_r16Present {
		ie.SidelinkParametersNR_r16 = new(SidelinkParametersNR_r16)
		if err = ie.SidelinkParametersNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SidelinkParametersNR_r16", err)
		}
	}
	if SidelinkParametersEUTRA_r16Present {
		ie.SidelinkParametersEUTRA_r16 = new(SidelinkParametersEUTRA_r16)
		if err = ie.SidelinkParametersEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SidelinkParametersEUTRA_r16", err)
		}
	}
	return nil
}
