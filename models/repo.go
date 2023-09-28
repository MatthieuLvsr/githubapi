package models

type Repos struct {
	// Id int `json: "id"`
	Name string `json: "name"`
	Html_url string `json: "html_url"`
	Description string `json: "description"`
	// Fork bool `json: "fork"`
	CreatedAt string `json: "created_at"`
	UpdatedAt string `json: "updated_at"`
}

func Tostring(repo Repos)[]string{
	var str []string
	str = append(str, repo.Name)
	str = append(str, repo.Html_url)
	str = append(str, repo.Description)
	str = append(str, repo.CreatedAt)
	str = append(str, repo.UpdatedAt)
	return str
}

// "id": 635232145,
//         "node_id": "R_kgDOJdzfkQ",
//         "name": "BlackJack",
//         "full_name": "MatthieuLvsr/BlackJack",
//         "private": false,
//         "owner": {
//             "login": "MatthieuLvsr",
//             "id": 72619668,
//             "node_id": "MDQ6VXNlcjcyNjE5NjY4",
//             "avatar_url": "https://avatars.githubusercontent.com/u/72619668?v=4",
//             "gravatar_id": "",
//             "url": "https://api.github.com/users/MatthieuLvsr",
//             "html_url": "https://github.com/MatthieuLvsr",
//             "followers_url": "https://api.github.com/users/MatthieuLvsr/followers",
//             "following_url": "https://api.github.com/users/MatthieuLvsr/following{/other_user}",
//             "gists_url": "https://api.github.com/users/MatthieuLvsr/gists{/gist_id}",
//             "starred_url": "https://api.github.com/users/MatthieuLvsr/starred{/owner}{/repo}",
//             "subscriptions_url": "https://api.github.com/users/MatthieuLvsr/subscriptions",
//             "organizations_url": "https://api.github.com/users/MatthieuLvsr/orgs",
//             "repos_url": "https://api.github.com/users/MatthieuLvsr/repos",
//             "events_url": "https://api.github.com/users/MatthieuLvsr/events{/privacy}",
//             "received_events_url": "https://api.github.com/users/MatthieuLvsr/received_events",
//             "type": "User",
//             "site_admin": false
//         },
//         "html_url": "https://github.com/MatthieuLvsr/BlackJack",
//         "description": null,
//         "fork": false,
//         "url": "https://api.github.com/repos/MatthieuLvsr/BlackJack",
//         "forks_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/forks",
//         "keys_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/keys{/key_id}",
//         "collaborators_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/collaborators{/collaborator}",
//         "teams_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/teams",
//         "hooks_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/hooks",
//         "issue_events_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/issues/events{/number}",
//         "events_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/events",
//         "assignees_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/assignees{/user}",
//         "branches_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/branches{/branch}",
//         "tags_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/tags",
//         "blobs_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/git/blobs{/sha}",
//         "git_tags_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/git/tags{/sha}",
//         "git_refs_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/git/refs{/sha}",
//         "trees_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/git/trees{/sha}",
//         "statuses_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/statuses/{sha}",
//         "languages_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/languages",
//         "stargazers_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/stargazers",
//         "contributors_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/contributors",
//         "subscribers_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/subscribers",
//         "subscription_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/subscription",
//         "commits_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/commits{/sha}",
//         "git_commits_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/git/commits{/sha}",
//         "comments_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/comments{/number}",
//         "issue_comment_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/issues/comments{/number}",
//         "contents_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/contents/{+path}",
//         "compare_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/compare/{base}...{head}",
//         "merges_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/merges",
//         "archive_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/{archive_format}{/ref}",
//         "downloads_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/downloads",
//         "issues_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/issues{/number}",
//         "pulls_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/pulls{/number}",
//         "milestones_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/milestones{/number}",
//         "notifications_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/notifications{?since,all,participating}",
//         "labels_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/labels{/name}",
//         "releases_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/releases{/id}",
//         "deployments_url": "https://api.github.com/repos/MatthieuLvsr/BlackJack/deployments",
//         "created_at": "2023-05-02T08:52:33Z",
//         "updated_at": "2023-05-02T08:57:24Z",
//         "pushed_at": "2023-05-02T10:10:49Z",
//         "git_url": "git://github.com/MatthieuLvsr/BlackJack.git",
//         "ssh_url": "git@github.com:MatthieuLvsr/BlackJack.git",
//         "clone_url": "https://github.com/MatthieuLvsr/BlackJack.git",
//         "svn_url": "https://github.com/MatthieuLvsr/BlackJack",
//         "homepage": null,
//         "size": 457,
//         "stargazers_count": 0,
//         "watchers_count": 0,
//         "language": "JavaScript",
//         "has_issues": true,
//         "has_projects": true,
//         "has_downloads": true,
//         "has_wiki": true,
//         "has_pages": false,
//         "has_discussions": false,
//         "forks_count": 0,
//         "mirror_url": null,
//         "archived": false,
//         "disabled": false,
//         "open_issues_count": 0,
//         "license": {
//             "key": "mit",
//             "name": "MIT License",
//             "spdx_id": "MIT",
//             "url": "https://api.github.com/licenses/mit",
//             "node_id": "MDc6TGljZW5zZTEz"
//         },
//         "allow_forking": true,
//         "is_template": false,
//         "web_commit_signoff_required": false,
//         "topics": [],
//         "visibility": "public",
//         "forks": 0,
//         "open_issues": 0,
//         "watchers": 0,
//         "default_branch": "main"
//     },