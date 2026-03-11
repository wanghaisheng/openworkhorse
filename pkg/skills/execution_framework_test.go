package skills

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExecutionFramework_ExecuteSkill(t *testing.T) {
	// 创建临时目录
	tmpDir := t.TempDir()
	
	// 创建HARNESS.md文件
	harnessContent := `# HARNESS.md - 项目核心约束

## 🚫 禁止清单
- ❌ 禁止全局变量
- ❌ 禁止循环依赖
- ❌ 禁止直接数据库操作

## ✅ 必须遵守
- ✅ 测试覆盖率 ≥ 85%
- ✅ 严格分层架构
- ✅ 完整错误处理

## 📊 质量标准
测试覆盖率 ≥ 85%
代码质量 ≥ 85%
安全评分 ≥ 90%
性能评分 ≥ 80%
`
	
	harnessPath := filepath.Join(tmpDir, "HARNESS.md")
	err := os.WriteFile(harnessPath, []byte(harnessContent), 0644)
	require.NoError(t, err)

	// 创建执行框架
	framework := NewExecutionFramework(tmpDir, true)
	
	// 测试技能内容
	skillContent := `# Test Skill

这是一个测试技能。

## 功能
- 功能1
- 功能2

## 使用方法
1. 步骤1
2. 步骤2
`

	// 执行技能
	ctx := context.Background()
	result, err := framework.ExecuteSkill(ctx, "test-skill", skillContent, "测试输入")
	
	// 验证结果
	require.NoError(t, err)
	assert.True(t, result.Success)
	assert.NotEmpty(t, result.Output)
	assert.True(t, result.RalphLoopPassed)
	assert.Greater(t, result.QualityScore, 0.0)
	
	// 验证元数据
	assert.Equal(t, "test-skill", result.Metadata["skill_name"])
	assert.Equal(t, "true", result.Metadata["enhanced_with_harness"])
}

func TestExecutionFramework_LoadHarnessRules(t *testing.T) {
	// 创建临时目录
	tmpDir := t.TempDir()
	
	// 创建HARNESS.md文件
	harnessContent := `# HARNESS.md

## 🚫 禁止清单
- ❌ 禁止全局变量
- ❌ 禁止循环依赖

## ✅ 必须遵守
- ✅ 测试覆盖率 ≥ 85%
- ✅ 严格分层架构

## 📊 质量标准
测试覆盖率 ≥ 90%
代码质量 ≥ 80%
`
	
	harnessPath := filepath.Join(tmpDir, "HARNESS.md")
	err := os.WriteFile(harnessPath, []byte(harnessContent), 0644)
	require.NoError(t, err)

	// 创建执行框架
	framework := NewExecutionFramework(tmpDir, true)
	
	// 加载规则
	rules, err := framework.loadHarnessRules()
	require.NoError(t, err)
	
	// 验证规则
	assert.Equal(t, 2, len(rules.Prohibitions))
	assert.Equal(t, 2, len(rules.Requirements))
	assert.Equal(t, 90, rules.QualityGate.TestCoverage)
	assert.Equal(t, 80.0, rules.QualityGate.CodeQuality)
}

func TestExecutionFramework_GetDefaultHarnessRules(t *testing.T) {
	// 创建执行框架（不创建HARNESS.md文件）
	tmpDir := t.TempDir()
	framework := NewExecutionFramework(tmpDir, true)
	
	// 获取默认规则
	rules := framework.getDefaultHarnessRules()
	
	// 验证默认规则
	assert.NotEmpty(t, rules.Prohibitions)
	assert.NotEmpty(t, rules.Requirements)
	assert.Equal(t, 85, rules.QualityGate.TestCoverage)
	assert.Equal(t, 85.0, rules.QualityGate.CodeQuality)
}

func TestExecutionFramework_BuildEnhancedSkillContent(t *testing.T) {
	// 创建临时目录
	tmpDir := t.TempDir()
	
	// 创建HARNESS.md文件
	harnessContent := `# HARNESS.md

## 🚫 禁止清单
- ❌ 禁止全局变量

## ✅ 必须遵守
- ✅ 测试覆盖率 ≥ 85%
`
	
	harnessPath := filepath.Join(tmpDir, "HARNESS.md")
	err := os.WriteFile(harnessPath, []byte(harnessContent), 0644)
	require.NoError(t, err)

	// 创建执行框架
	framework := NewExecutionFramework(tmpDir, true)
	
	// 获取规则
	rules, err := framework.loadHarnessRules()
	require.NoError(t, err)
	
	// 构建增强内容
	skillContent := `# Original Skill

原始技能内容
`
	
	enhanced := framework.buildEnhancedSkillContent("test-skill", skillContent, rules)
	
	// 验证增强内容
	assert.Contains(t, enhanced, "🛡️ HARNESS.md 集成约束")
	assert.Contains(t, enhanced, "🚫 禁止清单")
	assert.Contains(t, enhanced, "✅ 必须遵守")
	assert.Contains(t, enhanced, "🔄 Ralph Wiggum Loop")
	assert.Contains(t, enhanced, "📋 原始技能内容")
	assert.Contains(t, enhanced, "原始技能内容")
}

func TestExecutionFramework_ExecuteRalphWiggumLoop(t *testing.T) {
	// 创建临时目录
	tmpDir := t.TempDir()
	
	// 创建HARNESS.md文件
	harnessContent := `# HARNESS.md

## 🚫 禁止清单
- ❌ 禁止全局变量

## ✅ 必须遵守
- ✅ 测试覆盖率 ≥ 85%
`
	
	harnessPath := filepath.Join(tmpDir, "HARNESS.md")
	err := os.WriteFile(harnessPath, []byte(harnessContent), 0644)
	require.NoError(t, err)

	// 创建执行框架
	framework := NewExecutionFramework(tmpDir, true)
	
	// 获取规则
	rules, err := framework.loadHarnessRules()
	require.NoError(t, err)
	
	// 创建初始结果
	result := &ExecutionResult{
		Success:      true,
		Output:       "测试输出",
		Artifacts:    []string{"test.txt"},
		Violations:   []Violation{},
		QualityScore: 85.0,
	}
	
	// 执行Ralph Wiggum Loop
	ctx := context.Background()
	finalResult, err := framework.executeRalphWiggumLoop(ctx, "test-skill", result, rules)
	
	// 验证结果
	require.NoError(t, err)
	assert.True(t, finalResult.RalphLoopPassed)
	assert.Contains(t, finalResult.Metadata, "improvements")
}

func TestExecutionFramework_GetFrameworkStatus(t *testing.T) {
	// 创建临时目录
	tmpDir := t.TempDir()
	
	// 创建HARNESS.md文件
	harnessContent := `# HARNESS.md
测试内容
`
	
	harnessPath := filepath.Join(tmpDir, "HARNESS.md")
	err := os.WriteFile(harnessPath, []byte(harnessContent), 0644)
	require.NoError(t, err)

	// 创建执行框架
	framework := NewExecutionFramework(tmpDir, true)
	
	// 获取状态
	status := framework.GetFrameworkStatus()
	
	// 验证状态
	assert.Equal(t, tmpDir, status["workspace"])
	assert.Equal(t, harnessPath, status["harness_path"])
	assert.Equal(t, true, status["ralph_loop_enabled"])
	assert.Equal(t, true, status["harness_available"])
}

func TestSkillsLoader_ExecuteSkill(t *testing.T) {
	// 创建临时目录
	tmpDir := t.TempDir()
	
	// 创建技能目录
	skillsDir := filepath.Join(tmpDir, "skills")
	err := os.MkdirAll(skillsDir, 0755)
	require.NoError(t, err)
	
	// 创建技能文件
	skillDir := filepath.Join(skillsDir, "test-skill")
	err = os.MkdirAll(skillDir, 0755)
	require.NoError(t, err)
	
	skillContent := `# Test Skill

测试技能内容
`
	
	skillPath := filepath.Join(skillDir, "SKILL.md")
	err = os.WriteFile(skillPath, []byte(skillContent), 0644)
	require.NoError(t, err)
	
	// 创建HARNESS.md文件
	harnessContent := `# HARNESS.md

## 🚫 禁止清单
- ❌ 禁止全局变量

## ✅ 必须遵守
- ✅ 测试覆盖率 ≥ 85%
`
	
	harnessPath := filepath.Join(tmpDir, "HARNESS.md")
	err = os.WriteFile(harnessPath, []byte(harnessContent), 0644)
	require.NoError(t, err)

	// 创建技能加载器
	loader := NewSkillsLoader(tmpDir, "", "")
	
	// 执行技能
	ctx := context.Background()
	result, err := loader.ExecuteSkill(ctx, "test-skill", "测试输入")
	
	// 验证结果
	require.NoError(t, err)
	assert.True(t, result.Success)
	assert.NotEmpty(t, result.Output)
}

func TestExecutionFramework_ContextCancellation(t *testing.T) {
	// 创建临时目录
	tmpDir := t.TempDir()
	
	// 创建HARNESS.md文件
	harnessContent := `# HARNESS.md
测试内容
`
	
	harnessPath := filepath.Join(tmpDir, "HARNESS.md")
	err := os.WriteFile(harnessPath, []byte(harnessContent), 0644)
	require.NoError(t, err)

	// 创建执行框架
	framework := NewExecutionFramework(tmpDir, true)
	
	// 创建可取消的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	
	// 等待上下文取消
	time.Sleep(20 * time.Millisecond)
	
	// 执行技能（应该因为上下文取消而失败）
	skillContent := `# Test Skill
测试内容
`
	
	_, err = framework.ExecuteSkill(ctx, "test-skill", skillContent, "测试输入")
	assert.Error(t, err)
	assert.Equal(t, context.Canceled, err)
}

func BenchmarkExecutionFramework_ExecuteSkill(b *testing.B) {
	// 创建临时目录
	tmpDir := b.TempDir()
	
	// 创建HARNESS.md文件
	harnessContent := `# HARNESS.md

## 🚫 禁止清单
- ❌ 禁止全局变量

## ✅ 必须遵守
- ✅ 测试覆盖率 ≥ 85%
`
	
	harnessPath := filepath.Join(tmpDir, "HARNESS.md")
	err := os.WriteFile(harnessPath, []byte(harnessContent), 0644)
	if err != nil {
		b.Fatal(err)
	}

	// 创建执行框架
	framework := NewExecutionFramework(tmpDir, true)
	
	// 准备技能内容
	skillContent := `# Test Skill

这是一个测试技能。

## 功能
- 功能1
- 功能2

## 使用方法
1. 步骤1
2. 步骤2
`

	// 重置计时器
	b.ResetTimer()
	
	// 运行基准测试
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		_, err := framework.ExecuteSkill(ctx, "test-skill", skillContent, "测试输入")
		if err != nil {
			b.Fatal(err)
		}
	}
}
