# carneades-4-zh-cn 治理交接

2026-10-27；当前授权：用户要求按既定规范更新五个已有 GitHub 项目。工作分支 governance/project-baseline-20261027，基线 `630d60403e81ee239d8babb81c2144baa27c0a4f`。

已准备章程、协作规则、质量矩阵、资产来源与模板；实际 Issue、PR、平台约束和检查结果由本轮完成记录补充。当前尚未以本文件宣告合并、部署或业务验收。

下一执行动作：读取本仓 Issues 的唯一状态与阻塞，核对当前分支和最新检查，按具体用例推进。沿用现有 CI Go 1.26.x：`go build -o bin/carneades ./src/cmd/carneades`；`go test -vet=off ./src/cmd/carneades ./src/engine/caes/encoding/yaml`；`./bin/carneades check -f yaml examples/AGs/YAML/zh-CN/电子数据真实性.yml`。-vet=off 是现有局部测试约定，不代表全仓 vet 通过；图形输出可能另需 Graphviz。

具体产品待办：建立中文论证案例与当前版本检查基线；建立上游同步与中文版本交付检查。业务期限/预算尚未承诺；新外部费用、真实发信或数据恢复按任务补齐范围。

当前已知限制：默认分支 master；最近观察到的成功运行属于较早提交，当前版本需重新触发 PR 检查。既有中文 YAML 往返断言可用作局部基线。 2026-10-27 已补齐 project/workflow 权限并接入私有看板；独立人工批准和用户验收尚未取得。本轮不新建仓库，不引入 JEV 等本地对象。

## 实际任务入口

[本轮里程碑](https://github.com/2432antony-13/carneades-4-zh-cn/milestone/1)

- [governance #2](https://github.com/2432antony-13/carneades-4-zh-cn/issues/2)
- [quality #3](https://github.com/2432antony-13/carneades-4-zh-cn/issues/3)
- [release #4](https://github.com/2432antony-13/carneades-4-zh-cn/issues/4)

任务当前状态以[私有管理看板](https://github.com/users/2432antony-13/projects/1)的 Status 与 Blocked 为准；最终 PR、合并来源与运行结果在治理任务正文登记。这里不复制会随进度变化的第二套任务状态。

## 2026-10-27 首页整理与三语入口

- 本轮首页任务：https://github.com/2432antony-13/carneades-4-zh-cn/issues/7。
- README 顶部提供简体中文、繁体中文和 English；完整说明收进 docs/，旧版本在 docs/archive/ 保留。
- 三语覆盖当前仓库首页与核心使用说明；产品界面与专业内容尚未完成全量三语化，不能据此标为已完成。
