package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarLogMeasConfig_r16_IEs struct {
	AreaConfiguration_r16   *AreaConfiguration_r16                            `optional`
	Bt_NameList_r16         *BT_NameList_r16                                  `optional`
	Wlan_NameList_r16       *WLAN_NameList_r16                                `optional`
	Sensor_NameList_r16     *Sensor_NameList_r16                              `optional`
	LoggingDuration_r16     LoggingDuration_r16                               `madatory`
	ReportType              VarLogMeasConfig_r16_IEs_reportType               `madatory`
	EarlyMeasIndication_r17 *VarLogMeasConfig_r16_IEs_earlyMeasIndication_r17 `optional`
	AreaConfiguration_v1700 *AreaConfiguration_v1700                          `optional`
}

func (ie *VarLogMeasConfig_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.AreaConfiguration_r16 != nil, ie.Bt_NameList_r16 != nil, ie.Wlan_NameList_r16 != nil, ie.Sensor_NameList_r16 != nil, ie.EarlyMeasIndication_r17 != nil, ie.AreaConfiguration_v1700 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.AreaConfiguration_r16 != nil {
		if err = ie.AreaConfiguration_r16.Encode(w); err != nil {
			return utils.WrapError("Encode AreaConfiguration_r16", err)
		}
	}
	if ie.Bt_NameList_r16 != nil {
		if err = ie.Bt_NameList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Bt_NameList_r16", err)
		}
	}
	if ie.Wlan_NameList_r16 != nil {
		if err = ie.Wlan_NameList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Wlan_NameList_r16", err)
		}
	}
	if ie.Sensor_NameList_r16 != nil {
		if err = ie.Sensor_NameList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sensor_NameList_r16", err)
		}
	}
	if err = ie.LoggingDuration_r16.Encode(w); err != nil {
		return utils.WrapError("Encode LoggingDuration_r16", err)
	}
	if err = ie.ReportType.Encode(w); err != nil {
		return utils.WrapError("Encode ReportType", err)
	}
	if ie.EarlyMeasIndication_r17 != nil {
		if err = ie.EarlyMeasIndication_r17.Encode(w); err != nil {
			return utils.WrapError("Encode EarlyMeasIndication_r17", err)
		}
	}
	if ie.AreaConfiguration_v1700 != nil {
		if err = ie.AreaConfiguration_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode AreaConfiguration_v1700", err)
		}
	}
	return nil
}

func (ie *VarLogMeasConfig_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var AreaConfiguration_r16Present bool
	if AreaConfiguration_r16Present, err = r.ReadBool(); err != nil {
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
	var EarlyMeasIndication_r17Present bool
	if EarlyMeasIndication_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AreaConfiguration_v1700Present bool
	if AreaConfiguration_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	if AreaConfiguration_r16Present {
		ie.AreaConfiguration_r16 = new(AreaConfiguration_r16)
		if err = ie.AreaConfiguration_r16.Decode(r); err != nil {
			return utils.WrapError("Decode AreaConfiguration_r16", err)
		}
	}
	if Bt_NameList_r16Present {
		ie.Bt_NameList_r16 = new(BT_NameList_r16)
		if err = ie.Bt_NameList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Bt_NameList_r16", err)
		}
	}
	if Wlan_NameList_r16Present {
		ie.Wlan_NameList_r16 = new(WLAN_NameList_r16)
		if err = ie.Wlan_NameList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Wlan_NameList_r16", err)
		}
	}
	if Sensor_NameList_r16Present {
		ie.Sensor_NameList_r16 = new(Sensor_NameList_r16)
		if err = ie.Sensor_NameList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sensor_NameList_r16", err)
		}
	}
	if err = ie.LoggingDuration_r16.Decode(r); err != nil {
		return utils.WrapError("Decode LoggingDuration_r16", err)
	}
	if err = ie.ReportType.Decode(r); err != nil {
		return utils.WrapError("Decode ReportType", err)
	}
	if EarlyMeasIndication_r17Present {
		ie.EarlyMeasIndication_r17 = new(VarLogMeasConfig_r16_IEs_earlyMeasIndication_r17)
		if err = ie.EarlyMeasIndication_r17.Decode(r); err != nil {
			return utils.WrapError("Decode EarlyMeasIndication_r17", err)
		}
	}
	if AreaConfiguration_v1700Present {
		ie.AreaConfiguration_v1700 = new(AreaConfiguration_v1700)
		if err = ie.AreaConfiguration_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode AreaConfiguration_v1700", err)
		}
	}
	return nil
}
