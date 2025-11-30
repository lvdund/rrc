package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1610 struct {
	BandList_v1610                []BandParameters_v1610                      `lb:1,ub:maxSimultaneousBands,optional`
	Ca_ParametersNR_v1610         *CA_ParametersNR_v1610                      `optional`
	Ca_ParametersNRDC_v1610       *CA_ParametersNRDC_v1610                    `optional`
	PowerClass_v1610              *BandCombination_v1610_powerClass_v1610     `optional`
	PowerClassNRPart_r16          *BandCombination_v1610_powerClassNRPart_r16 `optional`
	FeatureSetCombinationDAPS_r16 *FeatureSetCombinationId                    `optional`
	Mrdc_Parameters_v1620         *MRDC_Parameters_v1620                      `optional`
}

func (ie *BandCombination_v1610) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.BandList_v1610) > 0, ie.Ca_ParametersNR_v1610 != nil, ie.Ca_ParametersNRDC_v1610 != nil, ie.PowerClass_v1610 != nil, ie.PowerClassNRPart_r16 != nil, ie.FeatureSetCombinationDAPS_r16 != nil, ie.Mrdc_Parameters_v1620 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.BandList_v1610) > 0 {
		tmp_BandList_v1610 := utils.NewSequence[*BandParameters_v1610]([]*BandParameters_v1610{}, aper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
		for _, i := range ie.BandList_v1610 {
			tmp_BandList_v1610.Value = append(tmp_BandList_v1610.Value, &i)
		}
		if err = tmp_BandList_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode BandList_v1610", err)
		}
	}
	if ie.Ca_ParametersNR_v1610 != nil {
		if err = ie.Ca_ParametersNR_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNR_v1610", err)
		}
	}
	if ie.Ca_ParametersNRDC_v1610 != nil {
		if err = ie.Ca_ParametersNRDC_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNRDC_v1610", err)
		}
	}
	if ie.PowerClass_v1610 != nil {
		if err = ie.PowerClass_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode PowerClass_v1610", err)
		}
	}
	if ie.PowerClassNRPart_r16 != nil {
		if err = ie.PowerClassNRPart_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PowerClassNRPart_r16", err)
		}
	}
	if ie.FeatureSetCombinationDAPS_r16 != nil {
		if err = ie.FeatureSetCombinationDAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode FeatureSetCombinationDAPS_r16", err)
		}
	}
	if ie.Mrdc_Parameters_v1620 != nil {
		if err = ie.Mrdc_Parameters_v1620.Encode(w); err != nil {
			return utils.WrapError("Encode Mrdc_Parameters_v1620", err)
		}
	}
	return nil
}

func (ie *BandCombination_v1610) Decode(r *aper.AperReader) error {
	var err error
	var BandList_v1610Present bool
	if BandList_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ca_ParametersNR_v1610Present bool
	if Ca_ParametersNR_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ca_ParametersNRDC_v1610Present bool
	if Ca_ParametersNRDC_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PowerClass_v1610Present bool
	if PowerClass_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PowerClassNRPart_r16Present bool
	if PowerClassNRPart_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var FeatureSetCombinationDAPS_r16Present bool
	if FeatureSetCombinationDAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mrdc_Parameters_v1620Present bool
	if Mrdc_Parameters_v1620Present, err = r.ReadBool(); err != nil {
		return err
	}
	if BandList_v1610Present {
		tmp_BandList_v1610 := utils.NewSequence[*BandParameters_v1610]([]*BandParameters_v1610{}, aper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
		fn_BandList_v1610 := func() *BandParameters_v1610 {
			return new(BandParameters_v1610)
		}
		if err = tmp_BandList_v1610.Decode(r, fn_BandList_v1610); err != nil {
			return utils.WrapError("Decode BandList_v1610", err)
		}
		ie.BandList_v1610 = []BandParameters_v1610{}
		for _, i := range tmp_BandList_v1610.Value {
			ie.BandList_v1610 = append(ie.BandList_v1610, *i)
		}
	}
	if Ca_ParametersNR_v1610Present {
		ie.Ca_ParametersNR_v1610 = new(CA_ParametersNR_v1610)
		if err = ie.Ca_ParametersNR_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNR_v1610", err)
		}
	}
	if Ca_ParametersNRDC_v1610Present {
		ie.Ca_ParametersNRDC_v1610 = new(CA_ParametersNRDC_v1610)
		if err = ie.Ca_ParametersNRDC_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNRDC_v1610", err)
		}
	}
	if PowerClass_v1610Present {
		ie.PowerClass_v1610 = new(BandCombination_v1610_powerClass_v1610)
		if err = ie.PowerClass_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode PowerClass_v1610", err)
		}
	}
	if PowerClassNRPart_r16Present {
		ie.PowerClassNRPart_r16 = new(BandCombination_v1610_powerClassNRPart_r16)
		if err = ie.PowerClassNRPart_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PowerClassNRPart_r16", err)
		}
	}
	if FeatureSetCombinationDAPS_r16Present {
		ie.FeatureSetCombinationDAPS_r16 = new(FeatureSetCombinationId)
		if err = ie.FeatureSetCombinationDAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode FeatureSetCombinationDAPS_r16", err)
		}
	}
	if Mrdc_Parameters_v1620Present {
		ie.Mrdc_Parameters_v1620 = new(MRDC_Parameters_v1620)
		if err = ie.Mrdc_Parameters_v1620.Decode(r); err != nil {
			return utils.WrapError("Decode Mrdc_Parameters_v1620", err)
		}
	}
	return nil
}
