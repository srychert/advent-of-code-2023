.PHONY: all

all: $(addprefix day, $(shell seq 1 25))

$(addprefix day, $(shell seq 1 25)):
	go run . -day $(subst day,,$@)