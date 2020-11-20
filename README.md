# Sommelier 3.0 Stargate 

Sommelier 3.0 is a test version of the Stargate binaries. It is intended to use in testing compatibiltiy with the post Stargate ugprade 

The biggest change is that there is no longer two seperate binaries. There is just `sommelier`.

The rest and new GRPC interfaces can be configured using the `app.toml`.

You can interact via the cli interface using a second instance of the binary while a full node is running.

Key an eye on our [audit](https://github.com/cosmosdevs/stargate/pull/8) of interface of changes to help with upgrading.

##

Sommelier is the first implementation of the Cosmos Hub, built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk).  Sommelier and other Cosmos Hubs allow fully sovereign blockchains to interact with one another using a protocol called [IBC](https://github.com/cosmos/ics/tree/master/ibc) that enables Inter-Blockchain Communication.

[![CircleCI](https://circleci.com/gh/peggyjv/sommelier/tree/master.svg?style=shield)](https://circleci.com/gh/peggyjv/sommelier/tree/master)
[![codecov](https://codecov.io/gh/peggyjv/sommelier/branch/master/graph/badge.svg)](https://codecov.io/gh/peggyjv/sommelier)
[![Go Report Card](https://goreportcard.com/badge/github.com/peggyjv/sommelier)](https://goreportcard.com/report/github.com/peggyjv/sommelier)
[![license](https://img.shields.io/github/license/peggyjv/sommelier.svg)](https://github.com/peggyjv/sommelier/blob/master/LICENSE)
[![LoC](https://tokei.rs/b1/github/peggyjv/sommelier)](https://github.com/peggyjv/sommelier)
[![GolangCI](https://golangci.com/badges/github.com/peggyjv/sommelier.svg)](https://golangci.com/r/github.com/peggyjv/sommelier)


## Talk to us!

We have active, helpful communities on Twitter, Discord, and Telegram.

* [Discord](https://discord.gg/huHEBUX)
* [Twitter](https://twitter.com/cosmos)
* [Telegram](https://t.me/cosmosproject)

## Archives

With each version of the Cosmos Hub, the chain is restarted from a new Genesis state.  We are currently on cosmoshub-3.  Archives of the state of cosmoshub-1 and cosmoshub-2 are available [here](./docs/resources/archives.md).

## Disambiguation

Sommelier is not related to the [React-Cosmos](https://github.com/react-cosmos/react-cosmos) project (yet). Many thanks to Evan Coury and Ovidiu (@skidding) for this Github organization name. Per our agreement, this disambiguation notice will stay here.
