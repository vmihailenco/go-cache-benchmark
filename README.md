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
zipfian cache=1000 keys=10000

      CACHE      | HIT RATE | TIME  | MEMORY  |  HITS  | MISSES  
-----------------+----------+-------+---------+--------+---------
  tinylfu        |     0.34 | 254ms | 0.20MiB | 339363 | 660637  
  clockpro       |     0.33 | 262ms | 0.40MiB | 328143 | 671857  
  arc            |     0.32 | 455ms | 0.49MiB | 319438 | 680562  
  ristretto      |     0.32 | 436ms | 3.52MiB | 320537 | 679463  
  directcache    |     0.48 | 218ms | 0.21MiB | 481128 | 518872  
  two-queue      |     0.32 | 456ms | 0.33MiB | 316910 | 683090  
  lru-groupcache |     0.28 | 266ms | 0.22MiB | 280730 | 719270  
  lru-hashicorp  |     0.28 | 288ms | 0.22MiB | 280377 | 719623  
  s4lru          |     0.33 | 198ms | 0.22MiB | 332599 | 667401  
  slru           |     0.33 | 300ms | 0.20MiB | 329509 | 670491  
  wtfcache       |     0.28 | 229ms | 0.19MiB | 279742 | 720258  


zipfian cache=1000 keys=100000

      CACHE      | HIT RATE | TIME  | MEMORY  |  HITS  | MISSES  
-----------------+----------+-------+---------+--------+---------
  tinylfu        |     0.29 | 242ms | 0.18MiB | 289752 | 710248  
  clockpro       |     0.29 | 259ms | 0.41MiB | 285591 | 714409  
  arc            |     0.28 | 446ms | 0.43MiB | 283353 | 716647  
  ristretto      |     0.24 | 468ms | 3.93MiB | 236928 | 763072  
  directcache    |     0.27 | 224ms | 0.21MiB | 274123 | 725877  
  two-queue      |     0.28 | 469ms | 0.32MiB | 276276 | 723724  
  lru-groupcache |     0.21 | 279ms | 0.21MiB | 208824 | 791176  
  lru-hashicorp  |     0.21 | 301ms | 0.22MiB | 207826 | 792174  
  s4lru          |     0.28 | 196ms | 0.22MiB | 277247 | 722753  
  slru           |     0.28 | 302ms | 0.19MiB | 276295 | 723705  
  wtfcache       |     0.21 | 244ms | 0.19MiB | 207185 | 792815  


zipfian cache=1000 keys=1000000

      CACHE      | HIT RATE | TIME  | MEMORY  |  HITS  | MISSES  
-----------------+----------+-------+---------+--------+---------
  tinylfu        |     0.29 | 248ms | 0.18MiB | 288601 | 711399  
  clockpro       |     0.29 | 260ms | 0.41MiB | 288899 | 711101  
  arc            |     0.29 | 456ms | 0.45MiB | 287758 | 712242  
  ristretto      |     0.22 | 495ms | 0.76MiB | 218458 | 781542  
  directcache    |     0.25 | 233ms | 0.21MiB | 254150 | 745850  
  two-queue      |     0.28 | 484ms | 0.34MiB | 278537 | 721463  
  lru-groupcache |     0.20 | 278ms | 0.22MiB | 200153 | 799847  
  lru-hashicorp  |     0.20 | 318ms | 0.21MiB | 201648 | 798352  
  s4lru          |     0.27 | 213ms | 0.21MiB | 272905 | 727095  
  slru           |     0.27 | 314ms | 0.18MiB | 270394 | 729606  
  wtfcache       |     0.20 | 245ms | 0.19MiB | 200961 | 799039  


zipfian cache=10000 keys=100000

      CACHE      | HIT RATE | TIME  | MEMORY  |  HITS  | MISSES  
-----------------+----------+-------+---------+--------+---------
  tinylfu        |     0.41 | 281ms | 1.99MiB | 412001 | 587999  
  clockpro       |     0.41 | 432ms | 3.54MiB | 406993 | 593007  
  arc            |     0.40 | 919ms | 4.16MiB | 395673 | 604327  
  ristretto      |     0.40 | 366ms | 7.02MiB | 396708 | 603292  
  directcache    |     0.37 | 249ms | 0.77MiB | 370036 | 629964  
  two-queue      |     0.40 | 730ms | 3.13MiB | 398233 | 601767  
  lru-groupcache |     0.37 | 364ms | 1.98MiB | 365455 | 634545  
  lru-hashicorp  |     0.36 | 425ms | 1.98MiB | 364845 | 635155  
  s4lru          |     0.39 | 313ms | 1.94MiB | 392852 | 607148  
  slru           |     0.39 | 410ms | 1.98MiB | 385549 | 614451  
  wtfcache       |     0.37 | 301ms | 1.68MiB | 365759 | 634241  


zipfian cache=10000 keys=1000000

      CACHE      | HIT RATE | TIME  | MEMORY  |  HITS  | MISSES  
-----------------+----------+-------+---------+--------+---------
  tinylfu        |     0.37 | 314ms | 1.97MiB | 367741 | 632259  
  clockpro       |     0.36 | 338ms | 3.59MiB | 363505 | 636495  
  arc            |     0.36 | 584ms | 4.19MiB | 362503 | 637497  
  ristretto      |     0.34 | 382ms | 7.10MiB | 338357 | 661643  
  directcache    |     0.31 | 230ms | 0.76MiB | 309275 | 690725  
  two-queue      |     0.36 | 567ms | 3.21MiB | 359047 | 640953  
  lru-groupcache |     0.30 | 329ms | 1.97MiB | 304014 | 695986  
  lru-hashicorp  |     0.30 | 354ms | 1.97MiB | 303598 | 696402  
  s4lru          |     0.33 | 223ms | 1.60MiB | 331456 | 668544  
  slru           |     0.33 | 333ms | 1.34MiB | 326529 | 673471  
  wtfcache       |     0.30 | 271ms | 1.67MiB | 303214 | 696786  


zipfian cache=10000 keys=10000000

      CACHE      | HIT RATE | TIME  | MEMORY  |  HITS  | MISSES  
-----------------+----------+-------+---------+--------+---------
  tinylfu        |     0.36 | 277ms | 1.95MiB | 364256 | 635744  
  clockpro       |     0.36 | 302ms | 3.60MiB | 360103 | 639897  
  arc            |     0.36 | 598ms | 4.16MiB | 360572 | 639428  
  ristretto      |     0.33 | 345ms | 7.14MiB | 331305 | 668695  
  directcache    |     0.30 | 234ms | 0.75MiB | 303046 | 696954  
  two-queue      |     0.35 | 575ms | 3.24MiB | 354729 | 645271  
  lru-groupcache |     0.30 | 321ms | 1.98MiB | 297185 | 702815  
  lru-hashicorp  |     0.30 | 352ms | 1.98MiB | 296834 | 703166  
  s4lru          |     0.33 | 219ms | 1.33MiB | 327179 | 672821  
  slru           |     0.32 | 328ms | 1.19MiB | 321550 | 678450  
  wtfcache       |     0.30 | 271ms | 1.67MiB | 298173 | 701827  


zipfian cache=100000 keys=1000000

      CACHE      | HIT RATE | TIME  |  MEMORY  |  HITS  | MISSES  
-----------------+----------+-------+----------+--------+---------
  tinylfu        |     0.46 | 465ms | 19.18MiB | 458753 | 541247  
  clockpro       |     0.46 | 480ms | 29.59MiB | 458491 | 541509  
  arc            |     0.46 | 980ms | 35.91MiB | 459607 | 540393  
  ristretto      |     0.45 | 432ms | 29.09MiB | 451846 | 548154  
  directcache    |     0.45 | 266ms |  6.24MiB | 449331 | 550669  
  two-queue      |     0.46 | 886ms | 28.82MiB | 459891 | 540109  
  lru-groupcache |     0.45 | 450ms | 17.75MiB | 445664 | 554336  
  lru-hashicorp  |     0.45 | 496ms | 17.76MiB | 445351 | 554649  
  s4lru          |     0.40 | 284ms |  9.57MiB | 404596 | 595404  
  slru           |     0.40 | 406ms |  8.07MiB | 395499 | 604501  
  wtfcache       |     0.45 | 382ms | 14.71MiB | 445871 | 554129  


zipfian cache=100000 keys=10000000

      CACHE      | HIT RATE | TIME  |  MEMORY  |  HITS  | MISSES  
-----------------+----------+-------+----------+--------+---------
  tinylfu        |     0.41 | 458ms | 18.75MiB | 411355 | 588645  
  clockpro       |     0.41 | 445ms | 29.73MiB | 411368 | 588632  
  arc            |     0.41 | 938ms | 35.14MiB | 411948 | 588052  
  ristretto      |     0.40 | 424ms | 29.01MiB | 401787 | 598213  
  directcache    |     0.40 | 253ms |  6.38MiB | 400520 | 599480  
  two-queue      |     0.41 | 878ms | 27.71MiB | 411830 | 588170  
  lru-groupcache |     0.40 | 441ms | 17.90MiB | 396088 | 603912  
  lru-hashicorp  |     0.40 | 459ms | 17.89MiB | 396648 | 603352  
  s4lru          |     0.38 | 252ms |  7.63MiB | 379652 | 620348  
  slru           |     0.37 | 382ms |  6.26MiB | 374019 | 625981  
  wtfcache       |     0.40 | 362ms | 14.84MiB | 396547 | 603453  


zipfian cache=100000 keys=100000000

      CACHE      | HIT RATE | TIME  |  MEMORY  |  HITS  | MISSES  
-----------------+----------+-------+----------+--------+---------
  tinylfu        |     0.41 | 416ms | 18.69MiB | 406565 | 593435  
  clockpro       |     0.41 | 447ms | 29.69MiB | 407329 | 592671  
  arc            |     0.41 | 932ms | 35.29MiB | 408391 | 591609  
  ristretto      |     0.40 | 424ms | 29.31MiB | 397395 | 602605  
  directcache    |     0.40 | 271ms |  6.38MiB | 396138 | 603862  
  two-queue      |     0.41 | 837ms | 28.69MiB | 407455 | 592545  
  lru-groupcache |     0.39 | 448ms | 17.90MiB | 391166 | 608834  
  lru-hashicorp  |     0.39 | 466ms | 17.90MiB | 391200 | 608800  
  s4lru          |     0.38 | 262ms |  7.49MiB | 378169 | 621831  
  slru           |     0.37 | 378ms |  6.12MiB | 372687 | 627313  
  wtfcache       |     0.39 | 356ms | 14.85MiB | 391717 | 608283  


zipfian cache=1000000 keys=10000000

      CACHE      | HIT RATE | TIME  |  MEMORY   |  HITS  | MISSES  
-----------------+----------+-------+-----------+--------+---------
  tinylfu        |     0.45 | 715ms | 138.28MiB | 447743 | 552257  
  clockpro       |     0.45 | 476ms |  84.26MiB | 447792 | 552208  
  arc            |     0.45 | 764ms |  98.46MiB | 447224 | 552776  
  ristretto      |     0.41 | 446ms | 105.33MiB | 409853 | 590147  
  directcache    |     0.45 | 323ms |  40.08MiB | 447663 | 552337  
  two-queue      |     0.45 | 786ms |  98.40MiB | 447587 | 552413  
  lru-groupcache |     0.45 | 582ms |  92.71MiB | 447619 | 552381  
  lru-hashicorp  |     0.45 | 644ms |  92.67MiB | 447984 | 552016  
  s4lru          |     0.44 | 486ms |  50.13MiB | 437228 | 562772  
  slru           |     0.43 | 620ms |  41.28MiB | 431884 | 568116  
  wtfcache       |     0.45 | 503ms | 103.43MiB | 448204 | 551796  


zipfian cache=1000000 keys=100000000

      CACHE      | HIT RATE | TIME  |  MEMORY   |  HITS  | MISSES  
-----------------+----------+-------+-----------+--------+---------
  tinylfu        |     0.44 | 736ms | 139.70MiB | 435255 | 564745  
  clockpro       |     0.44 | 489ms |  85.49MiB | 435297 | 564703  
  arc            |     0.44 | 787ms |  96.10MiB | 435752 | 564248  
  ristretto      |     0.40 | 452ms | 105.25MiB | 403229 | 596771  
  directcache    |     0.43 | 318ms |  40.09MiB | 434648 | 565352  
  two-queue      |     0.43 | 742ms |  96.25MiB | 434377 | 565623  
  lru-groupcache |     0.44 | 615ms |  94.14MiB | 435035 | 564965  
  lru-hashicorp  |     0.43 | 643ms |  94.16MiB | 434901 | 565099  
  s4lru          |     0.43 | 469ms |  48.95MiB | 426663 | 573337  
  slru           |     0.42 | 583ms |  40.27MiB | 423834 | 576166  
  wtfcache       |     0.44 | 500ms | 104.45MiB | 435936 | 564064  
```

To run this benchmark:

```shell
go run *.go
```
