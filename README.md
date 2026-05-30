## Gophercises: Quiz Game

This exercise is part of [gophercises](https://courses.calhoun.io/courses/cor_gophercises). This covers exercise #1: Quiz Game

This is a simple CLI based quiz game which prompts questions and answers based on the input CSV file.

Build instructions:
```
$ go build quiz.go
```

Usage:
```
$ ./quiz -h
```

Pass custom CSV file (default: problems.csv):
```
$ ./quiz -file /path/to/file
```

Set a timeout for the quiz in seconds (default: 30s):
```
$ ./quiz -limit 20
```

Start quiz:
```
$ ./quiz
```
