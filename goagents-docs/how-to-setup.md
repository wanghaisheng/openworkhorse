# 🚀 Go Agents 完整设置指南

## 🎯 概述

本指南基于已实现的 PO 系统变更，提供完整的 Go Agents 设置流程，包括 CLI 混合集成、技能系统配置和工作流管理。

## 📋 前置要求

### 系统要求
- **PicoClaw**: 已安装最新版本
- **Go**: 1.19+ (如果需要编译)
- **权限**: 工作区读写权限

### 验证环境
```bash
# 检查 PicoClaw 版本
picoclaw version

# 检查工作区
picoclaw status
```

## 🏗️ 完整设置流程

### Phase 1: CLI 混合集成设置

#### 1.1 验证 CLI 集成
```bash
# 检查 goagents 命令是否可用
picoclaw goagents --help

# 预期输出：
# Go Agents configuration management for PO system
#
# Usage:
#   picoclaw goagents [command]
#
# Available Commands:
#   workflow    Manage workflow configurations
#   phase       Manage phase configurations
#   team        Manage team configurations
#   task        Manage task configurations
#   config      Manage global configuration
```

#### 1.2 验证技能系统
```bash
# 检查技能系统
picoclaw skills list

# 预期输出应包含：
# - po-core
# - task-modes
# - team-roles
# - phase-templates
```

### Phase 2: PO 核心系统设置

#### 2.1 安装 PO 核心技能
```bash
# 安装 PO 核心技能
picoclaw skills install po-core

# 安装任务模式技能
picoclaw skills install standard-mode
picoclaw skills install free-mode
picoclaw skills install hybrid-mode

# 安装团队角色技能
picoclaw skills install role-analyst
picoclaw skills install role-architect
picoclaw skills install role-developer
picoclaw skills install role-qa

# 启用 PO 核心技能
picoclaw skills enable po-core
```

#### 2.2 验证 @go 触发机制
```bash
# 测试 @go 触发
@go "测试 PO 系统"

# 预期输出：
# 📋 **PO任务分析结果**
# 
# **需求**: 测试 PO 系统
# **类型**: 系统验证 (system_test)
# **复杂度**: 低 (0.2)
# **推荐模式**: Standard（标准化测试）
# **预估工期**: 1-2天
# 
# **阶段规划**:
# 1. discovery (1天) - 系统验证和需求确认
# 2. planning (1天) - 实施计划制定
# 
# **团队配置**:
# - discovery: analyst(主导) + qa(支持)
# 
# 是否开始执行？[Y/n/customize]
```

### Phase 3: 配置系统设置

#### 3.1 初始化配置系统
```bash
# 创建配置目录
mkdir -p ~/.picoclaw/workspace/.goagents/{workflows,phases,teams,tasks}

# 初始化全局配置
picoclaw goagents config init

# 验证配置
picoclaw goagents config validate
```

#### 3.2 设置标准工作流
```bash
# 创建标准开发工作流
picoclaw goagents workflow create standard-development

# 验证工作流
picoclaw goagents workflow list

# 预期输出：
# Available Workflows:
#   - standard-development
#   - mobile-app-development
#   - custom-workflow
```

#### 3.3 设置团队配置
```bash
# 创建发现阶段团队
picoclaw goagents team create discovery-team

# 创建开发阶段团队
picoclaw goagents team create development-team

# 验证团队配置
picoclaw goagents team list
```

#### 3.4 设置任务模板
```bash
# 创建市场分析任务
picoclaw goagents task create market-analysis

# 创建需求收集任务
picoclaw goagents task create requirement-gathering

# 验证任务模板
picoclaw goagents task list
```

### Phase 4: HARNESS.md 集成设置

#### 4.1 创建 HARNESS.md
```bash
# 在项目根目录创建 HARNESS.md
cat > HARNESS.md << 'EOF'
# HARNESS.md - 项目核心约束

## 禁止清单
- ❌ 禁止全局变量
- ❌ 禁止循环依赖
- ❌ 禁止直接数据库操作（必须走service层）
- ❌ 禁止硬编码配置值
- ❌ 禁止忽略错误处理

## 必须遵守
- ✅ 所有新功能必须有单元测试+集成测试覆盖≥85%
- ✅ 新模块必须有接口定义
- ✅ 严格分层（controller → service → repository）
- ✅ 所有公共方法必须有文档注释
- ✅ 错误处理必须完整和一致

## 模块职责矩阵
| 目录 | 职责 | 依赖方向 |
|------|------|---------|
| /domain/ | 核心业务实体与规则 | 无外部依赖 |
| /application/ | 用例协调器 | 只依赖domain |
| /adapters/ | 外部接口实现 | 只依赖application |
| /infrastructure/ | 基础设施实现 | 只依赖application |

## 命名规范
- 文件名：kebab-case
- 类名：PascalCase
- 函数名：camelCase
- 常量：UPPER_SNAKE_CASE
- 接口名：以 I 开头

## 代码质量标准
- 圈复杂度：≤10
- 函数长度：≤50行
- 文件长度：≤500行
- 测试覆盖率：≥85%
- Lint 通过率：100%

## Ralph Wiggum Loop
任务完成后必须执行：
1. 运行所有测试 + lint + 类型检查直到 100% 通过
2. 根据 HARNESS.md 自我审查（列出违规项）
3. 如果失败 → 分析根因 → 修复 → 重复
4. 只有在 green + 自我审查通过时才创建 PR
5. 如果架构有变更，更新 HARNESS.md/ADR
6. 提出 2–3 个后续改进建议
EOF
```

#### 4.2 验证 HARNESS.md 集成
```bash
# 测试 HARNESS.md 集成
@go "验证 HARNESS.md 集成"

# 预期输出应包含 HARNESS.md 检查结果
```

## 🎯 使用示例

### 基本使用流程
```bash
# 1. 启动 PO 项目
@go "开发电商购物车功能"

# 2. 查看项目状态
@go --status

# 3. 查看当前阶段
@go --phase-status

# 4. 阶段转换（需要 PO 批准）
@go --phase-transition discovery planning --reason "requirements_complete"
```

### 配置管理流程
```bash
# 1. 创建自定义工作流
picoclaw goagents workflow create ecommerce-development

# 2. 配置自定义团队
picoclaw goagents team create ecommerce-team

# 3. 创建自定义任务
picoclaw goagents task create checkout-feature

# 4. 验证配置
picoclaw goagents config validate
```

### 监控和维护
```bash
# 1. 日常监控
@go --observability --daily

# 2. 性能检查
@go --performance --weekly

# 3. 质量评估
@go --quality --monthly

# 4. 垃圾收集
@go --garbage-collection --weekly
```

## 🔧 高级配置

### 自定义任务模式
```bash
# 设置默认任务模式
picoclaw goagents config set default_task_mode hybrid

# 配置粒度控制
picoclaw goagents config set granularity.sweet_spot_lines [400,600]
picoclaw goagents config set granularity.max_files_affected 10
```

### 自定义质量门禁
```bash
# 设置质量门禁
picoclaw goagents config set quality_gates.test_coverage_min 90
picoclaw goagents config set quality_gates.lint_pass_rate 100
```

### 团队角色定制
```bash
# 自定义角色技能
picoclaw skills edit role-analyst --persona "资深电商分析师"

# 添加新角色
picoclaw skills install role-data-scientist
picoclaw skills enable role-data-scientist
```

## 🚨 故障排除

### 常见问题

#### 1. @go 命令无响应
```bash
# 检查 PO 技能状态
picoclaw skills status po-core

# 重新启用技能
picoclaw skills disable po-core
picoclaw skills enable po-core
```

#### 2. 配置文件错误
```bash
# 验证配置
picoclaw goagents config validate

# 修复配置
picoclaw goagents config repair
```

#### 3. 工作流无法加载
```bash
# 检查工作流文件
picoclaw goagents workflow show standard-development

# 重置工作流
picoclaw goagents workflow reset standard-development
```

### 调试模式
```bash
# 启用调试模式
export PICOLAW_DEBUG=true
export GOAGENTS_DEBUG=true

# 查看详细日志
picoclaw goagents --debug workflow list
```

## 📊 性能优化

### 缓存配置
```bash
# 设置技能缓存
picoclaw goagents config set skills.cache.enabled true
picoclaw goagents config set skills.cache.ttl 3600

# 设置配置缓存
picoclaw goagents config set config.cache.enabled true
```

### 并发配置
```bash
# 设置并发执行
picoclaw goagents config set system.max_concurrent_tasks 3
picoclaw goagents config set system.parallel_execution true
```

## 🎯 最佳实践

### 1. 项目组织
```bash
# 为不同项目创建专用配置
mkdir -p ~/.picoclaw/workspaces/{ecommerce,fintech,healthcare}
picoclaw --workspace ecommerce goagents workflow create ecommerce-flow
```

### 2. 版本管理
```bash
# 备份配置
picoclaw goagents config backup --name "pre-production"

# 恢复配置
picoclaw goagents config restore --name "pre-production"
```

### 3. 团队协作
```bash
# 导出配置
picoclaw goagents export --format yaml > team-config.yaml

# 导入配置
picoclaw goagents import team-config.yaml
```

## 📈 监控指标

### 关键指标
- **任务执行成功率**: ≥85%
- **质量门禁通过率**: ≥90%
- **平均响应时间**: <5秒
- **配置加载时间**: <2秒

### 监控命令
```bash
# 查看系统指标
picoclaw goagents metrics --system

# 查看项目指标
@go --metrics --project

# 查看性能报告
picoclaw goagents report --performance --last 7d
```

## 🎯 成功验证

### 完整验证清单
- [ ] CLI 混合集成正常工作
- [ ] @go 触发机制正常
- [ ] PO 核心技能已启用
- [ ] 配置系统已初始化
- [ ] HARNESS.md 已集成
- [ ] 标准工作流已创建
- [ ] 团队配置已设置
- [ ] 任务模板已创建
- [ ] 质量门禁已配置
- [ ] 监控系统已启用

### 最终测试
```bash
# 运行完整测试
@go "完整系统验证测试"

# 预期输出：
# ✅ PO 系统验证通过
# ✅ CLI 集成正常
# ✅ 配置系统正常
# ✅ 质量门禁正常
# ✅ 监控系统正常
# 
# 🎉 Go Agents 设置完成！
```

## 🎉 完成

恭喜！你已经成功设置了完整的 Go Agents 系统。现在可以：

1. **使用 @go 启动项目**: `@go "你的需求描述"`
2. **管理配置**: `picoclaw goagents config`
3. **监控系统**: `@go --observability`
4. **优化性能**: `picoclaw goagents metrics`

享受工业级的多 Agent 协作体验！🚀