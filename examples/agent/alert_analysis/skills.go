package main

import (
	"context"
	"fmt"
	"time"
)

// SkillType 技能类型
type SkillType string

const (
	// 数据采集技能
	LogQuerySkill          SkillType = "log_query"           // 日志查询
	MetricQuerySkill       SkillType = "metric_query"        // 监控指标查询
	TraceQuerySkill        SkillType = "trace_query"         // 链路追踪
	TopologyQuerySkill     SkillType = "topology_query"      // 拓扑分析
	ConfigQuerySkill       SkillType = "config_query"        // 配置查询

	// 分析技能
	TimeSeriesAnalysisSkill    SkillType = "timeseries_analysis"    // 时间序列分析
	CorrelationAnalysisSkill   SkillType = "correlation_analysis"   // 关联分析
	PatternMatchingSkill       SkillType = "pattern_matching"       // 模式匹配
	AnomalyDetectionSkill      SkillType = "anomaly_detection"      // 异常检测
	HistoryMatchSkill          SkillType = "history_match"          // 历史案例匹配
	RootCauseAnalysisSkill     SkillType = "rootcause_analysis"     // 根因分析
	FaultPropagationSkill      SkillType = "fault_propagation"      // 故障传播分析
	DependencyAnalysisSkill    SkillType = "dependency_analysis"    // 依赖分析
	CapacityAnalysisSkill      SkillType = "capacity_analysis"      // 容量分析
	TrendAnalysisSkill         SkillType = "trend_analysis"         // 趋势分析

	// 诊断技能
	SlowQueryAnalysisSkill   SkillType = "slowquery_analysis"   // 慢查询分析
	ErrorLogAnalysisSkill    SkillType = "errorlog_analysis"    // 错误日志分析
	NetworkDiagnosisSkill    SkillType = "network_diagnosis"    // 网络诊断
	ResourceLeakDetectionSkill SkillType = "resource_leak_detection" // 资源泄漏检测
	MemoryAnalysisSkill      SkillType = "memory_analysis"      // 内存分析
	CPUAnalysisSkill         SkillType = "cpu_analysis"         // CPU分析
	DiskAnalysisSkill        SkillType = "disk_analysis"        // 磁盘分析
	DatabaseAnalysisSkill    SkillType = "database_analysis"    // 数据库分析
	CacheAnalysisSkill       SkillType = "cache_analysis"       // 缓存分析
	LoadBalancerAnalysisSkill SkillType = "loadbalancer_analysis" // 负载均衡分析

	// 验证技能
	ConnectivityCheckSkill   SkillType = "connectivity_check"   // 连通性检查
	HealthCheckSkill         SkillType = "health_check"         // 健康检查
	PerformanceTestSkill     SkillType = "performance_test"     // 性能测试
	StressTestSkill          SkillType = "stress_test"          // 压力测试
	ValidationCheckSkill     SkillType = "validation_check"     // 验证检查
)

// Skill 技能接口
type Skill interface {
	Execute(ctx context.Context, input any) (*SkillResult, error)
	GetName() string
	GetType() SkillType
	GetDescription() string
}

// SkillResult 技能执行结果
type SkillResult struct {
	Success   bool                   `json:"success"`
	Data      map[string]any         `json:"data"`
	Message   string                 `json:"message"`
	Confidence float64               `json:"confidence"`
	Timestamp int64                  `json:"timestamp"`
	Duration  time.Duration          `json:"duration"`
}

// BaseSkill 基础技能
type BaseSkill struct {
	name        string
	skillType   SkillType
	description string
}

// GetName 获取技能名称
func (b *BaseSkill) GetName() string {
	return b.name
}

// GetType 获取技能类型
func (b *BaseSkill) GetType() SkillType {
	return b.skillType
}

// GetDescription 获取技能描述
func (b *BaseSkill) GetDescription() string {
	return b.description
}

// ==================== 数据采集技能 ====================

// LogQuerySkill 日志查询技能
type LogQuerySkillImpl struct {
	BaseSkill
}

func NewLogQuerySkill() *LogQuerySkillImpl {
	return &LogQuerySkillImpl{
		BaseSkill: BaseSkill{
			name:        "日志查询",
			skillType:   LogQuerySkill,
			description: "查询指定时间范围内的日志信息",
		},
	}
}

func (s *LogQuerySkillImpl) Execute(ctx context.Context, input any) (*SkillResult, error) {
	start := time.Now()
	// Mock 实现
	result := &SkillResult{
		Success: true,
		Data: map[string]any{
			"logCount":       1523,
			"errorLogs":      45,
			"warningLogs":    128,
			"sampleLogs": []string{
				"[ERROR] 2024-01-15 10:23:45 connection pool exhausted",
				"[WARN] 2024-01-15 10:23:46 query timeout: 5000ms",
				"[ERROR] 2024-01-15 10:23:47 failed to acquire connection",
			},
			"errorPatterns": []string{
				"connection pool exhausted",
				"query timeout",
				"failed to acquire connection",
			},
			"timeRange": "last 15 minutes",
		},
		Message:   "成功查询日志",
		Confidence: 0.95,
		Timestamp: time.Now().Unix(),
		Duration:  time.Since(start),
	}
	return result, nil
}

// MetricQuerySkill 监控指标查询技能
type MetricQuerySkillImpl struct {
	BaseSkill
}

func NewMetricQuerySkill() *MetricQuerySkillImpl {
	return &MetricQuerySkillImpl{
		BaseSkill: BaseSkill{
			name:        "监控指标查询",
			skillType:   MetricQuerySkill,
			description: "查询 CPU、内存、网络等监控指标",
		},
	}
}

func (s *MetricQuerySkillImpl) Execute(ctx context.Context, input any) (*SkillResult, error) {
	start := time.Now()
	result := &SkillResult{
		Success: true,
		Data: map[string]any{
			"cpu": map[string]any{
				"usage":      85.6,
				"load1m":    12.5,
				"load5m":    10.2,
				"load15m":   8.7,
				"cores":     16,
			},
			"memory": map[string]any{
				"total":       32.0,
				"used":        28.5,
				"free":        3.5,
				"usagePercent": 89.1,
			},
			"disk": map[string]any{
				"total":       500.0,
				"used":        350.0,
				"free":        150.0,
				"iops":        2500,
				"throughput":  "125 MB/s",
			},
			"network": map[string]any{
				"inBytes":    125000000,
				"outBytes":   98000000,
				"inPps":      8500,
				"outPps":     7200,
			},
			"database": map[string]any{
				"connections":      85,
				"maxConnections":   100,
				"qps":              2500,
				"slowQueryCount":   12,
				"avgQueryTime":     450,
			},
			"cache": map[string]any{
				"hitRate":         0.85,
				"memoryUsed":      2.5,
				"keyCount":        125000,
				"evictions":       125,
			},
		},
		Message:    "成功查询监控指标",
		Confidence: 0.98,
		Timestamp:   time.Now().Unix(),
		Duration:   time.Since(start),
	}
	return result, nil
}

// TraceQuerySkill 链路追踪技能
type TraceQuerySkillImpl struct {
	BaseSkill
}

func NewTraceQuerySkill() *TraceQuerySkillImpl {
	return &TraceQuerySkillImpl{
		BaseSkill: BaseSkill{
			name:        "链路追踪",
			skillType:   TraceQuerySkill,
			description: "查询分布式链路追踪信息",
		},
	}
}

func (s *TraceQuerySkillImpl) Execute(ctx context.Context, input any) (*SkillResult, error) {
	start := time.Now()
	result := &SkillResult{
		Success: true,
		Data: map[string]any{
			"traceId": "trace-1234567890",
			"spans": []map[string]any{
				{
					"service":     "api-gateway",
					"operation":  "POST /api/payment",
					"duration":   2340,
					"status":     "error",
				},
				{
					"service":    "payment-service",
					"operation": "processPayment",
					"duration":   2100,
					"status":     "error",
				},
				{
					"service":    "database",
					"operation": "query",
					"duration":   1850,
					"status":     "error",
				},
			},
			"totalDuration": 2340,
			"bottleneck":    "database",
			"errorSpans":    3,
		},
		Message:    "成功查询链路追踪信息",
		Confidence: 0.92,
		Timestamp:   time.Now().Unix(),
		Duration:   time.Since(start),
	}
	return result, nil
}

// TopologyQuerySkill 拓扑分析技能
type TopologyQuerySkillImpl struct {
	BaseSkill
}

func NewTopologyQuerySkill() *TopologyQuerySkillImpl {
	return &TopologyQuerySkillImpl{
		BaseSkill: BaseSkill{
			name:        "拓扑分析",
			skillType:   TopologyQuerySkill,
			description: "分析服务依赖关系和拓扑结构",
		},
	}
}

func (s *TopologyQuerySkillImpl) Execute(ctx context.Context, input any) (*SkillResult, error) {
	start := time.Now()
	result := &SkillResult{
		Success: true,
		Data: map[string]any{
			"services": []string{
				"api-gateway",
				"payment-service",
				"user-service",
				"order-service",
				"database",
				"redis-cache",
			},
			"dependencies": []map[string]string{
				{"from": "api-gateway", "to": "payment-service"},
				{"from": "api-gateway", "to": "user-service"},
				{"from": "api-gateway", "to": "order-service"},
				{"from": "payment-service", "to": "database"},
				{"from": "payment-service", "to": "redis-cache"},
				{"from": "user-service", "to": "database"},
				{"from": "order-service", "to": "database"},
			},
			"criticalPath": []string{
				"api-gateway",
				"payment-service",
				"database",
			},
			"impactScope": []string{
				"payment-service",
				"order-service",
			},
		},
		Message:    "成功分析服务拓扑",
		Confidence: 0.90,
		Timestamp:   time.Now().Unix(),
		Duration:   time.Since(start),
	}
	return result, nil
}

// ==================== 分析技能 ====================

// TimeSeriesAnalysisSkill 时间序列分析技能
type TimeSeriesAnalysisSkillImpl struct {
	BaseSkill
}

func NewTimeSeriesAnalysisSkill() *TimeSeriesAnalysisSkillImpl {
	return &TimeSeriesAnalysisSkillImpl{
		BaseSkill: BaseSkill{
			name:        "时间序列分析",
			skillType:   TimeSeriesAnalysisSkill,
			description: "分析指标的时间序列趋势和异常点",
		},
	}
}

func (s *TimeSeriesAnalysisSkillImpl) Execute(ctx context.Context, input any) (*SkillResult, error) {
	start := time.Now()
	result := &SkillResult{
		Success: true,
		Data: map[string]any{
			"trend": "exponential_increase",
			"anomalies": []map[string]any{
				{
					"timestamp":   "2024-01-15T10:23:00Z",
					"value":       2340,
					"baseline":    200,
					"deviation":   11.7,
					"type":        "spike",
				},
			},
			"baseline":     200,
			"currentValue": 2340,
			"changeRate":   1070.0,
			"forecast": []float64{
				2100, 2200, 2300, 2400, 2500,
			},
		},
		Message:    "检测到异常增长趋势",
		Confidence: 0.89,
		Timestamp:   time.Now().Unix(),
		Duration:   time.Since(start),
	}
	return result, nil
}

// CorrelationAnalysisSkill 关联分析技能
type CorrelationAnalysisSkillImpl struct {
	BaseSkill
}

func NewCorrelationAnalysisSkill() *CorrelationAnalysisSkillImpl {
	return &CorrelationAnalysisSkillImpl{
		BaseSkill: BaseSkill{
			name:        "关联分析",
			skillType:   CorrelationAnalysisSkill,
			description: "分析告警和指标之间的关联关系",
		},
	}
}

func (s *CorrelationAnalysisSkillImpl) Execute(ctx context.Context, input any) (*SkillResult, error) {
	start := time.Now()
	result := &SkillResult{
		Success: true,
		Data: map[string]any{
			"correlations": []map[string]any{
				{
					"metric1":       "database.connections",
					"metric2":       "api.latency",
					"correlation":   0.92,
					"relationship":  "strong_positive",
				},
				{
					"metric1":       "memory.usage",
					"metric2":       "api.error_rate",
					"correlation":   0.78,
					"relationship":  "moderate_positive",
				},
			},
			"rootCauseCandidates": []string{
				"database connection pool exhaustion",
				"high memory usage",
			},
		},
		Message:    "发现强关联关系",
		Confidence: 0.85,
		Timestamp:   time.Now().Unix(),
		Duration:   time.Since(start),
	}
	return result, nil
}

// PatternMatchingSkill 模式匹配技能
type PatternMatchingSkillImpl struct {
	BaseSkill
}

func NewPatternMatchingSkill() *PatternMatchingSkillImpl {
	return &PatternMatchingSkillImpl{
		BaseSkill: BaseSkill{
			name:        "模式匹配",
			skillType:   PatternMatchingSkill,
			description: "匹配已知的故障模式",
		},
	}
}

func (s *PatternMatchingSkillImpl) Execute(ctx context.Context, input any) (*SkillResult, error) {
	start := time.Now()
	result := &SkillResult{
		Success: true,
		Data: map[string]any{
			"matchedPatterns": []map[string]any{
				{
					"patternName":  "connection_pool_exhaustion",
					"similarity":   0.92,
					"description":  "数据库连接池耗尽",
					"symptoms": []string{
						"high latency",
						"connection timeout",
						"pool exhaustion errors",
					},
				},
				{
					"patternName":  "memory_leak",
					"similarity":  0.75,
					"description":  "内存泄漏",
					"symptoms": []string{
						"gradual memory increase",
						"GC pressure",
					},
				},
			},
			"primaryPattern": "connection_pool_exhaustion",
		},
		Message:    "匹配到故障模式",
		Confidence: 0.88,
		Timestamp:   time.Now().Unix(),
		Duration:   time.Since(start),
	}
	return result, nil
}

// HistoryMatchSkill 历史案例匹配技能
type HistoryMatchSkillImpl struct {
	BaseSkill
}

func NewHistoryMatchSkill() *HistoryMatchSkillImpl {
	return &HistoryMatchSkillImpl{
		BaseSkill: BaseSkill{
			name:        "历史案例匹配",
			skillType:   HistoryMatchSkill,
			description: "匹配历史相似故障案例",
		},
	}
}

func (s *HistoryMatchSkillImpl) Execute(ctx context.Context, input any) (*SkillResult, error) {
	start := time.Now()
	result := &SkillResult{
		Success: true,
		Data: map[string]any{
			"similarCases": []map[string]any{
				{
					"caseId":       "INC-2024-0115-001",
					"date":         "2024-01-15",
					"similarity":   0.94,
					"rootCause":    "database connection pool too small",
					"resolution":   "increased pool size from 50 to 100",
					"resolutionTime": "2 hours",
				},
				{
					"caseId":       "INC-2024-0108-003",
					"date":         "2024-01-08",
					"similarity":   0.87,
					"rootCause":    "slow queries causing connection buildup",
					"resolution":   "optimized queries and added indexes",
					"resolutionTime": "4 hours",
				},
			},
			"recommendedActions": []string{
				"increase database connection pool size",
				"check for slow queries",
				"monitor connection usage",
			},
		},
		Message:    "找到相似历史案例",
		Confidence: 0.91,
		Timestamp:   time.Now().Unix(),
		Duration:   time.Since(start),
	}
	return result, nil
}

// RootCauseAnalysisSkill 根因分析技能
type RootCauseAnalysisSkillImpl struct {
	BaseSkill
}

func NewRootCauseAnalysisSkill() *RootCauseAnalysisSkillImpl {
	return &RootCauseAnalysisSkillImpl{
		BaseSkill: BaseSkill{
			name:        "根因分析",
			skillType:   RootCauseAnalysisSkill,
			description: "深度分析故障根本原因",
		},
	}
}

func (s *RootCauseAnalysisSkillImpl) Execute(ctx context.Context, input any) (*SkillResult, error) {
	start := time.Now()
	result := &SkillResult{
		Success: true,
		Data: map[string]any{
			"rootCauses": []map[string]any{
				{
					"cause":       "database connection pool misconfiguration",
					"probability": 0.85,
					"evidence": []string{
						"connection pool exhausted errors in logs",
						"high correlation between connections and latency",
						"similar historical cases",
					},
				},
				{
					"cause":       "slow queries causing connection buildup",
					"probability": 0.65,
					"evidence": []string{
						"12 slow queries detected",
						"high database query times",
					},
				},
			},
			"primaryRootCause": "database connection pool misconfiguration",
			"causalityChain": []string{
				"high traffic load",
				"increased database queries",
				"connection pool saturation",
				"query timeouts",
				"api latency increase",
			},
		},
		Message:    "完成根因分析",
		Confidence: 0.84,
		Timestamp:   time.Now().Unix(),
		Duration:   time.Since(start),
	}
	return result, nil
}

// FaultPropagationSkill 故障传播分析技能
type FaultPropagationSkillImpl struct {
	BaseSkill
}

func NewFaultPropagationSkill() *FaultPropagationSkillImpl {
	return &FaultPropagationSkillImpl{
		BaseSkill: BaseSkill{
			name:        "故障传播分析",
			skillType:   FaultPropagationSkill,
			description: "分析故障在系统中的传播路径",
		},
	}
}

func (s *FaultPropagationSkillImpl) Execute(ctx context.Context, input any) (*SkillResult, error) {
	start := time.Now()
	result := &SkillResult{
		Success: true,
		Data: map[string]any{
			"propagationPath": []map[string]any{
				{
					"service":    "database",
					"status":     "root_cause",
					"timestamp":  "2024-01-15T10:23:30Z",
				},
				{
					"service":    "payment-service",
					"status":     "impacted",
					"timestamp":  "2024-01-15T10:23:35Z",
					"delay":      "5s",
				},
				{
					"service":    "api-gateway",
					"status":     "impacted",
					"timestamp":  "2024-01-15T10:23:40Z",
					"delay":      "10s",
				},
			},
			"affectedServices": []string{
				"payment-service",
				"order-service",
				"api-gateway",
			},
			"impactLevel": "high",
		},
		Message:    "完成故障传播分析",
		Confidence: 0.87,
		Timestamp:   time.Now().Unix(),
		Duration:   time.Since(start),
	}
	return result, nil
}

// ==================== 诊断技能 ====================

// SlowQueryAnalysisSkill 慢查询分析技能
type SlowQueryAnalysisSkillImpl struct {
	BaseSkill
}

func NewSlowQueryAnalysisSkill() *SlowQueryAnalysisSkillImpl {
	return &SlowQueryAnalysisSkillImpl{
		BaseSkill: BaseSkill{
			name:        "慢查询分析",
			skillType:   SlowQueryAnalysisSkill,
			description: "分析数据库慢查询",
		},
	}
}

func (s *SlowQueryAnalysisSkillImpl) Execute(ctx context.Context, input any) (*SkillResult, error) {
	start := time.Now()
	result := &SkillResult{
		Success: true,
		Data: map[string]any{
			"slowQueries": []map[string]any{
				{
					"query":      "SELECT * FROM orders WHERE user_id = ?",
					"avgTime":    1250,
					"count":      456,
					"rows":       15000,
				},
				{
					"query":      "SELECT * FROM payments WHERE status = 'pending'",
					"avgTime":    850,
					"count":      289,
					"rows":       8500,
				},
			},
			"recommendations": []string{
				"add index on orders.user_id",
				"add index on payments.status",
				"optimize query to use index",
			},
			"potentialImprovement": "65%",
		},
		Message:    "发现慢查询",
		Confidence: 0.92,
		Timestamp:   time.Now().Unix(),
		Duration:   time.Since(start),
	}
	return result, nil
}

// ErrorLogAnalysisSkill 错误日志分析技能
type ErrorLogAnalysisSkillImpl struct {
	BaseSkill
}

func NewErrorLogAnalysisSkill() *ErrorLogAnalysisSkillImpl {
	return &ErrorLogAnalysisSkillImpl{
		BaseSkill: BaseSkill{
			name:        "错误日志分析",
			skillType:   ErrorLogAnalysisSkill,
			description: "分析错误日志模式和趋势",
		},
	}
}

func (s *ErrorLogAnalysisSkillImpl) Execute(ctx context.Context, input any) (*SkillResult, error) {
	start := time.Now()
	result := &SkillResult{
		Success: true,
		Data: map[string]any{
			"errorStats": map[string]any{
				"totalErrors":     1523,
				"uniqueErrors":    45,
				"topError":        "connection pool exhausted",
				"errorRate":       0.05,
			},
			"topErrors": []map[string]any{
				{
					"message":   "connection pool exhausted",
					"count":     425,
					"trend":     "increasing",
					"service":   "payment-service",
				},
				{
					"message":   "query timeout",
					"count":     312,
					"trend":     "increasing",
					"service":   "database",
				},
			},
			"errorPatterns": []string{
				"connection-related errors",
				"timeout errors",
				"resource exhaustion",
			},
		},
		Message:    "完成错误日志分析",
		Confidence: 0.94,
		Timestamp:   time.Now().Unix(),
		Duration:   time.Since(start),
	}
	return result, nil
}

// ResourceLeakDetectionSkill 资源泄漏检测技能
type ResourceLeakDetectionSkillImpl struct {
	BaseSkill
}

func NewResourceLeakDetectionSkill() *ResourceLeakDetectionSkillImpl {
	return &ResourceLeakDetectionSkillImpl{
		BaseSkill: BaseSkill{
			name:        "资源泄漏检测",
			skillType:   ResourceLeakDetectionSkill,
			description: "检测内存、连接等资源泄漏",
		},
	}
}

func (s *ResourceLeakDetectionSkillImpl) Execute(ctx context.Context, input any) (*SkillResult, error) {
	start := time.Now()
	result := &SkillResult{
		Success: true,
		Data: map[string]any{
			"leaks": []map[string]any{
				{
					"type":        "database_connection",
					"severity":    "high",
					"leakRate":    "5 connections/min",
					"estimated":   150,
					"evidence": []string{
						"connections not released after timeout",
						"gradual increase in active connections",
					},
				},
			},
			"memoryAnalysis": map[string]any{
				"heapGrowth":     "gradual",
				"gcPressure":     "moderate",
				"potentialLeak":   false,
			},
		},
		Message:    "检测到连接泄漏",
		Confidence: 0.82,
		Timestamp:   time.Now().Unix(),
		Duration:   time.Since(start),
	}
	return result, nil
}

// NetworkDiagnosisSkill 网络诊断技能
type NetworkDiagnosisSkillImpl struct {
	BaseSkill
}

func NewNetworkDiagnosisSkill() *NetworkDiagnosisSkillImpl {
	return &NetworkDiagnosisSkillImpl{
		BaseSkill: BaseSkill{
			name:        "网络诊断",
			skillType:   NetworkDiagnosisSkill,
			description: "诊断网络连接和延迟问题",
		},
	}
}

func (s *NetworkDiagnosisSkillImpl) Execute(ctx context.Context, input any) (*SkillResult, error) {
	start := time.Now()
	result := &SkillResult{
		Success: true,
		Data: map[string]any{
			"connectivity": map[string]any{
				"dns":           "ok",
				"tcp":           "ok",
				"tls":           "ok",
			},
			"latency": map[string]any{
				"avg":          25,
				"p50":          20,
				"p95":          45,
				"p99":          80,
			},
			"packetLoss":    0.001,
			"bandwidth":     "1 Gbps",
			"conclusion":    "network is healthy",
		},
		Message:    "网络诊断正常",
		Confidence: 0.96,
		Timestamp:   time.Now().Unix(),
		Duration:   time.Since(start),
	}
	return result, nil
}

// MemoryAnalysisSkill 内存分析技能
type MemoryAnalysisSkillImpl struct {
	BaseSkill
}

func NewMemoryAnalysisSkill() *MemoryAnalysisSkillImpl {
	return &MemoryAnalysisSkillImpl{
		BaseSkill: BaseSkill{
			name:        "内存分析",
			skillType:   MemoryAnalysisSkill,
			description: "分析内存使用情况和问题",
		},
	}
}

func (s *MemoryAnalysisSkillImpl) Execute(ctx context.Context, input any) (*SkillResult, error) {
	start := time.Now()
	result := &SkillResult{
		Success: true,
		Data: map[string]any{
			"memoryUsage": map[string]any{
				"heap":          28.5,
				"stack":         0.5,
				"code":          0.3,
				"buffers":       2.7,
			},
			"gcStats": map[string]any{
				"gcCount":       125,
				"gcTime":        "2.5s",
				"gcPauseAvg":    "20ms",
				"gcPauseMax":    "150ms",
			},
			"conclusion":     "high memory usage but no leak detected",
		},
		Message:    "内存使用较高但无泄漏",
		Confidence: 0.88,
		Timestamp:   time.Now().Unix(),
		Duration:   time.Since(start),
	}
	return result, nil
}

// ==================== Skill Registry ====================

// SkillRegistry 技能注册表
type SkillRegistry struct {
	skills map[SkillType]Skill
}

// NewSkillRegistry 创建技能注册表
func NewSkillRegistry() *SkillRegistry {
	registry := &SkillRegistry{
		skills: make(map[SkillType]Skill),
	}

	// 注册所有技能
	registry.Register(NewLogQuerySkill())
	registry.Register(NewMetricQuerySkill())
	registry.Register(NewTraceQuerySkill())
	registry.Register(NewTopologyQuerySkill())
	registry.Register(NewTimeSeriesAnalysisSkill())
	registry.Register(NewCorrelationAnalysisSkill())
	registry.Register(NewPatternMatchingSkill())
	registry.Register(NewHistoryMatchSkill())
	registry.Register(NewRootCauseAnalysisSkill())
	registry.Register(NewFaultPropagationSkill())
	registry.Register(NewSlowQueryAnalysisSkill())
	registry.Register(NewErrorLogAnalysisSkill())
	registry.Register(NewResourceLeakDetectionSkill())
	registry.Register(NewNetworkDiagnosisSkill())
	registry.Register(NewMemoryAnalysisSkill())

	return registry
}

// Register 注册技能
func (r *SkillRegistry) Register(skill Skill) {
	r.skills[skill.GetType()] = skill
}

// Get 获取技能
func (r *SkillRegistry) Get(skillType SkillType) (Skill, bool) {
	skill, ok := r.skills[skillType]
	return skill, ok
}

// List 列出所有技能
func (r *SkillRegistry) List() []Skill {
	skills := make([]Skill, 0, len(r.skills))
	for _, skill := range r.skills {
		skills = append(skills, skill)
	}
	return skills
}

// ListByCategory 按类别列出技能
func (r *SkillRegistry) ListByCategory() map[string][]Skill {
	categories := map[string][]Skill{
		"数据采集": {},
		"分析能力": {},
		"诊断能力": {},
		"验证能力": {},
	}

	for _, skill := range r.skills {
		skillType := skill.GetType()
		switch {
		case skillType == LogQuerySkill || skillType == MetricQuerySkill ||
		     skillType == TraceQuerySkill || skillType == TopologyQuerySkill ||
		     skillType == ConfigQuerySkill:
			categories["数据采集"] = append(categories["数据采集"], skill)
		case skillType == TimeSeriesAnalysisSkill || skillType == CorrelationAnalysisSkill ||
		     skillType == PatternMatchingSkill || skillType == HistoryMatchSkill ||
		     skillType == RootCauseAnalysisSkill || skillType == FaultPropagationSkill ||
		     skillType == DependencyAnalysisSkill || skillType == CapacityAnalysisSkill ||
		     skillType == TrendAnalysisSkill || skillType == AnomalyDetectionSkill:
			categories["分析能力"] = append(categories["分析能力"], skill)
		case skillType == SlowQueryAnalysisSkill || skillType == ErrorLogAnalysisSkill ||
		     skillType == NetworkDiagnosisSkill || skillType == ResourceLeakDetectionSkill ||
		     skillType == MemoryAnalysisSkill || skillType == CPUAnalysisSkill ||
		     skillType == DiskAnalysisSkill || skillType == DatabaseAnalysisSkill ||
		     skillType == CacheAnalysisSkill || skillType == LoadBalancerAnalysisSkill:
			categories["诊断能力"] = append(categories["诊断能力"], skill)
		case skillType == ConnectivityCheckSkill || skillType == HealthCheckSkill ||
		     skillType == PerformanceTestSkill || skillType == StressTestSkill ||
		     skillType == ValidationCheckSkill:
			categories["验证能力"] = append(categories["验证能力"], skill)
		}
	}

	return categories
}

// PrintSkills 打印所有技能
func (r *SkillRegistry) PrintSkills() {
	fmt.Println("\n========== 可用技能列表 ==========")
	categories := r.ListByCategory()
	for categoryName, skills := range categories {
		if len(skills) > 0 {
			fmt.Printf("\n【%s】\n", categoryName)
			for _, skill := range skills {
				fmt.Printf("  - %s (%s)\n", skill.GetName(), skill.GetType())
				fmt.Printf("    %s\n", skill.GetDescription())
			}
		}
	}
	fmt.Println("\n==================================")
}
