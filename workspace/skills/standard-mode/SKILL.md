---
name: standard-mode
description: "标准化任务执行模式 - 基于模板的标准化执行流程，适合需求明确、技术成熟的项目。提供严格的步骤控制和质量保证。"
---

# Standard Mode Skill

## 模式概述

Standard Mode 是基于模板的标准化执行模式，适合：

- **需求明确**: 功能需求清晰，变更可能性低
- **技术成熟**: 使用成熟技术栈，风险可控
- **流程标准**: 有标准化的开发流程和最佳实践
- **质量要求高**: 对代码质量和交付质量有严格要求

## 执行原则

### 严格遵循模板
- 必须按照模板步骤顺序执行
- 跳过步骤需要特殊批准
- 每个步骤都有明确的输入输出

### 质量门禁强制
- 每个步骤必须通过质量门禁
- 自动化检查失败时停止执行
- 人工审查必须完成

### 标准化输出
- 统一的文档格式
- 标准化的代码结构
- 一致的测试覆盖

## 模板执行机制

### 任务分解策略

```yaml
execution:
  task_breakdown:
    strategy: "template_based"
    granularity: "medium"
    max_depth: 3
    
  subtask_execution:
    sequential: true
    dependency_check: strict
    quality_gates: mandatory
```

### 模板步骤结构

每个模板步骤包含：

```yaml
step_definition:
  name: "步骤名称"
  description: "详细描述"
  estimated_time: "时间估算"
  
  inputs:
    - "输入1"
    - "输入2"
    
  outputs:
    - "输出1"
    - "输出2"
    
  dependencies: ["前置步骤"]
  
  quality_gates:
    - gate: "门禁名称"
      threshold: "阈值"
      auto_check: true
      
  execution_script: "scripts/step-script.sh"
```

## 标准任务模板

### 市场分析任务模板

```yaml
template_id: "market-analysis-standard"
name: "标准化市场分析"
category: "analysis"
estimated_time: "2-3天"
required_agents: ["market_analyst"]
difficulty: "medium"

template_steps:
  step_1:
    name: "数据收集准备"
    description: "准备市场规模评估所需的数据源和工具"
    estimated_time: "2小时"
    inputs: ["项目背景", "目标市场定义"]
    outputs: ["数据源清单.md", "工具准备清单.xlsx"]
    quality_gates:
      - gate: "data_source_completeness"
        threshold: 90
        auto_check: true
      - gate: "tool_availability"
        threshold: 100
        auto_check: true
        
  step_2:
    name: "市场规模数据收集"
    description: "从多个数据源收集市场规模相关数据"
    estimated_time: "4小时"
    inputs: ["数据源清单.md", "工具准备清单.xlsx"]
    outputs: ["原始数据.xlsx", "数据收集日志.md"]
    depends_on: ["step_1"]
    quality_gates:
      - gate: "data_completeness"
        threshold: 85
        auto_check: true
      - gate: "data_accuracy"
        threshold: 90
        auto_check: false
        
  step_3:
    name: "数据清洗和整理"
    description: "清洗收集的原始数据，整理成标准格式"
    estimated_time: "3小时"
    inputs: ["原始数据.xlsx", "数据收集日志.md"]
    outputs: ["清洗后数据.xlsx", "数据清洗记录.md"]
    depends_on: ["step_2"]
    quality_gates:
      - gate: "data_consistency"
        threshold: 95
        auto_check: true
      - gate: "format_standardization"
        threshold: 100
        auto_check: true
        
  step_4:
    name: "市场规模分析"
    description: "基于清洗后的数据进行市场规模分析"
    estimated_time: "6小时"
    inputs: ["清洗后数据.xlsx", "数据清洗记录.md"]
    outputs: ["市场规模分析报告.md", "市场规模图表.png"]
    depends_on: ["step_3"]
    quality_gates:
      - gate: "analysis_depth"
        threshold: 80
        auto_check: false
      - gate: "insight_quality"
        threshold: 85
        auto_check: false
        
  step_5:
    name: "报告生成和审核"
    description: "生成最终的市场规模分析报告并进行审核"
    estimated_time: "2小时"
    inputs: ["市场规模分析报告.md", "市场规模图表.png"]
    outputs: ["最终报告.pdf", "审核记录.md"]
    depends_on: ["step_4"]
    quality_gates:
      - gate: "report_completeness"
        threshold: 100
        auto_check: true
      - gate: "review_approval"
        threshold: 100
        auto_check: false
```

### 功能开发任务模板

```yaml
template_id: "feature-development-standard"
name: "标准化功能开发"
category: "development"
estimated_time: "3-5天"
required_agents: ["developer", "qa"]
difficulty: "medium"

template_steps:
  step_1:
    name: "需求理解和拆解"
    description: "深入理解功能需求，拆解为具体实现任务"
    estimated_time: "4小时"
    inputs: ["需求文档", "技术规范"]
    outputs: ["任务拆解.md", "技术方案.md"]
    quality_gates:
      - gate: "requirement_understanding"
        threshold: 90
        auto_check: false
        
  step_2:
    name: "代码结构设计"
    description: "设计代码结构和模块划分"
    estimated_time: "6小时"
    inputs: ["任务拆解.md", "技术方案.md"]
    outputs: ["代码结构图.png", "接口设计.md"]
    depends_on: ["step_1"]
    quality_gates:
      - gate: "design_quality"
        threshold: 85
        auto_check: false
        
  step_3:
    name: "核心功能实现"
    description: "实现核心业务逻辑"
    estimated_time: "1-2天"
    inputs: ["代码结构图.png", "接口设计.md"]
    outputs: ["源代码/", "实现文档.md"]
    depends_on: ["step_2"]
    quality_gates:
      - gate: "code_quality"
        threshold: 80
        auto_check: true
      - gate: "functionality_correctness"
        threshold: 90
        auto_check: false
        
  step_4:
    name: "单元测试编写"
    description: "编写和执行单元测试"
    estimated_time: "1天"
    inputs: ["源代码/", "实现文档.md"]
    outputs: ["测试代码/", "测试报告.html"]
    depends_on: ["step_3"]
    quality_gates:
      - gate: "test_coverage"
        threshold: 85
        auto_check: true
      - gate: "test_pass_rate"
        threshold: 100
        auto_check: true
        
  step_5:
    name: "代码审查和优化"
    description: "进行代码审查和性能优化"
    estimated_time: "4小时"
    inputs: ["源代码/", "测试报告.html"]
    outputs: ["审查报告.md", "优化后代码/"]
    depends_on: ["step_4"]
    quality_gates:
      - gate: "code_review_approval"
        threshold: 100
        auto_check: false
      - gate: "performance_benchmark"
        threshold: 80
        auto_check: true
```

## 模板执行流程

### 1. 模板选择

```bash
# 根据任务类型自动选择模板
standard-mode select-template --task-type "market-analysis"

# 手动指定模板
standard-mode select-template --template-id "market-analysis-standard"
```

### 2. 参数配置

```bash
# 配置模板参数
standard-mode configure --parameter "market_region=global" --parameter "time_range=2024-2025"

# 设置质量门禁阈值
standard-mode set-gate-threshold --gate "data_completeness" --threshold 90
```

### 3. 执行启动

```bash
# 启动模板执行
standard-mode execute --template-id "market-analysis-standard" --project "ecommerce-analysis"

# 从特定步骤开始执行
standard-mode execute --template-id "market-analysis-standard" --from-step "step_3"
```

### 4. 进度监控

```bash
# 查看执行进度
standard-mode status --project "ecommerce-analysis"

# 查看详细步骤状态
standard-mode status --project "ecommerce-analysis" --detailed
```

## 质量保证机制

### 自动化检查

```bash
# 数据完整性检查
scripts/check-data-completeness.sh --input "原始数据.xlsx" --threshold 85

# 代码质量检查
scripts/run-code-quality.sh --source "src/" --threshold 80

# 测试覆盖率检查
scripts/check-test-coverage.sh --test-reports "reports/" --threshold 85
```

### 人工审查流程

1. **自我审查**: 执行者完成自我审查清单
2. **同行审查**: 同团队成员进行交叉审查
3. **专家审查**: 领域专家进行深度审查
4. **最终批准**: Phase Lead 进行最终批准

### 审查清单模板

```markdown
## [步骤名称] 审查清单

### 自动化检查
- [ ] 数据完整性检查通过 (≥85%)
- [ ] 格式标准化检查通过 (100%)
- [ ] 性能基准测试通过 (≥80%)

### 人工检查
- [ ] 逻辑正确性验证
- [ ] 业务规则符合性
- [ ] 输出质量评估
- [ ] 文档完整性

### 批准状态
- 执行者: _______________
- 审查者: _______________
- 批准者: _______________
- 日期: _______________
```

## 错误处理和恢复

### 常见错误类型

1. **质量门禁失败**
   - 自动化检查未通过
   - 人工审查未批准
   - 依赖步骤未完成

2. **执行超时**
   - 步骤执行时间超过预期
   - 资源不足导致延迟
   - 外部依赖不可用

3. **数据问题**
   - 输入数据格式错误
   - 数据不完整或不准确
   - 数据权限问题

### 恢复策略

```bash
# 重试失败的步骤
standard-mode retry-step --project "ecommerce-analysis" --step "step_2"

# 跳过步骤（需要批准）
standard-mode skip-step --project "ecommerce-analysis" --step "step_3" --reason "特殊原因说明"

# 回滚到之前步骤
standard-mode rollback --project "ecommerce-analysis" --to-step "step_1"

# 重新配置参数后执行
standard-mode re-execute --project "ecommerce-analysis" --reconfigure
```

## 性能优化

### 并行执行

对于独立的步骤，支持并行执行：

```yaml
execution:
  parallel_steps:
    - ["step_2", "step_3"]  # 这两个步骤可以并行执行
    - "step_4"              # 依赖前面的步骤
```

### 缓存机制

```bash
# 启用结果缓存
standard-mode configure --cache-enabled true --cache-ttl 24h

# 清理缓存
standard-mode clear-cache --project "ecommerce-analysis"
```

### 资源优化

```yaml
resource_optimization:
  cpu_limit: "2 cores"
  memory_limit: "4GB"
  disk_space: "10GB"
  network_bandwidth: "100Mbps"
```

## 监控和报告

### 实时监控

```bash
# 启动实时监控
standard-mode monitor --project "ecommerce-analysis" --real-time

# 查看资源使用情况
standard-mode resources --project "ecommerce-analysis"
```

### 执行报告

```bash
# 生成执行报告
standard-mode report --project "ecommerce-analysis" --format markdown

# 生成质量报告
standard-mode quality-report --project "ecommerce-analysis" --detailed
```

### 关键指标

- **执行效率**: 实际时间 vs 预估时间
- **质量指标**: 各质量门禁的通过率
- **资源利用率**: CPU、内存、存储使用情况
- **错误率**: 各步骤的错误发生频率

## 配置管理

### 全局配置

```yaml
# .goagents/config/standard-mode.yaml
standard_mode:
  default_quality_gates:
    data_completeness: 85
    code_quality: 80
    test_coverage: 85
    
  execution_limits:
    max_execution_time: "7 days"
    max_retry_attempts: 3
    timeout_per_step: "24 hours"
    
  notification_settings:
    slack_webhook: "${SLACK_WEBHOOK}"
    email_recipients: ["team@company.com"]
```

### 项目配置

```yaml
# .goagents/projects/ecommerce-analysis/standard-mode.yaml
project_config:
  template_overrides:
    market-analysis-standard:
      step_2:
        estimated_time: "6 hours"  # 覆盖默认时间
        
  custom_quality_gates:
    - gate: "business_insight_quality"
      threshold: 90
      reviewer: "business_analyst"
```

## 集成接口

### 与 Phase Manager 集成

```typescript
interface StandardModeIntegration {
  executeTemplate(templateId: string, projectId: string): Promise<ExecutionResult>;
  getExecutionStatus(projectId: string): Promise<ExecutionStatus>;
  retryFailedStep(projectId: string, stepId: string): Promise<void>;
}
```

### 与团队角色集成

```typescript
interface TeamRoleIntegration {
  assignStepToAgent(projectId: string, stepId: string, agentRole: string): Promise<void>;
  getAgentWorkload(agentRole: string): Promise<WorkloadInfo>;
  updateAgentSkills(agentRole: string, skills: string[]): Promise<void>;
}
```

## 扩展性

### 自定义模板

支持创建自定义模板：

1. 定义模板结构
2. 编写执行脚本
3. 设置质量门禁
4. 注册到系统

### 自定义质量门禁

支持添加自定义质量门禁：

1. 编写检查脚本
2. 定义通过标准
3. 配置审查流程
4. 集成到模板

### 自定义通知

支持添加自定义通知：

1. 实现通知接口
2. 配置触发条件
3. 设置消息模板
4. 注册到系统
