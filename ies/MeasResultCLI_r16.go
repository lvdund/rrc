package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultCLI_r16 struct {
	MeasResultListSRS_RSRP_r16 *MeasResultListSRS_RSRP_r16 `optional`
	MeasResultListCLI_RSSI_r16 *MeasResultListCLI_RSSI_r16 `optional`
}

func (ie *MeasResultCLI_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasResultListSRS_RSRP_r16 != nil, ie.MeasResultListCLI_RSSI_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasResultListSRS_RSRP_r16 != nil {
		if err = ie.MeasResultListSRS_RSRP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultListSRS_RSRP_r16", err)
		}
	}
	if ie.MeasResultListCLI_RSSI_r16 != nil {
		if err = ie.MeasResultListCLI_RSSI_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResultListCLI_RSSI_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultCLI_r16) Decode(r *uper.UperReader) error {
	var err error
	var MeasResultListSRS_RSRP_r16Present bool
	if MeasResultListSRS_RSRP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultListCLI_RSSI_r16Present bool
	if MeasResultListCLI_RSSI_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasResultListSRS_RSRP_r16Present {
		ie.MeasResultListSRS_RSRP_r16 = new(MeasResultListSRS_RSRP_r16)
		if err = ie.MeasResultListSRS_RSRP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultListSRS_RSRP_r16", err)
		}
	}
	if MeasResultListCLI_RSSI_r16Present {
		ie.MeasResultListCLI_RSSI_r16 = new(MeasResultListCLI_RSSI_r16)
		if err = ie.MeasResultListCLI_RSSI_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResultListCLI_RSSI_r16", err)
		}
	}
	return nil
}
