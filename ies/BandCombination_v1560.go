package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1560 struct {
	Ne_DC_BC                 *BandCombination_v1560_ne_DC_BC `optional`
	Ca_ParametersNRDC        *CA_ParametersNRDC              `optional`
	Ca_ParametersEUTRA_v1560 *CA_ParametersEUTRA_v1560       `optional`
	Ca_ParametersNR_v1560    *CA_ParametersNR_v1560          `optional`
}

func (ie *BandCombination_v1560) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ne_DC_BC != nil, ie.Ca_ParametersNRDC != nil, ie.Ca_ParametersEUTRA_v1560 != nil, ie.Ca_ParametersNR_v1560 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ne_DC_BC != nil {
		if err = ie.Ne_DC_BC.Encode(w); err != nil {
			return utils.WrapError("Encode Ne_DC_BC", err)
		}
	}
	if ie.Ca_ParametersNRDC != nil {
		if err = ie.Ca_ParametersNRDC.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNRDC", err)
		}
	}
	if ie.Ca_ParametersEUTRA_v1560 != nil {
		if err = ie.Ca_ParametersEUTRA_v1560.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersEUTRA_v1560", err)
		}
	}
	if ie.Ca_ParametersNR_v1560 != nil {
		if err = ie.Ca_ParametersNR_v1560.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNR_v1560", err)
		}
	}
	return nil
}

func (ie *BandCombination_v1560) Decode(r *aper.AperReader) error {
	var err error
	var Ne_DC_BCPresent bool
	if Ne_DC_BCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ca_ParametersNRDCPresent bool
	if Ca_ParametersNRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ca_ParametersEUTRA_v1560Present bool
	if Ca_ParametersEUTRA_v1560Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ca_ParametersNR_v1560Present bool
	if Ca_ParametersNR_v1560Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ne_DC_BCPresent {
		ie.Ne_DC_BC = new(BandCombination_v1560_ne_DC_BC)
		if err = ie.Ne_DC_BC.Decode(r); err != nil {
			return utils.WrapError("Decode Ne_DC_BC", err)
		}
	}
	if Ca_ParametersNRDCPresent {
		ie.Ca_ParametersNRDC = new(CA_ParametersNRDC)
		if err = ie.Ca_ParametersNRDC.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNRDC", err)
		}
	}
	if Ca_ParametersEUTRA_v1560Present {
		ie.Ca_ParametersEUTRA_v1560 = new(CA_ParametersEUTRA_v1560)
		if err = ie.Ca_ParametersEUTRA_v1560.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersEUTRA_v1560", err)
		}
	}
	if Ca_ParametersNR_v1560Present {
		ie.Ca_ParametersNR_v1560 = new(CA_ParametersNR_v1560)
		if err = ie.Ca_ParametersNR_v1560.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNR_v1560", err)
		}
	}
	return nil
}
