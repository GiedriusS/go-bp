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
	o := compress32WithVanillaImpl(t, []uint32{1, 2, 100, 10000})
	require.Equal(t, []uint32{1, 2, 100, 10000}, DecompressUnder128(o))
}
