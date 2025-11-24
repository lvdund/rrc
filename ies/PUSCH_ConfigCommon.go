package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_ConfigCommon struct {
	GroupHoppingEnabledTransformPrecoding *PUSCH_ConfigCommon_groupHoppingEnabledTransformPrecoding `optional`
	Pusch_TimeDomainAllocationList        *PUSCH_TimeDomainResourceAllocationList                   `optional`
	Msg3_DeltaPreamble                    *int64                                                    `lb:-1,ub:6,optional`
	P0_NominalWithGrant                   *int64                                                    `lb:-202,ub:24,optional`
}

func (ie *PUSCH_ConfigCommon) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.GroupHoppingEnabledTransformPrecoding != nil, ie.Pusch_TimeDomainAllocationList != nil, ie.Msg3_DeltaPreamble != nil, ie.P0_NominalWithGrant != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.GroupHoppingEnabledTransformPrecoding != nil {
		if err = ie.GroupHoppingEnabledTransformPrecoding.Encode(w); err != nil {
			return utils.WrapError("Encode GroupHoppingEnabledTransformPrecoding", err)
		}
	}
	if ie.Pusch_TimeDomainAllocationList != nil {
		if err = ie.Pusch_TimeDomainAllocationList.Encode(w); err != nil {
			return utils.WrapError("Encode Pusch_TimeDomainAllocationList", err)
		}
	}
	if ie.Msg3_DeltaPreamble != nil {
		if err = w.WriteInteger(*ie.Msg3_DeltaPreamble, &uper.Constraint{Lb: -1, Ub: 6}, false); err != nil {
			return utils.WrapError("Encode Msg3_DeltaPreamble", err)
		}
	}
	if ie.P0_NominalWithGrant != nil {
		if err = w.WriteInteger(*ie.P0_NominalWithGrant, &uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
			return utils.WrapError("Encode P0_NominalWithGrant", err)
		}
	}
	return nil
}

func (ie *PUSCH_ConfigCommon) Decode(r *uper.UperReader) error {
	var err error
	var GroupHoppingEnabledTransformPrecodingPresent bool
	if GroupHoppingEnabledTransformPrecodingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pusch_TimeDomainAllocationListPresent bool
	if Pusch_TimeDomainAllocationListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Msg3_DeltaPreamblePresent bool
	if Msg3_DeltaPreamblePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var P0_NominalWithGrantPresent bool
	if P0_NominalWithGrantPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if GroupHoppingEnabledTransformPrecodingPresent {
		ie.GroupHoppingEnabledTransformPrecoding = new(PUSCH_ConfigCommon_groupHoppingEnabledTransformPrecoding)
		if err = ie.GroupHoppingEnabledTransformPrecoding.Decode(r); err != nil {
			return utils.WrapError("Decode GroupHoppingEnabledTransformPrecoding", err)
		}
	}
	if Pusch_TimeDomainAllocationListPresent {
		ie.Pusch_TimeDomainAllocationList = new(PUSCH_TimeDomainResourceAllocationList)
		if err = ie.Pusch_TimeDomainAllocationList.Decode(r); err != nil {
			return utils.WrapError("Decode Pusch_TimeDomainAllocationList", err)
		}
	}
	if Msg3_DeltaPreamblePresent {
		var tmp_int_Msg3_DeltaPreamble int64
		if tmp_int_Msg3_DeltaPreamble, err = r.ReadInteger(&uper.Constraint{Lb: -1, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode Msg3_DeltaPreamble", err)
		}
		ie.Msg3_DeltaPreamble = &tmp_int_Msg3_DeltaPreamble
	}
	if P0_NominalWithGrantPresent {
		var tmp_int_P0_NominalWithGrant int64
		if tmp_int_P0_NominalWithGrant, err = r.ReadInteger(&uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
			return utils.WrapError("Decode P0_NominalWithGrant", err)
		}
		ie.P0_NominalWithGrant = &tmp_int_P0_NominalWithGrant
	}
	return nil
}
