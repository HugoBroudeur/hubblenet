// Code generated
// This file is a generated precompile contract config with stubbed abstract functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package bibliophile

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/precompile/contract"

	_ "embed"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// Gas costs for each function. These are set to 1 by default.
	// You should set a gas cost for each function in your contract.
	// Generally, you should not set gas costs very low as this may cause your network to be vulnerable to DoS attacks.
	// There are some predefined gas costs in contract/utils.go that you can use.
	GetNotionalPositionAndMarginGasCost                  uint64 = 1000 /* SET A GAS COST HERE */
	GetPositionSizesGasCost                              uint64 = 1000 /* SET A GAS COST HERE */
	ValidateLiquidationOrderAndDetermineFillPriceGasCost uint64 = 1000 /* SET A GAS COST HERE */
	ValidateOrdersAndDetermineFillPriceGasCost           uint64 = 1000 /* SET A GAS COST HERE */
)

// CUSTOM CODE STARTS HERE
// Reference imports to suppress errors from unused imports. This code and any unnecessary imports can be removed.
var (
	_ = abi.JSON
	_ = errors.New
	_ = big.NewInt
)

// Singleton StatefulPrecompiledContract and signatures.
var (

	// HubbleBibliophileRawABI contains the raw ABI of HubbleBibliophile contract.
	//go:embed contract.abi
	HubbleBibliophileRawABI string

	HubbleBibliophileABI = contract.ParseABI(HubbleBibliophileRawABI)

	HubbleBibliophilePrecompile = createHubbleBibliophilePrecompile()
)

// Some changes to hubble config manager will require us to keep old as well new version of logic
// Some logic changes may result in usedGas which will result in error during replay of blocks while syncing.
// Some logic changes may result in different state wihch will result in error in state during replay of blocks while syncing.
// We should track these logic change which can cause changes specified above; as releases in comments below

// Release 1
// in amm.go multiply1e6 is diving by 1e6(v1). Change was to fix this to multiply by 1e6(v2).
// This caused different marginFration and thus different output which cause issue while replay of blocks.

// Release 2 - Better Pricing Algorithm (backwards compatible, so will not need an activation time)

// IHubbleBibliophileOrder is an auto generated low-level Go binding around an user-defined struct.
type IHubbleBibliophileOrder struct {
	AmmIndex          *big.Int
	Trader            common.Address
	BaseAssetQuantity *big.Int
	Price             *big.Int
	Salt              *big.Int
	ReduceOnly        bool
}

type GetNotionalPositionAndMarginInput struct {
	Trader                 common.Address
	IncludeFundingPayments bool
	Mode                   uint8
}

type GetNotionalPositionAndMarginOutput struct {
	NotionalPosition *big.Int
	Margin           *big.Int
}

type ValidateLiquidationOrderAndDetermineFillPriceInput struct {
	Order      IHubbleBibliophileOrder
	FillAmount *big.Int
}

type ValidateOrdersAndDetermineFillPriceInput struct {
	Orders      [2]IHubbleBibliophileOrder
	OrderHashes [2][32]byte
	FillAmount  *big.Int
}

type ValidateOrdersAndDetermineFillPriceOutput struct {
	FillPrice *big.Int
	Mode0     uint8
	Mode1     uint8
}

// UnpackGetNotionalPositionAndMarginInput attempts to unpack [input] as GetNotionalPositionAndMarginInput
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackGetNotionalPositionAndMarginInput(input []byte) (GetNotionalPositionAndMarginInput, error) {
	inputStruct := GetNotionalPositionAndMarginInput{}
	err := HubbleBibliophileABI.UnpackInputIntoInterface(&inputStruct, "getNotionalPositionAndMargin", input)

	return inputStruct, err
}

// PackGetNotionalPositionAndMargin packs [inputStruct] of type GetNotionalPositionAndMarginInput into the appropriate arguments for getNotionalPositionAndMargin.
func PackGetNotionalPositionAndMargin(inputStruct GetNotionalPositionAndMarginInput) ([]byte, error) {
	return HubbleBibliophileABI.Pack("getNotionalPositionAndMargin", inputStruct.Trader, inputStruct.IncludeFundingPayments, inputStruct.Mode)
}

// PackGetNotionalPositionAndMarginOutput attempts to pack given [outputStruct] of type GetNotionalPositionAndMarginOutput
// to conform the ABI outputs.
func PackGetNotionalPositionAndMarginOutput(outputStruct GetNotionalPositionAndMarginOutput) ([]byte, error) {
	return HubbleBibliophileABI.PackOutput("getNotionalPositionAndMargin",
		outputStruct.NotionalPosition,
		outputStruct.Margin,
	)
}

func getNotionalPositionAndMargin(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	if remainingGas, err = contract.DeductGas(suppliedGas, GetNotionalPositionAndMarginGasCost); err != nil {
		return nil, 0, err
	}
	// attempts to unpack [input] into the arguments to the GetNotionalPositionAndMarginInput.
	// Assumes that [input] does not include selector
	// You can use unpacked [inputStruct] variable in your code
	inputStruct, err := UnpackGetNotionalPositionAndMarginInput(input)
	if err != nil {
		return nil, remainingGas, err
	}

	// CUSTOM CODE STARTS HERE
	output := GetNotionalPositionAndMargin(accessibleState.GetStateDB(), &inputStruct, accessibleState.GetBlockContext().Timestamp())
	packedOutput, err := PackGetNotionalPositionAndMarginOutput(output)
	if err != nil {
		return nil, remainingGas, err
	}

	// Return the packed output and the remaining gas
	return packedOutput, remainingGas, nil
}

// UnpackGetPositionSizesInput attempts to unpack [input] into the common.Address type argument
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackGetPositionSizesInput(input []byte) (common.Address, error) {
	res, err := HubbleBibliophileABI.UnpackInput("getPositionSizes", input)
	if err != nil {
		return common.Address{}, err
	}
	unpacked := *abi.ConvertType(res[0], new(common.Address)).(*common.Address)
	return unpacked, nil
}

// PackGetPositionSizes packs [trader] of type common.Address into the appropriate arguments for getPositionSizes.
// the packed bytes include selector (first 4 func signature bytes).
// This function is mostly used for tests.
func PackGetPositionSizes(trader common.Address) ([]byte, error) {
	return HubbleBibliophileABI.Pack("getPositionSizes", trader)
}

// PackGetPositionSizesOutput attempts to pack given posSizes of type []*big.Int
// to conform the ABI outputs.
func PackGetPositionSizesOutput(posSizes []*big.Int) ([]byte, error) {
	return HubbleBibliophileABI.PackOutput("getPositionSizes", posSizes)
}

func getPositionSizes(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	if remainingGas, err = contract.DeductGas(suppliedGas, GetPositionSizesGasCost); err != nil {
		return nil, 0, err
	}
	// attempts to unpack [input] into the arguments to the GetPositionSizesInput.
	// Assumes that [input] does not include selector
	// You can use unpacked [inputStruct] variable in your code
	inputStruct, err := UnpackGetPositionSizesInput(input)
	if err != nil {
		return nil, remainingGas, err
	}

	// CUSTOM CODE STARTS HERE
	output := getPosSizes(accessibleState.GetStateDB(), &inputStruct)
	packedOutput, err := PackGetPositionSizesOutput(output)
	if err != nil {
		return nil, remainingGas, err
	}

	// Return the packed output and the remaining gas
	return packedOutput, remainingGas, nil
}

// UnpackValidateLiquidationOrderAndDetermineFillPriceInput attempts to unpack [input] as ValidateLiquidationOrderAndDetermineFillPriceInput
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackValidateLiquidationOrderAndDetermineFillPriceInput(input []byte) (ValidateLiquidationOrderAndDetermineFillPriceInput, error) {
	inputStruct := ValidateLiquidationOrderAndDetermineFillPriceInput{}
	err := HubbleBibliophileABI.UnpackInputIntoInterface(&inputStruct, "validateLiquidationOrderAndDetermineFillPrice", input)

	return inputStruct, err
}

// PackValidateLiquidationOrderAndDetermineFillPrice packs [inputStruct] of type ValidateLiquidationOrderAndDetermineFillPriceInput into the appropriate arguments for validateLiquidationOrderAndDetermineFillPrice.
func PackValidateLiquidationOrderAndDetermineFillPrice(inputStruct ValidateLiquidationOrderAndDetermineFillPriceInput) ([]byte, error) {
	return HubbleBibliophileABI.Pack("validateLiquidationOrderAndDetermineFillPrice", inputStruct.Order, inputStruct.FillAmount)
}

// PackValidateLiquidationOrderAndDetermineFillPriceOutput attempts to pack given fillPrice of type *big.Int
// to conform the ABI outputs.
func PackValidateLiquidationOrderAndDetermineFillPriceOutput(fillPrice *big.Int) ([]byte, error) {
	return HubbleBibliophileABI.PackOutput("validateLiquidationOrderAndDetermineFillPrice", fillPrice)
}

func validateLiquidationOrderAndDetermineFillPrice(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	if remainingGas, err = contract.DeductGas(suppliedGas, ValidateLiquidationOrderAndDetermineFillPriceGasCost); err != nil {
		return nil, 0, err
	}
	// attempts to unpack [input] into the arguments to the ValidateLiquidationOrderAndDetermineFillPriceInput.
	// Assumes that [input] does not include selector
	// You can use unpacked [inputStruct] variable in your code
	inputStruct, err := UnpackValidateLiquidationOrderAndDetermineFillPriceInput(input)
	if err != nil {
		return nil, remainingGas, err
	}

	// CUSTOM CODE STARTS HERE
	output, err := ValidateLiquidationOrderAndDetermineFillPrice(accessibleState.GetStateDB(), &inputStruct)
	if err != nil {
		return nil, remainingGas, err
	}
	packedOutput, err := PackValidateLiquidationOrderAndDetermineFillPriceOutput(output)
	if err != nil {
		return nil, remainingGas, err
	}

	// Return the packed output and the remaining gas
	return packedOutput, remainingGas, nil
}

// UnpackValidateOrdersAndDetermineFillPriceInput attempts to unpack [input] as ValidateOrdersAndDetermineFillPriceInput
// assumes that [input] does not include selector (omits first 4 func signature bytes)
func UnpackValidateOrdersAndDetermineFillPriceInput(input []byte) (ValidateOrdersAndDetermineFillPriceInput, error) {
	inputStruct := ValidateOrdersAndDetermineFillPriceInput{}
	err := HubbleBibliophileABI.UnpackInputIntoInterface(&inputStruct, "validateOrdersAndDetermineFillPrice", input)

	return inputStruct, err
}

// PackValidateOrdersAndDetermineFillPrice packs [inputStruct] of type ValidateOrdersAndDetermineFillPriceInput into the appropriate arguments for validateOrdersAndDetermineFillPrice.
func PackValidateOrdersAndDetermineFillPrice(inputStruct ValidateOrdersAndDetermineFillPriceInput) ([]byte, error) {
	return HubbleBibliophileABI.Pack("validateOrdersAndDetermineFillPrice", inputStruct.Orders, inputStruct.OrderHashes, inputStruct.FillAmount)
}

// PackValidateOrdersAndDetermineFillPriceOutput attempts to pack given [outputStruct] of type ValidateOrdersAndDetermineFillPriceOutput
// to conform the ABI outputs.
func PackValidateOrdersAndDetermineFillPriceOutput(outputStruct ValidateOrdersAndDetermineFillPriceOutput) ([]byte, error) {
	return HubbleBibliophileABI.PackOutput("validateOrdersAndDetermineFillPrice",
		outputStruct.FillPrice,
		outputStruct.Mode0,
		outputStruct.Mode1,
	)
}

func validateOrdersAndDetermineFillPrice(accessibleState contract.AccessibleState, caller common.Address, addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
	if remainingGas, err = contract.DeductGas(suppliedGas, ValidateOrdersAndDetermineFillPriceGasCost); err != nil {
		return nil, 0, err
	}
	// attempts to unpack [input] into the arguments to the ValidateOrdersAndDetermineFillPriceInput.
	// Assumes that [input] does not include selector
	// You can use unpacked [inputStruct] variable in your code
	inputStruct, err := UnpackValidateOrdersAndDetermineFillPriceInput(input)
	if err != nil {
		return nil, remainingGas, err
	}

	// CUSTOM CODE STARTS HERE
	output, err := ValidateOrdersAndDetermineFillPrice(accessibleState.GetStateDB(), &inputStruct)
	if err != nil {
		return nil, remainingGas, err
	}
	packedOutput, err := PackValidateOrdersAndDetermineFillPriceOutput(*output)
	if err != nil {
		return nil, remainingGas, err
	}

	// Return the packed output and the remaining gas
	return packedOutput, remainingGas, nil
}

// createHubbleBibliophilePrecompile returns a StatefulPrecompiledContract with getters and setters for the precompile.

func createHubbleBibliophilePrecompile() contract.StatefulPrecompiledContract {
	var functions []*contract.StatefulPrecompileFunction

	abiFunctionMap := map[string]contract.RunStatefulPrecompileFunc{
		"getNotionalPositionAndMargin":                  getNotionalPositionAndMargin,
		"getPositionSizes":                              getPositionSizes,
		"validateLiquidationOrderAndDetermineFillPrice": validateLiquidationOrderAndDetermineFillPrice,
		"validateOrdersAndDetermineFillPrice":           validateOrdersAndDetermineFillPrice,
	}

	for name, function := range abiFunctionMap {
		method, ok := HubbleBibliophileABI.Methods[name]
		if !ok {
			panic(fmt.Errorf("given method (%s) does not exist in the ABI", name))
		}
		functions = append(functions, contract.NewStatefulPrecompileFunction(method.ID, function))
	}
	// Construct the contract with no fallback function.
	statefulContract, err := contract.NewStatefulPrecompileContract(nil, functions)
	if err != nil {
		panic(err)
	}
	return statefulContract
}
