package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CG_SDT_Configuration_r17 struct {
	Cg_SDT_RetransmissionTimer *int64                                            `lb:1,ub:64,optional`
	Sdt_SSB_Subset_r17         *CG_SDT_Configuration_r17_sdt_SSB_Subset_r17      `lb:4,ub:4,optional`
	Sdt_SSB_PerCG_PUSCH_r17    *CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17 `optional`
	Sdt_P0_PUSCH_r17           *int64                                            `lb:-16,ub:15,optional`
	Sdt_Alpha_r17              *CG_SDT_Configuration_r17_sdt_Alpha_r17           `optional`
	Sdt_DMRS_Ports_r17         *CG_SDT_Configuration_r17_sdt_DMRS_Ports_r17      `lb:8,ub:8,optional`
	Sdt_NrofDMRS_Sequences_r17 *int64                                            `lb:1,ub:2,optional`
}

func (ie *CG_SDT_Configuration_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Cg_SDT_RetransmissionTimer != nil, ie.Sdt_SSB_Subset_r17 != nil, ie.Sdt_SSB_PerCG_PUSCH_r17 != nil, ie.Sdt_P0_PUSCH_r17 != nil, ie.Sdt_Alpha_r17 != nil, ie.Sdt_DMRS_Ports_r17 != nil, ie.Sdt_NrofDMRS_Sequences_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Cg_SDT_RetransmissionTimer != nil {
		if err = w.WriteInteger(*ie.Cg_SDT_RetransmissionTimer, &aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
			return utils.WrapError("Encode Cg_SDT_RetransmissionTimer", err)
		}
	}
	if ie.Sdt_SSB_Subset_r17 != nil {
		if err = ie.Sdt_SSB_Subset_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sdt_SSB_Subset_r17", err)
		}
	}
	if ie.Sdt_SSB_PerCG_PUSCH_r17 != nil {
		if err = ie.Sdt_SSB_PerCG_PUSCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sdt_SSB_PerCG_PUSCH_r17", err)
		}
	}
	if ie.Sdt_P0_PUSCH_r17 != nil {
		if err = w.WriteInteger(*ie.Sdt_P0_PUSCH_r17, &aper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode Sdt_P0_PUSCH_r17", err)
		}
	}
	if ie.Sdt_Alpha_r17 != nil {
		if err = ie.Sdt_Alpha_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sdt_Alpha_r17", err)
		}
	}
	if ie.Sdt_DMRS_Ports_r17 != nil {
		if err = ie.Sdt_DMRS_Ports_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sdt_DMRS_Ports_r17", err)
		}
	}
	if ie.Sdt_NrofDMRS_Sequences_r17 != nil {
		if err = w.WriteInteger(*ie.Sdt_NrofDMRS_Sequences_r17, &aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode Sdt_NrofDMRS_Sequences_r17", err)
		}
	}
	return nil
}

func (ie *CG_SDT_Configuration_r17) Decode(r *aper.AperReader) error {
	var err error
	var Cg_SDT_RetransmissionTimerPresent bool
	if Cg_SDT_RetransmissionTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Sdt_SSB_Subset_r17Present bool
	if Sdt_SSB_Subset_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sdt_SSB_PerCG_PUSCH_r17Present bool
	if Sdt_SSB_PerCG_PUSCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sdt_P0_PUSCH_r17Present bool
	if Sdt_P0_PUSCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sdt_Alpha_r17Present bool
	if Sdt_Alpha_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sdt_DMRS_Ports_r17Present bool
	if Sdt_DMRS_Ports_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sdt_NrofDMRS_Sequences_r17Present bool
	if Sdt_NrofDMRS_Sequences_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Cg_SDT_RetransmissionTimerPresent {
		var tmp_int_Cg_SDT_RetransmissionTimer int64
		if tmp_int_Cg_SDT_RetransmissionTimer, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode Cg_SDT_RetransmissionTimer", err)
		}
		ie.Cg_SDT_RetransmissionTimer = &tmp_int_Cg_SDT_RetransmissionTimer
	}
	if Sdt_SSB_Subset_r17Present {
		ie.Sdt_SSB_Subset_r17 = new(CG_SDT_Configuration_r17_sdt_SSB_Subset_r17)
		if err = ie.Sdt_SSB_Subset_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sdt_SSB_Subset_r17", err)
		}
	}
	if Sdt_SSB_PerCG_PUSCH_r17Present {
		ie.Sdt_SSB_PerCG_PUSCH_r17 = new(CG_SDT_Configuration_r17_sdt_SSB_PerCG_PUSCH_r17)
		if err = ie.Sdt_SSB_PerCG_PUSCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sdt_SSB_PerCG_PUSCH_r17", err)
		}
	}
	if Sdt_P0_PUSCH_r17Present {
		var tmp_int_Sdt_P0_PUSCH_r17 int64
		if tmp_int_Sdt_P0_PUSCH_r17, err = r.ReadInteger(&aper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Sdt_P0_PUSCH_r17", err)
		}
		ie.Sdt_P0_PUSCH_r17 = &tmp_int_Sdt_P0_PUSCH_r17
	}
	if Sdt_Alpha_r17Present {
		ie.Sdt_Alpha_r17 = new(CG_SDT_Configuration_r17_sdt_Alpha_r17)
		if err = ie.Sdt_Alpha_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sdt_Alpha_r17", err)
		}
	}
	if Sdt_DMRS_Ports_r17Present {
		ie.Sdt_DMRS_Ports_r17 = new(CG_SDT_Configuration_r17_sdt_DMRS_Ports_r17)
		if err = ie.Sdt_DMRS_Ports_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sdt_DMRS_Ports_r17", err)
		}
	}
	if Sdt_NrofDMRS_Sequences_r17Present {
		var tmp_int_Sdt_NrofDMRS_Sequences_r17 int64
		if tmp_int_Sdt_NrofDMRS_Sequences_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode Sdt_NrofDMRS_Sequences_r17", err)
		}
		ie.Sdt_NrofDMRS_Sequences_r17 = &tmp_int_Sdt_NrofDMRS_Sequences_r17
	}
	return nil
}
