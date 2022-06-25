#!/usr/bin/env sh

set -e
set -x

# Set docker-compose project dirname
project="simulator-integration-tests"

# Extend timeout for pulling large dependencies
export COMPOSE_HTTP_TIMEOUT=200

if [ "$2" = "debug" ]; then
	DEBUG="true"
fi

cd "$(dirname "${0}")"

docker-compose -p "$project" build --parallel

cleanup() {
	if [ $DEBUG ]; then
		echo "Press Enter when you want the environment cleaned up: "
		read -r _
	fi
	# Clean up directory
	docker-compose -p "$project" logs tests
	docker-compose -p "$project" down --volumes
}
trap cleanup EXIT

docker-compose -p "$project" up -d nats simulator
docker-compose -p "$project" up --no-deps tests
exit_code=$(docker inspect ${project}_tests_1 --format='{{.State.ExitCode}}')

if [ $exit_code -ne 0 ]; then
	exit $exit_code
fi
