#!/bin/bash

function testAndProfile {
	BENCH=${1:-none}
	if [ "$2" == "-m" ]
	then
		PTYPE="-memprofile"
	else
		PTYPE="-cpuprofile"
	fi
	echo "Running $PTYPE on $1"
	go test -run=none -bench=$BENCH $PTYPE=prof && \
	go tool pprof go-profiling-workshop.test prof
}

case "$1" in
        1)
		testAndProfile ^BenchmarkLinearSort$ $2
		;;

        2)
		testAndProfile ^BenchmarkQuicksort$ $2
		;;

        3)
		testAndProfile ^BenchmarkLinearSortNoBuffer$ $2
		;;

        4)
		testAndProfile ^BenchmarkLinearSortExplorationBuffer$ $2
		;;

	*)
		echo "Usage $0 {1|2|3|4} [-m|-c]"

esac
