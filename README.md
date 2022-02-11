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
Enter the size of the grid, e.g. 5 9
5 9
Enter a rover and list of commands, e.g. 2 3 E LFRFF. Or an empty line to continue
2 3 E LFRFF
Enter a rover and list of commands, e.g. 2 3 E LFRFF. Or an empty line to continue

Final Rover States
(4, 4, E)
```

## Continuing Improvements

-   Move CLI interaction from the main file into another package to separate concerns and follow hexagonal architecture
-   Properly implement CLI parsing of rover and commands to meet the specification given
-   Rename movement package to something that makes a bit more sense
