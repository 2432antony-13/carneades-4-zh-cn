<p>
  <img align="right" src="docs/assets/tony.gif" width="12%" alt="Tony 24 · Carneades · 中文适配" />
  <img src="docs/assets/hero.png" width="85%" alt="Carneades · 中文适配 · 把论点、依据与反驳，放到同一张可检验的论证图中。" />
</p>

<br clear="all" />

**简体中文** · [繁體中文](docs/i18n/README.zh-TW.md) · [English](docs/i18n/README.en.md)

# Carneades · 中文适配

**把论点、依据与反驳，放到同一张可检验的论证图中。**

Carneades · 中文适配是一套面向中文学习与研究场景的形式论证工具。它把观点、前提、争点和支持或反驳关系组织为结构化模型，再用明确的语义规则检查和评估。你可以从中文教学示例开始，用命令行重复实验，也可以启动三语网页表单查看结果。项目保留 Carneades 4 上游引擎与来源记录，重点是让建模过程和推理依据能够被理解、复现与讨论。

## 适合谁

- 论证理论、法律推理与可解释 AI 的学习者和研究者。
- 需要验证论证图、比较不同语义或接入结构化数据的开发者。
- 希望以中文示例开展课堂演示和方法讨论的教师。

## 7 个核心模块

### 1. 结构化论证建模

用 statements、arguments、issues 和 assumptions 描述观点、前提与争点，将支持和反驳关系放入同一个模型。争点可以指定证明标准。

**结果：** 得到可保存、可修改的论证图输入，让讨论落到具体前提和关系上。<br />
**状态：** 已实现 · 文件建模

### 2. 输入校验与错误定位

check 命令先检查论证图的语法和语义约束，再报告模型中的问题。可从文件或标准输入读入，适合编辑后反复检查。

**结果：** 在评估前发现不完整或不一致的模型，避免把输入错误误认为推理结果。<br />
**状态：** 已实现 · CLI

### 3. 结构化论证评估

eval 命令在校验后应用模型中的理论，推导相应论证，再使用 grounded 语义更新陈述的接受状态。

**结果：** 得到带有评估结果的论证图，便于比较不同假设下结论如何变化。<br />
**状态：** 已实现 · CLI / 本地网页

### 4. 抽象论证语义计算

dung 命令读取 TGF 论证框架，支持 grounded、complete、preferred 和 stable 四种语义，并可查询论证的可信接受或怀疑接受情况。

**结果：** 查看可接受的论证集合，比较不同语义对同一冲突结构给出的结果。<br />
**状态：** 已实现 · CLI / 本地网页

### 5. 格式互通与图形导出

结构化输入支持 YAML、CAF、AIF、AGXML 和 LKIF；CLI 可输出 YAML、DOT、GraphML，网页另提供 PNG 和 SVG 图形输出。

**结果：** 既能保存机器可读结果，也能把论证关系带入图形工具或讨论材料。<br />
**状态：** 已实现 · 图形渲染需 Graphviz

### 6. 中文教学示例

仓库提供“电子数据真实性”YAML 示例，展示原始设备、哈希核验与编辑痕迹如何构成支持和反驳，并写有预期接受状态。

**结果：** 从一个可检查的小模型入手，理解中文文本、论证结构和测试断言如何对应。<br />
**状态：** 已提供 · 虚构教学模型

### 7. 三语本地操作界面

启动本地 Web 服务后，可通过表单上传模型、选择输入格式、语义和输出格式；首页、操作说明等页面支持简体中文、繁體中文与 English。

**结果：** 不必为每次演示重输完整命令，也能在不同语言下阅读相同功能说明。<br />
**状态：** 已实现 · 自行运行服务

## 如何使用

1. 先阅读中文教学示例，确认要表达的观点、前提与争点。
2. 运行 check 校验模型，再用 eval 或 dung 选择对应的评估方式。
3. 查看文本或图形结果，修改假设与关系后再次运行，记录输入和结果差异。

## 当前完成程度

- 当前交付为源码、中文资料、CLI 和可自行启动的三语 Web 界面；没有本轮已验收的公共在线服务。
- 命令行构建、局部测试与中文示例检查已纳入现有 CI，支持持续检查这条基础使用路径。
- 这是独立衍生项目，不是 Carneades 原作者或 Fraunhofer FOKUS 的官方中文发行。

## 本地开始

在仓库根目录执行。CI 使用 Go 1.26.x；需要 PNG/SVG 渲染时安装 Graphviz。下列先检查教学示例，再导出评估结果；如需网页，将模板目录替换为本机绝对路径。

```bash
go build -o bin/carneades ./src/cmd/carneades
./bin/carneades check -f yaml examples/AGs/YAML/zh-CN/电子数据真实性.yml
./bin/carneades eval -f yaml -t yaml examples/AGs/YAML/zh-CN/电子数据真实性.yml
./bin/carneades server -p 8080 -t "/absolute/path/to/src/web/templates" -l zh-CN
# http://localhost:8080/carneades/
```

## 范围与边界

- 形式评估只反映输入模型及选定规则，不自动验证证据真伪，也不构成司法结论。
- Web 端对超过 20 个论证的 Dung 框架仅支持 grounded 语义；CLI 与 Web 的可选输出并不完全相同。
- 局部测试使用 -vet=off，不能据此声称全仓 vet 通过。界面翻译不自动翻译输入文件；继续遵循 MPL-2.0 及上游归属要求。

## 文档与参与

[使用说明](https://github.com/2432antony-13/carneades-4-zh-cn/blob/master/docs/USAGE.md) · [中文示例](https://github.com/2432antony-13/carneades-4-zh-cn/tree/master/examples/AGs/YAML/zh-CN) · [任务与反馈](https://github.com/2432antony-13/carneades-4-zh-cn/issues) · [上游与许可声明](https://github.com/2432antony-13/carneades-4-zh-cn/blob/master/docs/NOTICE.zh-CN.md)
