[简体中文](../../README.md) · [繁體中文](../../docs/i18n/README.zh-TW.md) · **English**

# Carneades · Chinese adaptation

**Make claims, grounds and rebuttals inspectable.**

An independent Chinese adaptation of Carneades 4, preserving its formal argumentation engine, upstream history and traceable Chinese examples.

[Issues](https://github.com/2432antony-13/carneades-4-zh-cn/issues) · [Docs](../../docs/README.md)

## Current status

CLI and Chinese examples run in CI. This independent fork is not an official Chinese release by the original authors or Fraunhofer FOKUS.

## Quick start

CI uses Go 1.26.x. Graph export may require Graphviz. The scoped tests with -vet=off do not establish a repository-wide vet pass.

```bash
go build -o bin/carneades ./src/cmd/carneades
go test -vet=off ./src/cmd/carneades ./src/engine/caes/encoding/yaml
./bin/carneades check -f yaml examples/AGs/YAML/zh-CN/电子数据真实性.yml
```

## Capabilities and scope

Abstract argumentation semantics; structured argument evaluation; YAML example validation; graph export.

MPL-2.0 and upstream attribution are preserved. Formal reasoning results do not establish real-world facts or judicial conclusions.

[Source](../../src/) · [Docs](../../docs/README.md)

Repository homepages and usage entry points are available in three languages. Product interface coverage is tracked separately; documentation translation does not imply full product localization.
