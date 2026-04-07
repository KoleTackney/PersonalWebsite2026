---
title: "Building This Site with Go, Templ, and HTMX"
date: 2026-04-06
description: "Why I rebuilt my portfolio from Next.js to a Go server-rendered stack — and what I learned along the way."
tags: ["go", "templ", "htmx", "tailwind"]
---

## Why Rebuild?

My old portfolio was a Next.js static export. It worked, but it felt like shipping a tank to deliver a letter. For a personal site with mostly static content, I wanted something leaner.

Go + Templ + HTMX gives me:

- **Type-safe templates** that compile to Go code
- **Server-side rendering** with no client-side framework
- **HTMX** for the few interactive bits, without shipping a JS bundle
- A **single binary** I can deploy anywhere

## The Stack

The site is built with:

- **Go 1.25** with the standard library `net/http` router
- **Templ** for HTML components
- **HTMX** for dynamic interactions
- **Tailwind CSS v4** compiled via the standalone CLI

Blog posts are markdown files embedded into the binary at build time and parsed with Goldmark.

## What's Next

I'm planning to write more about the robotics work I've been doing at Muskoka Cabinets, and some thoughts on TypeScript vs Go developer experience from building SousChef.

Stay tuned.
