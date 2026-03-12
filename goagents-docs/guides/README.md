# 📚 Go Agents 指南索引

`goagents-docs/guides/` 是面向 Go Agents 用户的教学层说明，重点回答"在某类项目里，Go Agents 应该怎么启动、怎么分阶段、怎么选模式"。

## 🎯 阶段定义

为降低歧义，Go Agents 指南统一使用下面这套展示层阶段：

- `Discovery（发现阶段）` - 需求分析、用户研究、市场分析
- `Architecture（架构阶段）` - 系统架构设计、技术设计、数据库设计
- `Development（开发阶段）` - 前端开发、后端开发、数据库开发、API实现
- `Validation（验证阶段）` - 功能测试、性能测试、安全测试、可访问性测试

与 PicoClaw 仓库术语的对应关系：

- `Discovery` 对应仓库中的 `discovery`
- `Architecture` 对应仓库中的 `architecture`
- `Development` 对应仓库中的 `development`
- `Validation` 对应仓库中的 `validation`

## 🔄 任务模式

Task 模式对外统一使用：

- `Standard` - 标准化模板执行
- `Free` - Phase Lead 主导动态规划
- `Hybrid` - 混合执行模式

补充说明：

- **配置驱动**: 所有模式都基于配置文件驱动，无需手动编程
- **AI Specialist Team**: 基于AI Specialist团队理念的专业角色协作
- **智能协调**: 通过IntelligentCoordinator实现智能任务分配和协调
- **质量保证**: 多层次质量检查和质量门禁机制

## 📖 先看哪类

- 如果你先按业务类型落地，先看 `industry`
- 如果你先按项目特征判断，先看 `scenario`
- 如果你已经知道要用哪种执行方式，先看 `mode`
- 如果你想了解配置驱动的架构，先看 `architecture`

## 🏗️ Architecture

- [配置驱动架构](./architecture/config-driven-architecture.md): 了解Go Agents的配置驱动架构设计
- [AI Specialist团队](./architecture/ai-specialist-team.md): 了解AI Specialist团队理念和角色定义
- [智能协调系统](./architecture/intelligent-coordination.md): 了解智能协调和任务分配机制
- [质量保证系统](./architecture/quality-assurance.md): 了解多层次质量保证机制

## � Platform Integration

- [平台集成指南](./platform-integration.md): 了解如何通过企业协作平台调用Go Agents
- [飞书集成指南](./feishu-integration.md): 详细的飞书平台集成说明
- [钉钉集成](./platform-integration.md#钉钉集成): 钉钉平台集成配置
- [企业微信集成](./platform-integration.md#企业微信集成): 企业微信平台集成配置
- [Slack集成](./platform-integration.md#slack集成): Slack平台集成配置

## � Industry

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

- **第一次接触 Go Agents**: 先看 [配置驱动架构](./architecture/config-driven-architecture.md)
- **了解AI Specialist团队**: 先看 [AI Specialist团队](./architecture/ai-specialist-team.md)
- **平台集成用户**: 先看 [平台集成指南](./platform-integration.md)
- **飞书用户**: 先看 [飞书集成指南](./feishu-integration.md)
- **做电商业务**: 先看 [电商项目](./industry/ecommerce.md)
- **做金融科技**: 先看 [金融科技项目](./industry/fintech.md)
- **做成熟业务功能**: 先看 [常规项目](./scenario/routine-project.md)
- **做探索型需求**: 先看 [创新项目](./scenario/innovation-project.md)
- **已经确定要模板驱动**: 先看 [Standard 模式](./mode/standard-mode.md)

## 📋 快速开始

### 1. 基础设置
```bash
# 完成基础设置（配置驱动）
picoclaw goagents config init
picoclaw goagents phase setup discovery
picoclaw goagents team setup discovery-team
```

### 2. 平台集成设置
```bash
# 配置飞书集成
picoclaw channel enable feishu
picoclaw channel config feishu --app-id="cli_xxx" --app-secret="xxx"

# 配置其他平台
picoclaw channel enable dingtalk
picoclaw channel enable slack
```

### 3. 选择行业模板
```bash
# 选择电商模板（配置驱动）
picoclaw goagents workflow create ecommerce-development
picoclaw goagents milestone create requirements-analysis
picoclaw goagents task create business-analysis
```

### 4. 启动项目
```bash
# 启动项目（配置驱动）
@go "开发电商购物车功能"

# 或通过飞书等平台
# 在飞书中发送：@go "开发电商购物车功能"
```

### 5. 监控执行
```bash
# 监控执行状态
picoclaw goagents status
picoclaw goagents quality check

# 或通过平台查询
# 在飞书中发送：@go "查看项目状态"
```

## 🔧 高级指南

### 团队配置
- [团队角色配置](./team-configuration.md): 如何配置不同角色的团队
- [协作模式设置](./collaboration-modes.md): 如何设置团队协作模式
- [权限管理](./permission-management.md): 如何管理团队权限

### 配置管理
- [自定义工作流](./custom-workflows.md): 如何创建自定义工作流
- [Phase配置设计](./phase-configuration.md): 如何设计Phase配置
- [Milestone配置设计](./milestone-configuration.md): 如何设计Milestone配置
- [Team配置设计](./team-configuration.md): 如何设计Team配置

### 平台集成
- [平台集成指南](./platform-integration.md): 如何集成到企业协作平台
- [飞书集成](./feishu-integration.md): 详细的飞书集成说明
- [钉钉集成](./platform-integration.md#钉钉集成): 钉钉平台集成配置
- [企业微信集成](./platform-integration.md#企业微信集成): 企业微信平台集成配置

### 监控优化
- [性能监控](./performance-monitoring.md): 如何监控系统性能
- [质量指标](./quality-metrics.md): 如何跟踪质量指标
- [故障排除](./troubleshooting.md): 常见问题和解决方案

## 📊 总参考

- [完整设置指南](/e:/workspace/picoclaw/goagents-docs/how-to-setup.md)
- [实施指南](/e:/workspace/picoclaw/openspec/changes/po-system-implementation/IMPLEMENTATION-GUIDE.md)
- [配置驱动实现](/e:/workspace/picoclaw/openspec/changes/po-system-implementation/MILESTONE_TO_TASK_IMPLEMENTATION.md)
- [Task模板优化](/e:/workspace/picoclaw/openspec/changes/po-system-implementation/TASK_TEMPLATE_OPTIMIZATION.md)

## 🎯 使用建议

1. **新手用户**: 按顺序阅读 配置架构 → 行业模板 → 具体项目
2. **有经验用户**: 直接跳转到对应的行业或场景指南
3. **高级用户**: 参考高级指南进行深度定制
4. **开发者**: 参考实施指南了解技术细节

## 📞 获取帮助

- **设置问题**: 参考 [how-to-setup.md](/e:/workspace/picoclaw/goagents-docs/how-to-setup.md)
- **架构问题**: 参考 [配置驱动架构](./architecture/config-driven-architecture.md)
- **平台集成问题**: 参考 [平台集成指南](./platform-integration.md)
- **飞书集成问题**: 参考 [飞书集成指南](./feishu-integration.md)
- **使用问题**: 参考对应行业/场景指南
- **技术问题**: 参考 [troubleshooting.md](./troubleshooting.md)
- **配置问题**: 参考 [自定义工作流](./custom-workflows.md)

## 🔄 更新说明

### v2.0 配置驱动架构更新
- ✅ **配置驱动**: 完全基于配置文件驱动，无需手动编程
- ✅ **AI Specialist Team**: 基于AI Specialist团队理念的专业角色协作
- ✅ **智能协调**: 通过IntelligentCoordinator实现智能任务分配和协调
- ✅ **质量保证**: 多层次质量检查和质量门禁机制
- ✅ **Phase系统**: Discovery、Architecture、Development、Validation四个阶段
- ✅ **Team系统**: 基于角色的专业团队配置
- ✅ **Task模板**: 通用模板支持动态Task生成
- ✅ **平台集成**: 支持飞书、钉钉、企业微信、Slack等平台
- ✅ **企业协作**: 无缝集成到企业协作平台中
