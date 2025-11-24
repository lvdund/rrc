package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarLogMeasReport_r16 struct {
	AbsoluteTimeInfo_r16         AbsoluteTimeInfo_r16                       `madatory`
	TraceReference_r16           TraceReference_r16                         `madatory`
	TraceRecordingSessionRef_r16 []byte                                     `lb:2,ub:2,madatory`
	Tce_Id_r16                   []byte                                     `lb:1,ub:1,madatory`
	LogMeasInfoList_r16          LogMeasInfoList_r16                        `madatory`
	Plmn_IdentityList_r16        PLMN_IdentityList2_r16                     `madatory`
	SigLoggedMeasType_r17        VarLogMeasReport_r16_sigLoggedMeasType_r17 `madatory`
}

func (ie *VarLogMeasReport_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.AbsoluteTimeInfo_r16.Encode(w); err != nil {
		return utils.WrapError("Encode AbsoluteTimeInfo_r16", err)
	}
	if err = ie.TraceReference_r16.Encode(w); err != nil {
		return utils.WrapError("Encode TraceReference_r16", err)
	}
	if err = w.WriteOctetString(ie.TraceRecordingSessionRef_r16, &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteOctetString TraceRecordingSessionRef_r16", err)
	}
	if err = w.WriteOctetString(ie.Tce_Id_r16, &uper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return utils.WrapError("WriteOctetString Tce_Id_r16", err)
	}
	if err = ie.LogMeasInfoList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode LogMeasInfoList_r16", err)
	}
	if err = ie.Plmn_IdentityList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Plmn_IdentityList_r16", err)
	}
	if err = ie.SigLoggedMeasType_r17.Encode(w); err != nil {
		return utils.WrapError("Encode SigLoggedMeasType_r17", err)
	}
	return nil
}

func (ie *VarLogMeasReport_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.AbsoluteTimeInfo_r16.Decode(r); err != nil {
		return utils.WrapError("Decode AbsoluteTimeInfo_r16", err)
	}
	if err = ie.TraceReference_r16.Decode(r); err != nil {
		return utils.WrapError("Decode TraceReference_r16", err)
	}
	var tmp_os_TraceRecordingSessionRef_r16 []byte
	if tmp_os_TraceRecordingSessionRef_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadOctetString TraceRecordingSessionRef_r16", err)
	}
	ie.TraceRecordingSessionRef_r16 = tmp_os_TraceRecordingSessionRef_r16
	var tmp_os_Tce_Id_r16 []byte
	if tmp_os_Tce_Id_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return utils.WrapError("ReadOctetString Tce_Id_r16", err)
	}
	ie.Tce_Id_r16 = tmp_os_Tce_Id_r16
	if err = ie.LogMeasInfoList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode LogMeasInfoList_r16", err)
	}
	if err = ie.Plmn_IdentityList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Plmn_IdentityList_r16", err)
	}
	if err = ie.SigLoggedMeasType_r17.Decode(r); err != nil {
		return utils.WrapError("Decode SigLoggedMeasType_r17", err)
	}
	return nil
}
