package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1730 struct {
	Ca_ParametersNR_v1730   *CA_ParametersNR_v1730   `optional`
	Ca_ParametersNRDC_v1730 *CA_ParametersNRDC_v1730 `optional`
	BandList_v1730          []BandParameters_v1730   `lb:1,ub:maxSimultaneousBands,optional`
}

func (ie *BandCombination_v1730) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ca_ParametersNR_v1730 != nil, ie.Ca_ParametersNRDC_v1730 != nil, len(ie.BandList_v1730) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ca_ParametersNR_v1730 != nil {
		if err = ie.Ca_ParametersNR_v1730.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNR_v1730", err)
		}
	}
	if ie.Ca_ParametersNRDC_v1730 != nil {
		if err = ie.Ca_ParametersNRDC_v1730.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNRDC_v1730", err)
		}
	}
	if len(ie.BandList_v1730) > 0 {
		tmp_BandList_v1730 := utils.NewSequence[*BandParameters_v1730]([]*BandParameters_v1730{}, aper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
		for _, i := range ie.BandList_v1730 {
			tmp_BandList_v1730.Value = append(tmp_BandList_v1730.Value, &i)
		}
		if err = tmp_BandList_v1730.Encode(w); err != nil {
			return utils.WrapError("Encode BandList_v1730", err)
		}
	}
	return nil
}

func (ie *BandCombination_v1730) Decode(r *aper.AperReader) error {
	var err error
	var Ca_ParametersNR_v1730Present bool
	if Ca_ParametersNR_v1730Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ca_ParametersNRDC_v1730Present bool
	if Ca_ParametersNRDC_v1730Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BandList_v1730Present bool
	if BandList_v1730Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ca_ParametersNR_v1730Present {
		ie.Ca_ParametersNR_v1730 = new(CA_ParametersNR_v1730)
		if err = ie.Ca_ParametersNR_v1730.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNR_v1730", err)
		}
	}
	if Ca_ParametersNRDC_v1730Present {
		ie.Ca_ParametersNRDC_v1730 = new(CA_ParametersNRDC_v1730)
		if err = ie.Ca_ParametersNRDC_v1730.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNRDC_v1730", err)
		}
	}
	if BandList_v1730Present {
		tmp_BandList_v1730 := utils.NewSequence[*BandParameters_v1730]([]*BandParameters_v1730{}, aper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
		fn_BandList_v1730 := func() *BandParameters_v1730 {
			return new(BandParameters_v1730)
		}
		if err = tmp_BandList_v1730.Decode(r, fn_BandList_v1730); err != nil {
			return utils.WrapError("Decode BandList_v1730", err)
		}
		ie.BandList_v1730 = []BandParameters_v1730{}
		for _, i := range tmp_BandList_v1730.Value {
			ie.BandList_v1730 = append(ie.BandList_v1730, *i)
		}
	}
	return nil
}
