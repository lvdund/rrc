package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16 struct {
	OverlapPDSCHsFullyFreqTime_r16       *int64                                                                                        `lb:1,ub:2,optional`
	OverlapPDSCHsInTimePartiallyFreq_r16 *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_overlapPDSCHsInTimePartiallyFreq_r16 `optional`
	OutOfOrderOperationDL_r16            *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationDL_r16            `optional`
	OutOfOrderOperationUL_r16            *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationUL_r16            `optional`
	SeparateCRS_RateMatching_r16         *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_separateCRS_RateMatching_r16         `optional`
	DefaultQCL_PerCORESETPoolIndex_r16   *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_defaultQCL_PerCORESETPoolIndex_r16   `optional`
	MaxNumberActivatedTCI_States_r16     *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_maxNumberActivatedTCI_States_r16     `optional`
}

func (ie *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.OverlapPDSCHsFullyFreqTime_r16 != nil, ie.OverlapPDSCHsInTimePartiallyFreq_r16 != nil, ie.OutOfOrderOperationDL_r16 != nil, ie.OutOfOrderOperationUL_r16 != nil, ie.SeparateCRS_RateMatching_r16 != nil, ie.DefaultQCL_PerCORESETPoolIndex_r16 != nil, ie.MaxNumberActivatedTCI_States_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.OverlapPDSCHsFullyFreqTime_r16 != nil {
		if err = w.WriteInteger(*ie.OverlapPDSCHsFullyFreqTime_r16, &aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode OverlapPDSCHsFullyFreqTime_r16", err)
		}
	}
	if ie.OverlapPDSCHsInTimePartiallyFreq_r16 != nil {
		if err = ie.OverlapPDSCHsInTimePartiallyFreq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode OverlapPDSCHsInTimePartiallyFreq_r16", err)
		}
	}
	if ie.OutOfOrderOperationDL_r16 != nil {
		if err = ie.OutOfOrderOperationDL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode OutOfOrderOperationDL_r16", err)
		}
	}
	if ie.OutOfOrderOperationUL_r16 != nil {
		if err = ie.OutOfOrderOperationUL_r16.Encode(w); err != nil {
			return utils.WrapError("Encode OutOfOrderOperationUL_r16", err)
		}
	}
	if ie.SeparateCRS_RateMatching_r16 != nil {
		if err = ie.SeparateCRS_RateMatching_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SeparateCRS_RateMatching_r16", err)
		}
	}
	if ie.DefaultQCL_PerCORESETPoolIndex_r16 != nil {
		if err = ie.DefaultQCL_PerCORESETPoolIndex_r16.Encode(w); err != nil {
			return utils.WrapError("Encode DefaultQCL_PerCORESETPoolIndex_r16", err)
		}
	}
	if ie.MaxNumberActivatedTCI_States_r16 != nil {
		if err = ie.MaxNumberActivatedTCI_States_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumberActivatedTCI_States_r16", err)
		}
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16) Decode(r *aper.AperReader) error {
	var err error
	var OverlapPDSCHsFullyFreqTime_r16Present bool
	if OverlapPDSCHsFullyFreqTime_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var OverlapPDSCHsInTimePartiallyFreq_r16Present bool
	if OverlapPDSCHsInTimePartiallyFreq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var OutOfOrderOperationDL_r16Present bool
	if OutOfOrderOperationDL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var OutOfOrderOperationUL_r16Present bool
	if OutOfOrderOperationUL_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SeparateCRS_RateMatching_r16Present bool
	if SeparateCRS_RateMatching_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DefaultQCL_PerCORESETPoolIndex_r16Present bool
	if DefaultQCL_PerCORESETPoolIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumberActivatedTCI_States_r16Present bool
	if MaxNumberActivatedTCI_States_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if OverlapPDSCHsFullyFreqTime_r16Present {
		var tmp_int_OverlapPDSCHsFullyFreqTime_r16 int64
		if tmp_int_OverlapPDSCHsFullyFreqTime_r16, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode OverlapPDSCHsFullyFreqTime_r16", err)
		}
		ie.OverlapPDSCHsFullyFreqTime_r16 = &tmp_int_OverlapPDSCHsFullyFreqTime_r16
	}
	if OverlapPDSCHsInTimePartiallyFreq_r16Present {
		ie.OverlapPDSCHsInTimePartiallyFreq_r16 = new(MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_overlapPDSCHsInTimePartiallyFreq_r16)
		if err = ie.OverlapPDSCHsInTimePartiallyFreq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode OverlapPDSCHsInTimePartiallyFreq_r16", err)
		}
	}
	if OutOfOrderOperationDL_r16Present {
		ie.OutOfOrderOperationDL_r16 = new(MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationDL_r16)
		if err = ie.OutOfOrderOperationDL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode OutOfOrderOperationDL_r16", err)
		}
	}
	if OutOfOrderOperationUL_r16Present {
		ie.OutOfOrderOperationUL_r16 = new(MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_outOfOrderOperationUL_r16)
		if err = ie.OutOfOrderOperationUL_r16.Decode(r); err != nil {
			return utils.WrapError("Decode OutOfOrderOperationUL_r16", err)
		}
	}
	if SeparateCRS_RateMatching_r16Present {
		ie.SeparateCRS_RateMatching_r16 = new(MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_separateCRS_RateMatching_r16)
		if err = ie.SeparateCRS_RateMatching_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SeparateCRS_RateMatching_r16", err)
		}
	}
	if DefaultQCL_PerCORESETPoolIndex_r16Present {
		ie.DefaultQCL_PerCORESETPoolIndex_r16 = new(MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_defaultQCL_PerCORESETPoolIndex_r16)
		if err = ie.DefaultQCL_PerCORESETPoolIndex_r16.Decode(r); err != nil {
			return utils.WrapError("Decode DefaultQCL_PerCORESETPoolIndex_r16", err)
		}
	}
	if MaxNumberActivatedTCI_States_r16Present {
		ie.MaxNumberActivatedTCI_States_r16 = new(MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16_maxNumberActivatedTCI_States_r16)
		if err = ie.MaxNumberActivatedTCI_States_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumberActivatedTCI_States_r16", err)
		}
	}
	return nil
}
