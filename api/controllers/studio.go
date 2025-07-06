package controllers

import (
	"GeneReport_platform/api/dto"
	"GeneReport_platform/configs"
	"GeneReport_platform/internal/services"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"time"
)

type Studio struct {
}

var (
	StudioController = &Studio{}
)

// CreateFromThirdParty 从第三方创建
func (s *Studio) CreateFromThirdParty(ctx *gin.Context) {
	var req dto.CreateAllFromThirdPartyReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrResponse{
			Code:    http.StatusBadRequest,
			Message: "请求体格式错误,请重试",
		})
		return
	}

	//fixme sse分阶段响应
	if toResp, err := services.StudioServ.CreateAllFromThirdPartyOnChain(ctx.GetString("user_address"), req); err != nil {
		_ = ctx.Error(err)
		return
	} else {
		ctx.JSON(http.StatusOK, dto.Response{
			Code:    http.StatusOK,
			Message: "创建成功",
			Data:    toResp,
		})
	}

}

// GetCompletedProfileIds 获取后台中已经保存数据的该用户的profile id供用户选择
func (s *Studio) GetCompletedProfileIds(ctx *gin.Context) {
	userAddress := ctx.GetString("user_address")
	ids, err := services.StudioServ.GetProfileIdsByUser(userAddress)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Message: "获取成功",
		Data:    ids,
	})
}

// GetUncompletedProfileIDProgress 实时获取wegene请求任务进度
func (s *Studio) GetUncompletedProfileIDProgress(ctx *gin.Context) {
	ctxRedis := context.Background()
	var TaskCount int
	//获取该用户提交的正在处理中的数据请求任务
	userAddress := ctx.GetString("user_address")
	ids, err := services.StudioServ.GetUncompletedProfileIdsByUser(userAddress)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	TaskCount = len(ids.ProfileIds)
	// 检查任务数是否为 0
	if len(ids.ProfileIds) == 0 {
		ctx.Writer.Header().Set("Content-Type", "text/event-stream")
		ctx.Writer.Header().Set("Cache-Control", "no-cache")
		ctx.Writer.Header().Set("Connection", "keep-alive")
		_, err := fmt.Fprintf(ctx.Writer, "data: {\"info\": \"无正在处理中的任务\"}\n\n")
		if err != nil {
			log.Println("SSE 写入失败:", err)
		}
		ctx.Writer.Flush()
		return // 直接返回，关闭 SSE 连接
	}

	//创建redis订阅频道切片
	channels := make([]string, 0, TaskCount)
	//循环加入前缀
	for _, id := range ids.ProfileIds {
		channels = append(channels, "process:"+id)
	}

	//订阅redis消息
	pubSub := configs.RedisClient.Subscribe(ctxRedis, channels...)
	defer pubSub.Close()

	// 设置 SSE 响应头
	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Connection", "keep-alive")

	// 心跳机制
	go func() {
		for {
			time.Sleep(15 * time.Second)
			_, _ = fmt.Fprintf(ctx.Writer, "data: {\"keepalive\": true}\n\n")
			ctx.Writer.Flush()
		}
	}()

	// 获取初始进度
	for _, id := range ids.ProfileIds {
		progress, err := configs.RedisClient.Get(ctxRedis, "task:"+id).Result()
		if errors.Is(err, redis.Nil) {
			_, _ = fmt.Fprintf(ctx.Writer, "data: {\"taskID\": \"%s\", \"error\": \"任务不存在\"}\n\n", id)
			ctx.Writer.Flush()
			continue
		}
		if err != nil {
			_, _ = fmt.Fprintf(ctx.Writer, "data: {\"taskID\": \"%s\", \"error\": \"Redis 查询失败\"}\n\n", id)
			ctx.Writer.Flush()
			continue
		}
		//推送初始进度
		_, _ = fmt.Fprintf(ctx.Writer, "data: {\"taskID\": \"%s\", \"progress\": \"%s\"}\n\n", id, progress)
		ctx.Writer.Flush()
	}

	// 监听 Redis 频道消息
	ch := pubSub.Channel()
	for msg := range ch {
		profileID := msg.Channel[len("progress:"):]
		_, _ = fmt.Fprintf(ctx.Writer, "data: {\"taskID\": \"%s\", \"progress\": \"%s\"}\n\n", profileID, msg.Payload)
		ctx.Writer.Flush()
		// 如果任务完成，从活跃任务列表检查是否继续
		if msg.Payload == "completed" {
			err := pubSub.Unsubscribe(ctxRedis, "process:"+profileID)
			if err != nil {
				log.Println("取消订阅失败:", err)
			}

			if TaskCount--; TaskCount == 0 {
				break
			}
		}
	}
}
