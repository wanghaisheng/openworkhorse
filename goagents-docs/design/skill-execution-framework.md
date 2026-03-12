# 技能执行框架 - HARNESS.md、WBS 和 OpenSpec 集成

## 概述

技能执行框架为 PicoClaw 的技能系统提供了统一的执行环境，自动集成 HARNESS.md 约束、WBS 工作分解结构和 OpenSpec 规范驱动，以及 Ralph Wiggum Loop 质量保证机制。这样，每个技能在执行时都会自动遵循项目的质量标准和结构化任务分解，无需在每个技能定义中重复这些约束。

## 🎯 核心特性

### 1. 自动 HARNESS.md 集成
- 自动加载项目的 HARNESS.md 文件
- 将约束规则注入到技能执行过程中
- 支持自定义约束规则解析

### 2. WBS 工作分解结构集成 🆕
- 自动加载 WBS.md 工作分解结构
- 提供项目目标、Epic、Feature、Milestone 层级视图
- 支持任务粒度控制（300-800行代码）

### 3. OpenSpec 规范驱动 🆕
- 自动加载 OpenSpec 配置和规范
- 强制技能执行遵循具体业务规范
- 支持 /openspec:verify 和 /openspec:apply 命令

### 4. 粒度控制机制 🆕
- 智能任务粒度验证（300-800行）
- 文件数量控制（3-12个文件）
- 自动粒度建议和优化

### 5. Ralph Wiggum Loop 自动化（升级版）
- 自动执行质量检查循环
- 自我审查和违规检测
- 质量分数计算和改进建议
- WBS 更新和 OpenSpec 验证
- **智能化学习**: 基于历史数据优化策略
- **预测性质量门禁**: 根据风险动态调整
- 质量分数计算和改进建议

### 6. 零负担设计
- 技能开发者无需额外配置
- 默认启用所有质量保证机制
- 可选择性禁用特定功能

## 🏗️ 架构设计

```
技能执行框架
├── ExecutionFramework (执行框架)
│   ├── HARNESS.md 加载器
│   ├── WBS 工作分解结构加载器 🆕
│   ├── OpenSpec 配置加载器 🆕
│   ├── 约束规则解析器
│   ├── 粒度控制器 🆕
│   ├── 技能内容增强器
│   ├── Ralph Wiggum Loop 执行器
│   ├── **智能化学习** 🆕
│   └── **预测性质量门禁** 🆕
├── SkillsLoader (技能加载器)
│   ├── 技能发现和加载
│   ├── **上下文管理器** 🆕
│   ├── **专业化规范注册表** 🆕
│   ├── **学习和进化引擎** 🆕
│   ├── 执行框架集成
│   └── 上下文管理
├── **增强组件** 🆕
│   ├── ContextManager (上下文管理)
│   ├── SpecRegistry (专业化规范注册)
│   ├── LearningEngine (学习和进化)
│   ├── QualityPredictor (质量预测)
│   └── CollaborationHub (协作中心)
├── GranularityControl (粒度控制) 🆕
│   ├── 任务规模验证
│   ├── 文件数量控制
│   └── 粒度建议生成
├── WBSWorkBreakdown (工作分解) 🆕
│   ├── Epic/Feature/Milestone 层级
│   ├── 估算信息管理
│   └── 依赖关系跟踪
├── OpenSpecConfig (OpenSpec配置) 🆕
│   ├── AI工具集成
│   ├── 自定义Schema支持
│   └── 上下文规则管理
├── ExecutionResult (执行结果)
│   ├── 质量指标
│   ├── 违规记录
│   ├── **粒度验证结果** 🆕
│   ├── **WBS更新信息** 🆕
│   ├── **OpenSpec验证结果** 🆕
│   └── **学习洞察** 🆕
└── ObservabilityStack (可观测性栈) 🆕
    ├── MetricsCollector (指标收集)
    ├── DistributedTracing (分布式追踪)
    ├── PerformanceMonitor (性能监控)
    └── QualityTracker (质量跟踪)
```

## 📋 使用方法

### 基本使用

```go
package main

import (
    "context"
    "github.com/sipeed/picoclaw/pkg/skills"
)

func main() {
    // 创建技能加载器（自动启用HARNESS.md、WBS和OpenSpec集成）
    loader := skills.NewSkillsLoader(workspace, "", "")
    
    // 执行技能（自动集成质量保证和粒度控制）
    ctx := context.Background()
    result, err := loader.ExecuteSkill(ctx, "po-core", "开发电商购物车功能")
    
    if err != nil {
        // 处理错误
        return
    }
    
    // 检查执行结果
    if result.Success && result.RalphLoopPassed {
        fmt.Println("技能执行成功，质量检查通过")
        fmt.Printf("质量分数: %.1f\n", result.QualityScore)
        fmt.Printf("粒度验证: %v\n", result.Metadata["granularity_valid"])
    }
}
```

### WBS集成使用 🆕

```go
// 获取WBS信息
framework := loader.GetExecutionFramework()
wbs, err := framework.LoadWBS()
if err != nil {
    log.Printf("加载WBS失败: %v", err)
    return
}

// 查看相关里程碑
for _, milestone := range wbs.Milestones {
    if strings.Contains(milestone.Title, "电商") {
        fmt.Printf("找到相关里程碑: %s\n", milestone.Title)
        fmt.Printf("估算: %d行, %d小时\n", 
            milestone.Estimate.Lines, milestone.Estimate.Hours)
    }
}
```

### 粒度控制使用 🆕

```go
// 验证任务粒度
granularityControl := framework.GetGranularityControl()
result := granularityControl.ValidateTaskSize(estimatedLines, fileCount)

if !result.Valid {
    fmt.Printf("粒度验证失败: %s\n", result.Recommendation)
    // 根据建议调整任务范围
} else {
    fmt.Println("任务粒度合适，可以执行")
}
```

### OpenSpec集成使用 🆕

```go
// 加载OpenSpec配置
config, err := framework.LoadOpenSpecConfig()
if err != nil {
    log.Printf("加载OpenSpec配置失败: %v", err)
    return
}

// 检查规范要求
fmt.Printf("使用Schema: %s\n", config.DefaultSchema)
fmt.Printf("AI工具: %v\n", config.AITools)

// 验证代码与spec一致性
if config.ContextRules["validation_required"] == "true" {
    // 运行/openspec:verify
    verifyResult := framework.RunOpenSpecVerify("po-system-001")
    if !verifyResult.Passed {
        fmt.Printf("Spec验证失败: %s\n", verifyResult.Issues)
    }
}
```

### 配置选项

```go
// 禁用 Ralph Wiggum Loop（用于快速执行）
loader.SetRalphWiggumLoopEnabled(false)

// 获取执行框架状态
framework := loader.GetExecutionFramework()
status := framework.GetFrameworkStatus()

// 检查 HARNESS.md 是否可用
if status["harness_available"] == true {
    fmt.Println("HARNESS.md 约束已启用")
}
```

## 🔧 HARNESS.md 格式

### 基本结构

```markdown
# HARNESS.md - 项目核心约束

## 🚫 禁止清单
- ❌ 禁止全局变量
- ❌ 禁止循环依赖
- ❌ 禁止直接数据库操作

## ✅ 必须遵守
- ✅ 测试覆盖率 ≥ 85%
- ✅ 严格分层架构
- ✅ 完整错误处理

## 🔄 Ralph Wiggum Loop
1. 运行所有测试 + lint + 类型检查
2. 根据HARNESS.md进行自我审查
3. 如果失败 → 分析根因 → 修复 → 重复
4. 只有在绿色+自我审查干净时才创建PR
5. 如果架构变更，更新HARNESS.md/ADR
6. 提出2-3个改进建议

## 📊 质量标准
测试覆盖率 ≥ 85%
代码质量 ≥ 85%
安全评分 ≥ 90%
性能评分 ≥ 80%
```

### 支持的规则类型

#### 禁止清单 (Prohibitions)
- 格式：`- ❌ 禁止描述`
- 用途：定义不允许的行为和做法
- 示例：`- ❌ 禁止全局变量`

#### 必须遵守 (Requirements)
- 格式：`- ✅ 要求描述`
- 用途：定义必须遵守的标准和规范
- 示例：`- ✅ 测试覆盖率 ≥ 85%`

#### 质量标准 (Quality Standards)
- 格式：`指标名称 ≥ 数值`
- 用途：定义具体的质量指标
- 示例：`测试覆盖率 ≥ 85%`

## 🔄 Ralph Wiggum Loop 流程

### 执行步骤

1. **运行所有检查**
   ```bash
   go test ./... -v -race -coverprofile=coverage.out
   golangci-lint run ./...
   go vet ./...
   gosec ./...
   ```

2. **自我审查**
   - 根据 HARNESS.md 检查违规项
   - 评估代码质量
   - 识别改进机会

3. **根因分析和修复**
   - 分析失败原因
   - 实施修复方案
   - 重新运行检查

4. **创建 PR**
   - 确保所有检查通过
   - 自我审查干净
   - 文档完整

5. **更新文档**
   - 更新 HARNESS.md
   - 更新 ADR
   - 更新相关文档

6. **改进建议**
   - 提出 2-3 个改进建议
   - 记录经验教训
   - 制定改进计划

### 自动化集成

```go
// Ralph Wiggum Loop 自动执行
func (ef *ExecutionFramework) executeRalphWiggumLoop(
    ctx context.Context,
    skillName string,
    result *ExecutionResult,
    rules *HarnessRules,
) (*ExecutionResult, error) {
    
    // 1. 运行所有检查
    checksPassed := ef.runAllChecks(ctx, skillName)
    
    // 2. 执行自我审查
    selfReviewPassed := ef.performSelfReview(ctx, skillName, result, rules)
    
    // 3. 生成改进建议
    improvements := ef.generateImprovements(skillName, result)
    
    // 4. 计算质量分数
    result.QualityScore = ef.calculateQualityScore(result, rules)
    result.RalphLoopPassed = checksPassed && selfReviewPassed
    
    return result, nil
}
```

## 执行结果

### ExecutionResult 结构

```go
type ExecutionResult struct {
    Success        bool              `json:"success"`        // 执行是否成功
    Output         string            `json:"output"`         // 技能输出内容
    Artifacts      []string          `json:"artifacts"`      // 生成的文件列表
    Violations     []Violation       `json:"violations"`     // 违规记录
    QualityScore   float64           `json:"quality_score"`   // 质量分数 (0-100)
    RalphLoopPassed bool             `json:"ralph_loop_passed"` // Ralph Loop 是否通过
    Metadata       map[string]string `json:"metadata"`       // 执行元数据
}
```

### Violation 结构

```go
type Violation struct {
    Type        string `json:"type"`        // 违规类型: prohibition, requirement, quality_gate
    Description string `json:"description"` // 违规描述
    Severity    string `json:"severity"`    // 严重程度: critical, high, medium, low
    Suggestion  string `json:"suggestion"`  // 修复建议
}
```

### 示例输出

```json
{
    "success": true,
    "output": "技能执行完成，输出了需求分析报告",
    "artifacts": [
        "requirements-analysis.md",
        "user-stories.md",
        "market-research.xlsx"
    ],
    "violations": [
        {
            "type": "quality_gate",
            "description": "测试覆盖率不足",
            "severity": "medium",
            "suggestion": "增加单元测试覆盖率到85%以上"
        }
    ],
    "quality_score": 85.0,
    "ralph_loop_passed": true,
    "metadata": {
        "skill_name": "role-analyst",
        "execution_time": "2026-03-12T07:30:00Z",
        "enhanced_with_harness": "true",
        "improvements": "建议1: 优化代码结构;建议2: 增强测试覆盖率;建议3: 改进文档完整性"
    }
}
```

## 🛠️ 高级配置

### 自定义执行框架

```go
// 创建自定义执行框架
framework := skills.NewExecutionFramework(workspace, true)

// 加载自定义 HARNESS.md
rules, err := framework.LoadHarnessRules()
if err != nil {
    // 使用默认规则
    rules = framework.GetDefaultHarnessRules()
}

// 执行技能
result, err := framework.ExecuteSkill(ctx, skillName, skillContent, taskInput)
```

### 批量技能执行

```go
// 批量执行多个技能
skillsToExecute := []string{
    "role-analyst",
    "role-architect", 
    "role-developer",
    "role-qa",
}

for _, skillName := range skillsToExecute {
    result, err := loader.ExecuteSkill(ctx, skillName, "协作完成项目")
    if err != nil {
        log.Printf("技能 %s 执行失败: %v", skillName, err)
        continue
    }
    
    fmt.Printf("技能 %s: 成功=%v, 质量=%.1f\n", 
        skillName, result.Success, result.QualityScore)
}
```

### 错误处理

```go
result, err := loader.ExecuteSkill(ctx, "skill-name", "task input")
if err != nil {
    switch {
    case strings.Contains(err.Error(), "技能不存在"):
        log.Println("技能未找到，请检查技能名称")
    case strings.Contains(err.Error(), "HARNESS.md"):
        log.Println("HARNESS.md 解析失败，使用默认规则")
    case context.Canceled == err:
        log.Println("执行被取消")
    default:
        log.Printf("未知错误: %v", err)
    }
}
```

## 📈 性能考虑

### 执行开销

- **HARNESS.md 加载**: 一次性加载，后续执行复用
- **规则解析**: 缓存解析结果，避免重复解析
- **Ralph Loop**: 可选启用，根据需要开关

### 优化建议

1. **缓存 HARNESS.md 规则**
   ```go
   // 框架自动缓存规则，无需手动处理
   ```

2. **批量执行技能**
   ```go
   // 避免重复初始化开销
   ```

3. **选择性启用 Ralph Loop**
   ```go
   // 快速执行时禁用
   loader.SetRalphWiggumLoopEnabled(false)
   ```

## 🧪 测试

### 单元测试

```bash
# 运行执行框架测试
go test ./pkg/skills/execution_framework_test.go

# 运行技能加载器测试
go test ./pkg/skills/loader_test.go

# 运行所有技能相关测试
go test ./pkg/skills/...
```

### 集成测试

```bash
# 运行示例程序
go run examples/skill_execution_with_harness.go

# 测试特定技能
go run examples/skill_execution_with_harness.go po-core
```

### 基准测试

```bash
# 运行性能基准测试
go test -bench=. ./pkg/skills/execution_framework_test.go
```

## 🔮 扩展性

### 自定义规则解析器

```go
// 实现 CustomRuleParser 接口
type CustomRuleParser interface {
    Parse(content string) (*HarnessRules, error)
}

// 注册自定义解析器
framework.SetRuleParser(&MyCustomParser{})
```

### 自定义质量检查器

```go
// 实现 QualityChecker 接口
type QualityChecker interface {
    Check(ctx context.Context, skillName string, result *ExecutionResult) (bool, []Violation)
}

// 添加自定义检查器
framework.AddQualityChecker(&MyQualityChecker{})
```

### 自定义改进建议生成器

```go
// 实现 ImprovementGenerator 接口
type ImprovementGenerator interface {
    Generate(skillName string, result *ExecutionResult) []string
}

// 设置自定义生成器
framework.SetImprovementGenerator(&MyImprovementGenerator{})
```

## 📚 最佳实践

### 1. HARNESS.md 维护

- **定期更新**: 随着项目发展更新约束规则
- **团队共识**: 确保团队对规则达成共识
- **文档完整**: 提供清晰的规则说明和示例

### 2. 技能开发

- **专注业务逻辑**: 技能开发者只需关注业务逻辑
- **信任框架**: 依赖框架保证质量标准
- **测试充分**: 编写充分的测试用例

### 3. 执行监控

- **质量趋势**: 监控技能执行的质量趋势
- **违规分析**: 分析常见违规模式和原因
- **持续改进**: 基于数据持续改进流程

## 🤝 贡献指南

### 添加新的规则类型

1. 在 `parseHarnessRules` 中添加解析逻辑
2. 更新 `HarnessRules` 结构体
3. 添加相应的测试用例
4. 更新文档

### 扩展质量检查

1. 实现 `QualityChecker` 接口
2. 在 `executeRalphWiggumLoop` 中集成
3. 添加测试用例
4. 更新配置选项

### 性能优化

1. 分析性能瓶颈
2. 实施缓存策略
3. 优化算法复杂度
4. 添加基准测试

---

通过这个技能执行框架，PicoClaw 的技能系统获得了统一的质量保证机制，确保每个技能的执行都符合项目的最高标准，同时保持了技能开发的简洁性和高效性。
