package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1700 struct {
	Ca_ParametersNR_v1700                               *CA_ParametersNR_v1700   `optional`
	Ca_ParametersNRDC_v1700                             *CA_ParametersNRDC_v1700 `optional`
	Mrdc_Parameters_v1700                               *MRDC_Parameters_v1700   `optional`
	BandList_v1710                                      []BandParameters_v1710   `lb:1,ub:maxSimultaneousBands,optional`
	SupportedBandCombListPerBC_SL_RelayDiscovery_r17    *uper.BitString          `lb:1,ub:maxBandComb,optional`
	SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17 *uper.BitString          `lb:1,ub:maxBandComb,optional`
}

func (ie *BandCombination_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ca_ParametersNR_v1700 != nil, ie.Ca_ParametersNRDC_v1700 != nil, ie.Mrdc_Parameters_v1700 != nil, len(ie.BandList_v1710) > 0, ie.SupportedBandCombListPerBC_SL_RelayDiscovery_r17 != nil, ie.SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ca_ParametersNR_v1700 != nil {
		if err = ie.Ca_ParametersNR_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNR_v1700", err)
		}
	}
	if ie.Ca_ParametersNRDC_v1700 != nil {
		if err = ie.Ca_ParametersNRDC_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNRDC_v1700", err)
		}
	}
	if ie.Mrdc_Parameters_v1700 != nil {
		if err = ie.Mrdc_Parameters_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode Mrdc_Parameters_v1700", err)
		}
	}
	if len(ie.BandList_v1710) > 0 {
		tmp_BandList_v1710 := utils.NewSequence[*BandParameters_v1710]([]*BandParameters_v1710{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
		for _, i := range ie.BandList_v1710 {
			tmp_BandList_v1710.Value = append(tmp_BandList_v1710.Value, &i)
		}
		if err = tmp_BandList_v1710.Encode(w); err != nil {
			return utils.WrapError("Encode BandList_v1710", err)
		}
	}
	if ie.SupportedBandCombListPerBC_SL_RelayDiscovery_r17 != nil {
		if err = w.WriteBitString(ie.SupportedBandCombListPerBC_SL_RelayDiscovery_r17.Bytes, uint(ie.SupportedBandCombListPerBC_SL_RelayDiscovery_r17.NumBits), &uper.Constraint{Lb: 1, Ub: maxBandComb}, false); err != nil {
			return utils.WrapError("Encode SupportedBandCombListPerBC_SL_RelayDiscovery_r17", err)
		}
	}
	if ie.SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17 != nil {
		if err = w.WriteBitString(ie.SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17.Bytes, uint(ie.SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17.NumBits), &uper.Constraint{Lb: 1, Ub: maxBandComb}, false); err != nil {
			return utils.WrapError("Encode SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17", err)
		}
	}
	return nil
}

func (ie *BandCombination_v1700) Decode(r *uper.UperReader) error {
	var err error
	var Ca_ParametersNR_v1700Present bool
	if Ca_ParametersNR_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ca_ParametersNRDC_v1700Present bool
	if Ca_ParametersNRDC_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mrdc_Parameters_v1700Present bool
	if Mrdc_Parameters_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var BandList_v1710Present bool
	if BandList_v1710Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedBandCombListPerBC_SL_RelayDiscovery_r17Present bool
	if SupportedBandCombListPerBC_SL_RelayDiscovery_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17Present bool
	if SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ca_ParametersNR_v1700Present {
		ie.Ca_ParametersNR_v1700 = new(CA_ParametersNR_v1700)
		if err = ie.Ca_ParametersNR_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNR_v1700", err)
		}
	}
	if Ca_ParametersNRDC_v1700Present {
		ie.Ca_ParametersNRDC_v1700 = new(CA_ParametersNRDC_v1700)
		if err = ie.Ca_ParametersNRDC_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNRDC_v1700", err)
		}
	}
	if Mrdc_Parameters_v1700Present {
		ie.Mrdc_Parameters_v1700 = new(MRDC_Parameters_v1700)
		if err = ie.Mrdc_Parameters_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode Mrdc_Parameters_v1700", err)
		}
	}
	if BandList_v1710Present {
		tmp_BandList_v1710 := utils.NewSequence[*BandParameters_v1710]([]*BandParameters_v1710{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
		fn_BandList_v1710 := func() *BandParameters_v1710 {
			return new(BandParameters_v1710)
		}
		if err = tmp_BandList_v1710.Decode(r, fn_BandList_v1710); err != nil {
			return utils.WrapError("Decode BandList_v1710", err)
		}
		ie.BandList_v1710 = []BandParameters_v1710{}
		for _, i := range tmp_BandList_v1710.Value {
			ie.BandList_v1710 = append(ie.BandList_v1710, *i)
		}
	}
	if SupportedBandCombListPerBC_SL_RelayDiscovery_r17Present {
		var tmp_bs_SupportedBandCombListPerBC_SL_RelayDiscovery_r17 []byte
		var n_SupportedBandCombListPerBC_SL_RelayDiscovery_r17 uint
		if tmp_bs_SupportedBandCombListPerBC_SL_RelayDiscovery_r17, n_SupportedBandCombListPerBC_SL_RelayDiscovery_r17, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: maxBandComb}, false); err != nil {
			return utils.WrapError("Decode SupportedBandCombListPerBC_SL_RelayDiscovery_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_SupportedBandCombListPerBC_SL_RelayDiscovery_r17,
			NumBits: uint64(n_SupportedBandCombListPerBC_SL_RelayDiscovery_r17),
		}
		ie.SupportedBandCombListPerBC_SL_RelayDiscovery_r17 = &tmp_bitstring
	}
	if SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17Present {
		var tmp_bs_SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17 []byte
		var n_SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17 uint
		if tmp_bs_SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17, n_SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: maxBandComb}, false); err != nil {
			return utils.WrapError("Decode SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17,
			NumBits: uint64(n_SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17),
		}
		ie.SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17 = &tmp_bitstring
	}
	return nil
}
