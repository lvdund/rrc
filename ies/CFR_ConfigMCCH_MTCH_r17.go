package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CFR_ConfigMCCH_MTCH_r17 struct {
	LocationAndBandwidthBroadcast_r17 *LocationAndBandwidthBroadcast_r17 `optional`
	Pdsch_ConfigMCCH_r17              *PDSCH_ConfigBroadcast_r17         `optional`
	CommonControlResourceSetExt_r17   *ControlResourceSet                `optional`
}

func (ie *CFR_ConfigMCCH_MTCH_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.LocationAndBandwidthBroadcast_r17 != nil, ie.Pdsch_ConfigMCCH_r17 != nil, ie.CommonControlResourceSetExt_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.LocationAndBandwidthBroadcast_r17 != nil {
		if err = ie.LocationAndBandwidthBroadcast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode LocationAndBandwidthBroadcast_r17", err)
		}
	}
	if ie.Pdsch_ConfigMCCH_r17 != nil {
		if err = ie.Pdsch_ConfigMCCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_ConfigMCCH_r17", err)
		}
	}
	if ie.CommonControlResourceSetExt_r17 != nil {
		if err = ie.CommonControlResourceSetExt_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CommonControlResourceSetExt_r17", err)
		}
	}
	return nil
}

func (ie *CFR_ConfigMCCH_MTCH_r17) Decode(r *uper.UperReader) error {
	var err error
	var LocationAndBandwidthBroadcast_r17Present bool
	if LocationAndBandwidthBroadcast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_ConfigMCCH_r17Present bool
	if Pdsch_ConfigMCCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CommonControlResourceSetExt_r17Present bool
	if CommonControlResourceSetExt_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if LocationAndBandwidthBroadcast_r17Present {
		ie.LocationAndBandwidthBroadcast_r17 = new(LocationAndBandwidthBroadcast_r17)
		if err = ie.LocationAndBandwidthBroadcast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode LocationAndBandwidthBroadcast_r17", err)
		}
	}
	if Pdsch_ConfigMCCH_r17Present {
		ie.Pdsch_ConfigMCCH_r17 = new(PDSCH_ConfigBroadcast_r17)
		if err = ie.Pdsch_ConfigMCCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_ConfigMCCH_r17", err)
		}
	}
	if CommonControlResourceSetExt_r17Present {
		ie.CommonControlResourceSetExt_r17 = new(ControlResourceSet)
		if err = ie.CommonControlResourceSetExt_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CommonControlResourceSetExt_r17", err)
		}
	}
	return nil
}
