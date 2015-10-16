# Autobuild

> A simple build tool that integrates with github

## The Plan

1. Configure It

  - Tell it where your code lives
  - Tell it which branch you want to watch
  - Give it a bash script to run when that branch changes

2. Setup a webhook on github that pings Autobuild when someone does a push

## Expected behaviour

When you push to github, it will automatically pull those changes, and run a
build script.
