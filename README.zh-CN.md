# Carneades 4 中国证据论证研究版

本仓库以 [Carneades 4](https://github.com/carneades/carneades-4) 为代码基础，在保留上游可废止论证内核、提交历史和许可证的前提下，进行简体中文本地化，并研究适用于中国证据推理的 CLES-MoE 专家路由、事实账本和可审计论证架构。

## 目标完成度

| 老师提出的目标 | 当前判断 | 验收依据 |
|---|---|---|
| 对 Carneades 进行汉化 | 已完成首轮可运行原型 | 九个中文 Web 页面、`-l zh-CN` 切换、中文帮助和术语文档 |
| 在国内用于证据推理研究和教学 | 已具备演示条件 | 中文 YAML 可导入、评估和导出；已有电子数据真实性合成示例 |
| 建成可供国内实务使用的专家系统 | 尚未完成 | 中国证明标准、事实账本、专家路由、规则库和人工复核流程尚未成为可执行模块 |

因此，当前版本适合向老师展示“汉化基础已经完成，并已进入中国证据推理适配研究”，
不宜表述为已经能够自动审查真实案件证据。

## 项目边界

- 本项目保留官方仓库的提交历史和 MPL-2.0 许可证。
- 中文界面、中文文档和中文示例由本分支独立编写。
- 本项目不隶属于 Carneades 原作者或 Fraunhofer FOKUS。
- 本项目是研究与教学工具，不提供自动裁判或法律意见。
- 软件标签和权重是形式模型的计算结果，不等同于中国法上的事实认定或证明标准。
- 当前仓库尚未训练参数级神经 MoE；现阶段提供的是研究架构、中文语境适配和符号原型。

## 已完成的中文化

- Web 操作界面及帮助页面；
- 中文快速开始文档；
- 法律论证核心术语表；
- 中文 YAML 论证图示例；
- 中文文本 YAML 导入导出测试；
- 来源、许可和研究贡献记录；
- 中国证据推理 MoE 研究报告、来源台账和工程架构。

## 快速运行

项目需要 Go 1.18 或更高版本。PNG 和 SVG 导出还需要 Graphviz。

```bash
go build -o bin/carneades ./src/cmd/carneades
./bin/carneades server \
  -p 8080 \
  -t ./src/web/templates \
  -l zh-CN
```

然后访问 `http://localhost:8080/carneades/`。

英文界面仍可使用：

```bash
./bin/carneades server -p 8080 -t ./src/web/templates -l en
```

更多说明：

- [中文快速开始](docs/zh-CN/快速开始.md)
- [中英文术语表](docs/zh-CN/术语表.md)
- [来源、许可与原创性记录](docs/zh-CN/来源与许可.md)
- [汉化范围与未完成事项](docs/zh-CN/汉化范围.md)
- [两日交付清单](docs/zh-CN/两日交付清单.md)
- [CLES-MoE 中国证据推理架构](docs/zh-CN/中国证据推理架构.md)
- [中国证据推理 MoE 架构研究报告](research/zh-CN/中国证据推理MoE架构研究报告.md)
- [研究来源台账](research/zh-CN/研究来源台账.md)
- [研究贡献与来源说明](research/zh-CN/研究贡献与来源说明.md)

## 中文示例

`examples/AGs/YAML/zh-CN/电子数据真实性.yml` 是独立编写的教学示例，用于演示支持论证、反对论证、争点和中文文本。它不对应真实案件，也不构成对电子数据真实性的法律判断。

## 上游与许可证

- 上游项目：<https://github.com/carneades/carneades-4>
- 原始许可证：[MPL-2.0](LICENSE)
- 本地化声明：[NOTICE.zh-CN.md](NOTICE.zh-CN.md)

修改受 MPL-2.0 约束的源文件时，应继续遵守 MPL-2.0。中文研究材料引用论文或法律文件时，应遵守相应的引注、许可和学术规范。
