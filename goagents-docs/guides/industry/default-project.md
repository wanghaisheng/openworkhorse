# 🏗️ 通用项目 Go Agents 指南

## 🎯 概述

通用项目指南适合还没有明确行业模板、想先用默认配置启动的项目。Go Agents v2.0 采用配置驱动架构，为通用项目提供标准化的开发流程，确保项目能够快速启动并保持高质量。

## 🏗️ 推荐配置

### Phase配置
```bash
# 创建通用项目Phase配置
picoclaw goagents phase create default-discovery
picoclaw goagents phase create default-architecture
picoclaw goagents phase create default-development
picoclaw goagents phase create default-validation

# 配置特点：
# - Discovery: 需求分析 + 用户研究 + 市场分析
# - Architecture: 系统架构 + 技术设计 + 数据设计
# - Development: 功能开发 + 测试 + 集成
# - Validation: 质量验证 + 用户验收 + 部署
```

### Team配置
```bash
# 创建通用项目团队
picoclaw goagents team create default-discovery-team
picoclaw goagents team create default-architecture-team
picoclaw goagents team create default-development-team
picoclaw goagents team create default-validation-team

# 推荐团队构成：
# - Discovery: business_analyst + user_researcher + market_analyst
# - Architecture: system_architect + technical_architect + database_architect
# - Development: frontend_developer + backend_developer + database_developer
# - Validation: qa_engineer + performance_engineer + security_engineer
```

### 任务模式推荐
- **标准功能**: Standard 模式（标准化开发流程）
- **创新功能**: Free 模式（探索性开发）
- **混合项目**: Hybrid 模式（标准化 + 探索性）

## 📋 阶段指南

### Discovery（发现阶段）

#### 目标
- 明确项目需求
- 了解用户需求
- 分析市场机会
- 评估技术可行性

#### 典型任务
```bash
# 需求分析任务
picoclaw goagents milestone create requirements-analysis-milestone
picoclaw goagents task create business-analysis

# 用户研究任务
picoclaw goagents task create user-research

# 市场分析任务
picoclaw goagents task create market-analysis
```

#### 预期输出
- 需求分析报告
- 用户研究报告
- 市场分析报告
- 技术可行性评估

### Architecture（架构阶段）

#### 目标
- 设计系统架构
- 制定技术方案
- 设计数据库结构
- 规划开发计划

#### 典型任务
```bash
# 系统架构设计
picoclaw goagents milestone create architecture-design-milestone
picoclaw goagents task create system-architecture

# 技术设计
picoclaw goagents task create technical-design

# 数据库设计
picoclaw goagents task create database-design
```

#### 预期输出
- 系统架构文档
- 技术设计文档
- 数据库设计文档
- 开发计划

### Development（开发阶段）

#### 目标
- 实现系统功能
- 编写测试用例
- 进行集成测试
- 完善文档

#### 典型任务
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

#### 预期输出
- 功能代码
- 测试代码
- 技术文档
- 用户手册

### Validation（验证阶段）

#### 目标
- 验证功能完整性
- 测试系统性能
- 评估用户体验
- 准备项目交付

#### 典型任务
```bash
# 功能测试
picoclaw goagents milestone create validation-milestone
picoclaw goagents task create functional-testing

# 性能测试
picoclaw goagents task create performance-testing

# 用户验收测试
picoclaw goagents task create user-acceptance-testing
```

#### 预期输出
- 测试报告
- 性能报告
- 用户验收报告
- 交付文档

## 🎯 具体使用示例

### 示例1：企业管理系统
```bash
# 启动项目
@go "开发企业管理系统"

# PO 分析结果：
# 📋 **PO任务分析结果**
# 
# **需求**: 开发企业管理系统
# **推荐模式**: Standard（标准化企业应用）
# **任务分解**:
# - Discovery: 需求分析 + 用户调研
# - Architecture: 系统架构 + 技术选型
# - Development: 功能开发 + 测试
# - Validation: 质量验证 + 部署
```

### 示例2：数据分析平台
```bash
# 启动项目
@go "开发数据分析平台"

# PO 分析结果：
# 📋 **PO任务分析结果**
# 
# **需求**: 开发数据分析平台
# **推荐模式**: Hybrid（标准化 + 创新功能）
# **任务分解**:
# - 标准化任务: 基础平台开发
# - 探索性任务: 创新分析算法
```

## 🚀 快速开始

### 1. 初始化项目
```bash
# 初始化通用项目配置
picoclaw goagents config init
picoclaw goagents phase use default-discovery
picoclaw goagents team use default-discovery-team
```

### 2. 启动开发
```bash
# 启动项目开发
@go "开发[项目名称]"

# 监控执行状态
picoclaw goagents status
```

### 3. 质量检查
```bash
# 检查项目质量
picoclaw goagents quality check

# 查看质量报告
picoclaw goagents quality report
```

## 📊 质量保证

### 质量标准
- **代码质量**: ≥ 85%
- **测试覆盖率**: ≥ 80%
- **文档完整性**: ≥ 90%
- **用户满意度**: ≥ 85%

### 质量门禁
- 每个阶段都有质量门禁
- 不通过则不能进入下一阶段
- 支持质量改进和重新检查

## 🎯 最佳实践

### 1. 明确项目目标
- 设定清晰的项目目标
- 定义成功标准
- 建立验收标准

### 2. 合理规划时间
- 评估项目复杂度
- 制定合理的时间计划
- 预留缓冲时间

### 3. 保证质量
- 遵循质量标准
- 进行充分测试
- 持续改进质量

### 4. 有效沟通
- 建立沟通机制
- 定期同步进度
- 及时解决问题

---

**通用项目指南为Go Agents用户提供了标准化的项目开发流程，通过配置驱动架构确保项目的高质量交付！** 🚀
