// Copyright 2017 HootSuite Media Inc.
//
// Licensed under the Apache License, Version 2.0 (the License);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an AS IS BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Modified hereafter by contributors to runatlantis/atlantis.

package vcs

import (
	"fmt"

	"github.com/runatlantis/atlantis/server/events/models"
)

// NotConfiguredVCSClient is used as a placeholder when Atlantis isn't configured
// on startup to support a certain VCS host. For example, if there is no GitHub
// config then this client will be used which will error if it's ever called.
type NotConfiguredVCSClient struct {
	Host models.VCSHostType
}

func (a *NotConfiguredVCSClient) GetModifiedFiles(_ models.Repo, _ models.PullRequest) ([]string, error) {
	return nil, a.err()
}
func (a *NotConfiguredVCSClient) CreateComment(_ models.Repo, _ int, _ string, _ string) error {
	return a.err()
}
func (a *NotConfiguredVCSClient) HidePrevCommandComments(_ models.Repo, _ int, _ string) error {
	return nil
}
func (a *NotConfiguredVCSClient) ReactToComment(repo models.Repo, pullNum int, commentID int64, reaction string) error { // nolint: revive
	return nil
}
func (a *NotConfiguredVCSClient) PullIsApproved(_ models.Repo, _ models.PullRequest) (models.ApprovalStatus, error) {
	return models.ApprovalStatus{}, a.err()
}
func (a *NotConfiguredVCSClient) DiscardReviews(_ models.Repo, _ models.PullRequest) error {
	return nil
}
func (a *NotConfiguredVCSClient) PullIsMergeable(_ models.Repo, _ models.PullRequest, _ string) (bool, error) {
	return false, a.err()
}
func (a *NotConfiguredVCSClient) UpdateStatus(_ models.Repo, _ models.PullRequest, _ models.CommitStatus, _ string, _ string, _ string) error {
	return a.err()
}
func (a *NotConfiguredVCSClient) MergePull(_ models.PullRequest, _ models.PullRequestOptions) error {
	return a.err()
}
func (a *NotConfiguredVCSClient) MarkdownPullLink(_ models.PullRequest) (string, error) {
	return "", a.err()
}
func (a *NotConfiguredVCSClient) err() error {
	return fmt.Errorf("atlantis was not configured to support repos from %s", a.Host.String())
}
func (a *NotConfiguredVCSClient) GetTeamNamesForUser(_ models.Repo, _ models.User) ([]string, error) {
	return nil, a.err()
}

func (a *NotConfiguredVCSClient) SupportsSingleFileDownload(_ models.Repo) bool {
	return false
}

func (a *NotConfiguredVCSClient) GetFileContent(_ models.PullRequest, _ string) (bool, []byte, error) {
	return true, []byte{}, a.err()
}
func (a *NotConfiguredVCSClient) GetCloneURL(_ models.VCSHostType, _ string) (string, error) {
	return "", a.err()
}

func (a *NotConfiguredVCSClient) GetPullLabels(_ models.Repo, _ models.PullRequest) ([]string, error) {
	return nil, a.err()
}

func (a *NotConfiguredVCSClient) AddPullLabel(repo models.Repo, pull models.PullRequest, label string) error {
	return a.err()
}
