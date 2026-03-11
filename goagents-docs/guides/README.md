# 📚 Go Agents 指南索引

`goagents-docs/guides/` 是面向 Go Agents 用户的教学层说明，重点回答"在某类项目里，Go Agents 应该怎么启动、怎么分阶段、怎么选模式"。

## 🎯 阶段定义

为降低歧义，Go Agents 指南统一使用下面这套展示层阶段：

- `Research（调研与机会识别）`
- `Requirements（需求澄清与PRD形成）`
- `Planning（架构设计与实施计划）`
- `Development（开发与联调）`
- `Validation（测试与验收）`

与 PicoClaw 仓库术语的对应关系：

- `Research + Requirements` 对应仓库中的 `discovery`
- `Planning` 对应 `planning`
- `Development` 对应 `development`
- `Validation` 对应 `validation`

## 🔄 任务模式

Task 模式对外统一使用：

- `Standard` - 标准化模板执行
- `Free` - Phase Lead 主导动态规划
- `Hybrid` - 混合执行模式

补充说明：

- `Template` 只表示 Standard 模式下的模板机制，不是单独的正式模式名
- 顶层 [how-to-setup.md](/e:/workspace/picoclaw/goagents-docs/how-to-setup.md) 是完整设置指南，本目录是在其基础上按行业、场景、模式拆分后的子文档

## 📖 先看哪类

- 如果你先按业务类型落地，先看 `industry`
- 如果你先按项目特征判断，先看 `scenario`
- 如果你已经知道要用哪种执行方式，先看 `mode`

## 🏭 Industry

- [电商项目](./industry/ecommerce.md): 适合用户量大、交易复杂、性能要求高的电商项目
- [金融科技项目](./industry/fintech.md): 适合安全合规、风险控制、数据处理要求高的金融项目
- [移动应用项目](./industry/mobile-app.md): 适合用户体验敏感、迭代快、移动端发布导向的项目
- [游戏开发项目](./industry/game-development.md): 适合创意驱动、资源密集、里程碑明显的项目
- [全栈 Web 项目](./industry/fullstack-web.md): 适合架构、安全、DevOps 要求较高的项目
- [AI/ML 项目](./industry/ai-ml.md): 适合数据驱动、模型训练、算法优化的 AI 项目
- [通用项目](./industry/default-project.md): 适合还没有明确行业模板、想先用默认配置启动的项目

## 🎯 Scenario

- [常规项目](./scenario/routine-project.md): 适合需求较清晰、节奏稳定的项目，通常以 Standard 为主
- [创新项目](./scenario/innovation-project.md): 适合探索性强、不确定性高的项目，通常以 Free 为主
- [紧急项目](./scenario/urgent-project.md): 适合时间压力大、需要快速交付的项目，通常以 Hybrid 或 Standard 优先
- [重构项目](./scenario/refactoring-project.md): 适合现有系统重构、技术债务清理的项目
- [集成项目](./scenario/integration-project.md): 适合系统集成、API 对接、数据迁移的项目

## 🔄 Mode

- [Standard 模式](./mode/standard-mode.md): 适合模板化、质量一致性要求高的项目
- [Free 模式](./mode/free-mode.md): 适合需要 Phase Lead 主导规划的创新项目
- [Hybrid 模式](./mode/hybrid-mode.md): 适合标准任务和探索任务并存的项目

## 🚀 推荐入口

- **第一次接触 Go Agents**: 先看 [通用项目](./industry/default-project.md)
- **做电商业务**: 先看 [电商项目](./industry/ecommerce.md)
- **做金融科技**: 先看 [金融科技项目](./industry/fintech.md)
- **做成熟业务功能**: 先看 [常规项目](./scenario/routine-project.md)
- **做探索型需求**: 先看 [创新项目](./scenario/innovation-project.md)
- **已经确定要模板驱动**: 先看 [Standard 模式](./mode/standard-mode.md)

## 📋 快速开始

### 1. 基础设置
```bash
# 完成基础设置
picoclaw goagents config init
picoclaw skills install po-core
picoclaw skills enable po-core
```

### 2. 选择行业模板
```bash
# 选择电商模板
picoclaw goagents workflow create ecommerce-development
picoclaw goagents team create ecommerce-team
```

### 3. 启动项目
```bash
# 启动项目
@go "开发电商购物车功能"
```

## 🔧 高级指南

### 团队配置
- [团队角色配置](./team-configuration.md): 如何配置不同角色的团队
- [协作模式设置](./collaboration-modes.md): 如何设置团队协作模式
- [权限管理](./permission-management.md): 如何管理团队权限

### 配置管理
- [自定义工作流](./custom-workflows.md): 如何创建自定义工作流
- [任务模板设计](./task-template-design.md): 如何设计任务模板
- [质量门禁配置](./quality-gates.md): 如何配置质量门禁

### 监控优化
- [性能监控](./performance-monitoring.md): 如何监控系统性能
- [质量指标](./quality-metrics.md): 如何跟踪质量指标
- [故障排除](./troubleshooting.md): 常见问题和解决方案

## 📊 总参考

- [完整设置指南](/e:/workspace/picoclaw/goagents-docs/how-to-setup.md)
- [实施指南](/e:/workspace/picoclaw/openspec/changes/po-system-implementation/IMPLEMENTATION-GUIDE.md)
- [技能设计文档](/e:/workspace/picoclaw/openspec/changes/po-system-implementation/SKILL-DESIGN.md)

## 🎯 使用建议

1. **新手用户**: 按顺序阅读 设置指南 → 行业模板 → 具体项目
2. **有经验用户**: 直接跳转到对应的行业或场景指南
3. **高级用户**: 参考高级指南进行深度定制

## 📞 获取帮助

- **设置问题**: 参考 [how-to-setup.md](/e:/workspace/picoclaw/goagents-docs/how-to-setup.md)
- **使用问题**: 参考对应行业/场景指南
- **技术问题**: 参考 [troubleshooting.md](./troubleshooting.md)
- **配置问题**: 参考 [custom-workflows.md](./custom-workflows.md)
