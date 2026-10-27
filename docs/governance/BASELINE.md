# carneades-4-zh-cn 治理基线

观察日期：2026-10-27。基线：`630d60403e81ee239d8babb81c2144baa27c0a4f`；默认分支 `master`；仓库公开。这是冻结事实，不替代当前 PR 的检查。

默认分支 master；最近观察到的成功运行属于较早提交，当前版本需重新触发 PR 检查。既有中文 YAML 往返断言可用作局部基线。

## 检查入口与历史证据

沿用现有 CI Go 1.26.x：`go build -o bin/carneades ./src/cmd/carneades`；`go test -vet=off ./src/cmd/carneades ./src/engine/caes/encoding/yaml`；`./bin/carneades check -f yaml examples/AGs/YAML/zh-CN/电子数据真实性.yml`。-vet=off 是现有局部测试约定，不代表全仓 vet 通过；图形输出可能另需 Graphviz。

- [zh-CN foundation / 30254949078](https://github.com/2432antony-13/carneades-4-zh-cn/actions/runs/30254949078)：success，来源 `231c5487d477f1bf85c24744dee5e46ff42e24f9`。
- [zh-CN foundation / 30252959051](https://github.com/2432antony-13/carneades-4-zh-cn/actions/runs/30252959051)：success，来源 `df7b7a26cb1b206a40fed77806928147277e7467`。
- [zh-CN foundation / 30251723134](https://github.com/2432antony-13/carneades-4-zh-cn/actions/runs/30251723134)：success，来源 `e21334d82196235d65e2d55b3dfdafae4534240a`。

## 平台差异

公开仓库尚无 rulesets；待实际 PR 检查通过后配置并回读。 CLI 当前有 repo 权限，仍缺 project/workflow；Projects 和新工作流接入单列待办。Carneades 的 Issues 初始关闭，本轮为管理入口启用并核对。

## 本轮采用范围

采用材料来源：用户提供的 project-governance kit/skill v1.0.0 与 2026-10-25 工作区方案；用户 2026-10-27 授权五个现有仓库实施。此前跨仓复用是建议；未发生的集成不登记为已采用。

本轮 PR 与实际管理交付见 [HANDOVER.md](HANDOVER.md)。专业质量、浏览器使用和线上渠道必须另有证据。

## 上游管理

父仓为 carneades/carneades-4；本地 origin 指向本 fork。同步前固定上游提交与本地版本，分别检查术语/模板、中文示例和引擎变化，运行局部回归后经 PR 集成。本轮没有同步上游或建立新的源码远端。

## 平台接入更新（2026-10-27）

上文是治理启动时快照。当前账号已具备 project/workflow，任务状态统一在[私有管理看板](https://github.com/users/2432antony-13/projects/1)的 Status 字段维护，阻塞使用 Blocked / Blocker 字段；Issue 正文保留验收与解除条件。旧 status/blocked 标签停止作为状态源。 本轮 PR 及运行结果在关联任务中记录，私有仓套餐限制与产品验收继续单列。
