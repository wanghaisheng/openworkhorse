---
name: hybrid-mode
description: "混合模式 - 结合标准化和自由模式的优点，平衡效率与灵活性。适合部分标准化部分探索的项目，提供最优的执行策略。"
---

# Hybrid Mode Skill

## 模式概述

Hybrid Mode 结合了 Standard Mode 的标准化优势和 Free Mode 的灵活性，适合：

- **混合项目**: 部分需求明确，部分需要探索
- **渐进式创新**: 在成熟基础上进行创新
- **技术迁移**: 从旧技术向新技术过渡
- **复杂系统**: 包含多个不同特性的子系统

## 核心理念

### 分层执行策略

```yaml
execution_layers:
  standard_layer:
    description: "标准化执行层"
   适用场景: "需求明确、技术成熟的部分"
    执行方式: "基于模板的标准化流程"
    
  flexible_layer:
    description: "灵活执行层"
    适用场景: "需求不明确、需要探索的部分"
    执行方式: "Phase Lead 主导的动态规划"
    
  integration_layer:
    description: "集成协调层"
    适用场景: "两层之间的协调和集成"
    执行方式: "混合协调机制"
```

### 智能模式选择

```typescript
interface HybridModeSelector {
  selectExecutionMode(
    task: Task,
    projectContext: ProjectContext
  ): ExecutionMode {
    
    // 1. 分析任务特征
    const characteristics = this.analyzeTaskCharacteristics(task);
    
    // 2. 评估需求明确度
    const clarityScore = this.assessRequirementClarity(task);
    
    // 3. 评估技术成熟度
    const maturityScore = this.assessTechnicalMaturity(task);
    
    // 4. 评估风险等级
    const riskLevel = this.assessRiskLevel(task);
    
    // 5. 计算模式选择分数
    const standardScore = this.calculateStandardScore(
      clarityScore, maturityScore, riskLevel
    );
    const flexibleScore = this.calculateFlexibleScore(
      clarityScore, maturityScore, riskLevel
    );
    
    // 6. 选择执行模式
    if (standardScore > flexibleScore + threshold) {
      return ExecutionMode.STANDARD;
    } else if (flexibleScore > standardScore + threshold) {
      return ExecutionMode.FLEXIBLE;
    } else {
      return ExecutionMode.HYBRID;
    }
  }
}
```

## 混合执行策略

### 任务分类机制

```yaml
task_classification:
  standard_tasks:
    criteria:
      - "需求明确度 ≥ 80%"
      - "技术成熟度 ≥ 85%"
      - "风险等级 ≤ 3"
      - "有成熟的最佳实践"
    
    examples:
      - "用户认证功能"
      - "数据库CRUD操作"
      - "标准API接口"
      - "UI组件开发"
      
  flexible_tasks:
    criteria:
      - "需求明确度 ≤ 60%"
      - "技术成熟度 ≤ 70%"
      - "风险等级 ≥ 6"
      - "需要技术探索"
    
    examples:
      - "AI算法选择"
      - "新架构设计"
      - "性能优化方案"
      - "创新功能设计"
      
  hybrid_tasks:
    criteria:
      - "60% < 需求明确度 < 80%"
      - "70% < 技术成熟度 < 85%"
      - "3 < 风险等级 < 6"
      - "部分已知部分未知"
    
    examples:
      - "系统集成方案"
      - "技术迁移项目"
      - "功能增强项目"
      - "性能改进项目"
```

### 分层执行流程

```typescript
interface HybridExecutionEngine {
  async executeHybridProject(
    project: Project,
    executionPlan: HybridExecutionPlan
  ): Promise<ProjectResult> {
    
    // 1. 任务分类
    const classifiedTasks = await this.classifyTasks(project.tasks);
    
    // 2. 创建执行层
    const standardLayer = new StandardExecutionLayer(classifiedTasks.standard);
    const flexibleLayer = new FlexibleExecutionLayer(classifiedTasks.flexible);
    const integrationLayer = new IntegrationLayer(
      classifiedTasks.hybrid,
      standardLayer,
      flexibleLayer
    );
    
    // 3. 并行执行
    const results = await Promise.all([
      standardLayer.execute(),
      flexibleLayer.execute(),
      integrationLayer.coordinate()
    ]);
    
    // 4. 集成结果
    return await this.integrateResults(results);
  }
}
```

## 标准化执行层

### 模板化执行

```yaml
standard_execution:
  template_library:
    - "user-auth-standard"
    - "data-crud-standard"
    - "api-interface-standard"
    - "ui-component-standard"
    
  quality_gates:
    mandatory: true
    auto_check: true
    threshold: 85
    
  reporting:
    frequency: "daily"
    format: "standard"
    metrics: ["进度", "质量", "资源使用"]
```

### 标准任务示例

```yaml
standard_task_example:
  name: "用户认证功能开发"
  classification: "standard"
  reason: "需求明确，技术成熟，有标准流程"
  
  execution_plan:
    template: "user-auth-standard"
    estimated_time: "3-5天"
    team: ["backend_dev", "frontend_dev", "qa"]
    
  quality_gates:
    - gate: "security_compliance"
      threshold: 100
      auto_check: true
    - gate: "test_coverage"
      threshold: 90
      auto_check: true
    - gate: "performance_benchmark"
      threshold: 85
      auto_check: true
      
  deliverables:
    - "认证API/"
    - "登录页面/"
    - "安全测试报告.md"
    - "性能测试报告.md"
```

## 灵活执行层

### 动态规划

```yaml
flexible_execution:
  planning_approach:
    - "里程碑驱动"
    - "迭代式开发"
    - "快速验证"
    - "持续调整"
    
  team_structure:
    phase_lead: "技术专家"
    team_members: ["高级开发", "架构师", "研究员"]
    
  decision_making:
    approach: "团队共识"
    threshold: 0.7
    conflict_resolution: "专家仲裁"
```

### 灵活任务示例

```yaml
flexible_task_example:
  name: "AI推荐算法探索"
  classification: "flexible"
  reason: "技术新颖，需要探索多种方案"
  
  execution_plan:
    approach: "动态规划"
    estimated_time: "2-3周"
    team: ["ai_researcher", "data_scientist", "backend_dev"]
    
  milestones:
    - milestone: "算法调研"
      tasks:
        - "调研现有算法"
        - "评估适用性"
        - "选择候选方案"
        
    - milestone: "原型验证"
      tasks:
        - "实现算法原型"
        - "数据验证"
        - "性能评估"
        
    - milestone: "方案确定"
      tasks:
        - "对比分析"
        - "最终选择"
        - "集成方案"
        
  adaptation_rules:
    - "如果算法效果不佳，立即调整方案"
    - "如果发现新的算法，重新评估"
    - "如果数据不足，调整数据收集策略"
```

## 集成协调层

### 协调机制

```typescript
interface IntegrationCoordinator {
  coordinateLayers(
    standardLayer: StandardExecutionLayer,
    flexibleLayer: FlexibleExecutionLayer,
    hybridTasks: HybridTask[]
  ): Promise<CoordinationResult> {
    
    // 1. 识别依赖关系
    const dependencies = this.identifyDependencies(
      standardLayer.tasks,
      flexibleLayer.tasks,
      hybridTasks
    );
    
    // 2. 同步执行计划
    const synchronizedPlan = await this.synchronizePlans(
      standardLayer.plan,
      flexibleLayer.plan,
      dependencies
    );
    
    // 3. 协调资源分配
    const resourceAllocation = await this.coordinateResources(
      standardLayer.resources,
      flexibleLayer.resources,
      hybridTasks
    );
    
    // 4. 管理接口定义
    const interfaces = await this.defineInterfaces(
      standardLayer.outputs,
      flexibleLayer.inputs
    );
    
    // 5. 监控集成进度
    const integrationProgress = await this.monitorIntegration(
      synchronizedPlan,
      interfaces
    );
    
    return {
      synchronizedPlan,
      resourceAllocation,
      interfaces,
      integrationProgress
    };
  }
}
```

### 接口管理

```yaml
interface_management:
  standard_to_flexible:
    description: "标准化层向灵活层提供接口"
    interfaces:
      - name: "user_data_api"
        type: "REST API"
        specification: "api-spec.yaml"
        version: "v1.0"
        
      - name: "business_logic_service"
        type: "gRPC"
        specification: "proto-files/"
        version: "v2.1"
        
  flexible_to_standard:
    description: "灵活层向标准化层提供接口"
    interfaces:
      - name: "ai_recommendation_service"
        type: "REST API"
        specification: "ai-api-spec.yaml"
        version: "v0.1"
        
      - name: "analytics_data_provider"
        type: "Message Queue"
        specification: "message-schema.json"
        version: "v1.0"
```

## 混合任务处理

### 任务分解策略

```typescript
interface HybridTaskDecomposer {
  decomposeHybridTask(
    task: HybridTask
  ): Promise<DecompositionResult> {
    
    // 1. 识别标准化部分
    const standardParts = this.identifyStandardParts(task);
    
    // 2. 识别灵活部分
    const flexibleParts = this.identifyFlexibleParts(task);
    
    // 3. 定义集成点
    const integrationPoints = this.defineIntegrationPoints(
      standardParts,
      flexibleParts
    );
    
    // 4. 创建子任务
    const subtasks = await this.createSubtasks(
      standardParts,
      flexibleParts,
      integrationPoints
    );
    
    // 5. 建立依赖关系
    const dependencies = this.establishDependencies(subtasks);
    
    return {
      standardSubtasks: subtasks.filter(t => t.type === 'standard'),
      flexibleSubtasks: subtasks.filter(t => t.type === 'flexible'),
      integrationSubtasks: subtasks.filter(t => t.type === 'integration'),
      dependencies
    };
  }
}
```

### 混合任务示例

```yaml
hybrid_task_example:
  name: "电商平台升级项目"
  classification: "hybrid"
  reason: "包含标准化功能升级和创新功能添加"
  
  decomposition:
    standard_parts:
      - "用户系统升级（标准化）"
      - "订单处理优化（标准化）"
      - "支付接口更新（标准化）"
      
    flexible_parts:
      - "AI推荐系统（探索）"
      - "个性化营销策略（创新）"
      - "实时数据分析（新技术）"
      
    integration_points:
      - "用户数据与推荐系统集成"
      - "订单数据与AI分析集成"
      - "支付数据与个性化营销集成"
      
  execution_plan:
    phase_1:
      name: "标准化升级"
      duration: "2周"
      tasks:
        - "用户系统API升级"
        - "订单处理性能优化"
        - "支付接口标准化"
      execution_mode: "standard"
      
    phase_2:
      name: "创新功能开发"
      duration: "3周"
      tasks:
        - "AI推荐算法开发"
        - "个性化营销策略设计"
        - "实时数据分析平台"
      execution_mode: "flexible"
      
    phase_3:
      name: "系统集成测试"
      duration: "1周"
      tasks:
        - "接口集成测试"
        - "端到端功能测试"
        - "性能集成测试"
      execution_mode: "integration"
```

## 质量保证

### 分层质量门禁

```yaml
layered_quality_gates:
  standard_layer:
    gates:
      - gate: "template_compliance"
        threshold: 100
        auto_check: true
      - gate: "code_quality"
        threshold: 85
        auto_check: true
      - gate: "test_coverage"
        threshold: 90
        auto_check: true
        
  flexible_layer:
    gates:
      - gate: "innovation_score"
        threshold: 70
        auto_check: false
      - gate: "validation_results"
        threshold: 80
        auto_check: false
      - gate: "team_consensus"
        threshold: 0.7
        auto_check: false
        
  integration_layer:
    gates:
      - gate: "interface_compatibility"
        threshold: 100
        auto_check: true
      - gate: "end_to_end_testing"
        threshold: 90
        auto_check: true
      - gate: "performance_integration"
        threshold: 85
        auto_check: true
```

### 跨层协调检查

```typescript
interface CrossLayerValidator {
  validateCrossLayerIntegration(
    standardLayer: StandardExecutionResult,
    flexibleLayer: FlexibleExecutionResult,
    integrationLayer: IntegrationResult
  ): Promise<ValidationResult> {
    
    // 1. 接口兼容性检查
    const interfaceCompatibility = await this.checkInterfaceCompatibility(
      standardLayer.interfaces,
      flexibleLayer.interfaces
    );
    
    // 2. 数据一致性检查
    const dataConsistency = await this.checkDataConsistency(
      standardLayer.data,
      flexibleLayer.data
    );
    
    // 3. 性能集成检查
    const performanceIntegration = await this.checkPerformanceIntegration(
      standardLayer.performance,
      flexibleLayer.performance
    );
    
    // 4. 安全集成检查
    const securityIntegration = await this.checkSecurityIntegration(
      standardLayer.security,
      flexibleLayer.security
    );
    
    return {
      interfaceCompatibility,
      dataConsistency,
      performanceIntegration,
      securityIntegration,
      overallScore: this.calculateOverallScore([
        interfaceCompatibility,
        dataConsistency,
        performanceIntegration,
        securityIntegration
      ])
    };
  }
}
```

## 监控和报告

### 分层监控

```bash
# 查看各层执行状态
hybrid-mode status --project "ecommerce-upgrade" --by-layer

# 查看集成进度
hybrid-mode integration-status --project "ecommerce-upgrade"

# 查看跨层依赖
hybrid-mode dependencies --project "ecommerce-upgrade" --cross-layer
```

### 混合报告

```bash
# 生成混合执行报告
hybrid-mode report --project "ecommerce-upgrade" --format markdown

# 生成层间对比报告
hybrid-mode comparison-report --project "ecommerce-upgrade" --layers "standard,flexible"

# 生成集成质量报告
hybrid-mode integration-report --project "ecommerce-upgrade" --detailed
```

### 关键指标

- **层间协调效率**: 标准层和灵活层的协调效果
- **集成成功率**: 跨层集成的成功率和质量
- **混合执行效率**: 混合模式相对于纯模式的效率提升
- **适应性评分**: 项目应对变化的能力

## 配置管理

### 全局配置

```yaml
# .goagents/config/hybrid-mode.yaml
hybrid_mode:
  layer_distribution:
    standard_threshold: 0.7
    flexible_threshold: 0.3
    integration_overhead: 0.1
    
  coordination_rules:
    sync_frequency: "daily"
    conflict_resolution: "expert_arbitration"
    interface_versioning: "semantic"
    
  quality_gates:
    cross_layer_validation: true
    integration_testing: mandatory
    performance_monitoring: continuous
```

### 项目配置

```yaml
# .goagents/projects/ecommerce-upgrade/hybrid-mode.yaml
project_config:
  layer_allocation:
    standard_layer: 60
    flexible_layer: 30
    integration_layer: 10
    
  custom_interfaces:
    - name: "user_recommendation_bridge"
      source: "standard_user_service"
      target: "flexible_ai_service"
      protocol: "REST"
      
  adaptation_rules:
    - rule: "标准化部分延迟时，灵活部分可独立推进"
      condition: "standard_layer_delay > 2天"
      action: "enable_flexible_independent_execution"
```

## 集成接口

### 与 Standard Mode 集成

```typescript
interface StandardModeIntegration {
  executeStandardTasks(
    tasks: StandardTask[],
    configuration: StandardConfig
  ): Promise<StandardResult>;
  
  getStandardProgress(
    projectId: string
  ): Promise<StandardProgress>;
}
```

### 与 Free Mode 集成

```typescript
interface FreeModeIntegration {
  executeFlexibleTasks(
    tasks: FlexibleTask[],
    phaseLead: Agent
  ): Promise<FlexibleResult>;
  
  getFlexibleProgress(
    projectId: string
  ): Promise<FlexibleProgress>;
}
```

### 集成协调接口

```typescript
interface HybridCoordination {
  coordinateExecution(
    standardPlan: StandardPlan,
    flexiblePlan: FlexiblePlan
  ): Promise<CoordinatedPlan>;
  
  manageInterfaces(
    interfaces: InterfaceDefinition[]
  ): Promise<InterfaceStatus>;
  
  validateIntegration(
    results: ExecutionResult[]
  ): Promise<ValidationResult>;
}
```

## 扩展性

### 自定义分层策略

支持自定义分层策略：

1. 定义层次结构
2. 设置分配规则
3. 配置协调机制
4. 注册到系统

### 自定义集成模式

支持添加自定义集成模式：

1. 定义集成模式
2. 设置集成规则
3. 配置接口标准
4. 集成到系统

### 自定义适应性机制

支持添加自定义适应性机制：

1. 定义适应策略
2. 设置触发条件
3. 配置响应动作
4. 注册到系统
