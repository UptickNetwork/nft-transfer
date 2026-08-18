# testing

ICS-721 keeper tests use [`github.com/cosmos/ibc-go/v10/testing`](https://github.com/cosmos/ibc-go) with a local SimApp.

- `simapp/` is ibc-go v10.5.0 `testing/simapp` plus cosmos `x/nft` and this module. Do not import the deleted v8 chain/coordinator copies.
- `mock/` is the cosmos `x/nft` adapter (`Wrap`, `ClassMetadata`) required by `types.NFTKeeper`.
- Tests construct coordinators with `simapp.SetupTestingApp`.
