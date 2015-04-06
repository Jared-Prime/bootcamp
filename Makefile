COMMIT_HASH=`git rev-parse --short HEAD 2>/dev/null`
BUILD_DATE=`date +%FT%T%z`

LDFLAGS=-ldflags "-X github.com/Jared-Prime/bootcamp/exercise/exerciselib.CommitHash ${COMMIT_HASH} -X github.com/Jared-Prime/bootcamp/exercise/exerciselib.BuildDate ${BUILD_DATE}"

run-tests:
	go test github.com/Jared-Prime/bootcamp/exercise
