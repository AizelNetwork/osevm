package ibc_test

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	ibctesting "github.com/cosmos/ibc-go/v8/testing"
	evmosibc "github.com/AizelNetwork/osevm/ibc"
	precompilestestutil "github.com/AizelNetwork/osevm/precompiles/testutil"
	testconstants "github.com/AizelNetwork/osevm/testutil/constants"
	"github.com/stretchr/testify/require"
)

func init() {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("evmos", "evmospub")
}

func TestGetTransferSenderRecipient(t *testing.T) {
	testCases := []struct {
		name         string
		data         transfertypes.FungibleTokenPacketData
		expSender    string
		expRecipient string
		expError     bool
	}{
		{
			name:         "empty FungibleTokenPacketData",
			data:         transfertypes.FungibleTokenPacketData{},
			expSender:    "",
			expRecipient: "",
			expError:     true,
		},
		{
			name: "invalid sender",
			data: transfertypes.FungibleTokenPacketData{
				Sender:   "cosmos1",
				Receiver: "evmos1x2w87cvt5mqjncav4lxy8yfreynn273xn5335v",
				Amount:   "123456",
			},
			expSender:    "",
			expRecipient: "",
			expError:     true,
		},
		{
			name: "invalid recipient",
			data: transfertypes.FungibleTokenPacketData{
				Sender:   "cosmos1qql8ag4cluz6r4dz28p3w00dnc9w8ueulg2gmc",
				Receiver: "evmos1",
				Amount:   "123456",
			},
			expSender:    "",
			expRecipient: "",
			expError:     true,
		},
		{
			name: "valid - cosmos sender, evmos recipient",
			data: transfertypes.FungibleTokenPacketData{
				Sender:   "cosmos1qql8ag4cluz6r4dz28p3w00dnc9w8ueulg2gmc",
				Receiver: "evmos1x2w87cvt5mqjncav4lxy8yfreynn273xn5335v",
				Amount:   "123456",
			},
			expSender:    "evmos1qql8ag4cluz6r4dz28p3w00dnc9w8ueuafmxps",
			expRecipient: "evmos1x2w87cvt5mqjncav4lxy8yfreynn273xn5335v",
			expError:     false,
		},
		{
			name: "valid - evmos sender, cosmos recipient",
			data: transfertypes.FungibleTokenPacketData{
				Sender:   "evmos1x2w87cvt5mqjncav4lxy8yfreynn273xn5335v",
				Receiver: "cosmos1qql8ag4cluz6r4dz28p3w00dnc9w8ueulg2gmc",
				Amount:   "123456",
			},
			expSender:    "evmos1x2w87cvt5mqjncav4lxy8yfreynn273xn5335v",
			expRecipient: "evmos1qql8ag4cluz6r4dz28p3w00dnc9w8ueuafmxps",
			expError:     false,
		},
		{
			name: "valid - osmosis sender, evmos recipient",
			data: transfertypes.FungibleTokenPacketData{
				Sender:   "osmo1qql8ag4cluz6r4dz28p3w00dnc9w8ueuhnecd2",
				Receiver: "evmos1x2w87cvt5mqjncav4lxy8yfreynn273xn5335v",
				Amount:   "123456",
			},
			expSender:    "evmos1qql8ag4cluz6r4dz28p3w00dnc9w8ueuafmxps",
			expRecipient: "evmos1x2w87cvt5mqjncav4lxy8yfreynn273xn5335v",
			expError:     false,
		},
	}

	for _, tc := range testCases {
		sender, recipient, _, _, err := evmosibc.GetTransferSenderRecipient(tc.data)
		if tc.expError {
			require.Error(t, err, tc.name)
		} else {
			require.NoError(t, err, tc.name)
			require.Equal(t, tc.expSender, sender.String())
			require.Equal(t, tc.expRecipient, recipient.String())
		}
	}
}

func TestGetTransferAmount(t *testing.T) {
	testCases := []struct {
		name      string
		packet    channeltypes.Packet
		expAmount string
		expError  bool
	}{
		{
			name:      "empty packet",
			packet:    channeltypes.Packet{},
			expAmount: "",
			expError:  true,
		},
		{
			name:      "invalid packet data",
			packet:    channeltypes.Packet{Data: ibctesting.MockFailPacketData},
			expAmount: "",
			expError:  true,
		},
		{
			name: "invalid amount - empty",
			packet: channeltypes.Packet{
				Data: transfertypes.ModuleCdc.MustMarshalJSON(
					&transfertypes.FungibleTokenPacketData{
						Sender:   "cosmos1qql8ag4cluz6r4dz28p3w00dnc9w8ueulg2gmc",
						Receiver: "evmos1x2w87cvt5mqjncav4lxy8yfreynn273xn5335v",
						Amount:   "",
					},
				),
			},
			expAmount: "",
			expError:  true,
		},
		{
			name: "invalid amount - non-int",
			packet: channeltypes.Packet{
				Data: transfertypes.ModuleCdc.MustMarshalJSON(
					&transfertypes.FungibleTokenPacketData{
						Sender:   "cosmos1qql8ag4cluz6r4dz28p3w00dnc9w8ueulg2gmc",
						Receiver: "evmos1x2w87cvt5mqjncav4lxy8yfreynn273xn5335v",
						Amount:   "test",
					},
				),
			},
			expAmount: "test",
			expError:  true,
		},
		{
			name: "valid",
			packet: channeltypes.Packet{
				Data: transfertypes.ModuleCdc.MustMarshalJSON(
					&transfertypes.FungibleTokenPacketData{
						Sender:   "cosmos1qql8ag4cluz6r4dz28p3w00dnc9w8ueulg2gmc",
						Receiver: "evmos1x2w87cvt5mqjncav4lxy8yfreynn273xn5335v",
						Amount:   "10000",
					},
				),
			},
			expAmount: "10000",
			expError:  false,
		},
	}

	for _, tc := range testCases {
		amt, err := evmosibc.GetTransferAmount(tc.packet)
		if tc.expError {
			require.Error(t, err, tc.name)
		} else {
			require.NoError(t, err, tc.name)
			require.Equal(t, tc.expAmount, amt)
		}
	}
}

func TestGetReceivedCoin(t *testing.T) {
	testCases := []struct {
		name       string
		srcPort    string
		srcChannel string
		dstPort    string
		dstChannel string
		rawDenom   string
		rawAmount  string
		expCoin    sdk.Coin
	}{
		{
			"transfer unwrapped coin to destination which is not its source",
			"transfer",
			"channel-0",
			"transfer",
			"channel-0",
			"uosmo",
			"10",
			sdk.Coin{Denom: precompilestestutil.UosmoIbcdenom, Amount: math.NewInt(10)},
		},
		{
			"transfer ibc wrapped coin to destination which is its source",
			"transfer",
			"channel-0",
			"transfer",
			"channel-0",
			"transfer/channel-0/aevmos",
			"10",
			sdk.Coin{Denom: "aevmos", Amount: math.NewInt(10)},
		},
		{
			"transfer 2x ibc wrapped coin to destination which is its source",
			"transfer",
			"channel-0",
			"transfer",
			"channel-2",
			"transfer/channel-0/transfer/channel-1/uatom",
			"10",
			sdk.Coin{Denom: precompilestestutil.UatomIbcdenom, Amount: math.NewInt(10)},
		},
		{
			"transfer ibc wrapped coin to destination which is not its source",
			"transfer",
			"channel-0",
			"transfer",
			"channel-0",
			"transfer/channel-1/uatom",
			"10",
			sdk.Coin{Denom: precompilestestutil.UatomOsmoIbcdenom, Amount: math.NewInt(10)},
		},
	}

	for _, tc := range testCases {
		coin := evmosibc.GetReceivedCoin(tc.srcPort, tc.srcChannel, tc.dstPort, tc.dstChannel, tc.rawDenom, tc.rawAmount)
		require.Equal(t, tc.expCoin, coin)
	}
}

func TestGetSentCoin(t *testing.T) {
	baseDenom := testconstants.ExampleAttoDenom

	testCases := []struct {
		name      string
		rawDenom  string
		rawAmount string
		expCoin   sdk.Coin
	}{
		{
			"get unwrapped aevmos coin",
			baseDenom,
			"10",
			sdk.Coin{Denom: baseDenom, Amount: math.NewInt(10)},
		},
		{
			"get ibc wrapped aevmos coin",
			"transfer/channel-0/aevmos",
			"10",
			sdk.Coin{Denom: precompilestestutil.AevmosIbcdenom, Amount: math.NewInt(10)},
		},
		{
			"get ibc wrapped uosmo coin",
			"transfer/channel-0/uosmo",
			"10",
			sdk.Coin{Denom: precompilestestutil.UosmoIbcdenom, Amount: math.NewInt(10)},
		},
		{
			"get ibc wrapped uatom coin",
			"transfer/channel-1/uatom",
			"10",
			sdk.Coin{Denom: precompilestestutil.UatomIbcdenom, Amount: math.NewInt(10)},
		},
		{
			"get 2x ibc wrapped uatom coin",
			"transfer/channel-0/transfer/channel-1/uatom",
			"10",
			sdk.Coin{Denom: precompilestestutil.UatomOsmoIbcdenom, Amount: math.NewInt(10)},
		},
	}

	for _, tc := range testCases {
		coin := evmosibc.GetSentCoin(tc.rawDenom, tc.rawAmount)
		require.Equal(t, tc.expCoin, coin)
	}
}

func TestDeriveDecimalsFromDenom(t *testing.T) {
	testCases := []struct {
		name      string
		baseDenom string
		expDec    uint8
		expFail   bool
		expErrMsg string
	}{
		{
			name:      "fail: empty string",
			baseDenom: "",
			expDec:    0,
			expFail:   true,
			expErrMsg: "Base denom cannot be an empty string",
		},
		{
			name:      "fail: invalid prefix",
			baseDenom: "nevmos",
			expDec:    0,
			expFail:   true,
			expErrMsg: "Should be either micro ('u[...]') or atto ('a[...]'); got: \"nevmos\"",
		},
		{
			name:      "success: micro 'u' prefix",
			baseDenom: "uevmos",
			expDec:    6,
			expFail:   false,
			expErrMsg: "",
		},
		{
			name:      "success: atto 'a' prefix",
			baseDenom: "aevmos",
			expDec:    18,
			expFail:   false,
			expErrMsg: "",
		},
	}

	for _, tc := range testCases {
		dec, err := evmosibc.DeriveDecimalsFromDenom(tc.baseDenom)
		if tc.expFail {
			require.Error(t, err, tc.expErrMsg)
			require.Contains(t, err.Error(), tc.expErrMsg)
		} else {
			require.NoError(t, err)
		}
		require.Equal(t, tc.expDec, dec)
	}
}

func TestIsBaseDenomFromSourceChain(t *testing.T) {
	tests := []struct {
		name     string
		denom    string
		expected bool
	}{
		{
			name:     "one hop",
			denom:    "transfer/channel-0/uatom",
			expected: false,
		},
		{
			name:     "no hop with factory prefix",
			denom:    "factory/owner/uatom",
			expected: false,
		},
		{
			name:     "multi hop",
			denom:    "transfer/channel-0/transfer/channel-1/uatom",
			expected: false,
		},
		{
			name:     "no hop",
			denom:    "uatom",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evmosibc.IsBaseDenomFromSourceChain(tt.denom)
			require.Equal(t, tt.expected, result)
		})
	}
}
