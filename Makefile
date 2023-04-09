init:
	docker volume remove xn-temporal || echo ""
	docker volume create xn-temporal

build-temporalite:
	docker build -t dev/temporalite -f Temporalite.Dockerfile .

start-temporalite:
	docker kill xn-temporalite  || echo ""
	docker rm xn-temporalite  || echo ""
	docker run -d --name xn-temporalite -v xn-temporal:/app/db -p 8233:8233 -p 7233:7233 dev/temporalite:latest

stop-temporalite:
	docker kill xn-temporalite  || echo ""
	docker rm xn-temporalite  || echo ""