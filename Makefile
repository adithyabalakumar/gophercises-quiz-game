.PHONY: build
build: quiz-game

.PHONY: quiz-game
quiz-game:
	go build -o $@ quiz.go
	@echo "$@: build successful"

.PHONY: clean
clean:
	$(RM) quiz-game