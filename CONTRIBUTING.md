# Contributing to Nonsense

Open issues and pull requests at https://github.com/nonsense-project/nonsense.
Target the `main` branch and describe the behavior being changed, the reason,
and the validation performed.

Use the Go toolchain recorded in `go.mod`. Format changed Go files with
`gofmt`, run `go mod verify`, and run the relevant tests. For a broader check:

```bash
go test -short -p 2 -parallel 2 -timeout 30m ./...
```

Changes to consensus rules, network identifiers, genesis data, proof of work,
or the subsidy table require explicit compatibility review. Do not modify
these parameters merely to make a legacy test pass.

Release archives are produced by `scripts/build_release.py` and contain only
the node, CLI wallet, and required license. Preserve upstream attribution.
