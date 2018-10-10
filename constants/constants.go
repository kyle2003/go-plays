package constants

const BASE = "https://xxgege.net"

type ReapStatus uint8

const (
	REAP_STATUS_UNKNOWN ReapStatus = iota
	REAP_STATUS__DONE
	REAP_STATUS__NOTDONE
)
