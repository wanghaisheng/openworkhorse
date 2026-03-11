package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/sipeed/picoclaw/pkg/skills"
)

func main() {
	// 示例：演示如何使用集成HARNESS.md和Ralph Wiggum Loop的技能执行框架

	// 1. 设置工作目录
	workspace, err := os.Getwd()
	if err != nil {
		log.Fatal("获取工作目录失败:", err)
	}

	// 2. 创建技能加载器（自动启用HARNESS.md和Ralph Wiggum Loop）
	loader := skills.NewSkillsLoader(workspace, "", "")

	// 3. 检查执行框架状态
	framework := loader.GetExecutionFramework()
	status := framework.GetFrameworkStatus()
	
	fmt.Println("=== 执行框架状态 ===")
	fmt.Printf("工作目录: %s\n", status["workspace"])
	fmt.Printf("HARNESS.md路径: %s\n", status["harness_path"])
	fmt.Printf("Ralph Wiggum Loop启用: %v\n", status["ralph_loop_enabled"])
	fmt.Printf("HARNESS.md可用: %v\n", status["harness_available"])
	fmt.Println()

	// 4. 演示执行PO Core技能
	fmt.Println("=== 执行PO Core技能 ===")
	ctx := context.Background()
	
	result, err := loader.ExecuteSkill(ctx, "po-core", "开发电商购物车功能")
	if err != nil {
		log.Printf("技能执行失败: %v", err)
		return
	}

	fmt.Printf("技能执行成功: %v\n", result.Success)
	fmt.Printf("输出内容: %s\n", result.Output)
	fmt.Printf("质量分数: %.1f\n", result.QualityScore)
	fmt.Printf("Ralph Wiggum Loop通过: %v\n", result.RalphLoopPassed)
	
	if len(result.Violations) > 0 {
		fmt.Println("发现的违规项:")
		for _, violation := range result.Violations {
			fmt.Printf("  - [%s] %s (%s): %s\n", 
				violation.Type, violation.Description, violation.Severity, violation.Suggestion)
		}
	}
	
	fmt.Println("生成的文件:")
	for _, artifact := range result.Artifacts {
		fmt.Printf("  - %s\n", artifact)
	}
	
	fmt.Println("元数据:")
	for key, value := range result.Metadata {
		fmt.Printf("  %s: %s\n", key, value)
	}
	fmt.Println()

	// 5. 演示执行团队角色技能
	fmt.Println("=== 执行Analyst角色技能 ===")
	
	analystResult, err := loader.ExecuteSkill(ctx, "role-analyst", "分析电商购物车功能需求")
	if err != nil {
		log.Printf("Analyst技能执行失败: %v", err)
		return
	}

	fmt.Printf("Analyst技能执行成功: %v\n", analystResult.Success)
	fmt.Printf("输出内容: %s\n", analystResult.Output)
	fmt.Printf("质量分数: %.1f\n", analystResult.QualityScore)
	fmt.Println()

	// 6. 演示禁用Ralph Wiggum Loop
	fmt.Println("=== 禁用Ralph Wiggum Loop ===")
	loader.SetRalphWiggumLoopEnabled(false)
	
	quickResult, err := loader.ExecuteSkill(ctx, "role-analyst", "快速需求分析")
	if err != nil {
		log.Printf("快速技能执行失败: %v", err)
		return
	}

	fmt.Printf("快速执行成功: %v\n", quickResult.Success)
	fmt.Printf("Ralph Wiggum Loop通过: %v\n", quickResult.RalphLoopPassed)
	fmt.Println()

	// 7. 演示批量执行技能
	fmt.Println("=== 批量执行技能 ===")
	
	skillsToExecute := []string{
		"role-analyst",
		"role-architect", 
		"role-developer",
		"role-qa",
	}

	for _, skillName := range skillsToExecute {
		fmt.Printf("执行技能: %s\n", skillName)
		
		result, err := loader.ExecuteSkill(ctx, skillName, "协作完成电商项目")
		if err != nil {
			log.Printf("技能 %s 执行失败: %v", skillName, err)
			continue
		}
		
		fmt.Printf("  成功: %v, 质量分数: %.1f\n", result.Success, result.QualityScore)
		time.Sleep(100 * time.Millisecond) // 避免过快执行
	}
	
	fmt.Println()

	// 8. 演示错误处理
	fmt.Println("=== 错误处理演示 ===")
	
	_, err = loader.ExecuteSkill(ctx, "non-existent-skill", "测试不存在的技能")
	if err != nil {
		fmt.Printf("预期的错误: %v\n", err)
	}
	
	fmt.Println()

	// 9. 演示上下文取消
	fmt.Println("=== 上下文取消演示 ===")
	
	cancelCtx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()
	
	// 模拟长时间运行的技能
	go func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("技能执行被取消")
	}()
	
	_, err = loader.ExecuteSkill(cancelCtx, "po-core", "长时间运行的任务")
	if err != nil {
		fmt.Printf("上下文取消错误: %v\n", err)
	}
	
	fmt.Println()
	fmt.Println("=== 演示完成 ===")
}

// createTestHarnessFile 创建测试用的HARNESS.md文件
func createTestHarnessFile(workspace string) error {
	harnessContent := `# HARNESS.md - PicoClaw 项目核心约束

## 🚫 禁止清单
- ❌ 禁止全局变量: 避免全局状态，使用依赖注入
- ❌ 禁止循环依赖: 模块间不允许循环引用
- ❌ 禁止直接数据库操作: 必须通过数据访问层
- ❌ 禁止硬编码配置: 使用配置文件或环境变量
- ❌ 禁止忽略错误处理: 所有错误必须被妥善处理

## ✅ 必须遵守
- ✅ 测试覆盖率 ≥ 85%: 所有核心功能必须有测试覆盖
- ✅ 严格分层架构: 表现层、业务层、数据层分离
- ✅ 完整错误处理: 所有可能的错误场景都要处理
- ✅ 代码审查通过: 所有代码必须经过同行审查
- ✅ 文档完整性: 公共API必须有完整文档

## 🔄 Ralph Wiggum Loop

完成任务后，必须执行以下步骤：
1. 运行所有测试 + lint + 类型检查直到100%通过
2. 根据HARNESS.md进行自我审查（列出违规项）
3. 如果失败 → 分析根因 → 修复 → 重复
4. 只有在绿色+自我审查干净时才创建PR
5. 如果架构变更，更新HARNESS.md/ADR
6. 提出2-3个改进建议

## 📊 质量标准
测试覆盖率 ≥ 85%
代码质量 ≥ 85%
安全评分 ≥ 90%
性能评分 ≥ 80%
`

	harnessPath := filepath.Join(workspace, "HARNESS.md")
	return os.WriteFile(harnessPath, []byte(harnessContent), 0644)
}

// createTestSkills 创建测试用的技能文件
func createTestSkills(workspace string) error {
	skillsDir := filepath.Join(workspace, "skills")
	
	// 创建技能目录
	skillNames := []string{"po-core", "role-analyst", "role-architect", "role-developer", "role-qa"}
	
	for _, skillName := range skillNames {
		skillDir := filepath.Join(skillsDir, skillName)
		err := os.MkdirAll(skillDir, 0755)
		if err != nil {
			return err
		}
		
		skillContent := fmt.Sprintf(`# %s Skill

这是 %s 技能的测试内容。

## 功能描述
- 功能1: 基础功能
- 功能2: 高级功能
- 功能3: 扩展功能

## 使用方法
1. 步骤1: 准备工作
2. 步骤2: 执行任务
3. 步骤3: 验证结果

## 输出示例
技能执行完成，输出结果如下：
- 任务1: 完成
- 任务2: 完成
- 任务3: 完成

## 质量保证
- 代码质量: 优秀
- 测试覆盖: 90%
- 性能表现: 良好
- 安全合规: 通过
`, skillName, skillName)
		
		skillPath := filepath.Join(skillDir, "SKILL.md")
		err = os.WriteFile(skillPath, []byte(skillContent), 0644)
		if err != nil {
			return err
		}
	}
	
	return nil
}

func init() {
	// 初始化测试环境
	workspace, err := os.Getwd()
	if err != nil {
		return
	}
	
	// 创建测试文件
	if err := createTestHarnessFile(workspace); err != nil {
		log.Printf("创建HARNESS.md失败: %v", err)
	}
	
	if err := createTestSkills(workspace); err != nil {
		log.Printf("创建测试技能失败: %v", err)
	}
}
