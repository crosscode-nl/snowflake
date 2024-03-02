# snowflake

Go implementation of Twitter Snowflake

## Comparison

| Package                           | Performance | Drift | Blocking       | Custom Epoch      | Change bit size   | Synchronization | Timestamp    | ID     | Coverage |
|-----------------------------------|-------------|-------|----------------|-------------------|-------------------|-----------------|--------------|--------|----------|
| github.com/crosscode-nl/snowflake | 42.42 ns/op | yes   | yes, sleep     | yes, per instance | yes, per instance | CAS, unlimited  | 42 bits      | uint64 | 100%     |
| github.com/influxdata/snowflake   | 42.70 ns/op | yes   | no             | no                | no                | CAS, 100x, bug  | 42 bits      | uint64 | 88.9%    |
| github.com/bwmarrin/snowflake     | 244.1 ns/op | no    | yes, busy loop | yes, global       | yes, global       | Mutex           | 42 bits, bug | int64  | 92.9%    |
| github.com/godruoyi/go-snowflake  | 244.2 ns/op | no    | yes, busy loop | yes, global       | yes, global       | CAS, unlimited  | 41 bits      | uint64 | 91.4%    |

