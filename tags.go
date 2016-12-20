package main

import (
	"fmt"
)

const (
	tagsURI = "/repos/%s/%s/tags%s"
)

// Tag ...
type Tag struct {
	Name       string `json:"name"`
	Commit     Commit `json:"commit"`
	ZipBallURL string `json:"zipball_url"`
	TarBallURL string `json:"tarball_url"`
}

// LightweightTag holds the fields required for a lightweight tag.
type LightweightTag struct {
	Ref string `json:"ref"`
	SHA string `json:"sha"`
}

func (t *Tag) String() string {
	return t.Name + " (commit: " + t.Commit.Url + ")"
}

// Tags gets the tags associated with a repo.
func Tags(user, repo, token string) ([]Tag, error) {
	var tags []Tag

	if token != "" {
		token = "?access_token=" + token
	}

	err := GithubGet(fmt.Sprintf(tagsURI, user, repo, token), &tags)
	if err != nil {
		return nil, err
	}

	return tags, nil
}
