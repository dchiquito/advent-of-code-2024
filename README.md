# advent-of-code-2024

# Usage

## Pulling data from adventofcode.com
To save your puzzle inputs locally, first log in to [adventofcode.com](adventofcode.com), inspect the page in your browser, locate the `session` cookie, and save the value in a file called `.session`. This file is used by the CLI to authenticate with `adventofcode.com`.

Now run `go run ./cmd/pull 1` to pull the input for day `1` and save it in `data/01.txt`.

I set my `GOBIN` to `./bin`, which enables me to run `go install ./cmd/pull` to build the pull script and run it with `bin/pull 1` instead.

## Running puzzles
With `GOBIN` set to `./bin`, run `go install ./cmd/*` to rebuild all currently implemented solutions, or `go install ./cmd/day01` to only build a specific day.

TODO utility for running part1/part2
