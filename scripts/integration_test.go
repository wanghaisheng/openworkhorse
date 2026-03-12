package integration

import (
    "testing"
    "time"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "log/slog"

import (
    "github.com/sipeed/picoclaw/pkg/skills"
    "github.com/sipeed/picoclaw/.goagents/registry"
)

// TestWorkflowToAgentIntegration 测试Workflow到Agent的完整流程
func TestWorkflowToAgentIntegration(t *testing.T) {
    t.Log("开始测试Workflow到Agent集成")
    
    // 1. 创建智能协调器
    coordinator, err := NewIntelligentCoordinator(
        &skills.ExecutionFramework{},
        "product-development-team",
        "product-development",
    )
    assert.NoError(t, "创建智能协调器失败")
    
    // 2. 创建项目请求
    projectRequest := map[string]interface{}{
        "project_name": "AI Specialist团队项目",
        "project_type": "product_development",
        "team_size": 8,
        "timeline": "6个月",
        "budget": "800万",
        "requirements": []string{
            "多租户架构",
            "国际化功能",
            "高级搜索功能",
            "个性化推荐",
            "实时库存管理",
            "多支付网关集成",
            "移动端适配",
            "数据分析仪表板",
        },
        "constraints": []string{
            "技术栈": ["React", "Go", "PostgreSQL", "Redis", "Kubernetes"],
            "合规要求": ["PCI-DSS", "GDPR", "SOX"],
            "团队构成": "混合团队（内部+外部）",
            "交付方式": "渐进式交付",
        },
    }
    
    // 3. 执行智能协调
    start := time.Now()
    result, err := coordinator.ExecuteIntelligentCoordination(
        context.Background(),
        projectRequest,
    )
    duration := time.Since(start)
    
    // 4. 验证结果
    assert.NoError(t, "智能协调执行失败: %v", err)
    assert.True(t, result.Success, "协调应该成功")
    assert.Less(t, duration < 30*time.Second, "协调应该在30秒内完成")
    
    // 5. 验证完整性
    assert.NotNil(t, result.ExecutionPlan, "应该有执行计划")
    assert.NotEmpty(t, result.StepResults, "应该有执行结果")
    assert.NotEmpty(t, result.QualityResults, "应该有质量结果")
    
    // 6. 验证智能决策
    assert.Greater(t, len(result.IntelligenceData.Decisions), 0, "应该有智能决策")
    assert.Greater(t, len(result.IntelligenceData.Predictions), 0, "应该有质量预测")
    assert.Greater(t, len(result.IntelligenceData.LearningInsights), 0, "应该有学习洞察")
    
    t.Logf("✅ Workflow到Agent集成测试通过，耗时: %v", duration)
}

// TestPhaseToPhaseIntegration 测试Phase间协同
func TestPhaseToPhaseIntegration(t *testing.T) {
    t.Log("开始测试Phase到Phase集成")
    
    // 1. 创建智能协调器
    coordinator, err := NewIntelligentCoordinator(
        &skills.ExecutionFramework{},
        "product-development-team",
        "discovery",
    )
    assert.NoError(t, "创建智能协调器失败: %v", err)
    
    // 2. 创建Discovery Phase
    discoveryPhase, err := coordinator.executionFramework.LoadPhaseConfig("discovery")
    assert.NoError(t, "加载Discovery Phase配置失败: %v", err)
    
    // 3. 执行Discovery Phase
    start := time.Now()
    discoveryResult, err := coordinator.ExecutePhase(
        context.Background(),
        discoveryPhase,
        map[string]interface{}{
            "phase_id": "discovery",
            "project_request": "项目发现阶段",
        },
    )
    duration := time.Since(start)
    
    // 4. 验证结果
    assert.NoError(t, "Discovery Phase执行失败: %v", err)
    assert.True(t, discoveryResult.Success, "Discovery Phase应该成功")
    assert.Less(t, duration < 60*time.Second, "Discovery Phase应该在60秒内完成")
    
    // 5. 创建Architecture Phase
    architecturePhase, err := coordinator.executionFramework.LoadPhaseConfig("architecture")
    assert.NoError(t, "加载Architecture Phase配置失败: %v", err)
    
    // 6. 执行Architecture Phase
    start = time.Now()
    architectureResult, err := coordinator.ExecutePhase(
        context.Background(),
        architecturePhase,
        map[string]interface{}{
            "phase_id": "architecture",
            "project_request": "架构设计阶段",
            "discovery_results": discoveryResult,
        },
    )
    duration := time.Since(start)
    
    // 7. 验证协同效果
    assert.NoError(t, "Architecture Phase执行失败: %v", err)
    assert.True(t, architectureResult.Success, "Architecture Phase应该成功")
    assert.Less(t, duration < 60*time.Second, "Architecture Phase应该在60秒内完成")
    
    // 8. 验证数据传递
    assert.Contains(t, architectureResult.Input["discovery_results"], "discovery_results", "应该包含Discovery结果")
    
    t.Logf("✅ Phase到Phase集成测试通过，Discovery: %v, Architecture: %v", 
        discoveryResult.Success, architectureResult.Success)
}

// TestRoleToRoleIntegration 测试Role内协同
func TestRoleToRoleIntegration(t *testing.T) {
    t.Log("开始测试Role内协同")
    
    // 1. 创建智能协调器
    coordinator, err := NewIntelligentCoordinator(
        &skills.ExecutionFramework{},
        "discovery-team",
        "discovery",
    )
    assert.NoError(t, "创建智能协调器失败: %v", err)
    
    // 2. 加载团队配置
    teamConfig, err := coordinator.executionFramework.LoadTeamConfig("discovery-team")
    assert.NoError(t, "加载团队配置失败: %v", err)
    
    // 3. 激活Analyst Role
    analystRole, err := coordinator.executionFramework.LoadRoleConfig("analyst")
    assert.NoError(t, "加载Analyst Role配置失败: %v", err)
    
    // 4. 实例化多个Agent
    agents := make([]*registry.TeamMemberAgent, 0)
    for _, member := range teamConfig.Team.Members {
        if member.Role == "analyst" {
            agent, err := coordinator.CreateTeamMemberAgent(
                member.Agent,
                member.Role,
                member.Configuration,
            )
            assert.NoError(t, "创建Agent实例失败: %v", err)
            agents = append(agents, agent)
        }
    }
    
    // 5. 创建Role内部协调
    roleCoordination := &RoleInternalCoordination{
        RoleID: "analyst",
        Agents: agents,
        LeadAgent: agents[0], // agent_analyst_01作为Lead
        CoordinationPattern: "hierarchical",
    }
    
    // 6. 执行Role内部协同
    task := map[string]interface{}{
        "task_type": "业务分析",
        "description": "完整的业务分析任务",
        "priority": "high",
        "deadline": "2026-03-12",
    }
    
    start := time.Now()
    roleResult, err := roleCoordination.CoordinateRoleInternal(
        context.Background(),
        task,
    )
    duration := time.Since(start)
    
    // 7. 验证协同效果
    assert.NoError(t, "Role内部协同失败: %v", err)
    assert.True(t, roleResult.Success, "Role内部协同应该成功")
    assert.Less(t, duration < 30*time.Second, "Role内部协同应该在30秒内完成")
    
    // 8. 验证质量保证
    assert.NotNil(t, roleResult.QualityValidation, "应该有质量验证结果")
    assert.True(t, roleResult.QualityValidation.OverallPassed, "质量验证应该通过")
    
    // 9. 验证结果整合
    assert.NotEmpty(t, roleResult.IntegratedResult, "应该有整合结果")
    assert.Equal(t, roleResult.LeadAgentID, "agent_analyst_01", "Lead Agent应该是agent_analyst_01")
    
    t.Logf("✅ Role内协同测试通过，耗时: %v", duration)
}

// TestAgentToAgentIntegration 测试Agent间协同
func TestAgentToAgentIntegration(t *testing.T) {
    t.Log("开始测试Agent间协同")
    
    // 1. 创建智能协调器
    coordinator, err := NewIntelligentCoordinator(
        &skills.ExecutionFramework{},
        "discovery-team",
        "discovery",
    )
    assert.NoError(t, "创建智能协调器失败: %v", err)
    
    // 2. 加载团队配置
    teamConfig, err := coordinator.executionFramework.LoadTeamConfig("discovery-team")
    assert.NoError(t, "加载团队配置失败: %v", err)
    
    // 3. 实例化多个Agent
    agents := make([]*registry.TeamMemberAgent, 0)
    for _, member := range teamConfig.Team.Members {
        agent, err := coordinator.CreateTeamMemberAgent(
            member.Agent,
            member.Role,
            member.Configuration,
        )
        assert.NoError(t, "创建Agent实例失败: %v", err)
        agents = append(agents, agent)
    }
    
    // 4. 创建Agent协同
    agentCoordination := &AgentCoordination{
        Agents: agents,
        CoordinationPattern: "collaborative",
        SharedWorkspace: make(map[string]interface{}),
        ConflictResolution: "mediated",
    }
    
    // 5. 执行Agent协同任务
    task := map[string]interface{}{
        "task_type": "业务分析",
        "description": "需要多Agent协同的复杂任务",
        "priority": "high",
        "deadline": "2026-03-12",
    }
    
    start := time.Now()
    agentResult, err := agentCoordination.CoordinateAgentCollaborative(
        context.Background(),
        task,
    )
    duration := time.Since(start)
    
    // 6. 验证协同效果
    assert.NoError(t, "Agent协同失败: %v", err)
    assert.True(t, agentResult.Success, "Agent协同应该成功")
    assert.Less(t, duration < 45*time.Second, "Agent协同应该在45秒内完成")
    
    // 7. 验证结果整合
    assert.NotEmpty(t, agentResult.SharedResults, "应该有共享结果")
    assert.Equal(t, len(agentResult.SharedResults), len(agents), "所有Agent都应该有共享结果")
    
    // 8. 验证质量保证
    assert.NotNil(t, agentResult.QualityValidation, "应该有质量验证结果")
    assert.True(t, agentResult.QualityValidation.OverallPassed, "质量验证应该通过")
    
    t.Logf("✅ Agent间协同测试通过，耗时: %v", duration)
}

// TestConfigurationLoading 集成配置加载测试
func TestConfigurationLoading(t *testing.T) {
    t.Log("开始测试配置加载")
    
    // 1. 测试全局配置加载
    globalConfig, err := coordinator.executionFramework.LoadGlobalConfig()
    assert.NoError(t, "加载全局配置失败: %v", err)
    assert.NotNil(t, globalConfig, "全局配置不能为空")
    
    // 2. 测试Team配置加载
    teamConfig, err := coordinator.executionFramework.LoadTeamConfig("discovery-team")
    assert.NoError(t, "加载团队配置失败: %v", err)
    assert.NotNil(t, teamConfig, "团队配置不能为空")
    
    // 3. 测试Phase配置加载
    phaseConfig, err := coordinator.executionFramework.LoadPhaseConfig("discovery")
    assert.NoError(t, "加载Phase配置失败: %v", err)
    assert.NotNil(t, phaseConfig, "Phase配置不能为空")
    
    // 4. 测试配置一致性
    assert.Equal(t, globalConfig.Version, teamConfig.Version, "全局配置和团队配置版本应该一致")
    assert.Equal(t, globalConfig.Teams.SearchPaths, teamConfig.SearchPaths, "搜索路径应该一致")
    assert.Equal(t, globalConfig.Phases.SearchPaths, phaseConfig.SearchPaths, "搜索路径应该一致")
    
    t.Logf("✅ 配置加载测试通过")
}
```

### **阶段2: 性能优化**
```go
// pkg/performance/performance_test.go
package performance

import (
    "testing"
    "time"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "log/slog"
)

func TestSystemPerformance(t *testing.T) {
    t.Log("开始系统性能测试")
    
    // 1. 创建性能分析器
    analyzer := &performance.PerformanceAnalyzer{}
    report := analyzer.AnalyzeSystemPerformance()
    
    // 2. 验证性能指标
    assert.Greater(t, report.SystemMetrics.Throughput, 10, "系统吞吐量应该大于10个任务/小时")
    assert.Less(t, report.SystemMetrics.ResponseTime, 5.0, "系统响应时间应该小于5秒")
    assert.Less(t, report.SystemMetrics.ErrorRate, 0.05, "系统错误率应该小于5%")
    
    t.Logf("✅ 系统性能测试通过")
    t.Logf("吞吐量: %.2f 任务/小时", report.SystemMetrics.Throughput)
    t.Logf("响应时间: %.2f 秒", report.SystemMetrics.ResponseTime)
    t.Logf("错误率: %.2f%%", report.SystemMetrics.ErrorRate)
}

func TestAgentPerformance(t *testing.T) {
    t.Log("开始Agent性能测试")
    
    // 1. 创建性能分析器
    analyzer := &performance.PerformanceAnalyzer{}
    report := analyzer.AnalyzeAgentMetrics()
    
    // 2. 验证性能指标
    assert.Greater(t, report.AgentMetrics.InstantiationTime, 5.0, "Agent实例化时间应该小于5秒")
    assert.Greater(t, report.AgentMetrics.TaskExecutionTime, 30.0, "任务执行时间应该小于30秒")
    assert.Less(t, report.AgentMetrics.ErrorRate, 0.05, "Agent错误率应该小于5%")
    
    t.Logf("✅ Agent性能测试通过")
    t.Logf("实例化时间: %.2f 秒", report.AgentMetrics.InstantiationTime)
    t.Logf("执行时间: %.2f 秒", report.AgentMetrics.TaskExecutionTime)
    t.Logf("错误率: %.2f%%", report.AgentMetrics.ErrorRate)
}

func TestConfigurationPerformance(t *testing.T) {
    t.Log("开始配置性能测试")
    
    // 1. 创建性能分析器
    analyzer := &performance.PerformanceAnalyzer{}
    report := analyzer.AnalyzeConfigurationPerformance()
    
    // 2. 验证性能指标
    assert.Less(t, report.ConfigurationLoadingTime, 1.0, "配置加载时间应该小于1秒")
    assert.Greater(t, report.ConfigurationParsingTime, 0.5, "配置解析时间应该小于0.5秒")
    assert.Greater(t, report.ConfigurationValidationTime, 0.5, "配置验证时间应该小于0.5秒")
    
    t.Logf("✅ 配置性能测试通过")
    t.Logf("加载时间: %.2f 秒", report.ConfigurationLoadingTime)
    t.Logf("解析时间: %.2f 秒", report.ConfigurationParsingTime)
    t.Logf("验证时间: %.2f 秒", report.ConfigurationValidationTime)
}
```

### **阶段3: 用户体验优化**

#### 步骤1: CLI工具优化
```bash
# 创建CLI优化脚本
# scripts/cli_optimization.sh

#!/bin/bash
# CLI优化脚本
set -e

echo "🔧 优化CLI工具..."

# 1. 添加便捷命令
echo "添加便捷命令..."
goagents role create --interactive
goagents agent create --interactive
goagents coordinator test --verbose
goagents workflow test --verbose

# 2. 优化命令行界面
echo "优化命令行界面..."
goagents role list --format table
goagents agent list --format json
goagents workflow list --format tree

# 3. 优化帮助文档
echo "优化帮助文档..."
goagents help --verbose
goagents role --help
goagents agent --help
```

#### 步骤2: Web界面创建
```bash
# 创建Web界面
mkdir -p web/frontend
mkdir -p web/backend
mkdir -p web/static

# 创建前端应用
cd web/frontend
npm create react-app frontend
cd frontend
npm install
npm start

# 创建后端API
cd ../backend
npm install express
npm install cors
npm install morgan
npm install helmet
npm start
```

### **阶段4: 文档完善**

#### 步骤1: 更新用户指南
```markdown
# 更新用户指南
# goagents-docs/user-guide.md

## 🎯 快速开始

### 1. 创建团队配置
```bash
# 创建团队配置
goagents team create --name "my-team" --template "product-development"

# 添加团队成员
goagents agent create --role "analyst" --experience "senior"
goagents agent create --role "architect" --experience "senior"
goagents agent create --role "developer" --experience "mid"
goagents agent create --role "qa" --experience "junior"
```

### 2. 创建Phase配置
```bash
# 创建Phase配置
goagents phase create --name "discovery" --template "discovery"
goagents phase create --name "architecture" --template "architecture"
```

### 3. 执行Workflow
```bash
# 执行Workflow
goagents workflow execute --team "my-team" --phase "discovery"
goagents workflow execute --team "my-team" --phase "architecture"
goagents workflow execute --team "my-team" --phase "development"
goagents workflow execute --team "my-team" --phase "validation"
```

### 4. 监控和优化
```bash
# 监控系统状态
goagents status --verbose
goagents metrics --real-time

# 优化系统性能
goagents optimize --performance
goagents cache --clear
goagents cache --warmup
```

#### 步骤2: 更新开发者文档
```markdown
# 更新开发者指南
# goagents-docs/developer-guide.md

## 🛠️ 开发者指南

### 1. 系统架构
```
Workflow -> Phase -> Milestone -> Task -> Team -> Team Role -> Team Member Agent
```

### 2. 核心组件
- **IntelligentCoordinator**: 智能协调器
- **TeamManager**: 团队管理器
- **AgentFactory**: Agent实例工厂
- **ConfigManager**: 配置管理器
- **QualityManager**: 质量管理器

### 3. 执行行流程
```go
// 创建智能协调器
coordinator, err := NewIntelligentCoordinator(
    ef, 
    "product-development-team", 
    "product-development",
)

// 执行Workflow
result, err := coordinator.ExecuteIntelligentCoordination(
    context.Background(), 
    projectRequest,
)
```

### 4. 扩展指南
```go
// 添加新的Team Role
role := &registry.TeamRole{
    ID: "custom-role",
    Name: "自定义角色",
    Responsibilities: ["自定义职责"],
    Capabilities: ["自定义能力"],
}

// 添加新的Team Member Agent
agent := &registry.TeamMemberAgent{
    ID: "custom-agent",
    RoleID: "custom-role",
    Personalization: PersonalizationConfig{
        ExperienceLevel: "senior",
        Specialization: "custom-specialization",
    },
}
```

### 5. 最佳实践
```markdown
# 最佳实践
- 使用版本控制管理配置
- 使用环境分离配置
- 定期备份配置
- 验证配置有效性
- 定期性能监控
```

### 6. 错误处理
```bash
# 故障排除
goagents diagnose --team "my-team"
goagents diagnose --phase "discovery"
goagents diagnose --role "analyst"
goagents diagnose --agent "agent_analyst_01"
```

#### 步骤3: 更新管理文档
```markdown
# 更新管理指南
# goagents-docs/management-guide.md

## 🛠️ 管理指南

### 1. 团队管理
```yaml
# 团队配置示例
team:
  id: "product-development-team"
  members:
    - member_id: "agent_analyst_01"
      agent: "agent_analyst_01"
      role: "analyst"
      responsibilities: ["业务分析", "市场调研"]
    - member_id: "agent_architect_01"
      agent: "agent_architect_01"
      role: "architect"
      responsibilities: ["架构设计", "技术选型"]
```

### 2. 性能监控
```bash
# 性能监控
goagents metrics --real-time
goagents performance --analyze
goagents cache --stats
goagents resource --usage
```

### 3. 质量管理
```bash
# 质量检查
goagents quality check --team "my-team"
goagents quality check --phase "discovery"
goagents quality check --role "analyst"
goagents quality check --agent "agent_analyst_01"
```

### 4. 故障排除
```bash
# 故障排除
goagents diagnose --team "my-team"
goagents diagnose --phase "discovery"
goagents diagnose --role "analyst"
goagents diagnose --agent "agent_analyst_01"
```

#### 步骤4: 更新培训材料
```markdown
# 更新培训材料
# goagents-docs/training-materials.md

## 🎓 培训材料

### 1. 基础培训
- **Workflow概念**: Workflow、Phase、Milestone、Task
- **Team Role概念**: Team Role定义和职责
- **Agent实例化**: Agent实例化和个性化
- **智能协调**: 智能协调和质量保证
- **质量保证**: 多层次质量保证

### 2. 进阶培训
- **配置管理**: 高级配置管理技巧
- **性能优化**: 系统性能优化技巧
- **质量保证**: 高级质量保证技巧
- **扩展开发**: 系统扩展技巧

### 3. 实践案例
- **电商项目**: 电商项目完整案例
- **金融项目**: 金融项目完整案例
- **教育项目**: 教育项目完整案例
- **企业项目**: 企业项目完整案例
```

## 🎉 总结

### **实施优先级**
1. **系统集成验证** - 确保系统完整性
2. **性能优化** - 确保系统效率
3. **用户体验优化** - 确保易用性
4. **文档完善** - 确保文档完整性

### **预期效果**
- 🎯 **系统完整性**: 100%功能集成
- 🎯 **性能效率**: 提升50%性能
- 🎯 **用户体验**: 90%用户满意度
- 🎯 **文档完整性**: 100%文档覆盖

---

## 🎉 实施时间表

| 阶段 | 时间 | 主要任务 | 预期结果 |
|------|------|----------|----------|
| 1 | 1周 | 系统集成验证 | 95%集成测试通过率 |
| 2 | 2周 | 性能优化 | 50%性能提升 |
| 3 | 1周 | 用户体验优化 | 90%用户满意度 |
| 4 | 1周 | 文档完善 | 100%文档覆盖 |

---

**这个实施计划确保了AI Specialist团队系统的完整实现，提供了详细的实施步骤和预期效果！** 🚀

所有阶段都有明确的任务清单、验证标准和预期结果，确保实施过程的可控和可验证！
