package builder

import (
	"context"

	"github.com/aide-family/moon/api"
	adminapi "github.com/aide-family/moon/api/admin"
	sbscriberapi "github.com/aide-family/moon/api/admin/subscriber"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/pkg/helper/middleware"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"
	"github.com/aide-family/moon/pkg/util/types"
	"github.com/aide-family/moon/pkg/vobj"
)

var _ ISubscriberModuleBuilder = (*subscriberModuleBuilder)(nil)

type (
	subscriberModuleBuilder struct {
		ctx context.Context
	}

	ISubscriberModuleBuilder interface {
		WithUserSubscriberListRequest(*sbscriberapi.UserSubscriberListRequest) IUserSubscriberListRequestBuilder
		WithSubscriberStrategyRequest(*sbscriberapi.SubscriberStrategyRequest) ISubscriberStrategyRequestBuilder
		WithUnSubscriberRequest(*sbscriberapi.UnSubscriberRequest) IUnSubscriberRequestBuilder
		WithStrategySubscriberRequest(*sbscriberapi.StrategySubscriberRequest) IStrategySubscriberRequestBuilder
		DoSubscriberBuilder() IDoSubscriberBuilder
	}

	IUserSubscriberListRequestBuilder interface {
		ToBo() *bo.QueryUserSubscriberParams
	}

	userSubscriberListRequestBuilder struct {
		ctx context.Context
		*sbscriberapi.UserSubscriberListRequest
	}

	ISubscriberStrategyRequestBuilder interface {
		ToBo() *bo.SubscriberStrategyParams
	}

	subscriberStrategyRequestBuilder struct {
		ctx context.Context
		*sbscriberapi.SubscriberStrategyRequest
	}

	IUnSubscriberRequestBuilder interface {
		ToBo() *bo.UnSubscriberStrategyParams
	}

	unSubscriberRequestBuilder struct {
		ctx context.Context
		*sbscriberapi.UnSubscriberRequest
	}

	IStrategySubscriberRequestBuilder interface {
		ToBo() *bo.QueryStrategySubscriberParams
	}

	strategySubscriberRequestBuilder struct {
		ctx context.Context
		*sbscriberapi.StrategySubscriberRequest
	}

	IDoSubscriberBuilder interface {
		ToAPI(*bizmodel.StrategySubscriber) *adminapi.StrategySubscriberItem
		ToAPIs([]*bizmodel.StrategySubscriber) []*adminapi.StrategySubscriberItem
		ToStrategies([]*bizmodel.StrategySubscriber) []*adminapi.StrategyItem
	}

	doSubscriberBuilder struct {
		ctx context.Context
	}
)

func (d *doSubscriberBuilder) ToAPI(subscriber *bizmodel.StrategySubscriber) *adminapi.StrategySubscriberItem {
	if types.IsNil(subscriber) || types.IsNil(d) {
		return nil
	}

	return &adminapi.StrategySubscriberItem{
		Id:         subscriber.ID,
		User:       nil, // TODO get user info
		NotifyType: api.NotifyType(subscriber.AlarmNoticeType),
	}
}

func (d *doSubscriberBuilder) ToAPIs(subscribers []*bizmodel.StrategySubscriber) []*adminapi.StrategySubscriberItem {
	if types.IsNil(subscribers) || types.IsNil(d) {
		return nil
	}

	return types.SliceTo(subscribers, func(subscriber *bizmodel.StrategySubscriber) *adminapi.StrategySubscriberItem {
		return d.ToAPI(subscriber)
	})
}

func (d *doSubscriberBuilder) ToStrategies(subscribers []*bizmodel.StrategySubscriber) []*adminapi.StrategyItem {
	if types.IsNil(subscribers) || types.IsNil(d) {
		return nil
	}

	return types.SliceToWithFilter(subscribers, func(subscriber *bizmodel.StrategySubscriber) (*adminapi.StrategyItem, bool) {
		if types.IsNil(subscriber.Strategy) {
			return nil, false
		}
		strategyInfo := subscriber.Strategy
		return NewParamsBuild(d.ctx).StrategyModuleBuilder().DoStrategyBuilder().ToAPI(strategyInfo), true
	})
}

func (s *strategySubscriberRequestBuilder) ToBo() *bo.QueryStrategySubscriberParams {
	if types.IsNil(s) || types.IsNil(s.StrategySubscriberRequest) {
		return nil
	}

	return &bo.QueryStrategySubscriberParams{
		Page:       types.NewPagination(s.GetPagination()),
		StrategyID: s.GetStrategyId(),
		NotifyType: vobj.NotifyType(s.GetNotifyType()),
	}
}

func (u *unSubscriberRequestBuilder) ToBo() *bo.UnSubscriberStrategyParams {
	if types.IsNil(u) || types.IsNil(u.UnSubscriberRequest) {
		return nil
	}

	return &bo.UnSubscriberStrategyParams{
		StrategyID: u.GetStrategyId(),
		UserID:     middleware.GetUserID(u.ctx),
	}
}

func (s *subscriberStrategyRequestBuilder) ToBo() *bo.SubscriberStrategyParams {
	if types.IsNil(s) || types.IsNil(s.SubscriberStrategyRequest) {
		return nil
	}

	return &bo.SubscriberStrategyParams{
		StrategyID: s.GetStrategyId(),
		NotifyType: vobj.NotifyType(s.GetNotifyType()),
		UserID:     middleware.GetUserID(s.ctx),
	}
}

func (u *userSubscriberListRequestBuilder) ToBo() *bo.QueryUserSubscriberParams {
	if types.IsNil(u) || types.IsNil(u.UserSubscriberListRequest) {
		return nil
	}

	return &bo.QueryUserSubscriberParams{
		UserID:     middleware.GetUserID(u.ctx),
		NotifyType: vobj.NotifyType(u.GetNotifyType()),
		Page:       types.NewPagination(u.GetPagination()),
	}
}

func (s *subscriberModuleBuilder) WithUserSubscriberListRequest(request *sbscriberapi.UserSubscriberListRequest) IUserSubscriberListRequestBuilder {
	return &userSubscriberListRequestBuilder{ctx: s.ctx, UserSubscriberListRequest: request}
}

func (s *subscriberModuleBuilder) WithSubscriberStrategyRequest(request *sbscriberapi.SubscriberStrategyRequest) ISubscriberStrategyRequestBuilder {
	return &subscriberStrategyRequestBuilder{ctx: s.ctx, SubscriberStrategyRequest: request}
}

func (s *subscriberModuleBuilder) WithUnSubscriberRequest(request *sbscriberapi.UnSubscriberRequest) IUnSubscriberRequestBuilder {
	return &unSubscriberRequestBuilder{ctx: s.ctx, UnSubscriberRequest: request}
}

func (s *subscriberModuleBuilder) WithStrategySubscriberRequest(request *sbscriberapi.StrategySubscriberRequest) IStrategySubscriberRequestBuilder {
	return &strategySubscriberRequestBuilder{ctx: s.ctx, StrategySubscriberRequest: request}
}

func (s *subscriberModuleBuilder) DoSubscriberBuilder() IDoSubscriberBuilder {
	return &doSubscriberBuilder{ctx: s.ctx}
}
