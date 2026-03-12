---
name: ralph-wiggum-loop
description: "Ralph Wiggum Loop AI专家 - 专门负责持续质量改进循环的自动化质量保证专业技能"
---

# Ralph Wiggum Loop AI专家

## 🎯 技能角色

### 核心职责
- **持续改进**: 自动执行质量检查和改进循环
- **质量监控**: 实时监控项目质量和健康状态
- **问题检测**: 主动发现潜在的质量问题和风险
- **改进建议**: 提供具体的质量改进方案

### 专业能力
- **循环管理**: 管理完整的质量检查循环流程
- **多维度监控**: 监控代码质量、性能、安全、文档等多个维度
- **智能分析**: 分析质量趋势和问题模式
- **自动化执行**: 自动执行质量检查和简单修复

## 🔧 功能实现

### 输入格式
```yaml
input:
  project_path: string         # 项目路径
  loop_config: object          # 循环配置
  trigger_event: string        # 触发事件
  quality_targets: map[string]float  # 质量目标
  auto_improvement: boolean     # 自动改进开关
```

### 输出格式
```yaml
output:
  loop_status: string           # 循环状态：running/completed/paused
  quality_metrics: object       # 质量指标
  detected_issues: object[]     # 检测到的问题
  improvement_actions: object[] # 改进措施
  trend_analysis: object        # 趋势分析
  next_schedule: object        # 下次执行计划
```

## 🔄 Ralph Wiggum Loop设计

### 循环触发条件
```yaml
triggers:
  automatic:
    - milestone_start: "每个Milestone启动时"
    - quality_gate_failure: "质量门禁失败时"
    - code_commit: "重要代码提交后"
    - time_interval: "每24小时自动触发"
    
  manual:
    - user_request: "用户手动触发"
    - quality_review: "质量审查会议后"
    - release_preparation: "发布准备阶段"
```

### 循环执行步骤
```yaml
execution_steps:
  1. "上下文检查": 验证HARNESS.md、WBS.md、OpenSpec配置
  2. "质量检查": 运行代码质量、性能、安全检查
  3. "功能测试": 运行单元测试、集成测试、端到端测试
  4. "性能验证": 检查性能指标是否符合标准
  5. "文档一致性": 验证文档与代码的一致性
  6. "质量评分": 计算综合质量分数
  7. "改进建议": 生成具体的改进建议
  8. "状态更新": 更新WBS.md中的里程碑状态
  9. "趋势分析": 分析质量趋势和模式
  10. "循环记录": 记录循环执行历史
```

### 质量门禁定义
```yaml
quality_gates:
  code_quality:
    threshold: 80
    tools: ["golangci-lint", "gosec", "ineffassign"]
    auto_fix: true
    
  test_coverage:
    threshold: 85
    command: "go test -cover"
    auto_fix: false
    
  documentation:
    threshold: 90
    check: "public_api_documented"
    auto_fix: true
    
  performance:
    threshold: "baseline_+10%"
    benchmark: "go test -bench"
    auto_fix: false
    
  security:
    threshold: 0  # 零安全漏洞
    tools: ["gosec", "gosec-vuln"]
    auto_fix: false
```

## 🎯 使用示例

### 完整循环执行
```bash
@go "使用ralph-wiggum-loop技能执行完整质量循环：
项目：PO System Implementation
路径：/workspace/skills/po-core-v2
触发：milestone_completion
自动改进：true
质量目标：代码质量85，测试覆盖率90，文档95"
```

### 手动质量检查
```bash
@go "使用ralph-wiggum-loop技能进行质量检查：
目标：当前技能实现
检查类型：comprehensive
严格模式：true
自动修复：true"
```

### 趋势分析
```bash
@go "使用ralph-wiggum-loop技能分析质量趋势：
项目：PO Core v2
时间范围：最近7天
分析维度：code_quality, test_coverage, performance
预测模型：linear_trend"
```

## 🔍 质量检查实现

### 代码质量检查
```yaml
code_quality_checks:
  static_analysis:
    - tool: "golangci-lint"
      config: ".golangci.yaml"
      threshold: 80
      auto_fix: true
      
    - tool: "gosec"
      config: "gosec.json"
      threshold: 0  # 零安全问题
      auto_fix: false
      
  code_metrics:
    - cyclomatic_complexity: "< 10"
    - maintainability_index: "> 70"
    - code_duplication: "< 5%"
    - technical_debt: "< 8 hours"
```

### 测试质量检查
```yaml
test_quality_checks:
  coverage_analysis:
    - overall_coverage: ">= 85%"
    - unit_test_coverage: ">= 80%"
    - integration_test_coverage: ">= 70%"
    
  test_effectiveness:
    - mutation_score: ">= 70%"
    - test_execution_time: "< 5 minutes"
    - test_stability: "> 95%"
```

### 性能质量检查
```yaml
performance_checks:
  response_time:
    - api_endpoints: "< 100ms (95th percentile)"
    - database_queries: "< 50ms"
    - file_operations: "< 200ms"
    
  resource_usage:
    - memory_usage: "< 512MB"
    - cpu_usage: "< 50%"
    - disk_io: "< 100MB/s"
```

### 文档质量检查
```yaml
documentation_checks:
  completeness:
    - api_documentation: "100%"
    - code_comments: "> 80%"
    - readme_files: "100%"
    
  quality:
    - documentation_format: "markdown_standard"
    - examples_provided: "true"
    - troubleshooting_guide: "true"
```

## 📊 质量评分算法

### 评分维度
```yaml
scoring_dimensions:
  code_quality: 30%           # 代码质量权重
  test_coverage: 25%          # 测试覆盖率权重
  performance: 20%            # 性能权重
  documentation: 15%          # 文档权重
  security: 10%              # 安全权重
```

### 评分计算
```yaml
score_calculation:
  code_quality_score:
    lint_score: 40%
    metrics_score: 30%
    complexity_score: 30%
    
  test_coverage_score:
    overall_coverage: 50%
    unit_test_coverage: 30%
    integration_test_coverage: 20%
    
  performance_score:
    response_time: 40%
    resource_usage: 30%
    scalability: 30%
    
  documentation_score:
    completeness: 50%
    quality: 30%
    examples: 20%
    
  security_score:
    vulnerability_count: 60%
    security_practices: 40%
    
  total_score: "weighted_sum_of_all_dimensions"
```

## 🚨 问题检测和分类

### 问题严重性分类
```yaml
severity_levels:
  critical:
    - security_vulnerabilities
    - data_corruption_risks
    - system_crashes
    - performance_degradation
    
  major:
    - test_coverage_gaps
    - code_quality_issues
    - documentation_missing
    - performance_bottlenecks
    
  minor:
    - code_style_issues
    - minor_documentation_gaps
    - optimization_opportunities
```

### 自动改进措施
```yaml
auto_improvements:
  code_style:
    - formatting: "gofmt"
    - import_ordering: "goimports"
    - unused_code: "goimports"
    
  documentation:
    - api_docs: "自动生成API文档"
    - code_comments: "自动生成注释模板"
    - readme_update: "自动更新README"
    
  dependencies:
    - dependency_cleanup: "go mod tidy"
    - vulnerability_fix: "自动更新依赖"
    - version_update: "自动更新版本"
```

## 📈 趋势分析

### 质量趋势分析
```yaml
trend_analysis:
  metrics_tracked:
    - code_quality_trend: "代码质量趋势"
    - test_coverage_trend: "测试覆盖率趋势"
    - performance_trend: "性能趋势"
    - security_trend: "安全趋势"
    
  analysis_methods:
    - moving_average: "7天移动平均"
    - linear_regression: "线性回归预测"
    - anomaly_detection: "异常检测"
    - seasonal_analysis: "季节性分析"
```

### 预测模型
```yaml
prediction_models:
  quality_degradation:
    features: ["code_churn", "complexity_increase", "test_decrease"]
    model: "random_forest"
    accuracy: "> 85%"
    
  performance_regression:
    features: ["code_size", "dependency_count", "resource_usage"]
    model: "linear_regression"
    accuracy: "> 80%"
```

## 🔄 循环历史记录

### 执行历史
```yaml
execution_history:
  record_format:
    - timestamp: "执行时间"
    - trigger_event: "触发事件"
    - quality_score: "质量分数"
    - issues_found: "发现问题数"
    - improvements_made: "改进措施数"
    - execution_time: "执行时长"
    
  retention_policy:
    - keep_records: "90天"
    - archive_old: "true"
    - summary_reports: "monthly"
```

### 改进跟踪
```yaml
improvement_tracking:
  improvement_types:
    - auto_fixed: "自动修复的问题"
    - manual_required: "需要手动修复的问题"
    - in_progress: "正在进行中的改进"
    - completed: "已完成的改进"
    
  success_metrics:
    - fix_rate: "修复率"
    - regression_rate: "回归率"
    - improvement_velocity: "改进速度"
```

## 🎯 使用示例

### 技能级Ralph Wiggum Loop
```bash
@go "使用ralph-wiggum-loop技能为po-core-v2建立质量循环：
目标技能：po-core-v2
循环频率：每次执行后
质量目标：综合分数85+
自动改进：true
监控维度：code_quality, test_coverage, performance"

# 自动执行流程：
# 1. 检查HARNESS.md合规性
# 2. 运行代码质量检查
# 3. 执行测试覆盖率检查
# 4. 验证性能指标
# 5. 检查文档完整性
# 6. 计算综合质量分数
# 7. 生成改进建议
# 8. 自动修复简单问题
# 9. 更新质量趋势
# 10. 记录循环历史
```

### 项目级质量监控
```bash
@go "使用ralph-wiggum-loop技能监控整个PO System项目：
监控范围：所有技能和代码
检查频率：每日自动执行
报告格式：详细报告 + 摘要
预警机制：质量下降超过10%时告警"

# 输出：
# - 每日质量报告
# - 质量趋势图表
# - 问题清单和改进建议
# - 预警通知（如果需要）
```

## 📊 质量标准

### 检查准确性
- **问题检测准确率**: > 90%
- **误报率**: < 10%
- **漏报率**: < 5%

### 性能要求
- **循环执行时间**: < 5分钟
- **内存使用**: < 200MB
- **并发支持**: 支持多项目并行监控

### 可靠性要求
- **循环稳定性**: > 99%
- **数据完整性**: 100%
- **自动修复安全性**: 不破坏原有功能

## 🚀 集成方式

### 与技能执行框架集成
```yaml
skill_execution_integration:
  pre_execution:
    - "质量基线检查"
    - "环境验证"
    
  post_execution:
    - "质量变化检测"
    - "自动改进执行"
    - "循环记录更新"
```

### 与harness-integrator集成
```yaml
harness_integration:
  data_flow:
    - "harness-integrator提供合规检查结果"
    - "ralph-wiggum-loop基于结果执行循环"
    - "改进建议反馈给harness-integrator"
    
  coordination:
    - "共享质量指标"
    - "协调改进措施"
    - "统一质量标准"
```

### 与po-core-v2集成
```yaml
po_core_integration:
  execution_flow:
    1. "po-core-v2执行项目处理"
    2. "ralph-wiggum-loop自动触发质量检查"
    3. "生成质量报告和改进建议"
    4. "反馈给po-core-v2优化下次执行"
    
  quality_assurance:
    - "持续监控po-core-v2输出质量"
    - "自动优化执行参数"
    - "确保输出符合HARNESS.md标准"
```

---

这个技能专门负责Ralph Wiggum Loop的完整实现，确保PO System具有持续的质量改进能力，是整个质量保证体系的核心引擎。
