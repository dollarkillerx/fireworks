package git_models

import "time"

type GitlabWebHook struct {
	ObjectKind   string `json:"object_kind"` // tag_push, push, merge_request
	EventName    string `json:"event_name"`
	Before       string `json:"before"`
	After        string `json:"after"`
	Ref          string `json:"ref"` // refs/tags/v0.0.1-dev,refs/heads/master
	CheckoutSha  string `json:"checkout_sha"`
	Message      string `json:"message"`
	UserId       int    `json:"user_id"`
	UserName     string `json:"user_name"`
	UserUsername string `json:"user_username"`
	UserEmail    string `json:"user_email"`
	UserAvatar   string `json:"user_avatar"`
	ProjectId    int    `json:"project_id"`
	Project      struct {
		Id                int         `json:"id"`
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebUrl            string      `json:"web_url"`
		AvatarUrl         interface{} `json:"avatar_url"`
		GitSshUrl         string      `json:"git_ssh_url"`
		GitHttpUrl        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		CiConfigPath      string      `json:"ci_config_path"`
		Homepage          string      `json:"homepage"`
		Url               string      `json:"url"`
		SshUrl            string      `json:"ssh_url"`
		HttpUrl           string      `json:"http_url"`
	} `json:"project"`
	Commits []struct {
		Id        string    `json:"id"`
		Message   string    `json:"message"`
		Title     string    `json:"title"`
		Timestamp time.Time `json:"timestamp"`
		Url       string    `json:"url"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
		Added    []string      `json:"added"`
		Modified []string      `json:"modified"`
		Removed  []interface{} `json:"removed"`
	} `json:"commits"`
	TotalCommitsCount int `json:"total_commits_count"`
	PushOptions       struct {
	} `json:"push_options"`
	ObjectAttributes *ObjectAttributes `json:"object_attributes"` // mg
	Labels           []interface{}     `json:"labels"`
	Changes          struct {
		StateId struct {
			Previous int `json:"previous"`
			Current  int `json:"current"`
		} `json:"state_id"`
		UpdatedAt struct {
			Previous string `json:"previous"`
			Current  string `json:"current"`
		} `json:"updated_at"`
	} `json:"changes"`
	Repository struct {
		Name            string `json:"name"`
		Url             string `json:"url"`
		Description     string `json:"description"`
		Homepage        string `json:"homepage"`
		GitHttpUrl      string `json:"git_http_url"`
		GitSshUrl       string `json:"git_ssh_url"`
		VisibilityLevel int    `json:"visibility_level"`
	} `json:"repository"`
}

// ObjectAttributes mg
type ObjectAttributes struct {
	AssigneeId     interface{} `json:"assignee_id"`
	AuthorId       int         `json:"author_id"`
	CreatedAt      string      `json:"created_at"`
	Description    string      `json:"description"`
	HeadPipelineId interface{} `json:"head_pipeline_id"`
	Id             int         `json:"id"`
	Iid            int         `json:"iid"`
	LastEditedAt   interface{} `json:"last_edited_at"`
	LastEditedById interface{} `json:"last_edited_by_id"`
	MergeCommitSha string      `json:"merge_commit_sha"`
	MergeError     interface{} `json:"merge_error"`
	MergeParams    struct {
		ForceRemoveSourceBranch string `json:"force_remove_source_branch"`
	} `json:"merge_params"`
	MergeStatus               string      `json:"merge_status"`
	MergeUserId               interface{} `json:"merge_user_id"`
	MergeWhenPipelineSucceeds bool        `json:"merge_when_pipeline_succeeds"`
	MilestoneId               interface{} `json:"milestone_id"`
	SourceBranch              string      `json:"source_branch"` // 来源分支
	SourceProjectId           int         `json:"source_project_id"`
	StateId                   int         `json:"state_id"`
	TargetBranch              string      `json:"target_branch"` // 目标分支
	TargetProjectId           int         `json:"target_project_id"`
	TimeEstimate              int         `json:"time_estimate"`
	Title                     string      `json:"title"`
	UpdatedAt                 string      `json:"updated_at"`
	UpdatedById               interface{} `json:"updated_by_id"`
	Url                       string      `json:"url"`
	Source                    struct {
		Id                int         `json:"id"`
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebUrl            string      `json:"web_url"`
		AvatarUrl         interface{} `json:"avatar_url"`
		GitSshUrl         string      `json:"git_ssh_url"`
		GitHttpUrl        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		CiConfigPath      string      `json:"ci_config_path"`
		Homepage          string      `json:"homepage"`
		Url               string      `json:"url"`
		SshUrl            string      `json:"ssh_url"`
		HttpUrl           string      `json:"http_url"`
	} `json:"source"`
	Target struct {
		Id                int         `json:"id"`
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		WebUrl            string      `json:"web_url"`
		AvatarUrl         interface{} `json:"avatar_url"`
		GitSshUrl         string      `json:"git_ssh_url"`
		GitHttpUrl        string      `json:"git_http_url"`
		Namespace         string      `json:"namespace"`
		VisibilityLevel   int         `json:"visibility_level"`
		PathWithNamespace string      `json:"path_with_namespace"`
		DefaultBranch     string      `json:"default_branch"`
		CiConfigPath      string      `json:"ci_config_path"`
		Homepage          string      `json:"homepage"`
		Url               string      `json:"url"`
		SshUrl            string      `json:"ssh_url"`
		HttpUrl           string      `json:"http_url"`
	} `json:"target"`
	LastCommit struct {
		Id        string    `json:"id"`
		Message   string    `json:"message"`
		Title     string    `json:"title"`
		Timestamp time.Time `json:"timestamp"`
		Url       string    `json:"url"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
	} `json:"last_commit"`
	WorkInProgress      bool          `json:"work_in_progress"`
	TotalTimeSpent      int           `json:"total_time_spent"`
	HumanTotalTimeSpent interface{}   `json:"human_total_time_spent"`
	HumanTimeEstimate   interface{}   `json:"human_time_estimate"`
	AssigneeIds         []interface{} `json:"assignee_ids"`
	State               string        `json:"state"`
	Action              string        `json:"action"`
}

// push after: 0000000000000000000000000000000000000000 del

// mg:

/**
{"object_kind":"merge_request","event_type":"merge_request","user":{"name":"wangye","username":"dollarkiller","avatar_url":"https://gitlab.mvalley.com/uploads/-/system/user/avatar/8/avatar.png","email":"wangy@rimepevc.com"},"project":{"id":755,"name":"Data output","description":"针对数据团队的需求 进行数据导出","web_url":"https://gitlab.mvalley.com/dollarkiller/data-output","avatar_url":null,"git_ssh_url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","git_http_url":"https://gitlab.mvalley.com/dollarkiller/data-output.git","namespace":"wangye","visibility_level":20,"path_with_namespace":"dollarkiller/data-output","default_branch":"master","ci_config_path":"","homepage":"https://gitlab.mvalley.com/dollarkiller/data-output","url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","ssh_url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","http_url":"https://gitlab.mvalley.com/dollarkiller/data-output.git"},"object_attributes":{"assignee_id":null,"author_id":8,"created_at":"2022-03-29 09:04:23 UTC","description":"","head_pipeline_id":null,"id":10929,"iid":1,"last_edited_at":null,"last_edited_by_id":null,"merge_commit_sha":"1d9d5f032769ef143de075d007e246afc8595254","merge_error":null,"merge_params":{"force_remove_source_branch":"1"},"merge_status":"can_be_merged","merge_user_id":null,"merge_when_pipeline_succeeds":false,"milestone_id":null,"source_branch":"test","source_project_id":755,"state_id":3,"target_branch":"master","target_project_id":755,"time_estimate":0,"title":"Test","updated_at":"2022-03-29 09:04:26 UTC","updated_by_id":null,"url":"https://gitlab.mvalley.com/dollarkiller/data-output/-/merge_requests/1","source":{"id":755,"name":"Data output","description":"针对数据团队的需求 进行数据导出","web_url":"https://gitlab.mvalley.com/dollarkiller/data-output","avatar_url":null,"git_ssh_url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","git_http_url":"https://gitlab.mvalley.com/dollarkiller/data-output.git","namespace":"wangye","visibility_level":20,"path_with_namespace":"dollarkiller/data-output","default_branch":"master","ci_config_path":"","homepage":"https://gitlab.mvalley.com/dollarkiller/data-output","url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","ssh_url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","http_url":"https://gitlab.mvalley.com/dollarkiller/data-output.git"},"target":{"id":755,"name":"Data output","description":"针对数据团队的需求 进行数据导出","web_url":"https://gitlab.mvalley.com/dollarkiller/data-output","avatar_url":null,"git_ssh_url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","git_http_url":"https://gitlab.mvalley.com/dollarkiller/data-output.git","namespace":"wangye","visibility_level":20,"path_with_namespace":"dollarkiller/data-output","default_branch":"master","ci_config_path":"","homepage":"https://gitlab.mvalley.com/dollarkiller/data-output","url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","ssh_url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","http_url":"https://gitlab.mvalley.com/dollarkiller/data-output.git"},"last_commit":{"id":"95ea8ccd7afc70c68f4ca295825442ee0677222b","message":"px\n","title":"px","timestamp":"2022-03-29T17:03:05+08:00","url":"https://gitlab.mvalley.com/dollarkiller/data-output/-/commit/95ea8ccd7afc70c68f4ca295825442ee0677222b","author":{"name":"dollarkillerx","email":"adapawang@gmail.com"}},"work_in_progress":false,"total_time_spent":0,"human_total_time_spent":null,"human_time_estimate":null,"assignee_ids":[],"state":"merged","action":"merge"},"labels":[],"changes":{"state_id":{"previous":4,"current":3},"updated_at":{"previous":"2022-03-29 09:04:26 UTC","current":"2022-03-29 09:04:26 UTC"}},"repository":{"name":"Data output","url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","description":"针对数据团队的需求 进行数据导出","homepage":"https://gitlab.mvalley.com/dollarkiller/data-output"}}
*/

// push

/**
{"object_kind":"push","event_name":"push","before":"ec5dd4e338079d6a3db47c0b1bf7e8d1efc95ef3","after":"1d9d5f032769ef143de075d007e246afc8595254","ref":"refs/heads/master","checkout_sha":"1d9d5f032769ef143de075d007e246afc8595254","message":null,"user_id":8,"user_name":"wangye","user_username":"dollarkiller","user_email":"","user_avatar":"https://gitlab.mvalley.com/uploads/-/system/user/avatar/8/avatar.png","project_id":755,"project":{"id":755,"name":"Data output","description":"针对数据团队的需求 进行数据导出","web_url":"https://gitlab.mvalley.com/dollarkiller/data-output","avatar_url":null,"git_ssh_url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","git_http_url":"https://gitlab.mvalley.com/dollarkiller/data-output.git","namespace":"wangye","visibility_level":20,"path_with_namespace":"dollarkiller/data-output","default_branch":"master","ci_config_path":"","homepage":"https://gitlab.mvalley.com/dollarkiller/data-output","url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","ssh_url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","http_url":"https://gitlab.mvalley.com/dollarkiller/data-output.git"},"commits":[{"id":"389e4ef0aff6c7af05a59166546445824c2ab055","message":"hello world\n","title":"hello world","timestamp":"2022-03-29T17:02:05+08:00","url":"https://gitlab.mvalley.com/dollarkiller/data-output/-/commit/389e4ef0aff6c7af05a59166546445824c2ab055","author":{"name":"dollarkillerx","email":"adapawang@gmail.com"},"added":[],"modified":["README.md"],"removed":[]},{"id":"95ea8ccd7afc70c68f4ca295825442ee0677222b","message":"px\n","title":"px","timestamp":"2022-03-29T17:03:05+08:00","url":"https://gitlab.mvalley.com/dollarkiller/data-output/-/commit/95ea8ccd7afc70c68f4ca295825442ee0677222b","author":{"name":"dollarkillerx","email":"adapawang@gmail.com"},"added":[],"modified":["README.md"],"removed":[]},{"id":"1d9d5f032769ef143de075d007e246afc8595254","message":"Merge branch 'test' into 'master'\n\nTest\n\nSee merge request dollarkiller/data-output!1","title":"Merge branch 'test' into 'master'","timestamp":"2022-03-29T09:04:26+00:00","url":"https://gitlab.mvalley.com/dollarkiller/data-output/-/commit/1d9d5f032769ef143de075d007e246afc8595254","author":{"name":"wangye","email":"wangy@rimepevc.com"},"added":[],"modified":["README.md"],"removed":[]}],"total_commits_count":3,"push_options":{},"repository":{"name":"Data output","url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","description":"针对数据团队的需求 进行数据导出","homepage":"https://gitlab.mvalley.com/dollarkiller/data-output","git_http_url":"https://gitlab.mvalley.com/dollarkiller/data-output.git","git_ssh_url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","visibility_level":20}}

{"object_kind":"push","event_name":"push","before":"0000000000000000000000000000000000000000","after":"0878170886b4857e874ae0533d508d348bc1be7e","ref":"refs/heads/test","checkout_sha":"0878170886b4857e874ae0533d508d348bc1be7e","message":null,"user_id":8,"user_name":"wangye","user_username":"dollarkiller","user_email":"","user_avatar":"https://gitlab.mvalley.com/uploads/-/system/user/avatar/8/avatar.png","project_id":755,"project":{"id":755,"name":"Data output","description":"针对数据团队的需求 进行数据导出","web_url":"https://gitlab.mvalley.com/dollarkiller/data-output","avatar_url":null,"git_ssh_url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","git_http_url":"https://gitlab.mvalley.com/dollarkiller/data-output.git","namespace":"wangye","visibility_level":20,"path_with_namespace":"dollarkiller/data-output","default_branch":"master","ci_config_path":"","homepage":"https://gitlab.mvalley.com/dollarkiller/data-output","url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","ssh_url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","http_url":"https://gitlab.mvalley.com/dollarkiller/data-output.git"},"commits":[{"id":"0878170886b4857e874ae0533d508d348bc1be7e","message":"px\n","title":"px","timestamp":"2022-03-29T17:17:33+08:00","url":"https://gitlab.mvalley.com/dollarkiller/data-output/-/commit/0878170886b4857e874ae0533d508d348bc1be7e","author":{"name":"dollarkillerx","email":"adapawang@gmail.com"},"added":["a.txt"],"modified":[],"removed":[]}],"total_commits_count":1,"push_options":{},"repository":{"name":"Data output","url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","description":"针对数据团队的需求 进行数据导出","homepage":"https://gitlab.mvalley.com/dollarkiller/data-output","git_http_url":"https://gitlab.mvalley.com/dollarkiller/data-output.git","git_ssh_url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","visibility_level":20}}
*/

// del branch

/**
{"object_kind":"push","event_name":"push","before":"95ea8ccd7afc70c68f4ca295825442ee0677222b","after":"0000000000000000000000000000000000000000","ref":"refs/heads/test","checkout_sha":null,"message":null,"user_id":8,"user_name":"wangye","user_username":"dollarkiller","user_email":"","user_avatar":"https://gitlab.mvalley.com/uploads/-/system/user/avatar/8/avatar.png","project_id":755,"project":{"id":755,"name":"Data output","description":"针对数据团队的需求 进行数据导出","web_url":"https://gitlab.mvalley.com/dollarkiller/data-output","avatar_url":null,"git_ssh_url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","git_http_url":"https://gitlab.mvalley.com/dollarkiller/data-output.git","namespace":"wangye","visibility_level":20,"path_with_namespace":"dollarkiller/data-output","default_branch":"master","ci_config_path":"","homepage":"https://gitlab.mvalley.com/dollarkiller/data-output","url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","ssh_url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","http_url":"https://gitlab.mvalley.com/dollarkiller/data-output.git"},"commits":[],"total_commits_count":0,"push_options":{},"repository":{"name":"Data output","url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","description":"针对数据团队的需求 进行数据导出","homepage":"https://gitlab.mvalley.com/dollarkiller/data-output","git_http_url":"https://gitlab.mvalley.com/dollarkiller/data-output.git","git_ssh_url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","visibility_level":20}}
*/

// tag    "ref":"refs/tags/v0.0.2-dev",

/**
{"object_kind":"tag_push","event_name":"tag_push","before":"0000000000000000000000000000000000000000","after":"1d9d5f032769ef143de075d007e246afc8595254","ref":"refs/tags/v0.0.2-dev","checkout_sha":"1d9d5f032769ef143de075d007e246afc8595254","message":"","user_id":8,"user_name":"wangye","user_username":"dollarkiller","user_email":"","user_avatar":"https://gitlab.mvalley.com/uploads/-/system/user/avatar/8/avatar.png","project_id":755,"project":{"id":755,"name":"Data output","description":"针对数据团队的需求 进行数据导出","web_url":"https://gitlab.mvalley.com/dollarkiller/data-output","avatar_url":null,"git_ssh_url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","git_http_url":"https://gitlab.mvalley.com/dollarkiller/data-output.git","namespace":"wangye","visibility_level":20,"path_with_namespace":"dollarkiller/data-output","default_branch":"master","ci_config_path":"","homepage":"https://gitlab.mvalley.com/dollarkiller/data-output","url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","ssh_url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","http_url":"https://gitlab.mvalley.com/dollarkiller/data-output.git"},"commits":[{"id":"1d9d5f032769ef143de075d007e246afc8595254","message":"Merge branch 'test' into 'master'\n\nTest\n\nSee merge request dollarkiller/data-output!1","title":"Merge branch 'test' into 'master'","timestamp":"2022-03-29T09:04:26+00:00","url":"https://gitlab.mvalley.com/dollarkiller/data-output/-/commit/1d9d5f032769ef143de075d007e246afc8595254","author":{"name":"wangye","email":"wangy@rimepevc.com"},"added":[],"modified":["README.md"],"removed":[]}],"total_commits_count":1,"push_options":{},"repository":{"name":"Data output","url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","description":"针对数据团队的需求 进行数据导出","homepage":"https://gitlab.mvalley.com/dollarkiller/data-output","git_http_url":"https://gitlab.mvalley.com/dollarkiller/data-output.git","git_ssh_url":"ssh://git@gitlab.mvalley.com:9022/dollarkiller/data-output.git","visibility_level":20}}
*/
