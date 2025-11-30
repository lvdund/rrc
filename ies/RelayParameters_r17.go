package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RelayParameters_r17 struct {
	RelayUE_Operation_L2_r17                   *RelayParameters_r17_relayUE_Operation_L2_r17                   `optional`
	RemoteUE_Operation_L2_r17                  *RelayParameters_r17_remoteUE_Operation_L2_r17                  `optional`
	RemoteUE_PathSwitchToIdleInactiveRelay_r17 *RelayParameters_r17_remoteUE_PathSwitchToIdleInactiveRelay_r17 `optional`
}

func (ie *RelayParameters_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.RelayUE_Operation_L2_r17 != nil, ie.RemoteUE_Operation_L2_r17 != nil, ie.RemoteUE_PathSwitchToIdleInactiveRelay_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.RelayUE_Operation_L2_r17 != nil {
		if err = ie.RelayUE_Operation_L2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode RelayUE_Operation_L2_r17", err)
		}
	}
	if ie.RemoteUE_Operation_L2_r17 != nil {
		if err = ie.RemoteUE_Operation_L2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode RemoteUE_Operation_L2_r17", err)
		}
	}
	if ie.RemoteUE_PathSwitchToIdleInactiveRelay_r17 != nil {
		if err = ie.RemoteUE_PathSwitchToIdleInactiveRelay_r17.Encode(w); err != nil {
			return utils.WrapError("Encode RemoteUE_PathSwitchToIdleInactiveRelay_r17", err)
		}
	}
	return nil
}

func (ie *RelayParameters_r17) Decode(r *aper.AperReader) error {
	var err error
	var RelayUE_Operation_L2_r17Present bool
	if RelayUE_Operation_L2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var RemoteUE_Operation_L2_r17Present bool
	if RemoteUE_Operation_L2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var RemoteUE_PathSwitchToIdleInactiveRelay_r17Present bool
	if RemoteUE_PathSwitchToIdleInactiveRelay_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if RelayUE_Operation_L2_r17Present {
		ie.RelayUE_Operation_L2_r17 = new(RelayParameters_r17_relayUE_Operation_L2_r17)
		if err = ie.RelayUE_Operation_L2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode RelayUE_Operation_L2_r17", err)
		}
	}
	if RemoteUE_Operation_L2_r17Present {
		ie.RemoteUE_Operation_L2_r17 = new(RelayParameters_r17_remoteUE_Operation_L2_r17)
		if err = ie.RemoteUE_Operation_L2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode RemoteUE_Operation_L2_r17", err)
		}
	}
	if RemoteUE_PathSwitchToIdleInactiveRelay_r17Present {
		ie.RemoteUE_PathSwitchToIdleInactiveRelay_r17 = new(RelayParameters_r17_remoteUE_PathSwitchToIdleInactiveRelay_r17)
		if err = ie.RemoteUE_PathSwitchToIdleInactiveRelay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode RemoteUE_PathSwitchToIdleInactiveRelay_r17", err)
		}
	}
	return nil
}
