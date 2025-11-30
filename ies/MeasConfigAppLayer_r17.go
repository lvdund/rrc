package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasConfigAppLayer_r17 struct {
	MeasConfigAppLayerId_r17           MeasConfigAppLayerId_r17                `madatory`
	MeasConfigAppLayerContainer_r17    *[]byte                                 `lb:1,ub:8000,optional`
	ServiceType_r17                    *MeasConfigAppLayer_r17_serviceType_r17 `optional`
	PauseReporting_r17                 *bool                                   `optional`
	TransmissionOfSessionStartStop_r17 *bool                                   `optional`
	Ran_VisibleParameters_r17          *RAN_VisibleParameters_r17              `optional,setuprelease`
}

func (ie *MeasConfigAppLayer_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasConfigAppLayerContainer_r17 != nil, ie.ServiceType_r17 != nil, ie.PauseReporting_r17 != nil, ie.TransmissionOfSessionStartStop_r17 != nil, ie.Ran_VisibleParameters_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.MeasConfigAppLayerId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode MeasConfigAppLayerId_r17", err)
	}
	if ie.MeasConfigAppLayerContainer_r17 != nil {
		if err = w.WriteOctetString(*ie.MeasConfigAppLayerContainer_r17, &aper.Constraint{Lb: 1, Ub: 8000}, false); err != nil {
			return utils.WrapError("Encode MeasConfigAppLayerContainer_r17", err)
		}
	}
	if ie.ServiceType_r17 != nil {
		if err = ie.ServiceType_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ServiceType_r17", err)
		}
	}
	if ie.PauseReporting_r17 != nil {
		if err = w.WriteBoolean(*ie.PauseReporting_r17); err != nil {
			return utils.WrapError("Encode PauseReporting_r17", err)
		}
	}
	if ie.TransmissionOfSessionStartStop_r17 != nil {
		if err = w.WriteBoolean(*ie.TransmissionOfSessionStartStop_r17); err != nil {
			return utils.WrapError("Encode TransmissionOfSessionStartStop_r17", err)
		}
	}
	if ie.Ran_VisibleParameters_r17 != nil {
		tmp_Ran_VisibleParameters_r17 := utils.SetupRelease[*RAN_VisibleParameters_r17]{
			Setup: ie.Ran_VisibleParameters_r17,
		}
		if err = tmp_Ran_VisibleParameters_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ran_VisibleParameters_r17", err)
		}
	}
	return nil
}

func (ie *MeasConfigAppLayer_r17) Decode(r *aper.AperReader) error {
	var err error
	var MeasConfigAppLayerContainer_r17Present bool
	if MeasConfigAppLayerContainer_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ServiceType_r17Present bool
	if ServiceType_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PauseReporting_r17Present bool
	if PauseReporting_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TransmissionOfSessionStartStop_r17Present bool
	if TransmissionOfSessionStartStop_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ran_VisibleParameters_r17Present bool
	if Ran_VisibleParameters_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.MeasConfigAppLayerId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode MeasConfigAppLayerId_r17", err)
	}
	if MeasConfigAppLayerContainer_r17Present {
		var tmp_os_MeasConfigAppLayerContainer_r17 []byte
		if tmp_os_MeasConfigAppLayerContainer_r17, err = r.ReadOctetString(&aper.Constraint{Lb: 1, Ub: 8000}, false); err != nil {
			return utils.WrapError("Decode MeasConfigAppLayerContainer_r17", err)
		}
		ie.MeasConfigAppLayerContainer_r17 = &tmp_os_MeasConfigAppLayerContainer_r17
	}
	if ServiceType_r17Present {
		ie.ServiceType_r17 = new(MeasConfigAppLayer_r17_serviceType_r17)
		if err = ie.ServiceType_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ServiceType_r17", err)
		}
	}
	if PauseReporting_r17Present {
		var tmp_bool_PauseReporting_r17 bool
		if tmp_bool_PauseReporting_r17, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode PauseReporting_r17", err)
		}
		ie.PauseReporting_r17 = &tmp_bool_PauseReporting_r17
	}
	if TransmissionOfSessionStartStop_r17Present {
		var tmp_bool_TransmissionOfSessionStartStop_r17 bool
		if tmp_bool_TransmissionOfSessionStartStop_r17, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode TransmissionOfSessionStartStop_r17", err)
		}
		ie.TransmissionOfSessionStartStop_r17 = &tmp_bool_TransmissionOfSessionStartStop_r17
	}
	if Ran_VisibleParameters_r17Present {
		tmp_Ran_VisibleParameters_r17 := utils.SetupRelease[*RAN_VisibleParameters_r17]{}
		if err = tmp_Ran_VisibleParameters_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ran_VisibleParameters_r17", err)
		}
		ie.Ran_VisibleParameters_r17 = tmp_Ran_VisibleParameters_r17.Setup
	}
	return nil
}
