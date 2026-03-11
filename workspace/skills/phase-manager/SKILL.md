---
name: phase-manager
description: "阶段管理器 - 管理项目阶段转换，支持discovery、planning、development、validation四个阶段的流程控制和质量门禁。"
---

# Phase Manager Skill

## 支持的阶段

### Discovery 阶段
- **目标**: 完成市场调研、需求澄清和PRD形成
- **时长**: 1-2周
- **关键活动**: 市场分析、用户研究、需求定义
- **交付物**: 市场分析报告、用户研究报告、需求规格书

### Planning 阶段
- **目标**: 完成架构设计和实施计划
- **时长**: 1周
- **关键活动**: 技术选型、架构设计、API设计
- **交付物**: 架构文档、技术规范、实施计划

### Development 阶段
- **目标**: 完成功能开发和实现
- **时长**: 2-4周
- **关键活动**: 功能开发、单元测试、代码审查
- **交付物**: 功能代码、测试用例、技术文档

### Validation 阶段
- **目标**: 完成测试验证和质量保证
- **时长**: 3-5天
- **关键活动**: 集成测试、用户验收、性能测试
- **交付物**: 测试报告、验收文档、发布准备

## 阶段转换规则

### Discovery → Planning
**转换条件**:
- 需求文档已完成并批准
- 市场分析报告通过评审
- 技术可行性验证完成
- 利益相关者签字确认

**质量门禁**:
- [ ] 需求完整性检查 (≥90%)
- [ ] 市场数据验证 (≥85%)
- [ ] 利益相关者确认 (100%)
- [ ] PO 批准

**自动化检查**:
```bash
# 需求文档完整性检查
scripts/check-requirement-completeness.sh

# 市场数据验证
scripts/validate-market-data.sh

# 利益相关者确认检查
scripts/check-stakeholder-approval.sh
```

### Planning → Development
**转换条件**:
- 架构设计文档完成
- 技术选型决策确定
- 实施计划制定完成
- 资源分配确认

**质量门禁**:
- [ ] 架构评审通过 (≥85%)
- [ ] 技术风险评估完成
- [ ] 实施计划可行性验证
- [ ] 资源分配确认

**自动化检查**:
```bash
# 架构一致性检查
scripts/check-architecture-consistency.sh

# 技术风险评估
scripts/assess-technical-risks.sh

# 实施计划验证
scripts/validate-implementation-plan.sh
```

### Development → Validation
**转换条件**:
- 功能开发完成
- 单元测试通过
- 代码审查完成
- 集成测试准备就绪

**质量门禁**:
- [ ] 代码覆盖率 ≥85%
- [ ] 单元测试通过率 100%
- [ ] 代码审查完成
- [ ] 静态代码分析通过

**自动化检查**:
```bash
# 测试覆盖率检查
scripts/check-test-coverage.sh

# 代码质量检查
scripts/run-code-quality-checks.sh

# 集成测试准备
scripts/prepare-integration-tests.sh
```

### Validation → 完成
**转换条件**:
- 集成测试通过
- 用户验收完成
- 性能测试达标
- 发布准备完成

**质量门禁**:
- [ ] 集成测试通过率 100%
- [ ] 用户验收通过
- [ ] 性能指标达标
- [ ] 发布检查清单完成

**自动化检查**:
```bash
# 集成测试执行
scripts/run-integration-tests.sh

# 性能测试
scripts/run-performance-tests.sh

# 发布检查
scripts/check-release-readiness.sh
```

## 阶段状态跟踪

### 状态模型

```yaml
phase_states:
  - NOT_STARTED: 未开始
  - IN_PROGRESS: 进行中
  - UNDER_REVIEW: 审查中
  - QUALITY_GATE: 质量门禁
  - COMPLETED: 已完成
  - BLOCKED: 阻塞
  - CANCELLED: 取消
```

### 状态转换矩阵

| 当前状态 | 可转换到 | 触发条件 |
|---------|---------|----------|
| NOT_STARTED | IN_PROGRESS | 阶段启动命令 |
| IN_PROGRESS | UNDER_REVIEW | 主要工作完成 |
| UNDER_REVIEW | QUALITY_GATE | 审查完成 |
| QUALITY_GATE | COMPLETED | 所有门禁通过 |
| QUALITY_GATE | IN_PROGRESS | 门禁失败，需要修复 |
| IN_PROGRESS | BLOCKED | 遇到阻塞问题 |
| BLOCKED | IN_PROGRESS | 阻塞问题解决 |

## 阶段模板系统

### Discovery 阶段模板

```yaml
discovery_template:
  objective: "完成市场调研、需求澄清和PRD形成"
  duration: "1-2周"
  
  standard_tasks:
    - name: "market-analysis"
      description: "市场调研和分析"
      estimated_time: "2-3天"
      required_agents: ["analyst"]
      deliverables: ["market-analysis-report.md"]
      
    - name: "user-research"
      description: "用户研究和访谈"
      estimated_time: "1-2周"
      required_agents: ["analyst", "pm"]
      deliverables: ["user-research-report.pdf"]
      
    - name: "requirement-definition"
      description: "需求定义和规格"
      estimated_time: "2-3天"
      required_agents: ["pm", "analyst"]
      deliverables: ["requirement-spec.md"]
      
  quality_gates:
    - gate: "market_depth"
      threshold: "80%"
      reviewer: "phase_lead"
      
    - gate: "requirement_clarity"
      threshold: "85%"
      reviewer: "po"
```

### Planning 阶段模板

```yaml
planning_template:
  objective: "完成架构设计和实施计划"
  duration: "1周"
  
  standard_tasks:
    - name: "architecture-design"
      description: "系统架构设计"
      estimated_time: "2-3天"
      required_agents: ["architect"]
      deliverables: ["architecture-diagram.png", "architecture-doc.md"]
      
    - name: "technology-selection"
      description: "技术选型和评估"
      estimated_time: "1-2天"
      required_agents: ["architect", "tech-lead"]
      deliverables: ["tech-selection-report.md"]
      
    - name: "implementation-planning"
      description: "实施计划制定"
      estimated_time: "1-2天"
      required_agents: ["tech-lead", "pm"]
      deliverables: ["implementation-plan.md", "timeline.md"]
      
  quality_gates:
    - gate: "architecture_soundness"
      threshold: "85%"
      reviewer: "tech_lead"
      
    - gate: "plan_feasibility"
      threshold: "90%"
      reviewer: "po"
```

### Development 阶段模板

```yaml
development_template:
  objective: "完成功能开发和实现"
  duration: "2-4周"
  
  standard_tasks:
    - name: "feature-development"
      description: "功能开发实现"
      estimated_time: "2-3周"
      required_agents: ["developer"]
      deliverables: ["source-code/", "unit-tests/"]
      
    - name: "unit-testing"
      description: "单元测试编写"
      estimated_time: "3-5天"
      required_agents: ["developer", "qa"]
      deliverables: ["test-results.xml", "coverage-report.html"]
      
    - name: "code-review"
      description: "代码审查"
      estimated_time: "2-3天"
      required_agents: ["tech-lead", "senior-dev"]
      deliverables: ["review-comments.md", "approved-code/"]
      
  quality_gates:
    - gate: "test_coverage"
      threshold: "85%"
      reviewer: "qa"
      
    - gate: "code_quality"
      threshold: "80%"
      reviewer: "tech_lead"
```

### Validation 阶段模板

```yaml
validation_template:
  objective: "完成测试验证和质量保证"
  duration: "3-5天"
  
  standard_tasks:
    - name: "integration-testing"
      description: "集成测试执行"
      estimated_time: "2-3天"
      required_agents: ["qa", "developer"]
      deliverables: ["integration-test-report.md"]
      
    - name: "user-acceptance"
      description: "用户验收测试"
      estimated_time: "1-2天"
      required_agents: ["qa", "pm", "user"]
      deliverables: ["uat-report.md", "user-feedback.md"]
      
    - name: "performance-testing"
      description: "性能测试"
      estimated_time: "1天"
      required_agents: ["qa", "devops"]
      deliverables: ["performance-report.md", "benchmark-results.json"]
      
  quality_gates:
    - gate: "integration_success"
      threshold: "100%"
      reviewer: "qa_lead"
      
    - gate: "user_satisfaction"
      threshold: "85%"
      reviewer: "po"
```

## 阶段管理API

### 阶段启动

```bash
# 启动 discovery 阶段
phase-manager start discovery --project "ecommerce-cart"

# 启动 planning 阶段
phase-manager start planning --project "ecommerce-cart" --from discovery
```

### 阶段状态查询

```bash
# 查询当前阶段状态
phase-manager status --project "ecommerce-cart"

# 查询所有阶段状态
phase-manager status --project "ecommerce-cart" --all
```

### 阶段转换

```bash
# 手动触发阶段转换
phase-manager transition --project "ecommerce-cart" --from discovery --to planning

# 强制转换（跳过质量门禁）
phase-manager transition --project "ecommerce-cart" --from planning --to development --force
```

### 质量门禁检查

```bash
# 检查当前阶段质量门禁
phase-manager check-gates --project "ecommerce-cart" --phase discovery

# 重新运行失败的检查
phase-manager retry-gates --project "ecommerce-cart" --phase discovery --gate market_depth
```

## 配置文件

### 项目配置

```yaml
# .goagents/projects/ecommerce-cart.yaml
project:
  name: "ecommerce-cart"
  description: "电商购物车功能开发"
  
phases:
  discovery:
    start_date: "2026-03-12"
    assigned_team: ["analyst", "pm"]
    custom_tasks:
      - name: "competitor-analysis"
        description: "竞品分析"
        estimated_time: "2天"
        
  planning:
    dependencies: ["discovery"]
    assigned_team: ["architect", "tech-lead"]
    
quality_gates:
  strict_mode: true
  auto_promote: false
  notification_channels: ["slack", "email"]
```

### 全局配置

```yaml
# .goagents/config/phase-manager.yaml
phase_manager:
  default_phases: ["discovery", "planning", "development", "validation"]
  quality_gates:
    strict_mode: true
    parallel_checks: true
    timeout_minutes: 30
    
notifications:
  slack_webhook: "${SLACK_WEBHOOK_URL}"
  email_smtp: "${SMTP_CONFIG}"
  
reporting:
  dashboard_url: "https://dashboard.company.com/phases"
  export_formats: ["json", "yaml", "csv"]
```

## 监控和报告

### 阶段进度报告

```bash
# 生成阶段进度报告
phase-manager report --project "ecommerce-cart" --format markdown

# 生成跨项目阶段概览
phase-manager report --all-projects --format html
```

### 关键指标

- **阶段周期时间**: 各阶段的平均完成时间
- **质量门禁通过率**: 各阶段质量门禁的首次通过率
- **阻塞频率**: 阶段阻塞的发生频率和原因
- **资源利用率**: 各角色在不同阶段的投入情况

## 错误处理

### 常见错误场景

1. **阶段转换失败**
   - 检查质量门禁状态
   - 查看错误日志
   - 修复问题后重试

2. **质量门禁超时**
   - 检查自动化脚本状态
   - 增加超时时间配置
   - 手动执行检查

3. **团队资源冲突**
   - 查看资源分配情况
   - 调整阶段时间安排
   - 申请额外资源

### 恢复程序

```bash
# 查看阶段错误日志
phase-manager logs --project "ecommerce-cart" --phase discovery

# 重置阶段状态
phase-manager reset --project "ecommerce-cart" --phase discovery

# 回滚到上一阶段
phase-manager rollback --project "ecommerce-cart" --from planning --to discovery
```

## 集成接口

### 与 PO Core 集成

```typescript
interface PhaseManager {
  startPhase(projectId: string, phaseName: string): Promise<PhaseResult>;
  getPhaseStatus(projectId: string): Promise<PhaseStatus>;
  transitionPhase(projectId: string, fromPhase: string, toPhase: string): Promise<TransitionResult>;
  checkQualityGates(projectId: string, phaseName: string): Promise<GateResult[]>;
}
```

### 与团队角色集成

```typescript
interface TeamRoleIntegration {
  assignTeamMember(projectId: string, phaseName: string, role: string, member: string): Promise<void>;
  getTeamAssignments(projectId: string, phaseName: string): Promise<TeamAssignment[]>;
  updateWorkload(projectId: string, phaseName: string, memberId: string, workload: number): Promise<void>;
}
```

## 扩展性

### 自定义阶段

支持添加自定义阶段：

1. 定义阶段模板
2. 配置转换规则
3. 设置质量门禁
4. 注册到系统

### 自定义质量门禁

支持添加自定义质量门禁：

1. 编写检查脚本
2. 定义通过标准
3. 配置审查流程
4. 集成到阶段模板

### 自定义通知

支持添加自定义通知渠道：

1. 实现通知接口
2. 配置触发条件
3. 设置消息模板
4. 注册到系统
