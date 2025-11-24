package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LocationMeasurementIndication_IEs struct {
	MeasurementIndication    *LocationMeasurementInfo `madatory,setuprelease`
	LateNonCriticalExtension *[]byte                  `optional`
	NonCriticalExtension     interface{}              `optional`
}

func (ie *LocationMeasurementIndication_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_MeasurementIndication := utils.SetupRelease[*LocationMeasurementInfo]{
		Setup: ie.MeasurementIndication,
	}
	if err = tmp_MeasurementIndication.Encode(w); err != nil {
		return utils.WrapError("Encode MeasurementIndication", err)
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *LocationMeasurementIndication_IEs) Decode(r *uper.UperReader) error {
	var err error
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_MeasurementIndication := utils.SetupRelease[*LocationMeasurementInfo]{}
	if err = tmp_MeasurementIndication.Decode(r); err != nil {
		return utils.WrapError("Decode MeasurementIndication", err)
	}
	ie.MeasurementIndication = tmp_MeasurementIndication.Setup
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
