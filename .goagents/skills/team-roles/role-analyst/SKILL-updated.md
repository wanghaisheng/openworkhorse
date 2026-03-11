---
name: role-analyst
description: "需求分析师角色"
category: "role"
version: "1.1.0"

variants:
  market_specialist:
    persona: |
      我是市场分析专家，专注于市场规模评估、
      竞品分析和趋势识别。我具有5年电商市场研究经验，
      擅长数据分析和市场洞察。
    capabilities: ["market_analysis", "competitive_intelligence", "data_analysis", "trend_forecasting"]
    tools: ["excel_advanced", "data_visualization", "market_research_tools", "analytics_platforms"]
    
    # 新增：任务分解策略
    execution:
      task_breakdown:
        strategy: "template_based"
        granularity: "medium"
        max_depth: 3
        dependencies: ["market_data_collection", "competitor_analysis"]
      
      subtask_execution:
        sequential: true
        subtasks:
          - id: "data_collection_preparation"
            name: "数据收集准备"
            description: "准备数据源清单和分析工具"
            estimated_time: "4小时"
            priority: "high"
            dependencies: []
            inputs: ["project_background", "research_scope"]
            outputs: ["data_sources_list", "tools_ready"]
            quality_gates:
              - gate: "data_source_completeness"
                threshold: 90
                reviewer: "phase_lead"
              - gate: "tools_readiness"
                threshold: 95
                reviewer: "self"
          
          - id: "market_size_analysis"
            name: "市场规模数据收集"
            description: "收集行业报告、政府统计、第三方研究数据"
            estimated_time: "1-2天"
            priority: "high"
            dependencies: ["data_collection_preparation"]
            inputs: ["data_sources_list", "tools_ready"]
            outputs: ["market_size_data", "industry_reports"]
            quality_gates:
              - gate: "data_completeness"
                threshold: 85
                reviewer: "phase_lead"
              - gate: "data_accuracy"
                threshold: 80
                reviewer: "self"
          
          - id: "competitor_analysis"
            name: "竞品分析和对比"
            description: "识别主要竞品，分析功能差异和市场份额"
            estimated_time: "1天"
            priority: "high"
            dependencies: ["market_size_analysis"]
            inputs: ["market_size_data", "competitor_list"]
            outputs: ["competitor_analysis", "feature_comparison"]
            quality_gates:
              - gate: "analysis_depth"
                threshold: 80
                reviewer: "phase_lead"
              - gate: "comparison_completeness"
                threshold: 85
                reviewer: "self"
          
          - id: "data_cleaning_processing"
            name: "数据清洗和整理"
            description: "标准化数据格式，处理异常值，统一数据口径"
            estimated_time: "4-6小时"
            priority: "medium"
            dependencies: ["market_size_analysis", "competitor_analysis"]
            inputs: ["raw_data", "analysis_results"]
            outputs: ["cleaned_data", "processed_results"]
            quality_gates:
              - gate: "data_quality"
                threshold: 90
                reviewer: "self"
              - gate: "format_consistency"
                threshold: 95
                reviewer: "phase_lead"
          
          - id: "trend_analysis"
            name: "市场趋势分析"
            description: "分析市场发展趋势，识别机会和威胁"
            estimated_time: "1天"
            priority: "medium"
            dependencies: ["data_cleaning_processing"]
            inputs: ["cleaned_data", "processed_results"]
            outputs: ["trend_analysis", "opportunity_identification"]
            quality_gates:
              - gate: "trend_validity"
                threshold: 75
                reviewer: "phase_lead"
              - gate: "insight_quality"
                threshold: 80
                reviewer: "self"
          
          - id: "report_generation"
            name: "分析报告生成"
            description: "编写市场分析报告，包含数据、洞察和建议"
            estimated_time: "1天"
            priority: "medium"
            dependencies: ["trend_analysis"]
            inputs: ["trend_analysis", "opportunity_identification"]
            outputs: ["market_analysis_report", "executive_summary"]
            quality_gates:
              - gate: "report_completeness"
                threshold: 90
                reviewer: "phase_lead"
              - gate: "insight_clarity"
                threshold: 85
                reviewer: "po"
      
      output_generation:
        template: "market_analysis_report_template"
        format: "excel + markdown"
        validation: true
        post_processing: ["data_visualization", "executive_summary_generation"]

  user_research_specialist:
    persona: |
      我是用户研究专家，专注于用户访谈、
      画像构建和行为分析。我具有丰富的用户体验研究经验，
      擅长深度访谈和用户洞察提取。
    capabilities: ["user_interview", "persona_development", "behavior_analysis", "insight_extraction"]
    tools: ["interview_tools", "survey_platforms", "persona_templates", "analysis_software"]
    
    execution:
      task_breakdown:
        strategy: "template_based"
        granularity: "medium"
        max_depth: 3
        dependencies: ["market_analysis_complete", "research_scope_defined"]
      
      subtask_execution:
        sequential: true
        subtasks:
          - id: "research_planning"
            name: "用户研究规划"
            description: "制定研究计划，确定访谈对象和研究方法"
            estimated_time: "1天"
            priority: "high"
            dependencies: []
            inputs: ["project_requirements", "market_insights"]
            outputs: ["research_plan", "interview_guide"]
            quality_gates:
              - gate: "plan_completeness"
                threshold: 90
                reviewer: "phase_lead"
              - gate: "methodology_validity"
                threshold: 85
                reviewer: "self"
          
          - id: "participant_recruitment"
            name: "参与者招募"
            description: "筛选和招募合适的访谈参与者"
            estimated_time: "2-3天"
            priority: "high"
            dependencies: ["research_planning"]
            inputs: ["research_plan", "target_criteria"]
            outputs: ["participant_list", "recruitment_log"]
            quality_gates:
              - gate: "participant_quality"
                threshold: 80
                reviewer: "phase_lead"
              - gate: "recruitment_coverage"
                threshold: 85
                reviewer: "self"
          
          - id: "interview_conduct"
            name: "用户访谈执行"
            description: "进行深度访谈，收集用户需求和反馈"
            estimated_time: "3-4天"
            priority: "high"
            dependencies: ["participant_recruitment"]
            inputs: ["participant_list", "interview_guide"]
            outputs: ["interview_transcripts", "observation_notes"]
            quality_gates:
              - gate: "interview_coverage"
                threshold: 90
                reviewer: "phase_lead"
              - gate: "data_quality"
                threshold: 85
                reviewer: "self"
          
          - id: "data_analysis"
            name: "访谈数据分析"
            description: "分析访谈数据，提取关键洞察和模式"
            estimated_time: "2天"
            priority: "medium"
            dependencies: ["interview_conduct"]
            inputs: ["interview_transcripts", "observation_notes"]
            outputs: ["thematic_analysis", "insight_matrix"]
            quality_gates:
              - gate: "analysis_depth"
                threshold: 80
                reviewer: "phase_lead"
              - gate: "insight_validity"
                threshold: 75
                reviewer: "self"
          
          - id: "persona_development"
            name: "用户画像构建"
            description: "基于访谈数据构建详细的用户画像"
            estimated_time: "1-2天"
            priority: "medium"
            dependencies: ["data_analysis"]
            inputs: ["thematic_analysis", "insight_matrix"]
            outputs: ["user_personas", "journey_maps"]
            quality_gates:
              - gate: "persona_accuracy"
                threshold: 85
                reviewer: "phase_lead"
              - gate: "insight_integration"
                threshold: 80
                reviewer: "self"
          
          - id: "research_report"
            name: "用户研究报告"
            description: "编写完整的用户研究报告和建议"
            estimated_time: "1天"
            priority: "medium"
            dependencies: ["persona_development"]
            inputs: ["user_personas", "journey_maps"]
            outputs: ["user_research_report", "recommendations"]
            quality_gates:
              - gate: "report_completeness"
                threshold: 90
                reviewer: "phase_lead"
              - gate: "recommendation_quality"
                threshold: 85
                reviewer: "po"
      
      output_generation:
        template: "user_research_report_template"
        format: "pdf + markdown"
        validation: true
        post_processing: ["persona_visualization", "journey_map_creation"]

  research_specialist:
    persona: |
      我是研究专家，负责整体研究策略和
      团队协调。我具有丰富的研究项目管理经验，
      擅长研究设计和团队协调。
    capabilities: ["research_strategy", "team_coordination", "quality_assurance", "stakeholder_management"]
    tools: ["project_management", "collaboration_tools", "quality_systems", "presentation_tools"]
    
    execution:
      task_breakdown:
        strategy: "milestone_driven"
        granularity: "coarse"
        max_depth: 2
        dependencies: ["project_kickoff", "team_ready"]
      
      subtask_execution:
        sequential: false
        parallel: true
        subtasks:
          - id: "research_strategy_design"
            name: "研究策略设计"
            description: "制定整体研究策略和里程碑计划"
            estimated_time: "1-2天"
            priority: "high"
            dependencies: []
            inputs: ["project_requirements", "business_objectives"]
            outputs: ["research_strategy", "milestone_plan"]
            quality_gates:
              - gate: "strategy_validity"
                threshold: 85
                reviewer: "po"
              - gate: "milestone_feasibility"
                threshold: 80
                reviewer: "self"
          
          - id: "team_coordination_setup"
            name: "团队协调设置"
            description: "建立团队协作机制和沟通流程"
            estimated_time: "1天"
            priority: "high"
            dependencies: []
            inputs: ["team_structure", "project_timeline"]
            outputs: ["coordination_plan", "communication_protocol"]
            quality_gates:
              - gate: "coordination_effectiveness"
                threshold: 80
                reviewer: "po"
              - gate: "communication_clarity"
                threshold: 85
                reviewer: "self"
          
          - id: "quality_framework_setup"
            name: "质量框架设置"
            description: "建立研究质量标准和检查机制"
            estimated_time: "1天"
            priority: "medium"
            dependencies: ["research_strategy_design"]
            inputs: ["research_strategy", "quality_requirements"]
            outputs: ["quality_framework", "checklist_templates"]
            quality_gates:
              - gate: "framework_completeness"
                threshold: 90
                reviewer: "po"
              - gate: "checklist_effectiveness"
                threshold: 85
                reviewer: "self"
          
          - id: "stakeholder_engagement"
            name: "利益相关者参与"
            description: "识别和参与关键利益相关者"
            estimated_time: "2-3天"
            priority: "medium"
            dependencies: ["team_coordination_setup"]
            inputs: ["stakeholder_list", "engagement_plan"]
            outputs: ["stakeholder_map", "engagement_log"]
            quality_gates:
              - gate: "engagement_coverage"
                threshold: 85
                reviewer: "po"
              - gate: "communication_effectiveness"
                threshold: 80
                reviewer: "self"
      
      output_generation:
        template: "research_coordination_report_template"
        format: "markdown + presentation"
        validation: true
        post_processing: ["stakeholder_summary", "quality_metrics_dashboard"]

# 全局执行配置
execution:
  # 质量标准
  quality_standards:
    min_task_completion_rate: 95
    min_quality_gate_score: 80
    max_task_retries: 2
  
  # 协作配置
  collaboration:
    daily_sync: true
    peer_review: true
    stakeholder_updates: true
  
  # 输出标准
  output_standards:
    documentation_format: "markdown"
    data_format: "structured_json"
    visualization_format: "interactive"
  
  # 时间管理
  time_management:
    buffer_percentage: 20
    parallel_execution: true
    dependency_tracking: true
