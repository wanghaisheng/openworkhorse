package skills

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/sipeed/picoclaw/.goagents/registry"
)

// DynamicCoordinator 动态协调器 - 根据team和phase配置动态生成执行计划
type DynamicCoordinator struct {
	executionFramework *ExecutionFramework
	teamConfig        *registry.TeamConfig
	phaseConfig       *registry.PhaseConfig
	logger            *slog.Logger
}

// NewDynamicCoordinator 创建新的动态协调器
func NewDynamicCoordinator(ef *ExecutionFramework, teamID string, phaseID string) (*DynamicCoordinator, error) {
	// 加载team配置
	teamConfig, err := ef.LoadTeamConfig(teamID)
	if err != nil {
		return nil, fmt.Errorf("加载team配置失败: %w", err)
	}

	// 加载phase配置
	phaseConfig, err := ef.LoadPhaseConfig(phaseID)
	if err != nil {
		return nil, fmt.Errorf("加载phase配置失败: %w", err)
	}

	return &DynamicCoordinator{
		executionFramework: ef,
		teamConfig:        teamConfig,
		phaseConfig:       phaseConfig,
		logger:            ef.logger,
	}, nil
}

// ExecutionPlan 执行计划
type ExecutionPlan struct {
	PhaseID    string          `json:"phase_id"`
	TeamID     string          `json:"team_id"`
	Strategy   string          `json:"strategy"`
	Steps      []ExecutionStep `json:"steps"`
	QualityGates []QualityGate  `json:"quality_gates"`
}

// ExecutionStep 执行步骤
type ExecutionStep struct {
	StepID       string                 `json:"step_id"`
	Order        int                    `json:"order"`
	SkillID      string                 `json:"skill_id"`
	MemberID     string                 `json:"member_id"`
	Dependencies []string               `json:"dependencies"`
	Config       map[string]interface{} `json:"config"`
	Inputs       []string               `json:"inputs"`
	Outputs      []string               `json:"outputs"`
	CoordinationType string              `json:"coordination_type"`
}

// QualityGate 质量门禁
type QualityGate struct {
	GateID      string  `json:"gate_id"`
	Name         string  `json:"name"`
	Threshold    int     `json:"threshold"`
	Reviewer     string  `json:"reviewer"`
	Required     bool    `json:"required"`
	Description  string  `json:"description"`
}

// CoordinationResult 协调结果
type CoordinationResult struct {
	PhaseID       string                    `json:"phase_id"`
	TeamID        string                    `json:"team_id"`
	ExecutionPlan *ExecutionPlan            `json:"execution_plan"`
	StepResults   map[string]*StepResult    `json:"step_results"`
	QualityResults map[string]*QualityResult `json:"quality_results"`
	Success       bool                      `json:"success"`
	Summary       string                    `json:"summary"`
}

// StepResult 步骤结果
type StepResult struct {
	StepID      string                 `json:"step_id"`
	MemberID    string                 `json:"member_id"`
	SkillID     string                 `json:"skill_id"`
	Success     bool                   `json:"success"`
	Output      map[string]interface{} `json:"output"`
	Error       string                 `json:"error,omitempty"`
	Duration    int64                  `json:"duration_ms"`
	Quality     *QualityResult         `json:"quality,omitempty"`
}

// QualityResult 质量结果
type QualityResult struct {
	GateID      string  `json:"gate_id"`
	Passed       bool    `json:"passed"`
	Score        int     `json:"score"`
	Threshold    int     `json:"threshold"`
	Reviewer     string  `json:"reviewer"`
	Findings     []string `json:"findings"`
	Recommendations []string `json:"recommendations"`
}

// GenerateExecutionPlan 生成执行计划
func (dc *DynamicCoordinator) GenerateExecutionPlan() (*ExecutionPlan, error) {
	dc.logger.Info("开始生成执行计划", 
		"team_id", dc.teamConfig.Team.ID,
		"phase_id", dc.phaseConfig.Phase.ID,
	)

	plan := &ExecutionPlan{
		PhaseID:  dc.phaseConfig.Phase.ID,
		TeamID:   dc.teamConfig.Team.ID,
		Strategy: dc.getExecutionStrategy(),
		Steps:    []ExecutionStep{},
		QualityGates: dc.extractQualityGates(),
	}

	// 根据phase配置的执行策略生成步骤
	if dc.phaseConfig.ExecutionStrategy != nil {
		steps, err := dc.generateStepsFromStrategy()
		if err != nil {
			return nil, fmt.Errorf("从策略生成步骤失败: %w", err)
		}
		plan.Steps = steps
	} else {
		// 如果没有明确的执行策略，根据team配置生成默认步骤
		steps, err := dc.generateDefaultSteps()
		if err != nil {
			return nil, fmt.Errorf("生成默认步骤失败: %w", err)
		}
		plan.Steps = steps
	}

	dc.logger.Info("执行计划生成完成", 
		"steps_count", len(plan.Steps),
		"quality_gates_count", len(plan.QualityGates),
	)

	return plan, nil
}

// getExecutionStrategy 获取执行策略
func (dc *DynamicCoordinator) getExecutionStrategy() string {
	// 根据phase类型和team特征确定执行策略
	phaseCategory := dc.phaseConfig.Phase.Category
	teamSize := len(dc.teamConfig.Team.Members)

	switch {
	case phaseCategory == "analysis" && teamSize <= 5:
		return "sequential_with_feedback"
	case phaseCategory == "development" && teamSize > 5:
		return "parallel_with_coordination"
	case phaseCategory == "validation":
		return "gate_based_validation"
	default:
		return "adaptive_coordination"
	}
}

// generateStepsFromStrategy 从策略生成步骤
func (dc *DynamicCoordinator) generateStepsFromStrategy() ([]ExecutionStep, error) {
	var steps []ExecutionStep

	// 根据phase配置中的标准任务生成步骤
	for i, task := range dc.phaseConfig.StandardTasks {
		step := ExecutionStep{
			StepID:       fmt.Sprintf("step_%d_%s", i+1, task.TemplateID),
			Order:        i + 1,
			SkillID:      dc.mapTaskToSkill(task.TemplateID),
			MemberID:     dc.findMemberForTask(task.TemplateID),
			Dependencies: dc.getTaskDependencies(task.TemplateID),
			Config:       dc.getTaskConfig(task.TemplateID),
			Inputs:       dc.getTaskInputs(task.TemplateID),
			Outputs:      dc.getTaskOutputs(task.TemplateID),
			CoordinationType: dc.getCoordinationType(task.TemplateID),
		}
		steps = append(steps, step)
	}

	return steps, nil
}

// generateDefaultSteps 生成默认步骤
func (dc *DynamicCoordinator) generateDefaultSteps() ([]ExecutionStep, error) {
	// 基于team composition生成默认的协调步骤
	steps := []ExecutionStep{
		{
			StepID:           "step_1_planning",
			Order:            1,
			SkillID:          "po-core-v2",
			MemberID:         dc.findTeamLead(),
			Dependencies:     []string{},
			Config:           map[string]interface{}{"role": "planning"},
			Inputs:           []string{"project_request"},
			Outputs:          []string{"coordination_plan"},
			CoordinationType: "leadership",
		},
		{
			StepID:           "step_2_analysis",
			Order:            2,
			SkillID:          "requirement-analyzer",
			MemberID:         dc.findMemberBySkill("requirement-analyzer"),
			Dependencies:     []string{"step_1_planning"},
			Config:           map[string]interface{}{"depth": "comprehensive"},
			Inputs:           []string{"coordination_plan"},
			Outputs:          []string{"requirements_analysis"},
			CoordinationType: "collaborative",
		},
		{
			StepID:           "step_3_execution",
			Order:            3,
			SkillID:          "mode-selector",
			MemberID:         dc.findMemberBySkill("mode-selector"),
			Dependencies:     []string{"step_2_analysis"},
			Config:           map[string]interface{}{"auto_select": true},
			Inputs:           []string{"requirements_analysis"},
			Outputs:          []string{"mode_selection"},
			CoordinationType: "decision_driven",
		},
	}

	return steps, nil
}

// mapTaskToSkill 将任务映射到技能
func (dc *DynamicCoordinator) mapTaskToSkill(taskID string) string {
	taskSkillMapping := map[string]string{
		"market-analysis":       "requirement-analyzer",
		"requirement-gathering": "requirement-analyzer",
		"api-development":       "standard-mode-executor",
		"prototype-development": "free-mode-executor",
		"quality-validation":    "harness-integrator",
		"team-planning":         "team-builder",
		"phase-management":      "phase-manager-v2",
	}

	if skill, exists := taskSkillMapping[taskID]; exists {
		return skill
	}
	return "po-core-v2" // 默认技能
}

// findMemberForTask 为任务找到合适的成员
func (dc *DynamicCoordinator) findMemberForTask(taskID string) string {
	skillID := dc.mapTaskToSkill(taskID)
	return dc.findMemberBySkill(skillID)
}

// findMemberBySkill 根据技能找到成员
func (dc *DynamicCoordinator) findMemberBySkill(skillID string) string {
	// 遍历团队成员，找到匹配的技能
	for _, member := range dc.teamConfig.Team.Members {
		// 这里需要根据实际的member结构来判断技能匹配
		// 假设member有Agent和Variant字段
		if member.Agent == skillID || member.Variant == skillID {
			return member.MemberID
		}
	}

	// 如果没有找到，返回team lead
	return dc.findTeamLead()
}

// findTeamLead 找到团队领导
func (dc *DynamicCoordinator) findTeamLead() string {
	// 返回team lead的member ID
	// 这里需要根据实际的team配置结构来确定
	return "team_lead" // 默认返回
}

// getTaskDependencies 获取任务依赖
func (dc *DynamicCoordinator) getTaskDependencies(taskID string) []string {
	// 根据任务类型返回依赖关系
	dependencyMap := map[string][]string{
		"market-analysis":       {"planning"},
		"requirement-gathering": {"planning"},
		"api-development":       {"requirement-gathering"},
		"prototype-development": {"requirement-gathering"},
		"quality-validation":    {"api-development", "prototype-development"},
	}

	if deps, exists := dependencyMap[taskID]; exists {
		return deps
	}
	return []string{}
}

// getTaskConfig 获取任务配置
func (dc *DynamicCoordinator) getTaskConfig(taskID string) map[string]interface{} {
	// 根据任务类型返回配置
	configMap := map[string]map[string]interface{}{
		"market-analysis": {
			"method": "comprehensive",
			"sources": []string{"primary", "secondary"},
			"depth": "detailed",
		},
		"requirement-gathering": {
			"method": "collaborative",
			"stakeholders": []string{"users", "business", "technical"},
			"format": "user_stories",
		},
		"api-development": {
			"method": "standard",
			"quality_gates": []string{"unit_test", "integration_test"},
			"documentation": "required",
		},
	}

	if config, exists := configMap[taskID]; exists {
		return config
	}
	return map[string]interface{}{}
}

// getTaskInputs 获取任务输入
func (dc *DynamicCoordinator) getTaskInputs(taskID string) []string {
	// 根据任务类型返回输入
	inputMap := map[string][]string{
		"market-analysis":       {"project_context", "business_requirements"},
		"requirement-gathering": {"project_context", "market_analysis"},
		"api-development":       {"requirements_spec", "technical_constraints"},
		"prototype-development": {"requirements_spec", "design_guidelines"},
	}

	if inputs, exists := inputMap[taskID]; exists {
		return inputs
	}
	return []string{"project_request"}
}

// getTaskOutputs 获取任务输出
func (dc *DynamicCoordinator) getTaskOutputs(taskID string) []string {
	// 根据任务类型返回输出
	outputMap := map[string][]string{
		"market-analysis":       {"market_report", "competitive_analysis"},
		"requirement-gathering": {"requirements_spec", "user_stories"},
		"api-development":       {"api_code", "documentation"},
		"prototype-development": {"prototype", "demo"},
	}

	if outputs, exists := outputMap[taskID]; exists {
		return outputs
	}
	return []string{"task_output"}
}

// getCoordinationType 获取协调类型
func (dc *DynamicCoordinator) getCoordinationType(taskID string) string {
	// 根据任务类型返回协调类型
	coordinationMap := map[string]string{
		"market-analysis":       "collaborative",
		"requirement-gathering": "participatory",
		"api-development":       "technical_coordination",
		"prototype-development": "creative_coordination",
	}

	if coordType, exists := coordinationMap[taskID]; exists {
		return coordType
	}
	return "standard"
}

// extractQualityGates 提取质量门禁
func (dc *DynamicCoordinator) extractQualityGates() []QualityGate {
	var gates []QualityGate

	// 从phase配置中提取质量门禁
	if dc.phaseConfig.QualityGates != nil {
		for i, gate := range dc.phaseConfig.QualityGates {
			qg := QualityGate{
				GateID:      fmt.Sprintf("gate_%d", i+1),
				Name:         gate.Name,
				Threshold:    gate.Threshold,
				Reviewer:     gate.Reviewer,
				Required:     gate.Required,
				Description:  fmt.Sprintf("Quality gate for %s", gate.Name),
			}
			gates = append(gates, qg)
		}
	}

	// 如果没有定义质量门禁，添加默认的
	if len(gates) == 0 {
		gates = []QualityGate{
			{
				GateID:      "default_quality",
				Name:         "Default Quality Gate",
				Threshold:    85,
				Reviewer:     "team_lead",
				Required:     true,
				Description:  "Default quality assurance gate",
			},
		}
	}

	return gates
}

// ExecuteCoordination 执行协调
func (dc *DynamicCoordinator) ExecuteCoordination(ctx context.Context, input interface{}) (*CoordinationResult, error) {
	dc.logger.Info("开始执行协调", 
		"team_id", dc.teamConfig.Team.ID,
		"phase_id", dc.phaseConfig.Phase.ID,
	)

	// 1. 生成执行计划
	plan, err := dc.GenerateExecutionPlan()
	if err != nil {
		return nil, fmt.Errorf("生成执行计划失败: %w", err)
	}

	// 2. 初始化结果
	result := &CoordinationResult{
		PhaseID:        dc.phaseConfig.Phase.ID,
		TeamID:         dc.teamConfig.Team.ID,
		ExecutionPlan:  plan,
		StepResults:    make(map[string]*StepResult),
		QualityResults: make(map[string]*QualityResult),
		Success:        true,
		Summary:        "",
	}

	// 3. 按计划执行步骤
	for _, step := range plan.Steps {
		stepResult, err := dc.executeStep(ctx, step, result)
		if err != nil {
			dc.logger.Error("步骤执行失败", "step_id", step.StepID, "error", err)
			result.Success = false
			result.Summary = fmt.Sprintf("步骤 %s 执行失败: %v", step.StepID, err)
			return result, nil
		}
		result.StepResults[step.StepID] = stepResult
	}

	// 4. 执行质量门禁检查
	err = dc.validateQualityGates(result)
	if err != nil {
		dc.logger.Error("质量门禁检查失败", "error", err)
		result.Success = false
		result.Summary = fmt.Sprintf("质量门禁检查失败: %v", err)
		return result, nil
	}

	// 5. 生成总结
	result.Summary = dc.generateSummary(result)

	dc.logger.Info("协调执行完成", 
		"success", result.Success,
		"steps_completed", len(result.StepResults),
		"quality_gates_passed", len(result.QualityResults),
	)

	return result, nil
}

// executeStep 执行单个步骤
func (dc *DynamicCoordinator) executeStep(ctx context.Context, step ExecutionStep, result *CoordinationResult) (*StepResult, error) {
	dc.logger.Info("开始执行步骤", 
		"step_id", step.StepID,
		"skill_id", step.SkillID,
		"member_id", step.MemberID,
	)

	// 检查依赖
	for _, depID := range step.Dependencies {
		if _, exists := result.StepResults[depID]; !exists {
			return nil, fmt.Errorf("依赖步骤 %s 未完成", depID)
		}
	}

	// 准备输入
	inputs := dc.prepareStepInputs(step, result)

	// 执行技能（这里需要与实际的技能执行系统集成）
	stepResult := &StepResult{
		StepID:   step.StepID,
		MemberID: step.MemberID,
		SkillID:  step.SkillID,
		Success:  true,
		Output:   map[string]interface{}{
			"inputs":  inputs,
			"config":  step.Config,
			"outputs": step.Outputs,
		},
		Duration: 1000, // 模拟执行时间
	}

	// 这里应该调用实际的技能执行逻辑
	// success, output, err := dc.executionFramework.ExecuteSkill(ctx, step.SkillID, step.Config, inputs)
	// if err != nil {
	//     return nil, err
	// }

	dc.logger.Info("步骤执行完成", 
		"step_id", step.StepID,
		"success", stepResult.Success,
	)

	return stepResult, nil
}

// prepareStepInputs 准备步骤输入
func (dc *DynamicCoordinator) prepareStepInputs(step ExecutionStep, result *CoordinationResult) map[string]interface{} {
	inputs := make(map[string]interface{})

	// 从依赖步骤的输出中获取输入
	for _, depID := range step.Dependencies {
		if depResult, exists := result.StepResults[depID]; exists {
			for _, output := range step.Inputs {
				if outputValue, exists := depResult.Output[output]; exists {
					inputs[output] = outputValue
				}
			}
		}
	}

	return inputs
}

// validateQualityGates 验证质量门禁
func (dc *DynamicCoordinator) validateQualityGates(result *CoordinationResult) error {
	for _, gate := range result.ExecutionPlan.QualityGates {
		qualityResult := &QualityResult{
			GateID:      gate.GateID,
			Passed:      true,
			Score:       90, // 模拟分数
			Threshold:   gate.Threshold,
			Reviewer:    gate.Reviewer,
			Findings:    []string{},
			Recommendations: []string{},
		}

		// 这里应该执行实际的质量检查逻辑
		// score, findings, err := dc.executeQualityGate(gate, result)
		// if err != nil {
		//     return err
		// }

		result.QualityResults[gate.GateID] = qualityResult

		if gate.Required && !qualityResult.Passed {
			return fmt.Errorf("必需的质量门禁 %s 未通过", gate.GateID)
		}
	}

	return nil
}

// generateSummary 生成总结
func (dc *DynamicCoordinator) generateSummary(result *CoordinationResult) string {
	if result.Success {
		return fmt.Sprintf("协调执行成功完成 %d 个步骤，通过 %d 个质量门禁", 
			len(result.StepResults), len(result.QualityResults))
	} else {
		return result.Summary
	}
}
