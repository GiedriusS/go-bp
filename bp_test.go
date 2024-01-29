package bp

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func compress32WithVanillaImpl(t testing.TB, data []uint32) []byte {
	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })

	td := t.TempDir()

	f, err := os.Create(filepath.Join(td, "input.bin"))
	require.NoError(t, err)

	t.Cleanup(func() { f.Close() })

	bw := bufio.NewWriter(f)
	for _, v := range data {
		fmt.Fprintf(bw, "%d\n", v)
	}
	require.NoError(t, bw.Flush())

	c := exec.Command("./originalimpl", filepath.Join(td, "input.bin"), filepath.Join(td, "output.bin"))
	out, err := c.Output()
	require.NoError(t, err)
	t.Logf("originalimpl: %s", out)

	outContent, err := os.ReadFile(filepath.Join(td, "output.bin"))
	require.NoError(t, err)

	return outContent
}

func TestUnder128Decompression(t *testing.T) {
	in := []uint32{}

	const start = 1500000000
	for i := start; i < start+129; i++ {
		in = append(in, uint32(i))
	}
	o := compress32WithVanillaImpl(t, in)
	t.Log(o)
	require.Equal(t, in, DecompressUnder128(o))
}
