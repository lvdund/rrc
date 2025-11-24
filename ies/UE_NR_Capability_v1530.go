package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v1530 struct {
	Fdd_Add_UE_NR_Capabilities_v1530 *UE_NR_CapabilityAddXDD_Mode_v1530           `optional`
	Tdd_Add_UE_NR_Capabilities_v1530 *UE_NR_CapabilityAddXDD_Mode_v1530           `optional`
	Dummy                            *UE_NR_Capability_v1530_dummy                `optional`
	InterRAT_Parameters              *InterRAT_Parameters                         `optional`
	InactiveState                    *UE_NR_Capability_v1530_inactiveState        `optional`
	DelayBudgetReporting             *UE_NR_Capability_v1530_delayBudgetReporting `optional`
	NonCriticalExtension             *UE_NR_Capability_v1540                      `optional`
}

func (ie *UE_NR_Capability_v1530) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Fdd_Add_UE_NR_Capabilities_v1530 != nil, ie.Tdd_Add_UE_NR_Capabilities_v1530 != nil, ie.Dummy != nil, ie.InterRAT_Parameters != nil, ie.InactiveState != nil, ie.DelayBudgetReporting != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Fdd_Add_UE_NR_Capabilities_v1530 != nil {
		if err = ie.Fdd_Add_UE_NR_Capabilities_v1530.Encode(w); err != nil {
			return utils.WrapError("Encode Fdd_Add_UE_NR_Capabilities_v1530", err)
		}
	}
	if ie.Tdd_Add_UE_NR_Capabilities_v1530 != nil {
		if err = ie.Tdd_Add_UE_NR_Capabilities_v1530.Encode(w); err != nil {
			return utils.WrapError("Encode Tdd_Add_UE_NR_Capabilities_v1530", err)
		}
	}
	if ie.Dummy != nil {
		if err = ie.Dummy.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy", err)
		}
	}
	if ie.InterRAT_Parameters != nil {
		if err = ie.InterRAT_Parameters.Encode(w); err != nil {
			return utils.WrapError("Encode InterRAT_Parameters", err)
		}
	}
	if ie.InactiveState != nil {
		if err = ie.InactiveState.Encode(w); err != nil {
			return utils.WrapError("Encode InactiveState", err)
		}
	}
	if ie.DelayBudgetReporting != nil {
		if err = ie.DelayBudgetReporting.Encode(w); err != nil {
			return utils.WrapError("Encode DelayBudgetReporting", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v1530) Decode(r *uper.UperReader) error {
	var err error
	var Fdd_Add_UE_NR_Capabilities_v1530Present bool
	if Fdd_Add_UE_NR_Capabilities_v1530Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Tdd_Add_UE_NR_Capabilities_v1530Present bool
	if Tdd_Add_UE_NR_Capabilities_v1530Present, err = r.ReadBool(); err != nil {
		return err
	}
	var DummyPresent bool
	if DummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var InterRAT_ParametersPresent bool
	if InterRAT_ParametersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var InactiveStatePresent bool
	if InactiveStatePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var DelayBudgetReportingPresent bool
	if DelayBudgetReportingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Fdd_Add_UE_NR_Capabilities_v1530Present {
		ie.Fdd_Add_UE_NR_Capabilities_v1530 = new(UE_NR_CapabilityAddXDD_Mode_v1530)
		if err = ie.Fdd_Add_UE_NR_Capabilities_v1530.Decode(r); err != nil {
			return utils.WrapError("Decode Fdd_Add_UE_NR_Capabilities_v1530", err)
		}
	}
	if Tdd_Add_UE_NR_Capabilities_v1530Present {
		ie.Tdd_Add_UE_NR_Capabilities_v1530 = new(UE_NR_CapabilityAddXDD_Mode_v1530)
		if err = ie.Tdd_Add_UE_NR_Capabilities_v1530.Decode(r); err != nil {
			return utils.WrapError("Decode Tdd_Add_UE_NR_Capabilities_v1530", err)
		}
	}
	if DummyPresent {
		ie.Dummy = new(UE_NR_Capability_v1530_dummy)
		if err = ie.Dummy.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy", err)
		}
	}
	if InterRAT_ParametersPresent {
		ie.InterRAT_Parameters = new(InterRAT_Parameters)
		if err = ie.InterRAT_Parameters.Decode(r); err != nil {
			return utils.WrapError("Decode InterRAT_Parameters", err)
		}
	}
	if InactiveStatePresent {
		ie.InactiveState = new(UE_NR_Capability_v1530_inactiveState)
		if err = ie.InactiveState.Decode(r); err != nil {
			return utils.WrapError("Decode InactiveState", err)
		}
	}
	if DelayBudgetReportingPresent {
		ie.DelayBudgetReporting = new(UE_NR_Capability_v1530_delayBudgetReporting)
		if err = ie.DelayBudgetReporting.Decode(r); err != nil {
			return utils.WrapError("Decode DelayBudgetReporting", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(UE_NR_Capability_v1540)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
