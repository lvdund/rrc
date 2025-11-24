package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlinkPerCC_v1720 struct {
	MaxModulationOrderForMulticastDataRateCalculation_r17 *FeatureSetDownlinkPerCC_v1720_maxModulationOrderForMulticastDataRateCalculation_r17 `optional`
	Fdm_BroadcastUnicast_r17                              *FeatureSetDownlinkPerCC_v1720_fdm_BroadcastUnicast_r17                              `optional`
	Fdm_MulticastUnicast_r17                              *FeatureSetDownlinkPerCC_v1720_fdm_MulticastUnicast_r17                              `optional`
}

func (ie *FeatureSetDownlinkPerCC_v1720) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MaxModulationOrderForMulticastDataRateCalculation_r17 != nil, ie.Fdm_BroadcastUnicast_r17 != nil, ie.Fdm_MulticastUnicast_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MaxModulationOrderForMulticastDataRateCalculation_r17 != nil {
		if err = ie.MaxModulationOrderForMulticastDataRateCalculation_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxModulationOrderForMulticastDataRateCalculation_r17", err)
		}
	}
	if ie.Fdm_BroadcastUnicast_r17 != nil {
		if err = ie.Fdm_BroadcastUnicast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Fdm_BroadcastUnicast_r17", err)
		}
	}
	if ie.Fdm_MulticastUnicast_r17 != nil {
		if err = ie.Fdm_MulticastUnicast_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Fdm_MulticastUnicast_r17", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlinkPerCC_v1720) Decode(r *uper.UperReader) error {
	var err error
	var MaxModulationOrderForMulticastDataRateCalculation_r17Present bool
	if MaxModulationOrderForMulticastDataRateCalculation_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Fdm_BroadcastUnicast_r17Present bool
	if Fdm_BroadcastUnicast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Fdm_MulticastUnicast_r17Present bool
	if Fdm_MulticastUnicast_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MaxModulationOrderForMulticastDataRateCalculation_r17Present {
		ie.MaxModulationOrderForMulticastDataRateCalculation_r17 = new(FeatureSetDownlinkPerCC_v1720_maxModulationOrderForMulticastDataRateCalculation_r17)
		if err = ie.MaxModulationOrderForMulticastDataRateCalculation_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxModulationOrderForMulticastDataRateCalculation_r17", err)
		}
	}
	if Fdm_BroadcastUnicast_r17Present {
		ie.Fdm_BroadcastUnicast_r17 = new(FeatureSetDownlinkPerCC_v1720_fdm_BroadcastUnicast_r17)
		if err = ie.Fdm_BroadcastUnicast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Fdm_BroadcastUnicast_r17", err)
		}
	}
	if Fdm_MulticastUnicast_r17Present {
		ie.Fdm_MulticastUnicast_r17 = new(FeatureSetDownlinkPerCC_v1720_fdm_MulticastUnicast_r17)
		if err = ie.Fdm_MulticastUnicast_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Fdm_MulticastUnicast_r17", err)
		}
	}
	return nil
}
