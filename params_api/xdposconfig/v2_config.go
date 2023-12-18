package xdcposconfig

type V2Config struct {
	MaxMasternodes       int     `json:"maxMasternodes"`       // v2 max masternodes
	SwitchRound          uint64  `json:"switchRound"`          // v1 to v2 switch block number
	MinePeriod           int     `json:"minePeriod"`           // Miner mine period to mine a block
	TimeoutSyncThreshold int     `json:"timeoutSyncThreshold"` // send syncInfo after number of timeout
	TimeoutPeriod        int     `json:"timeoutPeriod"`        // Duration in ms
	CertThreshold        float64 `json:"certificateThreshold"` // Necessary number of messages from master nodes to form a certificate
}
