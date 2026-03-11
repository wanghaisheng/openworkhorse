---
name: free-mode
description: "自由模式 - Phase Lead主导的动态任务规划，适合创新项目、需求探索和技术验证。提供高度灵活性和适应性。"
---

# Free Mode Skill

## 模式概述

Free Mode 是 Phase Lead 主导的动态任务规划模式，适合：

- **创新项目**: 需要探索新技术或新方法
- **需求不明确**: 需要在执行过程中逐步明确需求
- **技术验证**: 需要验证技术可行性
- **高度不确定性**: 项目范围和目标可能发生变化

## 核心特点

### 动态任务生成
- Phase Lead 根据项目进展动态生成任务
- 支持任务优先级调整
- 可以随时添加或删除任务

### 灵活执行策略
- 不严格依赖预设模板
- 支持多种执行路径
- 允许根据实际情况调整方向

### 迭代式开发
- 小步快跑，快速验证
- 持续收集反馈
- 及时调整策略

## Phase Lead 权限配置

### 任务分析权限

```yaml
task_analysis:
  permissions:
    - "分析任务复杂度和范围"
    - "确定任务优先级"
    - "识别关键路径"
    - "评估资源需求"
    - "识别风险和依赖"
    
  tools:
    - "complexity_analyzer"
    - "priority_matrix"
    - "risk_assessment"
    - "resource_planner"
```

### 任务规划权限

```yaml
task_planning:
  permissions:
    - "制定任务执行计划"
    - "分解大任务为小任务"
    - "设置任务依赖关系"
    - "安排任务执行顺序"
    - "调整时间估算"
    
  tools:
    - "task_breakdown"
    - "dependency_mapper"
    - "timeline_planner"
    - "milestone_tracker"
```

### 任务分配权限

```yaml
task_assignment:
  permissions:
    - "分配任务给团队成员"
    - "协调团队协作"
    - "调整任务分配"
    - "解决资源冲突"
    - "优化工作负载"
    
  tools:
    - "team_allocator"
    - "workload_balancer"
    - "conflict_resolver"
    - "collaboration_coordinator"
```

### 任务调整权限

```yaml
task_adjustment:
  permissions:
    - "动态调整任务优先级"
    - "修改任务执行计划"
    - "添加或删除任务"
    - "调整任务时间估算"
    - "重新分配资源"
    
  tools:
    - "dynamic_planner"
    - "adaptive_scheduler"
    - "resource_reallocator"
    - "priority_adjuster"
```

## 动态任务生成流程

### 1. 里程碑分析

Phase Lead 首先分析里程碑目标：

```typescript
interface MilestoneAnalysis {
  analyzeMilestone(
    milestone: Milestone,
    projectContext: ProjectContext
  ): Promise<AnalysisResult> {
    
    // 1. 理解里程碑目标
    const objectives = await this.parseObjectives(milestone.description);
    
    // 2. 识别关键约束
    const constraints = await this.identifyConstraints(projectContext);
    
    // 3. 评估技术风险
    const risks = await this.assessTechnicalRisks(objectives, constraints);
    
    // 4. 确定成功标准
    const successCriteria = await this.defineSuccessCriteria(objectives);
    
    return {
      objectives,
      constraints,
      risks,
      successCriteria,
      complexity: this.calculateComplexity(objectives, risks),
      estimatedEffort: this.estimateEffort(objectives, constraints)
    };
  }
}
```

### 2. 任务分解策略

```typescript
interface TaskBreakdownStrategy {
  createTaskBreakdown(
    analysis: MilestoneAnalysis,
    teamConfig: TeamConfiguration
  ): Promise<TaskBreakdown> {
    
    // 1. 选择分解策略
    const strategy = this.selectBreakdownStrategy(analysis);
    
    // 2. 生成主要任务
    const mainTasks = await this.generateMainTasks(analysis, strategy);
    
    // 3. 分解子任务
    const subtasks = await this.breakdownSubtasks(mainTasks);
    
    // 4. 建立依赖关系
    const dependencies = await this.establishDependencies(subtasks);
    
    // 5. 估算工作量
    const estimates = await this.estimateEffort(subtasks);
    
    return {
      mainTasks,
      subtasks,
      dependencies,
      estimates,
      strategy
    };
  }
}
```

### 3. 动态任务生成

```typescript
interface DynamicTaskGenerator {
  async generateTasks(
    phase: ProjectPhase,
    milestone: Milestone,
    phaseLead: Agent,
    teamConfig: TeamConfiguration,
    projectContext: ProjectContext
  ): Promise<TaskExecutionPlan> {
    
    // 1. Phase Lead 分析里程碑目标
    const milestoneAnalysis = await phaseLead.analyzeMilestone(
      milestone,
      projectContext
    );
    
    // 2. Phase Lead 制定任务分解策略
    const taskBreakdownStrategy = await phaseLead.createTaskBreakdown(
      milestoneAnalysis,
      teamConfig
    );
    
    // 3. Phase Lead 生成具体任务
    const dynamicTasks = await phaseLead.generateTasks(
      taskBreakdownStrategy,
      teamConfig
    );
    
    // 4. 团队协作优化任务计划
    const optimizedTasks = await this.optimizeWithTeamInput(
      dynamicTasks,
      teamConfig
    );
    
    // 5. PO 审核任务计划
    const approvedTasks = await this.poAgent.approveTaskPlan(
      optimizedTasks,
      milestone
    );
    
    return approvedTasks;
  }
}
```

## 里程碑驱动示例

### 项目启动阶段

```yaml
project_kickoff_phase:
  milestone: "项目启动和需求分析"
  phase_lead: "analyst" (research_specialist)
  team: ["market_analyst", "user_researcher", "data_analyst"]
  
  dynamic_task_plan:
    task_1:
      name: "项目背景和目标调研"
      agent: "phase_lead"
      reasoning: "需要深入理解项目背景，明确项目目标和约束条件"
      estimated_time: "1-2天"
      deliverables: ["项目背景文档.md", "目标清单.md"]
      priority: "high"
      
    task_2:
      name: "关键利益相关者识别和访谈"
      agent: "phase_lead"
      reasoning: "项目涉及多个部门，需要识别关键决策者并了解他们的期望"
      estimated_time: "2-3天"
      deliverables: ["利益相关者清单.md", "访谈记录.md"]
      depends_on: ["task_1"]
      priority: "high"
      
    task_3:
      name: "需求范围和边界确定"
      agent: "phase_lead"
      reasoning: "明确项目范围，避免范围蔓延，确保项目可交付"
      estimated_time: "1天"
      deliverables: ["需求边界文档.md", "范围说明书.md"]
      depends_on: ["task_2"]
      priority: "medium"
      
    task_4:
      name: "市场调研和用户研究"
      agent: "market_analyst" + "user_researcher"
      reasoning: "需要团队协作收集市场数据和用户反馈"
      estimated_time: "1-2周"
      deliverables: ["市场分析报告.xlsx", "用户研究报告.pdf"]
      depends_on: ["task_3"]
      priority: "high"
      
    task_5:
      name: "竞争分析和技术调研"
      agent: "data_analyst"
      reasoning: "了解竞争格局和技术趋势，为技术选型提供依据"
      estimated_time: "3-5天"
      deliverables: ["竞争分析报告.md", "技术调研报告.md"]
      depends_on: ["task_4"]
      priority: "medium"
```

### 技术探索阶段

```yaml
technology_exploration_phase:
  milestone: "新技术可行性验证"
  phase_lead: "architect" (technology_specialist)
  team: ["tech_lead", "senior_dev", "devops_engineer"]
  
  dynamic_task_plan:
    task_1:
      name: "技术栈调研和对比"
      agent: "phase_lead"
      reasoning: "需要全面调研可选技术栈，对比优劣"
      estimated_time: "2-3天"
      deliverables: ["技术栈对比表.xlsx", "技术评估报告.md"]
      priority: "high"
      
    task_2:
      name: "核心概念验证(POC)"
      agent: "tech_lead" + "senior_dev"
      reasoning: "通过实际代码验证核心概念的可行性"
      estimated_time: "1-2周"
      deliverables: ["poc代码/", "验证报告.md", "性能测试结果.json"]
      depends_on: ["task_1"]
      priority: "high"
      
    task_3:
      name: "架构原型设计"
      agent: "phase_lead"
      reasoning: "基于POC结果设计系统架构原型"
      estimated_time: "3-5天"
      deliverables: ["架构原型图.png", "架构文档.md"]
      depends_on: ["task_2"]
      priority: "medium"
      
    task_4:
      name: "性能和扩展性测试"
      agent: "devops_engineer" + "senior_dev"
      reasoning: "验证技术方案的性能和扩展性表现"
      estimated_time: "1周"
      deliverables: ["性能测试报告.md", "扩展性分析.pdf"]
      depends_on: ["task_3"]
      priority: "medium"
      
    task_5:
      name: "风险评估和缓解方案"
      agent: "phase_lead"
      reasoning: "识别技术风险并制定相应的缓解方案"
      estimated_time: "2-3天"
      deliverables: ["风险评估报告.md", "缓解方案.md"]
      depends_on: ["task_4"]
      priority: "high"
```

## 任务调整机制

### 动态优先级调整

```typescript
interface DynamicPriorityAdjuster {
  adjustPriorities(
    tasks: Task[],
    context: ProjectContext,
    feedback: TeamFeedback[]
  ): Task[] {
    
    // 1. 分析当前进展
    const progress = this.analyzeProgress(tasks);
    
    // 2. 识别阻塞点
    const blockers = this.identifyBlockers(tasks, feedback);
    
    // 3. 评估紧急程度
    const urgency = this.assessUrgency(context, blockers);
    
    // 4. 重新计算优先级
    const adjustedTasks = tasks.map(task => ({
      ...task,
      priority: this.calculateNewPriority(task, progress, urgency)
    }));
    
    // 5. 排序任务
    return this.sortByPriority(adjustedTasks);
  }
}
```

### 任务添加和删除

```typescript
interface TaskManager {
  addTask(
    projectId: string,
    taskDefinition: TaskDefinition,
    reason: string
  ): Promise<Task> {
    
    // 1. 验证任务定义
    this.validateTaskDefinition(taskDefinition);
    
    // 2. 检查资源可用性
    const resources = this.checkResourceAvailability(taskDefinition);
    
    // 3. 评估影响
    const impact = this.assessImpact(taskDefinition, projectId);
    
    // 4. 创建任务
    const task = await this.createTask(taskDefinition, resources);
    
    // 5. 通知团队
    await this.notifyTeam(task, reason);
    
    return task;
  }
  
  removeTask(
    projectId: string,
    taskId: string,
    reason: string
  ): Promise<void> {
    
    // 1. 检查删除条件
    this.validateRemoval(taskId);
    
    // 2. 处理依赖关系
    await this.handleDependencies(taskId);
    
    // 3. 释放资源
    await this.releaseResources(taskId);
    
    // 4. 记录删除原因
    await this.logRemoval(taskId, reason);
    
    // 5. 通知团队
    await this.notifyTeamRemoval(taskId, reason);
  }
}
```

## 协作机制

### 团队决策流程

```yaml
decision_process:
  1._proposal:
    owner: "phase_lead"
    action: "提出任务调整建议"
    deliverable: "调整提案.md"
    
  2._team_review:
    participants: ["all_team_members"]
    action: "团队讨论和反馈"
    deliverable: "团队反馈.md"
    duration: "24小时"
    
  3._refinement:
    owner: "phase_lead"
    action: "根据反馈完善提案"
    deliverable: "完善后提案.md"
    
  4._final_approval:
    owner: "po"
    action: "最终批准"
    deliverable: "批准决策.md"
```

### 冲突解决

```typescript
interface ConflictResolver {
  resolveConflict(
    conflict: TaskConflict,
    team: TeamMember[]
  ): Promise<Resolution> {
    
    // 1. 分析冲突类型
    const conflictType = this.analyzeConflictType(conflict);
    
    // 2. 收集各方观点
    const perspectives = await this.collectPerspectives(conflict, team);
    
    // 3. 评估解决方案
    const solutions = await this.generateSolutions(conflictType, perspectives);
    
    // 4. 团队投票
    const vote = await this.teamVote(solutions, team);
    
    // 5. 执行解决方案
    return await this.implementSolution(vote.winningSolution);
  }
}
```

## 质量保证

### 迭代审查

```yaml
iteration_review:
  frequency: "每周"
  participants: ["phase_lead", "team_members", "po"]
  
  review_items:
    - "任务完成情况"
    - "质量指标达成"
    - "团队协作效果"
    - "风险评估更新"
    - "下一步计划调整"
    
  outputs:
    - "周报.md"
    - "风险清单.md"
    - "调整计划.md"
```

### 适应性评估

```typescript
interface AdaptabilityAssessment {
  assessAdaptability(
    project: Project,
    currentPlan: TaskExecutionPlan
  ): Promise<AdaptabilityReport> {
    
    // 1. 计划灵活性评估
    const planFlexibility = this.assessPlanFlexibility(currentPlan);
    
    // 2. 团队适应性评估
    const teamAdaptability = this.assessTeamAdaptability(project.team);
    
    // 3. 资源弹性评估
    const resourceElasticity = this.assessResourceElasticity(project.resources);
    
    // 4. 风险承受能力评估
    const riskTolerance = this.assessRiskTolerance(project);
    
    return {
      planFlexibility,
      teamAdaptability,
      resourceElasticity,
      riskTolerance,
      overallScore: this.calculateOverallScore([
        planFlexibility,
        teamAdaptability,
        resourceElasticity,
        riskTolerance
      ])
    };
  }
}
```

## 监控和报告

### 实时监控

```bash
# 查看当前任务状态
free-mode status --project "ai-explorer" --real-time

# 查看团队工作负载
free-mode workload --project "ai-explorer" --by-role

# 查看任务依赖关系
free-mode dependencies --project "ai-explorer" --graph
```

### 进度报告

```bash
# 生成周报
free-mode report --project "ai-explorer" --period "week" --format markdown

# 生成里程碑报告
free-mode milestone-report --project "ai-explorer" --milestone "poc-validation"

# 生成团队绩效报告
free-mode team-performance --project "ai-explorer" --detailed
```

### 关键指标

- **任务完成率**: 实际完成任务 vs 计划任务
- **计划调整频率**: 计划调整的次数和原因
- **团队协作效率**: 团队成员之间的协作效果
- **适应性评分**: 项目应对变化的能力

## 配置管理

### 全局配置

```yaml
# .goagents/config/free-mode.yaml
free_mode:
  authority_levels:
    phase_lead:
      can_add_tasks: true
      can_remove_tasks: true
      can_adjust_priorities: true
      can_reassign_resources: true
      
    team_member:
      can_suggest_tasks: true
      can_report_issues: true
      can_request_help: true
      
  decision_making:
    consensus_required: false
    vote_threshold: 0.6
    conflict_timeout: "48 hours"
    
  adaptation_rules:
    max_plan_changes_per_week: 5
    max_task_reassignments: 3
    emergency_override: true
```

### 项目配置

```yaml
# .goagents/projects/ai-explorer/free-mode.yaml
project_config:
  phase_lead: "architect"
  decision_frequency: "daily"
  adaptation_threshold: 0.7
  
  custom_rules:
    - rule: "技术探索任务可以优先级提升"
      condition: "发现关键技术突破"
      action: "自动提升优先级"
      
    - rule: "用户反馈需要立即响应"
      condition: "收到负面用户反馈"
      action: "24小时内必须处理"
```

## 集成接口

### 与 PO Core 集成

```typescript
interface FreeModeIntegration {
  generateDynamicPlan(
    milestone: Milestone,
    teamConfig: TeamConfiguration
  ): Promise<TaskExecutionPlan>;
  
  adjustPlan(
    projectId: string,
    adjustments: PlanAdjustment[]
  ): Promise<TaskExecutionPlan>;
  
  handleConflict(
    conflict: TaskConflict
  ): Promise<Resolution>;
}
```

### 与团队角色集成

```typescript
interface TeamRoleIntegration {
  assignPhaseLead(
    projectId: string,
    phaseName: string,
    agentId: string
  ): Promise<void>;
  
  collectTeamFeedback(
    projectId: string,
    topic: string
  ): Promise<TeamFeedback[]>;
  
  conductTeamVote(
    projectId: string,
    options: VoteOption[]
  ): Promise<VoteResult>;
}
```

## 扩展性

### 自定义决策流程

支持自定义决策流程：

1. 定义决策步骤
2. 设置参与者角色
3. 配置投票规则
4. 注册到系统

### 自定义适应性规则

支持添加自定义适应性规则：

1. 定义触发条件
2. 设置响应动作
3. 配置执行参数
4. 集成到系统

### 自定义冲突解决机制

支持添加自定义冲突解决机制：

1. 定义冲突类型
2. 设置解决策略
3. 配置参与规则
4. 注册到系统
