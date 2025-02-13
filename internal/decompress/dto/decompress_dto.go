package dto

type TimestampSummary struct {
	StartTime      string `json:"start_time"`
	TransferTime   string `json:"transfer_time"`
	DecompressTime string `json:"decompress_time"`
	MergeTime      string `json:"merge_time"`
}

type DurationSummary struct {
	TransferDuration   int64 `json:"transfer_duration"`
	DecompressDuration int64 `json:"decompress_duration"`
	MergeDuration      int64 `json:"merge_duration"`
	TotalDuration      int64 `json:"total_duration"`
}
