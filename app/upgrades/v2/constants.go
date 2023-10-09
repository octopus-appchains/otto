// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package v2

const (
	// UpgradeName defines the on-chain upgrade name.
	UpgradeName = "v2"

	genesisJson = `
{
  "params": {
    "enabled": true,
    "blocks_per_distribution_transmission": "1000",
    "distribution_transmission_channel": "",
    "provider_fee_pool_addr_str": "",
    "ccv_timeout_period": "2419200s",
    "transfer_timeout_period": "3600s",
    "consumer_redistribution_fraction": "0.75",
    "historical_entries": "10000",
    "unbonding_period": "1728000s",
    "soft_opt_out_threshold": "0.05"
  },
  "provider_client_id": "",
  "provider_channel_id": "",
  "new_chain": true,
  "provider_client_state": {
    "chain_id": "provider",
    "trust_level": {
      "numerator": "1",
      "denominator": "3"
    },
    "trusting_period": "1197504s",
    "unbonding_period": "1814400s",
    "max_clock_drift": "10s",
    "frozen_height": {
      "revision_number": "0",
      "revision_height": "0"
    },
    "latest_height": {
      "revision_number": "0",
      "revision_height": "516"
    },
    "proof_specs": [
      {
        "leaf_spec": {
          "hash": "SHA256",
          "prehash_key": "NO_HASH",
          "prehash_value": "SHA256",
          "length": "VAR_PROTO",
          "prefix": "AA=="
        },
        "inner_spec": {
          "child_order": [
            0,
            1
          ],
          "child_size": 33,
          "min_prefix_length": 4,
          "max_prefix_length": 12,
          "empty_child": null,
          "hash": "SHA256"
        },
        "max_depth": 0,
        "min_depth": 0,
        "prehash_key_before_comparison": false
      },
      {
        "leaf_spec": {
          "hash": "SHA256",
          "prehash_key": "NO_HASH",
          "prehash_value": "SHA256",
          "length": "VAR_PROTO",
          "prefix": "AA=="
        },
        "inner_spec": {
          "child_order": [
            0,
            1
          ],
          "child_size": 32,
          "min_prefix_length": 1,
          "max_prefix_length": 1,
          "empty_child": null,
          "hash": "SHA256"
        },
        "max_depth": 0,
        "min_depth": 0,
        "prehash_key_before_comparison": false
      }
    ],
    "upgrade_path": [
      "upgrade",
      "upgradedIBCState"
    ],
    "allow_update_after_expiry": false,
    "allow_update_after_misbehaviour": false
  },
  "provider_consensus_state": {
    "timestamp": "2023-10-09T04:34:05.698895300Z",
    "root": {
      "hash": "qxwho4ty35HaiX3CVkQtESVoU8xWW6yR51BuCaiKuu8="
    },
    "next_validators_hash": "8B9A10D2F42D3650DB5B427460B8688B114254BA75570EF13681C45AE41806FA"
  },
  "maturing_packets": [],
  "initial_val_set": [
    {
      "pub_key": {
        "ed25519": "6TtgljW2k9j+yFfIEmAfZG+Apfq2E5lG4yxqZkXbYkM="
      },
      "power": "99"
    }
  ],
  "height_to_valset_update_id": [],
  "outstanding_downtime_slashing": [],
  "pending_consumer_packets": {
    "list": []
  },
  "last_transmission_block_height": {
    "height": "0"
  },
  "preCCV": true
}
`
)
