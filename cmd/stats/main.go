package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/slack-go/slack"
)

// TODO: 全体的に名前付けがださい。簡潔で直感的なものにする。
type StatsResponse []struct {
	Date  string `json:"date"`
	Stats []struct {
		Metrics struct {
			Blocks           int `json:"blocks"`
			BounceDrops      int `json:"bounce_drops"`
			Bounces          int `json:"bounces"`
			Clicks           int `json:"clicks"`
			Deferred         int `json:"deferred"`
			Delivered        int `json:"delivered"`
			InvalidEmails    int `json:"invalid_emails"`
			Opens            int `json:"opens"`
			Processed        int `json:"processed"`
			Requests         int `json:"requests"`
			SpamReportDrops  int `json:"spam_report_drops"`
			SpamReports      int `json:"spam_reports"`
			UniqueClicks     int `json:"unique_clicks"`
			UniqueOpens      int `json:"unique_opens"`
			UnsubscribeDrops int `json:"unsubscribe_drops"`
			Unsubscribes     int `json:"unsubscribes"`
		} `json:"metrics"`
	} `json:"stats"`
}

type Stats struct {
	BounceRate     float64
	SpamReportRate float64
}

var (
	apiKey    string
	apiHost   string
	channelID string

	slackCli *slack.Client
)

func init() {
	channelID = os.Getenv("SLACK_CHANNEL_ID")
	apiKey = os.Getenv("SENDGRID_API_KEY")

	oauthToken := os.Getenv("SLACK_OAUTH_TOKEN")
	slackCli = slack.New(oauthToken)
}

func main() {
	ctx := context.Background()
	err := SendDailyStatsReport(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

// TODO: 長い。もっとわける。
func SendDailyStatsReport(ctx context.Context) error {
	apiPath := "/v3/stats"
	apiHost = "https://api.sendgrid.com"
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	req := sendgrid.GetRequest(
		apiKey,
		apiPath,
		apiHost,
	)
	req.Method = "GET"
	req.QueryParams = map[string]string{
		"start_date":    yesterday,
		"end_date":      yesterday,
		"aggregated_by": "day",
		"limit":         "1",
		"offset":        "1",
	}

	// SendgridAPIリクエスト実行
	resp, err := sendgrid.API(req)
	if err != nil {
		return errors.WithStack(err)
	}

	// JSONレスポンスのDecode
	var data StatsResponse
	err = json.Unmarshal([]byte(resp.Body), &data)
	if err != nil {
		return errors.WithStack(err)
	}

	// 必要要素の抽出
	m := data[0].Stats[0].Metrics
	// TODO: 計算式が汚い
	stats := Stats{
		BounceRate:     float64(m.Bounces) / float64(m.Requests),
		SpamReportRate: float64(m.SpamReports) / float64(m.Requests),
	}

	// Slack送信
	if err := slackPost(ctx, stats); err != nil {
		log.Printf("failed to post slack: %+v", err)
		return errors.WithStack(err)
	}

	return nil
}

func severity(s Stats) string {
	var (
		bounceRateThreshold     = 5.0
		spamReportRateThreshold = 0.1
	)

	switch {
	case s.BounceRate >= bounceRateThreshold || s.SpamReportRate >= spamReportRateThreshold:
		return "danger"
	default:
		return "good"
	}
}

func slackPost(ctx context.Context, stats Stats) error {
	attachment := slack.Attachment{
		Title: "SendGridメール送信状況(日次)",
		Color: severity(stats),
		Fields: []slack.AttachmentField{
			{
				Title: "不達率",
				Value: fmt.Sprintf("%.1f %%", stats.BounceRate),
			},
			{
				Title: "スパムメール報告発生率",
				Value: fmt.Sprintf("%.1f %%", stats.SpamReportRate),
			},
		},
		Footer: "from SendGrid API",
	}
	_, _, _, err := slackCli.SendMessageContext(ctx, channelID, slack.MsgOptionAttachments(attachment))

	return err
}
