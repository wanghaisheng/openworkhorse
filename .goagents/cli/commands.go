package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sipeed/picoclaw/.goagents/registry"
	"github.com/spf13/cobra"
)

// CLI 命令定义
var rootCmd = &cobra.Command{
	Use:   "goagents",
	Short: "Go Agents 配置管理工具",
	Long:  `Go Agents 配置管理工具，用于管理 Workflow、Phase、Team、Task 配置`,
}

var workflowCmd = &cobra.Command{
	Use:   "workflow",
	Short: "管理工作流配置",
	Long:  `创建、更新、删除、克隆工作流配置`,
}

var phaseCmd = &cobra.Command{
	Use:   "phase",
	Short: "管理阶段配置",
	Long:  `创建、更新、删除、克隆阶段配置`,
}

var teamCmd = &cobra.Command{
	Use:   "team",
	Short: "管理团队配置",
	Long:  `创建、更新、删除、克隆团队配置`,
}

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "管理任务配置",
	Long:  `创建、更新、删除、克隆任务配置`,
}

// 初始化命令
func init() {
	rootCmd.AddCommand(workflowCmd)
	rootCmd.AddCommand(phaseCmd)
	rootCmd.AddCommand(teamCmd)
	rootCmd.AddCommand(taskCmd)
	
	// Workflow 子命令
	workflowCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "列出所有工作流",
		Run:   listWorkflows,
	})
	
	workflowCmd.AddCommand(&cobra.Command{
		Use:   "create [workflow-id]",
		Short: "创建新工作流",
		Args:  cobra.ExactArgs(1),
		Run:   func(cmd *cobra.Command, args []string) {
			createWorkflow(args[0])
		},
	})
	
	workflowCmd.AddCommand(&cobra.Command{
		Use:   "update [workflow-id]",
		Short: "更新工作流",
		Args:  cobra.ExactArgs(1),
		Run:   func(cmd *cobra.Command, args []string) {
			updateWorkflow(args[0])
		},
	})
	
	workflowCmd.AddCommand(&cobra.Command{
		Use:   "delete [workflow-id]",
		Short: "删除工作流",
		Args:  cobra.ExactArgs(1),
		Run:   func(cmd *cobra.Command, args []string) {
			deleteConfig("workflow", args[0])
		},
	})
	
	workflowCmd.AddCommand(&cobra.Command{
		Use:   "clone [source-id] [new-id]",
		Short: "克隆工作流",
		Args:  cobra.ExactArgs(2),
		Run:   func(cmd *cobra.Command, args []string) {
			cloneConfig("workflow", args[0], args[1])
		},
	})
	
	// Phase 子命令
	phaseCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "列出所有阶段",
		Run:   listPhases,
	})
	
	phaseCmd.AddCommand(&cobra.Command{
		Use:   "create [phase-id]",
		Short: "创建新阶段",
		Args:  cobra.ExactArgs(1),
		Run:   func(cmd *cobra.Command, args []string) {
			createPhase(args[0])
		},
	})
	
	phaseCmd.AddCommand(&cobra.Command{
		Use:   "update [phase-id]",
		Short: "更新阶段",
		Args:  cobra.ExactArgs(1),
		Run:   func(cmd *cobra.Command, args []string) {
			updatePhase(args[0])
		},
	})
	
	// Team 子命令
	teamCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "列出所有团队",
		Run:   listTeams,
	})
	
	teamCmd.AddCommand(&cobra.Command{
		Use:   "create [team-id]",
		Short: "创建新团队",
		Args:  cobra.ExactArgs(1),
		Run:   func(cmd *cobra.Command, args []string) {
			createTeam(args[0])
		},
	})
	
	teamCmd.AddCommand(&cobra.Command{
		Use:   "update [team-id]",
		Short: "更新团队",
		Args:  cobra.ExactArgs(1),
		Run:   func(cmd *cobra.Command, args []string) {
			updateTeam(args[0])
		},
	})
	
	// Task 子命令
	taskCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "列出所有任务",
		Run:   listTasks,
	})
	
	taskCmd.AddCommand(&cobra.Command{
		Use:   "create [task-id]",
		Short: "创建新任务",
		Args:  cobra.ExactArgs(1),
		Run:   func(cmd *cobra.Command, args []string) {
			createTask(args[0])
		},
	})
	
	taskCmd.AddCommand(&cobra.Command{
		Use:   "update [task-id]",
		Short: "更新任务",
		Args:  cobra.ExactArgs(1),
		Run:   func(cmd *cobra.Command, args []string) {
			updateTask(args[0])
		},
	})
}

// 获取配置管理器
func getConfigManager() *registry.ConfigManager {
	homeDir, _ := os.UserHomeDir()
	configPath := filepath.Join(homeDir, ".picoclaw", "workspace", ".goagents")
	return registry.NewConfigManager(configPath)
}

// Workflow 命令实现
func listWorkflows(cmd *cobra.Command, args []string) {
	cm := getConfigManager()
	workflows, err := cm.loader.ListWorkflows()
	if err != nil {
		fmt.Printf("Error listing workflows: %v\n", err)
		return
	}
	
	fmt.Println("Available Workflows:")
	for _, workflow := range workflows {
		fmt.Printf("  - %s\n", workflow)
	}
}

func createWorkflow(workflowID string) {
	fmt.Printf("Creating workflow: %s\n", workflowID)
	// 这里应该启动交互式配置创建流程
	// 或者从模板创建
}

func updateWorkflow(workflowID string) {
	fmt.Printf("Updating workflow: %s\n", workflowID)
	// 这里应该加载现有配置并启动编辑流程
}

// Phase 命令实现
func listPhases(cmd *cobra.Command, args []string) {
	cm := getConfigManager()
	phases, err := cm.loader.ListPhases()
	if err != nil {
		fmt.Printf("Error listing phases: %v\n", err)
		return
	}
	
	fmt.Println("Available Phases:")
	for _, phase := range phases {
		fmt.Printf("  - %s\n", phase)
	}
}

func createPhase(phaseID string) {
	fmt.Printf("Creating phase: %s\n", phaseID)
	// 实现阶段创建逻辑
}

func updatePhase(phaseID string) {
	fmt.Printf("Updating phase: %s\n", phaseID)
	// 实现阶段更新逻辑
}

// Team 命令实现
func listTeams(cmd *cobra.Command, args []string) {
	cm := getConfigManager()
	teams, err := cm.loader.ListTeams()
	if err != nil {
		fmt.Printf("Error listing teams: %v\n", err)
		return
	}
	
	fmt.Println("Available Teams:")
	for _, team := range teams {
		fmt.Printf("  - %s\n", team)
	}
}

func createTeam(teamID string) {
	fmt.Printf("Creating team: %s\n", teamID)
	// 实现团队创建逻辑
}

func updateTeam(teamID string) {
	fmt.Printf("Updating team: %s\n", teamID)
	// 实现团队更新逻辑
}

// Task 命令实现
func listTasks(cmd *cobra.Command, args []string) {
	cm := getConfigManager()
	tasks, err := cm.loader.ListTasks()
	if err != nil {
		fmt.Printf("Error listing tasks: %v\n", err)
		return
	}
	
	fmt.Println("Available Tasks:")
	for _, task := range tasks {
		fmt.Printf("  - %s\n", task)
	}
}

func createTask(taskID string) {
	fmt.Printf("Creating task: %s\n", taskID)
	// 实现任务创建逻辑
}

func updateTask(taskID string) {
	fmt.Printf("Updating task: %s\n", taskID)
	// 实现任务更新逻辑
}

// 通用命令实现
func deleteConfig(configType, configID string) {
	cm := getConfigManager()
	if err := cm.DeleteConfig(configType, configID); err != nil {
		fmt.Printf("Error deleting %s %s: %v\n", configType, configID, err)
		return
	}
	
	fmt.Printf("Successfully deleted %s: %s\n", configType, configID)
}

func cloneConfig(configType, sourceID, newID string) {
	cm := getConfigManager()
	if err := cm.CloneConfig(configType, sourceID, newID); err != nil {
		fmt.Printf("Error cloning %s %s to %s: %v\n", configType, sourceID, newID, err)
		return
	}
	
	fmt.Printf("Successfully cloned %s %s to %s\n", configType, sourceID, newID)
}

// 主函数
func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
