package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNR struct {
	Dummy                                     *CA_ParametersNR_dummy                                     `optional`
	ParallelTxSRS_PUCCH_PUSCH                 *CA_ParametersNR_parallelTxSRS_PUCCH_PUSCH                 `optional`
	ParallelTxPRACH_SRS_PUCCH_PUSCH           *CA_ParametersNR_parallelTxPRACH_SRS_PUCCH_PUSCH           `optional`
	SimultaneousRxTxInterBandCA               *CA_ParametersNR_simultaneousRxTxInterBandCA               `optional`
	SimultaneousRxTxSUL                       *CA_ParametersNR_simultaneousRxTxSUL                       `optional`
	DiffNumerologyAcrossPUCCH_Group           *CA_ParametersNR_diffNumerologyAcrossPUCCH_Group           `optional`
	DiffNumerologyWithinPUCCH_GroupSmallerSCS *CA_ParametersNR_diffNumerologyWithinPUCCH_GroupSmallerSCS `optional`
	SupportedNumberTAG                        *CA_ParametersNR_supportedNumberTAG                        `optional`
}

func (ie *CA_ParametersNR) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Dummy != nil, ie.ParallelTxSRS_PUCCH_PUSCH != nil, ie.ParallelTxPRACH_SRS_PUCCH_PUSCH != nil, ie.SimultaneousRxTxInterBandCA != nil, ie.SimultaneousRxTxSUL != nil, ie.DiffNumerologyAcrossPUCCH_Group != nil, ie.DiffNumerologyWithinPUCCH_GroupSmallerSCS != nil, ie.SupportedNumberTAG != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Dummy != nil {
		if err = ie.Dummy.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy", err)
		}
	}
	if ie.ParallelTxSRS_PUCCH_PUSCH != nil {
		if err = ie.ParallelTxSRS_PUCCH_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode ParallelTxSRS_PUCCH_PUSCH", err)
		}
	}
	if ie.ParallelTxPRACH_SRS_PUCCH_PUSCH != nil {
		if err = ie.ParallelTxPRACH_SRS_PUCCH_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode ParallelTxPRACH_SRS_PUCCH_PUSCH", err)
		}
	}
	if ie.SimultaneousRxTxInterBandCA != nil {
		if err = ie.SimultaneousRxTxInterBandCA.Encode(w); err != nil {
			return utils.WrapError("Encode SimultaneousRxTxInterBandCA", err)
		}
	}
	if ie.SimultaneousRxTxSUL != nil {
		if err = ie.SimultaneousRxTxSUL.Encode(w); err != nil {
			return utils.WrapError("Encode SimultaneousRxTxSUL", err)
		}
	}
	if ie.DiffNumerologyAcrossPUCCH_Group != nil {
		if err = ie.DiffNumerologyAcrossPUCCH_Group.Encode(w); err != nil {
			return utils.WrapError("Encode DiffNumerologyAcrossPUCCH_Group", err)
		}
	}
	if ie.DiffNumerologyWithinPUCCH_GroupSmallerSCS != nil {
		if err = ie.DiffNumerologyWithinPUCCH_GroupSmallerSCS.Encode(w); err != nil {
			return utils.WrapError("Encode DiffNumerologyWithinPUCCH_GroupSmallerSCS", err)
		}
	}
	if ie.SupportedNumberTAG != nil {
		if err = ie.SupportedNumberTAG.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedNumberTAG", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNR) Decode(r *aper.AperReader) error {
	var err error
	var DummyPresent bool
	if DummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ParallelTxSRS_PUCCH_PUSCHPresent bool
	if ParallelTxSRS_PUCCH_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ParallelTxPRACH_SRS_PUCCH_PUSCHPresent bool
	if ParallelTxPRACH_SRS_PUCCH_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SimultaneousRxTxInterBandCAPresent bool
	if SimultaneousRxTxInterBandCAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SimultaneousRxTxSULPresent bool
	if SimultaneousRxTxSULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DiffNumerologyAcrossPUCCH_GroupPresent bool
	if DiffNumerologyAcrossPUCCH_GroupPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DiffNumerologyWithinPUCCH_GroupSmallerSCSPresent bool
	if DiffNumerologyWithinPUCCH_GroupSmallerSCSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportedNumberTAGPresent bool
	if SupportedNumberTAGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if DummyPresent {
		ie.Dummy = new(CA_ParametersNR_dummy)
		if err = ie.Dummy.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy", err)
		}
	}
	if ParallelTxSRS_PUCCH_PUSCHPresent {
		ie.ParallelTxSRS_PUCCH_PUSCH = new(CA_ParametersNR_parallelTxSRS_PUCCH_PUSCH)
		if err = ie.ParallelTxSRS_PUCCH_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode ParallelTxSRS_PUCCH_PUSCH", err)
		}
	}
	if ParallelTxPRACH_SRS_PUCCH_PUSCHPresent {
		ie.ParallelTxPRACH_SRS_PUCCH_PUSCH = new(CA_ParametersNR_parallelTxPRACH_SRS_PUCCH_PUSCH)
		if err = ie.ParallelTxPRACH_SRS_PUCCH_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode ParallelTxPRACH_SRS_PUCCH_PUSCH", err)
		}
	}
	if SimultaneousRxTxInterBandCAPresent {
		ie.SimultaneousRxTxInterBandCA = new(CA_ParametersNR_simultaneousRxTxInterBandCA)
		if err = ie.SimultaneousRxTxInterBandCA.Decode(r); err != nil {
			return utils.WrapError("Decode SimultaneousRxTxInterBandCA", err)
		}
	}
	if SimultaneousRxTxSULPresent {
		ie.SimultaneousRxTxSUL = new(CA_ParametersNR_simultaneousRxTxSUL)
		if err = ie.SimultaneousRxTxSUL.Decode(r); err != nil {
			return utils.WrapError("Decode SimultaneousRxTxSUL", err)
		}
	}
	if DiffNumerologyAcrossPUCCH_GroupPresent {
		ie.DiffNumerologyAcrossPUCCH_Group = new(CA_ParametersNR_diffNumerologyAcrossPUCCH_Group)
		if err = ie.DiffNumerologyAcrossPUCCH_Group.Decode(r); err != nil {
			return utils.WrapError("Decode DiffNumerologyAcrossPUCCH_Group", err)
		}
	}
	if DiffNumerologyWithinPUCCH_GroupSmallerSCSPresent {
		ie.DiffNumerologyWithinPUCCH_GroupSmallerSCS = new(CA_ParametersNR_diffNumerologyWithinPUCCH_GroupSmallerSCS)
		if err = ie.DiffNumerologyWithinPUCCH_GroupSmallerSCS.Decode(r); err != nil {
			return utils.WrapError("Decode DiffNumerologyWithinPUCCH_GroupSmallerSCS", err)
		}
	}
	if SupportedNumberTAGPresent {
		ie.SupportedNumberTAG = new(CA_ParametersNR_supportedNumberTAG)
		if err = ie.SupportedNumberTAG.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedNumberTAG", err)
		}
	}
	return nil
}
