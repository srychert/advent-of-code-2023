.PHONY: all

all: $(addprefix day, $(shell seq 1 25))

# Pattern rule to match day targets and handle the part and ex arguments
day%:
	@DAY=$(subst day,,$@); \
	PART=$(word 2, $(MAKECMDGOALS)); \
	EX=$(word 3, $(MAKECMDGOALS)); \
	if [ -n "$$EX" ]; then \
		EX=true; \
	else \
		EX=false; \
	fi; \
	$(MAKE) --no-print-directory run DAY=$$DAY PART=$$PART EX=$$EX

run:
	go run . -day $(DAY) -part $(PART) $(if $(filter true,$(EX)),-ex)

# Ignore the part and ex arguments as make targets
%:
	@:
