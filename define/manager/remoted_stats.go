package manager

type GetRemotedStatsResponse struct {
	Error int          `json:"error"`
	Data  RemotedStats `json:"data"`
}

type RemotedStats struct {
	DiscardedCount int `json:"discarded_count"`
	MsgSent        int `json:"msg_sent"`
	QueueSize      int `json:"queue_size"`
	CtrlMsgCount   int `json:"ctrl_msg_count"`
	EvtCount       int `json:"evt_count"`
	TCPSessions    int `json:"tcp_sessions"`
	TotalQueueSize int `json:"total_queue_size"`
}
