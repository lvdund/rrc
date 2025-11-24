package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DMRS_BundlingPUSCH_Config_r17 struct {
	Pusch_DMRS_Bundling_r17            *DMRS_BundlingPUSCH_Config_r17_pusch_DMRS_Bundling_r17            `optional`
	Pusch_TimeDomainWindowLength_r17   *int64                                                            `lb:2,ub:32,optional`
	Pusch_WindowRestart_r17            *DMRS_BundlingPUSCH_Config_r17_pusch_WindowRestart_r17            `optional`
	Pusch_FrequencyHoppingInterval_r17 *DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17 `optional`
}

func (ie *DMRS_BundlingPUSCH_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Pusch_DMRS_Bundling_r17 != nil, ie.Pusch_TimeDomainWindowLength_r17 != nil, ie.Pusch_WindowRestart_r17 != nil, ie.Pusch_FrequencyHoppingInterval_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pusch_DMRS_Bundling_r17 != nil {
		if err = ie.Pusch_DMRS_Bundling_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_DMRS_Bundling_r17", err)
		}
	}
	if ie.Pusch_TimeDomainWindowLength_r17 != nil {
		if err = w.WriteInteger(*ie.Pusch_TimeDomainWindowLength_r17, &uper.Constraint{Lb: 2, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode Pusch_TimeDomainWindowLength_r17", err)
		}
	}
	if ie.Pusch_WindowRestart_r17 != nil {
		if err = ie.Pusch_WindowRestart_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_WindowRestart_r17", err)
		}
	}
	if ie.Pusch_FrequencyHoppingInterval_r17 != nil {
		if err = ie.Pusch_FrequencyHoppingInterval_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_FrequencyHoppingInterval_r17", err)
		}
	}
	return nil
}

func (ie *DMRS_BundlingPUSCH_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var Pusch_DMRS_Bundling_r17Present bool
	if Pusch_DMRS_Bundling_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_TimeDomainWindowLength_r17Present bool
	if Pusch_TimeDomainWindowLength_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_WindowRestart_r17Present bool
	if Pusch_WindowRestart_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_FrequencyHoppingInterval_r17Present bool
	if Pusch_FrequencyHoppingInterval_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Pusch_DMRS_Bundling_r17Present {
		ie.Pusch_DMRS_Bundling_r17 = new(DMRS_BundlingPUSCH_Config_r17_pusch_DMRS_Bundling_r17)
		if err = ie.Pusch_DMRS_Bundling_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_DMRS_Bundling_r17", err)
		}
	}
	if Pusch_TimeDomainWindowLength_r17Present {
		var tmp_int_Pusch_TimeDomainWindowLength_r17 int64
		if tmp_int_Pusch_TimeDomainWindowLength_r17, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode Pusch_TimeDomainWindowLength_r17", err)
		}
		ie.Pusch_TimeDomainWindowLength_r17 = &tmp_int_Pusch_TimeDomainWindowLength_r17
	}
	if Pusch_WindowRestart_r17Present {
		ie.Pusch_WindowRestart_r17 = new(DMRS_BundlingPUSCH_Config_r17_pusch_WindowRestart_r17)
		if err = ie.Pusch_WindowRestart_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_WindowRestart_r17", err)
		}
	}
	if Pusch_FrequencyHoppingInterval_r17Present {
		ie.Pusch_FrequencyHoppingInterval_r17 = new(DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17)
		if err = ie.Pusch_FrequencyHoppingInterval_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_FrequencyHoppingInterval_r17", err)
		}
	}
	return nil
}
