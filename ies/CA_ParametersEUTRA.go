package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersEUTRA struct {
	MultipleTimingAdvance                       *CA_ParametersEUTRA_multipleTimingAdvance          `optional`
	SimultaneousRx_Tx                           *CA_ParametersEUTRA_simultaneousRx_Tx              `optional`
	SupportedNAICS_2CRS_AP                      *aper.BitString                                    `lb:1,ub:8,optional`
	AdditionalRx_Tx_PerformanceReq              *CA_ParametersEUTRA_additionalRx_Tx_PerformanceReq `optional`
	Ue_CA_PowerClass_N                          *CA_ParametersEUTRA_ue_CA_PowerClass_N             `optional`
	SupportedBandwidthCombinationSetEUTRA_v1530 *aper.BitString                                    `lb:1,ub:32,optional`
}

func (ie *CA_ParametersEUTRA) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MultipleTimingAdvance != nil, ie.SimultaneousRx_Tx != nil, ie.SupportedNAICS_2CRS_AP != nil, ie.AdditionalRx_Tx_PerformanceReq != nil, ie.Ue_CA_PowerClass_N != nil, ie.SupportedBandwidthCombinationSetEUTRA_v1530 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MultipleTimingAdvance != nil {
		if err = ie.MultipleTimingAdvance.Encode(w); err != nil {
			return utils.WrapError("Encode MultipleTimingAdvance", err)
		}
	}
	if ie.SimultaneousRx_Tx != nil {
		if err = ie.SimultaneousRx_Tx.Encode(w); err != nil {
			return utils.WrapError("Encode SimultaneousRx_Tx", err)
		}
	}
	if ie.SupportedNAICS_2CRS_AP != nil {
		if err = w.WriteBitString(ie.SupportedNAICS_2CRS_AP.Bytes, uint(ie.SupportedNAICS_2CRS_AP.NumBits), &aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode SupportedNAICS_2CRS_AP", err)
		}
	}
	if ie.AdditionalRx_Tx_PerformanceReq != nil {
		if err = ie.AdditionalRx_Tx_PerformanceReq.Encode(w); err != nil {
			return utils.WrapError("Encode AdditionalRx_Tx_PerformanceReq", err)
		}
	}
	if ie.Ue_CA_PowerClass_N != nil {
		if err = ie.Ue_CA_PowerClass_N.Encode(w); err != nil {
			return utils.WrapError("Encode Ue_CA_PowerClass_N", err)
		}
	}
	if ie.SupportedBandwidthCombinationSetEUTRA_v1530 != nil {
		if err = w.WriteBitString(ie.SupportedBandwidthCombinationSetEUTRA_v1530.Bytes, uint(ie.SupportedBandwidthCombinationSetEUTRA_v1530.NumBits), &aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode SupportedBandwidthCombinationSetEUTRA_v1530", err)
		}
	}
	return nil
}

func (ie *CA_ParametersEUTRA) Decode(r *aper.AperReader) error {
	var err error
	var MultipleTimingAdvancePresent bool
	if MultipleTimingAdvancePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SimultaneousRx_TxPresent bool
	if SimultaneousRx_TxPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedNAICS_2CRS_APPresent bool
	if SupportedNAICS_2CRS_APPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AdditionalRx_Tx_PerformanceReqPresent bool
	if AdditionalRx_Tx_PerformanceReqPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ue_CA_PowerClass_NPresent bool
	if Ue_CA_PowerClass_NPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedBandwidthCombinationSetEUTRA_v1530Present bool
	if SupportedBandwidthCombinationSetEUTRA_v1530Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MultipleTimingAdvancePresent {
		ie.MultipleTimingAdvance = new(CA_ParametersEUTRA_multipleTimingAdvance)
		if err = ie.MultipleTimingAdvance.Decode(r); err != nil {
			return utils.WrapError("Decode MultipleTimingAdvance", err)
		}
	}
	if SimultaneousRx_TxPresent {
		ie.SimultaneousRx_Tx = new(CA_ParametersEUTRA_simultaneousRx_Tx)
		if err = ie.SimultaneousRx_Tx.Decode(r); err != nil {
			return utils.WrapError("Decode SimultaneousRx_Tx", err)
		}
	}
	if SupportedNAICS_2CRS_APPresent {
		var tmp_bs_SupportedNAICS_2CRS_AP []byte
		var n_SupportedNAICS_2CRS_AP uint
		if tmp_bs_SupportedNAICS_2CRS_AP, n_SupportedNAICS_2CRS_AP, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode SupportedNAICS_2CRS_AP", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_SupportedNAICS_2CRS_AP,
			NumBits: uint64(n_SupportedNAICS_2CRS_AP),
		}
		ie.SupportedNAICS_2CRS_AP = &tmp_bitstring
	}
	if AdditionalRx_Tx_PerformanceReqPresent {
		ie.AdditionalRx_Tx_PerformanceReq = new(CA_ParametersEUTRA_additionalRx_Tx_PerformanceReq)
		if err = ie.AdditionalRx_Tx_PerformanceReq.Decode(r); err != nil {
			return utils.WrapError("Decode AdditionalRx_Tx_PerformanceReq", err)
		}
	}
	if Ue_CA_PowerClass_NPresent {
		ie.Ue_CA_PowerClass_N = new(CA_ParametersEUTRA_ue_CA_PowerClass_N)
		if err = ie.Ue_CA_PowerClass_N.Decode(r); err != nil {
			return utils.WrapError("Decode Ue_CA_PowerClass_N", err)
		}
	}
	if SupportedBandwidthCombinationSetEUTRA_v1530Present {
		var tmp_bs_SupportedBandwidthCombinationSetEUTRA_v1530 []byte
		var n_SupportedBandwidthCombinationSetEUTRA_v1530 uint
		if tmp_bs_SupportedBandwidthCombinationSetEUTRA_v1530, n_SupportedBandwidthCombinationSetEUTRA_v1530, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode SupportedBandwidthCombinationSetEUTRA_v1530", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_SupportedBandwidthCombinationSetEUTRA_v1530,
			NumBits: uint64(n_SupportedBandwidthCombinationSetEUTRA_v1530),
		}
		ie.SupportedBandwidthCombinationSetEUTRA_v1530 = &tmp_bitstring
	}
	return nil
}
