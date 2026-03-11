# Go Agents 配置系统使用指南

## 🎯 概述

Go Agents 配置系统提供了完整的 Workflow、Phase、Team、Task 配置管理能力，支持创建、更新、删除、克隆等操作。

## 📁 配置文件结构

```
~/.picoclaw/workspace/.goagents/
├── config.yaml               # 全局配置
├── workflows/                # 工作流配置
│   ├── standard-development.yaml
│   ├── mobile-app-development.yaml
│   └── custom-workflow.yaml
├── phases/                   # 阶段配置
│   ├── discovery.yaml
│   ├── planning.yaml
│   ├── development.yaml
│   └── validation.yaml
├── teams/                    # 团队配置
│   ├── discovery-team.yaml
│   ├── development-team.yaml
│   └── custom-team.yaml
├── tasks/                    # 任务模板
│   ├── market-analysis.yaml
│   ├── requirement-gathering.yaml
│   ├── api-development.yaml
│   └── custom-task.yaml
├── registry/                 # 配置管理器
│   ├── config-loader.go
│   ├── config-manager.go
│   └── types.go
└── cli/                      # 命令行工具
    └── commands.go
```

## 🚀 快速开始

### 1. 查看可用配置
```bash
# 列出所有工作流
goagents workflow list

# 列出所有阶段
goagents phase list

# 列出所有团队
goagents team list

# 列出所有任务
goagents task list
```

### 2. 创建新配置
```bash
# 创建新工作流
goagents workflow create my-workflow

# 创建新阶段
goagents phase create my-phase

# 创建新团队
goagents team create my-team

# 创建新任务
goagents task create my-task
```

### 3. 更新现有配置
```bash
# 更新工作流
goagents workflow update my-workflow

# 更新阶段
goagents phase update my-phase

# 更新团队
goagents team update my-team

# 更新任务
goagents task update my-task
```

### 4. 克隆配置
```bash
# 克隆工作流
goagents workflow clone standard-development my-custom-workflow

# 克隆阶段
goagents phase clone discovery my-discovery

# 克隆团队
goagents team clone discovery-team my-team

# 克隆任务
goagents task clone market-analysis my-custom-analysis
```

### 5. 删除配置
```bash
# 删除工作流
goagents workflow delete my-workflow

# 删除阶段
goagents phase delete my-phase

# 删除团队
goagents team delete my-team

# 删除任务
goagents task delete my-task
```

## 🔧 配置定制

### 工作流定制

#### 创建自定义工作流
```yaml
# ~/.picoclaw/workspace/.goagents/workflows/my-workflow.yaml
workflow:
  id: "my-workflow"
  name: "我的自定义工作流"
  description: "针对特定项目需求的自定义工作流"
  version: "1.0.0"
  
phases:
  - id: "custom-discovery"
    name: "自定义需求发现"
    team_config: "my-discovery-team"
    task_mode: "hybrid"
    estimated_duration: "1-2周"
    
  - id: "custom-development"
    name: "自定义开发"
    team_config: "my-development-team"
    task_mode: "standard"
    estimated_duration: "2-3周"
```

#### 更新现有工作流
```bash
# 更新工作流配置
goagents workflow update standard-development

# 然后编辑生成的配置文件，或使用交互式更新
```

### 阶段定制

#### 创建自定义阶段
```yaml
# ~/.picoclaw/workspace/.goagents/phases/my-phase.yaml
phase:
  id: "my-phase"
  name: "我的自定义阶段"
  description: "针对特定需求的阶段定义"
  
objectives:
  primary:
    - "完成特定目标"
    - "满足特定需求"
    
standard_tasks:
  - template_id: "my-custom-task"
    name: "自定义任务"
    estimated_time: "1-2天"
    
free_mode_config:
  phase_lead_authority:
    task_planning: true
    task_assignment: true
```

### 团队定制

#### 创建自定义团队
```yaml
# ~/.picoclaw/workspace/.goagents/teams/my-team.yaml
team:
  id: "my-team"
  name: "我的自定义团队"
  description: "针对特定项目需求的团队配置"
  
team_lead:
  agent: "po"
  variant: "senior_product_manager"
  
phases:
  custom_phase:
    phase_lead:
      agent: "analyst"
      variant: "domain_expert"
      
    team_members:
      - member_id: "specialist_1"
        agent: "analyst"
        variant: "technical_specialist"
```

### 任务定制

#### 创建自定义任务
```yaml
# ~/.picoclaw/workspace/.goagents/tasks/my-task.yaml
task:
  id: "my-task"
  name: "我的自定义任务"
  description: "针对特定需求的任务定义"
  
execution_steps:
  - step_id: "custom_step_1"
    name: "自定义步骤1"
    estimated_time: "2小时"
    activities:
      - "自定义活动1"
      - "自定义活动2"
    outputs:
      - name: "custom_output_1"
        format: "markdown"
```

## 🔄 配置加载机制

### 自动加载
配置系统会按以下优先级自动加载配置：

1. **用户自定义配置** (`~/.picoclaw/workspace/.goagents/`)
2. **项目级配置** (`./.goagents/`)
3. **系统默认配置** (内置)

### 缓存机制
- **缓存时间**: 5分钟
- **缓存大小**: 100个配置对象
- **自动刷新**: 配置文件修改时自动清除缓存

### 验证机制
每个配置加载时都会进行：
- **格式验证**: YAML 格式正确性
- **完整性验证**: 必需字段完整性
- **逻辑验证**: 配置逻辑一致性

## 🎛️ 高级功能

### 配置模板系统
```bash
# 从模板创建配置
goagents workflow create --template standard-development my-workflow
goagents phase create --template discovery my-phase
goagents team create --template discovery-team my-team
goagents task create --template market-analysis my-task
```

### 配置导入导出
```bash
# 导出配置
goagents export workflow standard-development --format yaml > workflow.yaml
goagents export all --format json > all-configs.json

# 导入配置
goagents import workflow.yaml
goagents import all-configs.json
```

### 配置验证
```bash
# 验证所有配置
goagents validate --all

# 验证特定配置
goagents validate workflow standard-development
goagents validate phase discovery
goagents validate team discovery-team
goagents validate task market-analysis
```

### 配置合并
```bash
# 合并配置
goagents merge workflow base.yaml custom.yaml --output merged.yaml
goagents merge team base-team.yaml custom-team.yaml --output merged-team.yaml
```

## 🔍 配置查询

### 条件查询
```bash
# 查询特定条件的配置
goagents query workflows --category development
goagents query phases --category analysis
goagents query teams --tag research
goagents query tasks --difficulty medium
```

### 依赖关系查询
```bash
# 查询配置依赖关系
goagents dependencies workflow standard-development
goagents dependencies phase discovery
goagents dependencies team discovery-team
```

### 使用统计查询
```bash
# 查询配置使用统计
goagents stats workflow standard-development
goagents stats --most-used
goagents stats --success-rate
```

## 🛠️ 集成开发

### API 接口
```go
// 使用配置管理器
import "github.com/sipeed/picoclaw/.goagents/registry"

cm := registry.NewConfigManager("/path/to/.goagents")

// 加载配置
workflow, err := cm.LoadWorkflow("standard-development")
if err != nil {
    log.Fatal(err)
}

// 更新配置
updates := map[string]interface{}{
    "name": "Updated Workflow Name",
    "description": "Updated description",
}
err = cm.UpdateWorkflow("standard-development", updates)
if err != nil {
    log.Fatal(err)
}
```

### 技能集成
```go
// 在技能中使用配置
func (s *MySkill) Execute() error {
    // 加载工作流配置
    workflow, err := s.configManager.LoadWorkflow(s.workflowID)
    if err != nil {
        return err
    }
    
    // 根据配置执行逻辑
    return s.executeWorkflow(workflow)
}
```

## 📊 监控和日志

### 配置变更日志
```bash
# 查看配置变更历史
goagents log workflow --last 10
goagents log phase --since "2026-03-01"
goagents log team --user "current-user"
```

### 性能监控
```bash
# 查看配置加载性能
goagents monitor --performance
goagents monitor --cache-hit-rate
goagents monitor --load-time
```

## 🔧 故障排除

### 常见问题

#### 配置文件不存在
```bash
# 错误：配置文件不存在
Error: failed to read workflow config: file does not exist

# 解决：检查文件路径和权限
ls -la ~/.picoclaw/workspace/.goagents/workflows/
goagents validate --all
```

#### 配置格式错误
```bash
# 错误：YAML 格式错误
Error: failed to parse workflow config: yaml: line 5: mapping values are not allowed in this context

# 解决：验证 YAML 格式
goagents validate workflow my-workflow
yamllint ~/.picoclaw/workspace/.goagents/workflows/my-workflow.yaml
```

#### 配置验证失败
```bash
# 错误：配置验证失败
Error: workflow validation failed: workflow ID is required

# 解决：检查必需字段
goagents validate --verbose workflow my-workflow
```

### 调试模式
```bash
# 启用调试模式
export GOAGENTS_DEBUG=true
goagents workflow list

# 查看详细加载过程
goagents workflow load my-workflow --verbose
```

## 📚 最佳实践

### 配置组织
1. **命名规范**: 使用 kebab-case 命名
2. **版本控制**: 为每个配置维护版本号
3. **文档完整**: 为每个配置提供详细描述
4. **分类管理**: 按项目、类型、用途分类

### 变更管理
1. **渐进式变更**: 避免大规模一次性变更
2. **向后兼容**: 保持配置格式向后兼容
3. **测试验证**: 变更后充分测试
4. **备份恢复**: 重要变更前备份

### 性能优化
1. **合理缓存**: 设置合适的缓存时间
2. **延迟加载**: 按需加载配置
3. **批量操作**: 批量处理配置操作
4. **监控指标**: 监控配置加载性能

---

通过这个配置系统，你可以灵活地管理和定制 Go Agents 的所有配置，实现完全个性化的多 Agent 协作平台！
