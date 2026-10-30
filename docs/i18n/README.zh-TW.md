<p>
  <img align="right" src="../assets/tony.gif" width="12%" alt="Tony 24 · Carneades · 中文适配" />
  <img src="../assets/hero.png" width="85%" alt="Carneades · 中文适配 · 把論點、依據與反駁，放到同一張可檢驗的論證圖中。" />
</p>

<br clear="all" />

[简体中文](../../README.md) · **繁體中文** · [English](README.en.md)

# Carneades · 中文适配

**把論點、依據與反駁，放到同一張可檢驗的論證圖中。**

Carneades · 中文適配是一套面向中文學習與研究場景的形式論證工具。它把觀點、前提、爭點和支援或反駁關係組織為結構化模型，再用明確的語義規則檢查和評估。你可以從中文教學示例開始，用命令列重複實驗，也可以啟動三語網頁表單檢視結果。專案保留 Carneades 4 上游引擎與來源記錄，重點是讓建模過程和推理依據能夠被理解、復現與討論。

## 適合誰

- 論證理論、法律推理與可解釋 AI 的學習者和研究者。
- 需要驗證論證圖、比較不同語義或接入結構化資料的開發者。
- 希望以中文示例開展課堂演示和方法討論的教師。

## 7 個核心模組

### 1. 結構化論證建模

用 statements、arguments、issues 和 assumptions 描述觀點、前提與爭點，將支援和反駁關係放入同一個模型。爭點可以指定證明標準。

**結果：** 得到可儲存、可修改的論證圖輸入，讓討論落到具體前提和關係上。<br />
**狀態：** 已實現 · 檔案建模

### 2. 輸入校驗與錯誤定位

check 命令先檢查論證圖的語法和語義約束，再報告模型中的問題。可從檔案或標準輸入讀入，適合編輯後反覆檢查。

**結果：** 在評估前發現不完整或不一致的模型，避免把輸入錯誤誤認為推理結果。<br />
**狀態：** 已實現 · CLI

### 3. 結構化論證評估

eval 命令在校驗後應用模型中的理論，推導相應論證，再使用 grounded 語義更新陳述的接受狀態。

**結果：** 得到帶有評估結果的論證圖，便於比較不同假設下結論如何變化。<br />
**狀態：** 已實現 · CLI / 本地網頁

### 4. 抽象論證語義計算

dung 命令讀取 TGF 論證框架，支援 grounded、complete、preferred 和 stable 四種語義，並可查詢論證的可信接受或懷疑接受情況。

**結果：** 檢視可接受的論證集合，比較不同語義對同一衝突結構給出的結果。<br />
**狀態：** 已實現 · CLI / 本地網頁

### 5. 格式互通與圖形匯出

結構化輸入支援 YAML、CAF、AIF、AGXML 和 LKIF；CLI 可輸出 YAML、DOT、GraphML，網頁另提供 PNG 和 SVG 圖形輸出。

**結果：** 既能儲存機器可讀結果，也能把論證關係帶入圖形工具或討論材料。<br />
**狀態：** 已實現 · 圖形渲染需 Graphviz

### 6. 中文教學示例

倉庫提供“電子資料真實性”YAML 示例，展示原始裝置、雜湊核驗與編輯痕跡如何構成支援和反駁，並寫有預期接受狀態。

**結果：** 從一個可檢查的小模型入手，理解中文文字、論證結構和測試斷言如何對應。<br />
**狀態：** 已提供 · 虛構教學模型

### 7. 三語本地操作介面

啟動本地 Web 服務後，可透過表單上傳模型、選擇輸入格式、語義和輸出格式；首頁、操作說明等頁面支援簡體中文、繁體中文與 English。

**結果：** 不必為每次演示重輸完整命令，也能在不同語言下閱讀相同功能說明。<br />
**狀態：** 已實現 · 自行執行服務

## 如何使用

1. 先閱讀中文教學示例，確認要表達的觀點、前提與爭點。
2. 執行 check 校驗模型，再用 eval 或 dung 選擇對應的評估方式。
3. 檢視文字或圖形結果，修改假設與關係後再次執行，記錄輸入和結果差異。

## 目前完成程度

- 當前交付為原始碼、中文資料、CLI 和可自行啟動的三語 Web 介面；沒有本輪已驗收的公共線上服務。
- 命令列構建、區域性測試與中文示例檢查已納入現有 CI，支援持續檢查這條基礎使用路徑。
- 這是獨立衍生專案，不是 Carneades 原作者或 Fraunhofer FOKUS 的官方中文發行。

## 本機開始

在倉庫根目錄執行。CI 使用 Go 1.26.x；需要 PNG/SVG 渲染時安裝 Graphviz。下列先檢查教學示例，再匯出評估結果；如需網頁，將模板目錄替換為本機絕對路徑。

```bash
go build -o bin/carneades ./src/cmd/carneades
./bin/carneades check -f yaml examples/AGs/YAML/zh-CN/电子数据真实性.yml
./bin/carneades eval -f yaml -t yaml examples/AGs/YAML/zh-CN/电子数据真实性.yml
./bin/carneades server -p 8080 -t "/absolute/path/to/src/web/templates" -l zh-CN
# http://localhost:8080/carneades/
```

## 範圍與邊界

- 形式評估只反映輸入模型及選定規則，不自動驗證證據真偽，也不構成司法結論。
- Web 端對超過 20 個論證的 Dung 框架僅支援 grounded 語義；CLI 與 Web 的可選輸出並不完全相同。
- 區域性測試使用 -vet=off，不能據此聲稱全倉 vet 透過。介面翻譯不自動翻譯輸入檔案；繼續遵循 MPL-2.0 及上游歸屬要求。

## 文件與參與

[使用說明](https://github.com/2432antony-13/carneades-4-zh-cn/blob/master/docs/USAGE.md) · [中文示例](https://github.com/2432antony-13/carneades-4-zh-cn/tree/master/examples/AGs/YAML/zh-CN) · [任務與反饋](https://github.com/2432antony-13/carneades-4-zh-cn/issues) · [上游與許可宣告](https://github.com/2432antony-13/carneades-4-zh-cn/blob/master/docs/NOTICE.zh-CN.md)
