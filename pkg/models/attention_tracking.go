package models

import (
	"time"

	"github.com/mynaparrot/plugnmeet-server/pkg/config"
	"github.com/mynaparrot/plugnmeet-server/pkg/dbmodels"
	dbservice "github.com/mynaparrot/plugnmeet-server/pkg/services/db"
	log "github.com/sirupsen/logrus"
)

// UserAttentionStats 用户注意力统计信息
type UserAttentionStats struct {
	RoomId           string  `json:"room_id"`
	UserId           string  `json:"user_id"`
	TotalActiveTime  int64   `json:"total_active_time"`
	TotalPassiveTime int64   `json:"total_passive_time"`
	LastActiveTime   int64   `json:"last_active_time"`
	AttentionRate    float64 `json:"attention_rate"`
}

type AttentionTrackingModel struct {
	app *config.AppConfig
	ds  *dbservice.DatabaseService
}

func NewAttentionTrackingModel(app *config.AppConfig, ds *dbservice.DatabaseService) *AttentionTrackingModel {
	if app == nil {
		app = config.GetConfig()
	}
	if ds == nil {
		ds = dbservice.New(app.DB)
	}

	return &AttentionTrackingModel{
		app: app,
		ds:  ds,
	}
}

// TrackUserAttention 记录用户的注意力状态
func (m *AttentionTrackingModel) TrackUserAttention(roomId, userId string, isActive bool) error {
	// 创建或更新用户的注意力记录
	attention := &dbmodels.AttentionTracking{
		RoomId:    roomId,
		UserId:    userId,
		IsActive:  isActive,
		Timestamp: time.Now().Unix(),
	}

	// 插入或更新数据库记录
	err := m.ds.InsertOrUpdateAttentionTracking(attention)
	if err != nil {
		log.Errorln("Error tracking user attention:", err)
		return err
	}

	return nil
}

// GetUserAttentionStats 获取用户的注意力统计信息
func (m *AttentionTrackingModel) GetUserAttentionStats(roomId, userId string) (*UserAttentionStats, error) {
	stats, err := m.ds.GetUserAttentionStats(roomId, userId)
	if err != nil {
		log.Errorln("Error getting user attention stats:", err)
		return nil, err
	}

	// 转换为协议消息
	protoStats := &UserAttentionStats{
		RoomId:           stats.RoomId,
		UserId:           stats.UserId,
		TotalActiveTime:  stats.TotalActiveTime,
		TotalPassiveTime: stats.TotalPassiveTime,
		LastActiveTime:   stats.LastActiveTime,
		AttentionRate:    stats.AttentionRate,
	}

	return protoStats, nil
}

// GetRoomAttentionStats 获取房间内所有用户的注意力统计信息
func (m *AttentionTrackingModel) GetRoomAttentionStats(roomId string) ([]*UserAttentionStats, error) {
	stats, err := m.ds.GetRoomAttentionStats(roomId)
	if err != nil {
		log.Errorln("Error getting room attention stats:", err)
		return nil, err
	}

	// 转换为协议消息数组
	var protoStats []*UserAttentionStats
	for _, stat := range stats {
		protoStat := &UserAttentionStats{
			RoomId:           stat.RoomId,
			UserId:           stat.UserId,
			TotalActiveTime:  stat.TotalActiveTime,
			TotalPassiveTime: stat.TotalPassiveTime,
			LastActiveTime:   stat.LastActiveTime,
			AttentionRate:    stat.AttentionRate,
		}
		protoStats = append(protoStats, protoStat)
	}

	return protoStats, nil
}

// CleanupOldAttentionRecords 清理旧的注意力记录
func (m *AttentionTrackingModel) CleanupOldAttentionRecords(beforeTime int64) error {
	err := m.ds.CleanupOldAttentionRecords(beforeTime)
	if err != nil {
		log.Errorln("Error cleaning up old attention records:", err)
		return err
	}

	return nil
}
