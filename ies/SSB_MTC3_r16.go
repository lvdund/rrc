package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SSB_MTC3_r16 struct {
	PeriodicityAndOffset_r16 SSB_MTC3_r16_periodicityAndOffset_r16 `lb:0,ub:4,madatory`
	Duration_r16             SSB_MTC3_r16_duration_r16             `madatory`
	Pci_List_r16             []PhysCellId                          `lb:1,ub:maxNrofPCIsPerSMTC,optional`
	Ssb_ToMeasure_r16        *SSB_ToMeasure                        `optional,setuprelease`
}

func (ie *SSB_MTC3_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Pci_List_r16) > 0, ie.Ssb_ToMeasure_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.PeriodicityAndOffset_r16.Encode(w); err != nil {
		return utils.WrapError("Encode PeriodicityAndOffset_r16", err)
	}
	if err = ie.Duration_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Duration_r16", err)
	}
	if len(ie.Pci_List_r16) > 0 {
		tmp_Pci_List_r16 := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, uper.Constraint{Lb: 1, Ub: maxNrofPCIsPerSMTC}, false)
		for _, i := range ie.Pci_List_r16 {
			tmp_Pci_List_r16.Value = append(tmp_Pci_List_r16.Value, &i)
		}
		if err = tmp_Pci_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Pci_List_r16", err)
		}
	}
	if ie.Ssb_ToMeasure_r16 != nil {
		tmp_Ssb_ToMeasure_r16 := utils.SetupRelease[*SSB_ToMeasure]{
			Setup: ie.Ssb_ToMeasure_r16,
		}
		if err = tmp_Ssb_ToMeasure_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_ToMeasure_r16", err)
		}
	}
	return nil
}

func (ie *SSB_MTC3_r16) Decode(r *uper.UperReader) error {
	var err error
	var Pci_List_r16Present bool
	if Pci_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_ToMeasure_r16Present bool
	if Ssb_ToMeasure_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.PeriodicityAndOffset_r16.Decode(r); err != nil {
		return utils.WrapError("Decode PeriodicityAndOffset_r16", err)
	}
	if err = ie.Duration_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Duration_r16", err)
	}
	if Pci_List_r16Present {
		tmp_Pci_List_r16 := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, uper.Constraint{Lb: 1, Ub: maxNrofPCIsPerSMTC}, false)
		fn_Pci_List_r16 := func() *PhysCellId {
			return new(PhysCellId)
		}
		if err = tmp_Pci_List_r16.Decode(r, fn_Pci_List_r16); err != nil {
			return utils.WrapError("Decode Pci_List_r16", err)
		}
		ie.Pci_List_r16 = []PhysCellId{}
		for _, i := range tmp_Pci_List_r16.Value {
			ie.Pci_List_r16 = append(ie.Pci_List_r16, *i)
		}
	}
	if Ssb_ToMeasure_r16Present {
		tmp_Ssb_ToMeasure_r16 := utils.SetupRelease[*SSB_ToMeasure]{}
		if err = tmp_Ssb_ToMeasure_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_ToMeasure_r16", err)
		}
		ie.Ssb_ToMeasure_r16 = tmp_Ssb_ToMeasure_r16.Setup
	}
	return nil
}
