package ies

// RRCSetup-v1700-IEs ::= SEQUENCE
type RrcsetupV1700 struct {
	SlConfigdedicatednrR17 *SlConfigdedicatednrR16
	SlL2remoteueConfigR17  *SlL2remoteueConfigR17
	Noncriticalextension   *RrcsetupV1700IesNoncriticalextension
}
