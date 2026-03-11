# 🎭 技能定义规范指南

## 📋 概述

本文档定义了 Go Agents 中技能（Skill）的标准化定义规范，包括提示词模板、配置结构和最佳实践。所有技能都应遵循此规范以确保一致性和可维护性。

## 🎯 **以终为始的设计理念**

### **终态目标**
技能的最终目标是：
1. **接收任务输入** - 能够理解并处理来自上层（Phase Lead、PO）的任务
2. **内部任务拆解** - 将任务分解为可执行的子任务序列
3. **执行任务** - 按照质量标准完成任务
4. **输出交付物** - 生成符合要求的输出物
5. **质量保证** - 通过质量门禁和自检机制

### **设计原则**
- **单一职责** - 每个技能专注于一个特定领域或角色
- **可组合性** - 技能可以组合成更复杂的工作流
- **可扩展性** - 支持变体（variants）和自定义配置
- **可测试性** - 技能行为可预测、可验证

## 🏗️ **技能文件结构**

### **标准文件结构**
```
~/.picoclaw/workspace/skills/
├── category-name/
│   ├── skill-name/
│   │   ├── SKILL.md              # 技能定义文件（必需）
│   │   ├── metadata.json         # 技能元数据（必需）
│   │   ├── variants/             # 变体定义（可选）
│   │   │   ├── variant-1.md
│   │   │   └── variant-2.md
│   │   ├── rules/                # 详细规则（可选）
│   │   │   ├── rule-1.md
│   │   │   └── rule-2.md
│   │   ├── templates/            # 提示词模板（必需）
│   │   │   ├── task-decomposition.md
│   │   │   ├── quality-check.md
│   │   │   └── deliverable-generation.md
│   │   ├── scripts/              # 执行脚本（可选）
│   │   │   ├── setup.sh
│   │   │   └── validate.sh
│   │   ├── examples/             # 使用示例（可选）
│   │   │   ├── basic-usage.md
│   │   │   └── advanced-usage.md
│   │   └── README.md             # 技能说明文档
│   └── README.md                # 类别说明文档
```

## 📝 **技能定义文件规范**

### **SKILL.md 文件模板**

```markdown
---
name: skill-name
description: "技能的简短描述，包含触发条件和使用场景"
category: "skill-category"
version: "1.0.0"
license: MIT

# 技能元数据
metadata:
  author: "技能作者"
  created_at: "2026-03-12"
  updated_at: "2026-03-12"
  tags: ["tag1", "tag2"]
  organization: "组织名称"
  references:
    - "相关文档链接1"
    - "相关文档链接2"

# 技能能力
capabilities:
  - "核心能力1"
  - "核心能力2"
  - "核心能力3"

# 所需依赖
requires:
  skills: ["dependency-skill-1", "dependency-skill-2"]
  tools: ["tool-1", "tool-2"]
  conflicts: ["conflicting-skill"]

# 触发条件
triggers:
  - "关键词1"
  - "关键词2"
  - "特定任务描述"

# 技能变体
variants:
  variant-1:
    persona: |
      我是[角色名称]，专注于[专业领域]。
      我的特长是[核心技能]，能够[具体能力]。
      我的优势是[独特优势]，擅长[特殊技能]。
    
    capabilities: ["variant-specific-capability-1", "variant-specific-capability-2"]
    tools: ["variant-specific-tool-1", "variant-specific-tool-2"]
    
    execution:
      task_breakdown:
        strategy: "template_based|milestone_driven|hybrid"
        granularity: "fine|medium|coarse"
        max_depth: 3
        auto_adjust: true
      
      subtask_execution:
        sequential: true
        parallel_threshold: 4
        quality_gates:
          - gate: "quality_gate_name"
            threshold: 85
            reviewer: "reviewer_role"
      
      ralph_wiggum_loop:
        enabled: true
        steps: ["step1", "step2", "step3"]
        auto_update: true

# 技能执行配置
execution:
  # 任务分解策略
  task_breakdown:
    strategy: "template_based"
    granularity: "medium"
    max_depth: 3
    auto_adjust: true
  
  # 子任务执行
  subtask_execution:
    sequential: true
    parallel_threshold: 4
    quality_gates:
      - gate: "task_completion"
        threshold: 100
        reviewer: "self"
      - gate: "code_quality"
        threshold: 85
        reviewer: "self"
  
  # Ralph Wiggum Loop
  ralph_wiggum_loop:
    enabled: true
    steps:
      - "run_tests_and_lint"
      - "self_review_vs_harness"
      - "update_wbs"
      - "suggest_next_tasks"
    
    quality_gates:
      - gate: "test_coverage"
        threshold: 85
      - gate: "code_quality"
        threshold: 80
    
    auto_update:
      wbs: true
      harness: false

# 输出物定义
outputs:
  supported_types:
    - "document"
    - "code"
    - "image"
    - "video"
    - "prototype"
    - "demo"
  
  quality_standards:
    document:
      format: "markdown"
      structure: "standard"
      validation: "markdown_lint"
    code:
      format: "source_code"
      coverage: ">=85%"
      style: "standard_lint"

# 质量门禁定义
quality_gates:
  task_level:
    - gate: "task_completion"
      threshold: 100
      reviewer: "self"
    - gate: "deliverable_quality"
      threshold: 85
      reviewer: "phase_lead"
  
  milestone_level:
    - gate: "milestone_success"
      threshold: 80
      reviewer: "po"

# 使用示例
examples:
  - title: "基本使用示例"
    description: "技能的基本使用方法"
    command: "@go \"[任务描述]\""
    expected_output: "[预期输出]"
  
  - title: "高级使用示例"
    description: "技能的高级使用方法"
    command: "@go --mode advanced \"[复杂任务描述]\""
    expected_output: "[预期输出]"

# 如何使用
usage:
  when_to_apply:
    - "使用场景1"
    - "使用场景2"
    - "使用场景3"
  
  command_examples:
    - command: "@go \"[示例命令]\""
      description: "命令说明"
      expected: "预期结果"

# 输出格式
output_format:
  success: "成功时的输出格式模板"
  error: "错误时的输出格式模板"
  progress: "进度更新格式模板"

# 故障排除
troubleshooting:
  common_issues:
    - issue: "常见问题1"
      solution: "解决方案1"
    - issue: "常见问题2"
      solution: "解决方案2"
```

### **metadata.json 文件模板**

```json
{
  "version": "1.0.0",
  "organization": "Go Agents Team",
  "date": "2026-03-12",
  "abstract": "技能的简短摘要，描述技能的主要功能和用途",
  "references": [
    "https://相关文档链接1",
    "https://相关文档链接2"
  ],
  "tags": ["tag1", "tag2", "tag3"],
  "category": "skill-category",
  "complexity": "basic|intermediate|advanced",
  "estimated_time": "1-2小时",
  "dependencies": ["skill-1", "skill-2"],
  "compatibility": {
    "picoclaw_version": ">=1.0.0",
    "platforms": ["web", "desktop", "mobile"]
  }
}
```

## 📋 **规则文件规范**

### **规则文件模板**

```markdown
---
title: 规则标题
impact: MEDIUM|HIGH|CRITICAL
impactDescription: "影响描述（如：20-50% 改进）"
tags: tag1, tag2, tag3
category: "规则分类"
---

## 规则标题

**Impact: MEDIUM (影响描述)**

规则的简要说明和重要性解释。应该清晰简洁，解释性能或质量影响。

**错误（问题描述）：**

```typescript
// 错误的代码示例
const bad = example()
```

**正确（解决方案）：**

```typescript
// 正确的代码示例
const good = example()
```

**解释：**
详细说明为什么正确的方法更好，以及具体的改进效果。

**参考：**[相关文档链接](https://example.com)

**相关规则：**
- [相关规则1](rule-1.md)
- [相关规则2](rule-2.md)
```

## 📋 **提示词模板规范**

### **1. 任务分解模板**

#### **templates/task-decomposition.md**
```markdown
# 任务分解提示词模板

## 模板用途
用于指导技能内部的任务分解过程，确保任务拆解的合理性和完整性。

## 基础模板
```
First, read and strictly follow HARNESS.md and WBS.md in the repository root.

当前任务：【任务描述】

请按以下结构进行任务拆解：
- 每个子任务规模控制在 350-750 行代码（含测试）
- 给出 Acceptance Criteria + 验证命令 + 规模预估
- 确保 100% 规则覆盖完整
- 用 Markdown 表格形式输出子任务清单

## 拆解要求
1. **粒度控制**：每个子任务 300-800 行，影响 3-12 个文件
2. **依赖关系**：明确子任务间的依赖关系
3. **质量标准**：每个子任务都有明确的质量门禁
4. **可交付性**：每个子任务都有可验证的输出物

## 输出格式
```markdown
| 子任务ID | 任务名称 | 预估时间 | 依赖关系 | 质量门禁 | 输出物 |
|---------|---------|---------|---------|---------|--------|
| subtask_1 | 数据收集准备 | 4小时 | - | 数据完整性90% | data_sources_list |
| subtask_2 | 市场分析 | 1-2天 | subtask_1 | 分析深度80% | analysis_report |
```

## 特殊处理
- **复杂任务**：优先使用 milestone_driven 策略
- **探索性任务**：允许动态调整粒度
- **标准化任务**：使用 template_based 策略
```

### **2. 质量检查模板**

#### **templates/quality-check.md**
```markdown
# 质量检查提示词模板

## 模板用途
用于指导技能执行过程中的质量检查和自检机制。

## 基础模板
```
请按照以下标准进行代码质量检查：

## 代码质量检查
1. **代码规范**：
   - 遵循 HARNESS.md 中的编码规范
   - 变量命名符合项目标准
   - 函数长度控制在合理范围
   - 注释覆盖率 >= 80%

2. **功能正确性**：
   - 所有功能点都已实现
   - 边界条件处理正确
   - 错误处理完善
   - 单元测试覆盖核心逻辑

3. **性能要求**：
   - 响应时间符合要求
   - 内存使用合理
   - 无明显性能瓶颈
   - 资源使用优化

## 测试要求
1. **单元测试**：
   - 测试覆盖率 >= 85%
   - 测试用例覆盖边界条件
   - 测试命名规范清晰
   - 测试数据准备充分

2. **集成测试**：
   - 关键路径测试通过
   - 接口调用正确
   - 数据一致性验证
   - 错误处理测试完整

## 检查命令
```bash
# 运行测试
npm test

# 代码检查
npm run lint

# 类型检查
npm run type-check

# 代码覆盖率
npm run coverage
```

## 输出要求
请提供：
1. 质量检查结果摘要
2. 发现的问题列表
3. 修复建议
4. 修复后的验证计划
```

### **3. 交付物生成模板**

#### **templates/deliverable-generation.md**
```markdown
# 交付物生成提示词模板

## 模板用途
用于指导技能生成符合要求的交付物。

## 基础模板
```
请按照以下标准生成交付物：

## 交付物要求
1. **内容完整性**：
   - 包含所有必需章节
   - 内容逻辑清晰
   - 结构层次合理
   - 格式符合规范

2. **质量标准**：
   - 语言表达准确
   - 专业术语正确
   - 无语法错误
   - 格式统一规范

3. **可读性**：
   - 章节长度适中
   - 重点内容突出
   - 图表清晰易懂
   - 示例丰富贴切

## 输出格式
- **文档类型**：Markdown 格式
- **代码类型**：符合语言规范
- **图片类型**：PNG 格式，清晰度 >= 300dpi
- **视频类型**：MP4 格式，时长适中

## 质量检查
生成完成后请检查：
1. 内容完整性验证
2. 格式规范检查
3. 质量门禁通过
4. 用户需求满足度
```

## 🎭 **变体定义规范**

### **变体文件结构**
```
variants/
├── variant-1.md
├── variant-2.md
└── variant-3.md
```

### **变体定义模板**
```markdown
---
name: variant-name
description: "变体描述"
parent_skill: "parent-skill-name"
persona_type: "specialist|generalist|expert"

# 角色定义
persona: |
  我是[角色名称]，专注于[专业领域]。
  我的特长是[核心技能]，能够[具体能力]。
  我的优势是[独特优势]，擅长[特殊技能]。

# 专业能力
capabilities:
  - "专业能力1"
  - "专业能力2"
  - "专业能力3"

# 工具集
tools:
  - "工具1"
  - "工具2"
  - "工具3"

# 执行配置
execution:
  task_breakdown:
    strategy: "template_based|milestone_driven|hybrid"
    granularity: "fine|medium|coarse"
    max_depth: 3
    auto_adjust: true
  
  subtask_execution:
    sequential: true
    parallel_threshold: 4
    quality_gates:
      - gate: "quality_gate_name"
        threshold: 85
        reviewer: "reviewer_role"
  
  ralph_wiggum_loop:
    enabled: true
    steps: ["step1", "step2", "step3"]
    auto_update: true

# 特殊配置
special_config:
  custom_setting_1: "value1"
  custom_setting_2: "value2"

# 使用场景
use_cases:
  - "适用场景1"
  - "适用场景2"
  - "适用场景3"
```

## 🔧 **脚本规范**

### **脚本文件结构**
```
scripts/
├── setup.sh              # 环境设置脚本
├── validate.sh           # 验证脚本
├── test.sh              # 测试脚本
└── cleanup.sh           # 清理脚本
```

### **脚本模板**
```bash
#!/bin/bash
# 脚本描述
set -e  # 失败时退出

# 全局变量
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SKILL_DIR="$(dirname "$SCRIPT_DIR")"

# 日志函数
log_info() {
    echo "INFO: $1" >&2
}

log_error() {
    echo "ERROR: $1" >&2
}

# 主要逻辑
main() {
    log_info "开始执行脚本..."
    
    # 脚本逻辑
    
    log_info "脚本执行完成"
}

# 清理函数
cleanup() {
    log_info "清理资源..."
    # 清理逻辑
}

# 设置清理陷阱
trap cleanup EXIT

# 执行主函数
main "$@"
```

## 📊 **质量标准**

### **代码质量标准**
- **测试覆盖率**: >= 85%
- **代码规范**: 100% 通过 lint 检查
- **类型检查**: 100% 通过类型检查
- **文档覆盖率**: >= 80%

### **交付物质量标准**
- **内容完整性**: 100% 覆盖要求
- **格式规范**: 100% 符合标准
- **专业准确性**: 95% 以上准确
- **用户满意度**: >= 4.5/5

### **性能标准**
- **响应时间**: < 5秒（简单任务）
- **执行效率**: > 80% 成功率
- **资源使用**: 合理的内存和 CPU 使用
- **并发处理**: 支持合理的并发任务

## 🚀 **最佳实践**

### **1. 技能设计原则**
- **单一职责**: 每个技能专注于一个特定领域
- **可组合性**: 技能可以组合成更复杂的工作流
- **可扩展性**: 支持变体和自定义配置
- **可测试性**: 技能行为可预测、可验证

### **2. 提示词设计原则**
- **清晰明确**: 提示词要具体、无歧义
- **结构化**: 使用标准化的模板结构
- **可复用**: 支持模板化使用
- **可维护**: 易于理解和修改

### **3. 质量保证原则**
- **多层次检查**: 子任务、任务、里程碑多级检查
- **自动化**: 尽可能使用自动化工具
- **一致性**: 保持与项目标准一致
- **可追溯**: 所有检查结果可追溯

### **4. 错误处理原则**
- **优雅降级**: 遇到错误时优雅处理
- **详细日志**: 提供详细的错误日志
- **恢复机制**: 支持错误恢复和重试
- **用户友好**: 提供用户友好的错误信息

## 📝 **文档要求**

### **README.md**
```markdown
# 技能名称

## 概述
技能的简要描述和主要功能

## 安装和配置
详细的安装步骤和配置说明

## 基本使用
基本的使用示例和命令

## 高级功能
高级功能的说明和使用方法

## 配置选项
详细的配置选项说明

## 故障排除
常见问题和解决方案

## 贡献指南
如何贡献代码和文档
```

### **CHANGELOG.md**
```markdown
# 变更日志

## [1.0.0] - 2026-03-12

### 新增
- 新功能1
- 新功能2

### 修复
- 修复问题1
- 修复问题2

### 变更
- 变更1
- 变更2
```

## 🔍 **验证和测试**

### **技能验证清单**
- [ ] 技能定义符合规范
- [ ] 提示词模板完整
- [ ] 质量门禁配置正确
- [ ] 输出物类型支持完整
- [ ] 错误处理机制完善
- [ ] 脚本可正常执行
- [ ] 文档完整准确

### **功能测试清单**
- [ ] 基本功能测试通过
- [ ] 边界条件测试通过
- [ ] 错误处理测试通过
- [ ] 性能测试通过
- [ ] 集成测试通过
- [ ] 变体功能测试通过

### **文档验证清单**
- [ ] 文档结构符合规范
- [ ] 内容准确完整
- [ ] 示例可执行
- [ ] 链接正确有效
- [ ] 格式规范统一

## 🎉 **总结**

遵循此规范定义的技能将具备：
1. **一致性** - 所有技能遵循统一的定义标准
2. **可维护性** - 清晰的结构和文档
3. **可扩展性** - 支持变体和自定义配置
4. **高质量** - 内置质量保证机制
5. **易用性** - 标准化的使用方式

通过统一的技能定义规范，我们可以构建一个高质量、可维护、可扩展的技能生态系统！🚀

## 📋 **提示词模板规范**

### **1. 任务分解模板**

#### **templates/task-decomposition.md**
```markdown
# 任务分解提示词模板

## 模板用途
用于指导技能内部的任务分解过程，确保任务拆解的合理性和完整性。

## 基础模板
```
First, read and strictly follow HARNESS.md and WBS.md in the repository root.

当前任务：【任务描述】

请按以下结构进行任务拆解：
- 每个子任务规模控制在 350-750 行代码（含测试）
- 给出 Acceptance Criteria + 验证命令 + 规模预估
- 确保 100% 规则覆盖完整
- 用 Markdown 表格形式输出子任务清单

## 拆解要求
1. **粒度控制**：每个子任务 300-800 行，影响 3-12 个文件
2. **依赖关系**：明确子任务间的依赖关系
3. **质量标准**：每个子任务都有明确的质量门禁
4. **可交付性**：每个子任务都有可验证的输出物

## 输出格式
```markdown
| 子任务ID | 任务名称 | 预估时间 | 依赖关系 | 质量门禁 | 输出物 |
|---------|---------|---------|---------|---------|--------|
| subtask_1 | 数据收集准备 | 4小时 | - | 数据完整性90% | data_sources_list |
| subtask_2 | 市场分析 | 1-2天 | subtask_1 | 分析深度80% | analysis_report |
```

## 特殊处理
- **复杂任务**：优先使用 milestone_driven 策略
- **探索性任务**：允许动态调整粒度
- **标准化任务**：使用 template_based 策略
```

### **2. 质量检查模板**

#### **templates/quality-check.md**
```markdown
# 质量检查提示词模板

## 模板用途
用于指导技能执行过程中的质量检查和自检机制。

## 基础模板
```
请按照以下标准进行代码质量检查：

## 代码质量检查
1. **代码规范**：
   - 遵循 HARNESS.md 中的编码规范
   - 变量命名符合项目标准
   - 函数长度控制在合理范围
   - 注释覆盖率 >= 80%

2. **功能正确性**：
   - 所有功能点都已实现
   - 边界条件处理正确
   - 错误处理完善
   - 单元测试覆盖核心逻辑

3. **性能要求**：
   - 响应时间符合要求
   - 内存使用合理
   - 无明显性能瓶颈
   - 资源使用优化

## 测试要求
1. **单元测试**：
   - 测试覆盖率 >= 85%
   - 测试用例覆盖边界条件
   - 测试命名规范清晰
   - 测试数据准备充分

2. **集成测试**：
   - 关键路径测试通过
   - 接口调用正确
   - 数据一致性验证
   - 错误处理测试完整

## 检查命令
```bash
# 运行测试
npm test

# 代码检查
npm run lint

# 类型检查
npm run type-check

# 代码覆盖率
npm run coverage
```

## 输出要求
请提供：
1. 质量检查结果摘要
2. 发现的问题列表
3. 修复建议
4. 修复后的验证计划
```

### **3. 交付物生成模板**

#### **templates/deliverable-generation.md**
```
# 交付物生成提示词模板

## 模板用途
用于指导技能生成符合要求的交付物。

## 基础模板
```
请按照以下标准生成交付物：

## 交付物要求
1. **内容完整性**：
   - 包含所有必需章节
   - 内容逻辑清晰
   - 结构层次合理
   - 格式符合规范

2. **质量标准**：
   - 语言表达准确
   - 专业术语正确
   - 无语法错误
   - 格式统一规范

3. **可读性**：
   - 章节长度适中
   - 重点内容突出
   - 图表清晰易懂
   - 示例丰富贴切

## 输出格式
- **文档类型**：Markdown 格式
- **代码类型**：符合语言规范
- **图片类型**：PNG 格式，清晰度 >= 300dpi
- **视频类型**：MP4 格式，时长适中

## 质量检查
生成完成后请检查：
1. 内容完整性验证
2. 格式规范检查
3. 质量门禁通过
4. 用户需求满足度
```

## 🎭 **变体定义规范**

### **变体文件结构**
```
variants/
├── variant-1.md
├── variant-2.md
└── variant-3.md
```

### **变体定义模板**
```markdown
---
name: variant-name
description: "变体描述"
parent_skill: "parent-skill-name"

# 角色定义
persona: |
  我是[角色名称]，专注于[专业领域]。
  我的特长是[核心技能]，能够[具体能力]。
  我的优势是[独特优势]，擅长[特殊技能]。

# 专业能力
capabilities:
  - "专业能力1"
  - "专业能力2"
  - "专业能力3"

# 工具集
tools:
  - "工具1"
  - "工具2"
  - "工具3"

# 执行配置
execution:
  task_breakdown:
    strategy: "template_based|milestone_driven|hybrid"
    granularity: "fine|medium|coarse"
    max_depth: 3
    auto_adjust: true
  
  subtask_execution:
    sequential: true
    parallel_threshold: 4
    quality_gates:
      - gate: "quality_gate_name"
        threshold: 85
        reviewer: "reviewer_role"
  
  ralph_wiggum_loop:
    enabled: true
    steps: ["step1", "step2", "step3"]
    auto_update: true

# 特殊配置
special_config:
  custom_setting_1: "value1"
  custom_setting_2: "value2"
```

## 🔧 **配置规范**

### **全局配置映射**
```yaml
# 技能配置与全局配置的映射关系
skill_config:
  # 任务分解策略
  task_breakdown:
    strategy: "template_based|milestone_driven|hybrid"
    granularity: "fine|medium|coarse"
    max_depth: 3
    auto_adjust: true
  
  # 质量控制
  quality_control:
    default_threshold: 80
    mandatory_gates: ["code_quality", "test_coverage"]
    escalation_path: ["self", "phase_lead", "po"]
  
  # 执行模式
  execution_mode:
    ralph_wiggum_loop: true
    auto_update: true
    quality_check_frequency: "per_subtask"
```

### **环境变量支持**
```yaml
# 环境变量配置
environment:
  # 调试模式
  debug_mode: ${DEBUG_MODE:false}
  
  # 日志级别
  log_level: ${LOG_LEVEL:info}
  
  # 超时设置
  timeout: ${TASK_TIMEOUT:3600}
  
  # 并发限制
  max_concurrent_tasks: ${MAX_CONCURRENT:3}
```

## 📊 **质量标准**

### **代码质量标准**
- **测试覆盖率**: >= 85%
- **代码规范**: 100% 通过 lint 检查
- **类型检查**: 100% 通过类型检查
- **文档覆盖率**: >= 80%

### **交付物质量标准**
- **内容完整性**: 100% 覆盖要求
- **格式规范**: 100% 符合标准
- **专业准确性**: 95% 以上准确
- **用户满意度**: >= 4.5/5

### **性能标准**
- **响应时间**: < 5秒（简单任务）
- **执行效率**: > 80% 成功率
- **资源使用**: 合理的内存和 CPU 使用
- **并发处理**: 支持合理的并发任务

## 🚀 **最佳实践**

### **1. 技能设计原则**
- **单一职责**: 每个技能专注于一个特定领域
- **可组合性**: 技能可以组合成更复杂的工作流
- **可扩展性**: 支持变体和自定义配置
- **可测试性**: 技能行为可预测、可验证

### **2. 提示词设计原则**
- **清晰明确**: 提示词要具体、无歧义
- **结构化**: 使用标准化的模板结构
- **可复用**: 支持模板化使用
- **可维护**: 易于理解和修改

### **3. 质量保证原则**
- **多层次检查**: 子任务、任务、里程碑多级检查
- **自动化**: 尽可能使用自动化工具
- **一致性**: 保持与项目标准一致
- **可追溯**: 所有检查结果可追溯

### **4. 错误处理原则**
- **优雅降级**: 遇到错误时优雅处理
- **详细日志**: 提供详细的错误日志
- **恢复机制**: 支持错误恢复和重试
- **用户友好**: 提供用户友好的错误信息

## 📝 **文档要求**

### **README.md**
- 技能概述和用途
- 安装和配置说明
- 基本使用示例
- 高级功能说明
- 故障排除指南

### **CHANGELOG.md**
- 版本变更记录
- 新功能说明
- 修复问题列表
- 兼容性说明

### **examples/** 目录
- 基本使用示例
- 高级使用示例
- 最佳实践案例
- 常见问题解答

## 🔍 **验证和测试**

### **技能验证清单**
- [ ] 技能定义符合规范
- [ ] 提示词模板完整
- [ ] 质量门禁配置正确
- [ ] 输出物类型支持完整
- [ ] 错误处理机制完善

### **功能测试清单**
- [ ] 基本功能测试通过
- [ ] 边界条件测试通过
- [ ] 错误处理测试通过
- [ ] 性能测试通过
- [ ] 集成测试通过

### **文档验证清单**
- [ ] 文档结构符合规范
- [ ] 内容准确完整
- [ ] 示例可执行
- [ ] 链接正确有效
- [ ] 格式规范统一

## 🎉 **总结**

遵循此规范定义的技能将具备：
1. **一致性** - 所有技能遵循统一的定义标准
2. **可维护性** - 清晰的结构和文档
3. **可扩展性** - 支持变体和自定义配置
4. **高质量** - 内置质量保证机制
5. **易用性** - 标准化的使用方式

通过统一的技能定义规范，我们可以构建一个高质量、可维护、可扩展的技能生态系统！🚀
