package consts

import "time"

const (
	RPCSETTINGS_MAXCAP              = 10
	RPCSETTINGS_DIAL_TIMEOUT        = 10 * time.Second
	RPCSETTINGS_DIAL_MAXTRY         = 3
	RPCSETTINGS_DIAL_RETRY_INTERVAL = 3 * time.Second
)
