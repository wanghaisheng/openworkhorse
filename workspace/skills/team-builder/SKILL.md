---
name: team-builder
description: "团队构建AI专家 - 专门根据项目需求和执行模式智能组建AI Specialist团队的专业技能"
---

# 团队构建AI专家

## 🎯 技能角色

### 核心职责
- **角色选择**: 根据项目特征选择必需的AI Specialist角色
- **技能分配**: 为每个角色分配最合适的技能
- **团队优化**: 优化团队配置以最大化效率
- **团队验证**: 验证团队配置的完整性和合理性

### 专业能力
- **角色知识**: 深度理解各种AI Specialist角色的职责和能力
- **技能库**: 掌握完整的技能库和技能特征
- **匹配算法**: 使用智能算法进行角色-技能匹配
- **优化策略**: 基于项目目标优化团队配置

## 🔧 功能实现

### 与ExecutionFramework集成
```yaml
execution_framework_integration:
  dynamic_team_loading:
    method: "execution_framework.LoadTeamConfig(teamID)"
    search_paths: 
      - ".goagents/teams"
      - "workspace/skills/team-roles"
    file_pattern: "{teamID}-team.yaml"
    cache_enabled: true
    cache_ttl: 3600 seconds
    
  team_config_structure:
    team_info: "团队基本信息（ID、名称、描述、版本）"
    metadata: "团队元数据（作者、创建时间、标签、适用阶段）"
    team_lead: "团队领导配置（代理、变体、职责、权限）"
    phases: "阶段特定配置（描述、持续时间、任务模式、阶段领导）"
    standard_tasks: "标准任务模板"
    quality_gates: "质量门禁配置"
    
  configuration_sources:
    primary: ".goagents/teams/{teamID}-team.yaml"
    fallback: "workspace/skills/team-roles/{teamID}.yaml"
    default: "execution_framework默认团队配置"
```

### 输入格式
```yaml
input:
  project_type: string         # 项目类型
  complexity: string           # 项目复杂度
  mode: string                 # 执行模式（从ExecutionFramework动态获取）
  requirements: object         # 项目需求
  constraints: map[string]string # 项目约束
  use_dynamic_config: boolean  # 使用动态配置加载
  team_preferences: map[string]string # 团队偏好设置
```

### 输出格式
```yaml
output:
  team_composition: map[string]string  # 角色 -> 技能映射
  core_roles: string[]                 # 核心角色列表
  support_roles: string[]              # 支持角色列表
  team_score: float                   # 团队配置评分
  validation_result: object            # 验证结果
  optimization_suggestions: string[]   # 优化建议
```

## 📋 角色定义

### 核心角色（所有项目必需）
```yaml
core_roles:
  po: "产品负责人 - 项目决策和协调"
  analyst: "需求分析师 - 需求分析和业务理解"
  architect: "架构师 - 系统架构设计"
  developer: "开发者 - 代码实现和开发"
  qa: "测试工程师 - 质量保证和测试"
```

### 领域特定角色
```yaml
domain_specific_roles:
  ecommerce:
    - "ui_designer"      # UI设计师
    - "product_manager"  # 产品经理
    - "payment_specialist" # 支付专家
    
  fintech:
    - "security_specialist"   # 安全专家
    - "compliance_officer"   # 合规官
    - "risk_manager"         # 风险经理
    
  gaming:
    - "game_designer"    # 游戏设计师
    - "game_artist"      # 游戏美术
    - "sound_designer"   # 音效设计师
    - "level_designer"   # 关卡设计师
    
  ai:
    - "ml_engineer"      # 机器学习工程师
    - "data_scientist"   # 数据科学家
    - "research_engineer" # 研究工程师
    
  healthcare:
    - "domain_expert"    # 领域专家
    - "compliance_officer" # 合规官
    - "security_specialist" # 安全专家
```

### 支持角色（根据复杂度添加）
```yaml
support_roles:
  high_complexity:
    - "senior_architect"    # 高级架构师
    - "performance_engineer" # 性能工程师
    - "tech_lead"          # 技术负责人
    
  medium_complexity:
    - "integration_specialist" # 集成专家
    - "devops_engineer"       # DevOps工程师
    
  business_complexity:
    - "business_analyst"   # 业务分析师
    - "stakeholder_manager" # 利益相关者经理
    
  large_team:
    - "team_lead"         # 团队负责人
    - "coordinator"       # 协调员
```

## 🎯 使用示例

### 完整团队构建
```bash
@go "使用team-builder技能组建AI Specialist团队：
项目：电商平台开发
模式：standard
需求：
- 功能需求：15个
- 技术复杂度：medium
- 业务复杂度：medium
- 领域：ecommerce
约束：高性能要求，安全要求"
```

### 基于模式的团队构建
```bash
@go "使用team-builder技能为standard模式组建团队：
项目类型：web应用
复杂度：medium
领域：general
团队规模：适中"
```

### 领域特定团队
```bash
@go "使用team-builder技能组建游戏开发团队：
项目：休闲益智游戏
模式：free
领域：gaming
复杂度：high
特殊需求：美术、音效、关卡设计"
```

## 🔍 匹配算法

### 角色选择算法
```yaml
role_selection:
  1. "添加所有核心角色"
  2. "根据项目领域添加领域特定角色"
  3. "根据复杂度添加支持角色"
  4. "根据模式调整角色配置"
  5. "验证角色完整性"
```

### 技能分配算法
```yaml
skill_assignment:
  1. "获取可用技能列表"
  2. "为每个角色生成候选技能"
  3. "计算技能匹配度"
  4. "选择最佳匹配技能"
  5. "验证技能可用性"
```

### 团队评分算法
```yaml
team_scoring:
  core_roles_score: 75%    # 核心角色完整性
  domain_roles_score: 15%   # 领域角色适配度
  support_roles_score: 10%  # 支持角色合理性
  
  total_score: "core_score + domain_score + support_score"
  max_score: 100
```

## 📊 技能匹配策略

### 精确匹配
```yaml
exact_matching:
  strategy: "skill_name == role_preference"
  priority: "highest"
  confidence: 100%
```

### 包含匹配
```yaml
contains_matching:
  strategy: "skill_name contains role_keyword"
  priority: "high"
  confidence: 80%
```

### 关键词匹配
```yaml
keyword_matching:
  strategy: "skill_keywords intersect role_keywords"
  priority: "medium"
  confidence: 60%
```

### 默认匹配
```yaml
default_matching:
  strategy: "use_predefined_default_skill"
  priority: "low"
  confidence: 40%
```

## 🎯 使用示例

### 电商项目团队
```bash
@go "使用team-builder技能构建电商团队：
输入：
- 领域：ecommerce
- 模式：standard
- 复杂度：medium

期望输出：
{
  "po": "po-core",
  "analyst": "business-analyst", 
  "architect": "system-architect",
  "developer": "fullstack-developer",
  "qa": "quality-engineer",
  "ui_designer": "ui-designer",
  "product_manager": "product-manager"
}"
```

### 游戏项目团队
```bash
@go "使用team-builder技能构建游戏团队：
输入：
- 领域：gaming
- 模式：free
- 复杂度：high

期望输出：
{
  "po": "po-core",
  "analyst": "game-analyst",
  "architect": "game-architect", 
  "developer": "game-developer",
  "qa": "game-qa",
  "game_designer": "game-designer",
  "game_artist": "game-artist",
  "sound_designer": "sound-designer",
  "senior_architect": "senior-game-architect",
  "performance_engineer": "game-performance-engineer"
}"
```

### AI项目团队
```bash
@go "使用team-builder技能构建AI团队：
输入：
- 领域：ai
- 模式：free
- 复杂度：high

期望输出：
{
  "po": "po-core",
  "analyst": "ai-analyst",
  "architect": "ml-architect",
  "developer": "ml-developer", 
  "qa": "ml-qa",
  "ml_engineer": "ml-engineer",
  "data_scientist": "data-scientist",
  "research_engineer": "research-engineer",
  "senior_architect": "senior-ml-architect"
}"
```

## 📊 验证标准

### 团队完整性验证
```yaml
completeness_validation:
  core_roles_required: true
  domain_roles_recommended: true
  support_roles_optional: true
  
  validation_rules:
    - "所有核心角色必须存在"
    - "领域特定角色强烈推荐"
    - "支持角色根据需要添加"
```

### 团队合理性验证
```yaml
rationality_validation:
  team_size_limits:
    min: 5
    max: 15
    optimal: 7-10
    
  role_conflicts:
    check: true
    resolution: "automatic"
    
  skill_availability:
    check: true
    fallback: "default_skill"
```

### 团队优化验证
```yaml
optimization_validation:
  score_threshold: 70
  improvement_suggestions: true
  
  optimization_criteria:
    - "团队配置评分"
    - "角色覆盖度"
    - "技能匹配度"
    - "项目适配度"
```

## 🚀 性能要求

### 处理时间
- **角色选择**: < 10秒
- **技能分配**: < 15秒
- **团队验证**: < 5秒
- **总计**: < 30秒

### 准确性要求
- **角色选择准确率**: > 90%
- **技能匹配准确率**: > 85%
- **团队配置合理性**: > 80%

---

这个技能专门负责AI Specialist团队的智能组建，确保每个项目都有最适合的团队配置，是PO Core系统协调能力的重要组成部分。
