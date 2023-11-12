// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package webhook

import "code.gitea.io/gitea/modules/util"

// HookEvents is a set of web hook events
type HookEvents struct {
	Create                   bool `json:"create"`
	Delete                   bool `json:"delete"`
	Fork                     bool `json:"fork"`
	Issues                   bool `json:"issues"`
	IssueAssign              bool `json:"issue_assign"`
	IssueLabel               bool `json:"issue_label"`
	IssueMilestone           bool `json:"issue_milestone"`
	IssueComment             bool `json:"issue_comment"`
	Push                     bool `json:"push"`
	PullRequest              bool `json:"pull_request"`
	PullRequestAssign        bool `json:"pull_request_assign"`
	PullRequestLabel         bool `json:"pull_request_label"`
	PullRequestMilestone     bool `json:"pull_request_milestone"`
	PullRequestComment       bool `json:"pull_request_comment"`
	PullRequestReview        bool `json:"pull_request_review"`
	PullRequestSync          bool `json:"pull_request_sync"`
	PullRequestReviewRequest bool `json:"pull_request_review_request"`
	Wiki                     bool `json:"wiki"`
	Repository               bool `json:"repository"`
	Release                  bool `json:"release"`
	Package                  bool `json:"package"`
}

// ParseHookEvents converts a list of strings to HookEvents
func ParseHookEvents(eventTypes []string) HookEvents {
	caseInsensitive := true
	return HookEvents{
		Create:                   util.SliceContainsString(eventTypes, HookEventCreate.Event(), caseInsensitive),
		Delete:                   util.SliceContainsString(eventTypes, HookEventDelete.Event(), caseInsensitive),
		Fork:                     util.SliceContainsString(eventTypes, HookEventFork.Event(), caseInsensitive),
		Issues:                   issuesHook(eventTypes, "issues_only", caseInsensitive),
		IssueAssign:              issuesHook(eventTypes, HookEventIssueAssign.Event(), caseInsensitive),
		IssueLabel:               issuesHook(eventTypes, HookEventIssueLabel.Event(), caseInsensitive),
		IssueMilestone:           issuesHook(eventTypes, HookEventIssueMilestone.Event(), caseInsensitive),
		IssueComment:             issuesHook(eventTypes, HookEventIssueComment.Event(), caseInsensitive),
		Push:                     util.SliceContainsString(eventTypes, HookEventPush.Event(), caseInsensitive),
		PullRequest:              pullHook(eventTypes, "pull_request_only", caseInsensitive),
		PullRequestAssign:        pullHook(eventTypes, HookEventPullRequestAssign.Event(), caseInsensitive),
		PullRequestLabel:         pullHook(eventTypes, HookEventPullRequestLabel.Event(), caseInsensitive),
		PullRequestMilestone:     pullHook(eventTypes, HookEventPullRequestMilestone.Event(), caseInsensitive),
		PullRequestComment:       pullHook(eventTypes, HookEventPullRequestComment.Event(), caseInsensitive),
		PullRequestReview:        pullHook(eventTypes, "pull_request_review", caseInsensitive),
		PullRequestReviewRequest: pullHook(eventTypes, HookEventPullRequestReviewRequest.Event(), caseInsensitive),
		PullRequestSync:          pullHook(eventTypes, HookEventPullRequestSync.Event(), caseInsensitive),
		Wiki:                     util.SliceContainsString(eventTypes, HookEventWiki.Event(), caseInsensitive),
		Repository:               util.SliceContainsString(eventTypes, HookEventRepository.Event(), caseInsensitive),
		Release:                  util.SliceContainsString(eventTypes, HookEventRelease.Event(), caseInsensitive),
	}
}

func issuesHook(events []string, event string, caseInsensitive bool) bool {
	return util.SliceContainsString(events, event, caseInsensitive) || util.SliceContainsString(events, HookEventIssues.Event(), caseInsensitive)
}

func pullHook(events []string, event string, caseInsensitive bool) bool {
	return util.SliceContainsString(events, event, caseInsensitive) || util.SliceContainsString(events, HookEventPullRequest.Event(), caseInsensitive)
}

// HookEvent represents events that will delivery hook.
type HookEvent struct {
	PushOnly       bool   `json:"push_only"`
	SendEverything bool   `json:"send_everything"`
	ChooseEvents   bool   `json:"choose_events"`
	BranchFilter   string `json:"branch_filter"`

	HookEvents `json:"events"`
}
