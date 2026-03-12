---
name: phase-manager-v2
description: "阶段管理AI专家 - 专门管理项目阶段、里程碑和质量门禁的专业技能"
---

# 阶段管理AI专家

## 🎯 技能角色

### 核心职责
- **阶段定义**: 根据执行模式定义项目阶段
- **里程碑规划**: 规划每个阶段的具体里程碑
- **质量门禁设置**: 设置每个阶段的质量检查点
- **依赖管理**: 管理阶段和里程碑之间的依赖关系

### 专业能力
- **项目管理**: 深度理解项目管理最佳实践
- **里程碑规划**: 掌握300-800行代码的里程碑粒度控制
- **质量保证**: 理解质量门禁的设置和验证
- **依赖分析**: 分析和管理复杂的依赖关系

## 🔧 功能实现

### 输入格式
```yaml
input:
  selected_mode: string           # 执行模式
  requirements: object           # 项目需求
  team_composition: map[string]string  # 团队组成
  constraints: map[string]string # 项目约束
```

### 输出格式
```yaml
output:
  phases: object[]                # 阶段列表
  total_milestones: int           # 总里程碑数
  estimated_duration: duration    # 预估总时长
  dependencies: map[string][]string  # 依赖关系
  quality_gates: object[]         # 质量门禁
  risk_assessment: object         # 风险评估
```

## 📋 模式阶段定义

### Standard模式阶段
```yaml
standard_phases:
  - name: "research_discovery"
    duration: "1周"
    required_roles: ["po", "analyst"]
    quality_gates: ["market_analysis_complete", "concept_clear", "feasibility_confirmed"]
    deliverables: ["市场分析报告", "概念设计文档", "可行性评估报告"]
    
  - name: "requirements_planning"
    duration: "1.5周"
    required_roles: ["po", "analyst", "architect"]
    quality_gates: ["requirements_complete", "architecture_approved", "plan_realistic"]
    deliverables: ["详细需求文档", "系统架构设计", "项目计划"]
    
  - name: "development_implementation"
    duration: "3周"
    required_roles: ["architect", "developer", "qa"]
    quality_gates: ["code_complete", "tests_pass", "performance_ok"]
    deliverables: ["系统代码", "测试套件", "部署文档"]
    
  - name: "validation_testing"
    duration: "1周"
    required_roles: ["qa", "po", "analyst"]
    quality_gates: ["user_acceptance", "quality_gates_pass", "documentation_complete"]
    deliverables: ["测试报告", "用户手册", "发布准备"]
```

### Free模式阶段
```yaml
free_phases:
  - name: "exploration_research"
    duration: "2周"
    required_roles: ["po", "analyst", "innovation_lead"]
    quality_gates: ["prototype_working", "concept_validated", "approach_clear"]
    deliverables: ["研究报告", "原型系统", "技术验证"]
    
  - name: "iterative_development"
    duration: "3周"
    required_roles: ["developer", "architect", "qa"]
    quality_gates: ["iteration_complete", "feedback_incorporated", "quality_maintained"]
    deliverables: ["迭代版本", "反馈报告", "优化方案"]
```

### Hybrid模式阶段
```yaml
hybrid_phases:
  - name: "planned_exploration"
    duration: "1.5周"
    required_roles: ["po", "analyst", "architect"]
    quality_gates: ["requirements_confirmed", "approach_selected", "plan_approved"]
    deliverables: ["需求文档", "技术方案", "实施计划"]
    
  - name: "structured_flexibility"
    duration: "3周"
    required_roles: ["developer", "architect", "qa", "integration_specialist"]
    quality_gates: ["structured_complete", "flexibility_maintained", "integration_successful"]
    deliverables: ["核心系统", "扩展接口", "集成文档"]
```

## 🎯 里程碑定义

### 里程碑粒度控制
```yaml
milestone_granularity:
  code_lines:
    min: 300
    max: 800
    optimal: 500
    
  duration:
    min: "1天"
    max: "1周"
    optimal: "2-3天"
    
  scope:
    single_responsibility: true
    atomic_operation: true
```

### 里程碑结构
```yaml
milestone_structure:
  id: string                     # 唯一标识符
  name: string                   # 里程碑名称
  description: string             # 详细描述
  estimated_lines: int           # 预估代码行数
  estimated_duration: duration    # 预估时长
  acceptance_criteria: string[]   # 验收标准
  verification_commands: string[] # 验证命令
  dependencies: string[]         # 依赖里程碑
  quality_gates: string[]         # 质量门禁
```

## 🎯 使用示例

### 完整阶段规划
```bash
@go "使用phase-manager-v2技能规划项目阶段：
模式：standard
需求：电商平台开发
团队：包含po、analyst、architect、developer、qa、ui_designer
约束：2个月交付，高性能要求"
```

### 基于模式的阶段定义
```bash
@go "使用phase-manager-v2技能为standard模式定义阶段：
项目类型：web应用
复杂度：medium
团队规模：7人
时间约束：8周"
```

### 里程碑规划
```bash
@go "使用phase-manager-v2技能规划里程碑：
阶段：development_implementation
总代码量：3000行
时间：3周
团队：architect、developer、qa"
```

## 🔍 质量门禁设置

### 通用质量门禁
```yaml
universal_quality_gates:
  - "code_quality_check"      # 代码质量检查
  - "test_coverage_check"     # 测试覆盖率检查
  - "documentation_check"     # 文档完整性检查
  - "performance_check"       # 性能检查
```

### 模式特定质量门禁
```yaml
mode_specific_gates:
  standard:
    - "template_compliance_check"  # 模板合规检查
    - "standard_process_check"     # 标准流程检查
    
  free:
    - "innovation_validation_check" # 创新验证检查
    - "flexibility_check"          # 灵活性检查
    
  hybrid:
    - "template_compliance_check"  # 模板合规检查
    - "innovation_validation_check" # 创新验证检查
```

### 复杂度特定质量门禁
```yaml
complexity_specific_gates:
  high_complexity:
    - "architecture_review_check"  # 架构审查检查
    - "security_audit_check"       # 安全审计检查
    - "scalability_check"          # 可扩展性检查
    
  medium_complexity:
    - "design_review_check"        # 设计审查检查
    - "integration_check"          # 集成检查
    
  low_complexity:
    - "basic_functionality_check"  # 基本功能检查
```

## 📊 依赖关系管理

### 阶段依赖
```yaml
phase_dependencies:
  sequential: true
  parallel_allowed: false
  
  dependency_rules:
    - "research_discovery -> requirements_planning"
    - "requirements_planning -> development_implementation"
    - "development_implementation -> validation_testing"
```

### 里程碑依赖
```yaml
milestone_dependencies:
  within_phase: true
  cross_phase: false
  
  dependency_types:
    - "hard_dependency"    # 强依赖
    - "soft_dependency"    # 弱依赖
    - "resource_dependency" # 资源依赖
```

## 🎯 使用示例

### Standard模式示例
```bash
@go "使用phase-manager-v2技能规划standard模式项目：
输入：
- 模式：standard
- 项目：电商平台
- 复杂度：medium
- 时间：8周

期望输出：
{
  "phases": [
    {
      "name": "research_discovery",
      "duration": "1周",
      "required_roles": ["po", "analyst"],
      "milestones": 3
    },
    {
      "name": "requirements_planning", 
      "duration": "1.5周",
      "required_roles": ["po", "analyst", "architect"],
      "milestones": 3
    },
    {
      "name": "development_implementation",
      "duration": "3周", 
      "required_roles": ["architect", "developer", "qa"],
      "milestones": 3
    },
    {
      "name": "validation_testing",
      "duration": "1周",
      "required_roles": ["qa", "po", "analyst"], 
      "milestones": 3
    }
  ],
  "total_milestones": 12,
  "estimated_duration": "6.5周"
}"
```

### Free模式示例
```bash
@go "使用phase-manager-v2技能规划free模式项目：
输入：
- 模式：free
- 项目：AI创新应用
- 复杂度：high
- 时间：12周

期望输出：
{
  "phases": [
    {
      "name": "exploration_research",
      "duration": "2周",
      "required_roles": ["po", "analyst", "innovation_lead"],
      "milestones": 3
    },
    {
      "name": "iterative_development",
      "duration": "3周",
      "required_roles": ["developer", "architect", "qa"], 
      "milestones": 3
    }
  ],
  "total_milestones": 6,
  "estimated_duration": "5周"
}"
```

## 📊 验证标准

### 阶段配置验证
```yaml
phase_validation:
  completeness_check: true
  dependency_check: true
  resource_check: true
  
  validation_rules:
    - "所有阶段必须有明确的目标"
    - "阶段依赖必须合理"
    - "资源分配必须平衡"
```

### 里程碑验证
```yaml
milestone_validation:
  granularity_check: true
  dependency_check: true
  feasibility_check: true
  
  validation_rules:
    - "里程碑粒度必须在300-800行"
    - "依赖关系必须无循环"
    - "时间估算必须合理"
```

### 质量门禁验证
```yaml
quality_gate_validation:
  coverage_check: true
  appropriateness_check: true
  measurability_check: true
  
  validation_rules:
    - "质量门禁必须可测量"
    - "必须覆盖关键质量属性"
    - "必须适合项目特征"
```

## 🚀 性能要求

### 处理时间
- **阶段定义**: < 15秒
- **里程碑规划**: < 20秒
- **依赖分析**: < 10秒
- **质量门禁设置**: < 5秒
- **总计**: < 50秒

### 准确性要求
- **阶段规划准确率**: > 85%
- **里程碑粒度控制**: > 90%
- **依赖关系正确性**: > 95%
- **质量门禁适配度**: > 80%

---

这个技能专门负责项目阶段管理和里程碑规划，确保项目按照最优的节奏推进，是PO Core系统项目管理能力的重要组成部分。
