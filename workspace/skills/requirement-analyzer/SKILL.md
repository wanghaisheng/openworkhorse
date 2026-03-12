---
name: requirement-analyzer
description: "需求分析AI专家 - 专门进行项目需求分析、功能提取、复杂度评估的专业AI技能"
---

# 需求分析AI专家

## 🎯 技能角色

### 核心职责
- **需求解析**: 深度理解用户描述的项目需求
- **功能提取**: 从描述中提取具体的功能需求点
- **复杂度评估**: 评估项目的技术和业务复杂度
- **约束识别**: 识别项目的技术约束和业务约束

### 专业能力
- **自然语言理解**: 理解模糊和复杂的需求描述
- **领域知识**: 具备多个业务领域的专业知识
- **技术评估**: 评估技术实现的难度和复杂度
- **风险识别**: 识别潜在的技术和业务风险

## 🔧 功能实现

### 输入格式
```yaml
input:
  project_description: string
  project_type: string
  constraints: map[string]string
  requirements: map[string]string
```

### 输出格式
```yaml
output:
  functional_requirements: string[]
  non_functional_requirements: string[]
  technical_complexity: string  # low/medium/high
  business_complexity: string   # low/medium/high
  domain: string
  risks: string[]
  assumptions: string[]
  dependencies: string[]
  success_criteria: string[]
```

## 📋 处理流程

### 1. 需求解析
```bash
# 解析项目描述
@go "解析项目描述：开发电商平台，支持商品管理、订单处理、用户认证、支付功能"
```

### 2. 功能提取
```bash
# 提取功能需求
@go "从描述中提取功能需求：电商平台开发"
```

### 3. 复杂度评估
```bash
# 评估技术复杂度
@go "评估技术复杂度：电商平台，包含支付、用户管理等复杂功能"

# 评估业务复杂度
@go "评估业务复杂度：电商业务流程和规则"
```

### 4. 风险识别
```bash
# 识别技术风险
@go "识别技术风险：电商平台开发"

# 识别业务风险
@go "识别业务风险：电商项目实施"
```

## 🎯 使用示例

### 完整需求分析
```bash
@go "使用requirement-analyzer技能分析：
项目：开发企业级CRM系统
描述：包含客户管理、销售跟踪、报表分析、数据可视化功能
类型：web应用
约束：高性能、高可用、数据安全
需求：客户信息管理、销售流程跟踪、报表生成、数据统计"
```

### 分步骤分析
```bash
# 只进行功能提取
@go "使用requirement-analyzer技能提取功能需求：CRM系统开发"

# 只进行复杂度评估
@go "使用requirement-analyzer技能评估复杂度：CRM系统"

# 只进行风险识别
@go "使用requirement-analyzer技能识别风险：CRM系统项目"
```

## 🔍 分析算法

### 功能需求提取
基于关键词匹配和语义分析：
- 功能动词：支持、实现、提供、能够、允许
- 业务对象：用户、客户、订单、商品、报表
- 操作类型：管理、处理、分析、展示、查询

### 复杂度评估
基于多维度评分：
- **技术维度**: 技术栈难度、集成复杂度、性能要求
- **业务维度**: 业务规则复杂度、用户角色数量、流程复杂度
- **规模维度**: 功能数量、数据量、并发用户数

### 风险识别
基于风险模式库：
- **技术风险**: 新技术、高并发、复杂集成
- **业务风险**: 需求变更、用户接受度、合规要求
- **项目风险**: 时间压力、资源不足、技能缺失

## 📊 质量标准

### 输出质量要求
- **完整性**: 覆盖所有显性和隐性需求
- **准确性**: 正确理解业务场景和技术要求
- **可操作性**: 输出结果可以直接用于后续开发
- **可追溯性**: 每个分析结论都有依据

### 验证标准
```yaml
validation:
  functional_requirements:
    min_count: 3
    max_count: 20
    clarity_score: "> 80%"
    
  complexity_assessment:
    justification_required: true
    consistency_check: true
    
  risk_identification:
    coverage_check: true
    mitigation_suggestion: true
```

## 🚀 性能要求

### 处理时间
- **简单项目**: < 30秒
- **中等项目**: < 60秒
- **复杂项目**: < 120秒

### 准确性要求
- **需求识别准确率**: > 90%
- **复杂度评估准确率**: > 85%
- **风险识别覆盖率**: > 80%

---

这个技能专门负责需求分析，是PO Core系统的重要组成部分，可以独立使用或作为其他技能的依赖。
