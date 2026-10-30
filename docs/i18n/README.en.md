<p>
  <img align="right" src="../assets/tony.gif" width="12%" alt="Tony 24 · Carneades · 中文适配" />
  <img src="../assets/hero.png" width="85%" alt="Carneades · 中文适配 · Make claims, reasons and rebuttals explicit—and inspect how they fit together." />
</p>

<br clear="all" />

[简体中文](../../README.md) · [繁體中文](README.zh-TW.md) · **English**

# Carneades · Chinese adaptation

**Make claims, reasons and rebuttals explicit—and inspect how they fit together.**

Carneades · Chinese adaptation is a formal argumentation toolkit for learning and research with Chinese material. It organizes claims, premises, issues and supporting or opposing arguments into a structured model, then checks and evaluates that model under explicit semantic rules. Start with a Chinese teaching example, repeat experiments from the command line, or launch a local web form in three languages. The project preserves the Carneades 4 engine and upstream provenance so that models and their reasoning can be understood, reproduced and discussed.

## Who it is for

- Students and researchers in argumentation theory, legal reasoning and explainable AI.
- Developers validating argument graphs, comparing semantics or integrating structured data.
- Educators who want a Chinese example for teaching and method demonstrations.

## 7 core capabilities

### 1. Structured argument models

Represent claims, premises and issues with statements, arguments, issues and assumptions. Put supporting and opposing arguments in the same model and assign a proof standard to an issue.

**Outcome：** A saved, editable argument graph that makes the premises and relationships in a discussion explicit.<br />
**Status：** Implemented · file-based modeling

### 2. Input checks and diagnostics

The check command validates syntactic and semantic constraints and reports problems in a model. It accepts a file or standard input, making it suitable for repeated checks while editing.

**Outcome：** Catch incomplete or inconsistent input before interpreting an evaluation result.<br />
**Status：** Implemented · CLI

### 3. Structured argument evaluation

After validation, eval applies any theory in the model to derive arguments, then uses grounded semantics to update statement acceptance labels.

**Outcome：** An evaluated graph for inspecting how conclusions change when assumptions change.<br />
**Status：** Implemented · CLI / local web

### 4. Abstract argumentation semantics

The dung command reads TGF frameworks and supports grounded, complete, preferred and stable semantics, including queries about credulous and skeptical acceptance.

**Outcome：** Inspect acceptable sets of arguments and compare how different semantics treat the same conflicts.<br />
**Status：** Implemented · CLI / local web

### 5. Interchange formats and graph output

Read structured graphs in YAML, CAF, AIF, AGXML or LKIF. The CLI exports YAML, DOT or GraphML; the web interface also offers PNG and SVG rendering.

**Outcome：** Keep machine-readable results or bring a visual argument map into another tool or a discussion.<br />
**Status：** Implemented · rendering requires Graphviz

### 6. Chinese teaching example

The electronic-data authenticity example shows how an original device, hash verification and editing traces can support or challenge a conclusion. It includes expected acceptance labels.

**Outcome：** A small, checkable model connecting Chinese text, argument structure and test expectations.<br />
**Status：** Available · fictional teaching model

### 7. Local interface in three languages

Launch the web service to upload a model and choose its input format, semantics and output format. The home page, forms and supporting instructions are available in Simplified Chinese, Traditional Chinese and English.

**Outcome：** Run demonstrations through forms and read the same functionality in your preferred language.<br />
**Status：** Implemented · self-hosted local service

## A typical journey

1. Read the Chinese teaching example and identify the claims, premises and issues you want to model.
2. Validate with check, then choose eval or dung for the type of evaluation you need.
3. Inspect text or graph output, revise assumptions or relationships, and record the differences between runs.

## Current delivery

- Delivered as source, Chinese documentation, a CLI and a locally runnable web interface. No public hosted service has been accepted in this delivery.
- The existing CI covers the command-line build, selected tests and checks of the Chinese example.
- An independent derivative project, not an official Chinese release by the Carneades authors or Fraunhofer FOKUS.

## Run locally

Run from the repository root. CI uses Go 1.26.x; install Graphviz for PNG/SVG rendering. Validate and evaluate the teaching example first. For the web interface, replace the template path with an absolute path on your machine.

```bash
go build -o bin/carneades ./src/cmd/carneades
./bin/carneades check -f yaml examples/AGs/YAML/zh-CN/电子数据真实性.yml
./bin/carneades eval -f yaml -t yaml examples/AGs/YAML/zh-CN/电子数据真实性.yml
./bin/carneades server -p 8080 -t "/absolute/path/to/src/web/templates" -l zh-CN
# http://localhost:8080/carneades/
```

## Scope and limits

- Formal evaluation reflects the supplied model and selected rules. It does not independently verify evidence or establish a judicial conclusion.
- The web interface supports only grounded semantics for Dung frameworks with more than 20 arguments. CLI and web output options differ.
- Selected tests use -vet=off; this does not establish that repository-wide vet passes. Interface translation does not translate input files. MPL-2.0 and upstream attribution remain applicable.

## Documentation and participation

[Usage guide](https://github.com/2432antony-13/carneades-4-zh-cn/blob/master/docs/USAGE.md) · [Chinese example](https://github.com/2432antony-13/carneades-4-zh-cn/tree/master/examples/AGs/YAML/zh-CN) · [Issues and feedback](https://github.com/2432antony-13/carneades-4-zh-cn/issues) · [Upstream attribution](https://github.com/2432antony-13/carneades-4-zh-cn/blob/master/docs/NOTICE.zh-CN.md)
