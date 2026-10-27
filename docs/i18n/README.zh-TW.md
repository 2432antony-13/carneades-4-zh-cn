<p><img src="../assets/tony.gif" width="112" height="112" alt="Tony 24 · Carneades"></p>

[简体中文](../../README.md) · **繁體中文** · [English](../../docs/i18n/README.en.md)

# Carneades · 中文適配

**讓論點、依據與反駁可以被檢驗。**

Carneades 4 的獨立中文適配分支，保留形式論證引擎、上游歷史與可追溯的中文範例。

[任務](https://github.com/2432antony-13/carneades-4-zh-cn/issues) · [文件](../../docs/README.md)

## 目前狀態

命令列與中文範例已接入 CI。這是獨立衍生專案，不是原作者或 Fraunhofer FOKUS 的官方中文發行。

## 快速開始

CI 使用 Go 1.26.x。圖形匯出可能需要 Graphviz；局部測試的 -vet=off 不代表全倉 vet 通過。

```bash
go build -o bin/carneades ./src/cmd/carneades
go test -vet=off ./src/cmd/carneades ./src/engine/caes/encoding/yaml
./bin/carneades check -f yaml examples/AGs/YAML/zh-CN/电子数据真实性.yml
```

## 能力與範圍

抽象論證語義；結構化論證評估；YAML 範例檢查；圖形匯出。

保留 MPL-2.0 及上游歸屬。形式推理結果不等於業務事實或司法結論。

[原始碼](../../src/) · [文件](../../docs/README.md)

倉庫首頁與產品介面提供 **简体中文 / 繁體中文 / English**。採用 Q 版 **Tony 24** 標識，GitHub 顯示 GIF，產品使用 SVG/CSS 動畫。

[語言覆蓋與原文範圍](../LOCALIZATION.md)
