package interceptor

import (
	"context"
	"encoding/json"
	"github.com/WildWeber-xiaocheng/ms_project/project-common/encrypts"
	"github.com/WildWeber-xiaocheng/ms_project/project-grpc/project"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/dao"
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/repo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

// CacheInterceptor 除了缓存拦截器，还可以实现日志拦截器
type CacheInterceptor struct {
	cache    repo.Cache
	cacheMap map[string]interface{}
}

// 这里没有用到这个数据结构，可以利用这个来设置参数，而不是在Cache函数里将参数设置死
type CacheRespOption struct {
	path   string
	typ    interface{}
	expire time.Duration
}

func New() *CacheInterceptor {
	cacheMap := make(map[string]interface{})
	cacheMap["/project.service.v1.ProjectService/FindProjectByMemId"] = &project.MyProjectResponse{}
	return &CacheInterceptor{cache: dao.Rc, cacheMap: cacheMap}
}

func (c *CacheInterceptor) Cache() grpc.ServerOption {
	return grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		respType := c.cacheMap[info.FullMethod]
		if respType == nil { //不在拦截范围
			return handler(ctx, req)
		}
		//先查询是否有缓存，有的话直接返回，没有的话先请求，然后存入缓存
		con, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		marshal, _ := json.Marshal(req)
		cacheKey := encrypts.Md5(string(marshal))
		respJson, _ := c.cache.Get(con, info.FullMethod+"::"+cacheKey)
		//有缓存
		if respJson != "" { //或者err == nil
			json.Unmarshal([]byte(respJson), &respType)
			zap.L().Info(info.FullMethod + "走了缓存")
			return respType, nil
		}
		//没有缓存
		resp, err = handler(ctx, req)
		bytes, _ := json.Marshal(resp)
		c.cache.Put(con, info.FullMethod+"::"+cacheKey, string(bytes), 5*time.Minute)
		zap.L().Info(info.FullMethod + "放入缓存")
		return
	})
}
