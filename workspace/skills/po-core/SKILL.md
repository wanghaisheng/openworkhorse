---
name: po-core
description: "产品经理核心 - 项目全生命周期管理，支持需求分析、模式选择、团队组建、阶段管理和多Agent协作协调。使用 @go 触发PO系统进行项目规划和管理。"
---

# PO Core Skill

## 核心职责

PO Core 是 PicoClaw 的产品经理核心技能，负责：

1. **需求分析**: 接收用户输入，分析任务类型和复杂度
2. **模式选择**: 根据项目特征选择合适的执行模式 (Standard/Free/Hybrid)
3. **团队组建**: 动态组建合适的团队配置
4. **阶段管理**: 管理项目阶段转换和质量门禁
5. **协作协调**: 协调多 Agent 协作执行

## 触发方式

使用 `@go` 命令触发 PO 系统：

```bash
@go "开发电商购物车功能"
@go "实现用户认证系统"
@go "设计新的数据分析平台"
```

## HARNESS.md 集成

每个任务开始前强制执行：
```bash
First, read and strictly follow HARNESS.md in repository root.
Do NOT deviate from its rules.
```

## 决策算法

### 需求分析流程

1. **解析用户输入**
   - 提取核心功能需求
   - 识别技术约束条件
   - 评估项目复杂度

2. **模式推荐**
   - **Standard Mode**: 标准化项目，需求明确，流程成熟
   - **Free Mode**: 创新项目，需求探索，高度灵活性
   - **Hybrid Mode**: 混合项目，部分标准化部分探索

3. **团队配置**
   - 根据项目阶段配置角色
   - 识别关键技能需求
   - 分配团队职责

4. **阶段规划**
   - discovery: 需求发现和市场调研
   - planning: 架构设计和实施计划
   - development: 功能开发和实现
   - validation: 测试验证和质量保证

## 输出格式

### 任务分析结果模板

```markdown
📋 **PO任务分析结果**

**需求**: {需求描述}
**类型**: {任务类型}
**复杂度**: {复杂度评分 1-10}
**预估工期**: {时间估算}

**推荐模式**: {模式名称}
**理由**: {选择理由}

**阶段规划**:
1. {阶段1名称} ({时间}) - {描述}
2. {阶段2名称} ({时间}) - {描述}
3. {阶段3名称} ({时间}) - {描述}

**团队配置**:
- {阶段1}: {角色配置}
- {阶段2}: {角色配置}
- {阶段3}: {角色配置}

是否开始执行？[Y/n/customize]
```

### 模式选择矩阵

| 项目特征 | Standard | Free | Hybrid |
|---------|----------|-------|--------|
| 需求明确度 | 高 | 低 | 中 |
| 技术成熟度 | 成熟 | 新兴 | 混合 |
| 团队经验 | 丰富 | 探索 | 中等 |
| 时间约束 | 紧张 | 宽松 | 中等 |
| 质量要求 | 稳定 | 创新 | 平衡 |

## 团队角色定义

### 核心角色

- **Analyst**: 需求分析师，负责需求分析和市场调研
- **Architect**: 架构师，负责技术架构设计
- **Developer**: 开发者，负责功能实现
- **QA**: 测试工程师，负责质量保证

### 角色能力矩阵

| 角色 | discovery | planning | development | validation |
|------|-----------|----------|-------------|------------|
| Analyst | 主导 | 支持 | 协调 | 审查 |
| Architect | 支持 | 主导 | 指导 | 审查 |
| Developer | 协调 | 实现 | 主导 | 测试 |
| QA | 验证 | 审查 | 测试 | 主导 |

## 质量门禁

### 阶段转换条件

- **discovery → planning**: 需求冻结 AND 市场分析批准
- **planning → development**: 架构批准 AND 实施计划就绪
- **development → validation**: 功能实现完成 AND 单元测试通过
- **validation → 完成**: 集成测试通过 AND 用户验收

### 自动化检查

- 需求完整性检查
- 架构一致性验证
- 代码质量扫描
- 测试覆盖率检查

## Ralph Wiggum Loop

任务完成后自动执行：
```bash
After implementing:
1. Run all tests + lint + type check until 100% green.
2. Self-review against HARNESS.md (list violations).
3. If fail → analyze root cause in /logs/ → fix → repeat.
4. Only create PR when green + self-review clean.
5. Update HARNESS.md/ADR if architecture changed.
6. Suggest 2–3 next improvements.
```

## 协作机制

### 层级协调

- PO → Phase Lead → Team Members
- 自上而下的目标分解和指导

### 同级协作

- Phase Leads 之间的协调
- 跨团队资源协调

### 领导协调

- PO 负责最终决策
- 冲突解决和优先级调整

## 错误处理

### 常见错误场景

1. **需求不明确**: 启动 discovery 阶段进行需求澄清
2. **技术风险**: 增加 planning 阶段进行技术验证
3. **资源不足**: 调整团队配置或延长项目周期
4. **质量问题**: 增加验证环节和代码审查

### 恢复策略

- 重新评估项目复杂度
- 调整执行模式
- 重新配置团队
- 修改阶段计划

## 使用示例

### 标准项目执行

```bash
@go "开发用户登录功能"

# PO 响应：
📋 **PO任务分析结果**

**需求**: 开发用户登录功能
**类型**: 功能开发
**复杂度**: 3/10
**预估工期**: 3-5天

**推荐模式**: Standard
**理由**: 标准化功能，技术成熟，流程明确

**阶段规划**:
1. discovery (1天) - 需求分析和技术选型
2. planning (1天) - 架构设计和API设计
3. development (2-3天) - 功能实现和单元测试
4. validation (1天) - 集成测试和用户验收

**团队配置**:
- discovery: analyst(主导) + architect(支持)
- planning: architect(主导) + analyst(支持)
- development: developer(主导) + qa(测试)
- validation: qa(主导) + developer(支持)

是否开始执行？[Y/n/customize]
```

### 创新项目执行

```bash
@go "探索AI驱动的推荐系统"

# PO 响应：
📋 **PO任务分析结果**

**需求**: 探索AI驱动的推荐系统
**类型**: 技术探索
**复杂度**: 8/10
**预估工期**: 2-3周

**推荐模式**: Free
**理由**: 技术新颖，需求不明确，需要大量探索

**阶段规划**:
1. discovery (1周) - 市场调研和技术研究
2. planning (3-4天) - 原型设计和验证
3. development (1-2周) - 迭代开发和优化
4. validation (2-3天) - 效果评估和用户测试

**团队配置**:
- discovery: analyst(主导) + architect(探索)
- planning: architect(主导) + analyst(验证)
- development: developer(实现) + architect(指导)
- validation: qa(测试) + analyst(评估)

是否开始执行？[Y/n/customize]
```

## 技能依赖

依赖以下技能：
- `phase-manager`: 阶段管理
- `standard-mode`: 标准化执行模式
- `free-mode`: 自由执行模式
- `hybrid-mode`: 混合执行模式
- `role-analyst`: 需求分析师角色
- `role-architect`: 架构师角色
- `role-developer`: 开发者角色
- `role-qa`: 测试工程师角色

## 配置文件

PO 系统使用 `.goagents/` 目录下的配置文件：

- `config.yaml`: 全局配置
- `workflows/`: 工作流定义
- `teams/`: 团队配置
- `phases/`: 阶段模板

## 扩展性

### 自定义模式

支持添加新的执行模式：
1. 创建新模式技能
2. 在 PO Core 中注册
3. 更新模式选择逻辑

### 自定义角色

支持添加新的团队角色：
1. 创建角色技能
2. 定义角色能力矩阵
3. 更新团队配置逻辑

### 自定义阶段

支持添加新的项目阶段：
1. 创建阶段模板
2. 定义转换条件
3. 更新阶段管理逻辑
