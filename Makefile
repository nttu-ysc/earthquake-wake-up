CONFIG_FILE := configs/config.yaml
EXAMPLE_CONFIG := configs/config.example.yaml

.PHONY: build config clean

build:
	go build -ldflags="-s -w" -o earthquake-wake-up .

config:
	@if [ ! -f "$(CONFIG_FILE)" ]; then \
        cp "$(EXAMPLE_CONFIG)" "$(CONFIG_FILE)"; \
        echo "Copied $(EXAMPLE_CONFIG) to $(CONFIG_FILE)"; \
    else \
        echo "$(CONFIG_FILE) already exists."; \
    fi

clean:
	rm -f earthquake-wake-up