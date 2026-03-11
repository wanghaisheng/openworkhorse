package skills

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sipeed/picoclaw/pkg/logger"
)

// ExecutionFramework 提供技能执行的通用框架
type ExecutionFramework struct {
	workspace       string
	harnessPath     string
	enableRalphLoop bool
	logger          *slog.Logger
}

// NewExecutionFramework 创建新的执行框架
func NewExecutionFramework(workspace string, enableRalphLoop bool) *ExecutionFramework {
	return &ExecutionFramework{
		workspace:       workspace,
		harnessPath:     filepath.Join(workspace, "HARNESS.md"),
		enableRalphLoop: enableRalphLoop,
		logger:          logger.GetLogger("skills.execution"),
	}
}

// ExecuteSkill 执行技能，自动集成HARNESS.md和Ralph Wiggum Loop
func (ef *ExecutionFramework) ExecuteSkill(
	ctx context.Context,
	skillName string,
	skillContent string,
	taskInput string,
) (*ExecutionResult, error) {
	
	ef.logger.Info("开始执行技能", 
		"skill", skillName, 
		"with_harness", ef.enableRalphLoop,
		"with_ralph_loop", ef.enableRalphLoop,
	)

	// 1. 加载HARNESS.md约束
	harnessRules, err := ef.loadHarnessRules()
	if err != nil {
		ef.logger.Warn("加载HARNESS.md失败，使用默认规则", "error", err)
		harnessRules = ef.getDefaultHarnessRules()
	}

	// 2. 构建增强的技能内容
	enhancedContent := ef.buildEnhancedSkillContent(skillName, skillContent, harnessRules)

	// 3. 执行技能任务
	result, err := ef.executeSkillTask(ctx, skillName, enhancedContent, taskInput)
	if err != nil {
		return nil, fmt.Errorf("技能执行失败: %w", err)
	}

	// 4. 如果启用Ralph Wiggum Loop，执行循环验证
	if ef.enableRalphLoop {
		ralphResult, err := ef.executeRalphWiggumLoop(ctx, skillName, result, harnessRules)
		if err != nil {
			ef.logger.Warn("Ralph Wiggum Loop执行失败", "error", err)
			// 不阻断主流程，记录警告
		} else {
			result = ralphResult
		}
	}

	ef.logger.Info("技能执行完成", "skill", skillName, "success", result.Success)
	return result, nil
}

// HarnessRules HARNESS.md规则定义
type HarnessRules struct {
	Prohibitions []string `json:"prohibitions"`
	Requirements []string `json:"requirements"`
	QualityGate  QualityGate `json:"quality_gate"`
}

// QualityGate 质量门禁定义
type QualityGate struct {
	TestCoverage    int     `json:"test_coverage"`
	CodeQuality     float64 `json:"code_quality"`
	SecurityScore   int     `json:"security_score"`
	PerformanceScore float64 `json:"performance_score"`
}

// ExecutionResult 技能执行结果
type ExecutionResult struct {
	Success        bool              `json:"success"`
	Output         string            `json:"output"`
	Artifacts      []string          `json:"artifacts"`
	Violations     []Violation       `json:"violations"`
	QualityScore   float64           `json:"quality_score"`
	RalphLoopPassed bool             `json:"ralph_loop_passed"`
	Metadata       map[string]string `json:"metadata"`
}

// Violation 违规记录
type Violation struct {
	Type        string `json:"type"`        // prohibition, requirement, quality_gate
	Description string `json:"description"`
	Severity    string `json:"severity"`    // critical, high, medium, low
	Suggestion  string `json:"suggestion"`
}

// loadHarnessRules 加载HARNESS.md规则
func (ef *ExecutionFramework) loadHarnessRules() (*HarnessRules, error) {
	if _, err := os.Stat(ef.harnessPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("HARNESS.md文件不存在: %s", ef.harnessPath)
	}

	content, err := os.ReadFile(ef.harnessPath)
	if err != nil {
		return nil, fmt.Errorf("读取HARNESS.md失败: %w", err)
	}

	return ef.parseHarnessRules(string(content))
}

// parseHarnessRules 解析HARNESS.md内容
func (ef *ExecutionFramework) parseHarnessRules(content string) (*HarnessRules, error) {
	rules := &HarnessRules{
		QualityGate: QualityGate{
			TestCoverage:     85,
			CodeQuality:      85.0,
			SecurityScore:    90,
			PerformanceScore: 80.0,
		},
	}

	lines := strings.Split(content, "\n")
	var currentSection string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// 解析章节
		if strings.HasPrefix(line, "## 🚫 禁止清单") {
			currentSection = "prohibitions"
			continue
		}
		if strings.HasPrefix(line, "## ✅ 必须遵守") {
			currentSection = "requirements"
			continue
		}
		if strings.HasPrefix(line, "## 📊 质量标准") {
			currentSection = "quality"
			continue
		}

		// 解析具体规则
		switch currentSection {
		case "prohibitions":
			if strings.HasPrefix(line, "- ❌") {
				rules.Prohibitions = append(rules.Prohibitions, ef.cleanRuleLine(line))
			}
		case "requirements":
			if strings.HasPrefix(line, "- ✅") {
				rules.Requirements = append(rules.Requirements, ef.cleanRuleLine(line))
			}
		case "quality":
			ef.parseQualityRules(line, rules)
		}
	}

	return rules, nil
}

// cleanRuleLine 清理规则行
func (ef *ExecutionFramework) cleanRuleLine(line string) string {
	// 移除标记符号，保留规则描述
	line = strings.ReplaceAll(line, "- ❌", "")
	line = strings.ReplaceAll(line, "- ✅", "")
	line = strings.TrimSpace(line)
	return line
}

// parseQualityRules 解析质量规则
func (ef *ExecutionFramework) parseQualityRules(line string, rules *HarnessRules) {
	if strings.Contains(line, "测试覆盖率") {
		if strings.Contains(line, "≥") {
			// 提取数字
			if num := ef.extractNumber(line); num > 0 {
				rules.QualityGate.TestCoverage = int(num)
			}
		}
	}
	if strings.Contains(line, "代码质量") {
		if num := ef.extractNumber(line); num > 0 {
			rules.QualityGate.CodeQuality = num
		}
	}
	if strings.Contains(line, "安全评分") {
		if num := ef.extractNumber(line); num > 0 {
			rules.QualityGate.SecurityScore = int(num)
		}
	}
	if strings.Contains(line, "性能评分") {
		if num := ef.extractNumber(line); num > 0 {
			rules.QualityGate.PerformanceScore = num
		}
	}
}

// extractNumber 从字符串中提取数字
func (ef *ExecutionFramework) extractNumber(line string) float64 {
	// 简单的数字提取逻辑
	// 实际实现可能需要更复杂的解析
	var num float64
	fmt.Sscanf(line, "%f", &num)
	return num
}

// getDefaultHarnessRules 获取默认HARNESS规则
func (ef *ExecutionFramework) getDefaultHarnessRules() *HarnessRules {
	return &HarnessRules{
		Prohibitions: []string{
			"禁止全局变量",
			"禁止循环依赖",
			"禁止直接数据库操作",
			"禁止硬编码配置",
			"禁止忽略错误处理",
		},
		Requirements: []string{
			"测试覆盖率 ≥ 85%",
			"严格分层架构",
			"完整错误处理",
			"代码审查通过",
			"文档完整性",
		},
		QualityGate: QualityGate{
			TestCoverage:     85,
			CodeQuality:      85.0,
			SecurityScore:    90,
			PerformanceScore: 80.0,
		},
	}
}

// buildEnhancedSkillContent 构建增强的技能内容
func (ef *ExecutionFramework) buildEnhancedSkillContent(skillName, skillContent string, rules *HarnessRules) string {
	var builder strings.Builder

	// 1. 添加HARNESS.md约束
	builder.WriteString("# 🛡️ HARNESS.md 集成约束\n\n")
	builder.WriteString("## 🚫 禁止清单\n")
	for _, prohibition := range rules.Prohibitions {
		builder.WriteString(fmt.Sprintf("- ❌ %s\n", prohibition))
	}
	builder.WriteString("\n")

	builder.WriteString("## ✅ 必须遵守\n")
	for _, requirement := range rules.Requirements {
		builder.WriteString(fmt.Sprintf("- ✅ %s\n", requirement))
	}
	builder.WriteString("\n")

	// 2. 添加Ralph Wiggum Loop说明
	if ef.enableRalphLoop {
		builder.WriteString("## 🔄 Ralph Wiggum Loop\n\n")
		builder.WriteString("完成任务后，必须执行以下步骤：\n")
		builder.WriteString("1. 运行所有测试 + lint + 类型检查直到100%通过\n")
		builder.WriteString("2. 根据HARNESS.md进行自我审查（列出违规项）\n")
		builder.WriteString("3. 如果失败 → 分析根因 → 修复 → 重复\n")
		builder.WriteString("4. 只有在绿色+自我审查干净时才创建PR\n")
		builder.WriteString("5. 如果架构变更，更新HARNESS.md/ADR\n")
		builder.WriteString("6. 提出2-3个改进建议\n\n")
	}

	// 3. 添加原始技能内容
	builder.WriteString("## 📋 原始技能内容\n\n")
	builder.WriteString(skillContent)

	return builder.String()
}

// executeSkillTask 执行技能任务
func (ef *ExecutionFramework) executeSkillTask(
	ctx context.Context,
	skillName string,
	enhancedContent string,
	taskInput string,
) (*ExecutionResult, error) {
	
	// 这里应该调用实际的AI模型执行技能
	// 目前返回模拟结果
	result := &ExecutionResult{
		Success: true,
		Output: fmt.Sprintf("技能 %s 执行完成，输入: %s", skillName, taskInput),
		Artifacts: []string{
			fmt.Sprintf("output_%s_%d.md", skillName, time.Now().Unix()),
		},
		Violations: []Violation{},
		QualityScore: 90.0,
		RalphLoopPassed: false,
		Metadata: map[string]string{
			"skill_name": skillName,
			"execution_time": time.Now().Format(time.RFC3339),
			"enhanced_with_harness": "true",
		},
	}

	// 模拟执行时间
	select {
	case <-time.After(100 * time.Millisecond):
		return result, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// executeRalphWiggumLoop 执行Ralph Wiggum Loop
func (ef *ExecutionFramework) executeRalphWiggumLoop(
	ctx context.Context,
	skillName string,
	result *ExecutionResult,
	rules *HarnessRules,
) (*ExecutionResult, error) {
	
	ef.logger.Info("开始执行Ralph Wiggum Loop", "skill", skillName)

	// 1. 模拟运行所有检查
	checksPassed := ef.runAllChecks(ctx, skillName)
	if !checksPassed {
		result.RalphLoopPassed = false
		result.Violations = append(result.Violations, Violation{
			Type:        "quality_gate",
			Description: "自动化检查未完全通过",
			Severity:    "high",
			Suggestion:  "修复失败的检查项后重新运行",
		})
		return result, nil
	}

	// 2. 模拟自我审查
	selfReviewPassed := ef.performSelfReview(ctx, skillName, result, rules)
	if !selfReviewPassed {
		result.RalphLoopPassed = false
		result.Violations = append(result.Violations, Violation{
			Type:        "requirement",
			Description: "自我审查发现违规项",
			Severity:    "medium",
			Suggestion:  "根据HARNESS.md修复违规项",
		})
		return result, nil
	}

	// 3. 模拟改进建议
	improvements := ef.generateImprovements(skillName, result)
	result.Metadata["improvements"] = strings.Join(improvements, ";")

	// 4. 更新质量分数
	result.QualityScore = ef.calculateQualityScore(result, rules)
	result.RalphLoopPassed = true

	ef.logger.Info("Ralph Wiggum Loop执行完成", 
		"skill", skillName, 
		"passed", result.RalphLoopPassed,
		"quality_score", result.QualityScore,
	)

	return result, nil
}

// runAllChecks 运行所有检查
func (ef *ExecutionFramework) runAllChecks(ctx context.Context, skillName string) bool {
	// 模拟检查过程
	// 实际实现应该运行真实的测试、lint等
	ef.logger.Info("运行所有检查", "skill", skillName)
	
	// 模拟90%的通过率
	return true
}

// performSelfReview 执行自我审查
func (ef *ExecutionFramework) performSelfReview(
	ctx context.Context,
	skillName string,
	result *ExecutionResult,
	rules *HarnessRules,
) bool {
	
	ef.logger.Info("执行自我审查", "skill", skillName)
	
	// 模拟自我审查过程
	// 实际实现应该分析输出内容是否符合HARNESS.md规则
	return true
}

// generateImprovements 生成改进建议
func (ef *ExecutionFramework) generateImprovements(skillName string, result *ExecutionResult) []string {
	return []string{
		fmt.Sprintf("建议1: 优化%s的代码结构", skillName),
		fmt.Sprintf("建议2: 增强%s的测试覆盖率", skillName),
		fmt.Sprintf("建议3: 改进%s的文档完整性", skillName),
	}
}

// calculateQualityScore 计算质量分数
func (ef *ExecutionFramework) calculateQualityScore(result *ExecutionResult, rules *HarnessRules) float64 {
	score := 100.0
	
	// 根据违规项扣分
	for _, violation := range result.Violations {
		switch violation.Severity {
		case "critical":
			score -= 20
		case "high":
			score -= 10
		case "medium":
			score -= 5
		case "low":
			score -= 2
		}
	}
	
	// 确保分数不低于0
	if score < 0 {
		score = 0
	}
	
	return score
}

// GetFrameworkStatus 获取框架状态
func (ef *ExecutionFramework) GetFrameworkStatus() map[string]interface{} {
	status := map[string]interface{}{
		"workspace": ef.workspace,
		"harness_path": ef.harnessPath,
		"ralph_loop_enabled": ef.enableRalphLoop,
	}
	
	// 检查HARNESS.md是否存在
	if _, err := os.Stat(ef.harnessPath); err == nil {
		status["harness_available"] = true
	} else {
		status["harness_available"] = false
	}
	
	return status
}
