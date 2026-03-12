# 🛠️ 技能导入转换器设计

## 📋 概述

设计一个技能导入转换器，能够将 GitHub 仓库中的原始 Markdown 技能定义（如 agency-agents）转换为符合 PicoClaw 技能定义规范的标准化技能。

## 🎯 **设计目标**

### **核心功能**
1. **自动发现** - 扫描 GitHub 仓库中的技能文件
2. **智能解析** - 解析原始 Markdown 中的技能信息
3. **格式转换** - 转换为 PicoClaw 标准格式
4. **元数据生成** - 自动生成符合规范的元数据
5. **质量验证** - 验证转换后的技能质量

### **支持的源格式**
- **Agency Agents** - 基于 Front Matter 的 Markdown 格式
- **Vercel Agent Skills** - 标准化的技能格式
- **自定义格式** - 可扩展的格式支持

## 🔍 **源格式分析**

### **Agency Agents 格式**
```markdown
---
name: Frontend Developer
description: Expert frontend developer specializing in modern web technologies
color: cyan
emoji: 🖥️
vibe: Builds responsive, accessible web apps with pixel-perfect precision.
---

# Frontend Developer Agent Personality

You are **Frontend Developer**, an expert frontend developer...

## 🧠 Your Identity & Memory
- **Role**: Modern web application and UI implementation specialist
- **Personality**: Detail-oriented, performance-focused, user-centric, technically precise
- **Memory**: You remember successful UI patterns, performance optimization techniques...
- **Experience**: You've seen applications succeed through great UX...

## 🎯 Your Core Mission
### Editor Integration Engineering
- Build editor extensions with navigation commands...
```

### **PicoClaw 目标格式**
```markdown
---
name: role-frontend-developer
description: "前端开发专家，专注于现代Web技术、UI框架和性能优化"
category: "team-roles"
version: "1.0.0"
license: MIT

# 技能元数据
metadata:
  author: "Go Agents Team"
  created_at: "2026-03-12"
  updated_at: "2026-03-12"
  tags: ["frontend", "react", "vue", "angular", "ui"]
  organization: "Go Agents"
  source_repository: "agency-agents"
  original_file: "engineering-frontend-developer.md"

# 技能能力
capabilities:
  - "modern_web_development"
  - "ui_framework_implementation"
  - "performance_optimization"
  - "accessibility_compliance"
  - "responsive_design"

# 所需依赖
requires:
  skills: ["po-core"]
  tools: ["node", "npm", "git", "vscode"]

# 触发条件
triggers:
  - "前端开发"
  - "UI实现"
  - "React组件"
  - "性能优化"
  - "响应式设计"

# 技能变体
variants:
  react_specialist:
    persona: |
      我是 React 专家，专注于现代 React 开发。
      我擅长组件设计、状态管理、性能优化。
    
    capabilities: ["react_hooks", "context_api", "performance_optimization"]
    tools: ["react", "typescript", "vite"]
    
  vue_specialist:
    persona: |
      我是 Vue 专家，专注于 Vue 生态系统。
      我擅长组件开发、状态管理、构建优化。
    
    capabilities: ["vue_composition", "pinia", "nuxt"]
    tools: ["vue", "typescript", "vite"]

# 技能执行配置
execution:
  task_breakdown:
    strategy: "template_based"
    granularity: "medium"
    max_depth: 3
    auto_adjust: true
  
  subtask_execution:
    sequential: true
    parallel_threshold: 4
    quality_gates:
      - gate: "code_quality"
        threshold: 85
        reviewer: "self"
      - gate: "accessibility"
        threshold: 90
        reviewer: "self"
  
  ralph_wiggum_loop:
    enabled: true
    steps:
      - "code_implementation"
      - "performance_optimization"
      - "accessibility_check"
      - "testing_validation"
    
    auto_update:
      wbs: true
      harness: false

# 输出物定义
outputs:
  supported_types:
    - "code"
    - "document"
    - "prototype"
  
  quality_standards:
    code:
      format: "typescript"
      coverage: ">=85%"
      style: "prettier"
    document:
      format: "markdown"
      structure: "standard"
      validation: "markdown_lint"

# 质量门禁定义
quality_gates:
  task_level:
    - gate: "code_quality"
      threshold: 85
      reviewer: "self"
    - gate: "performance"
      threshold: 80
      reviewer: "self"
    - gate: "accessibility"
      threshold: 90
      reviewer: "self"
  
  milestone_level:
    - gate: "user_experience"
      threshold: 85
      reviewer: "phase_lead"

# 使用示例
examples:
  - title: "React 组件开发"
    description: "开发高性能的 React 组件"
    command: "@go \"开发一个响应式的数据表格组件\""
    expected_output: "完整的 React 组件代码，包含类型定义、测试和文档"

# 如何使用
usage:
  when_to_apply:
    - "需要开发现代 Web 应用"
    - "需要实现复杂的 UI 组件"
    - "需要优化前端性能"
    - "需要确保可访问性"
  
  command_examples:
    - command: "@go \"开发一个响应式的用户界面\""
      description: "开发响应式用户界面"
      expected: "完整的 UI 组件和样式"

# 输出格式
output_format:
  success: "✅ 前端开发任务完成，代码质量优秀，性能优化到位"
  error: "❌ 开发过程中遇到问题，需要进一步优化"
  progress: "🔄 正在开发中，已完成 {progress}%"

# 故障排除
troubleshooting:
  common_issues:
    - issue: "性能问题"
      solution: "使用 React.memo、useMemo、useCallback 优化性能"
    - issue: "可访问性问题"
      solution: "添加 ARIA 标签，确保键盘导航"
    - issue: "浏览器兼容性"
      solution: "使用 polyfills 和渐进增强策略"
```

## 🏗️ **转换器架构**

### **整体架构**
```
┌─────────────────────────────────────────────────────────────┐
│                    技能导入转换器                              │
├─────────────────────────────────────────────────────────────┤
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐ │
│  │   源格式检测器   │  │   内容解析器     │  │   格式转换器     │ │
│  │ FormatDetector  │  │ ContentParser   │  │ FormatConverter │ │
│  └─────────────────┘  └─────────────────┘  └─────────────────┘ │
├─────────────────────────────────────────────────────────────┤
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐ │
│  │   元数据生成器   │  │   质量验证器     │  │   技能生成器     │ │
│  │ MetadataGen     │  │ QualityValidator │  │ SkillGenerator  │ │
│  └─────────────────┘  └─────────────────┘  └─────────────────┘ │
├─────────────────────────────────────────────────────────────┤
│                    输出管理层                                 │
│  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐ │
│  │   文件写入器     │  │   索引更新器     │  │   缓存管理器     │ │
│  │ FileWriter      │  │ IndexUpdater    │  │ CacheManager    │ │
│  └─────────────────┘  └─────────────────┘  └─────────────────┘ │
└─────────────────────────────────────────────────────────────┘
```

## 🔧 **核心组件设计**

### **1. 源格式检测器 (FormatDetector)**

#### **功能**
- 检测技能文件的格式类型
- 识别 Front Matter 结构
- 确定转换策略

#### **接口设计**
```go
type FormatDetector interface {
    DetectFormat(filePath string) (Format, error)
    GetSupportedFormats() []Format
    ValidateFormat(filePath string) error
}

type Format string

const (
    FormatAgencyAgents Format = "agency-agents"
    FormatVercelSkills Format = "vercel-skills"
    FormatCustom      Format = "custom"
)

type FormatInfo struct {
    Type        Format `json:"type"`
    Version     string `json:"version"`
    FrontMatter bool   `json:"front_matter"`
    Sections   []string `json:"sections"`
}
```

#### **实现示例**
```go
func (d *AgencyFormatDetector) DetectFormat(filePath string) (Format, error) {
    content, err := ioutil.ReadFile(filePath)
    if err != nil {
        return "", err
    }
    
    // 检查 Front Matter
    if strings.HasPrefix(string(content), "---\n") {
        // 检查是否包含 agency-agents 特定字段
        if strings.Contains(string(content), "name:") && 
           strings.Contains(string(content), "description:") &&
           strings.Contains(string(content), "color:") {
            return FormatAgencyAgents, nil
        }
    }
    
    return FormatUnknown, errors.New("unsupported format")
}
```

### **2. 内容解析器 (ContentParser)**

#### **功能**
- 解析 Front Matter
- 提取技能信息
- 识别关键章节

#### **接口设计**
```go
type ContentParser interface {
    ParseFrontMatter(content string) (map[string]interface{}, error)
    ParseSections(content string) (map[string]string, error)
    ExtractCapabilities(content string) ([]string, error)
    ExtractTriggers(content string) ([]string, error)
}

type ParsedContent struct {
    FrontMatter map[string]interface{} `json:"front_matter"`
    Sections    map[string]string       `json:"sections"`
    Metadata    SkillMetadata          `json:"metadata"`
}

type SkillMetadata struct {
    Name         string   `json:"name"`
    Description  string   `json:"description"`
    Category     string   `json:"category"`
    Capabilities []string `json:"capabilities"`
    Triggers     []string `json:"triggers"`
    Tools        []string `json:"tools"`
    Personality  string   `json:"personality"`
    Mission      string   `json:"mission"`
}
```

#### **实现示例**
```go
func (p *AgencyParser) ParseFrontMatter(content string) (map[string]interface{}, error) {
    parts := strings.SplitN(content, "---\n", 3)
    if len(parts) < 3 {
        return nil, errors.New("invalid front matter format")
    }
    
    fmContent := parts[1]
    var frontMatter map[string]interface{}
    
    err := yaml.Unmarshal([]byte(fmContent), &frontMatter)
    if err != nil {
        return nil, err
    }
    
    return frontMatter, nil
}

func (p *AgencyParser) ParseSections(content string) (map[string]string, error) {
    sections := make(map[string]string)
    
    // 使用正则表达式识别章节
    sectionRegex := regexp.MustCompile(`^#{1,3}\s+(.+)$\n([\s\S]*?)(?=\n#{1,3}\s+|$)`)
    matches := sectionRegex.FindAllStringSubmatch(content, -1)
    
    for _, match := range matches {
        if len(match) >= 3 {
            title := strings.TrimSpace(match[1])
            content := strings.TrimSpace(match[2])
            sections[title] = content
        }
    }
    
    return sections, nil
}
```

### **3. 格式转换器 (FormatConverter)**

#### **功能**
- 将源格式转换为目标格式
- 生成标准化的技能定义
- 处理变体和配置

#### **接口设计**
```go
type FormatConverter interface {
    ConvertToSkill(parsed ParsedContent) (*Skill, error)
    GenerateVariants(skill *Skill) ([]Variant, error)
    ConvertExecutionConfig(parsed ParsedContent) (*ExecutionConfig, error)
}

type ConversionConfig struct {
    TargetCategory    string            `json:"target_category"`
    DefaultVariant    string            `json:"default_variant"`
    QualityThreshold  int               `json:"quality_threshold"`
    CustomMappings    map[string]string `json:"custom_mappings"`
}
```

#### **实现示例**
```go
func (c *AgencyToPicoClawConverter) ConvertToSkill(parsed ParsedContent) (*Skill, error) {
    skill := &Skill{
        ID:          generateSkillID(parsed.Metadata.Name),
        Name:        parsed.Metadata.Name,
        Description: parsed.Metadata.Description,
        Category:    c.config.TargetCategory,
        Version:     "1.0.0",
        License:     "MIT",
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }
    
    // 转换能力
    skill.Capabilities = convertCapabilities(parsed.Metadata.Capabilities)
    
    // 转换触发条件
    skill.Triggers = convertTriggers(parsed.Metadata.Triggers)
    
    // 生成元数据
    skill.Metadata = generateMetadata(parsed)
    
    return skill, nil
}

func (c *AgencyToPicoClawConverter) generateMetadata(parsed ParsedContent) map[string]interface{} {
    return map[string]interface{}{
        "author":            "Go Agents Team",
        "source_repository": "agency-agents",
        "original_file":     parsed.Metadata.Name,
        "original_format":   "agency-agents",
        "conversion_date":   time.Now(),
        "personality":       parsed.Metadata.Personality,
        "mission":          parsed.Metadata.Mission,
    }
}
```

### **4. 元数据生成器 (MetadataGenerator)**

#### **功能**
- 生成符合规范的元数据
- 创建 metadata.json 文件
- 处理版本和依赖信息

#### **接口设计**
```go
type MetadataGenerator interface {
    GenerateMetadataJSON(skill *Skill) (*SkillMetadata, error)
    GenerateVariants(skill *Skill) ([]Variant, error)
    GenerateDependencies(skill *Skill) (*Dependencies, error)
}

type SkillMetadata struct {
    Version          string            `json:"version"`
    Organization     string            `json:"organization"`
    Date             string            `json:"date"`
    Abstract         string            `json:"abstract"`
    References       []string          `json:"references"`
    Tags             []string          `json:"tags"`
    Category         string            `json:"category"`
    Complexity       string            `json:"complexity"`
    EstimatedTime    string            `json:"estimated_time"`
    Dependencies     []string          `json:"dependencies"`
    Compatibility    CompatibilityInfo `json:"compatibility"`
    SourceInfo       SourceInfo        `json:"source_info"`
}

type SourceInfo struct {
    Repository       string `json:"repository"`
    OriginalFile     string `json:"original_file"`
    OriginalFormat   string `json:"original_format"`
    ConversionDate   string `json:"conversion_date"`
    ConversionTool   string `json:"conversion_tool"`
}
```

### **5. 质量验证器 (QualityValidator)**

#### **功能**
- 验证转换后的技能质量
- 检查格式规范
- 评估完整性

#### **接口设计**
```go
type QualityValidator interface {
    ValidateSkill(skill *Skill) (*ValidationResult, error)
    ValidateMetadata(metadata *SkillMetadata) error
    CalculateQualityScore(skill *Skill) (int, error)
}

type ValidationResult struct {
    Valid    bool     `json:"valid"`
    Errors   []string `json:"errors"`
    Warnings []string `json:"warnings"`
    Score    int      `json:"score"`
    Issues   []Issue  `json:"issues"`
}

type Issue struct {
    Type        string `json:"type"`        // error, warning, info
    Category    string `json:"category"`    // format, content, structure
    Message     string `json:"message"`
    Suggestion  string `json:"suggestion"`
    Severity    int    `json:"severity"`    // 1-10
}
```

## 📋 **转换流程**

### **1. 发现阶段**
```bash
# 扫描仓库
picoclaw skills import --repo agency-agents --scan

# 输出：
# 发现 45 个技能文件
# 格式：agency-agents
# 支持转换：42 个
```

### **2. 解析阶段**
```bash
# 解析技能
picoclaw skills import --repo agency-agents --parse engineering-frontend-developer.md

# 输出：
# ✅ 解析成功
# 名称：Frontend Developer
# 能力：12 个
# 触发条件：8 个
```

### **3. 转换阶段**
```bash
# 转换技能
picoclaw skills import --repo agency-agents --convert engineering-frontend-developer.md

# 输出：
# ✅ 转换成功
# 目标技能：role-frontend-developer
# 生成变体：3 个
# 质量分数：85
```

### **4. 验证阶段**
```bash
# 验证技能
picoclaw skills import --repo agency-agents --validate role-frontend-developer

# 输出：
# ✅ 验证通过
# 格式规范：100%
# 内容完整：95%
# 质量分数：85
```

### **5. 安装阶段**
```bash
# 安装技能
picoclaw skills import --repo agency-agents --install role-frontend-developer

# 输出：
# ✅ 安装成功
# 技能路径：~/.picoclaw/workspace/skills/team-roles/role-frontend-developer
# 可用命令：@go "前端开发任务"
```

## 🛠️ **命令行接口设计**

### **基础命令**
```bash
# 仓库扫描
picoclaw skills import --repo <repo-url> --scan
picoclaw skills import --repo <repo-url> --format agency-agents

# 技能转换
picoclaw skills import --repo <repo-url> --convert <skill-file>
picoclaw skills import --repo <repo-url> --convert-all

# 批量操作
picoclaw skills import --repo <repo-url> --batch --category team-roles
picoclaw skills import --repo <repo-url> --batch --quality-threshold 80

# 质量控制
picoclaw skills import --repo <repo-url> --validate <skill-name>
picoclaw skills import --repo <repo-url> --quality-report

# 安装管理
picoclaw skills import --repo <repo-url> --install <skill-name>
picoclaw skills import --repo <repo-url> --install-all
```

### **高级选项**
```bash
# 自定义配置
picoclaw skills import --repo <repo-url> --config custom-config.yaml
picoclaw skills import --repo <repo-url> --mapping field-mapping.json

# 输出控制
picoclaw skills import --repo <repo-url> --output-dir ./converted-skills
picoclaw skills import --repo <repo-url> --dry-run --verbose

# 并行处理
picoclaw skills import --repo <repo-url> --parallel 4
picoclaw skills import --repo <repo-url> --timeout 300
```

## 📊 **配置文件设计**

### **转换配置**
```yaml
# import-config.yaml
import:
  repositories:
    agency-agents:
      url: "https://github.com/msitarzewski/agency-agents"
      format: "agency-agents"
      branch: "main"
      target_category: "team-roles"
      quality_threshold: 80
  
  conversion:
    default_variant: "general"
    generate_variants: true
    include_examples: true
    preserve_original: true
  
  validation:
    strict_mode: false
    require_documentation: true
    require_examples: false
  
  output:
    directory: "./converted-skills"
    create_metadata: true
    generate_index: true
    compress_output: false
```

### **字段映射配置**
```yaml
# field-mapping.yaml
mappings:
  agency-agents:
    front_matter:
      name: "name"
      description: "description"
      category: "category"
    
    sections:
      "Your Core Mission": "mission"
      "Your Identity & Memory": "identity"
      "Critical Rules": "rules"
      "Technical Deliverables": "deliverables"
    
    capabilities:
      from_section: "Your Core Mission"
      pattern: "### (.+)"
    
    triggers:
      from_section: "Your Core Mission"
      pattern: "\\*\\* (.+) \\*\\*"
```

## 🚀 **实现计划**

### **Phase 1: 基础转换器 (Week 1-2)**
- [ ] 实现格式检测器
- [ ] 实现内容解析器
- [ ] 实现基础格式转换器
- [ ] 添加 Agency Agents 支持

### **Phase 2: 完善转换功能 (Week 3-4)**
- [ ] 实现元数据生成器
- [ ] 实现质量验证器
- [ ] 添加变体生成功能
- [ ] 实现批量转换

### **Phase 3: 命令行工具 (Week 5-6)**
- [ ] 实现 CLI 命令
- [ ] 添加配置文件支持
- [ ] 实现并行处理
- [ ] 添加进度显示

### **Phase 4: 扩展和优化 (Week 7-8)**
- [ ] 支持更多源格式
- [ ] 优化转换质量
- [ ] 添加转换报告
- [ ] 实现增量更新

## 🎉 **总结**

通过这个技能导入转换器，我们可以：

1. **自动化导入** - 自动发现和转换外部技能
2. **格式标准化** - 统一技能定义格式
3. **质量保证** - 确保转换后的技能质量
4. **批量处理** - 支持大规模技能导入
5. **可扩展性** - 支持多种源格式
6. **版本管理** - 跟踪转换历史和来源

这个转换器将大大丰富 Go Agents 的技能生态系统，让用户能够轻松使用来自各种来源的技能！🚀
