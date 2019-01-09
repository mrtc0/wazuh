package manager

type GetAnalysisdStatsResponse struct {
	Error int            `json:"error"`
	Data  AnalysisdStats `json:"data"`
}

type AnalysisdStats struct {
	ArchivesQueueSize         int `json:"archives_queue_size"`
	EventsDropped             int `json:"events_dropped"`
	AlertsQueueSize           int `json:"alerts_queue_size"`
	RuleMatchingQueueUsage    int `json:"rule_matching_queue_usage"`
	EventsProcessed           int `json:"events_processed"`
	EventQueueUsage           int `json:"event_queue_usage"`
	EventsEdps                int `json:"events_edps"`
	HostinfoEventsDecoded     int `json:"hostinfo_events_decoded"`
	SyscollectorEventsDecoded int `json:"syscollector_events_decoded"`
	RootcheckEdps             int `json:"rootcheck_edps"`
	FirewallQueueUsage        int `json:"firewall_queue_usage"`
	AlertsQueueUsage          int `json:"alerts_queue_usage"`
	FirewallQueueSize         int `json:"firewall_queue_size"`
	AlertsWritten             int `json:"alerts_written"`
	FirewallWritten           int `json:"firewall_written"`
	SyscheckQueueSize         int `json:"syscheck_queue_size"`
	EventsReceived            int `json:"events_received"`
	RootcheckQueueUsage       int `json:"rootcheck_queue_usage"`
	RootcheckEventsDecoded    int `json:"rootcheck_events_decoded"`
	RootcheckQueueSize        int `json:"rootcheck_queue_size"`
	SyscheckEdps              int `json:"syscheck_edps"`
	FtsWritten                int `json:"fts_written"`
	SyscheckQueueUsage        int `json:"syscheck_queue_usage"`
	OtherEventsEdps           int `json:"other_events_edps"`
	StatisticalQueueUsage     int `json:"statistical_queue_usage"`
	HostinfoEdps              int `json:"hostinfo_edps"`
	HostinfoQueueUsage        int `json:"hostinfo_queue_usage"`
	SyscheckEventsDecoded     int `json:"syscheck_events_decoded"`
	SyscollectorQueueUsage    int `json:"syscollector_queue_usage"`
	ArchivesQueueUsage        int `json:"archives_queue_usage"`
	StatisticalQueueSize      int `json:"statistical_queue_size"`
	TotalEventsDecoded        int `json:"total_events_decoded"`
	HostinfoQueueSize         int `json:"hostinfo_queue_size"`
	SyscollectorQueueSize     int `json:"syscollector_queue_size"`
	RuleMatchingQueueSize     int `json:"rule_matching_queue_size"`
	OtherEventsDecoded        int `json:"other_events_decoded"`
	EventQueueSize            int `json:"event_queue_size"`
	SyscollectorEdps          int `json:"syscollector_edps"`
}
