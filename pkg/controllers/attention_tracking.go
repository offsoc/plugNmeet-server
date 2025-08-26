package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/mynaparrot/plugnmeet-protocol/utils"
	"github.com/mynaparrot/plugnmeet-server/pkg/config"
	"github.com/mynaparrot/plugnmeet-server/pkg/models"
	dbservice "github.com/mynaparrot/plugnmeet-server/pkg/services/db"
)

// TrackAttentionReq 记录注意力请求
type TrackAttentionReq struct {
	RoomId   string `json:"room_id"`
	UserId   string `json:"user_id"`
	IsActive bool   `json:"is_active"`
}

// GetUserAttentionStatsReq 获取用户注意力统计请求
type GetUserAttentionStatsReq struct {
	RoomId string `json:"room_id"`
	UserId string `json:"user_id"`
}

// GetRoomAttentionStatsReq 获取房间注意力统计请求
type GetRoomAttentionStatsReq struct {
	RoomId string `json:"room_id"`
}

// AttentionTrackingRes 注意力跟踪响应
type AttentionTrackingRes struct {
	Status    bool                         `json:"status"`
	Msg       string                       `json:"msg"`
	Stats     *models.UserAttentionStats   `json:"stats,omitempty"`
	StatsList []*models.UserAttentionStats `json:"stats_list,omitempty"`
}

type AttentionTrackingController struct {
	app            *config.AppConfig
	ds             *dbservice.DatabaseService
	attentionModel *models.AttentionTrackingModel
}

func NewAttentionTrackingController() *AttentionTrackingController {
	app := config.GetConfig()
	ds := dbservice.New(app.DB)

	return &AttentionTrackingController{
		app:            app,
		ds:             ds,
		attentionModel: models.NewAttentionTrackingModel(app, ds),
	}
}

// TrackAttention 记录用户的注意力状态
func (c *AttentionTrackingController) TrackAttention(ctx *fiber.Ctx) error {
	var req TrackAttentionReq
	err := json.Unmarshal([]byte(ctx.FormValue("data")), &req)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(ctx, false, err.Error())
	}

	// 验证请求参数
	if req.RoomId == "" || req.UserId == "" {
		return utils.SendCommonProtoJsonResponse(ctx, false, "room_id and user_id are required")
	}

	// 记录用户的注意力状态
	err = c.attentionModel.TrackUserAttention(req.RoomId, req.UserId, req.IsActive)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(ctx, false, err.Error())
	}

	res := &AttentionTrackingRes{
		Status: true,
		Msg:    "attention tracked successfully",
	}

	return ctx.JSON(res)
}

// GetUserAttentionStats 获取用户的注意力统计信息
func (c *AttentionTrackingController) GetUserAttentionStats(ctx *fiber.Ctx) error {
	var req GetUserAttentionStatsReq
	err := json.Unmarshal([]byte(ctx.FormValue("data")), &req)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(ctx, false, err.Error())
	}

	// 验证请求参数
	if req.RoomId == "" || req.UserId == "" {
		return utils.SendCommonProtoJsonResponse(ctx, false, "room_id and user_id are required")
	}

	// 获取用户的注意力统计信息
	stats, err := c.attentionModel.GetUserAttentionStats(req.RoomId, req.UserId)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(ctx, false, err.Error())
	}

	res := &AttentionTrackingRes{
		Status: true,
		Msg:    "stats retrieved successfully",
		Stats:  stats,
	}

	return ctx.JSON(res)
}

// GetRoomAttentionStats 获取房间内所有用户的注意力统计信息
func (c *AttentionTrackingController) GetRoomAttentionStats(ctx *fiber.Ctx) error {
	var req GetRoomAttentionStatsReq
	err := json.Unmarshal([]byte(ctx.FormValue("data")), &req)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(ctx, false, err.Error())
	}

	// 验证请求参数
	if req.RoomId == "" {
		return utils.SendCommonProtoJsonResponse(ctx, false, "room_id is required")
	}

	// 获取房间内所有用户的注意力统计信息
	stats, err := c.attentionModel.GetRoomAttentionStats(req.RoomId)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(ctx, false, err.Error())
	}

	res := &AttentionTrackingRes{
		Status:    true,
		Msg:       "stats retrieved successfully",
		StatsList: stats,
	}

	return ctx.JSON(res)
}
