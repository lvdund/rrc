package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v1550 struct {
	ReducedCP_Latency    *UE_NR_Capability_v1550_reducedCP_Latency `optional`
	NonCriticalExtension *UE_NR_Capability_v1560                   `optional`
}

func (ie *UE_NR_Capability_v1550) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ReducedCP_Latency != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ReducedCP_Latency != nil {
		if err = ie.ReducedCP_Latency.Encode(w); err != nil {
			return utils.WrapError("Encode ReducedCP_Latency", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v1550) Decode(r *aper.AperReader) error {
	var err error
	var ReducedCP_LatencyPresent bool
	if ReducedCP_LatencyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ReducedCP_LatencyPresent {
		ie.ReducedCP_Latency = new(UE_NR_Capability_v1550_reducedCP_Latency)
		if err = ie.ReducedCP_Latency.Decode(r); err != nil {
			return utils.WrapError("Decode ReducedCP_Latency", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UE_NR_Capability_v1560)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
