package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type NR_NS_PmaxValue struct {
	AdditionalPmax             *P_Max                     `optional`
	AdditionalSpectrumEmission AdditionalSpectrumEmission `madatory`
}

func (ie *NR_NS_PmaxValue) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.AdditionalPmax != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.AdditionalPmax != nil {
		if err = ie.AdditionalPmax.Encode(w); err != nil {
			return utils.WrapError("Encode AdditionalPmax", err)
		}
	}
	if err = ie.AdditionalSpectrumEmission.Encode(w); err != nil {
		return utils.WrapError("Encode AdditionalSpectrumEmission", err)
	}
	return nil
}

func (ie *NR_NS_PmaxValue) Decode(r *aper.AperReader) error {
	var err error
	var AdditionalPmaxPresent bool
	if AdditionalPmaxPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if AdditionalPmaxPresent {
		ie.AdditionalPmax = new(P_Max)
		if err = ie.AdditionalPmax.Decode(r); err != nil {
			return utils.WrapError("Decode AdditionalPmax", err)
		}
	}
	if err = ie.AdditionalSpectrumEmission.Decode(r); err != nil {
		return utils.WrapError("Decode AdditionalSpectrumEmission", err)
	}
	return nil
}
