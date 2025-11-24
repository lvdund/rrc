package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LoggedMeasurementConfiguration_r16_IEs struct {
	TraceReference_r16           TraceReference_r16                                `madatory`
	TraceRecordingSessionRef_r16 []byte                                            `lb:2,ub:2,madatory`
	Tce_Id_r16                   []byte                                            `lb:1,ub:1,madatory`
	AbsoluteTimeInfo_r16         AbsoluteTimeInfo_r16                              `madatory`
	AreaConfiguration_r16        *AreaConfiguration_r16                            `optional`
	Plmn_IdentityList_r16        *PLMN_IdentityList2_r16                           `optional`
	Bt_NameList_r16              *BT_NameList_r16                                  `optional,setuprelease`
	Wlan_NameList_r16            *WLAN_NameList_r16                                `optional,setuprelease`
	Sensor_NameList_r16          *Sensor_NameList_r16                              `optional,setuprelease`
	LoggingDuration_r16          LoggingDuration_r16                               `madatory`
	ReportType                   LoggedMeasurementConfiguration_r16_IEs_reportType `madatory`
	LateNonCriticalExtension     *[]byte                                           `optional,ext`
	NonCriticalExtension         *LoggedMeasurementConfiguration_v1700_IEs         `optional,ext`
}

func (ie *LoggedMeasurementConfiguration_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.AreaConfiguration_r16 != nil, ie.Plmn_IdentityList_r16 != nil, ie.Bt_NameList_r16 != nil, ie.Wlan_NameList_r16 != nil, ie.Sensor_NameList_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
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
	if err = ie.AbsoluteTimeInfo_r16.Encode(w); err != nil {
		return utils.WrapError("Encode AbsoluteTimeInfo_r16", err)
	}
	if ie.AreaConfiguration_r16 != nil {
		if err = ie.AreaConfiguration_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AreaConfiguration_r16", err)
		}
	}
	if ie.Plmn_IdentityList_r16 != nil {
		if err = ie.Plmn_IdentityList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Plmn_IdentityList_r16", err)
		}
	}
	if ie.Bt_NameList_r16 != nil {
		tmp_Bt_NameList_r16 := utils.SetupRelease[*BT_NameList_r16]{
			Setup: ie.Bt_NameList_r16,
		}
		if err = tmp_Bt_NameList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Bt_NameList_r16", err)
		}
	}
	if ie.Wlan_NameList_r16 != nil {
		tmp_Wlan_NameList_r16 := utils.SetupRelease[*WLAN_NameList_r16]{
			Setup: ie.Wlan_NameList_r16,
		}
		if err = tmp_Wlan_NameList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Wlan_NameList_r16", err)
		}
	}
	if ie.Sensor_NameList_r16 != nil {
		tmp_Sensor_NameList_r16 := utils.SetupRelease[*Sensor_NameList_r16]{
			Setup: ie.Sensor_NameList_r16,
		}
		if err = tmp_Sensor_NameList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sensor_NameList_r16", err)
		}
	}
	if err = ie.LoggingDuration_r16.Encode(w); err != nil {
		return utils.WrapError("Encode LoggingDuration_r16", err)
	}
	if err = ie.ReportType.Encode(w); err != nil {
		return utils.WrapError("Encode ReportType", err)
	}
	return nil
}

func (ie *LoggedMeasurementConfiguration_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var AreaConfiguration_r16Present bool
	if AreaConfiguration_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Plmn_IdentityList_r16Present bool
	if Plmn_IdentityList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Bt_NameList_r16Present bool
	if Bt_NameList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Wlan_NameList_r16Present bool
	if Wlan_NameList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sensor_NameList_r16Present bool
	if Sensor_NameList_r16Present, err = r.ReadBool(); err != nil {
		return err
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
	if err = ie.AbsoluteTimeInfo_r16.Decode(r); err != nil {
		return utils.WrapError("Decode AbsoluteTimeInfo_r16", err)
	}
	if AreaConfiguration_r16Present {
		ie.AreaConfiguration_r16 = new(AreaConfiguration_r16)
		if err = ie.AreaConfiguration_r16.Decode(r); err != nil {
			return utils.WrapError("Decode AreaConfiguration_r16", err)
		}
	}
	if Plmn_IdentityList_r16Present {
		ie.Plmn_IdentityList_r16 = new(PLMN_IdentityList2_r16)
		if err = ie.Plmn_IdentityList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Plmn_IdentityList_r16", err)
		}
	}
	if Bt_NameList_r16Present {
		tmp_Bt_NameList_r16 := utils.SetupRelease[*BT_NameList_r16]{}
		if err = tmp_Bt_NameList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Bt_NameList_r16", err)
		}
		ie.Bt_NameList_r16 = tmp_Bt_NameList_r16.Setup
	}
	if Wlan_NameList_r16Present {
		tmp_Wlan_NameList_r16 := utils.SetupRelease[*WLAN_NameList_r16]{}
		if err = tmp_Wlan_NameList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Wlan_NameList_r16", err)
		}
		ie.Wlan_NameList_r16 = tmp_Wlan_NameList_r16.Setup
	}
	if Sensor_NameList_r16Present {
		tmp_Sensor_NameList_r16 := utils.SetupRelease[*Sensor_NameList_r16]{}
		if err = tmp_Sensor_NameList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sensor_NameList_r16", err)
		}
		ie.Sensor_NameList_r16 = tmp_Sensor_NameList_r16.Setup
	}
	if err = ie.LoggingDuration_r16.Decode(r); err != nil {
		return utils.WrapError("Decode LoggingDuration_r16", err)
	}
	if err = ie.ReportType.Decode(r); err != nil {
		return utils.WrapError("Decode ReportType", err)
	}
	return nil
}
