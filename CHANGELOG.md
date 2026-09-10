# Changelog

## v2.3.0 — 2026-09-10

Initial public Nonsense source publication from the deployed node snapshot.

- Independent Nonsense mainnet genesis, network identifiers, address prefixes,
  bootstrap peers, and FishHashPlus seed.
- Two billion NNN maximum supply parameter and integer monthly subsidy schedule.
- Full node and CLI wallet for Linux x64 and Windows x64.
- Module imports and project links point to `nonsense-project/nonsense`.
- English specifications, node and wallet instructions, and release build scripts.
- Updated inherited difficulty and DAG-window test fixtures for Nonsense genesis
  parameters and isolated the mock difficulty test from activation reset rules.
- Corrected stale port descriptions and removed a redundant boolean term
  reported by Go vet, without changing the expression's behavior.

The deployed source already reported version 2.3.0. This publication retains
that version and the network's consensus rules. No prior Git history was
available in the recovered source snapshot.
