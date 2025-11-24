package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_ConfigBroadcast_r17 struct {
	PdschConfigList_r17                []PDSCH_ConfigPTM_r17                       `lb:1,ub:maxNrofPDSCH_ConfigPTM_r17,madatory`
	Pdsch_TimeDomainAllocationList_r17 *PDSCH_TimeDomainResourceAllocationList_r16 `optional`
	RateMatchPatternToAddModList_r17   []RateMatchPattern                          `lb:1,ub:maxNrofRateMatchPatterns,optional`
	Lte_CRS_ToMatchAround_r17          *RateMatchPatternLTE_CRS                    `optional`
	Mcs_Table_r17                      *PDSCH_ConfigBroadcast_r17_mcs_Table_r17    `optional`
	XOverhead_r17                      *PDSCH_ConfigBroadcast_r17_xOverhead_r17    `optional`
}

func (ie *PDSCH_ConfigBroadcast_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Pdsch_TimeDomainAllocationList_r17 != nil, len(ie.RateMatchPatternToAddModList_r17) > 0, ie.Lte_CRS_ToMatchAround_r17 != nil, ie.Mcs_Table_r17 != nil, ie.XOverhead_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_PdschConfigList_r17 := utils.NewSequence[*PDSCH_ConfigPTM_r17]([]*PDSCH_ConfigPTM_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPDSCH_ConfigPTM_r17}, false)
	for _, i := range ie.PdschConfigList_r17 {
		tmp_PdschConfigList_r17.Value = append(tmp_PdschConfigList_r17.Value, &i)
	}
	if err = tmp_PdschConfigList_r17.Encode(w); err != nil {
		return utils.WrapError("Encode PdschConfigList_r17", err)
	}
	if ie.Pdsch_TimeDomainAllocationList_r17 != nil {
		if err = ie.Pdsch_TimeDomainAllocationList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_TimeDomainAllocationList_r17", err)
		}
	}
	if len(ie.RateMatchPatternToAddModList_r17) > 0 {
		tmp_RateMatchPatternToAddModList_r17 := utils.NewSequence[*RateMatchPattern]([]*RateMatchPattern{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
		for _, i := range ie.RateMatchPatternToAddModList_r17 {
			tmp_RateMatchPatternToAddModList_r17.Value = append(tmp_RateMatchPatternToAddModList_r17.Value, &i)
		}
		if err = tmp_RateMatchPatternToAddModList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode RateMatchPatternToAddModList_r17", err)
		}
	}
	if ie.Lte_CRS_ToMatchAround_r17 != nil {
		if err = ie.Lte_CRS_ToMatchAround_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Lte_CRS_ToMatchAround_r17", err)
		}
	}
	if ie.Mcs_Table_r17 != nil {
		if err = ie.Mcs_Table_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Mcs_Table_r17", err)
		}
	}
	if ie.XOverhead_r17 != nil {
		if err = ie.XOverhead_r17.Encode(w); err != nil {
			return utils.WrapError("Encode XOverhead_r17", err)
		}
	}
	return nil
}

func (ie *PDSCH_ConfigBroadcast_r17) Decode(r *uper.UperReader) error {
	var err error
	var Pdsch_TimeDomainAllocationList_r17Present bool
	if Pdsch_TimeDomainAllocationList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var RateMatchPatternToAddModList_r17Present bool
	if RateMatchPatternToAddModList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Lte_CRS_ToMatchAround_r17Present bool
	if Lte_CRS_ToMatchAround_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mcs_Table_r17Present bool
	if Mcs_Table_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var XOverhead_r17Present bool
	if XOverhead_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_PdschConfigList_r17 := utils.NewSequence[*PDSCH_ConfigPTM_r17]([]*PDSCH_ConfigPTM_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPDSCH_ConfigPTM_r17}, false)
	fn_PdschConfigList_r17 := func() *PDSCH_ConfigPTM_r17 {
		return new(PDSCH_ConfigPTM_r17)
	}
	if err = tmp_PdschConfigList_r17.Decode(r, fn_PdschConfigList_r17); err != nil {
		return utils.WrapError("Decode PdschConfigList_r17", err)
	}
	ie.PdschConfigList_r17 = []PDSCH_ConfigPTM_r17{}
	for _, i := range tmp_PdschConfigList_r17.Value {
		ie.PdschConfigList_r17 = append(ie.PdschConfigList_r17, *i)
	}
	if Pdsch_TimeDomainAllocationList_r17Present {
		ie.Pdsch_TimeDomainAllocationList_r17 = new(PDSCH_TimeDomainResourceAllocationList_r16)
		if err = ie.Pdsch_TimeDomainAllocationList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_TimeDomainAllocationList_r17", err)
		}
	}
	if RateMatchPatternToAddModList_r17Present {
		tmp_RateMatchPatternToAddModList_r17 := utils.NewSequence[*RateMatchPattern]([]*RateMatchPattern{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
		fn_RateMatchPatternToAddModList_r17 := func() *RateMatchPattern {
			return new(RateMatchPattern)
		}
		if err = tmp_RateMatchPatternToAddModList_r17.Decode(r, fn_RateMatchPatternToAddModList_r17); err != nil {
			return utils.WrapError("Decode RateMatchPatternToAddModList_r17", err)
		}
		ie.RateMatchPatternToAddModList_r17 = []RateMatchPattern{}
		for _, i := range tmp_RateMatchPatternToAddModList_r17.Value {
			ie.RateMatchPatternToAddModList_r17 = append(ie.RateMatchPatternToAddModList_r17, *i)
		}
	}
	if Lte_CRS_ToMatchAround_r17Present {
		ie.Lte_CRS_ToMatchAround_r17 = new(RateMatchPatternLTE_CRS)
		if err = ie.Lte_CRS_ToMatchAround_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Lte_CRS_ToMatchAround_r17", err)
		}
	}
	if Mcs_Table_r17Present {
		ie.Mcs_Table_r17 = new(PDSCH_ConfigBroadcast_r17_mcs_Table_r17)
		if err = ie.Mcs_Table_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Mcs_Table_r17", err)
		}
	}
	if XOverhead_r17Present {
		ie.XOverhead_r17 = new(PDSCH_ConfigBroadcast_r17_xOverhead_r17)
		if err = ie.XOverhead_r17.Decode(r); err != nil {
			return utils.WrapError("Decode XOverhead_r17", err)
		}
	}
	return nil
}
