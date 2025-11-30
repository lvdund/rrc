package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_UplinkTxSwitch_v1700 struct {
	BandCombination_v1700                     *BandCombination_v1700                  `optional`
	SupportedBandPairListNR_v1700             []ULTxSwitchingBandPair_v1700           `lb:1,ub:maxULTxSwitchingBandPairs,optional`
	UplinkTxSwitchingBandParametersList_v1700 []UplinkTxSwitchingBandParameters_v1700 `lb:1,ub:maxSimultaneousBands,optional`
}

func (ie *BandCombination_UplinkTxSwitch_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.BandCombination_v1700 != nil, len(ie.SupportedBandPairListNR_v1700) > 0, len(ie.UplinkTxSwitchingBandParametersList_v1700) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.BandCombination_v1700 != nil {
		if err = ie.BandCombination_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode BandCombination_v1700", err)
		}
	}
	if len(ie.SupportedBandPairListNR_v1700) > 0 {
		tmp_SupportedBandPairListNR_v1700 := utils.NewSequence[*ULTxSwitchingBandPair_v1700]([]*ULTxSwitchingBandPair_v1700{}, aper.Constraint{Lb: 1, Ub: maxULTxSwitchingBandPairs}, false)
		for _, i := range ie.SupportedBandPairListNR_v1700 {
			tmp_SupportedBandPairListNR_v1700.Value = append(tmp_SupportedBandPairListNR_v1700.Value, &i)
		}
		if err = tmp_SupportedBandPairListNR_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedBandPairListNR_v1700", err)
		}
	}
	if len(ie.UplinkTxSwitchingBandParametersList_v1700) > 0 {
		tmp_UplinkTxSwitchingBandParametersList_v1700 := utils.NewSequence[*UplinkTxSwitchingBandParameters_v1700]([]*UplinkTxSwitchingBandParameters_v1700{}, aper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
		for _, i := range ie.UplinkTxSwitchingBandParametersList_v1700 {
			tmp_UplinkTxSwitchingBandParametersList_v1700.Value = append(tmp_UplinkTxSwitchingBandParametersList_v1700.Value, &i)
		}
		if err = tmp_UplinkTxSwitchingBandParametersList_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode UplinkTxSwitchingBandParametersList_v1700", err)
		}
	}
	return nil
}

func (ie *BandCombination_UplinkTxSwitch_v1700) Decode(r *aper.AperReader) error {
	var err error
	var BandCombination_v1700Present bool
	if BandCombination_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedBandPairListNR_v1700Present bool
	if SupportedBandPairListNR_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var UplinkTxSwitchingBandParametersList_v1700Present bool
	if UplinkTxSwitchingBandParametersList_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	if BandCombination_v1700Present {
		ie.BandCombination_v1700 = new(BandCombination_v1700)
		if err = ie.BandCombination_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode BandCombination_v1700", err)
		}
	}
	if SupportedBandPairListNR_v1700Present {
		tmp_SupportedBandPairListNR_v1700 := utils.NewSequence[*ULTxSwitchingBandPair_v1700]([]*ULTxSwitchingBandPair_v1700{}, aper.Constraint{Lb: 1, Ub: maxULTxSwitchingBandPairs}, false)
		fn_SupportedBandPairListNR_v1700 := func() *ULTxSwitchingBandPair_v1700 {
			return new(ULTxSwitchingBandPair_v1700)
		}
		if err = tmp_SupportedBandPairListNR_v1700.Decode(r, fn_SupportedBandPairListNR_v1700); err != nil {
			return utils.WrapError("Decode SupportedBandPairListNR_v1700", err)
		}
		ie.SupportedBandPairListNR_v1700 = []ULTxSwitchingBandPair_v1700{}
		for _, i := range tmp_SupportedBandPairListNR_v1700.Value {
			ie.SupportedBandPairListNR_v1700 = append(ie.SupportedBandPairListNR_v1700, *i)
		}
	}
	if UplinkTxSwitchingBandParametersList_v1700Present {
		tmp_UplinkTxSwitchingBandParametersList_v1700 := utils.NewSequence[*UplinkTxSwitchingBandParameters_v1700]([]*UplinkTxSwitchingBandParameters_v1700{}, aper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
		fn_UplinkTxSwitchingBandParametersList_v1700 := func() *UplinkTxSwitchingBandParameters_v1700 {
			return new(UplinkTxSwitchingBandParameters_v1700)
		}
		if err = tmp_UplinkTxSwitchingBandParametersList_v1700.Decode(r, fn_UplinkTxSwitchingBandParametersList_v1700); err != nil {
			return utils.WrapError("Decode UplinkTxSwitchingBandParametersList_v1700", err)
		}
		ie.UplinkTxSwitchingBandParametersList_v1700 = []UplinkTxSwitchingBandParameters_v1700{}
		for _, i := range tmp_UplinkTxSwitchingBandParametersList_v1700.Value {
			ie.UplinkTxSwitchingBandParametersList_v1700 = append(ie.UplinkTxSwitchingBandParametersList_v1700, *i)
		}
	}
	return nil
}
