package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_FormatConfig struct {
	InterslotFrequencyHopping *PUCCH_FormatConfig_interslotFrequencyHopping `optional`
	AdditionalDMRS            *PUCCH_FormatConfig_additionalDMRS            `optional`
	MaxCodeRate               *PUCCH_MaxCodeRate                            `optional`
	NrofSlots                 *PUCCH_FormatConfig_nrofSlots                 `optional`
	Pi2BPSK                   *PUCCH_FormatConfig_pi2BPSK                   `optional`
	SimultaneousHARQ_ACK_CSI  *PUCCH_FormatConfig_simultaneousHARQ_ACK_CSI  `optional`
}

func (ie *PUCCH_FormatConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.InterslotFrequencyHopping != nil, ie.AdditionalDMRS != nil, ie.MaxCodeRate != nil, ie.NrofSlots != nil, ie.Pi2BPSK != nil, ie.SimultaneousHARQ_ACK_CSI != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.InterslotFrequencyHopping != nil {
		if err = ie.InterslotFrequencyHopping.Encode(w); err != nil {
			return utils.WrapError("Encode InterslotFrequencyHopping", err)
		}
	}
	if ie.AdditionalDMRS != nil {
		if err = ie.AdditionalDMRS.Encode(w); err != nil {
			return utils.WrapError("Encode AdditionalDMRS", err)
		}
	}
	if ie.MaxCodeRate != nil {
		if err = ie.MaxCodeRate.Encode(w); err != nil {
			return utils.WrapError("Encode MaxCodeRate", err)
		}
	}
	if ie.NrofSlots != nil {
		if err = ie.NrofSlots.Encode(w); err != nil {
			return utils.WrapError("Encode NrofSlots", err)
		}
	}
	if ie.Pi2BPSK != nil {
		if err = ie.Pi2BPSK.Encode(w); err != nil {
			return utils.WrapError("Encode Pi2BPSK", err)
		}
	}
	if ie.SimultaneousHARQ_ACK_CSI != nil {
		if err = ie.SimultaneousHARQ_ACK_CSI.Encode(w); err != nil {
			return utils.WrapError("Encode SimultaneousHARQ_ACK_CSI", err)
		}
	}
	return nil
}

func (ie *PUCCH_FormatConfig) Decode(r *uper.UperReader) error {
	var err error
	var InterslotFrequencyHoppingPresent bool
	if InterslotFrequencyHoppingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AdditionalDMRSPresent bool
	if AdditionalDMRSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxCodeRatePresent bool
	if MaxCodeRatePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NrofSlotsPresent bool
	if NrofSlotsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pi2BPSKPresent bool
	if Pi2BPSKPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SimultaneousHARQ_ACK_CSIPresent bool
	if SimultaneousHARQ_ACK_CSIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if InterslotFrequencyHoppingPresent {
		ie.InterslotFrequencyHopping = new(PUCCH_FormatConfig_interslotFrequencyHopping)
		if err = ie.InterslotFrequencyHopping.Decode(r); err != nil {
			return utils.WrapError("Decode InterslotFrequencyHopping", err)
		}
	}
	if AdditionalDMRSPresent {
		ie.AdditionalDMRS = new(PUCCH_FormatConfig_additionalDMRS)
		if err = ie.AdditionalDMRS.Decode(r); err != nil {
			return utils.WrapError("Decode AdditionalDMRS", err)
		}
	}
	if MaxCodeRatePresent {
		ie.MaxCodeRate = new(PUCCH_MaxCodeRate)
		if err = ie.MaxCodeRate.Decode(r); err != nil {
			return utils.WrapError("Decode MaxCodeRate", err)
		}
	}
	if NrofSlotsPresent {
		ie.NrofSlots = new(PUCCH_FormatConfig_nrofSlots)
		if err = ie.NrofSlots.Decode(r); err != nil {
			return utils.WrapError("Decode NrofSlots", err)
		}
	}
	if Pi2BPSKPresent {
		ie.Pi2BPSK = new(PUCCH_FormatConfig_pi2BPSK)
		if err = ie.Pi2BPSK.Decode(r); err != nil {
			return utils.WrapError("Decode Pi2BPSK", err)
		}
	}
	if SimultaneousHARQ_ACK_CSIPresent {
		ie.SimultaneousHARQ_ACK_CSI = new(PUCCH_FormatConfig_simultaneousHARQ_ACK_CSI)
		if err = ie.SimultaneousHARQ_ACK_CSI.Decode(r); err != nil {
			return utils.WrapError("Decode SimultaneousHARQ_ACK_CSI", err)
		}
	}
	return nil
}
