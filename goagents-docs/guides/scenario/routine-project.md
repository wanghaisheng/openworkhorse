# 📋 常规项目 Go Agents 指南

## 🎯 概述

常规项目具有需求较清晰、节奏稳定、可预测性高等特点。Go Agents v2.0 的 Standard 模式为常规项目提供高效的标准化执行流程，确保项目按时、按质、按预算交付。

## 🏗️ 推荐配置

### Phase配置
```bash
# 创建常规项目Phase配置
picoclaw goagents phase create routine-discovery
picoclaw goagents phase create routine-architecture
picoclaw goagents phase create routine-development
picoclaw goagents phase create routine-validation

# 配置特点：
# - Discovery: 标准需求分析 + 用户调研
# - Architecture: 标准架构设计 + 技术选型
# - Development: 标准功能开发 + 测试
# - Validation: 标准质量验证 + 部署
```

### Team配置
```bash
# 创建常规项目团队
picoclaw goagents team create routine-discovery-team
picoclaw goagents team create routine-architecture-team
picoclaw goagents team create routine-development-team
picoclaw goagents team create routine-validation-team

# 推荐团队构成：
# - Discovery: business_analyst + user_researcher
# - Architecture: system_architect + technical_architect
# - Development: frontend_developer + backend_developer + database_developer
# - Validation: qa_engineer + performance_engineer
```

### 任务模式推荐
- **主要推荐**: Standard 模式（标准化、质量稳定）
- **特殊情况**: Hybrid 模式（有少量创新需求时）

## 📋 项目特征

### 常规项目特点
- **需求明确**: 项目需求清晰明确
- **技术成熟**: 使用成熟的技术栈
- **流程稳定**: 开发流程稳定可靠
- **质量可控**: 质量标准明确可控

### 典型场景
- **功能迭代**: 现有产品功能迭代
- **系统升级**: 系统版本升级
- **维护开发**: 系统维护和开发
- **标准化开发**: 标准化产品开发

## 🚀 执行流程

### 1. 项目启动
```bash
# 初始化常规项目
picoclaw goagents phase use routine-discovery
picoclaw goagents team use routine-discovery-team
picoclaw goagents config set mode=standard

# 启动项目
@go "启动常规项目：[项目名称]"
```

### 2. 标准化执行
```bash
# 按标准流程执行
@go --standard "执行项目标准化开发流程"

# 自动任务分解：
# - 需求分析 → 架构设计 → 功能开发 → 质量验证
# - 每个步骤都有明确的输入输出和质量标准
# - 严格按照模板和标准执行
```

### 3. 质量控制
```bash
# 质量检查
picoclaw goagents quality check

# 质量门禁验证
picoclaw goagents quality gates

# 质量报告
picoclaw goagents quality report
```

## 🎯 阶段指南

### Discovery（发现阶段）

#### 标准化需求分析
```bash
# 需求分析任务
picoclaw goagents milestone create requirements-analysis-milestone
picoclaw goagents task create business-analysis

# 用户研究任务
picoclaw goagents task create user-research

# 需求规格化
picoclaw goagents task create requirement-specification
```

#### 标准化输出
- 需求规格文档
- 用户研究报告
- 项目范围定义
- 验收标准清单

### Architecture（架构阶段）

#### 标准化架构设计
```bash
# 系统架构设计
picoclaw goagents milestone create architecture-design-milestone
picoclaw goagents task create system-architecture

# 技术设计
picoclaw goagents task create technical-design

# 数据库设计
picoclaw goagents task create database-design
```

#### 标准化输出
- 系统架构文档
- 技术设计文档
- 数据库设计文档
- 开发计划

### Development（开发阶段）

#### 标准化功能开发
```bash
# 前端开发
picoclaw goagents milestone create development-milestone
picoclaw goagents task create frontend-development

# 后端开发
picoclaw goagents task create backend-development

# 数据库开发
picoclaw goagents task create database-development

# 测试开发
picoclaw goagents task create unit-testing
```

#### 标准化输出
- 功能代码
- 测试代码
- 技术文档
- 用户手册

### Validation（验证阶段）

#### 标准化质量验证
```bash
# 功能测试
picoclaw goagents milestone create validation-milestone
picoclaw goagents task create functional-testing

# 性能测试
picoclaw goagents task create performance-testing

# 用户验收测试
picoclaw goagents task create user-acceptance-testing
```

#### 标准化输出
- 测试报告
- 性能报告
- 用户验收报告
- 交付文档

## 📊 质量保证

### 标准化质量标准
```yaml
quality_standards:
  code_quality:
    complexity: "≤ 10"
    coverage: "≥ 80%"
    duplication: "≤ 5%"
  
  documentation_quality:
    completeness: "≥ 90%"
    accuracy: "≥ 95%"
    maintainability: "≥ 85%"
  
  testing_quality:
    unit_test_coverage: "≥ 80%"
    integration_test_coverage: "≥ 70%"
    acceptance_test_coverage: "≥ 85%"
```

### 标准化质量门禁
- 每个阶段都有严格的质量门禁
- 不通过则不能进入下一阶段
- 支持质量改进和重新检查

## 🎯 使用示例

### 示例1：用户管理系统升级
```bash
# 启动项目
@go "升级用户管理系统"

# PO 分析结果：
# 📋 **PO任务分析结果**
# 
# **需求**: 升级用户管理系统
# **推荐模式**: Standard（标准化升级项目）
# **执行计划**:
# - Discovery: 现有系统分析 + 升级需求定义
# - Architecture: 升级架构设计 + 兼容性设计
# - Development: 功能升级 + 测试
# - Validation: 质量验证 + 部署
```

### 示例2：报表系统开发
```bash
# 启动项目
@go "开发报表系统"

# PO 分析结果：
# 📋 **PO任务分析结果**
# 
# **需求**: 开发报表系统
# **推荐模式**: Standard（标准化报表系统）
# **执行计划**:
# - Discovery: 报表需求分析 + 数据源分析
# - Architecture: 报表架构设计 + 数据库设计
# - Development: 报表功能开发 + 测试
# - Validation: 报表验证 + 部署
```

## 🚀 最佳实践

### 1. 严格遵循标准
- 遵循项目开发标准
- 使用标准模板和工具
- 保持流程一致性

### 2. 注重质量
- 严格按照质量标准执行
- 进行充分的质量检查
- 持续改进质量

### 3. 有效沟通
- 建立标准化沟通机制
- 定期同步项目进度
- 及时解决问题

### 4. 风险管理
- 识别项目风险
- 制定应对策略
- 建立应急预案

## 📈 性能指标

### 效率指标
- **需求分析效率**: ≥ 90%
- **开发效率**: ≥ 85%
- **测试效率**: ≥ 90%
- **交付效率**: ≥ 95%

### 质量指标
- **代码质量**: ≥ 85%
- **测试覆盖率**: ≥ 80%
- **文档质量**: ≥ 90%
- **用户满意度**: ≥ 85%

### 成本指标
- **预算控制**: ≤ 100%
- **资源利用率**: ≥ 80%
- **返工率**: ≤ 5%
- **缺陷率**: ≤ 3%

---

**常规项目指南为Go Agents用户提供了高效的项目开发流程，通过标准化执行确保项目的成功交付！** 🚀
