package cli_test

import (
	"os"
	"testing"

	"github.com/churmd/wise-mars-rover/cli"
	"github.com/rhysd/go-fakeio"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	fakeStdin := fakeio.Stdin("4 8\n2 3 E LFRFF\n0 2 N FFLFRFF\n2 3 N FLLFR\n1 0 S FFRLF\n\n")
	defer fakeStdin.Restore()
	fakeStdout := fakeio.Stdout()
	defer fakeStdout.Restore()

	expectedOutput := "Enter the size of the grid, e.g. 4 8\n"
	expectedOutput += "Enter a rover and list of commands, e.g. 2 3 E LFRFF. Or an empty line to continue\n"
	expectedOutput += "Enter a rover and list of commands, e.g. 2 3 E LFRFF. Or an empty line to continue\n"
	expectedOutput += "Enter a rover and list of commands, e.g. 2 3 E LFRFF. Or an empty line to continue\n"
	expectedOutput += "Enter a rover and list of commands, e.g. 2 3 E LFRFF. Or an empty line to continue\n"
	expectedOutput += "Enter a rover and list of commands, e.g. 2 3 E LFRFF. Or an empty line to continue\n"
	expectedOutput += "Final Rover States\n"
	expectedOutput += "(4, 4, E)\n"
	expectedOutput += "(0, 4, W) LOST\n"
	expectedOutput += "(2, 3, W)\n"
	expectedOutput += "(1, 0, S) LOST\n"

	interactor := cli.New(os.Stdin, os.Stdout)
	interactor.Run()

	output, err := fakeStdout.String()
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, output)
}
