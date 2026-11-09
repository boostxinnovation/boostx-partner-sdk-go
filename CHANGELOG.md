# Changelog

## v1.2.0

Conforms to [boostx-partner-sdk-spec v1.2.0](https://github.com/boostxinnovation/boostx-partner-sdk-spec/tree/v1.2.0),
whose version number this SDK now carries: the spec and both reference SDKs are released together under one number.

Documentation only: no behaviour change. The README now says that a jackpot round keeps, on every attempt, the `X` of
the entry during which BoostX fixed its result, normally the entry that first sent it, and that `PlayerX` gives the X
the player saw on a jackpot round too, since BoostX computes it the same way. The README links the spec's rules instead
of the partner docs, and those links point to spec v1.2.0.

## v1.1.0

Conforms to [boostx-partner-sdk-spec v1.1.0](https://github.com/boostxinnovation/boostx-partner-sdk-spec/tree/v1.1.0).

`BuildGID`, `CreateGamePassToken`, `CreateSettlementToken` and `SubmitSettlement` now refuse a `Partner`, `User` or
`Bet` that contains `:` or is longer than 32, 64 or 96 characters, counted as Unicode code points
(`utf8.RuneCountInString`), with `ErrInvalidClaim` naming that id, before signing. A call that v1.0.0 signed fails when
its ids break either rule. BoostX refuses a gamepass whose ids break them with `401`, and the settlement builder refuses
them too, so a settlement is built only for ids a gamepass can carry: an id fault is a fault in your code. Only `:`
itself is refused, and ids are never trimmed or normalised. Each id runs its checks (valid UTF-8, no `:`, length) before
the next, after every other parameter check. Parsing an inbound token checks neither rule.

## v1.0.0

Initial release. Conforms to [boostx-partner-sdk-spec v1.0.0](https://github.com/boostxinnovation/boostx-partner-sdk-spec/tree/v1.0.0).
