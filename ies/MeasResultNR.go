package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultNR struct {
	PhysCellId         *PhysCellId                      `optional`
	MeasResult         *MeasResultNR_measResult         `optional`
	Cgi_Info           *CGI_InfoNR                      `optional,ext-1`
	ChoCandidate_r17   *MeasResultNR_choCandidate_r17   `optional,ext-2`
	ChoConfig_r17      []CondTriggerConfig_r16          `lb:1,ub:2,optional,ext-2`
	TriggeredEvent_r17 *MeasResultNR_triggeredEvent_r17 `optional,ext-2`
}

func (ie *MeasResultNR) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.Cgi_Info != nil || ie.ChoCandidate_r17 != nil || len(ie.ChoConfig_r17) > 0 || ie.TriggeredEvent_r17 != nil
	preambleBits := []bool{hasExtensions, ie.PhysCellId != nil, ie.MeasResult != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.PhysCellId != nil {
		if err = ie.PhysCellId.Encode(w); err != nil {
			return utils.WrapError("Encode PhysCellId", err)
		}
	}
	if ie.MeasResult != nil {
		if err = ie.MeasResult.Encode(w); err != nil {
			return utils.WrapError("Encode MeasResult", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.Cgi_Info != nil, ie.ChoCandidate_r17 != nil || len(ie.ChoConfig_r17) > 0 || ie.TriggeredEvent_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasResultNR", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Cgi_Info != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Cgi_Info optional
			if ie.Cgi_Info != nil {
				if err = ie.Cgi_Info.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Cgi_Info", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.ChoCandidate_r17 != nil, len(ie.ChoConfig_r17) > 0, ie.TriggeredEvent_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ChoCandidate_r17 optional
			if ie.ChoCandidate_r17 != nil {
				if err = ie.ChoCandidate_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ChoCandidate_r17", err)
				}
			}
			// encode ChoConfig_r17 optional
			if len(ie.ChoConfig_r17) > 0 {
				tmp_ChoConfig_r17 := utils.NewSequence[*CondTriggerConfig_r16]([]*CondTriggerConfig_r16{}, uper.Constraint{Lb: 1, Ub: 2}, false)
				for _, i := range ie.ChoConfig_r17 {
					tmp_ChoConfig_r17.Value = append(tmp_ChoConfig_r17.Value, &i)
				}
				if err = tmp_ChoConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ChoConfig_r17", err)
				}
			}
			// encode TriggeredEvent_r17 optional
			if ie.TriggeredEvent_r17 != nil {
				if err = ie.TriggeredEvent_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode TriggeredEvent_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *MeasResultNR) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var PhysCellIdPresent bool
	if PhysCellIdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasResultPresent bool
	if MeasResultPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if PhysCellIdPresent {
		ie.PhysCellId = new(PhysCellId)
		if err = ie.PhysCellId.Decode(r); err != nil {
			return utils.WrapError("Decode PhysCellId", err)
		}
	}
	if MeasResultPresent {
		ie.MeasResult = new(MeasResultNR_measResult)
		if err = ie.MeasResult.Decode(r); err != nil {
			return utils.WrapError("Decode MeasResult", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Cgi_InfoPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Cgi_Info optional
			if Cgi_InfoPresent {
				ie.Cgi_Info = new(CGI_InfoNR)
				if err = ie.Cgi_Info.Decode(extReader); err != nil {
					return utils.WrapError("Decode Cgi_Info", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ChoCandidate_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ChoConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			TriggeredEvent_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ChoCandidate_r17 optional
			if ChoCandidate_r17Present {
				ie.ChoCandidate_r17 = new(MeasResultNR_choCandidate_r17)
				if err = ie.ChoCandidate_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ChoCandidate_r17", err)
				}
			}
			// decode ChoConfig_r17 optional
			if ChoConfig_r17Present {
				tmp_ChoConfig_r17 := utils.NewSequence[*CondTriggerConfig_r16]([]*CondTriggerConfig_r16{}, uper.Constraint{Lb: 1, Ub: 2}, false)
				fn_ChoConfig_r17 := func() *CondTriggerConfig_r16 {
					return new(CondTriggerConfig_r16)
				}
				if err = tmp_ChoConfig_r17.Decode(extReader, fn_ChoConfig_r17); err != nil {
					return utils.WrapError("Decode ChoConfig_r17", err)
				}
				ie.ChoConfig_r17 = []CondTriggerConfig_r16{}
				for _, i := range tmp_ChoConfig_r17.Value {
					ie.ChoConfig_r17 = append(ie.ChoConfig_r17, *i)
				}
			}
			// decode TriggeredEvent_r17 optional
			if TriggeredEvent_r17Present {
				ie.TriggeredEvent_r17 = new(MeasResultNR_triggeredEvent_r17)
				if err = ie.TriggeredEvent_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode TriggeredEvent_r17", err)
				}
			}
		}
	}
	return nil
}
