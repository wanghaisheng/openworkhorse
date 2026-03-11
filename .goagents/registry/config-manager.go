package registry

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

// ConfigManager 配置管理器
type ConfigManager struct {
	loader    *ConfigLoader
	cache     map[string]interface{}
	cacheTTL  time.Duration
	basePath  string
}

// NewConfigManager 创建配置管理器
func NewConfigManager(basePath string) *ConfigManager {
	return &ConfigManager{
		loader:   NewConfigLoader(basePath),
		cache:    make(map[string]interface{}),
		cacheTTL: 10 * time.Minute,
		basePath:  basePath,
	}
}

// UpdateGlobalConfig 更新全局配置
func (cm *ConfigManager) UpdateGlobalConfig(updates map[string]interface{}) error {
	configPath := filepath.Join(cm.basePath, "config.yaml")
	
	// 加载现有配置
	config, err := cm.loader.LoadGlobalConfig()
	if err != nil {
		return fmt.Errorf("failed to load global config: %w", err)
	}
	
	// 应用更新
	if err := cm.applyConfigUpdates(config, updates); err != nil {
		return fmt.Errorf("failed to apply config updates: %w", err)
	}
	
	// 保存配置
	if err := cm.saveConfig(configPath, config); err != nil {
		return fmt.Errorf("failed to save global config: %w", err)
	}
	
	// 清除缓存
	cm.loader.ClearCache()
	
	return nil
}

// CreateWorkflow 创建新工作流
func (cm *ConfigManager) CreateWorkflow(workflow *WorkflowConfig) error {
	workflowPath := filepath.Join(cm.basePath, "workflows", workflow.ID+".yaml")
	
	// 验证工作流配置
	if err := cm.validateWorkflow(workflow); err != nil {
		return fmt.Errorf("workflow validation failed: %w", err)
	}
	
	// 保存工作流
	if err := cm.saveConfig(workflowPath, workflow); err != nil {
		return fmt.Errorf("failed to save workflow: %w", err)
	}
	
	// 清除缓存
	cm.loader.ClearCache()
	
	return nil
}

// UpdateWorkflow 更新工作流
func (cm *ConfigManager) UpdateWorkflow(workflowID string, updates map[string]interface{}) error {
	// 加载现有工作流
	workflow, err := cm.loader.LoadWorkflow(workflowID)
	if err != nil {
		return fmt.Errorf("failed to load workflow %s: %w", workflowID, err)
	}
	
	// 应用更新
	if err := cm.applyConfigUpdates(workflow, updates); err != nil {
		return fmt.Errorf("failed to apply workflow updates: %w", err)
	}
	
	// 保存工作流
	workflowPath := filepath.Join(cm.basePath, "workflows", workflowID+".yaml")
	if err := cm.saveConfig(workflowPath, workflow); err != nil {
		return fmt.Errorf("failed to save workflow: %w", err)
	}
	
	// 清除缓存
	cm.loader.ClearCache()
	
	return nil
}

// CreatePhase 创建新阶段
func (cm *ConfigManager) CreatePhase(phase *PhaseConfig) error {
	phasePath := filepath.Join(cm.basePath, "phases", phase.ID+".yaml")
	
	// 验证阶段配置
	if err := cm.validatePhase(phase); err != nil {
		return fmt.Errorf("phase validation failed: %w", err)
	}
	
	// 保存阶段
	if err := cm.saveConfig(phasePath, phase); err != nil {
		return fmt.Errorf("failed to save phase: %w", err)
	}
	
	// 清除缓存
	cm.loader.ClearCache()
	
	return nil
}

// UpdatePhase 更新阶段
func (cm *ConfigManager) UpdatePhase(phaseID string, updates map[string]interface{}) error {
	// 加载现有阶段
	phase, err := cm.loader.LoadPhase(phaseID)
	if err != nil {
		return fmt.Errorf("failed to load phase %s: %w", phaseID, err)
	}
	
	// 应用更新
	if err := cm.applyConfigUpdates(phase, updates); err != nil {
		return fmt.Errorf("failed to apply phase updates: %w", err)
	}
	
	// 保存阶段
	phasePath := filepath.Join(cm.basePath, "phases", phaseID+".yaml")
	if err := cm.saveConfig(phasePath, phase); err != nil {
		return fmt.Errorf("failed to save phase: %w", err)
	}
	
	// 清除缓存
	cm.loader.ClearCache()
	
	return nil
}

// CreateTeam 创建新团队
func (cm *ConfigManager) CreateTeam(team *TeamConfig) error {
	teamPath := filepath.Join(cm.basePath, "teams", team.ID+".yaml")
	
	// 验证团队配置
	if err := cm.validateTeam(team); err != nil {
		return fmt.Errorf("team validation failed: %w", err)
	}
	
	// 保存团队
	if err := cm.saveConfig(teamPath, team); err != nil {
		return fmt.Errorf("failed to save team: %w", err)
	}
	
	// 清除缓存
	cm.loader.ClearCache()
	
	return nil
}

// UpdateTeam 更新团队
func (cm *ConfigManager) UpdateTeam(teamID string, updates map[string]interface{}) error {
	// 加载现有团队
	team, err := cm.loader.LoadTeam(teamID)
	if err != nil {
		return fmt.Errorf("failed to load team %s: %w", teamID, err)
	}
	
	// 应用更新
	if err := cm.applyConfigUpdates(team, updates); err != nil {
		return fmt.Errorf("failed to apply team updates: %w", err)
	}
	
	// 保存团队
	teamPath := filepath.Join(cm.basePath, "teams", teamID+".yaml")
	if err := cm.saveConfig(teamPath, team); err != nil {
		return fmt.Errorf("failed to save team: %w", err)
	}
	
	// 清除缓存
	cm.loader.ClearCache()
	
	return nil
}

// CreateTask 创建新任务
func (cm *ConfigManager) CreateTask(task *TaskConfig) error {
	taskPath := filepath.Join(cm.basePath, "tasks", task.ID+".yaml")
	
	// 验证任务配置
	if err := cm.validateTask(task); err != nil {
		return fmt.Errorf("task validation failed: %w", err)
	}
	
	// 保存任务
	if err := cm.saveConfig(taskPath, task); err != nil {
		return fmt.Errorf("failed to save task: %w", err)
	}
	
	// 清除缓存
	cm.loader.ClearCache()
	
	return nil
}

// UpdateTask 更新任务
func (cm *ConfigManager) UpdateTask(taskID string, updates map[string]interface{}) error {
	// 加载现有任务
	task, err := cm.loader.LoadTask(taskID)
	if err != nil {
		return fmt.Errorf("failed to load task %s: %w", taskID, err)
	}
	
	// 应用更新
	if err := cm.applyConfigUpdates(task, updates); err != nil {
		return fmt.Errorf("failed to apply task updates: %w", err)
	}
	
	// 保存任务
	taskPath := filepath.Join(cm.basePath, "tasks", taskID+".yaml")
	if err := cm.saveConfig(taskPath, task); err != nil {
		return fmt.Errorf("failed to save task: %w", err)
	}
	
	// 清除缓存
	cm.loader.ClearCache()
	
	return nil
}

// DeleteConfig 删除配置
func (cm *ConfigManager) DeleteConfig(configType, configID string) error {
	var configPath string
	
	switch configType {
	case "workflow":
		configPath = filepath.Join(cm.basePath, "workflows", configID+".yaml")
	case "phase":
		configPath = filepath.Join(cm.basePath, "phases", configID+".yaml")
	case "team":
		configPath = filepath.Join(cm.basePath, "teams", configID+".yaml")
	case "task":
		configPath = filepath.Join(cm.basePath, "tasks", configID+".yaml")
	default:
		return fmt.Errorf("unknown config type: %s", configType)
	}
	
	// 检查文件是否存在
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return fmt.Errorf("config file does not exist: %s", configPath)
	}
	
	// 删除文件
	if err := os.Remove(configPath); err != nil {
		return fmt.Errorf("failed to delete config file: %w", err)
	}
	
	// 清除缓存
	cm.loader.ClearCache()
	
	return nil
}

// CloneConfig 克隆配置
func (cm *ConfigManager) CloneConfig(configType, configID, newID string) error {
	var sourcePath, destPath string
	
	switch configType {
	case "workflow":
		sourcePath = filepath.Join(cm.basePath, "workflows", configID+".yaml")
		destPath = filepath.Join(cm.basePath, "workflows", newID+".yaml")
	case "phase":
		sourcePath = filepath.Join(cm.basePath, "phases", configID+".yaml")
		destPath = filepath.Join(cm.basePath, "phases", newID+".yaml")
	case "team":
		sourcePath = filepath.Join(cm.basePath, "teams", configID+".yaml")
		destPath = filepath.Join(cm.basePath, "teams", newID+".yaml")
	case "task":
		sourcePath = filepath.Join(cm.basePath, "tasks", configID+".yaml")
		destPath = filepath.Join(cm.basePath, "tasks", newID+".yaml")
	default:
		return fmt.Errorf("unknown config type: %s", configType)
	}
	
	// 检查源文件是否存在
	if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
		return fmt.Errorf("source config file does not exist: %s", sourcePath)
	}
	
	// 读取源配置
	data, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to read source config: %w", err)
	}
	
	// 修改ID（简单替换）
	// 注意：这里需要更复杂的ID替换逻辑
	newData := cm.replaceConfigID(data, configID, newID)
	
	// 保存新配置
	if err := os.WriteFile(destPath, newData, 0644); err != nil {
		return fmt.Errorf("failed to save cloned config: %w", err)
	}
	
	// 清除缓存
	cm.loader.ClearCache()
	
	return nil
}

// ValidateAllConfigs 验证所有配置
func (cm *ConfigManager) ValidateAllConfigs() error {
	// 验证全局配置
	if err := cm.loader.ValidateConfig(); err != nil {
		return fmt.Errorf("global config validation failed: %w", err)
	}
	
	// 验证工作流配置
	workflows, err := cm.loader.ListWorkflows()
	if err != nil {
		return fmt.Errorf("failed to list workflows: %w", err)
	}
	
	for _, workflowID := range workflows {
		workflow, err := cm.loader.LoadWorkflow(workflowID)
		if err != nil {
			return fmt.Errorf("failed to load workflow %s: %w", workflowID, err)
		}
		
		if err := cm.validateWorkflow(workflow); err != nil {
			return fmt.Errorf("workflow %s validation failed: %w", workflowID, err)
		}
	}
	
	// 验证阶段配置
	phases, err := cm.loader.ListPhases()
	if err != nil {
		return fmt.Errorf("failed to list phases: %w", err)
	}
	
	for _, phaseID := range phases {
		phase, err := cm.loader.LoadPhase(phaseID)
		if err != nil {
			return fmt.Errorf("failed to load phase %s: %w", phaseID, err)
		}
		
		if err := cm.validatePhase(phase); err != nil {
			return fmt.Errorf("phase %s validation failed: %w", phaseID, err)
		}
	}
	
	// 验证团队配置
	teams, err := cm.loader.ListTeams()
	if err != nil {
		return fmt.Errorf("failed to list teams: %w", err)
	}
	
	for _, teamID := range teams {
		team, err := cm.loader.LoadTeam(teamID)
		if err != nil {
			return fmt.Errorf("failed to load team %s: %w", teamID, err)
		}
		
		if err := cm.validateTeam(team); err != nil {
			return fmt.Errorf("team %s validation failed: %w", teamID, err)
		}
	}
	
	// 验证任务配置
	tasks, err := cm.loader.ListTasks()
	if err != nil {
		return fmt.Errorf("failed to list tasks: %w", err)
	}
	
	for _, taskID := range tasks {
		task, err := cm.loader.LoadTask(taskID)
		if err != nil {
			return fmt.Errorf("failed to load task %s: %w", taskID, err)
		}
		
		if err := cm.validateTask(task); err != nil {
			return fmt.Errorf("task %s validation failed: %w", taskID, err)
		}
	}
	
	return nil
}

// 辅助方法

func (cm *ConfigManager) applyConfigUpdates(config interface{}, updates map[string]interface{}) error {
	// 使用反射或其他方法应用更新
	// 这里简化实现，实际需要更复杂的逻辑
	return nil
}

func (cm *ConfigManager) saveConfig(path string, config interface{}) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	
	// 确保目录存在
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	
	return nil
}

func (cm *ConfigManager) replaceConfigID(data []byte, oldID, newID string) []byte {
	// 简单的字符串替换，实际需要更复杂的解析和替换
	// 这里只是示例实现
	return data
}

func (cm *ConfigManager) validateWorkflow(workflow *WorkflowConfig) error {
	// 验证工作流配置的完整性
	if workflow.ID == "" {
		return fmt.Errorf("workflow ID is required")
	}
	if workflow.Name == "" {
		return fmt.Errorf("workflow name is required")
	}
	if len(workflow.Phases) == 0 {
		return fmt.Errorf("workflow must have at least one phase")
	}
	return nil
}

func (cm *ConfigManager) validatePhase(phase *PhaseConfig) error {
	// 验证阶段配置的完整性
	if phase.ID == "" {
		return fmt.Errorf("phase ID is required")
	}
	if phase.Name == "" {
		return fmt.Errorf("phase name is required")
	}
	return nil
}

func (cm *ConfigManager) validateTeam(team *TeamConfig) error {
	// 验证团队配置的完整性
	if team.ID == "" {
		return fmt.Errorf("team ID is required")
	}
	if team.Name == "" {
		return fmt.Errorf("team name is required")
	}
	if team.TeamLead.Agent == "" {
		return fmt.Errorf("team must have a team lead")
	}
	return nil
}

func (cm *ConfigManager) validateTask(task *TaskConfig) error {
	// 验证任务配置的完整性
	if task.ID == "" {
		return fmt.Errorf("task ID is required")
	}
	if task.Name == "" {
		return fmt.Errorf("task name is required")
	}
	if len(task.ExecutionSteps) == 0 {
		return fmt.Errorf("task must have at least one execution step")
	}
	return nil
}
