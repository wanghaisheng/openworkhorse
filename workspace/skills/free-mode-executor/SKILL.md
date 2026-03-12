---
name: free-mode-executor
description: "Free模式执行AI专家 - 专门负责Phase Lead主导的动态规划和灵活探索执行的专业技能"
---

# Free模式执行AI专家

## 🎯 技能角色

### 核心职责
- **动态规划**: Phase Lead主导的灵活项目规划
- **探索执行**: 支持迭代开发和原型驱动的执行方式
- **灵活调整**: 根据项目进展动态调整计划和策略
- **创新导向**: 鼓励创新和实验性的解决方案

### 专业能力
- **灵活规划**: 掌握动态项目规划方法
- **迭代开发**: 支持敏捷和原型驱动的开发流程
- **风险管理**: 在不确定性中识别和管理风险
- **创新思维**: 鼓励创新和非常规解决方案

## 🔧 功能实现

### 与ExecutionFramework集成
```yaml
execution_framework_integration:
  free_mode_configuration:
    method: "execution_framework.LoadPhaseConfig(phaseID)"
    focus: "free_mode_config"
    phase_lead_authority: "dynamic"
    iteration_support: true
    
  dynamic_planning:
    method: "registry.ConfigManager.LoadPhaseConfig()"
    config_section: "free_mode_config"
    adaptive_planning: true
    iteration_support: true
    
  official_types:
    phase_config: "registry.PhaseConfig"
    free_mode_config: "registry.FreeModeConfig"
    team_config: "registry.TeamConfig"
```

### 输入格式
```yaml
input:
  project_request: object       # 项目请求信息
  phase_lead: object           # Phase Lead信息
  free_mode_config: object     # Free模式配置
  iteration_count: int         # 迭代次数
  exploration_scope: object    # 探索范围
  innovation_level: string      # 创新程度要求
  use_official_registry: boolean # 使用官方registry
```

### 输出格式
```yaml
output:
  dynamic_plan: object         # 动态执行计划
  iteration_plan: object[]     # 迭代计划
  exploration_results: object[] # 探索结果
  innovation_suggestions: object[] # 创新建议
  risk_assessment: object      # 风险评估
  next_actions: object[]       # 下一步行动
```

## 📋 Free模式定义

### 模式特征
```yaml
free_mode_characteristics:
  execution_style: "dynamic"
  flexibility: "high"
  predictability: "low"
  quality_control: "adaptive"
  documentation: "minimal"
  risk_management: "reactive"
  innovation_focus: "high"
```

### 适用场景
```yaml
applicable_scenarios:
  project_types:
    - "AI研究和开发"
    - "创新产品开发"
    - "技术原型验证"
    - "探索性项目"
    - "前沿技术研究"
    
  team_characteristics:
    - "小规模团队（2-8人）"
    - "高技能专业团队"
    - "创新导向文化"
    - "高度自主性"
    
  requirements:
    - "需求不明确或需要探索"
    - "技术方案需要创新"
    - "允许失败和重试"
    - "时间和资源相对充裕"
```

## 🎯 使用示例

### 完整Free模式执行
```bash
@go "使用free-mode-executor技能执行Free模式项目：
项目：AI聊天机器人研发
Phase Lead：AI专家
迭代次数：3次
创新程度：high
探索范围：NLP技术、多模态交互
时间：4个月探索期
允许失败：true"
```

### 原型驱动开发
```bash
@go "使用free-mode-executor技能进行原型驱动开发：
项目：区块链投票系统
Phase Lead：区块链专家
原型类型：MVP原型
迭代周期：2周
验证重点：技术可行性
创新要求：共识机制创新"
```

## 🔄 动态规划机制

### Phase Lead主导规划
```yaml
phase_lead_authority:
  task_planning: true
  task_assignment: true
  resource_allocation: true
  technical_decisions: true
  priority_setting: true
  iteration_control: true
  
  decision_scope:
    - "技术方案选择"
    - "开发优先级"
    - "资源分配"
    - "迭代计划"
    - "质量标准"
```

### 迭代开发流程
```yaml
iteration_process:
  1. "探索阶段": 需求探索和技术调研
  2. "原型阶段": 快速原型开发
  3. "验证阶段": 原型验证和反馈收集
  4. "调整阶段": 根据反馈调整方向
  5. "深化阶段": 深入开发和优化"
  6. "评估阶段": 成果评估和下一步规划"
```

### 动态调整机制
```yaml
dynamic_adjustment:
  triggers:
    - "技术障碍出现"
    - "需求重大变更"
    - "市场环境变化"
    - "团队资源变化"
    - "创新机会发现"
    
  adjustment_types:
    - "技术方案调整"
    - "优先级重新排序"
    - "资源重新分配"
    - "时间计划调整"
    - "团队结构调整"
```

## 🔍 Free模式配置

### 官方registry配置
```yaml
# registry.PhaseConfig.FreeModeConfig
free_mode_config:
  phase_lead_authority:
    task_planning: true
    task_assignment: true
    resource_allocation: true
    technical_decisions: true
    priority_setting: true
    iteration_control: true
    
  iteration_config:
    default_duration: "2周"
    max_iterations: 5
    success_criteria: "prototype_validation"
    failure_tolerance: "medium"
    
  exploration_config:
    research_methods: ["prototype", "proof_of_concept", "experiment"]
    innovation_encouragement: "high"
    risk_tolerance: "medium"
    resource_allocation: "flexible"
    
  quality_config:
    adaptive_standards: true
    focus_on: ["functionality", "innovation", "user_feedback"]
    documentation_level: "minimal"
    testing_approach: "exploratory"
```

### 动态规划模板
```yaml
dynamic_planning_template:
  iteration_planning:
    objectives: "明确的迭代目标"
    success_metrics: "可衡量的成功指标"
    risk_assessment: "风险识别和应对"
    resource_allocation: "资源分配计划"
    
  exploration_framework:
    hypothesis: "需要验证的假设"
    experiment_design: "实验设计"
    success_criteria: "成功标准"
    failure_analysis: "失败分析框架"
    
  innovation_tracking:
    innovation_opportunities: "创新机会识别"
    patent_potential: "专利潜力评估"
    market_differentiation: "市场差异化"
    technical_advancement: "技术进步"
```

## 🎯 使用示例

### AI研究项目执行
```bash
@go "使用free-mode-executor技能执行AI研究项目：
项目：多模态AI助手研发
Phase Lead：AI研究专家
Free模式配置：
  phase_lead_authority: full
  iteration_config:
    default_duration: 3周
    max_iterations: 4
  exploration_config:
    research_methods: [prototype, experiment]
    innovation_encouragement: high

# 自动执行流程：
# 1. 加载官方registry.FreeModeConfig
# 2. Phase Lead主导动态规划
# 3. 执行迭代开发流程
# 4. 动态调整和优化
# 5. 生成创新建议和下一步行动
```

### 技术原型验证
```bash
@go "使用free-mode-executor技能进行技术原型验证：
项目：量子计算加密原型
Phase Lead：量子计算专家
原型目标：验证量子加密可行性
迭代周期：1周
验证重点：技术可行性
创新要求：算法创新

# 输出：
# - 动态原型开发计划
# - 迭代验证方案
# - 技术风险评估
# - 创新改进建议
# - 下一步研究方向
```

## 📊 质量标准

### Free模式质量指标
- **创新度**: > 80%（创新性评分）
- **原型验证率**: > 70%（原型成功验证）
- **迭代效率**: > 85%（迭代目标达成）
- **技术突破**: > 60%（技术突破程度）

### 动态规划质量
- **规划灵活性**: > 90%（规划调整能力）
- **响应速度**: < 24小时（问题响应时间）
- **资源利用率**: > 80%（资源使用效率）
- **风险控制**: > 85%（风险识别和控制）

## 🔄 与其他模式对比

### vs Standard模式
```yaml
standard_mode:
  planning: "模板化、预定义"
  execution: "严格按计划"
  quality: "标准化质量控制"
  documentation: "完整文档"
  
free_mode:
  planning: "动态、探索性"
  execution: "灵活迭代"
  quality: "适应性质量控制"
  documentation: "最小必要文档"
```

### vs Hybrid模式
```yaml
free_mode:
  flexibility: "最高"
  structure: "最小"
  innovation: "最高"
  risk: "最高"
  
hybrid_mode:
  flexibility: "中等"
  structure: "平衡"
  innovation: "中等"
  risk: "中等"
```

## 🚀 集成方式

### 与po-core-v2集成
```yaml
po_core_integration:
  mode_selection:
    condition: "high_innovation || unclear_requirements"
    mode: "free"
    executor: "free-mode-executor"
    
  execution_flow:
    1. "po-core-v2调用mode-selector选择free模式"
    2. "po-core-v2调用free-mode-executor执行项目"
    3. "free-mode-executor使用官方registry配置"
    4. "Phase Lead主导动态规划和迭代执行"
    5. "返回创新成果和下一步建议"
```

### 与registry集成
```yaml
registry_integration:
  config_loading:
    phase_config: "registry.ConfigManager.LoadPhaseConfig()"
    free_mode_config: "registry.PhaseConfig.FreeModeConfig"
    
  type_compatibility:
    phase_config: "registry.PhaseConfig"
    team_config: "registry.TeamConfig"
    system_config: "registry.SystemConfig"
    
  official_support:
    cli_tools: "goagents phase update"
    config_validation: "官方验证规则"
    documentation: "官方文档格式"
```

---

这个技能专门负责Free模式的动态规划和灵活执行，充分利用官方registry配置，支持Phase Lead主导的创新和探索性项目开发。
