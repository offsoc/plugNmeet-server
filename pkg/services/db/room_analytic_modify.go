package dbservice

import (
	"errors"

	"github.com/mynaparrot/plugnmeet-server/pkg/dbmodels"
	"gorm.io/gorm"
)

func (s *DatabaseService) InsertAnalyticsData(info *dbmodels.Analytics) (int64, error) {
	result := s.db.Create(info)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func (s *DatabaseService) DeleteAnalyticByFileId(fileId string) (int64, error) {
	cond := &dbmodels.Analytics{
		FileID: fileId,
	}

	result := s.db.Where(cond).Delete(&dbmodels.Analytics{})
	switch {
	case errors.Is(result.Error, gorm.ErrRecordNotFound):
		return 0, nil
	case result.Error != nil:
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

// SaveMeetingSummary 保存会议摘要
func (d *DatabaseService) SaveMeetingSummary(summary *dbmodels.MeetingSummary) error {
	return d.db.Save(summary).Error
}

// SaveMeetingTranscript 保存会议转录
func (d *DatabaseService) SaveMeetingTranscript(transcript *dbmodels.MeetingTranscript) error {
	return d.db.Save(transcript).Error
}

// SaveAIChatMessage 保存AI聊天消息
func (d *DatabaseService) SaveAIChatMessage(message *dbmodels.AIChatMessage) error {
	return d.db.Save(message).Error
}

// InsertOrUpdateAttentionTracking 插入或更新用户的注意力跟踪记录
func (d *DatabaseService) InsertOrUpdateAttentionTracking(attention *dbmodels.AttentionTracking) error {
	return d.db.Save(attention).Error
}

// CleanupOldAttentionRecords 清理旧的注意力记录
func (d *DatabaseService) CleanupOldAttentionRecords(beforeTime int64) error {
	return d.db.Where("timestamp < ?", beforeTime).Delete(&dbmodels.AttentionTracking{}).Error
}
