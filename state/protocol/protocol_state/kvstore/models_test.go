package kvstore

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/state/protocol/protocol_state"
)

// TestEncodeDecode tests encoding and decoding all supported model versions.
//   - VersionedEncode should return the correct version
//   - instances should be equal after encoding, then decoding
func TestEncodeDecode(t *testing.T) {
	t.Run("v0", func(t *testing.T) {
		model := &modelv0{}

		version, encoded, err := model.VersionedEncode()
		require.NoError(t, err)
		assert.Equal(t, uint64(0), version)

		decoded, err := VersionedDecode(version, encoded)
		require.NoError(t, err)
		assert.Equal(t, model, decoded)
	})

	t.Run("v1", func(t *testing.T) {
		model := &modelv1{
			InvalidEpochTransitionAttempted: rand.Int()%2 == 0,
		}

		version, encoded, err := model.VersionedEncode()
		require.NoError(t, err)
		assert.Equal(t, uint64(1), version)

		decoded, err := VersionedDecode(version, encoded)
		require.NoError(t, err)
		assert.Equal(t, model, decoded)
	})
}

// TestKVStoreAPI tests that all supported model versions satisfy the public interfaces.
//   - should be able to read/write supported keys
//   - should return the appropriate sentinel for unsupported keys
func TestKVStoreAPI(t *testing.T) {
	t.Run("v0", func(t *testing.T) {
		model := &modelv0{}

		assert.True(t, reflect.DeepEqual(model, model.Clone()))

		// v0
		assertModelIsUpgradable(t, model)

		version := model.GetProtocolStateVersion()
		assert.Equal(t, uint64(0), version)

		// v1
		err := model.SetInvalidEpochTransitionAttempted(true)
		assert.ErrorIs(t, err, ErrKeyNotSupported)

		_, err = model.GetInvalidEpochTransitionAttempted()
		assert.ErrorIs(t, err, ErrKeyNotSupported)
	})

	t.Run("v1", func(t *testing.T) {
		model := &modelv1{}

		assert.True(t, reflect.DeepEqual(model, model.Clone()))

		// v0
		assertModelIsUpgradable(t, model)

		version := model.GetProtocolStateVersion()
		assert.Equal(t, uint64(1), version)

		// v1
		err := model.SetInvalidEpochTransitionAttempted(true)
		assert.NoError(t, err)

		invalidEpochTransitionAttempted, err := model.GetInvalidEpochTransitionAttempted()
		assert.NoError(t, err)
		assert.Equal(t, true, invalidEpochTransitionAttempted)
	})
}

// TestKVStoreAPI_Clone tests that cloning of KV store correctly works. All versions need to be support this.
func TestKVStoreAPI_Clone(t *testing.T) {
	t.Run("v0", func(t *testing.T) {
		model := &modelv0{
			upgradableModel: upgradableModel{
				VersionUpgrade: &protocol_state.ViewBasedActivator[uint64]{
					Data:           13,
					ActivationView: 1000,
				},
			},
		}
		cpy := model.Clone()
		require.True(t, reflect.DeepEqual(model, cpy))

		model.VersionUpgrade.ActivationView++ // change
		require.False(t, reflect.DeepEqual(model, cpy))
	})
	t.Run("v1", func(t *testing.T) {
		model := &modelv1{
			upgradableModel: upgradableModel{
				VersionUpgrade: &protocol_state.ViewBasedActivator[uint64]{
					Data:           13,
					ActivationView: 1000,
				},
			},
			InvalidEpochTransitionAttempted: false,
		}
		cpy := model.Clone()
		require.True(t, reflect.DeepEqual(model, cpy))

		model.VersionUpgrade.ActivationView++ // change
		require.False(t, reflect.DeepEqual(model, cpy))
	})
}

// assertModelIsUpgradable tests that the model satisfies the version upgrade interface.
//   - should be able to set and get the upgrade version
//   - setting nil version upgrade should work
//
// This has to be tested for every model version since version upgrade should be supported by all models.
func assertModelIsUpgradable(t *testing.T, api protocol_state.KVStoreAPI) {
	oldVersion := api.GetProtocolStateVersion()
	activationView := uint64(1000)
	expected := &protocol_state.ViewBasedActivator[uint64]{
		Data:           oldVersion + 1,
		ActivationView: activationView,
	}

	// check if setting version upgrade works
	api.SetVersionUpgrade(expected)
	actual := api.GetVersionUpgrade()
	assert.Equal(t, expected, actual, "version upgrade should be set")

	// check if setting nil version upgrade works
	api.SetVersionUpgrade(nil)
	assert.Nil(t, api.GetVersionUpgrade(), "version upgrade should be nil")
}
