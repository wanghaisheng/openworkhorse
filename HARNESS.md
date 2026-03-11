# HARNESS.md - PicoClaw 项目核心约束

## 🎯 项目概述

本文档定义了 PicoClaw 项目的核心约束、质量标准和开发规范。所有项目参与者必须严格遵守这些规则，确保代码质量和项目成功。

## 🚫 禁止清单

### 代码规范禁止项

- ❌ **禁止全局变量**: 避免全局状态，使用依赖注入
- ❌ **禁止循环依赖**: 模块间不允许循环引用
- ❌ **禁止直接数据库操作**: 必须通过数据访问层
- ❌ **禁止硬编码配置**: 使用配置文件或环境变量
- ❌ **禁止忽略错误处理**: 所有错误必须被妥善处理
- ❌ **禁止提交调试代码**: 移除所有 console.log 和调试语句
- ❌ **禁止未经验证的第三方库**: 必须经过安全评估
- ❌ **禁止破坏向后兼容性**: 重大变更需要版本升级

### 架构设计禁止项

- ❌ **禁止单一职责违反**: 每个模块只负责一个功能
- ❌ **禁止紧耦合设计**: 模块间通过接口解耦
- ❌ **禁止重复代码**: 提取公共逻辑到共享模块
- ❌ **禁止过度设计**: 保持简洁实用的架构
- ❌ **禁止忽略性能考虑**: 设计阶段必须考虑性能影响

## ✅ 必须遵守

### 代码质量标准

- ✅ **测试覆盖率 ≥ 85%**: 所有核心功能必须有测试覆盖
- ✅ **严格分层架构**: 表现层、业务层、数据层分离
- ✅ **完整错误处理**: 所有可能的错误场景都要处理
- ✅ **代码审查通过**: 所有代码必须经过同行审查
- ✅ **文档完整性**: 公共API必须有完整文档
- ✅ **性能基准达标**: 满足性能要求才能发布
- ✅ **安全合规通过**: 通过安全扫描和合规检查

### 开发流程标准

- ✅ **需求明确**: 开始开发前需求必须清晰明确
- ✅ **设计评审**: 重要设计必须经过团队评审
- ✅ **测试驱动**: 关键功能采用TDD开发方式
- ✅ **持续集成**: 每次提交都要通过CI检查
- ✅ **版本控制**: 严格遵循Git工作流规范
- ✅ **发布流程**: 遵循标准发布流程和检查清单

## 🔄 Ralph Wiggum Loop

### 循环定义

Ralph Wiggum Loop 是一个自我迭代改进的闭环机制，确保代码质量和持续改进。

### 执行步骤

#### 步骤1: 运行所有检查
```bash
# 运行测试套件
go test ./... -v -race -coverprofile=coverage.out

# 运行代码检查
golangci-lint run ./...

# 运行类型检查
go vet ./...

# 运行安全扫描
gosec ./...

# 运行依赖检查
go mod tidy && go mod verify
```

#### 步骤2: 自我审查
根据 HARNESS.md 进行自我审查，列出所有违反项：

```markdown
## 自我审查清单

### HARNESS.md 合规性
- [x] 测试覆盖率达标 (85%+)
- [x] 代码检查通过
- [x] 类型检查通过
- [x] 安全检查通过
- [x] 依赖检查通过

### 代码质量
- [x] 错误处理完整
- [x] 文档完整
- [x] 性能考虑充分
- [x] 安全措施到位

### 架构设计
- [x] 分层架构正确
- [x] 模块职责单一
- [x] 接口设计合理
- [x] 依赖关系清晰

### 发现的违规项
1. **违规项1**: {具体描述}
   - **影响**: {影响评估}
   - **修复方案**: {修复方案}
   - **状态**: {待修复/已修复}

2. **违规项2**: {具体描述}
   - **影响**: {影响评估}
   - **修复方案**: {修复方案}
   - **状态**: {待修复/已修复}
```

#### 步骤3: 分析根因并修复
如果检查失败，分析根因并修复：

```bash
# 查看测试失败详情
go test ./... -v -run TestFailedFunction

# 查看代码检查问题
golangci-lint run ./... --verbose

# 分析性能问题
go test -bench=. -cpuprofile=cpu.prof
go tool pprof cpu.prof

# 修复问题
# ... 修复代码 ...

# 重新验证
go test ./... -v
golangci-lint run ./...
```

#### 步骤4: 创建PR
只有当所有检查通过且自我审查干净时才创建PR：

```markdown
## PR 检查清单

### 自动化检查
- [x] 所有测试通过 (100%)
- [x] 代码检查通过
- [x] 类型检查通过
- [x] 安全检查通过
- [x] 性能检查通过

### 自我审查
- [x] HARNESS.md 合规
- [x] 代码质量达标
- [x] 文档完整
- [x] 测试充分

### 人工审查
- [x] 代码审查通过
- [x] 设计评审通过
- [x] 安全评审通过

### 发布准备
- [x] 版本号更新
- [x] 变更日志更新
- [x] 文档更新
- [x] 发布说明准备
```

#### 步骤5: 更新 HARNESS.md/ADR
如果架构发生变化，更新相关文档：

```markdown
## 架构变更记录

### 变更描述
{变更的具体描述}

### 变更原因
{变更的原因和背景}

### 影响范围
- [ ] 代码结构
- [ ] API接口
- [ ] 数据模型
- [ ] 部署配置
- [ ] 文档

### 向后兼容性
- [ ] 向后兼容
- [ ] 破坏性变更 (需要版本升级)

### 更新文档
- [x] HARNESS.md
- [x] API文档
- [x] 架构文档
- [x] 部署文档
```

#### 步骤6: 建议改进
基于本次开发经验，提出2-3个改进建议：

```markdown
## 改进建议

### 建议1: {改进标题}
**问题描述**: {遇到的问题}
**建议方案**: {具体的改进方案}
**预期效果**: {改进后的效果}
**实施优先级**: {高/中/低}

### 建议2: {改进标题}
**问题描述**: {遇到的问题}
**建议方案**: {具体的改进方案}
**预期效果**: {改进后的效果}
**实施优先级**: {高/中/低}

### 建议3: {改进标题}
**问题描述**: {遇到的问题}
**建议方案**: {具体的改进方案}
**预期效果**: {改进后的效果}
**实施优先级**: {高/中/低}
```

## 📊 质量标准

### 代码质量指标

```yaml
code_quality_metrics:
  test_coverage:
    minimum: 85
    target: 90
    measurement: "go test -cover"
    
  code_complexity:
    maximum: 10
    measurement: "gocyclo -over 10"
    
  duplicate_code:
    maximum: 5
    measurement: "dupl"
    
  code_lines:
    maximum_function: 50
    maximum_file: 500
    measurement: "cloc"
```

### 性能标准

```yaml
performance_standards:
  api_response_time:
    p50: "< 100ms"
    p95: "< 200ms"
    p99: "< 500ms"
    
  database_query:
    simple: "< 10ms"
    complex: "< 100ms"
    
  memory_usage:
    idle: "< 100MB"
    peak: "< 1GB"
    
  cpu_usage:
    normal: "< 50%"
    peak: "< 80%"
```

### 安全标准

```yaml
security_standards:
  vulnerability_scan:
    critical: 0
    high: 0
    medium: "< 5"
    low: "< 10"
    
  dependency_check:
    known_vulnerabilities: 0
    outdated_packages: "< 5"
    
  code_security:
    sql_injection: 0
    xss: 0
    hard_coded_secrets: 0
```

## 🛠️ 开发工具配置

### Go工具链

```bash
# 安装必要的工具
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
go install github.com/gordonklaus/ineffassign@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
```

### golangci-lint 配置

```yaml
# .golangci.yml
run:
  timeout: 5m
  tests: true

linters:
  enable:
    - gofmt
    - goimports
    - govet
    - errcheck
    - staticcheck
    - unused
    - gosimple
    - structcheck
    - varcheck
    - ineffassign
    - deadcode
    - typecheck
    - gosec
    - misspell
    - unconvert
    - dupl
    - goconst
    - gocyclo

linters-settings:
  gocyclo:
    min-complexity: 10
  dupl:
    threshold: 100
  goconst:
    min-len: 3
    min-occurrences: 3

issues:
  exclude-use-default: false
  max-issues-per-linter: 0
  max-same-issues: 0
```

### Makefile

```makefile
# Makefile
.PHONY: test lint security clean

# 运行测试
test:
	go test ./... -v -race -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

# 代码检查
lint:
	golangci-lint run ./...

# 安全检查
security:
	gosec ./...

# 依赖检查
deps:
	go mod tidy
	go mod verify

# 所有检查
check: test lint security deps

# 代码格式化
fmt:
	go fmt ./...
	goimports -w .

# 清理
clean:
	rm -f coverage.out coverage.html
	go clean -testcache

# 构建
build:
	go build -o bin/picoclaw ./cmd/picoclaw

# Ralph Wiggum Loop
ralph: check
	@echo "✅ Ralph Wiggum Loop completed successfully!"
```

## 📋 检查清单

### 开发前检查

- [ ] 需求理解清晰
- [ ] 技术方案确定
- [ ] 开发环境准备
- [ ] 依赖库评估完成
- [ ] 安全考虑充分

### 开发中检查

- [ ] 代码规范遵循
- [ ] 错误处理完整
- [ ] 测试用例编写
- [ ] 文档同步更新
- [ ] 性能考虑充分

### 提交前检查

- [ ] 所有测试通过
- [ ] 代码检查通过
- [ ] 安全检查通过
- [ ] 自我审查完成
- [ ] 文档更新完成

### 发布前检查

- [ ] 代码审查通过
- [ ] 集成测试通过
- [ ] 性能测试通过
- [ ] 安全测试通过
- [ ] 用户验收通过

## 🚨 违规处理

### 违规等级

```yaml
violation_levels:
  critical:
    description: "严重违规，必须立即修复"
    examples:
      - "安全漏洞"
      - "数据丢失风险"
      - "系统崩溃"
    action: "立即停止相关操作，优先修复"
    
  high:
    description: "高级违规，需要尽快修复"
    examples:
      - "性能严重下降"
      - "功能异常"
      - "测试覆盖率不足"
    action: "24小时内修复"
    
  medium:
    description: "中级违规，需要计划修复"
    examples:
      - "代码质量问题"
      - "文档不完整"
      - "轻微性能问题"
    action: "3天内修复"
    
  low:
    description: "低级违规，可以后续修复"
    examples:
      - "代码风格问题"
      - "注释不规范"
      - "轻微优化机会"
    action: "1周内修复"
```

### 修复流程

```yaml
violation_fix_process:
  identification:
    - "自动化检查发现"
    - "代码审查发现"
    - "测试失败发现"
    - "用户反馈发现"
    
  prioritization:
    - "评估影响范围"
    - "确定修复优先级"
    - "分配修复责任人"
    - "设定修复时限"
    
  implementation:
    - "制定修复方案"
    - "实施代码修复"
    - "编写测试用例"
    - "验证修复效果"
    
  verification:
    - "运行所有检查"
    - "执行回归测试"
    - "进行性能验证"
    - "完成安全扫描"
    
  closure:
    - "更新相关文档"
    - "记录修复过程"
    - "总结经验教训"
    - "更新预防措施"
```

## 📚 参考资源

### 官方文档

- [Go官方代码规范](https://golang.org/doc/effective_go.html)
- [Go安全最佳实践](https://golang.org/doc/security)
- [Go性能优化指南](https://golang.org/doc/diagnostics)

### 工具文档

- [golangci-lint配置](https://golangci-lint.run/usage/configuration/)
- [gosec安全扫描](https://github.com/securecodewarrior/gosec)
- [Go测试覆盖率](https://blog.golang.org/cover)

### 最佳实践

- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [SOLID原则](https://en.wikipedia.org/wiki/SOLID)
- [测试驱动开发](https://en.wikipedia.org/wiki/Test-driven_development)

---

**重要提醒**: HARNESS.md 是项目的核心约束文档，任何修改都需要经过团队讨论和批准。所有参与者都有责任遵守这些规则，并在发现违规时及时报告。
