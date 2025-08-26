package dbservice

import (
	"errors"

	"github.com/mynaparrot/plugnmeet-server/pkg/dbmodels"
	"gorm.io/gorm"
)

func (s *DatabaseService) GetAnalytics(roomIds []string, offset, limit uint64, direction *string) ([]dbmodels.Analytics, int64, error) {
	var analytics []dbmodels.Analytics
	var total int64

	d := s.db.Model(&dbmodels.Analytics{})
	if len(roomIds) > 0 {
		d.Where("room_id IN ?", roomIds)
	}

	if err := d.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if limit == 0 {
		limit = 20
	}

	orderBy := "DESC"
	if direction != nil && *direction == "ASC" {
		orderBy = "ASC"
	}

	result := d.Offset(int(offset)).Limit(int(limit)).Order("id " + orderBy).Find(&analytics)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, 0, result.Error
	}

	return analytics, total, nil
}

func (s *DatabaseService) GetAnalyticByFileId(fileId string) (*dbmodels.Analytics, error) {
	info := new(dbmodels.Analytics)
	cond := &dbmodels.Analytics{
		FileID: fileId,
	}

	result := s.db.Where(cond).Take(info)
	switch {
	case errors.Is(result.Error, gorm.ErrRecordNotFound):
		return nil, nil
	case result.Error != nil:
		return nil, result.Error
	}

	return info, nil
}

func (s *DatabaseService) GetAnalyticByRoomTableId(roomTableId uint64) (*dbmodels.Analytics, error) {
	info := new(dbmodels.Analytics)
	cond := &dbmodels.Analytics{
		RoomTableID: roomTableId,
	}

	result := s.db.Where(cond).Take(info)
	switch {
	case errors.Is(result.Error, gorm.ErrRecordNotFound):
		return nil, nil
	case result.Error != nil:
		return nil, result.Error
	}

	return info, nil
}

// GetUserAttentionStats 获取用户的注意力统计信息
func (d *DatabaseService) GetUserAttentionStats(roomId, userId string) (*dbmodels.UserAttentionStats, error) {
	// 获取用户的所有注意力记录
	var records []dbmodels.AttentionTracking
	err := d.db.Where("room_id = ? AND user_id = ?", roomId, userId).Order("timestamp ASC").Find(&records).Error
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return &dbmodels.UserAttentionStats{
			RoomId:           roomId,
			UserId:           userId,
			TotalActiveTime:  0,
			TotalPassiveTime: 0,
			LastActiveTime:   0,
			AttentionRate:    0,
		}, nil
	}

	// 计算统计数据
	var totalActiveTime, totalPassiveTime, lastActiveTime int64
	for i, record := range records {
		if record.IsActive {
			lastActiveTime = record.Timestamp
			if i > 0 {
				// 计算与前一个记录的时间差
				timeDiff := record.Timestamp - records[i-1].Timestamp
				if records[i-1].IsActive {
					totalActiveTime += timeDiff
				} else {
					totalPassiveTime += timeDiff
				}
			}
		}
	}

	// 计算注意力比率
	var attentionRate float64
	totalTime := totalActiveTime + totalPassiveTime
	if totalTime > 0 {
		attentionRate = float64(totalActiveTime) / float64(totalTime)
	}

	stats := &dbmodels.UserAttentionStats{
		RoomId:           roomId,
		UserId:           userId,
		TotalActiveTime:  totalActiveTime,
		TotalPassiveTime: totalPassiveTime,
		LastActiveTime:   lastActiveTime,
		AttentionRate:    attentionRate,
	}

	return stats, nil
}

// GetRoomAttentionStats 获取房间内所有用户的注意力统计信息
func (d *DatabaseService) GetRoomAttentionStats(roomId string) ([]*dbmodels.UserAttentionStats, error) {
	// 获取房间内所有用户的ID
	var userIds []string
	err := d.db.Model(&dbmodels.AttentionTracking{}).
		Where("room_id = ?", roomId).
		Distinct("user_id").
		Pluck("user_id", &userIds).Error
	if err != nil {
		return nil, err
	}

	// 为每个用户获取注意力统计数据
	var stats []*dbmodels.UserAttentionStats
	for _, userId := range userIds {
		userStats, err := d.GetUserAttentionStats(roomId, userId)
		if err != nil {
			continue
		}
		stats = append(stats, userStats)
	}

	return stats, nil
}

// GetMeetingSummary 获取会议摘要
func (d *DatabaseService) GetMeetingSummary(roomId string) (*dbmodels.MeetingSummary, error) {
	var summary dbmodels.MeetingSummary
	err := d.db.Where("room_id = ?", roomId).First(&summary).Error
	if err != nil {
		return nil, err
	}
	return &summary, nil
}

// GetMeetingTranscripts 获取会议转录
func (d *DatabaseService) GetMeetingTranscripts(roomId string) ([]*dbmodels.MeetingTranscript, error) {
	var transcripts []*dbmodels.MeetingTranscript
	err := d.db.Where("room_id = ?", roomId).Order("timestamp ASC").Find(&transcripts).Error
	if err != nil {
		return nil, err
	}
	return transcripts, nil
}

// GetAIChatMessages 获取AI聊天消息
func (d *DatabaseService) GetAIChatMessages(roomId string) ([]*dbmodels.AIChatMessage, error) {
	var messages []*dbmodels.AIChatMessage
	err := d.db.Where("room_id = ?", roomId).Order("created_at ASC").Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
