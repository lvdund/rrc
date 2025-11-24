package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LogMeasReport_r16 struct {
	AbsoluteTimeStamp_r16        AbsoluteTimeInfo_r16                        `madatory`
	TraceReference_r16           TraceReference_r16                          `madatory`
	TraceRecordingSessionRef_r16 []byte                                      `lb:2,ub:2,madatory`
	Tce_Id_r16                   []byte                                      `lb:1,ub:1,madatory`
	LogMeasInfoList_r16          LogMeasInfoList_r16                         `madatory`
	LogMeasAvailable_r16         *LogMeasReport_r16_logMeasAvailable_r16     `optional`
	LogMeasAvailableBT_r16       *LogMeasReport_r16_logMeasAvailableBT_r16   `optional`
	LogMeasAvailableWLAN_r16     *LogMeasReport_r16_logMeasAvailableWLAN_r16 `optional`
}

func (ie *LogMeasReport_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.LogMeasAvailable_r16 != nil, ie.LogMeasAvailableBT_r16 != nil, ie.LogMeasAvailableWLAN_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.AbsoluteTimeStamp_r16.Encode(w); err != nil {
		return utils.WrapError("Encode AbsoluteTimeStamp_r16", err)
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
	if ie.LogMeasAvailable_r16 != nil {
		if err = ie.LogMeasAvailable_r16.Encode(w); err != nil {
			return utils.WrapError("Encode LogMeasAvailable_r16", err)
		}
	}
	if ie.LogMeasAvailableBT_r16 != nil {
		if err = ie.LogMeasAvailableBT_r16.Encode(w); err != nil {
			return utils.WrapError("Encode LogMeasAvailableBT_r16", err)
		}
	}
	if ie.LogMeasAvailableWLAN_r16 != nil {
		if err = ie.LogMeasAvailableWLAN_r16.Encode(w); err != nil {
			return utils.WrapError("Encode LogMeasAvailableWLAN_r16", err)
		}
	}
	return nil
}

func (ie *LogMeasReport_r16) Decode(r *uper.UperReader) error {
	var err error
	var LogMeasAvailable_r16Present bool
	if LogMeasAvailable_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LogMeasAvailableBT_r16Present bool
	if LogMeasAvailableBT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LogMeasAvailableWLAN_r16Present bool
	if LogMeasAvailableWLAN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.AbsoluteTimeStamp_r16.Decode(r); err != nil {
		return utils.WrapError("Decode AbsoluteTimeStamp_r16", err)
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
	if LogMeasAvailable_r16Present {
		ie.LogMeasAvailable_r16 = new(LogMeasReport_r16_logMeasAvailable_r16)
		if err = ie.LogMeasAvailable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode LogMeasAvailable_r16", err)
		}
	}
	if LogMeasAvailableBT_r16Present {
		ie.LogMeasAvailableBT_r16 = new(LogMeasReport_r16_logMeasAvailableBT_r16)
		if err = ie.LogMeasAvailableBT_r16.Decode(r); err != nil {
			return utils.WrapError("Decode LogMeasAvailableBT_r16", err)
		}
	}
	if LogMeasAvailableWLAN_r16Present {
		ie.LogMeasAvailableWLAN_r16 = new(LogMeasReport_r16_logMeasAvailableWLAN_r16)
		if err = ie.LogMeasAvailableWLAN_r16.Decode(r); err != nil {
			return utils.WrapError("Decode LogMeasAvailableWLAN_r16", err)
		}
	}
	return nil
}
