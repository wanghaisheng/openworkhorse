# 🚀 Open AI workhorse - 工业级多 Agent 协作平台

## 🎯 概述

Open AI workhorse  是基于 PicoClaw 的工业级多 Agent 协作平台，通过技能化实现和配置管理，为用户提供完整的项目全生命周期管理能力。

## 🏗️ 核心特性

### ✅ **零侵入性设计**
- 完全基于现有 PicoClaw 技能系统
- 无需修改核心代码，保持超轻量特性
- 向后兼容，不影响现有功能

### 🎭 **多 Agent 协作**
- **PO Core**: 产品经理核心，协调整个项目
- **团队角色**: Analyst、Architect、Developer、QA 等专业角色
- **智能切换**: 通过技能系统实现角色动态切换

### 🔄 **三种任务模式**
- **Standard Mode**: 标准化模板执行，适合常规项目
- **Free Mode**: Phase Lead 主导动态规划，适合创新项目
- **Hybrid Mode**: 混合执行，平衡效率与灵活性

### 🏗️ **完整工作流**
- **Research**: 调研与机会识别
- **Requirements**: 需求澄清与PRD形成
- **Planning**: 架构设计与实施计划
- **Development**: 开发与联调
- **Validation**: 测试与验收

### 🎯 **工业化集成**
- **HARNESS.md**: 集成 Harness Engineering 最佳实践
- **Ralph Wiggum Loop**: 自我迭代闭环机制
- **质量门禁**: 多层质量控制和自动化验证
- **粒度控制**: 300-800行任务粒度优化

## 🚀 快速开始

### 前置要求
- 已安装 PicoClaw 最新版本
- 工作区已初始化
- 基础 CLI 使用经验

### 5分钟快速设置
```bash
# 1. 验证环境
picoclaw version
picoclaw status

# 2. 安装 PO 核心技能
picoclaw skills install po-core
picoclaw skills enable po-core

# 3. 初始化配置系统
picoclaw wh config init

# 4. 验证安装
@go "测试 Go Agents 系统"

# 预期输出：
# 📋 **PO任务分析结果**
# **需求**: 测试 Go Agents 系统
# **推荐模式**: Standard（标准化测试）
# **预估工期**: 1-2天
# 
# 🎉 Go Agents 设置完成！
```

## 📚 文档导航

### 🎯 **新手用户**
1. [设置指南](./how-to-setup.md) - 完整的安装和配置指南
2. [指南索引](./guides/README.md) - 按行业、场景、模式分类的详细指南
3. [故障排除](./guides/troubleshooting.md) - 常见问题和解决方案

### 🏭 **行业指南**
- [电商项目](./guides/industry/ecommerce.md) - 电商专用配置和最佳实践
- [金融科技](./guides/industry/fintech.md) - 金融科技合规和安全要求
- [移动应用](./guides/industry/mobile-app.md) - 移动端开发和发布流程
- [AI/ML 项目](./guides/industry/ai-ml.md) - 数据驱动和模型开发

### 🎯 **场景指南**
- [常规项目](./guides/scenario/routine-project.md) - 需求明确的标准化项目
- [创新项目](./guides/scenario/innovation-project.md) - 探索性强的不确定性项目
- [紧急项目](./guides/scenario/urgent-project.md) - 时间压力大的快速交付项目
- [重构项目](./guides/scenario/refactoring-project.md) - 系统重构和技术债务清理

### 🔄 **模式指南**
- [Standard 模式](./guides/mode/standard-mode.md) - 模板化标准化执行
- [Free 模式](./guides/mode/free-mode.md) - Phase Lead 主导动态规划
- [Hybrid 模式](./guides/mode/hybrid-mode.md) - 标准与探索混合执行

## 🎮 使用示例

### 基本使用
```bash
# 启动项目
@go "开发电商购物车功能"

# 查看状态
@go --status

# 查看阶段进度
@go --phase-status

# 阶段转换
@go --phase-transition discovery planning --reason "requirements_complete"
```

### 配置管理
```bash
# 创建自定义工作流
picoclaw wh workflow create ecommerce-development

# 配置专业团队
picoclaw wh team create ecommerce-team

# 创建任务模板
picoclaw wh task create checkout-feature

# 验证配置
picoclaw wh config validate
```

### 监控和维护
```bash
# 日常监控
@go --observability --daily

# 性能检查
@go --performance --weekly

# 质量评估
@go --quality --monthly

# 垃圾收集
@go --garbage-collection --weekly
```

## 🏗️ 架构设计

### 📁 完整层级结构
```
Workflow
  -> Phase
    -> Milestone
      -> Task
        -> Team
          -> Team Role
            -> Team Member Agent
```

### 📁 目录结构
```
~/.picoclaw/workspace/
├── skills/                     # 技能实现
│   ├── po-core/               # PO 核心技能
│   ├── task-modes/            # 任务模式技能
│   ├── team-roles/            # 团队角色技能
│   └── workflows/             # 完整工作流
└── .wh/                # 配置系统
    ├── config.yaml           # 全局配置
    ├── workflows/            # 工作流配置
    ├── phases/               # 阶段配置（含里程碑）
    ├── teams/                # 团队配置
    ├── tasks/                # 任务模板
    ├── registry/             # 配置管理器
    └── cli/                  # 独立CLI工具
```

### 🔄 集成方式
- **技能系统**: 通过现有技能加载机制
- **CLI 集成**: 混合 `picoclaw wh` 命令
- **配置管理**: 独立的配置系统
- **触发机制**: `@go` 命令触发
- **层级执行**: 完整的 Workflow -> Phase -> Milestone -> Task -> Agent 执行流程

## 📊 性能指标

### 🎯 **目标指标**
- **任务成功率**: Standard 模式 ≥85%，Free 模式 ≥80%
- **代码质量**: 测试覆盖率 ≥85%，lint 通过率 100%
- **执行效率**: 3.5+ PR/工程师/天
- **响应时间**: PO 分析 <5秒

### 📈 **监控指标**
- **系统性能**: CPU、内存、响应时间
- **任务执行**: 成功率、执行时间、质量分数
- **协作效率**: 团队协作指标、沟通效率
- **用户满意度**: 使用体验、反馈评分

## 🎯 适用场景

### 🏭 **企业级应用**
- 大型复杂项目开发
- 多团队协作项目
- 质量要求高的项目
- 需要标准化流程的项目

### 🚀 **创新项目**
- 探索性强的项目
- 技术创新项目
- 业务模式创新
- 快速迭代项目

### 🔄 **维护项目**
- 系统重构项目
- 技术债务清理
- 性能优化项目
- 功能增强项目

## 🎉 核心优势

### ✅ **完全兼容**
- 基于现有 PicoClaw 架构
- 零侵入性设计
- 向后兼容保证
- 渐进式升级

### ✅ **工业级标准**
- Harness Engineering 集成
- Ralph Wiggum Loop 实现
- 质量门禁控制
- 粒度优化管理

### ✅ **高度灵活**
- 三种执行模式
- 自定义配置
- 模板化定制
- 行业适配

### ✅ **易于使用**
- 简单的 `@go` 触发
- 完整的 CLI 工具
- 详细的文档支持
- 友好的错误提示

## 🛠️ 开发和贡献

### 📋 开发环境
```bash
# 克隆仓库
git clone https://github.com/sipeed/picoclaw.git
cd picoclaw

# 安装依赖
go mod download

# 构建
go build ./cmd/picoclaw

# 测试
go test ./...
```

### 🔧 贡献指南
1. Fork 项目
2. 创建功能分支
3. 提交变更
4. 创建 Pull Request
5. 等待代码审查

### 📝 文档贡献
- 改进现有文档
- 添加新的指南
- 修复错误和过时信息
- 翻译文档

## 📞 支持和帮助

### 🆘 **获取帮助**
- [故障排除指南](./guides/troubleshooting.md) - 常见问题解决方案
- [设置指南](./how-to-setup.md) - 详细安装和配置
- [社区论坛](https://github.com/sipeed/picoclaw/discussions) - 社区讨论和帮助
- [GitHub Issues](https://github.com/sipeed/picoclaw/issues) - 问题报告和功能请求

### 📚 **学习资源**
- [官方文档](https://picoclaw.dev/docs) - 完整的官方文档
- [视频教程](https://youtube.com/picoclaw) - 视频教程和演示
- [博客文章](https://blog.picoclaw.dev) - 技术博客和最佳实践
- [社区案例](https://github.com/sipeed/picoclaw/examples) - 社区贡献的案例
- [企业级特性](./enterprise-features.md) - 企业级特性扩展指南

### 🎓 **培训课程**
- [基础课程](https://learn.picoclaw.dev/beginner) - Go Agents 基础使用
- [高级课程](https://learn.picoclaw.dev/advanced) - 高级配置和定制
- [企业培训](https://learn.picoclaw.dev/enterprise) - 企业级部署和管理
- [企业级特性](https://learn.picoclaw.dev/enterprise) - 企业级特性专项培训

## 🗺️ 发展路线图

### 🚀 **当前版本 (v1.0)**
- ✅ PO 核心系统
- ✅ 三种任务模式
- ✅ 基础配置管理
- ✅ CLI 混合集成

### 🎯 **下一版本 (v1.1)**
- 🔄 更多行业模板
- 🔄 高级监控功能
- 🔄 团队协作增强
- 🔄 性能优化
- 🔄 企业级特性扩展
- 🔄 输出物类型支持 (document, code, video, image, 3d_model, prototype, demo)
- 🔄 外部依赖管理

### 🌟 **未来版本 (v2.0)**
- 📋 AI 驱动的模式选择
- 📋 跨项目知识共享
- 📋 企业级权限管理
- 📋 云原生部署支持
- 📋 系统拆分图支持
- 📋 复杂依赖管理
- 📋 行业深度适配

## 📄 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

## 🙏 致谢

感谢所有为 Go Agents 项目做出贡献的开发者、用户和社区成员！

---

**通过 Go Agents，让每个项目都能享受工业级的多 Agent 协作体验！** 🚀
