print("# filter-polars.py", flush=True)

import os
import gc
import timeit
import polars as pl

from statistics import mean
from helpers import write_log, memory_usage, make_chk

# other questions ans info here
# https://github.com/h2oai/db-benchmark/tree/master

ver = pl.__version__
git = ""
task = "groupby"
solution = "polars"
fun = ".groupby"
cache = "TRUE"
on_disk = "FALSE"
data_names = [
    "G1_1e4_1e2_0_0", "G1_1e5_1e2_0_0", "G1_1e6_1e2_0_0", "G1_1e7_1e2_0_0",
    "G1_1e4_1e2_10_0", "G1_1e5_1e2_10_0", "G1_1e6_1e2_10_0", "G1_1e7_1e2_10_0"
]

for data_name in data_names:
    filepath = os.path.join("..", "testdata", data_name+".csv")
    print("loading dataset %s" % data_name, flush=True)

    x = pl.read_csv(filepath)
    in_rows = x.shape[0]


    ###################     QUESTION 1   ###################

    question = "sum v1 by id1"
    gc.collect()
    times = timeit.repeat(lambda: x.groupby("id1").agg(pl.sum("v1")), repeat=5, number=1)
    t = mean(times)

    gc.collect()
    ans =  x.groupby("id1").agg(pl.sum("v1"))
    print(ans.shape, flush=True)
    m = memory_usage()
    chk = [ans["v1"].cast(pl.Int64).sum()]

    write_log(task=task, data=data_name, in_rows=in_rows, question=question, out_rows=ans.shape[0], out_cols=ans.shape[1], solution=solution, version=ver, git=git, fun=fun, run=1, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=-1, on_disk=on_disk)
    del ans


    ###################     QUESTION 2   ###################

    question = "sum v1 by id1:id2"
    gc.collect()
    times = timeit.repeat(lambda: x.groupby(["id1","id2"]).agg(pl.sum("v1")), repeat=5, number=1)
    t = mean(times)

    gc.collect()
    ans =  x.groupby(["id1","id2"]).agg(pl.sum("v1"))
    print(ans.shape, flush=True)
    m = memory_usage()
    chk = [ans["v1"].cast(pl.Int64).sum()]

    write_log(task=task, data=data_name, in_rows=in_rows, question=question, out_rows=ans.shape[0], out_cols=ans.shape[1], solution=solution, version=ver, git=git, fun=fun, run=1, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=-1, on_disk=on_disk)
    del ans


    ###################     QUESTION 3   ###################

    question = "sum v1 mean v3 by id3"
    gc.collect()
    times = timeit.repeat(lambda: x.groupby("id3").agg([pl.sum("v1"), pl.mean("v3")]), repeat=5, number=1)
    t = mean(times)

    gc.collect()
    ans =  x.groupby("id3").agg([pl.sum("v1"), pl.mean("v3")])
    print(ans.shape, flush=True)
    m = memory_usage()
    chk = ans.select([pl.col("v1").cast(pl.Int64).sum(), pl.col("v3").sum()]).to_numpy().tolist()

    write_log(task=task, data=data_name, in_rows=in_rows, question=question, out_rows=ans.shape[0], out_cols=ans.shape[1], solution=solution, version=ver, git=git, fun=fun, run=1, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=-1, on_disk=on_disk)
    del ans


    ###################     QUESTION 4   ###################

    question = "mean v1:v3 by id4"
    gc.collect()
    times = timeit.repeat(lambda: x.groupby("id4").agg([pl.mean("v1"), pl.mean("v2"), pl.mean("v3")]), repeat=5, number=1)
    t = mean(times)

    gc.collect()
    ans =  x.groupby("id4").agg([pl.mean("v1"), pl.mean("v2"), pl.mean("v3")])
    print(ans.shape, flush=True)
    m = memory_usage()
    chk = ans.select([pl.col("v1").sum(), pl.col("v2").sum(), pl.col("v3").sum()]).to_numpy().tolist()

    write_log(task=task, data=data_name, in_rows=in_rows, question=question, out_rows=ans.shape[0], out_cols=ans.shape[1], solution=solution, version=ver, git=git, fun=fun, run=1, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=-1, on_disk=on_disk)
    del ans


    ###################     QUESTION 5   ###################

    question = "sum v1:v3 by id6"
    gc.collect()
    times = timeit.repeat(lambda: x.groupby("id6").agg([pl.sum("v1"), pl.sum("v2"), pl.sum("v3")]), repeat=5, number=1)
    t = mean(times)

    gc.collect()
    ans =  x.groupby("id6").agg([pl.sum("v1"), pl.sum("v2"), pl.sum("v3")])
    print(ans.shape, flush=True)
    m = memory_usage()
    chk = ans.select([pl.col("v1").cast(pl.Int64).sum(), pl.col("v2").cast(pl.Int64).sum(), pl.col("v3").sum()]).to_numpy().tolist()

    write_log(task=task, data=data_name, in_rows=in_rows, question=question, out_rows=ans.shape[0], out_cols=ans.shape[1], solution=solution, version=ver, git=git, fun=fun, run=1, time_sec=t, mem_gb=m, cache=cache, chk=make_chk(chk), chk_time_sec=-1, on_disk=on_disk)
    del ans
