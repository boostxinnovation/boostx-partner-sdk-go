# Changelog

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
