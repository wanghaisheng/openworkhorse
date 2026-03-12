---
name: mode-selector
description: "模式选择AI专家 - 专门分析项目特征并选择最优执行模式的专业AI技能"
---

# 模式选择AI专家

## 🎯 技能角色

### 核心职责
- **模式分析**: 分析Standard/Free/Hybrid三种模式的适用性
- **特征匹配**: 将项目特征与模式要求进行匹配
- **评分计算**: 计算各模式的适配度分数
- **推荐决策**: 基于评分结果推荐最优模式

### 专业能力
- **模式理解**: 深度理解三种执行模式的特点和适用场景
- **项目分析**: 分析项目的复杂度、创新性、时间压力等特征
- **决策算法**: 使用多维度评分算法进行模式选择
- **风险评估**: 评估不同模式下的项目风险

## 🔧 功能实现

### 输入格式
```yaml
input:
  requirements: object  # 来自requirement-analyzer的输出
  project_request: object  # 原始项目请求
  constraints: map[string]string
```

### 输出格式
```yaml
output:
  selected_mode: string  # standard/free/hybrid
  mode_scores: object    # 各模式的评分详情
  confidence: float      # 推荐置信度
  reasoning: string      # 选择理由
  alternatives: object   # 备选方案分析
```

## 📋 模式定义

### Standard模式
**适用场景**:
- 需求明确且稳定
- 技术方案成熟
- 项目复杂度低到中等
- 时间压力适中

**特点**:
- 模板化执行
- 标准化流程
- 高效率交付
- 低风险实施

### Free模式
**适用场景**:
- 需求不明确或需要探索
- 技术方案需要创新
- 项目复杂度高
- 时间压力相对宽松

**特点**:
- 灵活探索
- 迭代开发
- 创新导向
- 高风险高回报

### Hybrid模式
**适用场景**:
- 部分需求明确，部分需要探索
- 技术和业务复杂度中等
- 需要平衡效率和灵活性
- 时间压力适中

**特点**:
- 结构化灵活性
- 分阶段执行
- 风险平衡
- 渐进式交付

## 🎯 使用示例

### 完整模式选择
```bash
@go "使用mode-selector技能选择执行模式：
项目：电商平台开发
需求：功能需求明确，技术方案成熟
复杂度：中等
时间：2个月交付
约束：需要标准化流程"
```

### 基于需求分析结果
```bash
@go "使用mode-selector技能基于需求分析结果选择模式：
需求分析结果：
- 功能需求：15个
- 技术复杂度：medium
- 业务复杂度：low
- 风险：技术风险中等
项目类型：web应用"
```

## 🔍 评分算法

### 评分维度
```yaml
scoring_dimensions:
  project_type_fit:      # 项目类型适配度 (30分)
  complexity_match:       # 复杂度匹配度 (25分)
  requirement_clarity:    # 需求明确性 (20分)
  time_pressure:         # 时间压力 (15分)
  innovation_level:      # 创新程度 (10分)
```

### Standard模式评分
```yaml
standard_scoring:
  project_type_fit:
    ecommerce: +30
    enterprise: +30
    general: +25
    gaming: +10
    ai: +5
    
  complexity_match:
    low: +25
    medium: +20
    high: +5
    
  requirement_clarity:
    clear: +20
    partial: +10
    unclear: +0
    
  time_pressure:
    sufficient: +15
    moderate: +10
    tight: +5
    
  innovation_level:
    low: +10
    medium: +5
    high: +0
```

### Free模式评分
```yaml
free_scoring:
  project_type_fit:
    ai: +30
    gaming: +25
    research: +30
    innovation: +25
    ecommerce: +10
    
  complexity_match:
    high: +25
    medium: +15
    low: +5
    
  requirement_clarity:
    unclear: +20
    partial: +15
    clear: +5
    
  time_pressure:
    sufficient: +15
    moderate: +10
    tight: +0
    
  innovation_level:
    high: +10
    medium: +5
    low: +0
```

### Hybrid模式评分
```yaml
hybrid_scoring:
  project_type_fit:
    general: +25
    enterprise: +20
    ecommerce: +20
    gaming: +15
    
  complexity_match:
    medium: +25
    high: +15
    low: +10
    
  requirement_clarity:
    partial: +20
    clear: +15
    unclear: +5
    
  time_pressure:
    moderate: +15
    sufficient: +10
    tight: +5
    
  innovation_level:
    medium: +10
    high: +5
    low: +5
```

## 📊 决策逻辑

### 基础决策
```yaml
decision_logic:
  1. "计算各模式评分"
  2. "选择得分最高的模式"
  3. "检查分数差距"
  4. "如果差距<10分，选择hybrid模式"
  5. "计算推荐置信度"
  6. "生成选择理由"
```

### 置信度计算
```yaml
confidence_calculation:
  factors:
    - "最高分数与平均分数的差距"
    - "评分原因的数量和质量"
    - "项目特征的明确程度"
    
  formula: "(score_gap + reason_count * 5) / 110"
  range: "0.0 - 1.0"
```

## 🎯 使用示例

### 电商项目示例
```bash
@go "使用mode-selector技能分析：
项目：B2C电商平台
特征：
- 项目类型：ecommerce
- 复杂度：medium
- 需求明确性：clear
- 时间压力：moderate
- 创新程度：low

期望输出：Standard模式（高置信度）"
```

### AI项目示例
```bash
@go "使用mode-selector技能分析：
项目：AI智能助手
特征：
- 项目类型：ai
- 复杂度：high
- 需求明确性：unclear
- 时间压力：sufficient
- 创新程度：high

期望输出：Free模式（高置信度）"
```

### 混合项目示例
```bash
@go "使用mode-selector技能分析：
项目：企业级系统
特征：
- 项目类型：enterprise
- 复杂度：medium
- 需求明确性：partial
- 时间压力：moderate
- 创新程度：medium

期望输出：Hybrid模式（中等置信度）"
```

## 📊 质量标准

### 输出质量要求
- **决策合理性**: 选择模式有充分的理由支撑
- **评分准确性**: 评分算法反映真实适配度
- **置信度可靠**: 置信度计算准确反映决策确定性
- **解释清晰**: 提供清晰的选择理由和备选方案

### 验证标准
```yaml
validation:
  mode_selection:
    justification_required: true
    confidence_threshold: "> 0.6"
    
  scoring_accuracy:
    cross_validation: true
    historical_accuracy: "> 80%"
    
  reasoning_quality:
    clarity_score: "> 85%"
    completeness_score: "> 90%"
```

## 🚀 性能要求

### 处理时间
- **模式分析**: < 15秒
- **评分计算**: < 10秒
- **决策生成**: < 5秒
- **总计**: < 30秒

### 准确性要求
- **模式选择准确率**: > 85%
- **置信度预测准确率**: > 80%
- **评分一致性**: > 90%

## 🔄 模式执行器映射

### 执行器分配
```yaml
mode_executor_mapping:
  standard:
    executor_skill: "standard-mode-executor"
    execution_style: "template_based"
    quality_control: "strict"
    template_library: "enterprise_web_application, ecommerce_platform, financial_system"
    
  free:
    executor_skill: "free-mode-executor"  # 待实现
    execution_style: "dynamic"
    quality_control: "flexible"
    exploration_methods: "prototype_driven, iterative_development"
    
  hybrid:
    executor_skill: "hybrid-mode-executor"  # 待实现
    execution_style: "mixed"
    quality_control: "balanced"
    combination_strategy: "structured_flexibility"
```

### 模式选择到执行器的转换
```yaml
selection_to_execution:
  process_flow:
    1. "mode-selector分析项目特征"
    2. "计算各模式适配度分数"
    3. "选择最优执行模式"
    4. "映射到对应的执行器技能"
    5. "传递执行参数给执行器"
    
  parameter_mapping:
    project_request: "直接传递"
    team_composition: "直接传递"
    quality_requirements: "转换为质量门禁"
    schedule_constraints: "转换为执行计划"
```

---

这个技能专门负责执行模式选择，基于项目特征智能推荐最适合的开发模式，并映射到对应的执行器技能，是PO Core系统决策能力的重要组成部分。
