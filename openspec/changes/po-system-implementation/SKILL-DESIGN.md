# PO System Skill Design - 技能设计文档

## 概述

本文档详细描述了基于 PicoClaw 技能系统的 PO（产品经理）系统完整设计，包括核心技能、任务模式、团队角色和工业化集成。

## 🏗️ 完整层级结构设计

### 核心模型层级
基于核心模型实现完整的层级结构：
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
- **Workflow**: 项目整体流程和阶段定义
- **Phase**: 项目的主要阶段（discovery, planning, development, validation）
- **Milestone**: 阶段内的关键里程碑，挂载质量门禁
- **Task**: 具体的执行任务，包含内部子任务拆解
- **Team**: 执行任务的团队配置
- **Team Role**: 团队中的角色定义
- **Team Member Agent**: 具体的执行者

## 🎭 技能内部任务拆解

### 任务分解策略
每个技能都支持三种任务分解策略：

#### 1. Template-based（模板化）
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
      - id: "market_size_analysis"
        name: "市场规模数据收集"
        estimated_time: "1-2天"
        dependencies: ["data_collection_preparation"]
```

#### 2. Milestone-driven（里程碑驱动）
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
      - id: "team_coordination_setup"
        name: "团队协调设置"
        estimated_time: "1天"
```

#### 3. Hybrid（混合模式）
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

### 技能设计原则
- 完全基于现有技能系统实现
- 不修改 PicoClaw 核心代码
- 保持超轻量特性

### 2. 模块化设计
- 每个功能独立为技能
- 支持动态加载和卸载
- 便于维护和扩展

### 3. 工业化标准
- 集成 Harness Engineering 最佳实践
- 支持 Ralph Wiggum Loop
- 内置质量门禁

## 核心技能架构

### PO Core 技能

```yaml
---
name: po-core
description: "产品经理核心 - 项目全生命周期管理"
category: "orchestrator"
version: "1.0.0"
requires:
  skills: ["phase-manager", "task-modes", "team-roles"]
metadata: {
  "trigger": "@go",
  "harness_version": "2026.03",
  "supported_modes": ["standard", "free", "hybrid"]
}
---

# PO Core Skill

## 核心职责
1. **需求分析**: 接收用户输入，分析任务类型和复杂度
2. **模式选择**: 根据项目特征选择合适的执行模式
3. **团队组建**: 动态组建合适的团队配置
4. **阶段管理**: 管理项目阶段转换和质量门禁
5. **协作协调**: 协调多 Agent 协作执行

## HARNESS.md 集成
每个任务开始前强制执行：
```bash
First, read and strictly follow HARNESS.md in repository root.
Do NOT deviate from its rules.
```

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

## 决策算法
```typescript
class PODecisionEngine {
  async analyzeAndPlan(userInput: string): Promise<ProjectPlan> {
    // 1. 需求分析
    const analysis = await this.analyzeRequirement(userInput);
    
    // 2. 模式推荐
    const recommendedMode = await this.recommendMode(analysis);
    
    // 3. 团队配置
    const teamConfig = await this.buildTeam(analysis, recommendedMode);
    
    // 4. 阶段规划
    const phasePlan = await this.planPhases(analysis);
    
    return {
      analysis,
      mode: recommendedMode,
      team: teamConfig,
      phases: phasePlan
    };
  }
}
```

## 输出格式
```markdown
📋 **PO任务分析结果**

**需求**: {需求描述}
**类型**: {任务类型}
**复杂度**: {复杂度评分}
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
```

### Phase Manager 技能

```yaml
---
name: phase-manager
description: "阶段管理器 - 管理项目阶段转换"
category: "coordinator"
version: "1.0.0"
requires:
  skills: ["phase-templates"]
---

# Phase Manager

## 支持的阶段
- **discovery**: 需求发现和市场调研
- **planning**: 架构设计和实施计划
- **development**: 功能开发和实现
- **validation**: 测试验证和质量保证

## 阶段转换规则
```yaml
phase_transitions:
  discovery_to_planning:
    condition: "requirements_frozen AND market_analysis_approved"
    quality_gates: ["po_approval", "stakeholder_signoff"]
    automated_checks: ["requirement_completeness", "market_data_validation"]
    
  planning_to_development:
    condition: "architecture_approved AND implementation_plan_ready"
    quality_gates: ["technical_review", "resource_allocation"]
    automated_checks: ["architecture_validation", "plan_feasibility"]
    
  development_to_validation:
    condition: "feature_implementation_complete AND unit_tests_passing"
    quality_gates: ["code_review", "integration_test"]
    automated_checks: ["test_coverage", "code_quality", "performance_baseline"]
```

## 阶段状态跟踪
```typescript
class PhaseManager {
  private currentPhase: ProjectPhase;
  private phaseHistory: PhaseTransition[];
  private qualityGateResults: QualityGateResult[];
  
  async transitionToPhase(
    targetPhase: ProjectPhase,
    reason: string
  ): Promise<TransitionResult> {
    
    // 1. 验证转换条件
    const canTransition = await this.validateTransition(targetPhase);
    if (!canTransition.valid) {
      return { success: false, reason: canTransition.reason };
    }
    
    // 2. 执行质量门禁
    const gateResults = await this.executeQualityGates(targetPhase);
    if (!gateResults.allPassed) {
      return { success: false, reason: "Quality gates failed" };
    }
    
    // 3. 执行阶段转换
    const transition = await this.performTransition(targetPhase, reason);
    
    // 4. 更新状态
    this.updatePhaseState(transition);
    
    return { success: true, transition };
  }
}
```

## 阶段模板系统
每个阶段都有对应的模板，定义标准的工作流程和交付物：

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
```

## 任务模式技能设计

### Standard Mode 技能

```yaml
---
name: standard-mode
description: "标准化任务执行模式"
category: "execution"
version: "1.0.0"
requires:
  skills: ["task-templates"]
---

# Standard Mode Execution

## 模板执行机制
当 Phase Lead 分配任务时，Agent 严格按照模板步骤执行：

```typescript
class TemplateAgentExecutor {
  private assignedTask: TaskFromTemplate;
  private templateSteps: TemplateStep[];
  private executionProgress: StepProgress[];
  
  async executeTemplateTask(task: TaskFromTemplate): Promise<TaskResult> {
    // 1. Agent 解析模板任务
    const templateAnalysis = await this.parseTemplateTask(task);
    
    // 2. 加载对应的执行模板
    const executionTemplate = await this.loadExecutionTemplate(task.templateId);
    
    // 3. 按模板步骤顺序执行
    const results = await this.executeTemplateSteps(executionTemplate);
    
    // 4. 验证模板执行完整性
    const validation = await this.validateTemplateExecution(results);
    
    // 5. 生成标准化输出
    return await this.generateStandardOutputs(results);
  }
}
```

## 模板执行示例
```yaml
# 市场分析任务模板执行
template_execution:
  task: "market-analysis"
  agent: "market_analyst"
  template_version: "v2.1"
  deadline: "1天"
  deliverables: ["market_size_report.xlsx"]
  
  template_steps:
    step_1:
      name: "数据收集准备"
      description: "准备市场规模评估所需的数据源和工具"
      estimated_time: "2小时"
      inputs: ["项目背景", "目标市场定义"]
      outputs: ["数据源清单.md", "工具准备清单.xlsx"]
      quality_gates: ["数据源完整性", "工具可用性"]
      
    step_2:
      name: "市场规模数据收集"
      description: "从多个数据源收集市场规模相关数据"
      estimated_time: "4小时"
      inputs: ["数据源清单.md", "工具准备清单.xlsx"]
      outputs: ["原始数据.xlsx", "数据收集日志.md"]
      depends_on: ["step_1"]
      quality_gates: ["数据完整性", "数据准确性"]
      
    step_3:
      name: "数据清洗和整理"
      description: "清洗收集的原始数据，整理成标准格式"
      estimated_time: "3小时"
      inputs: ["原始数据.xlsx", "数据收集日志.md"]
      outputs: ["清洗后数据.xlsx", "数据清洗记录.md"]
      depends_on: ["step_2"]
      quality_gates: ["数据一致性", "格式标准化"]
```

## 严格遵循原则
```yaml
template_execution_discipline:
  step_sequence:
    mandatory: true  # 必须按模板顺序执行
    skip_requires_approval: true  # 跳过需要特殊批准
    
  io_standards:
    input_validation: true  # 严格验证输入
    output_format: "standard"  # 标准化输出格式
    template_compliance: true  # 确保符合模板要求
    
  quality_gates:
    mandatory: true  # 必须通过所有质量门禁
    auto_validation: true  # 自动化验证
    human_review: "phase_lead"  # 阶段负责人审核
```
```

### Free Mode 技能

```yaml
---
name: free-mode
description: "自由模式 - Phase Lead主导的动态任务规划"
category: "execution"
version: "1.0.0"
---

# Free Mode Execution

## Phase Lead 权限配置
```yaml
phase_lead_authority:
  task_analysis:
    - "分析任务复杂度和范围"
    - "确定任务优先级"
    - "识别关键路径"
    - "评估资源需求"
    
  task_planning:
    - "制定任务执行计划"
    - "分解大任务为小任务"
    - "设置任务依赖关系"
    - "安排任务执行顺序"
    
  task_assignment:
    - "分配任务给团队成员"
    - "协调团队协作"
    - "调整任务分配"
    - "解决资源冲突"
    
  task_adjustment:
    - "动态调整任务优先级"
    - "修改任务执行计划"
    - "添加或删除任务"
    - "调整任务时间估算"
```

## 动态任务生成流程
```typescript
class DynamicTaskGenerator {
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
```yaml
# 自由模式：项目启动阶段
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
      
    task_2:
      name: "关键利益相关者识别和访谈"
      agent: "phase_lead"
      reasoning: "项目涉及多个部门，需要识别关键决策者并了解他们的期望"
      estimated_time: "2-3天"
      deliverables: ["利益相关者清单.md", "访谈记录.md"]
      depends_on: ["task_1"]
      
    task_3:
      name: "需求范围和边界确定"
      agent: "phase_lead"
      reasoning: "明确项目范围，避免范围蔓延，确保项目可交付"
      estimated_time: "1天"
      deliverables: ["需求边界文档.md", "范围说明书.md"]
      depends_on: ["task_2"]
      
    task_4:
      name: "市场调研和用户研究"
      agent: "market_analyst" + "user_researcher"
      reasoning: "需要团队协作收集市场数据和用户反馈"
      estimated_time: "1-2周"
      deliverables: ["市场分析报告.xlsx", "用户研究报告.pdf"]
      depends_on: ["task_3"]
```
```

## 团队角色技能设计

### Role Analyst 技能

```yaml
---
name: role-analyst
description: "需求分析师角色"
category: "role"
version: "1.0.0"
persona: |
  我是专业需求分析师，具有5年电商产品经验。
  我擅长需求分解、市场调研、竞品分析。
  我的输出总是结构化、可执行的。
  
  我的工作方式：
  - 深度理解用户需求
  - 系统性分析市场环境
  - 数据驱动的决策支持
  - 清晰的文档输出

---

# Analyst Role

## 专业能力矩阵
```yaml
capabilities:
  market_research:
    level: "expert"
    tools: ["行业报告分析", "竞品功能对比", "市场规模评估"]
    deliverables: ["市场分析报告", "竞品对比表", "趋势预测"]
    
  user_research:
    level: "expert"
    tools: ["用户访谈", "问卷调查", "用户画像构建"]
    deliverables: ["用户研究报告", "用户画像", "用户旅程图"]
    
  requirement_analysis:
    level: "expert"
    tools: ["需求分解", "功能规格定义", "优先级排序"]
    deliverables: ["需求文档", "功能清单", "验收标准"]
    
  data_analysis:
    level: "advanced"
    tools: ["统计分析", "数据可视化", "洞察提取"]
    deliverables: ["数据分析报告", "洞察总结", "建议方案"]
```

## 工作流程
```yaml
workflow:
  step_1:
    name: "需求理解"
    activities: ["分析用户输入", "明确目标", "识别约束条件"]
    outputs: ["需求理解文档"]
    
  step_2:
    name: "信息收集"
    activities: ["市场调研", "用户访谈", "竞品分析"]
    outputs: ["调研数据", "访谈记录", "竞品资料"]
    
  step_3:
    name: "分析整合"
    activities: ["数据整理", "模式识别", "洞察提取"]
    outputs: ["分析报告", "关键发现"]
    
  step_4:
    name: "输出生成"
    activities: ["文档编写", "图表制作", "建议提出"]
    outputs: ["最终报告", "行动建议"]
```

## 输出模板
```markdown
📊 **需求分析结果**

**核心功能**:
- 功能1: {描述}
- 功能2: {描述}
- 功能3: {描述}

**技术考虑**:
- 架构要求: {要求}
- 性能考虑: {考虑}
- 集成需求: {需求}

**竞品分析**:
- 主要竞品: {竞品列表}
- 功能对比: {对比表}
- 差异化机会: {机会}

**输出文件**:
- brief.md: 需求简报
- research-findings.md: 调研发现
- user-stories.md: 用户故事
```

## 质量标准
```yaml
quality_standards:
  analysis_depth:
    description: "分析必须覆盖市场、用户、技术三个维度"
    threshold: "100%"
    
  data_support:
    description: "所有结论必须有数据或事实支撑"
    threshold: "95%"
    
  actionability:
    description: "输出必须能指导后续开发"
    threshold: "90%"
    
  clarity:
    description: "文档必须清晰易懂，结构化"
    threshold: "95%"
```
```

## 技能注册表系统

### 技能发现器 (Discovery)
```go
type Discovery interface {
    DiscoverLocalSkills() ([]Skill, error)
    DiscoverRemoteSkills(repository string) ([]Skill, error)
    UpdateIndex() error
    ValidateSkill(skill Skill) error
}

type LocalDiscovery struct {
    searchPaths []string
    cache       map[string]*Skill
}

func (d *LocalDiscovery) DiscoverLocalSkills() ([]Skill, error) {
    var skills []Skill
    
    for _, searchPath := range d.searchPaths {
        err := filepath.Walk(searchPath, func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }
            
            if info.IsDir() && filepath.Base(path) == "SKILL.md" {
                skill, err := d.parseSkillFile(path)
                if err != nil {
                    return err
                }
                skills = append(skills, *skill)
            }
            
            return nil
        })
        
        if err != nil {
            return nil, err
        }
    }
    
    return skills, nil
}
```

### 技能管理器 (Manager)
```go
type Manager interface {
    InstallSkill(skillID string, version string) error
    UninstallSkill(skillID string) error
    UpdateSkill(skillID string) error
    ListInstalledSkills() ([]Skill, error)
    GetSkillInfo(skillID string) (*Skill, error)
    ResolveDependencies(skillID string) ([]Dependency, error)
}

type SkillManager struct {
    installPath string
    registry    *SkillRegistry
}

func (m *SkillManager) InstallSkill(skillID string, version string) error {
    // 1. 下载技能
    skillData, err := m.downloadSkill(skillID, version)
    if err != nil {
        return err
    }
    
    // 2. 解压技能
    err = m.extractSkill(skillData, skillID)
    if err != nil {
        return err
    }
    
    // 3. 验证技能
    err = m.validateSkill(skillID)
    if err != nil {
        return err
    }
    
    // 4. 更新注册表
    err = m.registry.RegisterSkill(skillID)
    if err != nil {
        return err
    }
    
    return nil
}
```

### 技能搜索器 (Searcher)
```go
type Searcher interface {
    SearchSkills(query SearchQuery) ([]*SearchResult, error)
    GetSkillsByCategory(category string) ([]Skill, error)
    GetSkillsByTag(tag string) ([]Skill, error)
    GetRecommendedSkills(context string) ([]Skill, error)
    AutoComplete(prefix string) ([]string, error)
}

type SearchQuery struct {
    Query     string   `json:"query"`
    Category  string   `json:"category"`
    Tags      []string `json:"tags"`
    Author    string   `json:"author"`
    MinScore  int      `json:"min_score"`
    Limit     int      `json:"limit"`
    SortBy    string   `json:"sort_by"`    // name, score, updated
    SortOrder string   `json:"sort_order"` // asc, desc
}

func (s *SkillSearcher) SearchSkills(query SearchQuery) ([]*SearchResult, error) {
    var results []*SearchResult
    
    // 1. 从索引中获取候选技能
    candidates := s.index.GetCandidates(query)
    
    // 2. 计算匹配分数
    for _, candidate := range candidates {
        score := s.calculateScore(candidate, query)
        if score >= float64(query.MinScore) {
            result := &SearchResult{
                Skill: candidate,
                Score: score,
                Matches: s.getMatches(candidate, query),
            }
            results = append(results, result)
        }
    }
    
    // 3. 排序结果
    s.sortResults(results, query.SortBy, query.SortOrder)
    
    return results, nil
}
```

### 技能验证器 (Validator)
```go
type Validator interface {
    ValidateSkillFormat(skillPath string) (*ValidationResult, error)
    ValidateSkillSpec(skill Skill) (*ValidationResult, error)
    RunSkillTests(skillID string) (*TestResult, error)
    CalculateQualityScore(skill Skill) (int, error)
}

type ValidationResult struct {
    Valid    bool     `json:"valid"`
    Errors   []string `json:"errors"`
    Warnings []string `json:"warnings"`
    Score    int      `json:"score"`
    Issues   []Issue  `json:"issues"`
}

func (v *SkillValidator) ValidateSkillSpec(skill *Skill) (*ValidationResult, error) {
    result := &ValidationResult{
        Valid:    true,
        Errors:   []string{},
        Warnings: []string{},
        Issues:   []Issue{},
        Score:    100,
    }
    
    // 1. 基础格式验证
    for _, rule := range v.rules {
        err := rule.Validator(skill)
        if err != nil {
            result.Valid = false
            result.Errors = append(result.Errors, err.Error())
            result.Score -= rule.Severity
        }
    }
    
    // 2. 计算质量分数
    result.Score = v.calculateQualityScore(skill)
    
    return result, nil
}
```

## 技能导入转换器

### 格式检测器 (FormatDetector)
```go
type FormatDetector interface {
    DetectFormat(filePath string) (Format, error)
    GetSupportedFormats() []Format
    ValidateFormat(filePath string) error
}

type AgencyFormatDetector struct{}

func (d *AgencyFormatDetector) DetectFormat(filePath string) (Format, error) {
    content, err := ioutil.ReadFile(filePath)
    if err != nil {
        return "", err
    }
    
    contentStr := string(content)
    
    // 检查 Front Matter
    if strings.HasPrefix(contentStr, "---\n") {
        parts := strings.SplitN(contentStr, "---\n", 3)
        if len(parts) < 3 {
            return "", errors.New("invalid front matter format")
        }
        
        fmContent := parts[1]
        
        // 检查 Agency Agents 特定字段
        if strings.Contains(fmContent, "name:") && 
           strings.Contains(fmContent, "description:") &&
           strings.Contains(fmContent, "color:") &&
           strings.Contains(fmContent, "emoji:") {
            return FormatAgencyAgents, nil
        }
    }
    
    return FormatUnknown, errors.New("unsupported format")
}
```

### 内容解析器 (ContentParser)
```go
type ContentParser interface {
    ParseFrontMatter(content string) (map[string]interface{}, error)
    ParseSections(content string) (map[string]string, error)
    ExtractCapabilities(content string) ([]string, error)
    ExtractTriggers(content string) ([]string, error)
}

type AgencyParser struct{}

func (p *AgencyParser) ParseFrontMatter(content string) (map[string]interface{}, error) {
    parts := strings.SplitN(content, "---\n", 3)
    if len(parts) < 3 {
        return nil, errors.New("invalid front matter format")
    }
    
    fmContent := parts[1]
    var frontMatter map[string]interface{}
    
    err := yaml.Unmarshal([]byte(fmContent), &frontMatter)
    if err != nil {
        return nil, err
    }
    
    return frontMatter, nil
}

func (p *AgencyParser) ExtractCapabilities(content string) ([]string, error) {
    sections, err := p.ParseSections(content)
    if err != nil {
        return nil, err
    }
    
    var capabilities []string
    
    // 从 "Your Core Mission" 章节提取能力
    if missionSection, exists := sections["Your Core Mission"]; exists {
        // 提取项目符号点
        bulletRegex := regexp.MustCompile(`^\s*-\s+(.+)$`)
        matches := bulletRegex.FindAllStringSubmatch(missionSection, -1)
        
        for _, match := range matches {
            if len(match) >= 2 {
                capabilities = append(capabilities, strings.TrimSpace(match[1]))
            }
        }
    }
    
    return capabilities, nil
}
```

### 格式转换器 (FormatConverter)
```go
type FormatConverter interface {
    ConvertToSkill(parsed ParsedContent) (*Skill, error)
    GenerateVariants(skill *Skill) ([]Variant, error)
    ConvertExecutionConfig(parsed ParsedContent) (*ExecutionConfig, error)
}

type AgencyToPicoClawConverter struct {
    config *ConversionConfig
}

func (c *AgencyToPicoClawConverter) ConvertToSkill(parsed *ParsedContent) (*Skill, error) {
    skill := &Skill{
        ID:          c.generateSkillID(parsed.Metadata.Name),
        Name:        parsed.Metadata.Name,
        Description: parsed.Metadata.Description,
        Category:    c.config.TargetCategory,
        Version:     "1.0.0",
        License:     "MIT",
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }
    
    // 转换能力
    skill.Capabilities = c.convertCapabilities(parsed.Metadata.Capabilities)
    
    // 转换触发条件
    skill.Triggers = c.convertTriggers(parsed.Metadata.Triggers)
    
    // 生成元数据
    skill.Metadata = c.generateMetadata(parsed)
    
    return skill, nil
}
```

## 技能加载和管理

### 技能注册系统
```typescript
class SkillRegistry {
  private skills = new Map<string, Skill>();
  private dependencies = new Map<string, string[]>();
  
  async registerSkill(skillPath: string): Promise<void> {
    // 1. 加载技能定义
    const skill = await this.loadSkill(skillPath);
    
    // 2. 验证技能格式
    this.validateSkill(skill);
    
    // 3. 检查依赖关系
    await this.checkDependencies(skill);
    
    // 4. 注册技能
    this.skills.set(skill.name, skill);
    
    // 5. 更新索引
    await this.updateSkillIndex(skill);
  }
  
  async loadSkillsForMode(mode: string): Promise<Skill[]> {
    return Array.from(this.skills.values())
      .filter(skill => skill.supportedModes?.includes(mode));
  }
}
```

### 技能版本管理
```yaml
skill_versioning:
  format: "semantic_versioning"  # major.minor.patch
  compatibility:
    backward_compatible: true  # 向后兼容
    migration_support: true     # 支持迁移
    
  version_strategy:
    major: "破坏性变更"
    minor: "新功能添加"
    patch: "bug修复和优化"
```

## 质量保证

### 技能测试框架
```typescript
class SkillTestFramework {
  async testSkill(skillName: string): Promise<TestResult> {
    const tests = [
      this.testSkillLoading(skillName),
      this.testDependencyResolution(skillName),
      this.testExecutionFlow(skillName),
      this.testOutputFormat(skillName),
      this.testErrorHandling(skillName)
    ];
    
    const results = await Promise.all(tests);
    return this.aggregateTestResults(results);
  }
}
```

### 性能监控
```yaml
performance_metrics:
  skill_loading_time:
    target: "<2秒"
    alert_threshold: "5秒"
    
  execution_latency:
    target: "<5秒"
    alert_threshold: "10秒"
    
  memory_usage:
    target: "<50MB per skill"
    alert_threshold: "100MB"
```

## 部署和维护

### 技能部署流程
```bash
# 1. 技能开发完成
picoclaw skills validate po-core
picoclaw skills test po-core

# 2. 技能发布
picoclaw skills publish po-core --version "1.0.0"

# 3. 技能安装
picoclaw skills install po-core --version "latest"

# 4. 技能启用
picoclaw skills enable po-core
```

### 技能更新机制
```yaml
update_mechanism:
  auto_update: true  # 自动检查更新
  rollback_support: true  # 支持回滚
  dependency_management: true  # 自动管理依赖
  
  update_channels:
    - "stable": 稳定版本
    - "beta": 测试版本
    - "dev": 开发版本
```

## 总结

这个技能设计文档提供了：

1. **完整的 PO 系统架构** - 从核心功能到具体实现
2. **三种任务模式** - Standard/Free/Hybrid 完整设计
3. **团队角色体系** - 详细的角色定义和能力矩阵
4. **工业化集成** - Harness Engineering 最佳实践
5. **质量保证机制** - 测试、监控、版本管理

通过这个设计，PicoClaw 将升级为工业级的多 Agent 协作平台！
