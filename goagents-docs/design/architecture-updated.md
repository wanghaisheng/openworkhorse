# 🏗️ Go Agents 架构设计（完整层级版）

## 🎯 概述

基于核心模型的完整层级结构，Go Agents 实现了工业级的多 Agent 协作平台。

## 📋 完整层级结构

### 核心模型层级
```
Workflow
  -> Phase
    -> Milestone
      -> Task
        -> Team
          -> Team Role
            -> Team Member Agent
```

### 层级职责分工

#### 1. **Workflow（工作流）**
- **职责**: 项目整体流程和阶段定义
- **配置**: `~/.picoclaw/workspace/.goagents/workflows/`
- **示例**: `ecommerce-development.yaml`

#### 2. **Phase（阶段）**
- **职责**: 项目的主要阶段
- **阶段类型**: discovery, planning, development, validation
- **配置**: `~/.picoclaw/workspace/.goagents/phases/`

#### 3. **Milestone（里程碑）**
- **职责**: 阶段内的关键里程碑，挂载质量门禁
- **特点**: 支持复杂项目的里程碑管理
- **质量门禁**: 挂载在里程碑级别

#### 4. **Task（任务）**
- **职责**: 具体的执行任务，包含内部子任务拆解
- **特点**: 每个任务都有详细的内部任务分解
- **执行**: 通过技能系统执行

#### 5. **Team（团队）**
- **职责**: 执行任务的团队配置
- **配置**: `~/.picoclaw/workspace/.goagents/teams/`
- **结构**: Phase Lead + Team Members

#### 6. **Team Role（团队角色）**
- **职责**: 团队中的角色定义
- **实现**: 通过技能系统定义角色
- **变体**: 支持角色的专业化变体

#### 7. **Team Member Agent（团队成员）**
- **职责**: 具体的执行者
- **实现**: 通过技能变体实现
- **执行**: 包含内部任务拆解

## 🎭 技能内部任务拆解

### 任务分解策略

#### 1. **Template-based（模板化）**
```yaml
execution:
  task_breakdown:
    strategy: "template_based"
    granularity: "medium"
    max_depth: 3
  
  subtask_execution:
    sequential: true
    subtasks:
      - id: "data_collection_preparation"
        name: "数据收集准备"
        estimated_time: "4小时"
        quality_gates:
          - gate: "data_source_completeness"
            threshold: 90
```

#### 2. **Milestone-driven（里程碑驱动）**
```yaml
execution:
  task_breakdown:
    strategy: "milestone_driven"
    granularity: "coarse"
    max_depth: 2
  
  subtask_execution:
    parallel: true
    subtasks:
      - id: "research_strategy_design"
        name: "研究策略设计"
        estimated_time: "1-2天"
```

#### 3. **Hybrid（混合模式）**
```yaml
execution:
  task_breakdown:
    strategy: "hybrid"
    granularity: "fine"
    max_depth: 4
  
  subtask_execution:
    mixed: true
    subtasks:
      - id: "standard_part"
        name: "标准化部分"
        execution_type: "template"
      - id: "dynamic_part"
        name: "动态部分"
        execution_type: "milestone_driven"
```

## 🔄 执行流程

### 完整执行示例
```bash
# 用户输入
@go "开发电商购物车功能"

# 执行流程：
🎯 **Workflow**: ecommerce-development
  🔄 **Phase**: discovery
    🎯 **Milestone 1**: 项目启动和需求分析
      🔧 **Task**: project-background-research
        🎭 **Skill**: role-analyst (research_specialist)
          📝 **SubTask 1**: stakeholder_identification
          📝 **SubTask 2**: requirement_gathering
          📝 **SubTask 3**: scope_definition
    🎯 **Milestone 2**: 市场调研和用户研究
      🔧 **Task**: market-analysis
        🎭 **Skill**: role-analyst (market_specialist)
          📝 **SubTask 1**: data_collection_preparation
          📝 **SubTask 2**: market_size_analysis
          📝 **SubTask 3**: competitor_analysis
          📝 **SubTask 4**: data_cleaning_processing
          📝 **SubTask 5**: trend_analysis
          📝 **SubTask 6**: report_generation
```

### 执行引擎实现
```go
// 支持完整层级的执行引擎
func (we *WorkflowExecutorUpdated) ExecuteWorkflow(workflowID string, userInput string) error {
  // 1. 执行每个阶段
  for _, phase := range workflow.Phases {
    // 2. 执行每个里程碑
    for _, milestone := range phaseConfig.Milestones {
      // 3. 执行每个任务
      for _, task := range milestone.Tasks {
        // 4. 执行技能内部子任务
        err := we.executionEngine.ExecuteSkillWithSubtasks(skill, variant, task, userInput)
      }
    }
  }
}
```

## 🎯 三种任务模式

### 1. **Standard 模式（模板驱动）**
- **特点**: 基于预定义模板执行
- **适用**: 标准化项目，质量一致性要求高
- **任务分解**: template_based 策略
- **质量控制**: 严格的质量门禁

### 2. **Free 模式（Phase Lead 主导）**
- **特点**: Phase Lead 动态规划
- **适用**: 创新项目，探索性强
- **任务分解**: milestone_driven 策略
- **灵活性**: 高度灵活的任务调整

### 3. **Hybrid 模式（混合执行）**
- **特点**: 标准化 + 灵活结合
- **适用**: 复杂项目，平衡效率与创新
- **任务分解**: hybrid 策略
- **适应性**: 根据项目特征选择

## 📊 质量控制体系

### 多层次质量门禁

#### 1. **SubTask 级别**
```yaml
quality_gates:
  - gate: "data_source_completeness"
    threshold: 90
    reviewer: "phase_lead"
  - gate: "data_accuracy"
    threshold: 80
    reviewer: "self"
```

#### 2. **Task 级别**
```yaml
quality_gates:
  - gate: "task_completeness"
    threshold: 85
    reviewer: "phase_lead"
  - gate: "output_quality"
    threshold: 80
    reviewer: "self"
```

#### 3. **Milestone 级别**
```yaml
quality_gates:
  - gate: "requirement_completeness"
    threshold: 90
    reviewer: "po"
  - gate: "market_analysis_depth"
    threshold: 80
    reviewer: "phase_lead"
```

#### 4. **Phase 级别**
```yaml
quality_gates:
  - gate: "phase_objectives_met"
    threshold: 85
    reviewer: "po"
  - gate: "deliverables_complete"
    threshold: 90
    reviewer: "stakeholder"
```

## 🔧 配置系统

### 目录结构
```
~/.picoclaw/workspace/
├── skills/                     # 技能实现
│   ├── po-core/               # PO 核心技能
│   ├── task-modes/            # 任务模式技能
│   ├── team-roles/            # 团队角色技能
│   └── workflows/             # 完整工作流
└── .goagents/                # 配置系统
    ├── config.yaml           # 全局配置
    ├── workflows/            # 工作流配置
    ├── phases/               # 阶段配置（含里程碑）
    ├── teams/                # 团队配置
    ├── tasks/                # 任务模板
    ├── registry/             # 配置管理器
    └── cli/                  # 独立CLI工具
```

### 配置文件示例

#### Workflow 配置
```yaml
# ecommerce-development.yaml
workflow:
  id: "ecommerce-development"
  phases:
    - id: "discovery"
      team_config: "discovery-team"
      task_mode: "standard"
    - id: "planning"
      team_config: "planning-team"
      task_mode: "standard"
```

#### Phase 配置（含里程碑）
```yaml
# discovery.yaml
phase:
  milestones:
    milestone_1:
      name: "项目启动和需求分析"
      success_criteria:
        - "项目目标明确"
        - "需求范围确定"
      quality_gates:
        - gate: "stakeholder_alignment"
          threshold: 85
          reviewer: "po"
      tasks:
        - task_id: "project-background-research"
          agent: "phase_lead"
          variant: "research_specialist"
```

#### Team 配置
```yaml
# discovery-team.yaml
team:
  phases:
    discovery:
      phase_lead:
        agent: "analyst"
        variant: "research_specialist"
      team_members:
        - member_id: "market_analyst"
          agent: "analyst"
          variant: "market_specialist"
```

#### Skill 配置（含内部拆解）
```yaml
# role-analyst/SKILL.md
variants:
  market_specialist:
    execution:
      task_breakdown:
        strategy: "template_based"
        granularity: "medium"
      subtask_execution:
        subtasks:
          - id: "data_collection_preparation"
          - id: "market_size_analysis"
          - id: "competitor_analysis"
```

## 🎯 核心优势

### 1. **完整层级结构**
- 符合核心模型的完整层级
- 支持复杂项目的里程碑管理
- 灵活的任务分解策略

### 2. **技能内部拆解**
- 每个技能都有详细的内部任务分解
- 支持三种分解策略
- 完整的依赖关系管理

### 3. **多层次质量控制**
- SubTask、Task、Milestone、Phase 四级质量门禁
- 自动化质量检查
- 多层次审核机制

### 4. **三种执行模式**
- Standard 模式：模板驱动
- Free 模式：Phase Lead 主导
- Hybrid 模式：混合执行

### 5. **零侵入设计**
- 完全基于现有 PicoClaw 技能系统
- 不修改核心代码
- 保持超轻量特性

## 🚀 技术实现

### 1. **配置加载器**
```go
type ConfigLoader struct {
    workflowLoader  *WorkflowLoader
    phaseLoader     *PhaseLoader
    teamLoader      *TeamLoader
    skillLoader     *SkillLoader
}
```

### 2. **执行引擎**
```go
type WorkflowExecutorUpdated struct {
    configLoader    *ConfigLoader
    skillsLoader    *SkillsLoader
    executionEngine *SkillExecutionEngineUpdated
}
```

### 3. **技能执行引擎**
```go
type SkillExecutionEngineUpdated struct {
    skillRegistry map[string]*SkillDefinition
}
```

### 4. **CLI 集成**
```go
// 混合 CLI 集成
cmd.AddCommand(
    goagents.NewGoagentsCommand(), // 新增 goagents 命令
)
```

## 🎉 总结

通过完整层级结构的设计，Go Agents 实现了：

1. **工业级架构** - 完整的层级结构和质量控制
2. **高度灵活性** - 三种执行模式和任务分解策略
3. **零侵入设计** - 完全基于现有技能系统
4. **可扩展性** - 模块化设计和配置驱动
5. **用户友好** - 简单的 @go 触发和完整的 CLI 工具

这个设计让 Go Agents 真正成为工业级的多 Agent 协作平台！🚀
