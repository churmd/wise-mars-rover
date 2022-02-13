# wise-mars-rover

Write a program that takes in commands and moves one or more robots around the surface of Mars!

## How to run

Through docker:

```
make docker_build
make docker_run
```

If you have Go installed:

```
go run main.go
```

And then follow the prompts printed out to the terminal.

Note: Not properly implemented reading rover and commands, just have each variable space separated without brackets/commas. E.g. use `2 3 E LFRFF` rather than `(2, 3, E) LFRFF`.

Example run:

```
docker run -it mars-rover
Enter the size of the grid, e.g. 4 8
4 8
Enter a rover and list of commands, e.g. 2 3 E LFRFF. Or an empty line to continue
0 2 N FFLFRFF
Enter a rover and list of commands, e.g. 2 3 E LFRFF. Or an empty line to continue

Final Rover States
(0, 4, W) LOST
```

## Continuing Improvements

-   Properly implement CLI parsing of rover and commands to meet the specification given
-   Rename packages to something that make a bit more sense
