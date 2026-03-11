# 🛠️ Go Agents 故障排除指南

## 🎯 概述

本指南提供 Go Agents 使用过程中常见问题的诊断和解决方案，帮助用户快速定位和解决问题，确保系统稳定运行。

## 📋 问题分类

### 🔧 CLI 问题
### 🎯 技能问题
### ⚙️ 配置问题
### 🚀 执行问题
### 📊 性能问题
### 🔒 安全问题

## 🔧 CLI 问题

### 问题1: `picoclaw goagents` 命令不存在
```bash
# 错误信息
bash: picoclaw goagents: command not found

# 诊断步骤
1. 检查 PicoClaw 版本
picoclaw version

2. 检查 goagents 集成
picoclaw --help | grep goagents

# 解决方案
# 1. 重新编译安装
cd cmd/picoclaw
go build -o picoclaw .
sudo mv picoclaw /usr/local/bin/

# 2. 检查 PATH 环境变量
echo $PATH | grep picoclaw

# 3. 验证安装
picoclaw goagents --help
```

### 问题2: goagents 命令权限不足
```bash
# 错误信息
Error: permission denied while accessing goagents configuration

# 诊断步骤
1. 检查工作区权限
ls -la ~/.picoclaw/workspace/

2. 检查配置文件权限
ls -la ~/.picoclaw/workspace/.goagents/

# 解决方案
# 1. 修复权限
chmod -R 755 ~/.picoclaw/workspace/
chmod -R 644 ~/.picoclaw/workspace/.goagents/config.yaml

# 2. 重新初始化
picoclaw goagents config init --force
```

### 问题3: goagents 子命令不可用
```bash
# 错误信息
Error: unknown command "workflow" for "goagents"

# 诊断步骤
1. 检查 goagents 版本
picoclaw goagents version

2. 检查可用命令
picoclaw goagents --help

# 解决方案
# 1. 更新到最新版本
picoclaw update

# 2. 重新安装 goagents
picoclaw goagents install --force
```

## 🎯 技能问题

### 问题1: @go 触发无响应
```bash
# 测试命令
@go "测试消息"

# 诊断步骤
1. 检查 PO 技能状态
picoclaw skills status po-core

2. 检查技能是否启用
picoclaw skills list | grep po-core

3. 检查技能配置
picoclaw skills show po-core

# 解决方案
# 1. 重新启用技能
picoclaw skills disable po-core
picoclaw skills enable po-core

# 2. 重新安装技能
picoclaw skills uninstall po-core
picoclaw skills install po-core

# 3. 检查触发器配置
picoclaw skills edit po-core --trigger "@go"
```

### 问题2: 技能加载失败
```bash
# 错误信息
Error: failed to load skill po-core

# 诊断步骤
1. 检查技能文件完整性
ls -la ~/.picoclaw/workspace/skills/po-core/

2. 检查技能文件格式
picoclaw skills validate po-core

3. 检查依赖技能
picoclaw skills dependencies po-core

# 解决方案
# 1. 修复技能文件
picoclaw skills repair po-core

# 2. 重新下载技能
picoclaw skills download po-core

# 3. 检查依赖关系
picoclaw skills install-dependencies po-core
```

### 问题3: 技能执行错误
```bash
# 错误信息
Error: skill execution failed: po-core

# 诊断步骤
1. 查看详细错误日志
picoclaw goagents logs --skill po-core --verbose

2. 检查技能配置
picoclaw skills debug po-core

3. 测试技能执行
picoclaw skills test po-core

# 解决方案
# 1. 调试模式执行
picoclaw goagents --debug @go "测试消息"

# 2. 重置技能状态
picoclaw skills reset po-core

# 3. 更新技能版本
picoclaw skills update po-core
```

## ⚙️ 配置问题

### 问题1: 配置文件格式错误
```bash
# 错误信息
Error: invalid YAML format in config.yaml

# 诊断步骤
1. 验证配置文件
picoclaw goagents config validate

2. 检查 YAML 语法
yamllint ~/.picoclaw/workspace/.goagents/config.yaml

3. 查看具体错误
picoclaw goagents config validate --verbose

# 解决方案
# 1. 自动修复配置
picoclaw goagents config repair

# 2. 重新生成配置
picoclaw goagents config init --force

# 3. 手动编辑配置
vim ~/.picoclaw/workspace/.goagents/config.yaml
```

### 问题2: 工作流配置丢失
```bash
# 错误信息
Error: workflow not found: standard-development

# 诊断步骤
1. 检查工作流列表
picoclaw goagents workflow list

2. 检查工作流文件
ls -la ~/.picoclaw/workspace/.goagents/workflows/

3. 检查配置路径
picoclaw goagents config show paths

# 解决方案
# 1. 重新创建工作流
picoclaw goagents workflow create standard-development

# 2. 从模板恢复
picoclaw goagents workflow restore --template standard

# 3. 导入备份配置
picoclaw goagents import workflow-backup.yaml
```

### 问题3: 团队配置错误
```bash
# 错误信息
Error: invalid team configuration

# 诊断步骤
1. 验证团队配置
picoclaw goagents team validate discovery-team

2. 检查角色定义
picoclaw goagents team show discovery-team --verbose

3. 检查角色技能
picoclaw goagents team check-skills discovery-team

# 解决方案
# 1. 修复团队配置
picoclaw goagents team repair discovery-team

# 2. 重新定义角色
picoclaw goagents team redefine-role discovery-team analyst

# 3. 重新创建团队
picoclaw goagents team create discovery-team --force
```

## 🚀 执行问题

### 问题1: 任务执行卡住
```bash
# 诊断步骤
1. 检查执行状态
picoclaw goagents status --current-task

2. 查看执行日志
picoclaw goagents logs --current --verbose

3. 检查资源使用
picoclaw goagents metrics --resource

# 解决方案
# 1. 强制停止任务
picoclaw goagents stop --force

# 2. 重启执行
picoclaw goagents restart --from-last-checkpoint

# 3. 降级执行模式
picoclaw goagents execute --mode safe
```

### 问题2: 阶段转换失败
```bash
# 错误信息
Error: phase transition failed: discovery to planning

# 诊断步骤
1. 检查阶段状态
picoclaw goagents phase status

2. 检查转换条件
picoclaw goagents phase check-transition discovery planning

3. 查看质量门禁
picoclaw goagents quality-gates show --phase discovery

# 解决方案
# 1. 修复质量门禁
picoclaw goagents quality-gates fix --phase discovery

# 2. 强制转换（谨慎使用）
picoclaw goagents phase transition discovery planning --force

# 3. 重新执行阶段
picoclaw goagents phase retry discovery
```

### 问题3: 质量门禁失败
```bash
# 错误信息
Error: quality gate failed: test_coverage

# 诊断步骤
1. 查看质量门禁详情
picoclaw goagents quality-gates show --failed

2. 检查具体失败项
picoclaw goagents quality check --detailed

3. 生成修复建议
picoclaw goagents quality suggest-fix

# 解决方案
# 1. 自动修复
picoclaw goagents quality auto-fix

# 2. 手动修复指导
picoclaw goagents quality fix-guide --gate test_coverage

# 3. 调整质量标准
picoclaw goagents config set quality_gates.test_coverage_min 80
```

## 📊 性能问题

### 问题1: 响应时间过长
```bash
# 诊断步骤
1. 检查响应时间
picoclaw goagents metrics --response-time

2. 分析性能瓶颈
picoclaw goagents performance analyze

3. 查看资源使用
picoclaw goagents metrics --resource

# 解决方案
# 1. 优化配置
picoclaw goagents config set performance.cache_enabled true
picoclaw goagents config set performance.parallel_execution true

# 2. 清理缓存
picoclaw goagents cache clear --all

# 3. 升级资源配置
picoclaw goagents config set performance.memory_limit 2GB
picoclaw goagents config set performance.cpu_limit 4
```

### 问题2: 内存使用过高
```bash
# 诊断步骤
1. 检查内存使用
picoclaw goagents metrics --memory

2. 分析内存泄漏
picoclaw goagents memory analyze

3. 查看缓存状态
picoclaw goagents cache status

# 解决方案
# 1. 清理内存
picoclaw goagents memory cleanup

# 2. 调整缓存大小
picoclaw goagents config set cache.max_size 50

# 3. 重启服务
picoclaw goagents restart --graceful
```

### 问题3: 并发处理能力不足
```bash
# 诊断步骤
1. 检查并发状态
picoclaw goagents metrics --concurrency

2. 分析并发瓶颈
picoclaw goagents concurrency analyze

3. 查看队列状态
picoclaw goagents queue status

# 解决方案
# 1. 增加并发数
picoclaw goagents config set concurrency.max_workers 8

# 2. 优化队列处理
picoclaw goagents queue optimize

# 3. 负载均衡
picoclaw goagents load-balance enable
```

## 🔒 安全问题

### 问题1: 权限验证失败
```bash
# 错误信息
Error: permission denied: access denied to configuration

# 诊断步骤
1. 检查用户权限
picoclaw goagents auth check

2. 查看权限配置
picoclaw goagents auth show

3. 验证身份
picoclaw goagents auth verify

# 解决方案
# 1. 重新认证
picoclaw goagents auth login

# 2. 修复权限
picoclaw goagents auth fix-permissions

# 3. 重置权限
picoclaw goagents auth reset --user current
```

### 问题2: 配置文件泄露风险
```bash
# 诊断步骤
1. 检查文件权限
ls -la ~/.picoclaw/workspace/.goagents/

2. 扫描敏感信息
picoclaw goagents security scan --secrets

3. 检查访问日志
picoclaw goagents security audit-log

# 解决方案
# 1. 修复文件权限
chmod 600 ~/.picoclaw/workspace/.goagents/config.yaml

# 2. 加密敏感配置
picoclaw goagents security encrypt --config

# 3. 设置访问控制
picoclaw goagents security acl enable
```

## 🔧 调试工具

### 调试模式
```bash
# 启用调试模式
export PICOLAW_DEBUG=true
export GOAGENTS_DEBUG=true

# 调试执行
picoclaw goagents --debug @go "测试消息"

# 查看调试日志
picoclaw goagents logs --debug --tail 100
```

### 性能分析
```bash
# 性能分析
picoclaw goagents profile --cpu --memory

# 生成性能报告
picoclaw goagents performance report --output performance.html

# 实时监控
picoclaw goagents monitor --real-time
```

### 健康检查
```bash
# 系统健康检查
picoclaw goagents health-check --full

# 组件健康检查
picoclaw goagents health-check --component skills
picoclaw goagents health-check --component config
picoclaw goagents health-check --component execution

# 生成健康报告
picoclaw goagents health-report --output health.json
```

## 📞 获取帮助

### 自动诊断
```bash
# 运行自动诊断
picoclaw goagents diagnose --auto

# 生成诊断报告
picoclaw goagents diagnose --report > diagnosis.txt

# 发送诊断报告
picoclaw goagents diagnose --send-support
```

### 社区支持
```bash
# 查看常见问题
picoclaw goagents faq

# 搜索解决方案
picoclaw goagents search "permission denied"

# 获取社区帮助
picoclaw goagents community-help --issue "goagents command not found"
```

### 官方支持
```bash
# 生成支持请求
picoclaw goagents support-request --template bug-report

# 收集系统信息
picoclaw goagents system-info > system-info.txt

# 提交支持请求
picoclaw goagents support-submit --files system-info.txt,diagnosis.txt
```

## 🎯 预防措施

### 定期维护
```bash
# 定期清理
picoclaw goagents cleanup --schedule weekly

# 定期备份
picoclaw goagents backup --schedule daily

# 定期更新
picoclaw goagents update --schedule monthly
```

### 监控告警
```bash
# 设置监控告警
picoclaw goagents monitor --alert-threshold cpu:80,memory:90

# 配置告警通知
picoclaw goagents alert configure --email admin@example.com

# 测试告警
picoclaw goagents alert test
```

### 版本管理
```bash
# 检查版本兼容性
picoclaw goagents version-compatibility

# 锁定稳定版本
picoclaw goagents version lock stable

# 版本回滚
picoclaw goagents version rollback previous
```

## 🎉 总结

通过本故障排除指南，用户可以：

1. **快速诊断问题** - 系统化的问题诊断流程
2. **有效解决问题** - 详细的解决方案和步骤
3. **预防问题发生** - 主动的维护和监控措施
4. **获取专业帮助** - 多层次的支持渠道

记住，大部分问题都可以通过简单的重启、重新配置或更新来解决。如果问题持续存在，请及时寻求专业支持！🚀
