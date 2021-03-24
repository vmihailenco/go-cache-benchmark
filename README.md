<p align="center">
  <a href="https://uptrace.dev/?utm_source=gh-redis&utm_campaign=gh-redis-banner1">
    <img src="https://raw.githubusercontent.com/uptrace/roadmap/master/banner1.png" alt="All-in-one tool to optimize performance and monitor errors & logs">
  </a>
</p>

# Cache comparison benchmark for Go

This benchmark compares cache algorithms using scrambled zipfian distribution (a few occur very
often while many others occur rarely). Other distributions are supported too, but they produce
similar results. You may also want to check
[cachetest](https://github.com/dgryski/trifles/tree/master/cachetest).

The following libraries are supported:

- https://pkg.go.dev/github.com/dgryski/go-tinylfu
- https://pkg.go.dev/github.com/dgryski/go-clockpro
- https://pkg.go.dev/github.com/dgraph-io/ristretto
- https://pkg.go.dev/github.com/hashicorp/golang-lru (LRU, ARC, TwoQueue)
- https://pkg.go.dev/github.com/golang/groupcache/lru
- https://pkg.go.dev/github.com/dgryski/go-s4lru

And the results are:

<details>
  <summary>TLDR</summary>

- TinyLFU works best for small number of keys (~ 100k). TinyLFU memory overhead can be tweaked with
  the 2nd argument.
- Clock-pro has significantly smaller memory usage with large number of keys (~ 1m).
- Segmented LRU has even smaller memory usage, but has inconsistent hit rate.
- Ristretto can still be a good choice if you need additional features it provides.
</details>

```
zipfian cache=1000 keys=2000

      CACHE      | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
-----------------+----------+--------+---------+--------+---------
  tinylfu        |     0.62 |  530ms | 0.19MiB | 620898 | 379102
  clockpro       |     0.61 |  473ms | 0.23MiB | 606905 | 393095
  arc            |     0.60 | 1014ms | 0.35MiB | 604548 | 395452
  ristretto      |     0.62 | 1005ms | 0.79MiB | 623662 | 376338
  two-queue      |     0.60 |  867ms | 0.28MiB | 602307 | 397693
  lru-groupcache |     0.60 |  577ms | 0.19MiB | 601762 | 398238
  lru-hashicorp  |     0.60 |  647ms | 0.19MiB | 601209 | 398791
  s4lru          |     0.62 |  401ms | 0.19MiB | 616082 | 383918
  slru           |     0.61 |  704ms | 0.20MiB | 614580 | 385420


zipfian cache=1000 keys=10000

      CACHE      | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
-----------------+----------+--------+---------+--------+---------
  tinylfu        |     0.34 |  595ms | 0.19MiB | 338356 | 661644
  clockpro       |     0.33 |  655ms | 0.40MiB | 328305 | 671695
  arc            |     0.32 | 1366ms | 0.49MiB | 320453 | 679547
  ristretto      |     0.32 | 1197ms | 4.15MiB | 317582 | 682418
  two-queue      |     0.32 | 1246ms | 0.33MiB | 318107 | 681893
  lru-groupcache |     0.28 |  728ms | 0.22MiB | 280302 | 719698
  lru-hashicorp  |     0.28 |  810ms | 0.22MiB | 281400 | 718600
  s4lru          |     0.33 |  419ms | 0.22MiB | 332631 | 667369
  slru           |     0.33 |  802ms | 0.20MiB | 329516 | 670484


zipfian cache=1000 keys=100000

      CACHE      | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
-----------------+----------+--------+---------+--------+---------
  tinylfu        |     0.29 |  577ms | 0.19MiB | 289710 | 710290
  clockpro       |     0.29 |  691ms | 0.41MiB | 286546 | 713454
  arc            |     0.28 | 1328ms | 0.47MiB | 283291 | 716709
  ristretto      |     0.23 | 1228ms | 4.10MiB | 234963 | 765037
  two-queue      |     0.28 | 1360ms | 0.33MiB | 276287 | 723713
  lru-groupcache |     0.21 |  739ms | 0.20MiB | 207823 | 792177
  lru-hashicorp  |     0.21 |  871ms | 0.22MiB | 206861 | 793139
  s4lru          |     0.28 |  436ms | 0.22MiB | 276586 | 723414
  slru           |     0.28 |  805ms | 0.18MiB | 275731 | 724269


zipfian cache=10000 keys=20000

      CACHE      | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
-----------------+----------+--------+---------+--------+---------
  tinylfu        |     0.66 |  603ms | 1.91MiB | 658562 | 341438
  clockpro       |     0.65 |  523ms | 2.50MiB | 646865 | 353135
  arc            |     0.65 | 1128ms | 3.46MiB | 645484 | 354516
  ristretto      |     0.66 |  914ms | 4.43MiB | 657079 | 342921
  two-queue      |     0.65 | 1005ms | 2.92MiB | 645196 | 354804
  lru-groupcache |     0.64 |  665ms | 1.83MiB | 642420 | 357580
  lru-hashicorp  |     0.64 |  843ms | 1.83MiB | 643063 | 356937
  s4lru          |     0.65 |  458ms | 1.83MiB | 646058 | 353942
  slru           |     0.64 |  818ms | 2.02MiB | 640791 | 359209


zipfian cache=10000 keys=100000

      CACHE      | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
-----------------+----------+--------+---------+--------+---------
  tinylfu        |     0.41 |  709ms | 1.99MiB | 411491 | 588509
  clockpro       |     0.41 |  704ms | 3.53MiB | 406171 | 593829
  arc            |     0.40 | 1472ms | 4.19MiB | 396756 | 603244
  ristretto      |     0.40 | 1085ms | 7.11MiB | 398042 | 601958
  two-queue      |     0.40 | 1374ms | 3.06MiB | 398142 | 601858
  lru-groupcache |     0.37 |  847ms | 1.98MiB | 365646 | 634354
  lru-hashicorp  |     0.36 |  995ms | 1.98MiB | 364871 | 635129
  s4lru          |     0.39 |  487ms | 1.94MiB | 392482 | 607518
  slru           |     0.39 |  917ms | 1.97MiB | 385109 | 614891


zipfian cache=10000 keys=1000000

      CACHE      | HIT RATE |  TIME  | MEMORY  |  HITS  | MISSES
-----------------+----------+--------+---------+--------+---------
  tinylfu        |     0.37 |  703ms | 1.97MiB | 366269 | 633731
  clockpro       |     0.36 |  719ms | 3.59MiB | 364253 | 635747
  arc            |     0.36 | 1430ms | 4.19MiB | 361819 | 638181
  ristretto      |     0.34 | 1129ms | 7.02MiB | 338664 | 661336
  two-queue      |     0.36 | 1456ms | 3.22MiB | 357508 | 642492
  lru-groupcache |     0.30 |  898ms | 1.96MiB | 303337 | 696663
  lru-hashicorp  |     0.30 | 1052ms | 1.97MiB | 303367 | 696633
  s4lru          |     0.33 |  476ms | 1.55MiB | 331624 | 668376
  slru           |     0.33 |  870ms | 1.34MiB | 326673 | 673327


zipfian cache=100000 keys=200000

      CACHE      | HIT RATE |  TIME  |  MEMORY  |  HITS  | MISSES
-----------------+----------+--------+----------+--------+---------
  tinylfu        |     0.67 |  720ms | 18.82MiB | 668610 | 331390
  clockpro       |     0.67 |  599ms | 21.74MiB | 667152 | 332848
  arc            |     0.66 | 1283ms | 30.13MiB | 664304 | 335696
  ristretto      |     0.67 | 1071ms | 26.58MiB | 667405 | 332595
  two-queue      |     0.66 | 1118ms | 25.94MiB | 663983 | 336017
  lru-groupcache |     0.66 |  707ms | 17.02MiB | 663243 | 336757
  lru-hashicorp  |     0.66 |  751ms | 17.05MiB | 662651 | 337349
  s4lru          |     0.59 |  560ms | 16.52MiB | 591193 | 408807
  slru           |     0.55 |  833ms | 16.65MiB | 551634 | 448366


zipfian cache=100000 keys=1000000

      CACHE      | HIT RATE |  TIME  |  MEMORY  |  HITS  | MISSES
-----------------+----------+--------+----------+--------+---------
  tinylfu        |     0.46 |  735ms | 19.17MiB | 458678 | 541322
  clockpro       |     0.46 |  675ms | 29.61MiB | 458839 | 541161
  arc            |     0.46 | 1445ms | 35.96MiB | 460111 | 539889
  ristretto      |     0.45 | 1164ms | 29.51MiB | 454872 | 545128
  two-queue      |     0.46 | 1308ms | 28.79MiB | 459213 | 540787
  lru-groupcache |     0.45 |  744ms | 17.75MiB | 446016 | 553984
  lru-hashicorp  |     0.45 |  813ms | 17.76MiB | 445783 | 554217
  s4lru          |     0.40 |  528ms |  9.19MiB | 403832 | 596168
  slru           |     0.39 |  802ms |  8.00MiB | 393951 | 606049


zipfian cache=100000 keys=10000000

      CACHE      | HIT RATE |  TIME  |  MEMORY  |  HITS  | MISSES
-----------------+----------+--------+----------+--------+---------
  tinylfu        |     0.41 |  723ms | 18.76MiB | 412350 | 587650
  clockpro       |     0.41 |  641ms | 29.72MiB | 412315 | 587685
  arc            |     0.41 | 1500ms | 35.10MiB | 412989 | 587011
  ristretto      |     0.41 | 1249ms | 30.08MiB | 406083 | 593917
  two-queue      |     0.41 | 1357ms | 27.72MiB | 411963 | 588037
  lru-groupcache |     0.40 |  742ms | 17.89MiB | 396386 | 603614
  lru-hashicorp  |     0.40 |  810ms | 17.88MiB | 395939 | 604061
  s4lru          |     0.38 |  498ms |  7.47MiB | 380136 | 619864
  slru           |     0.38 |  819ms |  7.15MiB | 376288 | 623712


zipfian cache=1000000 keys=2000000

      CACHE      | HIT RATE |  TIME  |  MEMORY   |  HITS  | MISSES
-----------------+----------+--------+-----------+--------+---------
  tinylfu        |     0.51 |  791ms | 131.57MiB | 506342 | 493658
  clockpro       |     0.51 |  661ms |  78.41MiB | 506761 | 493239
  arc            |     0.51 | 1013ms |  78.20MiB | 506726 | 493274
  ristretto      |     0.51 | 1236ms | 193.88MiB | 505084 | 494916
  two-queue      |     0.51 |  981ms |  78.14MiB | 507041 | 492959
  lru-groupcache |     0.51 |  813ms |  85.83MiB | 507737 | 492263
  lru-hashicorp  |     0.51 |  886ms |  85.89MiB | 507129 | 492871
  s4lru          |     0.49 |  607ms |  55.10MiB | 493332 | 506668
  slru           |     0.48 |  867ms |  47.28MiB | 482858 | 517142


zipfian cache=1000000 keys=10000000

      CACHE      | HIT RATE |  TIME  |  MEMORY   |  HITS  | MISSES
-----------------+----------+--------+-----------+--------+---------
  tinylfu        |     0.45 |  778ms | 138.20MiB | 448427 | 551573
  clockpro       |     0.45 |  625ms |  84.20MiB | 448418 | 551582
  arc            |     0.45 | 1100ms |  98.42MiB | 447664 | 552336
  ristretto      |     0.45 | 1301ms | 153.94MiB | 446974 | 553026
  two-queue      |     0.45 | 1103ms |  98.41MiB | 447758 | 552242
  lru-groupcache |     0.45 |  817ms |  92.67MiB | 447916 | 552084
  lru-hashicorp  |     0.45 |  882ms |  92.69MiB | 447766 | 552234
  s4lru          |     0.44 |  645ms |  49.73MiB | 437770 | 562230
  slru           |     0.43 |  844ms |  41.19MiB | 431229 | 568771
```

To run this benchmark:

```shell
go run *.go
```
