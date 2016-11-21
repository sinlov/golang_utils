# golang_utils

this is self golang utils

# base use

- add lib of lib

```sh
go get -u -v github.com/bmizerany/assert
```

- clone lib of utils

```sh
go get -u -v github.com/sinlov/golang_utils
```

# Use

## cfg

```golang
import (
    "github.com/sinlov/golang_utils/cfg"
)
	cfg := new(cfg.Cfg)
	cfg.InitCfg("config.conf")
	daemon := cfg.Read("ServerSet", "daemon")
	port := cfg.Read("ServerSet", "port")
	fmt.Printf("daemon: %v, port: %v \n", daemon, port)
```

## randomplus

- TimeSeed

random number by time

```golang
import (
    "github.com/sinlov/golang_utils/randomplus"
)
    random, sed := randomplus.TimeSeed(999)
    fmt.Printf("random %v, time sed %v", random, sed.Unix())
```

- Positive will new random number by size

Like size 8 is [10000000, 99999999]

- PositiveNegative new random by size

Like size 4 is [-9999, -1000] to [1000, 9999]

```golang
    got, err := randomplus.PositiveNegative(8)
		if err != nil {
			fmt.Errorf("PositiveNegative() error = %v", err)

		}
	got, err := randomplus.Positive(5)
		if err != nil {
			fmt.Errorf("Postive() error = %v", err)
}
```

#License

---

Copyright 2016 sinlovgm@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.