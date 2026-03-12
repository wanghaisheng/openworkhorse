---
name: po-core-v2
description: "PO Core v2 - 完全基于技能系统的产品负责人AI专家，协调多个专业技能实现完整的项目处理流程"
---

# PO Core v2 - 基于技能系统的产品负责人AI专家

## 🎯 技能设计原则

### 零侵入性设计
- **完全基于现有技能系统**: 不修改`pkg`包结构
- **技能驱动实现**: 所有功能通过技能定义和执行框架实现
- **可插拔架构**: 可以独立部署和更新

### AI Specialist专业化
- **产品决策引擎**: 智能需求分析和项目规划
- **模式选择器**: 基于项目特征选择最优执行模式
- **团队协调器**: 智能组建和管理AI Specialist团队
- **质量保证者**: 确保输出符合HARNESS.md和OpenSpec标准

## 🏗️ 技能协作架构

### 核心技能协调
```yaml
skill_orchestration:
  primary_skill: "po-core-v2"
  dependency_skills:
    - "requirement-analyzer"
    - "mode-selector"
    - "team-builder"
    - "phase-manager-v2"
    - "harness-integrator"
    - "ralph-wiggum-loop"  # 新增Ralph Wiggum Loop技能
  
  execution_flow:
    1. "接收项目请求"
    2. "调用harness-integrator进行预检查"
    3. "调用requirement-analyzer进行需求分析"
    4. "调用mode-selector选择执行模式"
    5. "调用team-builder组建团队"
    6. "调用phase-manager-v2规划阶段"
    7. "调用harness-integrator进行后检查"
    8. "调用ralph-wiggum-loop执行质量循环"
    9. "生成完整项目响应"
```

### 技能间数据流
```yaml
data_flow:
  input: "project_request"
  
  step_1:
    skill: "requirement-analyzer"
    input: "project_request"
    output: "requirements_analysis"
    
  step_2:
    skill: "mode-selector"
    input: "requirements_analysis + project_request"
    output: "mode_selection"
    
  step_3:
    skill: "team-builder"
    input: "requirements_analysis + mode_selection"
    output: "team_composition"
    
  step_4:
    skill: "phase-manager-v2"
    input: "mode_selection + requirements_analysis + team_composition"
    output: "phase_plan"
    
  final_output:
    combine: "requirements_analysis + mode_selection + team_composition + phase_plan"
    format: "project_response"
```

## 📋 完整功能实现

### 1. 项目请求处理
通过技能协调处理完整的项目请求：

```bash
# 完整项目处理
@go "使用po-core-v2技能处理完整项目：
项目：开发企业级CRM系统
描述：包含客户管理、销售跟踪、报表分析、数据可视化功能
类型：web应用
复杂度：medium
时间：3个月交付
约束：高性能、高可用、数据安全
需求：客户信息管理、销售流程跟踪、报表生成、数据统计"
```

### 2. 自动技能协调
PO Core v2自动协调子技能的处理流程：

```yaml
automatic_coordination:
  trigger: "@go 使用po-core-v2技能处理：[项目描述]"
  
  auto_steps:
    1. "解析项目请求参数"
    2. "调用requirement-analyzer技能"
    3. "等待需求分析完成"
    4. "调用mode-selector技能"
    5. "等待模式选择完成"
    6. "调用team-builder技能"
    7. "等待团队组建完成"
    8. "调用phase-manager-v2技能"
    9. "等待阶段规划完成"
    10. "整合所有结果生成项目响应"
```

### 3. 分步骤处理
支持分步骤调用子技能：

```bash
# 只进行需求分析
@go "使用po-core-v2技能执行步骤1：需求分析 - 企业级CRM系统"

# 只进行模式选择
@go "使用po-core-v2技能执行步骤2：模式选择 - 基于CRM系统需求分析结果"

# 只进行团队组建
@go "使用po-core-v2技能执行步骤3：团队组建 - 基于CRM系统模式和需求"

# 只进行阶段规划
@go "使用po-core-v2技能执行步骤4：阶段规划 - 基于CRM系统完整信息"
```

## 🔧 技能实现细节

### 技能协调器实现
```yaml
skill_coordinator:
  implementation: "skill_execution_framework"
  
  coordination_methods:
    sequential: "按顺序执行技能"
    parallel: "并行执行独立技能"
    conditional: "基于条件选择技能"
    
  error_handling:
    skill_failure: "使用默认值或跳过"
    timeout: "设置合理超时时间"
    retry: "失败重试机制"
```

### 数据结构定义
```yaml
data_structures:
  project_request:
    name: string
    description: string
    type: string
    complexity: string
    deadline: datetime
    requirements: map[string]string
    constraints: map[string]string
    
  requirements_analysis:
    functional_requirements: string[]
    non_functional_requirements: string[]
    technical_complexity: string
    business_complexity: string
    domain: string
    risks: string[]
    assumptions: string[]
    
  mode_selection:
    selected_mode: string
    mode_scores: object
    confidence: float
    reasoning: string
    
  team_composition:
    core_roles: map[string]string
    support_roles: map[string]string
    team_score: float
    validation_result: object
    
  phase_plan:
    phases: object[]
    total_milestones: int
    estimated_duration: duration
    dependencies: map[string][]string
    quality_gates: object[]
    
  project_response:
    project_id: string
    status: string
    selected_mode: string
    team_composition: map[string]string
    execution_plan: object
    quality_gates: string[]
    created_at: datetime
```

## 🎯 使用示例

### 完整项目处理示例
```bash
@go "使用po-core-v2技能处理完整项目：
项目：B2C电商平台开发
描述：支持商品展示、购物车、订单管理、支付系统、用户管理功能
类型：web应用
复杂度：medium
时间：2个月交付
约束：高性能、高可用、支付安全
需求：
- 商品管理：商品信息CRUD、分类管理、搜索功能
- 订单管理：购物车、下单、支付、订单查询
- 用户管理：注册、登录、个人信息管理
- 支付系统：多种支付方式、交易安全"

# 自动执行流程：
# 1. 调用requirement-analyzer分析需求
# 2. 调用mode-selector选择standard模式
# 3. 调用team-builder组建电商专业团队
# 4. 调用phase-manager-v2规划4个阶段12个里程碑
# 5. 生成完整的项目响应
```

### 游戏项目示例
```bash
@go "使用po-core-v2技能处理游戏项目：
项目：休闲益智游戏开发
描述：包含关卡设计、角色系统、积分排行榜、音效系统
类型：game
复杂度：high
时间：4个月交付
约束：流畅运行、适配多平台
需求：
- 游戏玩法：核心游戏机制、关卡设计
- 角色系统：角色设计、升级系统
- 积分系统：积分计算、排行榜
- 音效系统：背景音乐、音效"

# 自动执行流程：
# 1. 调用requirement-analyzer分析游戏需求
# 2. 调用mode-selector选择free模式（创新性高）
# 3. 调用team-builder组建游戏专业团队
# 4. 调用phase-manager-v2规划探索+迭代开发
# 5. 生成游戏项目响应
```

### AI项目示例
```bash
@go "使用po-core-v2技能处理AI项目：
项目：智能客服系统
描述：基于NLP的智能问答、意图识别、知识图谱
类型：ai
复杂度：high
时间：6个月交付
约束：高准确率、实时响应
需求：
- NLP处理：自然语言理解、意图识别
- 知识管理：知识图谱构建、更新
- 问答系统：智能回答、多轮对话
- 性能要求：实时响应、高并发"

# 自动执行流程：
# 1. 调用requirement-analyzer分析AI需求
# 2. 调用mode-selector选择free模式（探索性强）
# 3. 调用team-builder组建AI专业团队
# 4. 调用phase-manager-v2规划研究+开发阶段
# 5. 生成AI项目响应
```

## 🔍 质量保证

### HARNESS.md集成
每个技能执行前自动检查HARNESS.md：
```bash
First, read and strictly follow HARNESS.md in repository root.
Do NOT deviate from its rules.
```

### 技能协调质量保证
```yaml
coordination_quality:
  skill_execution_validation: true
  data_flow_integrity: true
  error_handling_robustness: true
  performance_monitoring: true
```

### Ralph Wiggum Loop
通过技能执行框架自动执行质量检查循环：
```yaml
ralph_wiggum_loop:
  triggers:
    - "po_core_execution_start"
    - "subskill_completion"
    - "quality_gate_failure"
    - "milestone_completion"
  
  checks:
    - "harness_compliance"
    - "skill_coordination_quality"
    - "data_integrity"
    - "output_completeness"
    - "user_satisfaction"
```

## 📊 技能配置

### 技能依赖配置
```yaml
skill_dependencies:
  required_skills:
    - name: "requirement-analyzer"
      version: ">=1.0.0"
      critical: true
    - name: "mode-selector"
      version: ">=1.0.0"
      critical: true
    - name: "team-builder"
      version: ">=1.0.0"
      critical: true
    - name: "phase-manager-v2"
      version: ">=1.0.0"
      critical: true
    - name: "harness-integrator"
      version: ">=1.0.0"
      critical: true
    - name: "ralph-wiggum-loop"
      version: ">=1.0.0"
      critical: true  # Ralph Wiggum Loop是质量保证核心
      
  optional_skills:
    - name: "quality-analyzer"
      version: ">=1.0.0"
      critical: false
```

### 执行参数配置
```yaml
execution_parameters:
  timeout_settings:
    harness_pre_check: 30      # HARNESS.md预检查
    requirement_analysis: 60
    mode_selection: 30
    team_building: 45
    phase_planning: 60
    harness_post_check: 30   # HARNESS.md后检查
    ralph_wiggum_loop: 60     # Ralph Wiggum Loop质量循环
    total_execution: 360      # 增加总执行时间
    
  retry_policy:
    max_retries: 3
    backoff_strategy: "exponential"
    
  quality_thresholds:
    min_confidence: 0.7
    min_team_score: 70
    min_phase_coverage: 0.8
    min_harness_compliance: 0.9
    min_quality_score: 0.85    # Ralph Wiggum Loop质量分数要求
```

## 🚀 部署和使用

### 技能部署
```bash
# 部署PO Core v2及完整依赖技能（包含质量保证体系）
@go "部署技能组合：po-core-v2 + requirement-analyzer + mode-selector + team-builder + phase-manager-v2 + harness-integrator + ralph-wiggum-loop"

# 验证技能组合
@go "验证技能组合：po-core-v2完整功能测试"
```

### 技能更新
```bash
# 更新PO Core v2
@go "更新技能：po-core-v2"

# 更新依赖技能
@go "更新技能组合：po-core-v2依赖技能"
```

## 🎉 优势总结

### 1. 零侵入性
- ✅ 不修改任何`pkg`包结构
- ✅ 完全基于现有技能系统
- ✅ 可以独立部署和更新

### 2. 智能协调
- ✅ 自动协调多个专业技能
- ✅ 智能处理技能间数据流
- ✅ 自动错误处理和恢复

### 3. 专业能力
- ✅ 每个子技能代表一个专业领域
- ✅ 可以独立使用或组合使用
- ✅ 支持灵活的技能组合

### 4. 质量保证
- ✅ 集成HARNESS.md和Ralph Wiggum Loop
- ✅ 自动质量检查和验证
- ✅ 持续改进和学习

---

这个设计完全基于技能系统，实现了零侵入性的PO Core功能，通过智能协调多个专业AI技能，提供了完整的项目处理能力，真正体现了AI Specialist的专业化和协作特性。
