#Go profiling workshop

The objective of this workshop is to learn how to identify bottlenecks in go binaries.

##Profiling through tests

One of the tools Go provides for code profiling is its testing framework. A test or benchmark can be specifically compiled to be profiled.

'When CPU profiling is enabled, the Go program stops about 100 times per second and records a sample consisting of the program counters on the currently executing goroutine's stack.'

Read go's full [blog entry](http://blog.golang.org/profiling-go-programs) on profiling.

The following example generates a CPU profile of one benchmark in this package and writes the output into the "prof" file.

`go test -run=none -bench=BenchmarkLinearSortNoBuffer -cpuprofile=prof github.com/jrlangford/go-profiling-workshop`

A memory profile can be generated too by substituting the -cpuprofile param for -memprofile

##PProf

Go includes this tool to analyze profile data. It can be run as an interactive shell.

To execute pprof run the following command:
`go tool pprof [opts] <binary> <profile_data>`

The commands it supports include the following:

###top
List the functions that appeared the most during sampling

According to go's profiling blog entry 'In the `go tool pprof` output, there is a row for each function that appeared in a sample. The first two columns show the number of samples in which the function was running (as opposed to waiting for a called function to return), as a raw count and as a percentage of total samples. ... The fourth and fifth columns show the number of samples in which the function appeared (either running or waiting for a called function to return).'

###list 
`list <funcname>` shows the code section where 'funcname' appears along with the time each line in that section appeared during sampling

###web
Generates a graph in SVG format representing function calls and displays it in a browser. The "Graphviz" package is required for it to work.

###weblist
Displays the output of list in an interactive web page
