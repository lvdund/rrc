package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CommonLocationInfo_r16 struct {
	Gnss_TOD_msec_r16      *[]byte `optional`
	LocationTimestamp_r16  *[]byte `optional`
	LocationCoordinate_r16 *[]byte `optional`
	LocationError_r16      *[]byte `optional`
	LocationSource_r16     *[]byte `optional`
	VelocityEstimate_r16   *[]byte `optional`
}

func (ie *CommonLocationInfo_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Gnss_TOD_msec_r16 != nil, ie.LocationTimestamp_r16 != nil, ie.LocationCoordinate_r16 != nil, ie.LocationError_r16 != nil, ie.LocationSource_r16 != nil, ie.VelocityEstimate_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Gnss_TOD_msec_r16 != nil {
		if err = w.WriteOctetString(*ie.Gnss_TOD_msec_r16, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode Gnss_TOD_msec_r16", err)
		}
	}
	if ie.LocationTimestamp_r16 != nil {
		if err = w.WriteOctetString(*ie.LocationTimestamp_r16, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LocationTimestamp_r16", err)
		}
	}
	if ie.LocationCoordinate_r16 != nil {
		if err = w.WriteOctetString(*ie.LocationCoordinate_r16, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LocationCoordinate_r16", err)
		}
	}
	if ie.LocationError_r16 != nil {
		if err = w.WriteOctetString(*ie.LocationError_r16, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LocationError_r16", err)
		}
	}
	if ie.LocationSource_r16 != nil {
		if err = w.WriteOctetString(*ie.LocationSource_r16, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LocationSource_r16", err)
		}
	}
	if ie.VelocityEstimate_r16 != nil {
		if err = w.WriteOctetString(*ie.VelocityEstimate_r16, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode VelocityEstimate_r16", err)
		}
	}
	return nil
}

func (ie *CommonLocationInfo_r16) Decode(r *aper.AperReader) error {
	var err error
	var Gnss_TOD_msec_r16Present bool
	if Gnss_TOD_msec_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LocationTimestamp_r16Present bool
	if LocationTimestamp_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LocationCoordinate_r16Present bool
	if LocationCoordinate_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LocationError_r16Present bool
	if LocationError_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LocationSource_r16Present bool
	if LocationSource_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var VelocityEstimate_r16Present bool
	if VelocityEstimate_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Gnss_TOD_msec_r16Present {
		var tmp_os_Gnss_TOD_msec_r16 []byte
		if tmp_os_Gnss_TOD_msec_r16, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode Gnss_TOD_msec_r16", err)
		}
		ie.Gnss_TOD_msec_r16 = &tmp_os_Gnss_TOD_msec_r16
	}
	if LocationTimestamp_r16Present {
		var tmp_os_LocationTimestamp_r16 []byte
		if tmp_os_LocationTimestamp_r16, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LocationTimestamp_r16", err)
		}
		ie.LocationTimestamp_r16 = &tmp_os_LocationTimestamp_r16
	}
	if LocationCoordinate_r16Present {
		var tmp_os_LocationCoordinate_r16 []byte
		if tmp_os_LocationCoordinate_r16, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LocationCoordinate_r16", err)
		}
		ie.LocationCoordinate_r16 = &tmp_os_LocationCoordinate_r16
	}
	if LocationError_r16Present {
		var tmp_os_LocationError_r16 []byte
		if tmp_os_LocationError_r16, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LocationError_r16", err)
		}
		ie.LocationError_r16 = &tmp_os_LocationError_r16
	}
	if LocationSource_r16Present {
		var tmp_os_LocationSource_r16 []byte
		if tmp_os_LocationSource_r16, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LocationSource_r16", err)
		}
		ie.LocationSource_r16 = &tmp_os_LocationSource_r16
	}
	if VelocityEstimate_r16Present {
		var tmp_os_VelocityEstimate_r16 []byte
		if tmp_os_VelocityEstimate_r16, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode VelocityEstimate_r16", err)
		}
		ie.VelocityEstimate_r16 = &tmp_os_VelocityEstimate_r16
	}
	return nil
}
