package repo

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	root, err := ioutil.TempDir("", "TestRepo")
	require.Nil(t, err)
	defer os.RemoveAll(root)

	err = Initialize("", "Secp256k1")
	require.NotNil(t, err)
	err = Initialize(root, "Secp256k1")
	require.Nil(t, err)

	_, err = LoadPrivateKey(root)
	require.Nil(t, err)

	_, err = LoadNodePrivateKey("t")
	require.NotNil(t, err)
	_, err = LoadNodePrivateKey(root)
	require.Nil(t, err)

	pathRoot, err := PathRoot()
	require.Nil(t, err)
	homeRoot, err := homedir.Expand(DefaultPathRoot)
	require.Nil(t, err)
	require.Equal(t, homeRoot, pathRoot)

	rootWithDefault, err := PathRootWithDefault(root)
	require.Nil(t, err)
	require.Equal(t, root, rootWithDefault)

	err = InitConfig(filepath.Join(root, "pier.toml"))
	require.Nil(t, err)

	keyPath := KeyPath(root)
	require.Equal(t, filepath.Join(root, KeyName), keyPath)

	nodeKeyPath := NodeKeyPath(root)
	require.Equal(t, filepath.Join(root, NodeKeyName), nodeKeyPath)

	_, err = GetAPI("")
	require.NotNil(t, err)
	_, err = GetAPI(root)
	require.Nil(t, err)

	_, err = PluginPath()
	require.Nil(t, err)
}

func TestSetPath(t *testing.T) {
	root, err := ioutil.TempDir("", "TestRepo")
	require.Nil(t, err)
	defer os.RemoveAll(root)
	err = Initialize(root, "Secp256k1")
	require.Nil(t, err)

	SetPath(root)
	path, err := PathRoot()
	require.Nil(t, err)
	require.Equal(t, root, path)
	path, err = PathRootWithDefault("")
	require.Nil(t, err)
	require.Equal(t, root, path)
}
