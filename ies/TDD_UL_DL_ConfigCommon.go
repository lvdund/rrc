package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type TDD_UL_DL_ConfigCommon struct {
	ReferenceSubcarrierSpacing SubcarrierSpacing  `madatory`
	Pattern1                   TDD_UL_DL_Pattern  `madatory`
	Pattern2                   *TDD_UL_DL_Pattern `optional`
}

func (ie *TDD_UL_DL_ConfigCommon) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Pattern2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ReferenceSubcarrierSpacing.Encode(w); err != nil {
		return utils.WrapError("Encode ReferenceSubcarrierSpacing", err)
	}
	if err = ie.Pattern1.Encode(w); err != nil {
		return utils.WrapError("Encode Pattern1", err)
	}
	if ie.Pattern2 != nil {
		if err = ie.Pattern2.Encode(w); err != nil {
			return utils.WrapError("Encode Pattern2", err)
		}
	}
	return nil
}

func (ie *TDD_UL_DL_ConfigCommon) Decode(r *aper.AperReader) error {
	var err error
	var Pattern2Present bool
	if Pattern2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ReferenceSubcarrierSpacing.Decode(r); err != nil {
		return utils.WrapError("Decode ReferenceSubcarrierSpacing", err)
	}
	if err = ie.Pattern1.Decode(r); err != nil {
		return utils.WrapError("Decode Pattern1", err)
	}
	if Pattern2Present {
		ie.Pattern2 = new(TDD_UL_DL_Pattern)
		if err = ie.Pattern2.Decode(r); err != nil {
			return utils.WrapError("Decode Pattern2", err)
		}
	}
	return nil
}
