package registry

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

// ConfigLoader 配置加载器
type ConfigLoader struct {
	configPath    string
	cache        map[string]interface{}
	cacheMutex   sync.RWMutex
	cacheTTL     time.Duration
	lastLoadTime time.Time
}

// NewConfigLoader 创建新的配置加载器
func NewConfigLoader(configPath string) *ConfigLoader {
	return &ConfigLoader{
		configPath: configPath,
		cache:      make(map[string]interface{}),
		cacheTTL:   5 * time.Minute, // 5分钟缓存
	}
}

// LoadGlobalConfig 加载全局配置
func (cl *ConfigLoader) LoadGlobalConfig() (*GlobalConfig, error) {
	configPath := filepath.Join(cl.configPath, "config.yaml")
	
	cl.cacheMutex.RLock()
	if cached, exists := cl.cache["global"]; exists && time.Since(cl.lastLoadTime) < cl.cacheTTL {
		cl.cacheMutex.RUnlock()
		return cached.(*GlobalConfig), nil
	}
	cl.cacheMutex.RUnlock()

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read global config: %w", err)
	}

	var config GlobalConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse global config: %w", err)
	}

	cl.cacheMutex.Lock()
	cl.cache["global"] = &config
	cl.lastLoadTime = time.Now()
	cl.cacheMutex.Unlock()

	return &config, nil
}

// LoadWorkflow 加载工作流配置
func (cl *ConfigLoader) LoadWorkflow(workflowID string) (*WorkflowConfig, error) {
	workflowPath := filepath.Join(cl.configPath, "workflows", workflowID+".yaml")
	
	cl.cacheMutex.RLock()
	cacheKey := "workflow:" + workflowID
	if cached, exists := cl.cache[cacheKey]; exists && time.Since(cl.lastLoadTime) < cl.cacheTTL {
		cl.cacheMutex.RUnlock()
		return cached.(*WorkflowConfig), nil
	}
	cl.cacheMutex.RUnlock()

	data, err := os.ReadFile(workflowPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read workflow config %s: %w", workflowID, err)
	}

	var workflow WorkflowConfig
	if err := yaml.Unmarshal(data, &workflow); err != nil {
		return nil, fmt.Errorf("failed to parse workflow config %s: %w", workflowID, err)
	}

	cl.cacheMutex.Lock()
	cl.cache[cacheKey] = &workflow
	cl.lastLoadTime = time.Now()
	cl.cacheMutex.Unlock()

	return &workflow, nil
}

// LoadPhase 加载阶段配置
func (cl *ConfigLoader) LoadPhase(phaseID string) (*PhaseConfig, error) {
	phasePath := filepath.Join(cl.configPath, "phases", phaseID+".yaml")
	
	cl.cacheMutex.RLock()
	cacheKey := "phase:" + phaseID
	if cached, exists := cl.cache[cacheKey]; exists && time.Since(cl.lastLoadTime) < cl.cacheTTL {
		cl.cacheMutex.RUnlock()
		return cached.(*PhaseConfig), nil
	}
	cl.cacheMutex.RUnlock()

	data, err := os.ReadFile(phasePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read phase config %s: %w", phaseID, err)
	}

	var phase PhaseConfig
	if err := yaml.Unmarshal(data, &phase); err != nil {
		return nil, fmt.Errorf("failed to parse phase config %s: %w", phaseID, err)
	}

	cl.cacheMutex.Lock()
	cl.cache[cacheKey] = &phase
	cl.lastLoadTime = time.Now()
	cl.cacheMutex.Unlock()

	return &phase, nil
}

// LoadTeam 加载团队配置
func (cl *ConfigLoader) LoadTeam(teamID string) (*TeamConfig, error) {
	teamPath := filepath.Join(cl.configPath, "teams", teamID+".yaml")
	
	cl.cacheMutex.RLock()
	cacheKey := "team:" + teamID
	if cached, exists := cl.cache[cacheKey]; exists && time.Since(cl.lastLoadTime) < cl.cacheTTL {
		cl.cacheMutex.RUnlock()
		return cached.(*TeamConfig), nil
	}
	cl.cacheMutex.RUnlock()

	data, err := os.ReadFile(teamPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read team config %s: %w", teamID, err)
	}

	var team TeamConfig
	if err := yaml.Unmarshal(data, &team); err != nil {
		return nil, fmt.Errorf("failed to parse team config %s: %w", teamID, err)
	}

	cl.cacheMutex.Lock()
	cl.cache[cacheKey] = &team
	cl.lastLoadTime = time.Now()
	cl.cacheMutex.Unlock()

	return &team, nil
}

// LoadTask 加载任务配置
func (cl *ConfigLoader) LoadTask(taskID string) (*TaskConfig, error) {
	taskPath := filepath.Join(cl.configPath, "tasks", taskID+".yaml")
	
	cl.cacheMutex.RLock()
	cacheKey := "task:" + taskID
	if cached, exists := cl.cache[cacheKey]; exists && time.Since(cl.lastLoadTime) < cl.cacheTTL {
		cl.cacheMutex.RUnlock()
		return cached.(*TaskConfig), nil
	}
	cl.cacheMutex.RUnlock()

	data, err := os.ReadFile(taskPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read task config %s: %w", taskID, err)
	}

	var task TaskConfig
	if err := yaml.Unmarshal(data, &task); err != nil {
		return nil, fmt.Errorf("failed to parse task config %s: %w", taskID, err)
	}

	cl.cacheMutex.Lock()
	cl.cache[cacheKey] = &task
	cl.lastLoadTime = time.Now()
	cl.cacheMutex.Unlock()

	return &task, nil
}

// ListWorkflows 列出所有可用的工作流
func (cl *ConfigLoader) ListWorkflows() ([]string, error) {
	workflowPath := filepath.Join(cl.configPath, "workflows")
	
	entries, err := os.ReadDir(workflowPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read workflows directory: %w", err)
	}

	var workflows []string
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".yaml" {
			workflowID := entry.Name()[:len(entry.Name())-5] // 移除 .yaml 扩展名
			workflows = append(workflows, workflowID)
		}
	}

	return workflows, nil
}

// ListPhases 列出所有可用的阶段
func (cl *ConfigLoader) ListPhases() ([]string, error) {
	phasePath := filepath.Join(cl.configPath, "phases")
	
	entries, err := os.ReadDir(phasePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read phases directory: %w", err)
	}

	var phases []string
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".yaml" {
			phaseID := entry.Name()[:len(entry.Name())-5]
			phases = append(phases, phaseID)
		}
	}

	return phases, nil
}

// ListTeams 列出所有可用的团队
func (cl *ConfigLoader) ListTeams() ([]string, error) {
	teamPath := filepath.Join(cl.configPath, "teams")
	
	entries, err := os.ReadDir(teamPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read teams directory: %w", err)
	}

	var teams []string
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".yaml" {
			teamID := entry.Name()[:len(entry.Name())-5]
			teams = append(teams, teamID)
		}
	}

	return teams, nil
}

// ListTasks 列出所有可用的任务
func (cl *ConfigLoader) ListTasks() ([]string, error) {
	taskPath := filepath.Join(cl.configPath, "tasks")
	
	entries, err := os.ReadDir(taskPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read tasks directory: %w", err)
	}

	var tasks []string
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".yaml" {
			taskID := entry.Name()[:len(entry.Name())-5]
			tasks = append(tasks, taskID)
		}
	}

	return tasks, nil
}

// ClearCache 清除缓存
func (cl *ConfigLoader) ClearCache() {
	cl.cacheMutex.Lock()
	cl.cache = make(map[string]interface{})
	cl.lastLoadTime = time.Time{}
	cl.cacheMutex.Unlock()
}

// ValidateConfig 验证配置完整性
func (cl *ConfigLoader) ValidateConfig() error {
	// 检查必需的目录结构
	requiredDirs := []string{"workflows", "phases", "teams", "tasks"}
	for _, dir := range requiredDirs {
		dirPath := filepath.Join(cl.configPath, dir)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			return fmt.Errorf("required directory missing: %s", dir)
		}
	}

	// 验证全局配置
	_, err := cl.LoadGlobalConfig()
	if err != nil {
		return fmt.Errorf("global config validation failed: %w", err)
	}

	return nil
}
