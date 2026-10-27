# Carneades 4 简体中文本地化

本仓库是 [Carneades 4](https://github.com/carneades/carneades-4) 的非官方简体中文本地化分支。Carneades 是一个使用 Go 实现的计算论证系统，可以评估和可视化结构化论证图，也可以计算 Dung 抽象论证框架的扩展。

## 项目边界

- 本项目保留官方仓库的提交历史和 MPL-2.0 许可证。
- 中文界面、中文文档和中文示例由本分支独立编写。
- 本项目不隶属于 Carneades 原作者或 Fraunhofer FOKUS。
- 本项目是研究与教学工具，不提供自动裁判或法律意见。
- 软件标签和权重是形式模型的计算结果，不等同于中国法上的事实认定或证明标准。

## 已完成的中文化

- Web 操作界面及帮助页面；
- 中文快速开始文档；
- 法律论证核心术语表；
- 中文 YAML 论证图示例；
- 中文文本 YAML 导入导出测试；
- 来源、许可和原创性记录。

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

- [中文快速开始](../zh-CN/快速开始.md)
- [中英文术语表](../zh-CN/术语表.md)
- [来源、许可与原创性记录](../zh-CN/来源与许可.md)
- [汉化范围与未完成事项](../zh-CN/汉化范围.md)
- [两日交付清单](../zh-CN/两日交付清单.md)

## 中文示例

`examples/AGs/YAML/zh-CN/电子数据真实性.yml` 是独立编写的教学示例，用于演示支持论证、反对论证、争点和中文文本。它不对应真实案件，也不构成对电子数据真实性的法律判断。

## 上游与许可证

- 上游项目：<https://github.com/carneades/carneades-4>
- 原始许可证：[MPL-2.0](../../LICENSE)
- 本地化声明：[NOTICE.zh-CN.md](../NOTICE.zh-CN.md)

修改受 MPL-2.0 约束的源文件时，应继续遵守 MPL-2.0。独立中文研究材料如引用论文或法律文件，应使用规范引注，不得复制受版权保护的图表和大段表达。
