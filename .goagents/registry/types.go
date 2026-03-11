package registry

import "time"

// GlobalConfig 全局配置结构
type GlobalConfig struct {
	Version      string            `yaml:"version"`
	LastUpdated  string            `yaml:"last_updated"`
	System       SystemConfig      `yaml:"system"`
	Skills       SkillsConfig      `yaml:"skills"`
	Workflows    WorkflowsConfig   `yaml:"workflows"`
	Teams        TeamsConfig       `yaml:"teams"`
	Phases       PhasesConfig      `yaml:"phases"`
	Tasks        TasksConfig       `yaml:"tasks"`
	Logging      LoggingConfig     `yaml:"logging"`
	Monitoring   MonitoringConfig  `yaml:"monitoring"`
}

// SystemConfig 系统配置
type SystemConfig struct {
	DefaultTaskMode   string            `yaml:"default_task_mode"`
	SupportedModes     []string          `yaml:"supported_modes"`
	QualityGates       QualityGatesConfig `yaml:"quality_gates"`
	Granularity        GranularityConfig  `yaml:"granularity"`
}

// QualityGatesConfig 质量门禁配置
type QualityGatesConfig struct {
	TestCoverageMin   int `yaml:"test_coverage_min"`
	LintPassRate     int `yaml:"lint_pass_rate"`
	TaskSizeMin      int `yaml:"task_size_min"`
	TaskSizeMax      int `yaml:"task_size_max"`
}

// GranularityConfig 粒度控制配置
type GranularityConfig struct {
	SweetSpotLines    []int `yaml:"sweet_spot_lines"`
	MaxFilesAffected  int   `yaml:"max_files_affected"`
	EstimatedTimeRange []int `yaml:"estimated_time_range"` // minutes
}

// SkillsConfig 技能配置
type SkillsConfig struct {
	SearchPaths []string       `yaml:"search_paths"`
	Cache      SkillsCacheConfig `yaml:"cache"`
}

// SkillsCacheConfig 技能缓存配置
type SkillsCacheConfig struct {
	Enabled  bool `yaml:"enabled"`
	TTL      int  `yaml:"ttl"`       // seconds
	MaxSize  int  `yaml:"max_size"`
}

// WorkflowsConfig 工作流配置
type WorkflowsConfig struct {
	Default    string   `yaml:"default"`
	SearchPaths []string `yaml:"search_paths"`
}

// TeamsConfig 团队配置
type TeamsConfig struct {
	Default    string   `yaml:"default"`
	SearchPaths []string `yaml:"search_paths"`
}

// PhasesConfig 阶段配置
type PhasesConfig struct {
	Supported   []string `yaml:"supported"`
	SearchPaths []string `yaml:"search_paths"`
}

// TasksConfig 任务配置
type TasksConfig struct {
	SearchPaths    []string          `yaml:"search_paths"`
	TemplateCache  TemplateCacheConfig `yaml:"template_cache"`
}

// TemplateCacheConfig 模板缓存配置
type TemplateCacheConfig struct {
	Enabled      bool `yaml:"enabled"`
	MaxTemplates int  `yaml:"max_templates"`
}

// LoggingConfig 日志配置
type LoggingConfig struct {
	Level       string `yaml:"level"`
	Structured  bool   `yaml:"structured"`
	OutputDir   string `yaml:"output_dir"`
}

// MonitoringConfig 监控配置
type MonitoringConfig struct {
	Enabled bool     `yaml:"enabled"`
	Metrics []string `yaml:"metrics"`
}

// WorkflowConfig 工作流配置结构
type WorkflowConfig struct {
	ID               string               `yaml:"id"`
	Name             string               `yaml:"name"`
	Description       string               `yaml:"description"`
	Version          string               `yaml:"version"`
	Category         string               `yaml:"category"`
	Metadata         WorkflowMetadata     `yaml:"metadata"`
	ApplicableScenarios []string           `yaml:"applicable_scenarios"`
	Phases           []PhaseConfig       `yaml:"phases"`
	Transitions      []TransitionConfig    `yaml:"transitions"`
	WorkflowConfig   WorkflowLevelConfig `yaml:"workflow_config"`
	PerformanceBenchmarks PerformanceBenchmarks `yaml:"performance_benchmarks"`
	UsageStats       UsageStats           `yaml:"usage_stats"`
}

// WorkflowMetadata 工作流元数据
type WorkflowMetadata struct {
	Author      string    `yaml:"author"`
	Created     string    `yaml:"created"`
	LastUpdated string    `yaml:"last_updated"`
	Tags        []string  `yaml:"tags"`
}

// PhaseConfig 阶段配置结构
type PhaseConfig struct {
	ID               string               `yaml:"id"`
	Name             string               `yaml:"name"`
	Description       string               `yaml:"description"`
	Version          string               `yaml:"version"`
	Category         string               `yaml:"category"`
	Metadata         PhaseMetadata        `yaml:"metadata"`
	Objectives       ObjectivesConfig     `yaml:"objectives"`
	SuccessCriteria  SuccessCriteriaConfig `yaml:"success_criteria"`
	StandardTasks    []TaskTemplate      `yaml:"standard_tasks"`
	FreeModeConfig   FreeModeConfig       `yaml:"free_mode_config"`
	RoleRequirements RoleRequirementsConfig `yaml:"role_requirements"`
	Inputs           InputsConfig         `yaml:"inputs"`
	Outputs          OutputsConfig        `yaml:"outputs"`
	QualityGates     []QualityGate       `yaml:"quality_gates"`
	ToolsAndResources ToolsResourcesConfig `yaml:"tools_and_resources"`
	TimeEstimation   TimeEstimationConfig  `yaml:"time_estimation"`
	RiskFactors      RiskFactorsConfig     `yaml:"risk_factors"`
	CompletionChecklist CompletionChecklist  `yaml:"completion_checklist"`
}

// PhaseMetadata 阶段元数据
type PhaseMetadata struct {
	Author string   `yaml:"author"`
	Created string   `yaml:"created"`
	Tags    []string `yaml:"tags"`
}

// ObjectivesConfig 目标配置
type ObjectivesConfig struct {
	Primary   []string `yaml:"primary"`
	Secondary []string `yaml:"secondary"`
}

// SuccessCriteriaConfig 成功标准配置
type SuccessCriteriaConfig struct {
	Mandatory []string `yaml:"mandatory"`
	Optional  []string `yaml:"optional"`
}

// TaskTemplate 任务模板结构
type TaskTemplate struct {
	TemplateID     string            `yaml:"template_id"`
	Name           string            `yaml:"name"`
	Description    string            `yaml:"description"`
	EstimatedTime  string            `yaml:"estimated_time"`
	Priority       string            `yaml:"priority"`
	RequiredAgents []string          `yaml:"required_agents"`
	Deliverables   []Deliverable     `yaml:"deliverables"`
	QualityGates   []QualityGate      `yaml:"quality_gates"`
}

// Deliverable 交付物结构
type Deliverable struct {
	Name     string `yaml:"name"`
	Format   string `yaml:"format"`
	Template string `yaml:"template"`
}

// QualityGate 质量门禁结构
type QualityGate struct {
	Gate        string `yaml:"gate"`
	Threshold   int    `yaml:"threshold"`
	Reviewer    string `yaml:"reviewer"`
	CheckType   string `yaml:"check_type"`
}

// FreeModeConfig 自由模式配置
type FreeModeConfig struct {
	PhaseLeadAuthority PhaseLeadAuthorityConfig `yaml:"phase_lead_authority"`
	DynamicTaskRules   DynamicTaskRulesConfig   `yaml:"dynamic_task_rules"`
	QualityControl     QualityControlConfig     `yaml:"quality_control"`
}

// PhaseLeadAuthorityConfig 阶段领导权限配置
type PhaseLeadAuthorityConfig struct {
	TaskPlanning      bool `yaml:"task_planning"`
	TaskAssignment    bool `yaml:"task_assignment"`
	TaskAdjustment    bool `yaml:"task_adjustment"`
	MilestoneDefinition bool `yaml:"milestone_definition"`
}

// DynamicTaskRulesConfig 动态任务规则配置
type DynamicTaskRulesConfig struct {
	MaxTasksPerMilestone int     `yaml:"max_tasks_per_milestone"`
	MinTasksPerMilestone int     `yaml:"min_tasks_per_milestone"`
	TaskSizeRange        []int   `yaml:"task_size_range"` // lines of code
}

// QualityControlConfig 质量控制配置
type QualityControlConfig struct {
	PeerReview          bool `yaml:"peer_review"`
	MilestoneCheckpoints bool `yaml:"milestone_checkpoints"`
	FinalApproval       string `yaml:"final_approval"`
}

// RoleRequirementsConfig 角色需求配置
type RoleRequirementsConfig struct {
	MandatoryRoles []RoleRequirement `yaml:"mandatory_roles"`
	OptionalRoles  []RoleRequirement `yaml:"optional_roles"`
}

// RoleRequirement 角色需求结构
type RoleRequirement struct {
	Role             string   `yaml:"role"`
	AgentType        string   `yaml:"agent_type"`
	ExperienceLevel  string   `yaml:"experience_level"`
	Specializations []string `yaml:"specializations"`
}

// InputsConfig 输入配置
type InputsConfig struct {
	Required  []string `yaml:"required"`
	Optional  []string `yaml:"optional"`
}

// OutputsConfig 输出配置
type OutputsConfig struct {
	Artifacts []Artifact `yaml:"artifacts"`
	Metrics   []string   `yaml:"metrics"`
}

// Artifact 交付物结构
type Artifact struct {
	Name     string `yaml:"name"`
	Format   string `yaml:"format"`
	Template string `yaml:"template"`
}

// ToolsResourcesConfig 工具和资源配置
type ToolsResourcesConfig struct {
	RequiredTools []string `yaml:"required_tools"`
	OptionalTools []string `yaml:"optional_tools"`
}

// TimeEstimationConfig 时间估算配置
type TimeEstimationConfig struct {
	BaseDuration      string            `yaml:"base_duration"`
	ComplexityFactors map[string]float64 `yaml:"complexity_factors"`
	TeamSizeFactors  map[string]float64 `yaml:"team_size_factors"`
}

// RiskFactorsConfig 风险因素配置
type RiskFactorsConfig struct {
	HighRisk     []Risk     `yaml:"high_risk"`
	MediumRisk   []Risk     `yaml:"medium_risk"`
	MitigationStrategies []string `yaml:"mitigation_strategies"`
}

// Risk 风险结构
type Risk struct {
	Risk        string  `yaml:"risk"`
	Probability string  `yaml:"probability"`
	Impact     string  `yaml:"impact"`
	Mitigation  string  `yaml:"mitigation"`
}

// CompletionChecklist 完成检查清单
type CompletionChecklist struct {
	Mandatory []string `yaml:"mandatory"`
	Optional  []string `yaml:"optional"`
}

// TransitionConfig 转换配置结构
type TransitionConfig struct {
	From           string `yaml:"from"`
	To             string `yaml:"to"`
	Condition      string `yaml:"condition"`
	AutoTransition bool   `yaml:"auto_transition"`
}

// WorkflowLevelConfig 工作流级别配置
type WorkflowLevelConfig struct {
	ParallelExecution ParallelExecutionConfig `yaml:"parallel_execution"`
	Rollback        RollbackConfig        `yaml:"rollback"`
	Notifications  NotificationsConfig     `yaml:"notifications"`
}

// ParallelExecutionConfig 并行执行配置
type ParallelExecutionConfig struct {
	Enabled            bool `yaml:"enabled"`
	MaxParallelPhases  int  `yaml:"max_parallel_phases"`
}

// RollbackConfig 回滚配置
type RollbackConfig struct {
	Enabled                 bool     `yaml:"enabled"`
	AutoRollbackOnFailure   bool     `yaml:"auto_rollback_on_failure"`
	RollbackCheckpoints     []string `yaml:"rollback_checkpoints"`
}

// NotificationsConfig 通知配置
type NotificationsConfig struct {
	PhaseStart         bool `yaml:"phase_start"`
	PhaseComplete      bool `yaml:"phase_complete"`
	QualityGateFailure bool `yaml:"quality_gate_failure"`
	WorkflowComplete  bool `yaml:"workflow_complete"`
}

// PerformanceBenchmarks 性能基准
type PerformanceBenchmarks struct {
	TotalDuration    string            `yaml:"total_duration"`
	PhaseEfficiency  map[string]float64 `yaml:"phase_efficiency"`
	QualityScoreTarget float64           `yaml:"quality_score_target"`
}

// UsageStats 使用统计
type UsageStats struct {
	TimesUsed      int     `yaml:"times_used"`
	SuccessRate    float64 `yaml:"success_rate"`
	AverageDuration float64 `yaml:"average_duration"`
	LastUsed       *string `yaml:"last_used"`
}

// TeamConfig 团队配置结构
type TeamConfig struct {
	ID               string               `yaml:"id"`
	Name             string               `yaml:"name"`
	Description       string               `yaml:"description"`
	Version          string               `yaml:"version"`
	Category         string               `yaml:"category"`
	Metadata         TeamMetadata        `yaml:"metadata"`
	TeamLead         TeamLeadConfig      `yaml:"team_lead"`
	Phases           map[string]PhaseTeamConfig `yaml:"phases"`
	Collaboration    CollaborationConfig  `yaml:"collaboration"`
	QualityAssurance QualityAssuranceConfig `yaml:"quality_assurance"`
	SkillRequirements SkillRequirementsConfig `yaml:"skill_requirements"`
	PerformanceMetrics PerformanceMetricsConfig `yaml:"performance_metrics"`
	ResourceAllocation ResourceAllocationConfig `yaml:"resource_allocation"`
	RiskManagement   RiskManagementConfig   `yaml:"risk_management"`
	DynamicAdjustments DynamicAdjustmentsConfig `yaml:"dynamic_adjustments"`
	SuccessCriteria  SuccessCriteriaConfig `yaml:"success_criteria"`
}

// TeamMetadata 团队元数据
type TeamMetadata struct {
	Author           string   `yaml:"author"`
	Created          string   `yaml:"created"`
	Tags             []string `yaml:"tags"`
	ApplicablePhases []string `yaml:"applicable_phases"`
}

// TeamLeadConfig 团队领导配置
type TeamLeadConfig struct {
	Agent         string   `yaml:"agent"`
	Variant       string   `yaml:"variant"`
	Responsibilities []string `yaml:"responsibilities"`
	AuthorityLevel string   `yaml:"authority_level"`
	DecisionScope []string `yaml:"decision_scope"`
}

// PhaseTeamConfig 阶段团队配置
type PhaseTeamConfig struct {
	Description       string               `yaml:"description"`
	EstimatedDuration string               `yaml:"estimated_duration"`
	TaskMode         string               `yaml:"task_mode"`
	PhaseLead        PhaseLeadConfig      `yaml:"phase_lead"`
	TeamMembers      []TeamMemberConfig   `yaml:"team_members"`
}

// PhaseLeadConfig 阶段领导配置
type PhaseLeadConfig struct {
	Agent           string   `yaml:"agent"`
	Variant         string   `yaml:"variant"`
	Responsibilities []string `yaml:"responsibilities"`
	Capabilities    []string `yaml:"capabilities"`
	Priority        string   `yaml:"priority"`
	ExperienceLevel string   `yaml:"experience_level"`
	Specialization string   `yaml:"specialization"`
}

// TeamMemberConfig 团队成员配置
type TeamMemberConfig struct {
	MemberID        string   `yaml:"member_id"`
	Agent           string   `yaml:"agent"`
	Variant         string   `yaml:"variant"`
	Responsibilities []string `yaml:"responsibilities"`
	Capabilities    []string `yaml:"capabilities"`
	ExperienceLevel string   `yaml:"experience_level"`
	Specialization string   `yaml:"specialization"`
	ReportingTo     string   `yaml:"reporting_to"`
	Workload        string   `yaml:"workload"`
}

// CollaborationConfig 协作配置
type CollaborationConfig struct {
	Pattern           string               `yaml:"pattern"`
	Communication    CommunicationConfig  `yaml:"communication"`
	DecisionMaking   DecisionMakingConfig `yaml:"decision_making"`
	ConflictResolution ConflictResolutionConfig `yaml:"conflict_resolution"`
}

// CommunicationConfig 通信配置
type CommunicationConfig struct {
	PrimaryChannel    string   `yaml:"primary_channel"`
	SecondaryChannels []string `yaml:"secondary_channels"`
	MeetingFrequency  string   `yaml:"meeting_frequency"`
}

// DecisionMakingConfig 决策配置
type DecisionMakingConfig struct {
	ConsensusRequired  bool   `yaml:"consensus_required"`
	PhaseLeadAuthority string  `yaml:"phase_lead_authority"`
	EscalationPath    string  `yaml:"escalation_path"`
}

// ConflictResolutionConfig 冲突解决配置
type ConflictResolutionConfig struct {
	PrimaryMethod      string `yaml:"primary_method"`
	EscalationTrigger  string `yaml:"escalation_trigger"`
	FinalAuthority     string `yaml:"final_authority"`
}

// QualityAssuranceConfig 质量保证配置
type QualityAssuranceConfig struct {
	Reviews           ReviewConfig           `yaml:"reviews"`
	QualityGates      []QualityGate         `yaml:"quality_gates"`
	DeliverableStandards DeliverableStandardsConfig `yaml:"deliverable_standards"`
}

// ReviewConfig 审查配置
type ReviewConfig struct {
	PeerReview           bool `yaml:"peer_review"`
	LeadReview           bool `yaml:"lead_review"`
	CrossFunctionalReview bool `yaml:"cross_functional_review"`
}

// DeliverableStandardsConfig 交付物标准配置
type DeliverableStandardsConfig struct {
	Templates         bool `yaml:"templates"`
	FormatConsistency bool `yaml:"format_consistency"`
	VersionControl    bool `yaml:"version_control"`
}

// SkillRequirementsConfig 技能需求配置
type SkillRequirementsConfig struct {
	PhaseLead     SkillRequirement `yaml:"phase_lead"`
	TeamMembers   CommonSkillsConfig `yaml:"team_members"`
}

// CommonSkillsConfig 通用技能配置
type CommonSkillsConfig struct {
	CommonSkills    []string          `yaml:"common_skills"`
	SpecializedSkills map[string][]string `yaml:"specialized_skills"`
}

// PerformanceMetricsConfig 性能指标配置
type PerformanceMetricsConfig struct {
	Productivity     ProductivityConfig     `yaml:"productivity"`
	Collaboration  CollaborationConfig  `yaml:"collaboration"`
	Quality         QualityConfig         `yaml:"quality"`
}

// ProductivityConfig 生产力配置
type ProductivityConfig struct {
	TasksPerWeek         int     `yaml:"tasks_per_week"`
	DeliverablesPerWeek  int     `yaml:"deliverables_per_week"`
	QualityScoreTarget   float64 `yaml:"quality_score_target"`
}

// QualityConfig 质量配置
type QualityConfig struct {
	PeerReviewPassRate     float64 `yaml:"peer_review_pass_rate"`
	StakeholderSatisfaction float64 `yaml:"stakeholder_satisfaction"`
	DeliverableOnTimeRate float64 `yaml:"deliverable_on_time_rate"`
}

// ResourceAllocationConfig 资源配置配置
type ResourceAllocationConfig struct {
	Tools    ToolsConfig `yaml:"tools"`
	Budget   BudgetConfig `yaml:"budget"`
}

// ToolsConfig 工具配置
type ToolsConfig struct {
	SharedTools    []string                    `yaml:"shared_tools"`
	IndividualTools map[string][]string         `yaml:"individual_tools"`
}

// BudgetConfig 预算配置
type BudgetConfig struct {
	MarketResearchBudget string `yaml:"market_research_budget"`
	UserResearchBudget   string `yaml:"user_research_budget"`
	ToolsBudget         string `yaml:"tools_budget"`
}

// RiskManagementConfig 风险管理配置
type RiskManagementConfig struct {
	IdentifiedRisks []Risk    `yaml:"identified_risks"`
	MitigationStrategies []string `yaml:"mitigation_strategies"`
}

// DynamicAdjustmentsConfig 动态调整配置
type DynamicAdjustmentsConfig struct {
	ScalingRules    ScalingRulesConfig    `yaml:"scaling_rules"`
	RoleFlexibility RoleFlexibilityConfig `yaml:"role_flexibility"`
}

// ScalingRulesConfig 扩展规则配置
type ScalingRulesConfig struct {
	AddMemberCriteria    string `yaml:"add_member_criteria"`
	RemoveMemberCriteria string `yaml:"remove_member_criteria"`
}

// RoleFlexibilityConfig 角色灵活性配置
type RoleFlexibilityConfig struct {
	CrossTraining   bool `yaml:"cross_training"`
	RoleRotation    bool `yaml:"role_rotation"`
	SkillDevelopment bool `yaml:"skill_development"`
}

// TaskConfig 任务配置结构
type TaskConfig struct {
	ID                string                `yaml:"id"`
	Name              string                `yaml:"name"`
	Description        string                `yaml:"description"`
	Category          string                `yaml:"category"`
	Version           string                `yaml:"version"`
	Metadata          TaskMetadata          `yaml:"metadata"`
	ApplicableScenarios []string              `yaml:"applicable_scenarios"`
	Prerequisites     PrerequisitesConfig    `yaml:"prerequisites"`
	Objectives        ObjectivesConfig      `yaml:"objectives"`
	ExecutionSteps    []ExecutionStep       `yaml:"execution_steps"`
	RequiredSkills    SkillsConfig         `yaml:"required_skills"`
	QualityStandards  QualityStandardsConfig `yaml:"quality_standards"`
	ToolsRequired     ToolsConfig          `yaml:"tools_required"`
	RiskFactors       RiskFactorsConfig    `yaml:"risk_factors"`
	SuccessMetrics    SuccessMetricsConfig  `yaml:"success_metrics"`
	FollowUpActions   FollowUpActionsConfig `yaml:"follow_up_actions"`
	VersionHistory    []VersionHistory      `yaml:"version_history"`
}

// TaskMetadata 任务元数据
type TaskMetadata struct {
	Author string   `yaml:"author"`
	Created string   `yaml:"created"`
	Tags    []string `yaml:"tags"`
	ApplicablePhases []string `yaml:"applicable_phases"`
	Difficulty string `yaml:"difficulty"`
	EstimatedTime string `yaml:"estimated_time"`
}

// PrerequisitesConfig 前置条件配置
type PrerequisitesConfig struct {
	Required  []string `yaml:"required"`
	Optional  []string `yaml:"optional"`
}

// ExecutionStep 执行步骤结构
type ExecutionStep struct {
	StepID       string         `yaml:"step_id"`
	Name         string         `yaml:"name"`
	Description   string         `yaml:"description"`
	EstimatedTime string         `yaml:"estimated_time"`
	Priority     string         `yaml:"priority"`
	DependsOn    []string       `yaml:"depends_on"`
	Inputs       InputsConfig   `yaml:"inputs"`
	Activities   []string       `yaml:"activities"`
	Outputs      OutputsConfig  `yaml:"outputs"`
	QualityGates []QualityGate  `yaml:"quality_gates"`
}

// QualityStandardsConfig 质量标准配置
type QualityStandardsConfig struct {
	DataQuality     DataQualityConfig     `yaml:"data_quality"`
	AnalysisQuality AnalysisQualityConfig `yaml:"analysis_quality"`
	DeliverableQuality DeliverableQualityConfig `yaml:"deliverable_quality"`
}

// DataQualityConfig 数据质量配置
type DataQualityConfig struct {
	Completeness string `yaml:"completeness"`
	Accuracy    string `yaml:"accuracy"`
	Consistency string `yaml:"consistency"`
}

// AnalysisQualityConfig 分析质量配置
type AnalysisQualityConfig struct {
	MethodologySoundness   string `yaml:"methodology_soundness"`
	InsightRelevance      string `yaml:"insight_relevance"`
	RecommendationActionability string `yaml:"recommendation_actionability"`
}

// DeliverableQualityConfig 交付物质量配置
type DeliverableQualityConfig struct {
	FormatCompliance   string `yaml:"format_compliance"`
	TemplateAdherence  string `yaml:"template_adherence"`
	ClarityReadability string `yaml:"clarity_readability"`
}

// SuccessMetricsConfig 成功指标配置
type SuccessMetricsConfig struct {
	Quantitative []string `yaml:"quantitative"`
	Qualitative []string `yaml:"qualitative"`
}

// FollowUpActionsConfig 后续行动配置
type FollowUpActionsConfig struct {
	Immediate  []string `yaml:"immediate"`
	ShortTerm  []string `yaml:"short_term"`
	LongTerm   []string `yaml:"long_term"`
}

// VersionHistory 版本历史
type VersionHistory struct {
	Version string `yaml:"version"`
	Date    string `yaml:"date"`
	Changes string `yaml:"changes"`
	Author  string `yaml:"author"`
}
