package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LocationInfo_r16 struct {
	CommonLocationInfo_r16  *CommonLocationInfo_r16    `optional`
	Bt_LocationInfo_r16     *LogMeasResultListBT_r16   `optional`
	Wlan_LocationInfo_r16   *LogMeasResultListWLAN_r16 `optional`
	Sensor_LocationInfo_r16 *Sensor_LocationInfo_r16   `optional`
}

func (ie *LocationInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.CommonLocationInfo_r16 != nil, ie.Bt_LocationInfo_r16 != nil, ie.Wlan_LocationInfo_r16 != nil, ie.Sensor_LocationInfo_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CommonLocationInfo_r16 != nil {
		if err = ie.CommonLocationInfo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CommonLocationInfo_r16", err)
		}
	}
	if ie.Bt_LocationInfo_r16 != nil {
		if err = ie.Bt_LocationInfo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Bt_LocationInfo_r16", err)
		}
	}
	if ie.Wlan_LocationInfo_r16 != nil {
		if err = ie.Wlan_LocationInfo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Wlan_LocationInfo_r16", err)
		}
	}
	if ie.Sensor_LocationInfo_r16 != nil {
		if err = ie.Sensor_LocationInfo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sensor_LocationInfo_r16", err)
		}
	}
	return nil
}

func (ie *LocationInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	var CommonLocationInfo_r16Present bool
	if CommonLocationInfo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Bt_LocationInfo_r16Present bool
	if Bt_LocationInfo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Wlan_LocationInfo_r16Present bool
	if Wlan_LocationInfo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sensor_LocationInfo_r16Present bool
	if Sensor_LocationInfo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if CommonLocationInfo_r16Present {
		ie.CommonLocationInfo_r16 = new(CommonLocationInfo_r16)
		if err = ie.CommonLocationInfo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CommonLocationInfo_r16", err)
		}
	}
	if Bt_LocationInfo_r16Present {
		ie.Bt_LocationInfo_r16 = new(LogMeasResultListBT_r16)
		if err = ie.Bt_LocationInfo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Bt_LocationInfo_r16", err)
		}
	}
	if Wlan_LocationInfo_r16Present {
		ie.Wlan_LocationInfo_r16 = new(LogMeasResultListWLAN_r16)
		if err = ie.Wlan_LocationInfo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Wlan_LocationInfo_r16", err)
		}
	}
	if Sensor_LocationInfo_r16Present {
		ie.Sensor_LocationInfo_r16 = new(Sensor_LocationInfo_r16)
		if err = ie.Sensor_LocationInfo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sensor_LocationInfo_r16", err)
		}
	}
	return nil
}
