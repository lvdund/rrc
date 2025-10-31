package ies

// RACH-ConfigCommon-ra-SupervisionInfo ::= SEQUENCE
type RachConfigcommonRaSupervisioninfo struct {
	Preambletransmax             Preambletransmax
	RaResponsewindowsize         RachConfigcommonRaSupervisioninfoRaResponsewindowsize
	MacContentionresolutiontimer RachConfigcommonRaSupervisioninfoMacContentionresolutiontimer
}
