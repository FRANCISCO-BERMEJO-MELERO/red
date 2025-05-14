package keeper

import (
	"context"

	"red/x/test/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Holamundo(goCtx context.Context, msg *types.MsgHolamundo) (*types.MsgHolamundoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Imprime "Hola mundo" en los logs del nodo
	ctx.Logger().Info("Hola mundo")

	return &types.MsgHolamundoResponse{}, nil
}
