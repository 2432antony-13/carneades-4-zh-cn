**简体中文** · [繁體中文](docs/i18n/README.zh-TW.md) · [English](docs/i18n/README.en.md)

# Carneades · 中文适配

**让论点、依据与反驳可以被检验。**

Carneades 4 的独立中文适配分支，保留形式论证引擎、上游历史与可追溯的中文示例。

[任务](https://github.com/2432antony-13/carneades-4-zh-cn/issues) · [文档](docs/README.md)

## 当前状态

命令行与中文示例已接入 CI。它是独立衍生项目，不是原作者或 Fraunhofer FOKUS 的官方中文发行。

## 快速开始

CI 使用 Go 1.26.x。图形导出可能需要 Graphviz；局部测试的 -vet=off 不代表全仓 vet 通过。

```bash
go build -o bin/carneades ./src/cmd/carneades
go test -vet=off ./src/cmd/carneades ./src/engine/caes/encoding/yaml
./bin/carneades check -f yaml examples/AGs/YAML/zh-CN/电子数据真实性.yml
```

## 能力与范围

抽象论证语义；结构化论证评估；YAML 示例检查；图形导出。

保留 MPL-2.0 及上游归属。形式推理结果不等于业务事实或司法结论。

[源码](src/) · [文档](docs/README.md)

当前首页和使用入口提供三语；各产品界面的语言覆盖单独记录，不以文档翻译代替产品本地化。
