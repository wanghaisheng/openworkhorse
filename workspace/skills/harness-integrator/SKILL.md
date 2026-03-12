---
name: harness-integrator
description: "HARNESS.md集成AI专家 - 专门负责HARNESS.md规则的自动加载、解析、验证和执行的专业技能"
---

# HARNESS.md集成AI专家

## 🎯 技能角色

### 核心职责
- **规则解析**: 解析HARNESS.md中的约束和规则
- **自动验证**: 自动验证代码和行为是否符合HARNESS.md
- **违规检测**: 检测潜在的HARNESS.md违规
- **合规建议**: 提供合规性改进建议

### 专业能力
- **规则理解**: 深度理解HARNESS.md的所有约束和标准
- **自动检查**: 自动执行HARNESS.md定义的检查流程
- **违规分析**: 分析违规的根本原因和影响
- **改进指导**: 提供具体的合规性改进方案

## 🔧 功能实现

### 输入格式
```yaml
input:
  target_path: string          # 检查目标路径
  harness_path: string         # HARNESS.md文件路径
  check_type: string           # 检查类型：code/architecture/process
  strict_mode: boolean         # 严格模式检查
  auto_fix: boolean           # 自动修复简单违规
```

### 输出格式
```yaml
output:
  compliance_status: string     # 合规状态：compliant/partial_violations/critical_violations
  violations: object[]         # 违规项列表
  score: float                 # 合规分数（0-100）
  recommendations: string[]     # 改进建议
  auto_fixes: object[]         # 自动修复结果
  detailed_report: string      # 详细报告
```

## 📋 HARNESS.md规则解析

### 禁止清单解析
```yaml
prohibited_items:
  code_standards:
    - global_variables: "避免全局状态，使用依赖注入"
    - circular_dependencies: "模块间不允许循环引用"
    - direct_database_operations: "必须通过数据访问层"
    - hardcoded_configuration: "使用配置文件或环境变量"
    - unhandled_errors: "所有错误必须被妥善处理"
    - debug_code: "移除所有console.log和调试语句"
    - unverified_third_party_libraries: "必须经过安全评估"
    - backward_compatibility_breaks: "重大变更需要版本升级"
    
  architecture_design:
    - single_responsibility_violations: "每个模块只负责一个功能"
    - tight_coupling: "模块间通过接口解耦"
    - duplicate_code: "提取公共逻辑到共享模块"
    - over_engineering: "保持简洁实用的架构"
    - performance_ignorance: "设计阶段必须考虑性能影响"
```

### 必须遵守解析
```yaml
required_standards:
  code_quality:
    - test_coverage: "测试覆盖率 ≥ 85%"
    - layered_architecture: "严格分层架构"
    - complete_error_handling: "所有可能错误都要处理"
    - code_review: "所有代码必须经过同行审查"
    - documentation_completeness: "公共API必须有完整文档"
    - performance_benchmarks: "满足性能要求才能发布"
    - security_compliance: "通过安全扫描和合规检查"
    
  development_process:
    - clear_requirements: "需求必须清晰明确"
    - design_review: "重要设计必须经过团队评审"
    - test_driven: "关键功能采用TDD开发方式"
    - continuous_integration: "每次提交都要通过CI检查"
    - version_control: "严格遵循Git工作流规范"
    - release_process: "遵循标准发布流程和检查清单"
```

## 🎯 使用示例

### 完整合规检查
```bash
@go "使用harness-integrator技能进行完整合规检查：
目标：当前项目
路径：/workspace/skills/po-core-v2
检查类型：code
严格模式：true
自动修复：true"
```

### 架构合规检查
```bash
@go "使用harness-integrator技能检查架构合规：
目标：po-core-v2技能架构
检查类型：architecture
严格模式：true
自动修复：false"
```

### 流程合规检查
```bash
@go "使用harness-integrator技能检查开发流程：
目标：项目开发流程
检查类型：process
严格模式：false
自动修复：false"
```

## 🔍 自动检查实现

### 代码质量检查
```yaml
code_quality_checks:
  test_coverage:
    command: "go test ./... -coverprofile=coverage.out"
    threshold: 85
    auto_fix: false
    
  code_linting:
    command: "golangci-lint run ./..."
    threshold: 0  # 零错误
    auto_fix: true
    
  type_checking:
    command: "go vet ./..."
    threshold: 0
    auto_fix: false
    
  security_scanning:
    command: "gosec ./..."
    threshold: 0
    auto_fix: false
    
  dependency_check:
    command: "go mod tidy && go mod verify"
    threshold: 0
    auto_fix: true
```

### 架构设计检查
```yaml
architecture_checks:
  layer_violation:
    check: "import_analysis"
    pattern: "cross_layer_imports"
    auto_fix: false
    
  circular_dependency:
    check: "dependency_analysis"
    pattern: "circular_imports"
    auto_fix: false
    
  single_responsibility:
    check: "module_analysis"
    pattern: "multiple_responsibilities"
    auto_fix: false
    
  coupling_analysis:
    check: "coupling_metrics"
    pattern: "tight_coupling"
    auto_fix: false
```

### 开发流程检查
```yaml
process_checks:
  git_workflow:
    check: "git_history"
    pattern: "commit_message_format"
    auto_fix: false
    
  ci_cd_pipeline:
    check: "pipeline_status"
    pattern: "failed_jobs"
    auto_fix: false
    
  documentation:
    check: "doc_coverage"
    pattern: "missing_documentation"
    auto_fix: true
```

## 📊 合规评分算法

### 评分维度
```yaml
scoring_dimensions:
  compliance_weight: 70%      # 合规性权重
  quality_weight: 20%         # 代码质量权重
  process_weight: 10%         # 流程合规权重
```

### 评分计算
```yaml
score_calculation:
  compliance_score:
    critical_violations: -50
    major_violations: -20
    minor_violations: -5
    perfect_compliance: 100
    
  quality_score:
    test_coverage: 25%
    code_quality: 25%
    documentation: 25%
    performance: 25%
    
  process_score:
    git_workflow: 30%
    ci_cd_status: 30%
    documentation: 20%
    review_process: 20%
    
  total_score: "compliance_score * 0.7 + quality_score * 0.2 + process_score * 0.1"
```

## 🚨 违规检测和报告

### 违规分类
```yaml
violation_categories:
  critical:
    - security_vulnerabilities
    - data_corruption_risks
    - system_stability_issues
    
  major:
    - performance_problems
    - maintainability_issues
    - test_coverage_gaps
    
  minor:
    - code_style_issues
    - documentation_gaps
    - process_deviations
```

### 报告格式
```yaml
compliance_report:
  summary:
    overall_status: "partial_violations"
    compliance_score: 78.5
    total_violations: 12
    
  violations_by_category:
    critical: 1
    major: 5
    minor: 6
    
  detailed_violations:
    - category: "major"
      description: "测试覆盖率不足"
      location: "pkg/skills/execution_framework.go"
      recommendation: "增加单元测试覆盖率到85%以上"
      auto_fix_available: false
      
    - category: "minor"
      description: "文档不完整"
      location: "pkg/skills/loader.go"
      recommendation: "补充公共API文档"
      auto_fix_available: true
      
  auto_fix_results:
    fixed_count: 3
    remaining_count: 9
    success_rate: 25%
```

## 🔧 自动修复功能

### 可自动修复的问题
```yaml
auto_fixable_issues:
  code_style:
    - formatting_issues: "使用gofmt自动格式化"
    - import_ordering: "使用goimports自动整理"
    - unused_variables: "使用goimports自动移除"
    
  documentation:
    - missing_comments: "自动生成基础注释"
    - function_documentation: "自动生成函数文档模板"
    
  dependencies:
    - unused_imports: "自动移除未使用的导入"
    - dependency_cleanup: "自动整理依赖关系"
```

### 自动修复流程
```yaml
auto_fix_process:
  1. "识别可自动修复的问题"
  2. "生成修复方案"
  3. "执行自动修复"
  4. "验证修复结果"
  5. "更新违规状态"
  6. "生成修复报告"
```

## 🎯 使用示例

### 技能级HARNESS.md集成
```bash
@go "使用harness-integrator技能集成HARNESS.md到po-core-v2：
目标技能：po-core-v2
集成方式：自动检查
检查频率：每次执行前
违规处理：自动报告和建议修复"

# 自动执行流程：
# 1. 解析HARNESS.md规则
# 2. 检查po-core-v2技能文档
# 3. 检查依赖技能文档
# 4. 生成合规报告
# 5. 提供改进建议
```

### 项目级HARNESS.md检查
```bash
@go "使用harness-integrator技能检查整个项目：
目标：PO System Implementation项目
检查范围：所有技能和代码
严格模式：true
自动修复：true"

# 输出：
# - 完整的合规报告
# - 违规项列表
# - 自动修复结果
# - 改进建议
```

## 📊 质量标准

### 检查准确性
- **规则解析准确率**: > 95%
- **违规检测准确率**: > 90%
- **误报率**: < 5%

### 性能要求
- **检查速度**: < 30秒（中等项目）
- **内存使用**: < 100MB
- **并发支持**: 支持多项目并行检查

### 可靠性要求
- **检查稳定性**: > 99%
- **报告完整性**: 100%
- **自动修复安全性**: 不破坏原有功能

## 🚀 集成方式

### 与技能执行框架集成
```yaml
skill_execution_integration:
  pre_execution:
    - "自动检查HARNESS.md合规性"
    - "阻止严重违规的执行"
    
  post_execution:
    - "检查生成内容的合规性"
    - "自动修复简单违规"
    - "生成合规报告"
```

### 与Ralph Wiggum Loop集成
```yaml
ralph_wiggum_loop_integration:
  triggers:
    - "skill_execution_start"
    - "code_generation_complete"
    - "milestone_completion"
    
  actions:
    - "harness_compliance_check"
    - "violation_reporting"
    - "auto_fix_execution"
    - "quality_score_update"
```

---

这个技能专门负责HARNESS.md的集成，确保所有技能和代码都符合项目约束标准，是PO System Implementation质量保证体系的重要组成部分。
