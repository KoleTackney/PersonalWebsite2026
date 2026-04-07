package web

import (
	"bytes"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"PersonalWebsite2026/content"

	"github.com/yuin/goldmark"
)

type Post struct {
	Slug        string
	Title       string
	Date        time.Time
	DateDisplay string
	Description string
	Tags        []string
	HTML        string
}

var posts []Post
var postsBySlug map[string]Post

func init() {
	postsBySlug = make(map[string]Post)
	md := goldmark.New()

	entries, err := fs.ReadDir(content.BlogFiles, "blog")
	if err != nil {
		log.Printf("Warning: could not read blog directory: %v", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".md" {
			continue
		}

		data, err := fs.ReadFile(content.BlogFiles, "blog/"+entry.Name())
		if err != nil {
			log.Printf("Warning: could not read %s: %v", entry.Name(), err)
			continue
		}

		post := parsePost(entry.Name(), data, md)
		posts = append(posts, post)
		postsBySlug[post.Slug] = post
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})
}

func parsePost(filename string, data []byte, md goldmark.Markdown) Post {
	slug := strings.TrimSuffix(filename, ".md")
	post := Post{Slug: slug}

	content := string(data)

	// Parse YAML front matter
	if strings.HasPrefix(content, "---") {
		parts := strings.SplitN(content[3:], "---", 2)
		if len(parts) == 2 {
			parseFrontMatter(parts[0], &post)
			content = strings.TrimSpace(parts[1])
		}
	}

	// Render markdown to HTML
	var buf bytes.Buffer
	if err := md.Convert([]byte(content), &buf); err != nil {
		log.Printf("Warning: markdown conversion failed for %s: %v", filename, err)
	}
	post.HTML = buf.String()

	return post
}

func parseFrontMatter(fm string, post *Post) {
	for _, line := range strings.Split(fm, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		key, value, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}

		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)
		value = strings.Trim(value, "\"")

		switch key {
		case "title":
			post.Title = value
		case "date":
			if t, err := time.Parse("2006-01-02", value); err == nil {
				post.Date = t
				post.DateDisplay = t.Format("January 2, 2006")
			}
		case "description":
			post.Description = value
		case "tags":
			value = strings.Trim(value, "[]")
			for _, tag := range strings.Split(value, ",") {
				tag = strings.TrimSpace(tag)
				tag = strings.Trim(tag, "\"")
				if tag != "" {
					post.Tags = append(post.Tags, tag)
				}
			}
		}
	}
}

func BlogListHandler(w http.ResponseWriter, r *http.Request) {
	component := BlogListPage(posts)
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error rendering blog list: %v", err)
	}
}

func BlogPostHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	post, ok := postsBySlug[slug]
	if !ok {
		http.NotFound(w, r)
		return
	}

	component := BlogPostPage(post)
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error rendering blog post: %v", err)
	}
}
