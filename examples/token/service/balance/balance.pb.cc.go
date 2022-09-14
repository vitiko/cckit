// Code generated by protoc-gen-cc-gateway. DO NOT EDIT.
// source: token/service/balance/balance.proto

/*
Package balance contains
  *   chaincode methods names {service_name}Chaincode_{method_name}
  *   chaincode interface definition {service_name}Chaincode
  *   chaincode gateway definition {service_name}}Gateway
  *   chaincode service to cckit router registration func
*/
package balance

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

// BalanceServiceChaincode method names
const (

	// BalanceServiceChaincodeMethodPrefix allows to use multiple services with same method names in one chaincode
	BalanceServiceChaincodeMethodPrefix = ""

	BalanceServiceChaincode_GetBalance = BalanceServiceChaincodeMethodPrefix + "GetBalance"

	BalanceServiceChaincode_ListBalances = BalanceServiceChaincodeMethodPrefix + "ListBalances"

	BalanceServiceChaincode_ListAddressBalances = BalanceServiceChaincodeMethodPrefix + "ListAddressBalances"

	BalanceServiceChaincode_Transfer = BalanceServiceChaincodeMethodPrefix + "Transfer"
)

// BalanceServiceChaincode chaincode methods interface
type BalanceServiceChaincode interface {
	GetBalance(cckit_router.Context, *GetBalanceRequest) (*Balance, error)

	ListBalances(cckit_router.Context, *emptypb.Empty) (*Balances, error)

	ListAddressBalances(cckit_router.Context, *ListAddressBalancesRequest) (*Balances, error)

	Transfer(cckit_router.Context, *TransferRequest) (*TransferResponse, error)
}

// RegisterBalanceServiceChaincode registers service methods as chaincode router handlers
func RegisterBalanceServiceChaincode(r *cckit_router.Group, cc BalanceServiceChaincode) error {

	r.Query(BalanceServiceChaincode_GetBalance,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.GetBalance(ctx, ctx.Param().(*GetBalanceRequest))
		},
		cckit_defparam.Proto(&GetBalanceRequest{}))

	r.Query(BalanceServiceChaincode_ListBalances,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.ListBalances(ctx, ctx.Param().(*emptypb.Empty))
		},
		cckit_defparam.Proto(&emptypb.Empty{}))

	r.Query(BalanceServiceChaincode_ListAddressBalances,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.ListAddressBalances(ctx, ctx.Param().(*ListAddressBalancesRequest))
		},
		cckit_defparam.Proto(&ListAddressBalancesRequest{}))

	r.Invoke(BalanceServiceChaincode_Transfer,
		func(ctx cckit_router.Context) (interface{}, error) {
			return cc.Transfer(ctx, ctx.Param().(*TransferRequest))
		},
		cckit_defparam.Proto(&TransferRequest{}))

	return nil
}

//go:embed balance.swagger.json
var BalanceServiceSwagger []byte

// NewBalanceServiceGateway creates gateway to access chaincode method via chaincode service
func NewBalanceServiceGateway(sdk cckit_sdk.SDK, channel, chaincode string, opts ...cckit_gateway.Opt) *BalanceServiceGateway {
	return NewBalanceServiceGatewayFromInstance(
		cckit_gateway.NewChaincodeInstanceService(
			sdk,
			&cckit_gateway.ChaincodeLocator{Channel: channel, Chaincode: chaincode},
			opts...,
		))
}

func NewBalanceServiceGatewayFromInstance(chaincodeInstance cckit_gateway.ChaincodeInstance) *BalanceServiceGateway {
	return &BalanceServiceGateway{
		ChaincodeInstance: chaincodeInstance,
	}
}

// gateway implementation
// gateway can be used as kind of SDK, GRPC or REST server ( via grpc-gateway or clay )
type BalanceServiceGateway struct {
	ChaincodeInstance cckit_gateway.ChaincodeInstance
}

func (c *BalanceServiceGateway) Invoker() cckit_gateway.ChaincodeInstanceInvoker {
	return cckit_gateway.NewChaincodeInstanceServiceInvoker(c.ChaincodeInstance)
}

// ServiceDef returns service definition
func (c *BalanceServiceGateway) ServiceDef() cckit_gateway.ServiceDef {
	return cckit_gateway.NewServiceDef(
		_BalanceService_serviceDesc.ServiceName,
		BalanceServiceSwagger,
		&_BalanceService_serviceDesc,
		c,
		RegisterBalanceServiceHandlerFromEndpoint,
	)
}

func (c *BalanceServiceGateway) GetBalance(ctx context.Context, in *GetBalanceRequest) (*Balance, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker().Query(ctx, BalanceServiceChaincode_GetBalance, []interface{}{in}, &Balance{}); err != nil {
		return nil, err
	} else {
		return res.(*Balance), nil
	}
}

func (c *BalanceServiceGateway) ListBalances(ctx context.Context, in *emptypb.Empty) (*Balances, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker().Query(ctx, BalanceServiceChaincode_ListBalances, []interface{}{in}, &Balances{}); err != nil {
		return nil, err
	} else {
		return res.(*Balances), nil
	}
}

func (c *BalanceServiceGateway) ListAddressBalances(ctx context.Context, in *ListAddressBalancesRequest) (*Balances, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker().Query(ctx, BalanceServiceChaincode_ListAddressBalances, []interface{}{in}, &Balances{}); err != nil {
		return nil, err
	} else {
		return res.(*Balances), nil
	}
}

func (c *BalanceServiceGateway) Transfer(ctx context.Context, in *TransferRequest) (*TransferResponse, error) {
	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker().Invoke(ctx, BalanceServiceChaincode_Transfer, []interface{}{in}, &TransferResponse{}); err != nil {
		return nil, err
	} else {
		return res.(*TransferResponse), nil
	}
}

// BalanceServiceChaincodeResolver interface for service resolver
type (
	BalanceServiceChaincodeResolver interface {
		Resolve(ctx cckit_router.Context) (BalanceServiceChaincode, error)
	}

	BalanceServiceChaincodeLocalResolver struct {
		service BalanceServiceChaincode
	}

	BalanceServiceChaincodeLocatorResolver struct {
		locatorResolver cckit_gateway.ChaincodeLocatorResolver
		service         BalanceServiceChaincode
	}
)

func NewBalanceServiceChaincodeLocalResolver(service BalanceServiceChaincode) *BalanceServiceChaincodeLocalResolver {
	return &BalanceServiceChaincodeLocalResolver{
		service: service,
	}
}

func (r *BalanceServiceChaincodeLocalResolver) Resolve(ctx cckit_router.Context) (BalanceServiceChaincode, error) {
	if r.service == nil {
		return nil, errors.New("service not set for local chaincode resolver")
	}

	return r.service, nil
}

func NewBalanceServiceChaincodeResolver(locatorResolver cckit_gateway.ChaincodeLocatorResolver) *BalanceServiceChaincodeLocatorResolver {
	return &BalanceServiceChaincodeLocatorResolver{
		locatorResolver: locatorResolver,
	}
}

func (r *BalanceServiceChaincodeLocatorResolver) Resolve(ctx cckit_router.Context) (BalanceServiceChaincode, error) {
	if r.service != nil {
		return r.service, nil
	}

	locator, err := r.locatorResolver(ctx, _BalanceService_serviceDesc.ServiceName)
	if err != nil {
		return nil, err
	}

	r.service = NewBalanceServiceChaincodeStubInvoker(locator)
	return r.service, nil
}

type BalanceServiceChaincodeStubInvoker struct {
	Invoker cckit_gateway.ChaincodeStubInvoker
}

func NewBalanceServiceChaincodeStubInvoker(locator *cckit_gateway.ChaincodeLocator) *BalanceServiceChaincodeStubInvoker {
	return &BalanceServiceChaincodeStubInvoker{
		Invoker: &cckit_gateway.LocatorChaincodeStubInvoker{Locator: locator},
	}
}

func (c *BalanceServiceChaincodeStubInvoker) GetBalance(ctx cckit_router.Context, in *GetBalanceRequest) (*Balance, error) {

	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker.Query(ctx.Stub(), BalanceServiceChaincode_GetBalance, []interface{}{in}, &Balance{}); err != nil {
		return nil, err
	} else {
		return res.(*Balance), nil
	}

}

func (c *BalanceServiceChaincodeStubInvoker) ListBalances(ctx cckit_router.Context, in *emptypb.Empty) (*Balances, error) {

	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker.Query(ctx.Stub(), BalanceServiceChaincode_ListBalances, []interface{}{in}, &Balances{}); err != nil {
		return nil, err
	} else {
		return res.(*Balances), nil
	}

}

func (c *BalanceServiceChaincodeStubInvoker) ListAddressBalances(ctx cckit_router.Context, in *ListAddressBalancesRequest) (*Balances, error) {

	var inMsg interface{} = in
	if v, ok := inMsg.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return nil, err
		}
	}

	if res, err := c.Invoker.Query(ctx.Stub(), BalanceServiceChaincode_ListAddressBalances, []interface{}{in}, &Balances{}); err != nil {
		return nil, err
	} else {
		return res.(*Balances), nil
	}

}

func (c *BalanceServiceChaincodeStubInvoker) Transfer(ctx cckit_router.Context, in *TransferRequest) (*TransferResponse, error) {

	return nil, cckit_gateway.ErrInvokeMethodNotAllowed

}