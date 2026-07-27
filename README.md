# Carneades 4 中国证据论证研究版

本分支在 [Carneades 4](https://github.com/carneades/carneades-4) 可废止论证内核基础上，
完成简体中文界面、中文结构化论证示例和中国证据推理适配研究。

## 当前结论

| 目标 | 状态 | 说明 |
|---|---|---|
| 将 Carneades 系统汉化并运行 | 已完成首轮原型 | 可通过 `-l zh-CN` 启动九个中文 Web 页面 |
| 作为国内证据推理研究模型 | 已形成可演示基础 | 可运行中文电子数据真实性论证图，保留支持、反对和截断关系 |
| 作为国内实务专家系统使用 | 尚未完成 | 中国证明标准、事实账本、专家路由和法律规则库仍需进入可执行内核 |

建议依次查看：

1. [简体中文说明](README.zh-CN.md)
2. [中国语境适配范围](docs/zh-CN/汉化范围.md)
3. [中文电子数据真实性示例](examples/AGs/YAML/zh-CN/电子数据真实性.yml)
4. [中国证据推理架构研究报告](research/zh-CN/中国证据推理MoE架构研究报告.md)

当前定位是研究与教学原型。形式计算结果不等于证据采信、事实认定或法律意见。

---

## Upstream project

These are the source files of Version 4 of the
Carneades argumentation system, written in the Go programming language.

> This fork adds an independent Simplified Chinese localization. It is not an
> official release of the Carneades authors or Fraunhofer FOKUS. The original
> project history and MPL-2.0 license are preserved.

This source code is subject to the terms of the Mozilla Public
License, version 2.0 (MPL-2.0). If a copy of the MPL was not
distributed with this software, it is also available online at
<http://mozilla.org/MPL/2.0/>.  For futher information about the MPL see <http://www.mozilla.org/MPL/2.0/FAQ.html>.

This version of Carneades consists of:

- An implementation of a solver for Dung abstract argumentation frameworks,
using grounded, complete, preferred and stable semantics. Argumentation Frameworks
can be represented using the [Trivial Graph Format](https://en.wikipedia.org/wiki/Trivial_Graph_Format). The computed extensions can be exported to DOT, GraphML and plain text.
- An evaluator for structured arguments, based on a new version of the 
Carneades Argument Evaluation Structures (CAES) formal model of argument. 
- New features of version 4 of Carneades include:
    * Support for cyclic argument graphs, cumulative arguments, issue-based information systems (IBIS) and multicriteria decision analysis.
    * Argument graphs can be represented in AGXML, AIF, LKIF, CAF and YAML and exported to DOT, GraphML, PNG, SVG and YAML.
    * User-definable argument weighing functions, for computing the relative weights of arguments based on their properties, such as the authority or effective date of the argumentation scheme applied, or the labels (in, out, undecided) of premises.  The failure of a premise can weaken or even strengthen an argument without defeating it entirely.
    * Automatic argument construction by applying argumentation schemes to assumptions, via an inference engine implemented using [Constraint Handling Rules](https://dtai.cs.kuleuven.be/CHR/). 

Carneades can visualize both Dung abstract argumentation frameworks and CAES argument graphs using [DOT](http://www.graphviz.org/content/dot-language) and [GraphML](https://en.wikipedia.org/wiki/GraphML).  We recommend using the free [yEd](http://www.yworks.com/yed) GraphML editor to view the GraphML files.


Please see the Carneades blog at <https://carneades.github.io/> for
announcements about the further development of this version.

See the INSTALL.md file in the same directory as this README for
installation instructions.
