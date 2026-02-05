package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

// mockChatModel 是一个模拟的 ChatModel 实现
// 在实际使用中，应该使用 eino-ext 中的 openai 或其他实现
type mockChatModel struct{}

func (m *mockChatModel) Generate(ctx context.Context, messages []*schema.Message, options ...model.Option) (*schema.Message, error) {
	return &schema.Message{
		Role:    schema.Assistant,
		Content: "模拟响应",
	}, nil
}

func (m *mockChatModel) Stream(ctx context.Context, messages []*schema.Message, options ...model.Option) (*schema.StreamReader[*schema.Message], error) {
	return nil, nil
}

func (m *mockChatModel) BindTools(tools []*schema.ToolInfo) error {
	return nil
}

// AlertAnalysisSystem 告警分析系统
type AlertAnalysisSystem struct {
	manager *ManagerAgent
	pool    *AlarmPool
}

// NewAlertAnalysisSystem 创建告警分析系统
func NewAlertAnalysisSystem(ctx context.Context, chatModel model.ChatModel) (*AlertAnalysisSystem, error) {
	// 创建告警池
	pool := NewAlarmPool(1000, 1*time.Hour)

	// 创建专家 Agents
	experts := make([]*ExpertAgent, 0, 3)

	// 故障诊断专家
	faultExpert, err := NewExpertAgent(ctx, FaultDiagnosisExpert, chatModel)
	if err != nil {
		return nil, fmt.Errorf("failed to create fault expert: %w", err)
	}
	experts = append(experts, faultExpert)

	// 性能分析专家
	perfExpert, err := NewExpertAgent(ctx, PerformanceAnalysisExpert, chatModel)
	if err != nil {
		return nil, fmt.Errorf("failed to create performance expert: %w", err)
	}
	experts = append(experts, perfExpert)

	// 业务影响专家
	businessExpert, err := NewExpertAgent(ctx, BusinessImpactExpert, chatModel)
	if err != nil {
		return nil, fmt.Errorf("failed to create business expert: %w", err)
	}
	experts = append(experts, businessExpert)

	// 打印每个专家具备的技能
	fmt.Println("\n========== 专家技能配置 ==========")
	for _, expert := range experts {
		expert.PrintSkills()
	}
	fmt.Println("====================================")

	// 创建 Manager Agent
	manager, err := NewManagerAgent(ctx, chatModel, experts, 3) // 最多3轮迭代
	if err != nil {
		return nil, fmt.Errorf("failed to create manager: %w", err)
	}

	return &AlertAnalysisSystem{
		manager: manager,
		pool:    pool,
	}, nil
}

// Analyze 执行告警分析
func (s *AlertAnalysisSystem) Analyze(ctx context.Context, alarms []Alarm) (*FinalReport, error) {
	fmt.Printf("\n========== 开始告警分析 ==========\n")
	fmt.Printf("告警数量: %d\n\n", len(alarms))

	// 1. 场景识别与计划生成
	fmt.Printf("[步骤 1] Manager 进行场景识别和计划生成\n")
	plan, err := s.manager.ScenarioRecognition(ctx, alarms)
	if err != nil {
		return nil, fmt.Errorf("scenario recognition failed: %w", err)
	}
	fmt.Printf("场景: %s, 优先级: %s, 预计轮次: %d\n",
		plan.Scenario, plan.Priority, plan.EstimatedRounds)
	fmt.Printf("策略: %s\n\n", plan.Strategy)

	// 2. 多轮迭代分析
	var iterations []*IterationResult
	analysisContext := AnalysisContext{
		Alarms:    alarms,
		Iteration: 1,
		MaxRounds: plan.EstimatedRounds,
		PrevResults: []string{},
	}

	for i := 1; i <= plan.EstimatedRounds; i++ {
		fmt.Printf("\n[步骤 2.%d] 第 %d 轮迭代分析\n", i, i)
		analysisContext.Iteration = i

		iteration, err := s.manager.AnalyzeIteration(ctx, analysisContext)
		if err != nil {
			return nil, fmt.Errorf("iteration %d failed: %w", i, err)
		}
		iterations = append(iterations, iteration)

		// 打印专家结果
		for _, result := range iteration.ExpertResults {
			if result != nil {
				fmt.Printf("  - %s: 置信度 %.2f\n", result.ExpertName, result.Confidence)
			}
		}

		// 打印决策
		fmt.Printf("  - Manager 决策: %s\n", iteration.Decision.Decision)
		fmt.Printf("  - 理由: %s\n", iteration.Decision.Reasoning)

		// 如果不需要继续分析，提前结束
		if !iteration.Decision.ContinueAnalysis {
			fmt.Printf("  → 提前结束迭代\n")
			break
		}

		// 准备下一轮上下文
		analysisContext.PrevResults = append(analysisContext.PrevResults,
			fmt.Sprintf("第%d轮: %s", i, iteration.Decision.Decision))
	}

	// 3. 生成最终报告
	fmt.Printf("\n[步骤 3] Manager 生成最终报告\n")
	report, err := s.manager.GenerateReport(ctx, &analysisContext, iterations)
	if err != nil {
		return nil, fmt.Errorf("report generation failed: %w", err)
	}

	fmt.Printf("\n========== 分析完成 ==========\n")
	return report, nil
}

func main() {
	ctx := context.Background()

	// 创建 Chat Model (使用模拟实现，实际应该使用 openai.NewChatModel)
	chatModel := &mockChatModel{}

	// 演示技能库使用
	fmt.Println("\n========== 演示：技能库使用 ==========")
	demonstrateSkills(ctx)
	fmt.Println("======================================\n")

	// 创建告警分析系统
	system, err := NewAlertAnalysisSystem(ctx, chatModel)
	if err != nil {
		log.Fatalf("Failed to create alert analysis system: %v", err)
	}

	// 创建模拟告警
	alarms := CreateMockAlarms()
	fmt.Printf("系统初始化完成，检测到 %d 个告警\n", len(alarms))

	// 执行分析
	report, err := system.Analyze(ctx, alarms)
	if err != nil {
		log.Fatalf("Failed to analyze alerts: %v", err)
	}

	// 打印报告
	fmt.Printf("\n========== 最终报告 ==========\n")
	fmt.Printf("【总结】\n%s\n\n", report.Summary)
	fmt.Printf("【根本原因】\n%s\n\n", report.RootCause)
	fmt.Printf("【业务影响】\n%s\n\n", report.BusinessImpact)
	fmt.Printf("【性能分析】\n%s\n\n", report.PerformanceAnalysis)
	fmt.Printf("【处理优先级】%s\n\n", report.Priority)
	fmt.Printf("【建议措施】\n")
	for i, rec := range report.Recommendations {
		fmt.Printf("  %d. %s\n", i+1, rec)
	}
	fmt.Printf("\n【迭代轮次】%d\n", report.IterationsCompleted)
	fmt.Printf("==============================\n")

	// 示例：如何集成真实的 OpenAI 模型
	fmt.Printf("\n========== 使用说明 ==========\n")
	fmt.Printf("要使用真实的 OpenAI 模型，请：\n")
	fmt.Printf("1. 安装 eino-ext: go get github.com/cloudwego/eino-ext/components/model/openai\n")
	fmt.Printf("2. 替换 mockChatModel 为:\n")
	fmt.Printf(`
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		Model:  "gpt-4",
		APIKey: os.Getenv("OPENAI_API_KEY"),
	})
`)
	fmt.Printf("3. 取消注释代码中的 TODO 部分，启用实际的 ChatModel 调用\n")
	fmt.Printf("==============================\n")
}

// demonstrateSkills 演示技能库使用
func demonstrateSkills(ctx context.Context) {
	// 创建技能注册表
	registry := NewSkillRegistry()

	// 打印所有可用技能
	registry.PrintSkills()

	// 演示几个核心技能的使用
	fmt.Println("\n========== 技能执行演示 ==========")

	// 1. 日志查询技能
	logSkill, _ := registry.Get(LogQuerySkill)
	result, _ := logSkill.Execute(ctx, nil)
	fmt.Printf("\n[%s]\n", logSkill.GetName())
	fmt.Printf("  结果: %s\n", result.Message)
	fmt.Printf("  查询到日志数: %v\n", result.Data["logCount"])
	fmt.Printf("  错误日志数: %v\n", result.Data["errorLogs"])

	// 2. 监控指标查询技能
	metricSkill, _ := registry.Get(MetricQuerySkill)
	result, _ = metricSkill.Execute(ctx, nil)
	fmt.Printf("\n[%s]\n", metricSkill.GetName())
	fmt.Printf("  结果: %s\n", result.Message)
	if cpu, ok := result.Data["cpu"].(map[string]any); ok {
		fmt.Printf("  CPU 使用率: %.1f%%\n", cpu["usage"])
	}

	// 3. 时间序列分析技能
	tsSkill, _ := registry.Get(TimeSeriesAnalysisSkill)
	result, _ = tsSkill.Execute(ctx, nil)
	fmt.Printf("\n[%s]\n", tsSkill.GetName())
	fmt.Printf("  结果: %s\n", result.Message)
	fmt.Printf("  趋势: %v\n", result.Data["trend"])

	// 4. 关联分析技能
	corrSkill, _ := registry.Get(CorrelationAnalysisSkill)
	result, _ = corrSkill.Execute(ctx, nil)
	fmt.Printf("\n[%s]\n", corrSkill.GetName())
	fmt.Printf("  结果: %s\n", result.Message)

	// 5. 慢查询分析技能
	slowQuerySkill, _ := registry.Get(SlowQueryAnalysisSkill)
	result, _ = slowQuerySkill.Execute(ctx, nil)
	fmt.Printf("\n[%s]\n", slowQuerySkill.GetName())
	fmt.Printf("  结果: %s\n", result.Message)
	if queries, ok := result.Data["slowQueries"].([]map[string]any); ok {
		fmt.Printf("  发现慢查询数: %d\n", len(queries))
	}

	// 6. 历史案例匹配技能
	historySkill, _ := registry.Get(HistoryMatchSkill)
	result, _ = historySkill.Execute(ctx, nil)
	fmt.Printf("\n[%s]\n", historySkill.GetName())
	fmt.Printf("  结果: %s\n", result.Message)
	if cases, ok := result.Data["similarCases"].([]map[string]any); ok {
		fmt.Printf("  相似案例数: %d\n", len(cases))
	}

	// 7. 根因分析技能
	rcSkill, _ := registry.Get(RootCauseAnalysisSkill)
	result, _ = rcSkill.Execute(ctx, nil)
	fmt.Printf("\n[%s]\n", rcSkill.GetName())
	fmt.Printf("  结果: %s\n", result.Message)
	fmt.Printf("  主要根因: %v\n", result.Data["primaryRootCause"])

	fmt.Println("\n==================================")
}
