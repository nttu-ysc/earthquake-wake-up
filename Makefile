build:
	go build -ldflags="-s -w" -o earthquake-wake-up .

config:
	cp configs/config.example.yaml configs/config.yaml

clean:
	rm -f earthquake-wake-up