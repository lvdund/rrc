package utils

import (
	"bytes"
	"fmt"
	"io"

	"rrc/aper"
)

func encodeTransferMessage(ies []NgapMessageIE) (w []byte, err error) {
	var buf bytes.Buffer
	aw := aper.NewWriter(&buf)
	if err = aw.WriteBool(aper.Zero); err != nil {
		return
	}
	if len(ies) == 0 {
		err = fmt.Errorf("empty message")
		return
	}

	if err = aper.WriteSequenceOf[NgapMessageIE](ies, aw, &aper.Constraint{
		Lb: 0,
		Ub: int64(aper.POW_16 - 1),
	}, false); err != nil {
		return
	}

	err = aw.Close()
	w = buf.Bytes()
	return
}

func encodeMessage(w io.Writer, present uint8, procedureCode int64, criticality aper.Enumerated, ies []NgapMessageIE) (err error) {
	aw := aper.NewWriter(w)
	if err = aw.WriteBool(aper.Zero); err != nil {
		return
	}
	if err = aw.WriteChoice(uint64(present), 2, true); err != nil {
		return
	}
	pCode := ProcedureCode{
		Value: aper.Integer(procedureCode),
	}
	if err = pCode.Encode(aw); err != nil {
		return
	}
	cr := Criticality{
		Value: criticality,
	}
	if err = cr.Encode(aw); err != nil {
		return
	}
	if len(ies) == 0 {
		err = fmt.Errorf("empty message")
		return
	}

	var buf bytes.Buffer
	cW := aper.NewWriter(&buf) //container writer
	cW.WriteBool(aper.Zero)
	if err = aper.WriteSequenceOf[NgapMessageIE](ies, cW, &aper.Constraint{
		Lb: 0,
		Ub: int64(aper.POW_16 - 1),
	}, false); err != nil {
		return
	}

	if err = cW.Close(); err != nil { //finalize
		return
	}
	if err = aw.WriteOpenType(buf.Bytes()); err != nil {
		return
	}
	err = aw.Close()
	return
}

// represent an IE in Ngap messages
type NgapMessageIE struct {
	Id          ProtocolIEID //protocol IE identity
	Criticality Criticality
	Value       aper.AperMarshaller //open type
}

func (ie NgapMessageIE) Encode(w *aper.AperWriter) (err error) {
	//1. encode protocol Ie Id
	if err = ie.Id.Encode(w); err != nil {
		return
	}
	//2. encode criticality
	if err = ie.Criticality.Encode(w); err != nil {
		return
	}
	//3. encode NgapIE
	//encode IE into a byte array first
	var buf bytes.Buffer
	ieW := aper.NewWriter(&buf)
	if err = ie.Value.Encode(ieW); err != nil {
		return
	}
	ieW.Close()
	//then write the array as open type
	err = w.WriteOpenType(buf.Bytes())
	return
}

type ProcedureCode struct {
	Value aper.Integer `aper:"valueLB:0,valueUB:255"`
}

func (ie *ProcedureCode) Decode(r *aper.AperReader) error {
	if v, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return err
	} else {
		ie.Value = aper.Integer(v)
	}
	return nil
}
func (ie *ProcedureCode) Encode(r *aper.AperWriter) (err error) {
	if err = r.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return err
	}
	return nil
}

type TriggeringMessage struct {
	Value aper.Enumerated `aper:"valueLB:0,valueUB:2"`
}

func (ie *TriggeringMessage) Decode(r *aper.AperReader) error {
	if v, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return err
	} else {
		ie.Value = aper.Enumerated(v)
	}
	return nil
}
func (ie *TriggeringMessage) Encode(r *aper.AperWriter) (err error) {
	if err = r.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return err
	}
	return nil
}

type Criticality struct {
	Value aper.Enumerated `aper:"valueLB:0,valueUB:2"`
}

func (ie *Criticality) Decode(r *aper.AperReader) error {
	if v, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return err
	} else {
		ie.Value = aper.Enumerated(v)
	}
	return nil
}
func (ie *Criticality) Encode(r *aper.AperWriter) (err error) {
	if err = r.WriteEnumerate(uint64(ie.Value), aper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return err
	}
	return nil
}

type ProtocolIEID struct {
	Value aper.Integer `aper:"valueLB:0,valueUB:65535"`
}

func (ie *ProtocolIEID) Decode(r *aper.AperReader) error {
	if v, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return err
	} else {
		ie.Value = aper.Integer(v)
	}
	return nil
}
func (ie *ProtocolIEID) Encode(r *aper.AperWriter) (err error) {
	if err = r.WriteInteger(int64(ie.Value), &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return err
	}
	return nil
}

func msgErrors(err1, err2 error) error {
	if err1 == nil && err2 == nil {
		return nil
	}
	if err1 == nil {
		return err2
	}
	if err2 == nil {
		return err1
	}
	return fmt.Errorf("%v: %v", err1, err2)
}

// NgapPdu - Present
const (
	NgapPresentNothing uint8 = iota
	NgapPduInitiatingMessage
	NgapPduSuccessfulOutcome
	NgapPduUnsuccessfulOutcome
)

