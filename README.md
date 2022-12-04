# Advent of Code

My take on [Advent of Code](https://adventofcode.com/).

After giving up half-way through the challenges last year, I'll give it another try this year.

I've chosen to do this using [Go](https://go.dev/), since I am already using it for some while and want to get better at it and learn new aspects of it. That is also the reason why you might find some solutions over-engineered.

Each day is organized in a separate package, all being "forked" from a template.

For 2022 you find a `Makefile` that provides targets to 
- print all solutions: `make run`
- run the tests for specifc day packages: `make test 3` (with 3 being the day number)
- create a new day package from the template: `make new 3` (again, 3 being the day number)

## References

Below are some resources I used throughout my journey.

- Get started: https://go.dev/doc/tutorial/getting-started
- `bufio.Scanner`: https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go#16615559
- Testing basics: https://blog.alexellis.io/golang-writing-unit-tests/
- https://www.digitalocean.com/community/tutorials/understanding-arrays-and-slices-in-go
