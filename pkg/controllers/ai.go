package controllers

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/mynaparrot/plugnmeet-protocol/utils"
	"github.com/mynaparrot/plugnmeet-server/pkg/config"
	"github.com/mynaparrot/plugnmeet-server/pkg/dbmodels"
	aiservice "github.com/mynaparrot/plugnmeet-server/pkg/services/ai"
	dbservice "github.com/mynaparrot/plugnmeet-server/pkg/services/db"
)

// ActionItem 行动项
type ActionItem struct {
	Description string `json:"description"`
	Assignee    string `json:"assignee"`
	DueDate     int64  `json:"due_date"`
}

// MeetingSummaryRes 会议摘要响应
type MeetingSummaryRes struct {
	Status      bool          `json:"status"`
	Msg         string        `json:"msg"`
	RoomId      string        `json:"room_id"`
	Summary     string        `json:"summary"`
	KeyPoints   []string      `json:"key_points"`
	ActionItems []*ActionItem `json:"action_items"`
	CreatedAt   int64         `json:"created_at"`
}

// AIChatReq AI聊天请求
type AIChatReq struct {
	Message string        `json:"message"`
	Context []ChatMessage `json:"context"`
}

// ChatMessage 聊天消息
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// AIChatRes AI聊天响应
type AIChatRes struct {
	Status    bool   `json:"status"`
	Msg       string `json:"msg"`
	RoomId    string `json:"room_id"`
	Reply     string `json:"reply"`
	CreatedAt int64  `json:"created_at"`
}

// AITranslationReq AI翻译请求
type AITranslationReq struct {
	RoomId     string `json:"room_id"`
	Text       string `json:"text"`
	SourceLang string `json:"source_lang"`
	TargetLang string `json:"target_lang"`
}

// AITranslationRes AI翻译响应
type AITranslationRes struct {
	Status         bool   `json:"status"`
	Msg            string `json:"msg"`
	RoomId         string `json:"room_id"`
	TranslatedText string `json:"translated_text"`
	SourceLang     string `json:"source_lang"`
	TargetLang     string `json:"target_lang"`
}

// TranscriptReq 转录请求
type TranscriptReq struct {
	RoomId   string `json:"room_id"`
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	Text     string `json:"text"`
	Language string `json:"language"`
}

// AIController AI控制器
type AIController struct {
	app       *config.AppConfig
	ds        *dbservice.DatabaseService
	aiService *aiservice.AIService
}

// NewAIController 创建新的AI控制器
func NewAIController(app *config.AppConfig, ds *dbservice.DatabaseService) *AIController {
	if app == nil {
		app = config.GetConfig()
	}
	if ds == nil {
		ds = dbservice.New(app.DB)
	}

	return &AIController{
		app:       app,
		ds:        ds,
		aiService: aiservice.NewAIService(app, ds),
	}
}

// HandleGenerateMeetingSummary 处理生成会议摘要请求
func (c *AIController) HandleGenerateMeetingSummary(ctx *fiber.Ctx) error {
	isAdmin := ctx.Locals("isAdmin")
	if isAdmin != true {
		return utils.SendCommonProtoJsonResponse(ctx, false, "only admin can perform this task")
	}

	roomId := ctx.Locals("roomId").(string)

	// 获取会议转录
	transcripts, err := c.aiService.GetTranscripts(roomId)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(ctx, false, err.Error())
	}

	// 构造AI摘要请求
	aiReq := &aiservice.AISummaryRequest{
		RoomId:      roomId,
		Transcripts: transcripts,
	}

	// 生成会议摘要
	result, err := c.aiService.GenerateMeetingSummary(context.Background(), aiReq)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(ctx, false, err.Error())
	}

	// 转换为协议消息
	var keyPoints []string
	for _, kp := range result.KeyPoints {
		keyPoints = append(keyPoints, kp)
	}

	var actionItems []*ActionItem
	for _, aiItem := range result.ActionItems {
		actionItem := &ActionItem{
			Description: aiItem.Description,
			Assignee:    aiItem.Assignee,
			DueDate:     aiItem.DueDate,
		}
		actionItems = append(actionItems, actionItem)
	}

	protoResult := &MeetingSummaryRes{
		Status:      true,
		Msg:         "success",
		RoomId:      result.RoomId,
		Summary:     result.Summary,
		KeyPoints:   keyPoints,
		ActionItems: actionItems,
		CreatedAt:   result.CreatedAt,
	}

	return ctx.JSON(protoResult)
}

// HandleGetMeetingSummary 处理获取会议摘要请求
func (c *AIController) HandleGetMeetingSummary(ctx *fiber.Ctx) error {
	roomId := ctx.Locals("roomId").(string)

	// 获取会议摘要
	result, err := c.aiService.GetMeetingSummary(roomId)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(ctx, false, err.Error())
	}

	// 转换为协议消息
	var keyPoints []string
	for _, kp := range result.KeyPoints {
		keyPoints = append(keyPoints, kp)
	}

	var actionItems []*ActionItem
	for _, aiItem := range result.ActionItems {
		actionItem := &ActionItem{
			Description: aiItem.Description,
			Assignee:    aiItem.Assignee,
			DueDate:     aiItem.DueDate,
		}
		actionItems = append(actionItems, actionItem)
	}

	protoResult := &MeetingSummaryRes{
		Status:      true,
		Msg:         "success",
		RoomId:      result.RoomId,
		Summary:     result.Summary,
		KeyPoints:   keyPoints,
		ActionItems: actionItems,
		CreatedAt:   result.CreatedAt,
	}

	return ctx.JSON(protoResult)
}

// HandleProcessChatMessage 处理AI聊天消息请求
func (c *AIController) HandleProcessChatMessage(ctx *fiber.Ctx) error {
	roomId := ctx.Locals("roomId").(string)
	userId := ctx.Locals("requestedUserId").(string)

	var req AIChatReq
	err := json.Unmarshal([]byte(ctx.FormValue("data")), &req)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(ctx, false, err.Error())
	}

	// 构造AI聊天请求
	aiReq := &aiservice.AIChatRequest{
		RoomId:  roomId,
		UserId:  userId,
		Message: req.Message,
		Context: make([]aiservice.ChatMessage, len(req.Context)),
	}

	// 转换上下文
	for i, msg := range req.Context {
		aiReq.Context[i] = aiservice.ChatMessage{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}

	// 处理聊天消息
	response, err := c.aiService.ProcessChatMessage(context.Background(), aiReq)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(ctx, false, err.Error())
	}

	// 保存聊天消息到数据库
	chatMsg := &dbmodels.AIChatMessage{
		RoomId:    roomId,
		UserId:    userId,
		Message:   req.Message,
		Reply:     response.Reply,
		CreatedAt: response.CreatedAt,
	}
	err = c.ds.SaveAIChatMessage(chatMsg)
	if err != nil {
		// 记录错误但不中断流程
		// logrus.Errorln("Error saving AI chat message:", err)
	}

	protoResponse := &AIChatRes{
		Status:    true,
		Msg:       "success",
		RoomId:    response.RoomId,
		Reply:     response.Reply,
		CreatedAt: response.CreatedAt,
	}

	return ctx.JSON(protoResponse)
}

// HandleTranslateText 处理文本翻译请求
func (c *AIController) HandleTranslateText(ctx *fiber.Ctx) error {
	var req AITranslationReq
	err := json.Unmarshal([]byte(ctx.FormValue("data")), &req)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(ctx, false, err.Error())
	}

	// 构造AI翻译请求
	aiReq := &aiservice.AITranslationRequest{
		RoomId:     req.RoomId,
		Text:       req.Text,
		SourceLang: req.SourceLang,
		TargetLang: req.TargetLang,
	}

	// 翻译文本
	translatedText, err := c.aiService.TranslateText(context.Background(), aiReq)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(ctx, false, err.Error())
	}

	protoResponse := &AITranslationRes{
		Status:         true,
		Msg:            "success",
		RoomId:         req.RoomId,
		TranslatedText: translatedText,
		SourceLang:     req.SourceLang,
		TargetLang:     req.TargetLang,
	}

	return ctx.JSON(protoResponse)
}

// HandleSaveTranscript 处理保存转录请求
func (c *AIController) HandleSaveTranscript(ctx *fiber.Ctx) error {
	var req TranscriptReq
	err := json.Unmarshal([]byte(ctx.FormValue("data")), &req)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(ctx, false, err.Error())
	}

	// 保存转录
	err = c.aiService.SaveTranscript(req.RoomId, req.UserId, req.UserName, req.Text, req.Language)
	if err != nil {
		return utils.SendCommonProtoJsonResponse(ctx, false, err.Error())
	}

	return utils.SendCommonProtoJsonResponse(ctx, true, "success")
}
