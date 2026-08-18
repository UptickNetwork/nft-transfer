# nft-transfer

`nft-transfer` is an application that enables cross chain NFT transfer. It implements all functions defined in [ics-721](https://github.com/cosmos/ibc/tree/main/spec/app/ics-721-nft-transfer) using go language.

This repository is the [Uptick Network](https://github.com/UptickNetwork) fork of [`bianjieai/nft-transfer`](https://github.com/bianjieai/nft-transfer), upgraded to **ibc-go v10** and Cosmos SDK v0.53. The Go module path remains `github.com/bianjieai/nft-transfer`; consumers should `replace` it to this repository, for example:

```go
replace github.com/bianjieai/nft-transfer => github.com/UptickNetwork/nft-transfer v1.2.0-ibc-v10
```
