---
name: standard-mode-executor
description: "Standard模式执行AI专家 - 专门负责基于模板的标准化项目执行流程的专业技能"
---

# Standard模式执行AI专家

## 🎯 技能角色

### 核心职责
- **模板执行**: 基于预定义模板执行标准化项目流程
- **流程控制**: 严格控制项目执行的每个阶段和步骤
- **质量保证**: 确保每个步骤都符合标准化质量要求
- **进度管理**: 监控和管理项目执行进度

### 专业能力
- **模板管理**: 管理和维护各种项目执行模板
- **流程编排**: 编排和优化标准化执行流程
- **质量控制**: 实施严格的质量门禁和检查点
- **进度跟踪**: 实时跟踪项目执行状态和里程碑

## 🔧 功能实现

### 输入格式
```yaml
input:
  project_request: object       # 项目请求信息
  execution_template: string    # 执行模板名称
  quality_gates: object[]      # 质量门禁配置
  schedule_config: object      # 进度配置
  team_composition: map[string]string  # 团队组成
```

### 输出格式
```yaml
output:
  execution_plan: object        # 详细执行计划
  milestone_schedule: object[] # 里程碑时间表
  quality_checkpoints: object[] # 质量检查点
  progress_tracking: object    # 进度跟踪信息
  risk_assessment: object      # 风险评估
  completion_status: string    # 完成状态
```

## 📋 Standard模式定义

### 模式特征
```yaml
standard_mode_characteristics:
  execution_style: "template_based"
  flexibility: "low"
  predictability: "high"
  quality_control: "strict"
  documentation: "comprehensive"
  risk_management: "proactive"
```

### 适用场景
```yaml
applicable_scenarios:
  project_types:
    - "企业级应用开发"
    - "金融系统开发"
    - "政府项目"
    - "大型电商平台"
    
  team_characteristics:
    - "中大型团队（10-50人）"
    - "多角色协作"
    - "严格的质量要求"
    - "明确的交付时间"
    
  requirements:
    - "需求相对稳定"
    - "技术栈成熟"
    - "合规性要求高"
    - "可预测的结果"
```

## 🎯 使用示例

### 完整Standard模式执行
```bash
@go "使用standard-mode-executor技能执行Standard模式项目：
项目：企业级CRM系统
模板：enterprise_web_application
团队：15人（包含po、analyst、architect、developer、qa）
时间：3个月
质量要求：严格遵循HARNESS.md"
```

### 模板选择和执行
```bash
@go "使用standard-mode-executor技能选择并执行模板：
项目类型：电商平台
规模：中大型
技术栈：Go + React + PostgreSQL
团队：20人
约束：高可用、高性能、安全合规"
```

## 📊 执行模板库

### 企业级Web应用模板
```yaml
enterprise_web_application:
  phases:
    - name: "需求分析和规划"
      duration: "2周"
      deliverables: ["需求文档", "技术方案", "项目计划"]
      quality_gates: ["需求完整性", "方案可行性", "计划合理性"]
      
    - name: "系统设计和架构"
      duration: "2周"
      deliverables: ["架构设计", "数据库设计", "API设计"]
      quality_gates: ["架构合理性", "设计完整性", "性能考虑"]
      
    - name: "核心功能开发"
      duration: "6周"
      deliverables: ["核心模块", "单元测试", "集成测试"]
      quality_gates: ["代码质量", "测试覆盖率", "功能完整性"]
      
    - name: "系统集成和测试"
      duration: "2周"
      deliverables: ["完整系统", "端到端测试", "性能测试"]
      quality_gates: ["系统稳定性", "性能达标", "安全合规"]
      
    - name: "部署和上线"
      duration: "1周"
      deliverables: ["部署文档", "运维手册", "用户培训"]
      quality_gates: ["部署成功", "文档完整", "用户满意"]
      
  total_duration: "13周"
  team_size: "10-20人"
  quality_standard: "enterprise_grade"
```

### 电商平台模板
```yaml
ecommerce_platform:
  phases:
    - name: "市场调研和需求分析"
      duration: "1.5周"
      deliverables: ["市场分析报告", "用户需求文档", "竞品分析"]
      quality_gates: ["市场洞察", "需求清晰", "竞争策略"]
      
    - name: "产品设计和原型"
      duration: "2周"
      deliverables: ["产品设计", "交互原型", "UI设计"]
      quality_gates: ["设计一致性", "用户体验", "视觉效果"]
      
    - name: "技术架构设计"
      duration: "1.5周"
      deliverables: ["技术架构", "数据库设计", "安全方案"]
      quality_gates: ["架构可扩展", "安全可靠", "性能优化"]
      
    - name: "核心商城功能"
      duration: "4周"
      deliverables: ["商品管理", "订单系统", "支付系统"]
      quality_gates: ["功能完整", "性能达标", "支付安全"]
      
    - name: "用户和营销功能"
      duration: "3周"
      deliverables: ["用户系统", "营销工具", "数据分析"]
      quality_gates: ["用户体验", "营销效果", "数据准确"]
      
    - name: "测试和优化"
      duration: "2周"
      deliverables: ["完整测试", "性能优化", "安全加固"]
      quality_gates: ["质量达标", "性能优化", "安全合规"]
      
  total_duration: "14周"
  team_size: "12-25人"
  quality_standard: "commercial_grade"
```

### 金融系统模板
```yaml
financial_system:
  phases:
    - name: "合规和需求分析"
      duration: "3周"
      deliverables: ["合规文档", "业务需求", "风险评估"]
      quality_gates: ["合规审查", "需求准确", "风险可控"]
      
    - name: "安全架构设计"
      duration: "2周"
      deliverables: ["安全架构", "加密方案", "审计设计"]
      quality_gates: ["安全标准", "加密强度", "审计完整"]
      
    - name: "核心交易系统"
      duration: "6周"
      deliverables: ["交易引擎", "风控系统", "清算系统"]
      quality_gates: ["交易准确", "风控有效", "清算正确"]
      
    - name: "监管和报表系统"
      duration: "3周"
      deliverables: ["监管报表", "审计日志", "合规检查"]
      quality_gates: ["报表准确", "日志完整", "合规达标"]
      
    - name: "安全测试和上线"
      duration: "2周"
      deliverables: ["安全测试", "渗透测试", "上线准备"]
      quality_gates: ["安全通过", "无漏洞", "准备充分"]
      
  total_duration: "16周"
  team_size: "15-30人"
  quality_standard: "financial_grade"
```

## 🔍 执行流程控制

### 阶段执行控制
```yaml
phase_execution_control:
  pre_phase:
    - "验证前置条件"
    - "检查资源准备"
    - "确认团队就绪"
    - "设置质量门禁"
    
  during_phase:
    - "监控执行进度"
    - "检查质量指标"
    - "管理风险问题"
    - "协调团队协作"
    
  post_phase:
    - "验证交付成果"
    - "执行质量检查"
    - "更新项目状态"
    - "准备下一阶段"
```

### 质量门禁控制
```yaml
quality_gate_control:
  gate_types:
    - "entry_gate": "阶段入口检查"
    - "progress_gate": "进行中检查"
    - "exit_gate": "阶段出口检查"
    
  check_items:
    - "code_quality": "代码质量检查"
    - "test_coverage": "测试覆盖率检查"
    - "documentation": "文档完整性检查"
    - "performance": "性能指标检查"
    - "security": "安全合规检查"
    
  failure_handling:
    - "block_execution": "阻止继续执行"
    - "require_fixes": "要求修复问题"
    - "escalate": "升级处理"
    - "record_issue": "记录问题"
```

### 进度管理控制
```yaml
progress_management:
  tracking_metrics:
    - "milestone_completion": "里程碑完成率"
    - "task_completion": "任务完成率"
    - "resource_utilization": "资源利用率"
    - "quality_score": "质量分数"
    
  schedule_control:
    - "baseline_schedule": "基准进度计划"
    - "actual_progress": "实际进度"
    - "variance_analysis": "偏差分析"
    - "recovery_plan": "恢复计划"
    
  reporting:
    - "daily_status": "每日状态报告"
    - "weekly_summary": "周度总结报告"
    - "milestone_review": "里程碑评审报告"
    - "quality_report": "质量报告"
```

## 📈 风险管理

### 风险识别
```yaml
risk_identification:
  technical_risks:
    - "技术选型风险"
    - "架构设计风险"
    - "性能瓶颈风险"
    - "安全漏洞风险"
    
  project_risks:
    - "进度延期风险"
    - "资源不足风险"
    - "需求变更风险"
    - "团队协作风险"
    
  quality_risks:
    - "质量标准风险"
    - "测试覆盖风险"
    - "文档缺失风险"
    - "合规风险"
```

### 风险评估
```yaml
risk_assessment:
  probability_levels:
    - "high": "发生概率 > 70%"
    - "medium": "发生概率 30-70%"
    - "low": "发生概率 < 30%"
    
  impact_levels:
    - "critical": "影响项目成功"
    - "major": "影响项目进度"
    - "minor": "影响项目质量"
    
  risk_matrix:
    high_critical: "立即处理"
    high_major: "优先处理"
    high_minor: "密切监控"
    medium_critical: "制定计划"
    medium_major: "定期检查"
    medium_minor: "记录跟踪"
    low_critical: "预防措施"
    low_major: "关注观察"
    low_minor: "接受风险"
```

### 风险应对
```yaml
risk_response:
  mitigation_strategies:
    - "avoidance": "避免风险"
    - "reduction": "降低风险"
    - "transfer": "转移风险"
    - "acceptance": "接受风险"
    
  contingency_planning:
    - "backup_plans": "备用计划"
    - "resource_reserves": "资源储备"
    - "time_buffers": "时间缓冲"
    - "alternative_solutions": "替代方案"
```

## 🎯 使用示例

### 企业级项目执行
```bash
@go "使用standard-mode-executor技能执行企业级CRM项目：
项目：大型CRM系统
模板：enterprise_web_application
团队：
  - po: 1人
  - analyst: 2人
  - architect: 2人
  - developer: 8人
  - qa: 2人
时间：13周
质量：企业级标准
约束：高可用、高性能、安全合规"

# 自动执行流程：
# 1. 选择enterprise_web_application模板
# 2. 生成13周详细执行计划
# 3. 设置5个阶段的质量门禁
# 4. 配置进度跟踪和风险管理
# 5. 开始标准化执行流程
```

### 电商平台项目执行
```bash
@go "使用standard-mode-executor技能执行电商平台项目：
项目：B2C电商平台
模板：ecommerce_platform
团队：
  - po: 1人
  - analyst: 2人
  - architect: 2人
  - developer: 10人
  - qa: 3人
  - ui_designer: 2人
时间：14周
质量：商业级标准
约束：高并发、支付安全、用户体验"

# 输出：
# - 14周详细执行计划
# - 6个阶段的质量检查点
# - 进度跟踪配置
# - 风险管理计划
# - 团队协作方案
```

## 📊 质量标准

### 模板质量
- **模板完整性**: 100%
- **流程标准化**: > 95%
- **质量门禁覆盖**: 100%
- **风险识别准确率**: > 90%

### 执行质量
- **进度准确性**: > 90%
- **质量达标率**: > 95%
- **风险控制有效性**: > 85%
- **团队协作效率**: > 80%

## 🚀 集成方式

### 与po-core-v2集成
```yaml
po_core_integration:
  execution_flow:
    1. "po-core-v2调用mode-selector选择standard模式"
    2. "po-core-v2调用standard-mode-executor执行项目"
    3. "standard-mode-executor基于模板执行标准化流程"
    4. "返回执行结果和质量报告给po-core-v2"
    
  data_exchange:
    - "项目请求信息"
    - "团队组成信息"
    - "执行计划结果"
    - "质量检查报告"
    - "进度跟踪信息"
```

### 与harness-integrator集成
```yaml
harness_integration:
  quality_assurance:
    - "harness-integrator提供HARNESS.md合规检查"
    - "standard-mode-executor在关键节点执行检查"
    - "确保所有输出符合项目约束标准"
    
  compliance_monitoring:
    - "代码质量合规"
    - "文档完整性合规"
    - "安全标准合规"
    - "流程规范合规"
```

### 与ralph-wiggum-loop集成
```yaml
ralph_wiggum_integration:
  continuous_improvement:
    - "ralph-wiggum-loop监控standard模式执行质量"
    - "分析执行效率和问题模式"
    - "提供模板优化建议"
    - "更新最佳实践"
    
  quality_tracking:
    - "跟踪质量分数趋势"
    - "识别质量改进机会"
    - "优化质量门禁设置"
    - "提升执行效率"
```

---

这个技能专门负责Standard模式的标准化执行，确保项目按照预定义模板高质量、高效率地完成，是PO System任务模式系统的核心组件。
