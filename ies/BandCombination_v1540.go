package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1540 struct {
	BandList_v1540        []BandParameters_v1540 `lb:1,ub:maxSimultaneousBands,madatory`
	Ca_ParametersNR_v1540 *CA_ParametersNR_v1540 `optional`
}

func (ie *BandCombination_v1540) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ca_ParametersNR_v1540 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_BandList_v1540 := utils.NewSequence[*BandParameters_v1540]([]*BandParameters_v1540{}, aper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	for _, i := range ie.BandList_v1540 {
		tmp_BandList_v1540.Value = append(tmp_BandList_v1540.Value, &i)
	}
	if err = tmp_BandList_v1540.Encode(w); err != nil {
		return utils.WrapError("Encode BandList_v1540", err)
	}
	if ie.Ca_ParametersNR_v1540 != nil {
		if err = ie.Ca_ParametersNR_v1540.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNR_v1540", err)
		}
	}
	return nil
}

func (ie *BandCombination_v1540) Decode(r *aper.AperReader) error {
	var err error
	var Ca_ParametersNR_v1540Present bool
	if Ca_ParametersNR_v1540Present, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_BandList_v1540 := utils.NewSequence[*BandParameters_v1540]([]*BandParameters_v1540{}, aper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	fn_BandList_v1540 := func() *BandParameters_v1540 {
		return new(BandParameters_v1540)
	}
	if err = tmp_BandList_v1540.Decode(r, fn_BandList_v1540); err != nil {
		return utils.WrapError("Decode BandList_v1540", err)
	}
	ie.BandList_v1540 = []BandParameters_v1540{}
	for _, i := range tmp_BandList_v1540.Value {
		ie.BandList_v1540 = append(ie.BandList_v1540, *i)
	}
	if Ca_ParametersNR_v1540Present {
		ie.Ca_ParametersNR_v1540 = new(CA_ParametersNR_v1540)
		if err = ie.Ca_ParametersNR_v1540.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNR_v1540", err)
		}
	}
	return nil
}
