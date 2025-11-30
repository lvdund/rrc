package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16_Enum_t310_Expiry                aper.Enumerated = 0
	RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16_Enum_randomAccessProblem        aper.Enumerated = 1
	RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16_Enum_rlc_MaxNumRetx             aper.Enumerated = 2
	RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16_Enum_beamFailureRecoveryFailure aper.Enumerated = 3
	RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16_Enum_lbtFailure_r16             aper.Enumerated = 4
	RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16_Enum_bh_rlfRecoveryFailure      aper.Enumerated = 5
	RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16_Enum_t312_expiry_r17            aper.Enumerated = 6
	RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16_Enum_spare1                     aper.Enumerated = 7
)

type RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16 struct {
	Value aper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16", err)
	}
	return nil
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16) Decode(r *aper.AperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RLF_Report_r16_nr_RLF_Report_r16_rlf_Cause_r16", err)
	}
	ie.Value = aper.Enumerated(v)
	return nil
}
