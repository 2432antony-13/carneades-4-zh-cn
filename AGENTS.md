# carneades-4-zh-cn 协作约定

先读 [项目章程](PROJECT.md)、[工作流](docs/WORKFLOW.md) 与任务正文。保留用户和其他任务的未提交工作，核对远端、分支、基线 SHA；使用短期任务分支，目标默认分支为 `master`。

- 用户已经授权的常规实现和管理操作持续有效；不得从模板推导无关外发、费用或破坏性操作权限。
- 本轮仅涉及现有 GitHub 仓库；不创建新仓库，不引入 JEV 等本地项目依赖。
- 沿用现有 CI Go 1.26.x：`go build -o bin/carneades ./src/cmd/carneades`；`go test -vet=off ./src/cmd/carneades ./src/engine/caes/encoding/yaml`；`./bin/carneades check -f yaml examples/AGs/YAML/zh-CN/电子数据真实性.yml`。-vet=off 是现有局部测试约定，不代表全仓 vet 通过；图形输出可能另需 Graphviz。
- 行为改变覆盖原失败条件；文档改动检查事实、链接和差异。未知结果不得填“通过”。
- 保留上游历史、来源与 MPL-2.0 标识；中文翻译、示例和引擎改动分别记载；形式论证结果不冒充业务事实或司法结论。
- 任务状态用一个 status 标签，阻塞单列；PR 引用任务。合并、部署、用户验收分别记录，不自动关闭尚未满足交付验收的任务。
- 交接保留来源 SHA、命令/运行、输入输出范围、限制和下一动作。自查不称独立批准。
