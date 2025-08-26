package dbmodels

// AttentionTracking 注意力跟踪记录
type AttentionTracking struct {
	Id        int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	RoomId    string `gorm:"column:room_id;NOT NULL" json:"room_id"`
	UserId    string `gorm:"column:user_id;NOT NULL" json:"user_id"`
	IsActive  bool   `gorm:"column:is_active;NOT NULL" json:"is_active"`
	Timestamp int64  `gorm:"column:timestamp;NOT NULL" json:"timestamp"`
}

func (m *AttentionTracking) TableName() string {
	return "attention_tracking"
}

// UserAttentionStats 用户注意力统计
type UserAttentionStats struct {
	RoomId           string  `json:"room_id"`
	UserId           string  `json:"user_id"`
	TotalActiveTime  int64   `json:"total_active_time"`
	TotalPassiveTime int64   `json:"total_passive_time"`
	LastActiveTime   int64   `json:"last_active_time"`
	AttentionRate    float64 `json:"attention_rate"`
}
