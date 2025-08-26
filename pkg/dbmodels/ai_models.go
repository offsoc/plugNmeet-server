package dbmodels

// MeetingSummary 会议摘要
type MeetingSummary struct {
	Id          int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	RoomId      string `gorm:"column:room_id;NOT NULL" json:"room_id"`
	SummaryText string `gorm:"column:summary_text;type:text" json:"summary_text"`
	KeyPoints   string `gorm:"column:key_points;type:text" json:"key_points"`
	ActionItems string `gorm:"column:action_items;type:text" json:"action_items"`
	CreatedAt   int64  `gorm:"column:created_at;NOT NULL" json:"created_at"`
}

func (m *MeetingSummary) TableName() string {
	return "meeting_summaries"
}

// MeetingTranscript 会议转录
type MeetingTranscript struct {
	Id        int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	RoomId    string `gorm:"column:room_id;NOT NULL" json:"room_id"`
	UserId    string `gorm:"column:user_id;NOT NULL" json:"user_id"`
	UserName  string `gorm:"column:user_name;NOT NULL" json:"user_name"`
	Text      string `gorm:"column:text;type:text;NOT NULL" json:"text"`
	Language  string `gorm:"column:language;NOT NULL" json:"language"`
	Timestamp int64  `gorm:"column:timestamp;NOT NULL" json:"timestamp"`
}

func (m *MeetingTranscript) TableName() string {
	return "meeting_transcripts"
}

// AIChatMessage AI聊天消息
type AIChatMessage struct {
	Id        int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	RoomId    string `gorm:"column:room_id;NOT NULL" json:"room_id"`
	UserId    string `gorm:"column:user_id;NOT NULL" json:"user_id"`
	Message   string `gorm:"column:message;type:text;NOT NULL" json:"message"`
	Reply     string `gorm:"column:reply;type:text" json:"reply"`
	CreatedAt int64  `gorm:"column:created_at;NOT NULL" json:"created_at"`
}

func (m *AIChatMessage) TableName() string {
	return "ai_chat_messages"
}
