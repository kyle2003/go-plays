package constants

const BASE = "https://xxgege.net"

type ReapStatus uint8

const (
	REAP_STATUS__UNKNOWN ReapStatus = iota
	REAP_STATUS__DONE
	REAP_STATUS__NOTDONE
)

type DownloadStatus uint8

const (
	DOWNLOAD_STATUS__UNKNOWN DownloadStatus = iota
	DOWNLOAD_STATUS__DONE
	DOWNLOAD_STATUS__NOTDONE
)

type Bool uint8

const (
	BOOL__UNKNOWN Bool = iota
	BOOL__TRUE
	BOOL__FALSE
)
