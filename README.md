## Decision Log and Status Server

This sample server implements the APIs for OPA's [Decision Log](https://www.openpolicyagent.org/docs/edge/management/#decision-logs)
and [Status](https://www.openpolicyagent.org/docs/edge/management/#status) management features.

The steps below show an example of monitoring the number of decision logs dropped when the rate limit is exceeded
via a status update.

### Steps

#### 1. Start Bundle server

Start a server that serves the bundle resource `bundle.tar.gz`. The gzipped tarball contains the policy defined in
`allow.rego`.

```bash
python3 -m http.server
```

#### 2. Start Decision Log / Status server

In another terminal start the Decision Log / Status server. It logs the decision log and status events to console.

```bash
go run server.go
```

#### 3. Run OPA

In another terminal run OPA with the config defined in `config.yaml`. The decision log config sets the
rate limit to `1 decision per second`.

```bash
opa run -s -c config.yaml
```

#### 4. Run a load test

The below test will attack OPA at a rate of `10 requests per sec`.

```bash
echo 'POST http://localhost:8181/v1/data/authz/allow' | vegeta attack --duration=1s -rate=10 | tee results.bin | vegeta report
```

#### 5. Examine status update

Since the rate limit is `1 decision per second`, **9** decision log events should be dropped. The value of the
`counter_decision_logs_dropped` metric in the status update that's printed to console should be equal to `9`. Below is
an example of the status update.

```json
{
 "bundles": {
  "authz": {
   "last_request": "2021-04-03T00:54:27.419103Z",
   "last_successful_activation": "2021-04-03T00:54:27.419678Z",
   "last_successful_download": "2021-04-03T00:54:27.419103Z",
   "last_successful_request": "2021-04-03T00:54:27.419103Z",
   "metrics": {
    "timer_rego_data_parse_ns": 6846,
    "timer_rego_load_bundles_ns": 6464822,
    "timer_rego_module_compile_ns": 396674,
    "timer_rego_module_parse_ns": 37185
   },
   "name": "authz"
  }
 },
 "labels": {
  "id": "adbc9ab8-73b8-43cf-963a-3d7aa652da26",
  "version": "0.28.0-dev"
 },
 "metrics": {
  "prometheus": {
   "counter_decision_logs_dropped": 9,
   "go_gc_duration_seconds": {
    "help": "A summary of the pause duration of garbage collection cycles.",
    "metric": [
     {
      "summary": {
       "quantile": [
        {
         "quantile": 0,
         "value": 0.000065092
        },
        {
         "quantile": 0.25,
         "value": 0.000065092
        },
        {
         "quantile": 0.5,
         "value": 0.000161704
        },
        {
         "quantile": 0.75,
         "value": 0.000161704
        },
        {
         "quantile": 1,
         "value": 0.000161704
        }
       ],
       "sample_count": 2,
       "sample_sum": 0.000226796
      }
     }
    ],
    "name": "go_gc_duration_seconds",
    "type": 2
   },
   "go_goroutines": {
    "help": "Number of goroutines that currently exist.",
    "metric": [
     {
      "gauge": {
       "value": 21
      }
     }
    ],
    "name": "go_goroutines",
    "type": 1
   },
   "go_info": {
    "help": "Information about the Go environment.",
    "metric": [
     {
      "gauge": {
       "value": 1
      },
      "label": [
       {
        "name": "version",
        "value": "go1.15.2"
       }
      ]
     }
    ],
    "name": "go_info",
    "type": 1
   },
   "go_memstats_alloc_bytes": {
    "help": "Number of bytes allocated and still in use.",
    "metric": [
     {
      "gauge": {
       "value": 3747720
      }
     }
    ],
    "name": "go_memstats_alloc_bytes",
    "type": 1
   },
   "go_memstats_alloc_bytes_total": {
    "help": "Total number of bytes allocated, even if freed.",
    "metric": [
     {
      "counter": {
       "value": 7207112
      }
     }
    ],
    "name": "go_memstats_alloc_bytes_total",
    "type": 0
   },
   "go_memstats_buck_hash_sys_bytes": {
    "help": "Number of bytes used by the profiling bucket hash table.",
    "metric": [
     {
      "gauge": {
       "value": 1446757
      }
     }
    ],
    "name": "go_memstats_buck_hash_sys_bytes",
    "type": 1
   },
   "go_memstats_frees_total": {
    "help": "Total number of frees.",
    "metric": [
     {
      "counter": {
       "value": 78526
      }
     }
    ],
    "name": "go_memstats_frees_total",
    "type": 0
   },
   "go_memstats_gc_cpu_fraction": {
    "help": "The fraction of this program's available CPU time used by the GC since the program started.",
    "metric": [
     {
      "gauge": {
       "value": 0.000018730051565049067
      }
     }
    ],
    "name": "go_memstats_gc_cpu_fraction",
    "type": 1
   },
   "go_memstats_gc_sys_bytes": {
    "help": "Number of bytes used for garbage collection system metadata.",
    "metric": [
     {
      "gauge": {
       "value": 5299992
      }
     }
    ],
    "name": "go_memstats_gc_sys_bytes",
    "type": 1
   },
   "go_memstats_heap_alloc_bytes": {
    "help": "Number of heap bytes allocated and still in use.",
    "metric": [
     {
      "gauge": {
       "value": 3747720
      }
     }
    ],
    "name": "go_memstats_heap_alloc_bytes",
    "type": 1
   },
   "go_memstats_heap_idle_bytes": {
    "help": "Number of heap bytes waiting to be used.",
    "metric": [
     {
      "gauge": {
       "value": 60252160
      }
     }
    ],
    "name": "go_memstats_heap_idle_bytes",
    "type": 1
   },
   "go_memstats_heap_inuse_bytes": {
    "help": "Number of heap bytes that are in use.",
    "metric": [
     {
      "gauge": {
       "value": 6037504
      }
     }
    ],
    "name": "go_memstats_heap_inuse_bytes",
    "type": 1
   },
   "go_memstats_heap_objects": {
    "help": "Number of allocated objects.",
    "metric": [
     {
      "gauge": {
       "value": 29121
      }
     }
    ],
    "name": "go_memstats_heap_objects",
    "type": 1
   },
   "go_memstats_heap_released_bytes": {
    "help": "Number of heap bytes released to OS.",
    "metric": [
     {
      "gauge": {
       "value": 59867136
      }
     }
    ],
    "name": "go_memstats_heap_released_bytes",
    "type": 1
   },
   "go_memstats_heap_sys_bytes": {
    "help": "Number of heap bytes obtained from system.",
    "metric": [
     {
      "gauge": {
       "value": 66289664
      }
     }
    ],
    "name": "go_memstats_heap_sys_bytes",
    "type": 1
   },
   "go_memstats_last_gc_time_seconds": {
    "help": "Number of seconds since 1970 of last garbage collection.",
    "metric": [
     {
      "gauge": {
       "value": 1617411232.391993
      }
     }
    ],
    "name": "go_memstats_last_gc_time_seconds",
    "type": 1
   },
   "go_memstats_lookups_total": {
    "help": "Total number of pointer lookups.",
    "metric": [
     {
      "counter": {
       "value": 0
      }
     }
    ],
    "name": "go_memstats_lookups_total",
    "type": 0
   },
   "go_memstats_mallocs_total": {
    "help": "Total number of mallocs.",
    "metric": [
     {
      "counter": {
       "value": 107647
      }
     }
    ],
    "name": "go_memstats_mallocs_total",
    "type": 0
   },
   "go_memstats_mcache_inuse_bytes": {
    "help": "Number of bytes in use by mcache structures.",
    "metric": [
     {
      "gauge": {
       "value": 13888
      }
     }
    ],
    "name": "go_memstats_mcache_inuse_bytes",
    "type": 1
   },
   "go_memstats_mcache_sys_bytes": {
    "help": "Number of bytes used for mcache structures obtained from system.",
    "metric": [
     {
      "gauge": {
       "value": 16384
      }
     }
    ],
    "name": "go_memstats_mcache_sys_bytes",
    "type": 1
   },
   "go_memstats_mspan_inuse_bytes": {
    "help": "Number of bytes in use by mspan structures.",
    "metric": [
     {
      "gauge": {
       "value": 124712
      }
     }
    ],
    "name": "go_memstats_mspan_inuse_bytes",
    "type": 1
   },
   "go_memstats_mspan_sys_bytes": {
    "help": "Number of bytes used for mspan structures obtained from system.",
    "metric": [
     {
      "gauge": {
       "value": 131072
      }
     }
    ],
    "name": "go_memstats_mspan_sys_bytes",
    "type": 1
   },
   "go_memstats_next_gc_bytes": {
    "help": "Number of heap bytes when next garbage collection will take place.",
    "metric": [
     {
      "gauge": {
       "value": 6027440
      }
     }
    ],
    "name": "go_memstats_next_gc_bytes",
    "type": 1
   },
   "go_memstats_other_sys_bytes": {
    "help": "Number of bytes used for other system allocations.",
    "metric": [
     {
      "gauge": {
       "value": 1315203
      }
     }
    ],
    "name": "go_memstats_other_sys_bytes",
    "type": 1
   },
   "go_memstats_stack_inuse_bytes": {
    "help": "Number of bytes in use by the stack allocator.",
    "metric": [
     {
      "gauge": {
       "value": 819200
      }
     }
    ],
    "name": "go_memstats_stack_inuse_bytes",
    "type": 1
   },
   "go_memstats_stack_sys_bytes": {
    "help": "Number of bytes obtained from system for stack allocator.",
    "metric": [
     {
      "gauge": {
       "value": 819200
      }
     }
    ],
    "name": "go_memstats_stack_sys_bytes",
    "type": 1
   },
   "go_memstats_sys_bytes": {
    "help": "Number of bytes obtained from system.",
    "metric": [
     {
      "gauge": {
       "value": 75318272
      }
     }
    ],
    "name": "go_memstats_sys_bytes",
    "type": 1
   },
   "go_threads": {
    "help": "Number of OS threads created.",
    "metric": [
     {
      "gauge": {
       "value": 14
      }
     }
    ],
    "name": "go_threads",
    "type": 1
   },
   "http_request_duration_seconds": {
    "help": "A histogram of duration for requests.",
    "metric": [
     {
      "histogram": {
       "bucket": [
        {
         "cumulative_count": 0,
         "upper_bound": 0.000001
        },
        {
         "cumulative_count": 0,
         "upper_bound": 0.000005
        },
        {
         "cumulative_count": 0,
         "upper_bound": 0.00001
        },
        {
         "cumulative_count": 0,
         "upper_bound": 0.00005
        },
        {
         "cumulative_count": 0,
         "upper_bound": 0.0001
        },
        {
         "cumulative_count": 9,
         "upper_bound": 0.0005
        },
        {
         "cumulative_count": 9,
         "upper_bound": 0.001
        },
        {
         "cumulative_count": 10,
         "upper_bound": 0.01
        },
        {
         "cumulative_count": 10,
         "upper_bound": 0.1
        },
        {
         "cumulative_count": 10,
         "upper_bound": 1
        }
       ],
       "sample_count": 10,
       "sample_sum": 0.004671040999999999
      },
      "label": [
       {
        "name": "code",
        "value": "200"
       },
       {
        "name": "handler",
        "value": "v1/data"
       },
       {
        "name": "method",
        "value": "post"
       }
      ]
     }
    ],
    "name": "http_request_duration_seconds",
    "type": 4
   }
  }
 },
 "plugins": {
  "bundle": {
   "state": "OK"
  },
  "decision_logs": {
   "state": "OK"
  },
  "discovery": {
   "state": "OK"
  },
  "status": {
   "state": "OK"
  }
 }
}
```
