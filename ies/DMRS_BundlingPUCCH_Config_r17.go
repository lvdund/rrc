package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DMRS_BundlingPUCCH_Config_r17 struct {
	Pucch_DMRS_Bundling_r17            *DMRS_BundlingPUCCH_Config_r17_pucch_DMRS_Bundling_r17            `optional`
	Pucch_TimeDomainWindowLength_r17   *int64                                                            `lb:2,ub:8,optional`
	Pucch_WindowRestart_r17            *DMRS_BundlingPUCCH_Config_r17_pucch_WindowRestart_r17            `optional`
	Pucch_FrequencyHoppingInterval_r17 *DMRS_BundlingPUCCH_Config_r17_pucch_FrequencyHoppingInterval_r17 `optional`
}

func (ie *DMRS_BundlingPUCCH_Config_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Pucch_DMRS_Bundling_r17 != nil, ie.Pucch_TimeDomainWindowLength_r17 != nil, ie.Pucch_WindowRestart_r17 != nil, ie.Pucch_FrequencyHoppingInterval_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pucch_DMRS_Bundling_r17 != nil {
		if err = ie.Pucch_DMRS_Bundling_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_DMRS_Bundling_r17", err)
		}
	}
	if ie.Pucch_TimeDomainWindowLength_r17 != nil {
		if err = w.WriteInteger(*ie.Pucch_TimeDomainWindowLength_r17, &uper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Pucch_TimeDomainWindowLength_r17", err)
		}
	}
	if ie.Pucch_WindowRestart_r17 != nil {
		if err = ie.Pucch_WindowRestart_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_WindowRestart_r17", err)
		}
	}
	if ie.Pucch_FrequencyHoppingInterval_r17 != nil {
		if err = ie.Pucch_FrequencyHoppingInterval_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pucch_FrequencyHoppingInterval_r17", err)
		}
	}
	return nil
}

func (ie *DMRS_BundlingPUCCH_Config_r17) Decode(r *uper.UperReader) error {
	var err error
	var Pucch_DMRS_Bundling_r17Present bool
	if Pucch_DMRS_Bundling_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pucch_TimeDomainWindowLength_r17Present bool
	if Pucch_TimeDomainWindowLength_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pucch_WindowRestart_r17Present bool
	if Pucch_WindowRestart_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pucch_FrequencyHoppingInterval_r17Present bool
	if Pucch_FrequencyHoppingInterval_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Pucch_DMRS_Bundling_r17Present {
		ie.Pucch_DMRS_Bundling_r17 = new(DMRS_BundlingPUCCH_Config_r17_pucch_DMRS_Bundling_r17)
		if err = ie.Pucch_DMRS_Bundling_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_DMRS_Bundling_r17", err)
		}
	}
	if Pucch_TimeDomainWindowLength_r17Present {
		var tmp_int_Pucch_TimeDomainWindowLength_r17 int64
		if tmp_int_Pucch_TimeDomainWindowLength_r17, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Pucch_TimeDomainWindowLength_r17", err)
		}
		ie.Pucch_TimeDomainWindowLength_r17 = &tmp_int_Pucch_TimeDomainWindowLength_r17
	}
	if Pucch_WindowRestart_r17Present {
		ie.Pucch_WindowRestart_r17 = new(DMRS_BundlingPUCCH_Config_r17_pucch_WindowRestart_r17)
		if err = ie.Pucch_WindowRestart_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_WindowRestart_r17", err)
		}
	}
	if Pucch_FrequencyHoppingInterval_r17Present {
		ie.Pucch_FrequencyHoppingInterval_r17 = new(DMRS_BundlingPUCCH_Config_r17_pucch_FrequencyHoppingInterval_r17)
		if err = ie.Pucch_FrequencyHoppingInterval_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pucch_FrequencyHoppingInterval_r17", err)
		}
	}
	return nil
}
