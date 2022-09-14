// Code generated by protoc-gen-cc-gateway. DO NOT EDIT.
// source: token/service/config/config.proto

/*
Package config contains
  *   chaincode methods names {service_name}Chaincode_{method_name}
  *   chaincode interface definition {service_name}Chaincode
  *   chaincode gateway definition {service_name}}Gateway
  *   chaincode service to cckit router registration func
*/
package config

import (
	context "context"
	_ "embed"
	errors "errors"

	cckit_gateway "github.com/hyperledger-labs/cckit/gateway"
	cckit_router "github.com/hyperledger-labs/cckit/router"
	cckit_defparam "github.com/hyperledger-labs/cckit/router/param/defparam"
	cckit_sdk "github.com/hyperledger-labs/cckit/sdk"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ConfigServiceChaincode method names
const (

	// ConfigServiceChaincodeMethodPrefix allows to use multiple services with same method names in one chaincode
	ConfigServiceChaincodeMethodPrefix = ""

	ConfigServiceChaincode_GetConfig = ConfigServiceChaincodeMethodPrefix + "GetConfig"

	ConfigServiceChaincode_SetConfig = ConfigServiceChaincodeMethodPrefix + "SetConfig"

	ConfigServiceChaincode_GetToken = ConfigServiceChaincodeMethodPrefix + "GetToken"

	ConfigServiceChaincode_GetDefaultToken = ConfigServiceChaincodeMethodPrefix + "GetDefaultToken"

	ConfigServiceChaincode_CreateTokenType = ConfigServiceChaincodeMethodPrefix + "CreateTokenType"

	ConfigServiceChaincode_GetTokenType = ConfigServiceChaincodeMethodPrefix + "GetTokenType"

	ConfigServiceChaincode_ListTokenTypes = ConfigServiceChaincodeMethodPrefix + "ListTokenTypes"

	ConfigServiceChaincode_UpdateTokenType = ConfigServiceChaincodeMethodPrefix + "UpdateTokenType"

	ConfigServiceChaincode_DeleteTokenType = ConfigServiceChaincodeMethodPrefix + "DeleteTokenType"

	ConfigServiceChaincode_GetTokenGroups = ConfigServiceChaincodeMethodPrefix + "GetTokenGroups"

	ConfigServiceChaincode_CreateTokenGroup = ConfigServiceChaincodeMethodPrefix + "CreateTokenGroup"

	ConfigServiceChaincode_GetTokenGroup = ConfigServiceChaincodeMethodPrefix + "GetTokenGroup"

	ConfigServiceChaincode_DeleteTokenGroup = ConfigServiceChaincodeMethodPrefix + "DeleteTokenGroup"
)

// ConfigServiceChaincode chaincode methods interface
type ConfigServiceChaincode interface {
	GetConfig(cckit_router.Context, *emptypb.Empty) (*Config, error)

	SetConfig(cckit_router.Context, *Config) (*Config, error)

	GetToken(cckit_router.Context, *TokenId) (*Token, error)

	GetDefaultToken(cckit_router.Context, *emptypb.Empty) (*Token, error)

	CreateTokenType(cckit_router.Context, *CreateTokenTypeRequest) (*TokenType, error)

	GetTokenType(cckit_router.Context, *TokenTypeId) (*TokenType, error)

	ListTokenTypes(cckit_router.Context, *emptypb.Empty) (*TokenTypes, error)

	UpdateTokenType(cckit_router.Context, *UpdateTokenTypeRequest) (*TokenType, error)

	DeleteTokenType(cckit_router.Context, *TokenTypeId) (*TokenType, error)

	GetTokenGroups(cckit_router.Context, *TokenTypeId) (*TokenGroups, error)

	CreateTokenGroup(cckit_router.Context, *CreateTokenGroupRequest) (*TokenGroup, error)

	GetTokenGroup(cckit_router.Context, *TokenGroupId) (*TokenGroup, error)

	DeleteTokenGroup(cckit_router.Context, *TokenGroupId) (*Token, error)
}

// RegisterConfigServiceChaincode registers service methods as chaincode router handlers
func RegisterConfigServiceChaincode(r *cckit_router.Group, cc ConfigServiceChaincode) error {

	r.Query(ConfigServiceChaincode_GetConfig,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.GetConfig(ctx, ctx.Param().(*emptypb.Empty))
		},
		cckit_defparam.Proto(&emptypb.Empty{}))

	r.Invoke(ConfigServiceChaincode_SetConfig,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.SetConfig(ctx, ctx.Param().(*Config))
		},
		cckit_defparam.Proto(&Config{}))

	r.Query(ConfigServiceChaincode_GetToken,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.GetToken(ctx, ctx.Param().(*TokenId))
		},
		cckit_defparam.Proto(&TokenId{}))

	r.Query(ConfigServiceChaincode_GetDefaultToken,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.GetDefaultToken(ctx, ctx.Param().(*emptypb.Empty))
		},
		cckit_defparam.Proto(&emptypb.Empty{}))

	r.Invoke(ConfigServiceChaincode_CreateTokenType,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.CreateTokenType(ctx, ctx.Param().(*CreateTokenTypeRequest))
		},
		cckit_defparam.Proto(&CreateTokenTypeRequest{}))

	r.Query(ConfigServiceChaincode_GetTokenType,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.GetTokenType(ctx, ctx.Param().(*TokenTypeId))
		},
		cckit_defparam.Proto(&TokenTypeId{}))

	r.Query(ConfigServiceChaincode_ListTokenTypes,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.ListTokenTypes(ctx, ctx.Param().(*emptypb.Empty))
		},
		cckit_defparam.Proto(&emptypb.Empty{}))

	r.Invoke(ConfigServiceChaincode_UpdateTokenType,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.UpdateTokenType(ctx, ctx.Param().(*UpdateTokenTypeRequest))
		},
		cckit_defparam.Proto(&UpdateTokenTypeRequest{}))

	r.Invoke(ConfigServiceChaincode_DeleteTokenType,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.DeleteTokenType(ctx, ctx.Param().(*TokenTypeId))
		},
		cckit_defparam.Proto(&TokenTypeId{}))

	r.Query(ConfigServiceChaincode_GetTokenGroups,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.GetTokenGroups(ctx, ctx.Param().(*TokenTypeId))
		},
		cckit_defparam.Proto(&TokenTypeId{}))

	r.Invoke(ConfigServiceChaincode_CreateTokenGroup,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.CreateTokenGroup(ctx, ctx.Param().(*CreateTokenGroupRequest))
		},
		cckit_defparam.Proto(&CreateTokenGroupRequest{}))

	r.Query(ConfigServiceChaincode_GetTokenGroup,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.GetTokenGroup(ctx, ctx.Param().(*TokenGroupId))
		},
		cckit_defparam.Proto(&TokenGroupId{}))

	r.Invoke(ConfigServiceChaincode_DeleteTokenGroup,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.DeleteTokenGroup(ctx, ctx.Param().(*TokenGroupId))
		},
		cckit_defparam.Proto(&TokenGroupId{}))

	return nil
}

//go:embed config.swagger.json
var ConfigServiceSwagger []byte

// NewConfigServiceGateway creates gateway to access chaincode method via chaincode service
func NewConfigServiceGateway(sdk cckit_sdk.SDK, channel, chaincode string, opts ...cckit_gateway.Opt) *ConfigServiceGateway {
	return NewConfigServiceGatewayFromInstance(
		cckit_gateway.NewChaincodeInstanceService(
			sdk,
			&cckit_gateway.ChaincodeLocator{Channel: channel, Chaincode: chaincode},
			opts...,
		))
}

func NewConfigServiceGatewayFromInstance(chaincodeInstance cckit_gateway.ChaincodeInstance) *ConfigServiceGateway {
	return &ConfigServiceGateway{
		ChaincodeInstance: chaincodeInstance,
	}
}

// gateway implementation
// gateway can be used as kind of SDK, GRPC or REST server ( via grpc-gateway or clay )
type ConfigServiceGateway struct {
	ChaincodeInstance cckit_gateway.ChaincodeInstance
}

func (c *ConfigServiceGateway) Invoker() cckit_gateway.ChaincodeInstanceInvoker {
	return cckit_gateway.NewChaincodeInstanceServiceInvoker(c.ChaincodeInstance)
}

// ServiceDef returns service definition
func (c *ConfigServiceGateway) ServiceDef() cckit_gateway.ServiceDef {
	return cckit_gateway.NewServiceDef(
		_ConfigService_serviceDesc.ServiceName,
		ConfigServiceSwagger,
		&_ConfigService_serviceDesc,
		c,
		RegisterConfigServiceHandlerFromEndpoint,
	)
}

func (c *ConfigServiceGateway) GetConfig(ctx context.Context, in *emptypb.Empty) (*Config, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker().Query(ctx, ConfigServiceChaincode_GetConfig, []interface{}{in}, &Config{}); err != nil {
		return nil, err
	} else {
		return res.(*Config), nil
	}
}

func (c *ConfigServiceGateway) SetConfig(ctx context.Context, in *Config) (*Config, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker().Invoke(ctx, ConfigServiceChaincode_SetConfig, []interface{}{in}, &Config{}); err != nil {
		return nil, err
	} else {
		return res.(*Config), nil
	}
}

func (c *ConfigServiceGateway) GetToken(ctx context.Context, in *TokenId) (*Token, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker().Query(ctx, ConfigServiceChaincode_GetToken, []interface{}{in}, &Token{}); err != nil {
		return nil, err
	} else {
		return res.(*Token), nil
	}
}

func (c *ConfigServiceGateway) GetDefaultToken(ctx context.Context, in *emptypb.Empty) (*Token, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker().Query(ctx, ConfigServiceChaincode_GetDefaultToken, []interface{}{in}, &Token{}); err != nil {
		return nil, err
	} else {
		return res.(*Token), nil
	}
}

func (c *ConfigServiceGateway) CreateTokenType(ctx context.Context, in *CreateTokenTypeRequest) (*TokenType, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker().Invoke(ctx, ConfigServiceChaincode_CreateTokenType, []interface{}{in}, &TokenType{}); err != nil {
		return nil, err
	} else {
		return res.(*TokenType), nil
	}
}

func (c *ConfigServiceGateway) GetTokenType(ctx context.Context, in *TokenTypeId) (*TokenType, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker().Query(ctx, ConfigServiceChaincode_GetTokenType, []interface{}{in}, &TokenType{}); err != nil {
		return nil, err
	} else {
		return res.(*TokenType), nil
	}
}

func (c *ConfigServiceGateway) ListTokenTypes(ctx context.Context, in *emptypb.Empty) (*TokenTypes, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker().Query(ctx, ConfigServiceChaincode_ListTokenTypes, []interface{}{in}, &TokenTypes{}); err != nil {
		return nil, err
	} else {
		return res.(*TokenTypes), nil
	}
}

func (c *ConfigServiceGateway) UpdateTokenType(ctx context.Context, in *UpdateTokenTypeRequest) (*TokenType, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker().Invoke(ctx, ConfigServiceChaincode_UpdateTokenType, []interface{}{in}, &TokenType{}); err != nil {
		return nil, err
	} else {
		return res.(*TokenType), nil
	}
}

func (c *ConfigServiceGateway) DeleteTokenType(ctx context.Context, in *TokenTypeId) (*TokenType, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker().Invoke(ctx, ConfigServiceChaincode_DeleteTokenType, []interface{}{in}, &TokenType{}); err != nil {
		return nil, err
	} else {
		return res.(*TokenType), nil
	}
}

func (c *ConfigServiceGateway) GetTokenGroups(ctx context.Context, in *TokenTypeId) (*TokenGroups, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker().Query(ctx, ConfigServiceChaincode_GetTokenGroups, []interface{}{in}, &TokenGroups{}); err != nil {
		return nil, err
	} else {
		return res.(*TokenGroups), nil
	}
}

func (c *ConfigServiceGateway) CreateTokenGroup(ctx context.Context, in *CreateTokenGroupRequest) (*TokenGroup, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker().Invoke(ctx, ConfigServiceChaincode_CreateTokenGroup, []interface{}{in}, &TokenGroup{}); err != nil {
		return nil, err
	} else {
		return res.(*TokenGroup), nil
	}
}

func (c *ConfigServiceGateway) GetTokenGroup(ctx context.Context, in *TokenGroupId) (*TokenGroup, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker().Query(ctx, ConfigServiceChaincode_GetTokenGroup, []interface{}{in}, &TokenGroup{}); err != nil {
		return nil, err
	} else {
		return res.(*TokenGroup), nil
	}
}

func (c *ConfigServiceGateway) DeleteTokenGroup(ctx context.Context, in *TokenGroupId) (*Token, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker().Invoke(ctx, ConfigServiceChaincode_DeleteTokenGroup, []interface{}{in}, &Token{}); err != nil {
		return nil, err
	} else {
		return res.(*Token), nil
	}
}

// ConfigServiceChaincodeResolver interface for service resolver
type (
	ConfigServiceChaincodeResolver interface {
		Resolve(ctx cckit_router.Context) (ConfigServiceChaincode, error)
	}

	ConfigServiceChaincodeLocalResolver struct {
		service ConfigServiceChaincode
	}

	ConfigServiceChaincodeLocatorResolver struct {
		locatorResolver cckit_gateway.ChaincodeLocatorResolver
		service         ConfigServiceChaincode
	}
)

func NewConfigServiceChaincodeLocalResolver(service ConfigServiceChaincode) *ConfigServiceChaincodeLocalResolver {
	return &ConfigServiceChaincodeLocalResolver{
		service: service,
	}
}

func (r *ConfigServiceChaincodeLocalResolver) Resolve(ctx cckit_router.Context) (ConfigServiceChaincode, error) {
	if r.service == nil {
		return nil, errors.New("service not set for local chaincode resolver")
	}

	return r.service, nil
}

func NewConfigServiceChaincodeResolver(locatorResolver cckit_gateway.ChaincodeLocatorResolver) *ConfigServiceChaincodeLocatorResolver {
	return &ConfigServiceChaincodeLocatorResolver{
		locatorResolver: locatorResolver,
	}
}

func (r *ConfigServiceChaincodeLocatorResolver) Resolve(ctx cckit_router.Context) (ConfigServiceChaincode, error) {
	if r.service != nil {
		return r.service, nil
	}

	locator, err := r.locatorResolver(ctx, _ConfigService_serviceDesc.ServiceName)
	if err != nil {
		return nil, err
	}

	r.service = NewConfigServiceChaincodeStubInvoker(locator)
	return r.service, nil
}

type ConfigServiceChaincodeStubInvoker struct {
	Invoker cckit_gateway.ChaincodeStubInvoker
}

func NewConfigServiceChaincodeStubInvoker(locator *cckit_gateway.ChaincodeLocator) *ConfigServiceChaincodeStubInvoker {
	return &ConfigServiceChaincodeStubInvoker{
		Invoker: &cckit_gateway.LocatorChaincodeStubInvoker{Locator: locator},
	}
}

func (c *ConfigServiceChaincodeStubInvoker) GetConfig(ctx cckit_router.Context, in *emptypb.Empty) (*Config, error) {

	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker.Query(ctx.Stub(), ConfigServiceChaincode_GetConfig, []interface{}{in}, &Config{}); err != nil {
		return nil, err
	} else {
		return res.(*Config), nil
	}

}

func (c *ConfigServiceChaincodeStubInvoker) SetConfig(ctx cckit_router.Context, in *Config) (*Config, error) {

	return nil, cckit_gateway.ErrInvokeMethodNotAllowed

}

func (c *ConfigServiceChaincodeStubInvoker) GetToken(ctx cckit_router.Context, in *TokenId) (*Token, error) {

	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker.Query(ctx.Stub(), ConfigServiceChaincode_GetToken, []interface{}{in}, &Token{}); err != nil {
		return nil, err
	} else {
		return res.(*Token), nil
	}

}

func (c *ConfigServiceChaincodeStubInvoker) GetDefaultToken(ctx cckit_router.Context, in *emptypb.Empty) (*Token, error) {

	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker.Query(ctx.Stub(), ConfigServiceChaincode_GetDefaultToken, []interface{}{in}, &Token{}); err != nil {
		return nil, err
	} else {
		return res.(*Token), nil
	}

}

func (c *ConfigServiceChaincodeStubInvoker) CreateTokenType(ctx cckit_router.Context, in *CreateTokenTypeRequest) (*TokenType, error) {

	return nil, cckit_gateway.ErrInvokeMethodNotAllowed

}

func (c *ConfigServiceChaincodeStubInvoker) GetTokenType(ctx cckit_router.Context, in *TokenTypeId) (*TokenType, error) {

	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker.Query(ctx.Stub(), ConfigServiceChaincode_GetTokenType, []interface{}{in}, &TokenType{}); err != nil {
		return nil, err
	} else {
		return res.(*TokenType), nil
	}

}

func (c *ConfigServiceChaincodeStubInvoker) ListTokenTypes(ctx cckit_router.Context, in *emptypb.Empty) (*TokenTypes, error) {

	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker.Query(ctx.Stub(), ConfigServiceChaincode_ListTokenTypes, []interface{}{in}, &TokenTypes{}); err != nil {
		return nil, err
	} else {
		return res.(*TokenTypes), nil
	}

}

func (c *ConfigServiceChaincodeStubInvoker) UpdateTokenType(ctx cckit_router.Context, in *UpdateTokenTypeRequest) (*TokenType, error) {

	return nil, cckit_gateway.ErrInvokeMethodNotAllowed

}

func (c *ConfigServiceChaincodeStubInvoker) DeleteTokenType(ctx cckit_router.Context, in *TokenTypeId) (*TokenType, error) {

	return nil, cckit_gateway.ErrInvokeMethodNotAllowed

}

func (c *ConfigServiceChaincodeStubInvoker) GetTokenGroups(ctx cckit_router.Context, in *TokenTypeId) (*TokenGroups, error) {

	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker.Query(ctx.Stub(), ConfigServiceChaincode_GetTokenGroups, []interface{}{in}, &TokenGroups{}); err != nil {
		return nil, err
	} else {
		return res.(*TokenGroups), nil
	}

}

func (c *ConfigServiceChaincodeStubInvoker) CreateTokenGroup(ctx cckit_router.Context, in *CreateTokenGroupRequest) (*TokenGroup, error) {

	return nil, cckit_gateway.ErrInvokeMethodNotAllowed

}

func (c *ConfigServiceChaincodeStubInvoker) GetTokenGroup(ctx cckit_router.Context, in *TokenGroupId) (*TokenGroup, error) {

	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker.Query(ctx.Stub(), ConfigServiceChaincode_GetTokenGroup, []interface{}{in}, &TokenGroup{}); err != nil {
		return nil, err
	} else {
		return res.(*TokenGroup), nil
	}

}

func (c *ConfigServiceChaincodeStubInvoker) DeleteTokenGroup(ctx cckit_router.Context, in *TokenGroupId) (*Token, error) {

	return nil, cckit_gateway.ErrInvokeMethodNotAllowed

}
