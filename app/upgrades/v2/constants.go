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
    "timestamp": "2023-10-13T01:55:04.299344144Z",
    "root": {
      "hash": "AmjSjjCwwycdGU7J+6IZKcmSdmyz1m730jez6+xV/P8="
    },
    "next_validators_hash": "E67A5BABE54E98BF6CCF9E992188049B4F3968F0DEA33058CE69B69349F22A8E"
  },
  "maturing_packets": [],
  "initial_val_set": [
    {
      "pub_key": {
        "ed25519": "2CKdA3Sbl1hh6+Exdqy7LfspfGcgUtNhV1VwUAZcy7c="
      },
      "power": "10000"
    },
    {
      "pub_key": {
        "ed25519": "s6mgS6YztAJ0V0GyV0Rc6g4E3Y/QLFI4wZ3c5LEDaks="
      },
      "power": "10000"
    },
    {
      "pub_key": {
        "ed25519": "xaTzPWfA7PS06PDJw/p6p/QWSBqYH71PeeUqlPSv0QM="
      },
      "power": "10000"
    },
    {
      "pub_key": {
        "ed25519": "2JOX9X/hXjMVN1KfOPxZzRruKKb45kZW1nICBXHaXBI="
      },
      "power": "10000"
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
