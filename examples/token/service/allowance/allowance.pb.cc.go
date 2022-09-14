// Code generated by protoc-gen-cc-gateway. DO NOT EDIT.
// source: token/service/allowance/allowance.proto

/*
Package allowance contains
  *   chaincode methods names {service_name}Chaincode_{method_name}
  *   chaincode interface definition {service_name}Chaincode
  *   chaincode gateway definition {service_name}}Gateway
  *   chaincode service to cckit router registration func
*/
package allowance

import (
	context "context"
	_ "embed"
	errors "errors"

	cckit_gateway "github.com/hyperledger-labs/cckit/gateway"
	cckit_router "github.com/hyperledger-labs/cckit/router"
	cckit_defparam "github.com/hyperledger-labs/cckit/router/param/defparam"
	cckit_sdk "github.com/hyperledger-labs/cckit/sdk"
)

// AllowanceServiceChaincode method names
const (

	// AllowanceServiceChaincodeMethodPrefix allows to use multiple services with same method names in one chaincode
	AllowanceServiceChaincodeMethodPrefix = ""

	AllowanceServiceChaincode_GetAllowance = AllowanceServiceChaincodeMethodPrefix + "GetAllowance"

	AllowanceServiceChaincode_Approve = AllowanceServiceChaincodeMethodPrefix + "Approve"

	AllowanceServiceChaincode_TransferFrom = AllowanceServiceChaincodeMethodPrefix + "TransferFrom"
)

// AllowanceServiceChaincode chaincode methods interface
type AllowanceServiceChaincode interface {
	GetAllowance(cckit_router.Context, *AllowanceRequest) (*Allowance, error)

	Approve(cckit_router.Context, *ApproveRequest) (*Allowance, error)

	TransferFrom(cckit_router.Context, *TransferFromRequest) (*TransferFromResponse, error)
}

// RegisterAllowanceServiceChaincode registers service methods as chaincode router handlers
func RegisterAllowanceServiceChaincode(r *cckit_router.Group, cc AllowanceServiceChaincode) error {

	r.Query(AllowanceServiceChaincode_GetAllowance,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.GetAllowance(ctx, ctx.Param().(*AllowanceRequest))
		},
		cckit_defparam.Proto(&AllowanceRequest{}))

	r.Invoke(AllowanceServiceChaincode_Approve,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.Approve(ctx, ctx.Param().(*ApproveRequest))
		},
		cckit_defparam.Proto(&ApproveRequest{}))

	r.Invoke(AllowanceServiceChaincode_TransferFrom,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.TransferFrom(ctx, ctx.Param().(*TransferFromRequest))
		},
		cckit_defparam.Proto(&TransferFromRequest{}))

	return nil
}

//go:embed allowance.swagger.json
var AllowanceServiceSwagger []byte

// NewAllowanceServiceGateway creates gateway to access chaincode method via chaincode service
func NewAllowanceServiceGateway(sdk cckit_sdk.SDK, channel, chaincode string, opts ...cckit_gateway.Opt) *AllowanceServiceGateway {
	return NewAllowanceServiceGatewayFromInstance(
		cckit_gateway.NewChaincodeInstanceService(
			sdk,
			&cckit_gateway.ChaincodeLocator{Channel: channel, Chaincode: chaincode},
			opts...,
		))
}

func NewAllowanceServiceGatewayFromInstance(chaincodeInstance cckit_gateway.ChaincodeInstance) *AllowanceServiceGateway {
	return &AllowanceServiceGateway{
		ChaincodeInstance: chaincodeInstance,
	}
}

// gateway implementation
// gateway can be used as kind of SDK, GRPC or REST server ( via grpc-gateway or clay )
type AllowanceServiceGateway struct {
	ChaincodeInstance cckit_gateway.ChaincodeInstance
}

func (c *AllowanceServiceGateway) Invoker() cckit_gateway.ChaincodeInstanceInvoker {
	return cckit_gateway.NewChaincodeInstanceServiceInvoker(c.ChaincodeInstance)
}

// ServiceDef returns service definition
func (c *AllowanceServiceGateway) ServiceDef() cckit_gateway.ServiceDef {
	return cckit_gateway.NewServiceDef(
		_AllowanceService_serviceDesc.ServiceName,
		AllowanceServiceSwagger,
		&_AllowanceService_serviceDesc,
		c,
		RegisterAllowanceServiceHandlerFromEndpoint,
	)
}

func (c *AllowanceServiceGateway) GetAllowance(ctx context.Context, in *AllowanceRequest) (*Allowance, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker().Query(ctx, AllowanceServiceChaincode_GetAllowance, []interface{}{in}, &Allowance{}); err != nil {
		return nil, err
	} else {
		return res.(*Allowance), nil
	}
}

func (c *AllowanceServiceGateway) Approve(ctx context.Context, in *ApproveRequest) (*Allowance, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker().Invoke(ctx, AllowanceServiceChaincode_Approve, []interface{}{in}, &Allowance{}); err != nil {
		return nil, err
	} else {
		return res.(*Allowance), nil
	}
}

func (c *AllowanceServiceGateway) TransferFrom(ctx context.Context, in *TransferFromRequest) (*TransferFromResponse, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker().Invoke(ctx, AllowanceServiceChaincode_TransferFrom, []interface{}{in}, &TransferFromResponse{}); err != nil {
		return nil, err
	} else {
		return res.(*TransferFromResponse), nil
	}
}

// AllowanceServiceChaincodeResolver interface for service resolver
type (
	AllowanceServiceChaincodeResolver interface {
		Resolve(ctx cckit_router.Context) (AllowanceServiceChaincode, error)
	}

	AllowanceServiceChaincodeLocalResolver struct {
		service AllowanceServiceChaincode
	}

	AllowanceServiceChaincodeLocatorResolver struct {
		locatorResolver cckit_gateway.ChaincodeLocatorResolver
		service         AllowanceServiceChaincode
	}
)

func NewAllowanceServiceChaincodeLocalResolver(service AllowanceServiceChaincode) *AllowanceServiceChaincodeLocalResolver {
	return &AllowanceServiceChaincodeLocalResolver{
		service: service,
	}
}

func (r *AllowanceServiceChaincodeLocalResolver) Resolve(ctx cckit_router.Context) (AllowanceServiceChaincode, error) {
	if r.service == nil {
		return nil, errors.New("service not set for local chaincode resolver")
	}

	return r.service, nil
}

func NewAllowanceServiceChaincodeResolver(locatorResolver cckit_gateway.ChaincodeLocatorResolver) *AllowanceServiceChaincodeLocatorResolver {
	return &AllowanceServiceChaincodeLocatorResolver{
		locatorResolver: locatorResolver,
	}
}

func (r *AllowanceServiceChaincodeLocatorResolver) Resolve(ctx cckit_router.Context) (AllowanceServiceChaincode, error) {
	if r.service != nil {
		return r.service, nil
	}

	locator, err := r.locatorResolver(ctx, _AllowanceService_serviceDesc.ServiceName)
	if err != nil {
		return nil, err
	}

	r.service = NewAllowanceServiceChaincodeStubInvoker(locator)
	return r.service, nil
}

type AllowanceServiceChaincodeStubInvoker struct {
	Invoker cckit_gateway.ChaincodeStubInvoker
}

func NewAllowanceServiceChaincodeStubInvoker(locator *cckit_gateway.ChaincodeLocator) *AllowanceServiceChaincodeStubInvoker {
	return &AllowanceServiceChaincodeStubInvoker{
		Invoker: &cckit_gateway.LocatorChaincodeStubInvoker{Locator: locator},
	}
}

func (c *AllowanceServiceChaincodeStubInvoker) GetAllowance(ctx cckit_router.Context, in *AllowanceRequest) (*Allowance, error) {

	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker.Query(ctx.Stub(), AllowanceServiceChaincode_GetAllowance, []interface{}{in}, &Allowance{}); err != nil {
		return nil, err
	} else {
		return res.(*Allowance), nil
	}

}

func (c *AllowanceServiceChaincodeStubInvoker) Approve(ctx cckit_router.Context, in *ApproveRequest) (*Allowance, error) {

	return nil, cckit_gateway.ErrInvokeMethodNotAllowed

}

func (c *AllowanceServiceChaincodeStubInvoker) TransferFrom(ctx cckit_router.Context, in *TransferFromRequest) (*TransferFromResponse, error) {

	return nil, cckit_gateway.ErrInvokeMethodNotAllowed

}
