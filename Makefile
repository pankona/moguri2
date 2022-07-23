ACTION_NUM ?= -1

.PHONY: intall-tools
install-tools:
	go install github.com/cosmtrek/air@latest

# Convenience commands to test APIs
.PHONY: current-interaction
current-interaction:
	@curl -X GET http://localhost:3000/current_interaction

.PHONY: interact
interact:
	@curl -X POST http://localhost:3000/interact -d '{"action_num": $(ACTION_NUM)}'
